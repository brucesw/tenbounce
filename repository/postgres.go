package repository

import (
	"database/sql"
	"fmt"
	"tenbounce/model"

	_ "github.com/lib/pq"
)

type Postgres struct {
	dataSourceName string
}

func NewPostgresRepository(dataSourceName string) *Postgres {
	return &Postgres{
		dataSourceName: dataSourceName,
	}
}

func (r *Postgres) GetUser(userID string) (model.User, error) {
	return model.User{}, fmt.Errorf("user '%s' not found", userID)
}

func (r *Postgres) ListPoints(userID string) ([]model.Point, error) {
	db, err := sql.Open("postgres", r.dataSourceName)
	if err != nil {
		return nil, err
	}

	var points = []model.Point{}

	rows, err := db.Query("SELECT * FROM points WHERE user_id = $1", userID)
	if err != nil {
		fmt.Println(fmt.Errorf("db query: %w", err))
		return nil, fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var point model.Point

		if err := rows.Scan(&point.ID, &point.Timestamp, &point.UserID, &point.PointTypeID, &point.Value); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		points = append(points, point)
	}

	return points, nil
}

func (r *Postgres) SavePoint(p *model.Point) error {
	return nil
}

func (r *Postgres) ListPointTypes() ([]model.PointType, error) {
	return []model.PointType{}, nil
}
