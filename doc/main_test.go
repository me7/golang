package docs

import "testing"
import "fmt"

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Errorf("fail add test")
	}
}

func ExampleAdd() {
	fmt.Println(Add(2, 2))
	// Output: 4
}
