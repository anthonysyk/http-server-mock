package stdout

import (
	"io"
	"os"
)

func Record(fn func()) string {
	old := os.Stdout
	// start recording stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()

	// end recording
	_ = w.Close()
	result, _ := io.ReadAll(r)
	os.Stdout = old
	return string(result)
}
