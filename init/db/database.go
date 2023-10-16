package db

import (
	"api-estadia-express/init/logger"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"sync"
	"time"
)

var lock = &sync.Mutex{}
var postgresConnection *sql.DB

func ConnectionToPostgres() *sql.DB {
	if postgresConnection == nil {
		lock.Lock()
		defer lock.Unlock()
		if postgresConnection == nil {
			psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))

			db, err := sql.Open("postgres", psqlconn)
			if err != nil {
				logger.Fatal("error connecting to database", err)
			}

			ctx, c := context.WithTimeout(context.Background(), 10*time.Second)
			defer c()

			if err = db.PingContext(ctx); err != nil {
				logger.Fatal("error connecting to database", err)
			}

			postgresConnection = db
		}
	}
	return postgresConnection
}
