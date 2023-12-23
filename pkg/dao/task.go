package dao

import (
	"context"
	"errors"

	"github.com/Yapcheekian/task/pkg/model"
)

var ErrTaskNotFound = errors.New("task not found ")

type TaskDAO interface {
	List(ctx context.Context) ([]*model.Task, error)
	Create(ctx context.Context, task *model.Task) error
	Update(ctx context.Context, task *model.Task) error
	Delete(ctx context.Context, id int) error
}
