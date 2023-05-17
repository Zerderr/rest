package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"homework-5/internal/pkg/repository"
	"homework-5/internal/pkg/repository/postgresql"
	"testing"
)

func TestCreateStudent(t *testing.T) {
	A.SetUp(t)
	defer A.TearDown()
	tdb, err := NewTestDB(context.Background())
	universityRepo := postgresql.NewUniversity(tdb)
	univ_id, err := universityRepo.Add(context.Background(), &repository.University{
		Name:     "test",
		Facility: "test",
	})

	userRepo := postgresql.NewStudent(tdb)
	res, err := userRepo.Add(context.Background(), &repository.Student{
		Name:   "test",
		Grades: 250,
		UnivID: uint64(univ_id),
	})

	assert.NoError(t, err)
	assert.Equal(t, uint64(1), res)
}

func TestCreateUniversity(t *testing.T) {
	A.SetUp(t)
	defer A.TearDown()
	tdb, err := NewTestDB(context.Background())
	universityRepo := postgresql.NewUniversity(tdb)
	res, err := universityRepo.Add(context.Background(), &repository.University{
		Name:     "test",
		Facility: "test",
	})
	assert.NoError(t, err)
	assert.Equal(t, int64(1), res)
}
