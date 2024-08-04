package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

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

func tempPostgresRoute(g *echo.Group) {
	g.POST("/postgres", func(c echo.Context) error {

		// TODO(bruce): mysql??
		db, err := sql.Open("postgres",
			"postgresUser:postgresPW@tcp(127.0.0.1:5455)/hello")
		if err != nil {
			log.Fatal(err)
		}

		err = db.Ping()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "ping db")
		}
		defer db.Close()

		fmt.Println("gotme")
		return c.JSON(http.StatusOK, nil)
	})

}
