package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"homework-5/internal/pkg/repository"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func (s *server) deleteStudent(ctx context.Context, req *http.Request) (bool, int) {
	id, err := getStudentID(req.URL)
	if err != nil {
		fmt.Errorf("can't parse id: %s", err)
		return false, http.StatusBadRequest
	}
	var deleted bool
	deleted, err = s.studentRepo.Delete(ctx, id)
	if err != nil {
		fmt.Errorf("can't parse id: %s", err)
		return false, http.StatusInternalServerError
	}

	return deleted, http.StatusOK
}
func (s *server) getStudent(ctx context.Context, req *http.Request) ([]byte, int) {
	id, err := getStudentID(req.URL)
	if err != nil {
		fmt.Errorf("can't parse id: %s", err)
		return nil, http.StatusBadRequest
	}
	var students *repository.Student
	students, err = s.studentRepo.GetById(ctx, id)
	if err != nil {
		fmt.Errorf("can't parse id: %s", err)
		return nil, http.StatusInternalServerError
	}

	if students == nil {
		return nil, http.StatusInternalServerError
	}
	su := &serverStudent{}
	su.ID = students.ID
	su.Name = &students.Name
	su.Grades = students.Grades
	su.UnivID = students.UnivID
	data, err := json.Marshal(su)
	if err != nil {
		fmt.Errorf("can't marshal student with id: %d. Error: %s", id, err)
		fmt.Printf("can't marshal student with id: %d. Error: %s", id, err)
		return nil, http.StatusInternalServerError
	}

	return data, http.StatusOK
}

func (s *server) createStudent(ctx context.Context, req *http.Request) (uint, int) {
	student, err := getStudentData(req.Body)
	if err != nil {
		return 0, http.StatusBadRequest
	}
	id, err := s.studentRepo.Add(ctx, &repository.Student{Name: *student.Name, Grades: student.Grades, UnivID: student.UnivID})
	if err != nil {
		return 0, http.StatusInternalServerError
	}

	return uint(id), http.StatusOK
}

func (s *server) updateStudent(ctx context.Context, req *http.Request) (bool, int) {
	student, err := getStudentData(req.Body)
	fmt.Println(student)
	if err != nil {
		fmt.Println(err)
		return false, http.StatusBadRequest
	}
	updated, err := s.studentRepo.Update(ctx, &repository.Student{ID: student.ID, Name: *student.Name})
	if err != nil {
		fmt.Println(err)
		return false, http.StatusInternalServerError
	}

	return updated, http.StatusOK
}

func getStudentData(reader io.ReadCloser) (serverStudent, error) {
	body, err := io.ReadAll(reader)
	errWrongType := errors.New("wrong input type")
	if err != nil {
		return serverStudent{}, err
	}
	data := serverStudent{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, errWrongType
	}

	return data, nil
}

func getStudentID(reqUrl *url.URL) (uint64, error) {
	idStr := reqUrl.Query().Get("id")
	if len(idStr) == 0 {
		return 0, errors.New("can't get id")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("can't get id")
	}

	return uint64(id), nil
}
