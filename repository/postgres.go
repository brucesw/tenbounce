package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"tenbounce/model"
	"time"

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
			return model.User{}, errors.New("user not found")
		}

		return model.User{}, fmt.Errorf("scan row: %w", err)
	}

	return user, nil
}

func (r *Postgres) ListUsers() ([]model.User, error) {
	db, err := r.lazyPostgresDB()
	if err != nil {
		return []model.User{}, fmt.Errorf("lazy postgres db: %w", err)
	}

	var users = []model.User{}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User

		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *Postgres) GetPoint(pointID string) (model.Point, error) {
	db, err := r.lazyPostgresDB()
	if err != nil {
		return model.Point{}, fmt.Errorf("lazy postgres db: %w", err)
	}

	var point = model.Point{}

	var row = db.QueryRow("SELECT * FROM points WHERE id = $1", pointID)
	if err := row.Scan(&point.ID, &point.Timestamp, &point.UserID, &point.PointTypeID, &point.Value, &point.CreatedByUserID); err != nil {
		if err == sql.ErrNoRows {
			return model.Point{}, errors.New("point not found")
		}

		return model.Point{}, fmt.Errorf("scan row: %w", err)
	}

	return point, nil
}

func (r *Postgres) ListPoints(userID string) ([]model.Point, error) {
	db, err := r.lazyPostgresDB()
	if err != nil {
		return nil, fmt.Errorf("lazy postgres db: %w", err)
	}

	var points = []model.Point{}

	rows, err := db.Query("SELECT * FROM points WHERE user_id = $1", userID)
	if err != nil {
		return nil, fmt.Errorf("db query: %w", fmt.Errorf("db query: %w", err))
	}
	defer rows.Close()

	for rows.Next() {
		var point model.Point

		if err := rows.Scan(&point.ID, &point.Timestamp, &point.UserID, &point.PointTypeID, &point.Value, &point.CreatedByUserID); err != nil {
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

	var row = db.QueryRow("INSERT INTO points (timestamp, user_id, point_type_id, value, created_by) VALUES ($1, $2, $3, $4, $5) RETURNING id", p.Timestamp, p.UserID, p.PointTypeID, p.Value, p.CreatedByUserID)

	if err = row.Scan(&p.ID); err != nil {
		return fmt.Errorf("scan row for id: %w", err)
	}

	return nil
}

func (r *Postgres) DeletePoint(pointID string) error {
	db, err := r.lazyPostgresDB()
	if err != nil {
		return fmt.Errorf("lazy postgres db: %w", err)
	}

	_, err = db.Exec("DELETE FROM points WHERE id = $1", pointID)
	if err != nil {
		return fmt.Errorf("db exec: %w", err)
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

func (r *Postgres) CreatePointType(p *model.PointType) error {
	db, err := r.lazyPostgresDB()
	if err != nil {
		return fmt.Errorf("lazy postgres db: %w", err)
	}

	var row = db.QueryRow("INSERT INTO point_types (name) VALUES ($1) RETURNING id", p.Name)

	if err = row.Scan(&p.ID); err != nil {
		return fmt.Errorf("scan row for id: %w", err)
	}

	return nil
}

func (r *Postgres) GetStatsSummary() ([]model.StatsSummary, error) {
	db, err := r.lazyPostgresDB()
	if err != nil {
		return nil, fmt.Errorf("lazy postgres db: %w", err)
	}

	type row struct {
		userID        string
		userName      string
		timestamp     time.Time
		pointTypeID   model.PointTypeID
		pointValue    model.PointValue
		pointTypeName model.PointTypeName
	}

	var rows []row

	queryRows, err := db.Query(`
		SELECT
			users.id,
			users.name,
			points.timestamp,
			points.point_type_id,
			points.value,
			point_types.name
		FROM points
		LEFT JOIN users ON points.user_id = users.id
		LEFT JOIN point_types ON points.point_type_id = point_types.id
	`)
	if err != nil {
		return nil, fmt.Errorf("db query: %w", err)
	}
	defer queryRows.Close()

	for queryRows.Next() {
		var r row

		if err := queryRows.Scan(
			&r.userID,
			&r.userName,
			&r.timestamp,
			&r.pointTypeID,
			&r.pointValue,
			&r.pointTypeName,
		); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		rows = append(rows, r)
	}

	var statsSummaries = []model.StatsSummary{}

	// Construct stats summaries based on the retrieved rows
	for _, row := range rows {
		var statsSummary model.StatsSummary
		// Check if the user already exists in statsSummaries
		found := false
		for _, ss := range statsSummaries {
			if ss.UserID == row.userID {
				statsSummary = ss
				found = true
				break
			}
		}

		// If the user is not found, create a new StatsSummary
		if !found {
			statsSummary = model.StatsSummary{
				UserID:   row.userID,
				UserName: row.userName,
			}
			statsSummaries = append(statsSummaries, statsSummary)
		}

		// Create a new MiniPoint for the current row
		miniPoint := model.MiniPoint{
			Value:     row.pointValue,
			Timestamp: row.timestamp,
		}

		// Find the corresponding stat for the PointTypeID
		var statFound bool
		for j, stat := range statsSummary.Stats {
			if stat.PointTypeID == row.pointTypeID {
				// Append the MiniPoint to the existing stat's Values
				statsSummaries[len(statsSummaries)-1].Stats[j].Values = append(stat.Values, miniPoint)
				statFound = true
				break
			}
		}

		// If stat for the PointTypeID wasn't found, create a new one
		if !statFound {
			newStat := model.Stat{
				PointTypeID:   row.pointTypeID,
				PointTypeName: row.pointTypeName,
				Values:        []model.MiniPoint{miniPoint},
			}
			statsSummaries[len(statsSummaries)-1].Stats = append(statsSummaries[len(statsSummaries)-1].Stats, newStat)
		}
	}

	return statsSummaries, nil
}
