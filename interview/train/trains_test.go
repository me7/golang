package main

import "testing"

type test struct {
	res      string
	function func(...byte) string
}

var ans = []string{"9", "5", "13", "22", "NO SUCH ROUTE", "2", "3", "9", "9", "7"}

func TestTrains(t *testing.T) {
	// Init our node graph
	g := NewGraph()

	// Populate from file
	g.LoadFromFile("graph.db")

	if r := g.Distance('A', 'B', 'C'); r != ans[0] {
		t.Error("Expected", ans[0], "got", r)
	}
	if r := g.Distance('A', 'D'); r != ans[1] {
		t.Error("Expected", ans[1], "got", r)
	}
	if r := g.Distance('A', 'D', 'C'); r != ans[2] {
		t.Error("Expected", ans[2], "got", r)
	}
	if r := g.Distance('A', 'E', 'B', 'C', 'D'); r != ans[3] {
		t.Error("Expected", ans[3], "got", r)
	}
	if r := g.Distance('A', 'E', 'D'); r != ans[4] {
		t.Error("Expected", ans[4], "got", r)
	}
	if r := g.Trips('C', 'C', 0, 3); r != ans[5] {
		t.Error("Expected", ans[5], "got", r)
	}
	if r := g.Trips('A', 'C', 4, 4); r != ans[6] {
		t.Error("Expected", ans[6], "got", r)
	}
	if r := g.ShortestRouteLength('A', 'C', 10); r != ans[7] {
		t.Error("Expected", ans[7], "got", r)
	}
	if r := g.ShortestRouteLength('B', 'B', 10); r != ans[8] {
		t.Error("Expected", ans[8], "got", r)
	}
	if r := g.PossibleRoutes('C', 'C', 30); r != ans[9] {
		t.Error("Expected", ans[9], "got", r)
	}
}
