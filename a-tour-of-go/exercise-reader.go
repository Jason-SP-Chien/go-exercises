package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(a []byte) (n int, e error) {
	for {
		n = copy(a, []byte{'A'})
		return
	}
}

func main() {
	reader.Validate(MyReader{})
}
