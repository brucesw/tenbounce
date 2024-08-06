package repository

import (
	"fmt"
	"tenbounce/model"
)

type Postgres struct {
}

func NewPostgresRepository() *Postgres {
	return &Postgres{}
}

func (r *Postgres) GetUser(userID string) (model.User, error) {
	return model.User{}, fmt.Errorf("user '%s' not found", userID)
}

func (r *Postgres) ListPoints(userID string) ([]model.Point, error) {
	return []model.Point{}, nil
}

func (r *Postgres) SavePoint(p *model.Point) error {
	return nil
}

func (r *Postgres) ListPointTypes() ([]model.PointType, error) {
	return []model.PointType{}, nil
}
