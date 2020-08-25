package prototype_test

import (
	"prototype"
	"testing"
)

var expectPrint1 string = `  Dir1
    File1
    File2
    File3
`

var expectPrint2 string = `  Dir1_clone
    File1_clone
    File2_clone
    File3_clone
`

func TestPrototype(t *testing.T) {
	file1 := &prototype.File{Name: "File1"}
	file2 := &prototype.File{Name: "File2"}
	file3 := &prototype.File{Name: "File3"}
	dir1 := &prototype.Dir{Name: "Dir1", Nodes: []prototype.Inode{file1, file2, file3}}
	if dir1.Print("  ") != expectPrint1 {
		t.Fatal("Got:\n" + dir1.Print("  ") + "But expect:" + expectPrint1 + "#")
	}

	cloneDir := dir1.Clone()
	if cloneDir.Print("  ") != expectPrint2 {
		t.Fatal("Got:\n" + cloneDir.Print("  ") + "But expect:" + expectPrint2 + "#")
	}
}
