package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Edge struct holds node links
type Edge struct {
	weight int
	end    Node
}

// Node struct holds edges and the hash maps key for convinence
type Node struct {
	ID     byte
	weight int
	edges  []Edge
}

// Graph type
type Graph struct {
	nodes map[byte]Node
}

// NewGraph creates a graph object and inits some stuff
func NewGraph() *Graph {
	g := Graph{nodes: make(map[byte]Node)}
	return &g
}

// LoadFromFile graph def from CSV formated file, treats rows and cols as a single 1D map
func (g *Graph) LoadFromFile(fname string) error {
	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		for _, value := range record {
			// Sanity
			if len(value) != 3 {
				return errors.New("Input file invalid")
			}

			// Grab node ids
			a := []byte(value)[0]
			b := []byte(value)[1]

			// Grab distance weight
			w, err := strconv.Atoi(string(value[len(value)-1]))
			if err != nil {
				return err
			}

			// If nodes don't exist make sure they do
			aNode, ok := g.nodes[a]
			if !ok {
				g.nodes[a] = Node{ID: a, weight: w}
			}
			_, ok = g.nodes[b]
			if !ok {
				g.nodes[b] = Node{ID: b, weight: w}
			}

			// Populate edges
			aNode.edges = append(g.nodes[a].edges, Edge{end: g.nodes[b], weight: w})

			// Push back changes
			g.nodes[a] = aNode
		}
	}
	return nil
}

// ABDistance test
func (g *Graph) ABDistance(a, b byte) int {
	// Check for nil map(runtime errors are bad)
	if aNode, ok := g.nodes[a]; ok {
		for _, edge := range aNode.edges {
			// Detect nil end
			if edge.end.ID == 0 {
				continue
			}

			// On match return distance
			if edge.end.ID == b {
				return edge.weight
			}
		}
	}
	return -1
}

// Distance result
func (g *Graph) Distance(stops ...byte) string {
	// Distance result
	var res int

	// Loop over variadic node list
	for i, stop := range stops[:len(stops)-1] {
		// Check if route is possible, then accumulate distance
		if t := g.ABDistance(stop, stops[i+1]); t != -1 {
			res += t
		} else {
			fmt.Println("NO SUCH ROUTE")
			return "NO SUCH ROUTE"
		}
	}

	fmt.Println(res)
	return strconv.Itoa(res)
}

// TripsCheck test
func (g *Graph) TripsCheck(a, b byte, maxhops int) (bool, int) {
	// Check for nil map(runtime errors are bad)
	if aNode, ok := g.nodes[a]; ok && maxhops > 0 {
		// Do a quick edge check
		for _, edge := range aNode.edges {
			// Detect nil end
			if edge.end.ID == 0 {
				continue
			}

			// Every populated edge is a possible path
			if a == b {
				return true, maxhops - 1
			}
		}

		// Recursive check
		for _, edge := range aNode.edges {
			// Detect nil end
			if edge.end.ID == 0 {
				continue
			}

			return g.TripsCheck(edge.end.ID, b, maxhops-1)
		}
	}
	return false, 0
}

// Trips result
func (g *Graph) Trips(a, b byte, minhops, maxhops int) string {
	// Number of trips
	res := 0

	// Check for nil map(runtime errors are bad)
	if aNode, ok := g.nodes[a]; ok {
		for _, edge := range aNode.edges {
			// Detect nil end
			if edge.end.ID == 0 {
				continue
			}

			// Check if path was resolved and remaining hops fits within
			// the requested paramaters
			if ok, remHops := g.TripsCheck(edge.end.ID, b, maxhops); ok {
				if minhops > 0 && remHops < (maxhops-minhops) {
					continue
				}
				res++
			}
		}
	}

	fmt.Println(res)
	return strconv.Itoa(res)
}

// RouteLengthCheck test
func (g *Graph) RouteLengthCheck(a, b byte, maxhops int) int {
	// Don't traverse ourself, just return the distance of last the node
	if a == b {
		return g.nodes[b].weight
	}

	// Check for nil map(runtime errors are bad)
	// Make sure we haven't exceeded maxhops
	if aNode, ok := g.nodes[a]; ok && maxhops > 0 {
		//fmt.Print("->", string(a), "->", string(b))
		for _, edge := range aNode.edges {
			// Detect nil end
			if edge.end.ID == 0 {
				continue
			}

			// Recursively resolve path accumulating distance
			r := g.RouteLengthCheck(edge.end.ID, b, maxhops-1)

			// A negitive value means the path could not be resolved
			if r == -1 {
				continue
			}

			return r
		}
	}
	return -1
}

// ShortestRouteLength result
func (g *Graph) ShortestRouteLength(a, b byte, maxhops int) string {
	// Route paths
	res := []int{}

	// Check for nil map(runtime errors are bad)
	if aNode, ok := g.nodes[a]; ok {
		for _, edge := range aNode.edges {
			// Detect nil end
			if edge.end.ID == 0 {
				continue
			}

			// Recursively resolve path accumulating distance
			r := g.RouteLengthCheck(edge.end.ID, b, maxhops-1)
			if r == -1 {
				continue
			}

			// Append path distance
			res = append(res, r+edge.weight)
		}
	}

	// Get smallest value
	// (could be done with search built-in but this makes less sense)
	min := 1000
	for _, v := range res {
		if v < min {
			min = v
		}
	}

	fmt.Println(min)
	return strconv.Itoa(min)
}

// PossibleRoutesCheck test
func (g *Graph) PossibleRoutesCheck(a, b byte, maxhops int, start byte, route *int) {
	// Increment previous
	if start == b {
		*route++
	}
	if aNode, ok := g.nodes[a]; ok && maxhops > 0 {
		// Traverse edges
		for _, edge := range aNode.edges {
			// Detect nil end
			if edge.end.ID == 0 {
				continue
			}

			// Increment routes if not at end of chain
			*route++

			// Continue search
			g.RouteLengthCheck(edge.end.ID, b, maxhops-1)
		}
	}
}

// PossibleRoutes result
func (g *Graph) PossibleRoutes(a, b byte, maxhops int) string {
	res := 0
	if aNode, ok := g.nodes[a]; ok && maxhops > 0 {
		// Traverse edges
		for _, edge := range aNode.edges {
			// Detect nil end
			if edge.end.ID == 0 {
				continue
			}

			// Every non-nil path, thats a route
			if a == b {
				res++
			}

			// Begin recursive search up to maxhops
			g.PossibleRoutesCheck(edge.end.ID, b, maxhops-1, a, &res)
		}
	}

	fmt.Println(res)
	return strconv.Itoa(res)
}

func main() {

}
