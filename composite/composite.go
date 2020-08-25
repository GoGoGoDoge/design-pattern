package composite

import (
	"bytes"
)

type filenode interface {
	Print(sep string) string
}

type Dir struct {
	name     string
	children []filenode
}

func NewDir(name string) *Dir {
	return &Dir{name: name}
}

func (d *Dir) Add(node filenode) {
	d.children = append(d.children, node)
}

func (d *Dir) Print(sep string) string {
	var buf bytes.Buffer
	buf.WriteString(d.name)
	sep += sep
	for _, node := range d.children {
		// we can make use of the print here
		buf.WriteString("\n")
		buf.WriteString(sep + "|\n")
		buf.WriteString(sep + "->")
		buf.WriteString(node.Print(sep))
	}
	return buf.String()
}

type File struct {
	name string
}

func NewFile(name string) *File {
	return &File{
		name: name,
	}
}

func (f *File) Print(sep string) string {
	return f.name
}
