package procgen

import "math"
import "math/rand"

type Tree struct {
	Branches      []TreeNode
	liveTreeNodes []TreeNode
}

type TreeNode struct {
	X, Y, Z  float64
	Children []TreeNode
}

type branchFactor struct {
	BranchesCount float64
	Spanning      float64
	Length        float64
}

type WeatherCondition struct {
	heat     float64
	humidity float64
}

func (this *Tree) Grow(wc WeatherCondition) {
	//TODO compute branch factor
	//
	var newNodes []TreeNode
	for _, treenode := range this.liveTreeNodes {
		newNodes = append(newNodes, treenode.grow(computeBranchFactor(wc))...)
	}
	this.liveTreeNodes = newNodes
}

func computeBranchFactor(wc WeatherCondition) (bf branchFactor) {
	bf.BranchesCount = 4.0 //FIXME some random rounding here
	bf.Spanning = 1.0 * math.Pi
	bf.Length = 1.0
	return
}

func (this *TreeNode) grow(bf branchFactor) (children []TreeNode) {
	angle := bf.Spanning / bf.BranchesCount //FIXME
	for i := 0; i < int(bf.BranchesCount); i++ {
		curAngle := angle*i + rand.Float64*angle
		//TODO compute where the new node is
	}
	return
}
