package input

import (
	"context"
	"encoding/json"
	"fmt"
	"homework-5/internal/pkg/repository"
	"homework-5/internal/pkg/repository/postgresql"
)

func universityCrud(ctx context.Context, inputSlice []string, universityRepo *postgresql.UniversityRepo) {

	input := inputSlice[fieldsName]
	data := repository.University{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch inputSlice[methodName] {
	case getByIdInput:
		row, errGet := universityRepo.List(ctx, data.ID)
		if errGet != nil {
			fmt.Println(err)
		}
		output, errJson := json.Marshal(&row)
		if errJson != nil {
			fmt.Println(errJson)
		}
		fmt.Println(string(output))
	case addInput:
		university := &repository.University{Name: inputSlice[2]}
		id, errAdd := universityRepo.Add(ctx, university)
		if errAdd != nil {
			fmt.Println(errAdd)
			return
		}
		fmt.Println(id)
	case deleteInput:
		deleted, errDel := universityRepo.Delete(ctx, data.ID)
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
		updated, errUp := universityRepo.Update(ctx, &data)
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
		fmt.Println("wrong command input")
	}
}
