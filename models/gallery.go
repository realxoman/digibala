package models

type Gallery struct {
	ID        int64
	Name      string
	TypeFile  string
	Size      string
	Direction string
	Hidden    bool
}

func NewGallery() *Gallery {
	gallery = new(Gallery)
	return &Gallery{}
}

// TODO Create
func Create() bool {
	gallery = NewGallery()
	//Todo to create
	return true
}

// TODO update
func Update() {

}

// TODO Delete
func delete() {

}
