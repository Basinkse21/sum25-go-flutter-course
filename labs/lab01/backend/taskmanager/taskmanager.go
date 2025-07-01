package taskmanager

import (
	"errors"
	"time"
)

// Predefined errors
var (
	ErrTaskNotFound = errors.New("task not found")
	ErrEmptyTitle   = errors.New("title cannot be empty")
)

// Task represents a single task
type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
	CreatedAt   time.Time
}

// TaskManager manages a collection of tasks
type TaskManager struct {
	tasks  map[int]Task
	nextID int
}

// NewTaskManager creates a new task manager
func NewTaskManager() *TaskManager {
<<<<<<< HEAD
	// TODO: Implement this function
	return nil
}

// AddTask adds a new task to the manager, returns an error if the title is empty, and increments the nextID
func (tm *TaskManager) AddTask(title, description string) (Task, error) {
	// TODO: Implement this function
	return Task{}, nil
=======
	return &TaskManager{
		tasks:  make(map[int]*Task),
		nextID: 1,
	}
}

// AddTask adds a new task to the manager
func (tm *TaskManager) AddTask(title, description string) (*Task, error) {
	if title == "" {
		return nil, ErrEmptyTitle
	}

	task := &Task{
		ID:          tm.nextID,
		Title:       title,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
	}

	tm.tasks[tm.nextID] = task
	tm.nextID++

	return task, nil
>>>>>>> a0b7266 (lab01: реализованы базовые задачи Go и Flutter)
}

// UpdateTask updates an existing task, returns an error if the title is empty or the task is not found
func (tm *TaskManager) UpdateTask(id int, title, description string, done bool) error {
<<<<<<< HEAD
	// TODO: Implement this function
=======
	if id <= 0 {
		return ErrInvalidID
	}

	task, exists := tm.tasks[id]
	if !exists {
		return ErrTaskNotFound
	}

	if title == "" {
		return ErrEmptyTitle
	}

	task.Title = title
	task.Description = description
	task.Done = done

>>>>>>> a0b7266 (lab01: реализованы базовые задачи Go и Flutter)
	return nil
}

// DeleteTask removes a task from the manager, returns an error if the task is not found
func (tm *TaskManager) DeleteTask(id int) error {
<<<<<<< HEAD
	// TODO: Implement this function
	return nil
}

// GetTask retrieves a task by ID, returns an error if the task is not found
func (tm *TaskManager) GetTask(id int) (Task, error) {
	// TODO: Implement this function
	return Task{}, nil
}

// ListTasks returns all tasks, optionally filtered by done status, returns an empty slice if no tasks are found
func (tm *TaskManager) ListTasks(filterDone *bool) []Task {
	// TODO: Implement this function
	return nil
=======
	if id <= 0 {
		return ErrInvalidID
	}

	if _, exists := tm.tasks[id]; !exists {
		return ErrTaskNotFound
	}

	delete(tm.tasks, id)
	return nil
}

// GetTask retrieves a task by ID
func (tm *TaskManager) GetTask(id int) (*Task, error) {
	if id <= 0 {
		return nil, ErrInvalidID
	}

	task, exists := tm.tasks[id]
	if !exists {
		return nil, ErrTaskNotFound
	}

	return task, nil
}

// ListTasks returns all tasks, optionally filtered by done status
func (tm *TaskManager) ListTasks(filterDone *bool) []*Task {
	var tasks []*Task

	for _, task := range tm.tasks {
		if filterDone == nil || *filterDone == task.Done {
			tasks = append(tasks, task)
		}
	}

	return tasks
>>>>>>> a0b7266 (lab01: реализованы базовые задачи Go и Flutter)
}
