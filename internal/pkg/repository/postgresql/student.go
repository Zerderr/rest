package postgresql

import (
	"context"
	"homework-5/internal/pkg/db"
	"homework-5/internal/pkg/repository"
)

type StudentRepo struct {
	db db.DBops
}

func NewStudent(db db.DBops) *StudentRepo {
	return &StudentRepo{db: db}
}

func (r *StudentRepo) Add(ctx context.Context, data *repository.Student) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, "INSERT INTO student(name, grades, univ_apply_id) VALUES ($1, $2, $3) RETURNING id", data.Name, data.Grades, data.UnivID).Scan(&id)
	return id, err
}
func (r *StudentRepo) GetById(ctx context.Context, id uint64) ([]*repository.Student, error) {
	students := make([]*repository.Student, 10)
	err := r.db.Select(ctx, &students, "select id, name, grades, univ_apply_id from student where id = $1", id)
	return students, err
}
func (r *StudentRepo) Update(ctx context.Context, data *repository.Student) (bool, error) {
	result, err := r.db.Exec(ctx, "update student set name = $1 where id = $2", data.Name, data.ID)
	return result.RowsAffected() > 0, err
}

func (r *StudentRepo) Delete(ctx context.Context, id uint64) (bool, error) {
	result, err := r.db.Exec(ctx, "DELETE FROM student where id=$1", id)
	if err != nil {
		return result.RowsAffected() > 0, err
	}
	return result.RowsAffected() > 0, err
}
