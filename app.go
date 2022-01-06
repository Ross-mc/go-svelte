package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/ross-mc/go-svelte-proj/server"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

const PORT = "8080"

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID   int64 `bun:",pk,autoincrement"`
	Name string
}

func main() {
	http.HandleFunc("/", server.Router)
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Error loading .env file")
	}
	dsn := "postgres://postgres:@fullstack-postgres:5432/fullstack_api?user=ross&password=password&sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	table := db.NewCreateTable().Model((*User)(nil)).Table("Users").IfNotExists()
	fmt.Println(table)
	fmt.Printf("Listening for requests on port: %v\n", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
