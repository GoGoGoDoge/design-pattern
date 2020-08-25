package composite_test

import (
	"composite"
	"testing"
)

func TestComposite(t *testing.T) {
	dir1 := composite.NewDir("dir1")
	dir1.Add(composite.NewFile("file1"))
	dir2 := composite.NewDir("dir2")
	dir2.Add(composite.NewFile("file2"))
	dir1.Add(dir2)

	expectResult := `dir1
##|
##->file1
##|
##->dir2
####|
####->file2`

	if dir1.Print("#") != expectResult {
		t.Fatal(dir1.Print("#"))
	}
}
