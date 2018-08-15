package model

type QueryService interface {
	Find(offsetNum, limit int) (int, []*Image, error)
	// FindByID(id string) (*Image, error)
}