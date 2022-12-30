package usecase

import "testing"

func TestPrototype(t *testing.T) {
	file1 := NewFile("file1")
	file2 := NewFile("file2")
	file3 := NewFile("file3")

	children := []Inode{file1, file2, file3}

	folder := NewFolder("folder1", children)
	copyFolder := folder.Clone()

	if folder.GetName() != copyFolder.GetName() {
		t.Error("Objects are different")
	}
}
