package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"tenbounce/model"

	_ "github.com/lib/pq"
)

type Postgres struct {
	dataSourceName string
	db             *sql.DB
}

func NewPostgresRepository(dataSourceName string) *Postgres {
	return &Postgres{
		dataSourceName: dataSourceName,
	}
}

// lazyPostgresDB returns the *sql.DB if it already exists,
// otherwise instantiates one, attaches it to the Postgres Repository
// then returns it.
func (r *Postgres) lazyPostgresDB() (*sql.DB, error) {
	if r.db != nil {
		return r.db, nil
	}

	if db, err := sql.Open("postgres", r.dataSourceName); err != nil {
		return nil, fmt.Errorf("sql open postgres: %w", err)
	} else {
		r.db = db
	}

	return r.db, nil
}

func (r *Postgres) GetUser(userID string) (model.User, error) {
	db, err := r.lazyPostgresDB()
	if err != nil {
		return model.User{}, fmt.Errorf("lazy postgres db: %w", err)
	}

	var user = model.User{}

	var row = db.QueryRow("SELECT * FROM users WHERE id = $1", userID)
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, errors.New("no user found")
		}

		return model.User{}, fmt.Errorf("scan row: %w", err)
	}

	return user, nil
}

func (r *Postgres) ListPoints(userID string) ([]model.Point, error) {
	db, err := r.lazyPostgresDB()
	if err != nil {
		return nil, fmt.Errorf("lazy postgres db: %w", err)
	}

	var points = []model.Point{}

	rows, err := db.Query("SELECT * FROM points WHERE user_id = $1", userID)
	if err != nil {
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

func (r *Postgres) CreatePoint(p *model.Point) error {
	db, err := r.lazyPostgresDB()
	if err != nil {
		return fmt.Errorf("lazy postgres db: %w", err)
	}

	_, err = db.Exec("INSERT INTO points (timestamp, user_id, point_type_id, value) VALUES (CURRENT_TIMESTAMP, $1, $2, $3)", p.UserID, p.PointTypeID, p.Value)
	if err != nil {
		return fmt.Errorf("exec db: %w", err)
	}

	return nil
}

func (r *Postgres) ListPointTypes() ([]model.PointType, error) {
	db, err := r.lazyPostgresDB()
	if err != nil {
		return nil, fmt.Errorf("lazy postgres db: %w", err)
	}

	var pointTypes = []model.PointType{}

	rows, err := db.Query("SELECT * FROM point_types")
	if err != nil {
		return nil, fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var pointType = model.PointType{}

		if err := rows.Scan(&pointType.ID, &pointType.Name); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		pointTypes = append(pointTypes, pointType)
	}

	return pointTypes, nil
}
