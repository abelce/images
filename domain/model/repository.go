package model

type Repository interface {
	Save(a *Article) error
	NewIdentity() string
	UpdateByID(id string, article *Article)(*Article, error)
}