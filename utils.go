package gofb

func swapFloat32(x, y *float32) {
	*x, *y = *y, *x
}
