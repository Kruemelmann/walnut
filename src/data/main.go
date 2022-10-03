package main

import (
	"context"
	"database/sql"
	"log"

	role "github.com/kruemelmann/walnut/src/data/template"
	_ "github.com/lib/pq"
)

//go:generate go run ./gen/repogen.go -entry role

func main() {
	log.Println("Starting walnut-data service")

	ctx := populateConextwithDB(context.Background())

	//TODO remove this example after its running
	_, err := role.NewRoleRepository(ctx)
	if err != nil {
		log.Fatalf("error %s\n", err)
	}
}

func populateConextwithDB(ctx context.Context) context.Context {
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/bookstore")
	if err != nil {
		log.Fatal(err)
	}
	ctx = context.WithValue(ctx, "db_session", db)
	ctx = context.WithValue(ctx, "db_kind", "postgres")
	return ctx
}
