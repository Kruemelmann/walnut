package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/kruemelmann/walnut/src/data/role"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting walnut-data service")

	//TODO isolate to own function
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/bookstore")
	if err != nil {
		log.Fatal(err)
	}

	//TODO isolate to own function
	ctx := context.WithValue(context.Background(), "db_session", db)
	ctx = context.WithValue(ctx, "db_kind", "postgres")

	rolerepo, err := role.NewRoleRepository(ctx)
	if err != nil {
		log.Fatalf("error %s\n", err)
	}

	//FIXME remove thats just a test
	roles, err := rolerepo.GetRoles()
	if err != nil {
		log.Fatalf("error %s\n", err)
	}

	for k, v := range roles {
		fmt.Println(k, v)
	}
}
