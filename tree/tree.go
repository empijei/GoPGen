package gopgen

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
