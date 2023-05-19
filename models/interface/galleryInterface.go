package _interface

import "digibala/models"

type GalleryInterface interface {
	CreateGallery() models.Gallery
	UpdateGallery() bool
	deleteGallery() bool
}
