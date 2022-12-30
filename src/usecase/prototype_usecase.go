package usecase

type Inode interface {
	GetName() string
	SetName(string)
	Clone() Inode
}

// File concrete prototype
type File struct {
	name string
}

func NewFile(name string) *File {
	return &File{
		name: name,
	}
}

func (file *File) GetName() string {
	return file.name
}

func (file *File) SetName(name string) {
	file.name = name
}

func (file *File) Clone() Inode {
	return &File{name: file.name}
}

// Folder concrete prototype
type Folder struct {
	name     string
	children []Inode
}

func NewFolder(name string, children []Inode) *Folder {
	return &Folder{
		name:     name,
		children: children,
	}
}

func (folder *Folder) GetName() string {
	return folder.name
}

func (folder *Folder) SetName(name string) {
	folder.name = name
}

func (folder *Folder) Clone() Inode {
	folderClone := &Folder{name: folder.name}

	var tempChildren []Inode
	for _, child := range folder.children {
		childClone := child.Clone()
		tempChildren = append(tempChildren, childClone)
	}

	folderClone.children = tempChildren
	return folderClone
}
