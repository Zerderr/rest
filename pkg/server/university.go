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

func (s *server) deleteUniversity(ctx context.Context, req *http.Request) (bool, int) {
	id, err := getUniversityID(req.URL)
	if err != nil {
		fmt.Errorf("can't parse id: %s", err)
		return false, http.StatusBadRequest
	}
	var deleted bool
	deleted, err = s.universityRepo.Delete(ctx, id)
	if err != nil {
		fmt.Errorf("can't parse id: %s", err)
		return false, http.StatusInternalServerError
	}

	return deleted, http.StatusOK
}
func (s *server) getUniversity(ctx context.Context, req *http.Request) ([]byte, int) {
	id, err := getUniversityID(req.URL)
	if err != nil {
		fmt.Errorf("can't parse id: %s", err)
		return nil, http.StatusBadRequest
	}
	var universities *repository.University
	universities, err = s.universityRepo.GetById(ctx, id)
	if err != nil {
		fmt.Errorf("can't parse id: %s", err)
		return nil, http.StatusInternalServerError
	}

	if universities == nil {
		return nil, http.StatusInternalServerError
	}
	su := &serverUniversity{}
	su.ID = universities.ID
	su.Name = &universities.Name
	su.Facility = &universities.Facility
	data, err := json.Marshal(su)
	if err != nil {
		fmt.Errorf("can't marshal student with id: %d. Error: %s", id, err)
		fmt.Printf("can't marshal student with id: %d. Error: %s", id, err)
		return nil, http.StatusInternalServerError
	}

	return data, http.StatusOK
}

func (s *server) createUniversity(ctx context.Context, req *http.Request) (int64, int) {
	university, err := getUniversityData(req.Body)
	if err != nil {
		return 0, http.StatusBadRequest
	}
	id, err := s.universityRepo.Add(ctx, &repository.University{Name: *university.Name, Facility: *university.Facility})
	if err != nil {
		return 0, http.StatusInternalServerError
	}

	return int64(id), http.StatusOK
}

func (s *server) updateUniversity(ctx context.Context, req *http.Request) (bool, int) {
	university, err := getUniversityData(req.Body)
	if err != nil {
		fmt.Println(err)
		return false, http.StatusBadRequest
	}
	updated, err := s.universityRepo.Update(ctx, &repository.University{ID: university.ID, Name: *university.Name})
	if err != nil {
		fmt.Println(err)
		return false, http.StatusInternalServerError
	}

	return updated, http.StatusOK
}

func getUniversityData(reader io.ReadCloser) (serverUniversity, error) {
	body, err := io.ReadAll(reader)
	//errWrongType := errors.New("wrong input type")
	if err != nil {
		return serverUniversity{}, errors.New("wrong input type")
	}
	data := serverUniversity{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, errors.New("wrong input type")
	}

	return data, nil
}

func getUniversityID(reqUrl *url.URL) (uint64, error) {
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
