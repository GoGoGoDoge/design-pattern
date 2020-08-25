package prototype

import (
	"bytes"
	"fmt"
)

type Inode interface {
	Print(string) string
	Clone() Inode // 这个就是关键
}

type File struct {
	Name string
}

func (f *File) Print(sep string) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln(sep + f.Name))
	return buf.String()
}

// Clone should be light weight using shallow copy
func (f *File) Clone() Inode {
	return &File{Name: f.Name + "_clone"}
}

type Dir struct {
	Nodes []Inode
	Name  string
}

func (d *Dir) Print(sep string) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln(sep + d.Name))
	for _, i := range d.Nodes {
		buf.WriteString(i.Print(sep + sep))
	}
	return buf.String()
}

// Clone should be light weight using shallow copy
func (d *Dir) Clone() Inode {
	cloneDir := &Dir{Name: d.Name + "_clone"}
	var tempNodes []Inode
	for _, i := range d.Nodes {
		tempNodes = append(tempNodes, i.Clone())
	}
	cloneDir.Nodes = tempNodes
	return cloneDir
}
