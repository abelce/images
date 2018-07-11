package model

type QueryService interface {
	FindByID(id string) (*Article, error)
}