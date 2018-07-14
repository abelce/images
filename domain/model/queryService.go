package model

type QueryService interface {
	Find(offsetNum, limit int) (int, []*Article, error)
	FindByID(id string) (*Article, error)
}