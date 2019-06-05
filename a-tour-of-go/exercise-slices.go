package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		res[i] = make([]uint8, dx)
	}
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			res[i][j] = tf(i, j)
		}
	}
	return res
}
func tf(x, y int) uint8 {
	return uint8(x ^ y)
}
func main() {
	pic.Show(Pic)
}
