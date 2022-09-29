package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/kruemelmann/walnut/src/data/role"
	_ "github.com/lib/pq"
)

func populateConextwithDB(ctx context.Context) context.Context {
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/bookstore")
	if err != nil {
		log.Fatal(err)
	}
	ctx = context.WithValue(ctx, "db_session", db)
	ctx = context.WithValue(ctx, "db_kind", "postgres")
	return ctx
}

func main() {
	log.Println("Starting walnut-data service")

	ctx := populateConextwithDB(context.Background())

	//TODO isolate to own function
	//FIXME remove thats just a test
	rolerepo, err := role.NewRoleRepository(ctx)
	if err != nil {
		log.Fatalf("error %s\n", err)
	}
	roles, err := rolerepo.GetRoles()
	if err != nil {
		log.Fatalf("error %s\n", err)
	}
	for k, v := range roles {
		fmt.Println(k, v)
	}
	//FIXME =========== END
}
