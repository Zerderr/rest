package input

import (
	"context"
	"encoding/json"
	"fmt"
	"homework-5/internal/pkg/repository"
	"homework-5/internal/pkg/repository/postgresql"
)

func studentCrud(ctx context.Context, inputSlice []string, studentRepo postgresql.StudentRepo) {

	input := inputSlice[fieldsName]
	data := repository.Student{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch inputSlice[methodName] {
	case getByIdInput:
		row, errGet := studentRepo.GetById(ctx, data.ID)
		if errGet != nil {
			fmt.Println(err)
			return
		}
		output, errJson := json.Marshal(&row)
		if errJson != nil {
			fmt.Println(errJson)
			return
		}
		fmt.Println(string(output))

	case addInput:
		id, errAdd := studentRepo.Add(ctx, &data)
		if errAdd != nil {
			fmt.Println(errAdd)
			return
		}
		fmt.Println(id)
	case deleteInput:
		deleted, errDel := studentRepo.Delete(ctx, data.ID)
		if errDel != nil {
			fmt.Println(errDel)
			return
		}
		if deleted {
			fmt.Println("Deleted")
		} else {
			fmt.Println("id not found")
		}
	case updateInput:
		updated, errUp := studentRepo.Update(ctx, &data)
		if errUp != nil {
			fmt.Println(errUp)
			return
		}
		if updated {
			fmt.Println("Updated")
		} else {
			fmt.Println("id not found")
		}
	default:
		fmt.Println("wrong method input")
	}
}
