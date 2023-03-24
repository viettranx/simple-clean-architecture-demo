package inmem

import (
	"context"
	"github.com/google/uuid"
	"simple-clean-architecture-demo/common"
	"simple-clean-architecture-demo/modules/task"
	"simple-clean-architecture-demo/modules/task/entity"
	"sync"
)

// This storage has only task data, we don't have anything about User data
var initData = []entity.Task{
	{
		Id:          "57606a8d-9348-4cc2-ac0b-c7886108e65e",
		Title:       "This a task 1",
		Description: "Description of task 1",
		Status:      entity.StatusDoing,
	},
	{
		Id:          "46dacc98-4084-4901-b90e-267d4219f374",
		Title:       "This a task 2",
		Description: "Description of task 2",
		Status:      entity.StatusDoing,
	},
	{
		Id:          "9fb2f6a6-dcca-4d17-bd58-4bda471d914d",
		Title:       "This a task 3",
		Description: "Description of task 3",
		Status:      entity.StatusDone,
	},
	{
		Id:          "767f3519-44e0-4a25-a60c-a769a8bdbfc1",
		Title:       "This a task 4",
		Description: "Description of task 3",
		Status:      entity.StatusDoing,
	},
}

type inMemStorage struct {
	db []entity.Task
	l  *sync.RWMutex
}

func NewInMemStorage() task.Repository {
	return &inMemStorage{
		db: initData,
		l:  new(sync.RWMutex),
	}
}

func (s *inMemStorage) InsertNewTask(ctx context.Context, data *entity.TaskCreationData) error {
	s.l.Lock()
	defer s.l.Unlock()

	newTaskId := uuid.New().String()
	newTask := entity.Task{
		Id:          newTaskId,
		Title:       data.Title,
		Description: data.Description,
		Status:      data.Status,
	}

	s.db = append(s.db, newTask)
	data.Id = newTaskId // carry inserted id

	return nil
}

func (s *inMemStorage) GetTaskById(ctx context.Context, id string) (*entity.Task, error) {
	s.l.RLock()
	defer s.l.RUnlock()

	for i, t := range s.db {
		if t.Id == id {
			foundTask := s.db[i] // copy task
			return &foundTask, nil
		}
	}

	return nil, entity.ErrTaskNotFound
}

func (s *inMemStorage) FindTasks(ctx context.Context, f *entity.Filter, p *common.Paging) ([]entity.Task, error) {
	s.l.RLock()
	defer s.l.RUnlock()

	tasks := make([]entity.Task, len(s.db))

	// we should copy it to avoid data racing when return s.db directly
	// in this example, we just return all tasks
	copy(tasks, s.db)
	p.Total = int64(len(tasks))

	return tasks, nil
}

func (s *inMemStorage) UpdateTask(ctx context.Context, id string, data *entity.TaskPatchData) error {
	s.l.Lock()
	defer s.l.Unlock()

	var foundTask *entity.Task

	// In memory storage, we have to loop find task.
	// In DBMS, we should not do this
	for i, t := range s.db {
		if t.Id == id {
			foundTask = &s.db[i]
			break
		}
	}

	if foundTask != nil {
		if data.Title != nil {
			foundTask.Title = *data.Title
		}

		if data.Description != nil {
			foundTask.Description = *data.Description
		}

		if data.Status != nil {
			foundTask.Status = *data.Status
		}
	}

	return nil
}

func (s *inMemStorage) DeleteTask(ctx context.Context, id string) error {
	s.l.Lock()
	defer s.l.Unlock()

	for i := range s.db {
		if s.db[i].Id == id {
			s.db = append(s.db[:i], s.db[i+1:]...)
			break
		}
	}

	return nil
}
