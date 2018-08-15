package model

type Repository interface {
	Save(a *Image) error
	NewIdentity() string
	// UpdateByID(id string, article *Image)(*Image, error)
}