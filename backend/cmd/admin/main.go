// Package main is the entry point for the Admin CLI tool.
package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/myorg/myapp/backend/internal/application"
	"github.com/myorg/myapp/backend/internal/infra"
)

func main() {
	// 1. Setup Database
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./local.db"
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	// 2. Initialize Layers
	userRepo := infra.NewSQLiteUserRepository(db)
	userService := application.NewUserService(userRepo)

	// 3. Define Commands
	registerCmd := flag.NewFlagSet("register", flag.ExitOnError)
	email := registerCmd.String("email", "", "User email")
	password := registerCmd.String("password", "", "User password")

	if len(os.Args) < 2 {
		fmt.Println("expected 'register' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "register":
		registerCmd.Parse(os.Args[2:])
		if *email == "" || *password == "" {
			fmt.Println("email and password are required")
			registerCmd.Usage()
			os.Exit(1)
		}

		user, err := userService.Register(context.Background(), *email, *password)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("User registered successfully: %s (ID: %s)\n", user.Email, user.ID.String())

	default:
		fmt.Printf("unknown subcommand: %s\n", os.Args[1])
		os.Exit(1)
	}
}
