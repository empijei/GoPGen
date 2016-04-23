package gopgen

import (
	"math"
	"math/rand"
)

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
	angle := bf.Spanning / bf.BranchesCount
	for i := 0; i < int(bf.BranchesCount); i++ {
		curAngle := angle*float64(i) + rand.Float64()*angle //TODO control this more?
		deltax := bf.Length * math.Cos(curAngle)
		deltay := bf.Length * math.Sin(curAngle)
		children = append(children, TreeNode{X: this.X + deltax, Y: this.Y + deltay})
	}
	return
}
