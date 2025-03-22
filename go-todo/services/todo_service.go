package services

import (
	"errors"
	"go-todo/models"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type TodoServiceInterface interface {
	GetAllTodos() ([]models.Todo, error)
	CreateTodo(todo *models.Todo) error
	GetTodoById(todo *models.Todo, id string) (models.Todo, error)
	UpdateTodoById(todo *models.Todo, id string) error
	DeleteTodoById(todo *models.Todo, id string) error
}

type DBInterface interface {
	Create(value interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
}

// TodoService handles the business logic and interacts with the database
type TodoService struct {
	DB DBInterface
}

func NewTodoService(db DBInterface) *TodoService {
	return &TodoService{DB: db}
}

func (s *TodoService) GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo
	err := s.DB.Find(&todos).Error

	return todos, err
}

func (s *TodoService) CreateTodo(todo *models.Todo) error {

	// checks before adding
	err := s.DB.Create(todo).Error

	return err
}

func (s *TodoService) GetTodoById(todo *models.Todo, id string) (models.Todo, error) {
	var err error = nil

	s.DB.First(todo, id)

	return *todo, err
}

// UpdateTodo applies business logic and calls the DB
func (s *TodoService) UpdateTodoById(todo *models.Todo, id string) error {
	var err error = nil
	var temp models.Todo
	s.DB.First(&temp, id)

	// log.Print("DEBUG temp:", temp)
	// log.Print("DEBUG todo:", todo)
	if todo.Title == "" {
		return errors.New("title cannot be empty")
	}
	temp.Completed = todo.Completed
	temp.Title = todo.Title

	copier.Copy(&todo, &temp)

	s.DB.Save(temp)

	return err
}

func (s *TodoService) DeleteTodoById(todo *models.Todo, id string) error {
	var err error = nil

	s.DB.First(todo, id)

	s.DB.Delete(todo, id)

	return err
}
