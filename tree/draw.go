package gopgen

import (
	"fmt"
	"io"
)

const svgHeader = "<svg xmlns='http://www.w3.org/2000/svg' " +
	"style='stroke: grey; fill: white; stroke-width: 0.7' " +
	"width='%d' height='%d'>"
const line = "<line x1='%f' y1='%f' x2='%f' y2='%f' style='stroke:rgb(0,255,0);stroke-width:2' />"
const svgTrailer = "</svg>"

func (this *Tree) Draw(w io.Writer, width, height int) {
	if _, err := w.Write([]byte(fmt.Sprintf(svgHeader, width, height))); err == nil {
		drawBranch(this.Root, w, width, height)
	} else {
		fmt.Println(err)
		return
	}
	_, _ = w.Write([]byte(svgTrailer))
}

func drawBranch(node *TreeNode, w io.Writer, width, height int) {
	node = svgAdapter(node, width, height)
	for _, child := range node.Children {
		adaptedChild := svgAdapter(child, width, height)
		_, _ = w.Write([]byte(fmt.Sprintf(line, node.X, node.Y, adaptedChild.X, adaptedChild.Y)))
		drawBranch(child, w, width, height)
	}
}
func svgAdapter(n *TreeNode, width, height int) *TreeNode {
	b := &TreeNode{}
	b.X = float64(width)/2.0 + n.X
	b.Y = float64(height) - n.Y
	b.Children = n.Children
	return b
}
