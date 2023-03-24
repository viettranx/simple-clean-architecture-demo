package task

import (
	"context"
	"github.com/gin-gonic/gin"
	"simple-clean-architecture-demo/common"
	"simple-clean-architecture-demo/modules/task/entity"
)

type Repository interface {
	InsertNewTask(ctx context.Context, data *entity.TaskCreationData) error
	GetTaskById(ctx context.Context, id string) (*entity.Task, error)
	FindTasks(ctx context.Context, filter *entity.Filter, paging *common.Paging) ([]entity.Task, error)
	UpdateTask(ctx context.Context, id string, data *entity.TaskPatchData) error
	DeleteTask(ctx context.Context, id string) error
}

type Business interface {
	CreateNewTask(ctx context.Context, data *entity.TaskCreationData) error
	ListTasks(ctx context.Context, filter *entity.Filter, paging *common.Paging) ([]entity.Task, error)
	GetTaskDetails(ctx context.Context, id string) (*entity.Task, error)
	UpdateTask(ctx context.Context, id string, data *entity.TaskPatchData) error
	DeleteTask(ctx context.Context, id string) error
}

type API interface {
	CreateTaskHdl() gin.HandlerFunc
	ListTaskHdl() gin.HandlerFunc
	GetTaskHdl() gin.HandlerFunc
	UpdateTaskHdl() gin.HandlerFunc
	DeleteTaskHdl() gin.HandlerFunc
}
