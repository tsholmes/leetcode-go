package main

import (
	"path"
	"sort"
	"strings"
)

func main() {

}

type fnode struct {
	files map[string][]byte
	dirs  map[string]*fnode
}

type FileSystem struct {
	root *fnode
}

func Constructor() FileSystem {
	return FileSystem{
		root: &fnode{
			files: map[string][]byte{},
			dirs:  map[string]*fnode{},
		},
	}
}

func (fs *FileSystem) node(path string) *fnode {
	if path == "/" {
		return fs.root
	}
	path = strings.TrimSuffix(path, "/")
	ps := strings.Split(path[1:], "/")
	f := fs.root
	for _, p := range ps {
		if _, ok := f.dirs[p]; !ok {
			f.dirs[p] = &fnode{
				files: map[string][]byte{},
				dirs:  map[string]*fnode{},
			}
		}
		f = f.dirs[p]
	}
	return f
}

func (fs *FileSystem) Ls(fpath string) []string {
	if fpath != "/" {
		fp := fs.node(path.Dir(fpath))
		if _, ok := fp.files[path.Base(fpath)]; ok {
			return []string{path.Base(fpath)}
		}
	}

	f := fs.node(fpath)
	var res []string
	for k := range f.dirs {
		res = append(res, k)
	}
	for k := range f.files {
		res = append(res, k)
	}
	sort.Strings(res)
	return res
}

func (fs *FileSystem) Mkdir(path string) {
	fs.node(path)
}

func (fs *FileSystem) AddContentToFile(filePath string, content string) {
	f := fs.node(path.Dir(filePath))
	fname := path.Base(filePath)
	f.files[fname] = append(f.files[fname], content...)
}

func (fs *FileSystem) ReadContentFromFile(filePath string) string {
	f := fs.node(path.Dir(filePath))
	return string(f.files[path.Base(filePath)])
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ls(path);
 * obj.Mkdir(path);
 * obj.AddContentToFile(filePath,content);
 * param_4 := obj.ReadContentFromFile(filePath);
 */
