package src

import (
	"testing"
)

func TestFileSystem(t *testing.T) {
	f := FileSystemConstructor()
	f.Create("/leet", 1)
	f.Create("/leet/code", 2)
	f.Get("/leet/code")
	f.Create("/c/d", 3)
	f.Get("/c")
}