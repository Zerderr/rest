package input

import (
	"bufio"
	"context"
	"fmt"
	"homework-5/internal/pkg/db"
	"homework-5/internal/pkg/repository/postgresql"
	"os"
	"strings"
)

func InputConsole(ctx context.Context, database *db.Database) {
	for {
		studentRepo := postgresql.NewStudent(database)
		universityRepo := postgresql.NewUniversity(database)
		in := bufio.NewScanner(os.Stdin)
		in.Scan()
		if err := in.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Input error: ", err)
			return
		}
		inputSlice := strings.Fields(in.Text())
		if len(inputSlice) != correctRequestLen {
			fmt.Println("Invalid request")
			continue
		}
		switch inputSlice[tableName] {
		case studentInput:
			studentCrud(ctx, inputSlice, studentRepo)
		case universityInput:
			universityCrud(ctx, inputSlice, universityRepo)
		default:
			fmt.Println("wrong table input")

		}

	}
}
