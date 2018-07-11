package model

// import (
// 	"admin/domain/model"
// )

type Repository interface {
	Save(a *Article) (Article, error)
}