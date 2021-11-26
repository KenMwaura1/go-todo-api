package todo

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func (repository *TodoRepository) Create(todo *Todo) error {
	if err := repository.db.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func (repository *TodoRepository) FindAll() ([]Todo, error) {
	var todos []Todo
	if err := repository.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (repository *TodoRepository) Find(id uint) (*Todo, error) {
	var todo Todo
	err := repository.database.Find(&todo, id).Error
	if todo.Name == "" {
		err = errors.New("todo not found")
	}
	return &todo, err
}

func (repository *TodoRepository) Update(todo *Todo) error {
	if err := repository.db.Save(todo).Error; err != nil {
		return err
	}
	return nil
}

func (repository *TodoRepository) Save(user Todo) (Todo, error) {
	err := repository.db.Save(user).Error
	return user, err

}

func NewTodoRepository(database *gorm.DB) *TodoRepository {
	return &TodoRepository{db: database}
}
