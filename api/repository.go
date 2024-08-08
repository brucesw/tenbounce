package api

import "tenbounce/model"

type Repository interface {
	UserRepository
	PointRepository
	PointTypeRepository
}

type UserRepository interface {
	GetUser(userID string) (model.User, error)
	ListUsers() ([]model.User, error)
}
type PointRepository interface {
	ListPoints(userID string) ([]model.Point, error)
	CreatePoint(p *model.Point) error
}
type PointTypeRepository interface {
	ListPointTypes() ([]model.PointType, error)
}
