package server

import (
	"context"
	"fmt"
	"homework-5/internal/pkg/repository"
	"homework-5/internal/pkg/repository/postgresql"
	"log"
	"net/http"
	"strconv"
)

type serverStudent struct {
	ID     uint64  `db:"id" json:"id,omitempty"`
	Name   *string `db:"name" json:"name,omitempty"`
	Grades int16   `db:"grades" json:"grades,omitempty"`
	UnivID uint64  `db:"univ_apply_id" json:"univ_id"`
}

type serverUniversity struct {
	ID       uint64  `db:"id" json:"id,omitempty"`
	Name     *string `db:"univ_name" json:"name,omitempty"`
	Facility *string `db:"facility" json:"facility,omitempty"`
}

const (
	studentPath    = "/student"
	universityPath = "/university"
	port           = ":9001"
)

type server struct {
	studentRepo    postgresql.StudentRepo
	universityRepo repository.UniversityRepo
}

func CreateServer(ctx context.Context, sr postgresql.StudentRepo, ur repository.UniversityRepo) {
	serv := server{
		studentRepo:    sr,
		universityRepo: ur,
	}
	serveMux := http.NewServeMux()

	serveMux.HandleFunc(studentPath, func(res http.ResponseWriter, req *http.Request) {
		if req == nil {
			return
		}
		switch req.Method {
		case http.MethodGet:
			data, status := serv.getStudent(ctx, req)
			res.WriteHeader(status)
			res.Write(data)
		case http.MethodPost:
			id, status := serv.createStudent(ctx, req)
			res.WriteHeader(status)
			res.Write([]byte(strconv.Itoa(int(id))))
		case http.MethodPut:
			updated, status := serv.updateStudent(ctx, req)
			res.WriteHeader(status)
			if updated {
				res.Write([]byte("\nupdated"))
			} else {
				res.Write([]byte("\nnot updated"))
			}

		case http.MethodDelete:
			deleted, status := serv.deleteStudent(ctx, req)
			res.WriteHeader(status)
			if deleted {
				res.Write([]byte("\nDeleted"))
			} else {
				res.Write([]byte("\nnot deleted"))
			}
		default:
			fmt.Printf("unsupported method: [%s]", req.Method)
			res.WriteHeader(http.StatusNotImplemented)
		}
	})
	serveMux.HandleFunc(universityPath, func(res http.ResponseWriter, req *http.Request) {
		if req == nil {
			return
		}
		switch req.Method {
		case http.MethodGet:
			data, status := serv.getUniversity(ctx, req)
			res.WriteHeader(status)
			res.Write(data)
		case http.MethodPost:
			id, status := serv.createUniversity(ctx, req)
			res.WriteHeader(status)
			res.Write([]byte(strconv.Itoa(int(id))))
		case http.MethodPut:
			updated, status := serv.updateUniversity(ctx, req)
			res.WriteHeader(status)
			if updated {
				res.Write([]byte("\nupdated"))
			} else {
				res.Write([]byte("\nnot updated"))
			}

		case http.MethodDelete:
			deleted, status := serv.deleteUniversity(ctx, req)
			res.WriteHeader(status)
			if deleted {
				res.Write([]byte("\nDeleted"))
			} else {
				res.Write([]byte("\nnot deleted"))
			}
		default:
			fmt.Printf("unsupported method: [%s]", req.Method)
			res.WriteHeader(http.StatusNotImplemented)
		}
	})
	if err := http.ListenAndServe(port, serveMux); err != nil {
		log.Fatal(err)
	}

}
