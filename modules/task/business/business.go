package business

import (
	"context"
	"simple-clean-architecture-demo/common"
	"simple-clean-architecture-demo/modules/task"
	"simple-clean-architecture-demo/modules/task/entity"
)

type business struct {
	repository task.Repository
}

func NewBusiness(repository task.Repository) task.Business {
	return &business{repository: repository}
}

func (biz *business) CreateNewTask(ctx context.Context, data *entity.TaskCreationData) error {
	// Validate creation data
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repository.InsertNewTask(ctx, data); err != nil {
		return err
	}

	return nil
}

func (biz *business) ListTasks(ctx context.Context, filter *entity.Filter, paging *common.Paging) ([]entity.Task, error) {
	return biz.repository.FindTasks(ctx, filter, paging)
}

func (biz *business) GetTaskDetails(ctx context.Context, id string) (*entity.Task, error) {
	return biz.repository.GetTaskById(ctx, id)
}

func (biz *business) UpdateTask(ctx context.Context, id string, data *entity.TaskPatchData) error {
	// Validate creation data
	if err := data.Validate(); err != nil {
		return err
	}

	foundTask, err := biz.repository.GetTaskById(ctx, id)

	if err != nil {
		return err
	}

	if foundTask.Status == entity.StatusDone {
		return entity.ErrCannotUpdateTask
	}

	if err := biz.repository.UpdateTask(ctx, id, data); err != nil {
		return entity.ErrCannotUpdateTask
	}

	return nil
}

func (biz *business) DeleteTask(ctx context.Context, id string) error {
	_, err := biz.repository.GetTaskById(ctx, id)

	if err != nil {
		return err
	}

	if err := biz.repository.DeleteTask(ctx, id); err != nil {
		return entity.ErrCannotDeleteTask
	}

	return nil
}
