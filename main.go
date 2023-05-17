package main

import (
	"context"
	"fmt"
	"homework-5/internal/pkg/db"
	"homework-5/internal/pkg/repository/postgresql"
	"homework-5/internal/pkg/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	database, err := db.NewDB(ctx)
	defer database.GetPool().Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	studentRepo := postgresql.NewStudent(database)
	univrsityRepo := postgresql.NewUniversity(database)
	server.CreateServer(ctx, studentRepo, univrsityRepo)

	//var wg sync.WaitGroup
	//wg.Add(1)
	//
	//go input.InputConsole(ctx, database)
	//wg.Wait()
}

//curl -v -X GET "http://localhost:9001/?id=4"
