package gopgen

import (
	"math"
	"testing"
)

var growTests = []struct {
	bf   branchFactor
	this TreeNode
}{
	{branchFactor{4.0, 1.0 * math.Pi, 1.0}, TreeNode{}},
}

func TestGrow(t *testing.T) {
	for _, tt := range growTests {
		t.Logf("%#v", tt.this.grow(tt.bf))
	}
}
