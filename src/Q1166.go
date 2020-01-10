package src

import "strings"

type FileSystem struct {
	path map[string]int
}

func FileSystemConstructor() FileSystem {
	return FileSystem{
		path: make(map[string]int),
	}
}

func (this *FileSystem) Create(path string, value int) bool {
	//judge whether it is a primary directory
	last := strings.LastIndex(path, "/")
	if last != 0 {
		parent := path[0:last]
		if _, ok := this.path[parent]; !ok {
			return false
		}
	}

	this.path[path] = value
	return true
}

func (this *FileSystem) Get(path string) int {
	if value, ok := this.path[path]; !ok {
		return -1
	} else {
		return value
	}
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Create(path,value);
 * param_2 := obj.Get(path);
 */
