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

func (repository *TodoRespository) Find(id uint) (*Todo, error) {
	var todo Todo
	err := repository.database.Find(&todo, id).Error
	if todo.Name == "" {
		err = errors.New("todo not found")
	}
	return &todo, err
}

func (repository *TodoRespository) Update(todo *Todo) error {
	if err := repository.db.Save(todo).Error; err != nil {
		return err
	}
	return nil
}

func (repository *TodoRespository) Save(user Todo) (Todo, error) {
	err := repository.db.Save(user).Error
	return user, err

}

func NewTodoRepository(database *gorm.DB) *TodoRespository {
	return &TodoRespository{db: database}
}
