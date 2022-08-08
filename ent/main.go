package main

import (
	"context"
	"fmt"
	"log"
	"main/ent"
	"main/ent/user"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(context.Background())
	if err != nil {
		log.Fatalf("failed creating user: %w", err)
	}
	fmt.Println(u)

	users, err := client.User.
		Query().
		Where(user.Name("a8m")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		All(context.Background())
	if err != nil {
		log.Fatalf("failed get user: %w", err)
	}
	fmt.Println(users)

}
