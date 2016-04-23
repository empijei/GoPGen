package gopgen

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"testing"
)

var drawTests = []struct {
	bf   branchFactor
	this TreeNode
}{
	{branchFactor{4.0, 1.0 * math.Pi, 100.0}, TreeNode{}},
}

func TestDraw(t *testing.T) {
	for _, _ = range drawTests {
		node := TreeNode{}
		testTree := Tree{liveTreeNodes: []*TreeNode{&node}}
		testTree.Root = &node
		//t.Logf("%#v", children)
		//tree := Tree(children)
		t.Logf("%#v", testTree)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/svg+xml")
			testTree.Grow(WeatherCondition{})
			fmt.Printf("%#v\n", testTree)
			testTree.Draw(w, 400, 400)
			testTree.Draw(os.Stdout, 400, 400)
		})
		_ = http.ListenAndServe("localhost:8000", nil)
	}
}
