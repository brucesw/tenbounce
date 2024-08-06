package api

import (
	"fmt"
	"net/http"
	"tenbounce/repository"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
)

/*
docker run \
--name myPostgresDb \
-p 5455:5432 \
-e POSTGRES_USER=postgresUser \
-e POSTGRES_PASSWORD=postgresPW \
-e POSTGRES_DB=postgresDB \
-d \
postgres:16-alpine
*/

// docker exec -it fa58d2a22133 sh

// psql -d myPostgresDb -U postgresUser

/*
CREATE TABLE points (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	timestamp timestamp,
	user_id uuid,
	point_type_id uuid,
	value real
);
*/

/*
INSERT INTO points (timestamp, user_id, point_type_id, value) VALUES
(CURRENT_TIMESTAMP, '550e8400-e29b-41d4-a716-446655440000', '4e4b2b1c-5063-425a-a409-71b431068f78', 20.01),
(CURRENT_TIMESTAMP, '550e8400-e29b-41d4-a716-446655440000', '4e4b2b1c-5063-425a-a409-71b431068f78', 12.01),
(CURRENT_TIMESTAMP, '550e8400-e29b-41d4-a716-446655440000', '4e4b2b1c-5063-425a-a409-71b431068f78', 13.01),
(CURRENT_TIMESTAMP, '550e8400-e29b-41d4-a716-446655440000', '4e4b2b1c-5063-425a-a409-71b431068f78', 14.01),
(CURRENT_TIMESTAMP, '550e8400-e29b-41d4-a716-446655440000', '4e4b2b1c-5063-425a-a409-71b431068f78', 15.01),
(CURRENT_TIMESTAMP, '123e4567-e89b-12d3-a456-426614174000', '4e4b2b1c-5063-425a-a409-71b431068f78', 20.02),
(CURRENT_TIMESTAMP, '123e4567-e89b-12d3-a456-426614174000', '4e4b2b1c-5063-425a-a409-71b431068f78', 12.02),
(CURRENT_TIMESTAMP, '123e4567-e89b-12d3-a456-426614174000', '4e4b2b1c-5063-425a-a409-71b431068f78', 13.02),
(CURRENT_TIMESTAMP, '123e4567-e89b-12d3-a456-426614174000', '4e4b2b1c-5063-425a-a409-71b431068f78', 14.02),
(CURRENT_TIMESTAMP, '123e4567-e89b-12d3-a456-426614174000', '4e4b2b1c-5063-425a-a409-71b431068f78', 15.02);
*/

/*
CREATE TABLE users (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	name text,
	email text
);
*/

/*
INSERT INTO users (id, name, email) VALUES
('550e8400-e29b-41d4-a716-446655440000', 'Bruce Szudera Wienand', 'bszuderaw@gmail.com'),
('123e4567-e89b-12d3-a456-426614174000', 'Derek Therrien', 'dtherrien2503@gmail.com'),
('987fbc97-4bed-5078-889f-8c6e44d66b00', 'Lourens Willekes', 'lourw95@gmail.com');
*/
func tempPostgresRoute(g *echo.Group) {
	g.POST("/postgres", func(c echo.Context) error {

		var psqlInfo = "host=127.0.0.1 port=5455 user=postgresUser password=postgresPW dbname=postgresDB sslmode=disable"

		var p Repository = repository.NewPostgresRepository(psqlInfo)

		if points, err := p.ListPoints("550e8400-e29b-41d4-a716-446655440000"); err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Errorf("list points: %w", err))
		} else {
			return c.JSON(http.StatusOK, points)
		}
	})
}
