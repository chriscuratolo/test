package foo

func Boo(x int) *int {
	if x+1 != 100 {
		return nil
	}
	return &x
}
