package api

import "tenbounce/model"

type Repository interface {
	GetUser(userID string) (model.User, error)
	ListPoints(userID string) ([]model.Point, error)
	SavePoint(p *model.Point) error
	ListPointTypes() ([]model.PointType, error)
}

// TODO(bruce): Interface embedding
