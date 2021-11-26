package todo

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type TodoRespository struct {
	db *gorm.DB
}

func (repository *TodoRespository) Create(todo *Todo) error {
	if err := repository.db.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func (repository *TodoRespository) FindAll() ([]Todo, error) {
	var todos []Todo
	if err := repository.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (repository *TodoRespository) FindById(id uint) (*Todo, error) {
	var todo Todo
	if err := repository.db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}
