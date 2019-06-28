package src

import (
	"path/filepath"
	"strings"
)

func simplifyPath(path string) string {
	if path == "/"{
		return path
	}

	stringSlice := strings.Split(path, "/")
	//remove last /

	var pathStack  []string
	for _,v :=range stringSlice{
		switch v {
			case "":
			case ".":
				break
			default:
				pathStack = append(pathStack, "/")
				pathStack = append(pathStack, v)
			case "..":
				if pathLength := len(pathStack); pathLength>2{
					pathStack = pathStack[:pathLength - 2]
				}else{
					pathStack = []string{}
				}
		}
	}
	if len(pathStack) == 0 {
		return "/"
	}

	result := ""
	for _,v :=range  pathStack{
		result+=v
	}
	return result
}

func simplifyPath2(path string) string {
	return filepath.Clean(path);
}

func simplifyPath3(path string) string {
	dirs := strings.Split(path, "/")
	dirs = dirs[1:]

	i := 0
	for i < len(dirs) {
		if dirs[i] == "" || dirs[i] == "." {
			dirs = append(dirs[:i], dirs[i+1:]...)
		} else if dirs[i] == ".." {
			if i == 0 {
				dirs = append(dirs[:i], dirs[i+1:]...)
			} else {
				dirs = append(dirs[:i-1], dirs[i+1:]...)
				i--
			}
		} else {
			i++
		}
	}

	return "/" + strings.Join(dirs, "/")
}