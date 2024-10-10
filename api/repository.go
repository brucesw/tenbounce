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
	GetPoint(pointID string) (model.Point, error)
	ListPoints(userID string) ([]model.Point, error)
	CreatePoint(p *model.Point) error
	DeletePoint(pointID string) error
}
type PointTypeRepository interface {
	ListPointTypes() ([]model.PointType, error)
	CreatePointType(p *model.PointType) error
}
