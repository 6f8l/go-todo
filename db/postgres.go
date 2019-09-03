package db

import (
	"database/sql"

	"../schema"
)

type Postgres struct {
	DB *sql.DB
}

func (p *Postgres) Close() {
}

func (p *Postgres) Insert(todo *schema.Todo) (int, error) {
	return 0, nil
}

func (p *Postgres) Delete(id int) error {
	return nil
}

func (p *Postgres) GetAll() ([]schema.Todo, error) {
	return nil, nil
}
