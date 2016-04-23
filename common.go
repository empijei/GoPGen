package gopgen

type Drawable interface {
	Draw(w Writer, width int, height int)
}

type Writer interface {
	Write([]byte) (int, error)
}
