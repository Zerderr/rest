package tests

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"homework-5/internal/pkg/db"
	"strings"
	"sync"
	"testing"
)

const (
	tHost     = "localhost"
	tPort     = 5438
	tUser     = "test"
	tPassword = "test"
	tDbname   = "test_db"
)

type TDB struct {
	sync.Mutex
	db *db.Database
}

func newDataBase(cluster *pgxpool.Pool) *db.Database {
	return &db.Database{Cluster: cluster}
}

func NewTestDB(ctx context.Context) (*db.Database, error) {
	dsn := generateTestDsn()
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return newDataBase(pool), nil
}

func generateTestDsn() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", tHost, tPort, tUser, tPassword, tDbname)
}

func (d *TDB) SetUp(t *testing.T) {
	t.Helper()
	ctx := context.Background()
	d.Lock()
	d.Truncate(ctx)

}

func (d *TDB) TearDown() {
	defer d.Unlock()
	d.Truncate(context.Background())

}

func (d *TDB) Truncate(ctx context.Context) {
	var tables []string
	err := d.db.Select(ctx, &tables, "SELECT table_name from information_schema.tables where table_schema ='public' AND table_type='BASE TABLE' and table_name!='goose_db_version'")
	if err != nil {
		fmt.Println(err)
	}
	//Чистим таблички
	q := fmt.Sprintf("Truncate table %s", strings.Join(tables, ","))
	if _, err := d.db.Exec(ctx, q); err != nil {
		panic(err)
	}
	//Обновляем id так как он serial и после каждого теста увеличивается
	update := "ALTER SEQUENCE university_id_seq RESTART WITH 1"
	if _, err := d.db.Exec(ctx, update); err != nil {
		panic(err)
	}
	update = "UPDATE student SET id=nextval('university_id_seq')"
	if _, err := d.db.Exec(ctx, update); err != nil {
		panic(err)
	}
	update = "ALTER SEQUENCE student_id_seq RESTART WITH 1"
	if _, err := d.db.Exec(ctx, update); err != nil {
		panic(err)
	}
	update = "UPDATE student SET id=nextval('student_id_seq')"
	if _, err := d.db.Exec(ctx, update); err != nil {
		panic(err)
	}

}

var (
	A = TDB{
		Mutex: sync.Mutex{},
		db:    nil,
	}
	err error
)

func init() {

	A.db, err = NewTestDB(context.Background())

}
