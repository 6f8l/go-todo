package db

import "../schema"

type Sample struct{}

func (s *Sample) Colse() {}

func (s *Sample) Insert(todo *schema.Todo) (int, error) {
	return 0, nil
}

func (s *Sample) Delete(id int) error {
	return nil
}

func (s *Sample) GetAll() ([]schema.Todo, error) {
	return nil, nil
}
