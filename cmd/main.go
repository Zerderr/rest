package main

import (
	"context"
	"fmt"
	"homework-5/internal/pkg/db"
	"homework-5/internal/pkg/input"
	"sync"
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
	var wg sync.WaitGroup
	wg.Add(1)
	go input.InputConsole(ctx, database)
	wg.Wait()
}
