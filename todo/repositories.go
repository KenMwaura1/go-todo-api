package todo

import (
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

func (repository *TodoRepository) FindById(id uint) (*Todo, error) {
	var todo Todo
	if err := repository.db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
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

func (repository *TodoRepository) Delete(id uint) error {
	if err := repository.db.Delete(&Todo{}, id).Error; err != nil {
		return err
	}
	return nil
}

func NewTodoRepository(database *gorm.DB) *TodoRepository {
	return &TodoRepository{db: database}
}
