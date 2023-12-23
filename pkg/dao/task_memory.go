package dao

import (
	"context"

	"github.com/Yapcheekian/task/pkg/model"
	"github.com/samber/lo"
)

var _ TaskDAO = (*MemoryTaskDAO)(nil)

type MemoryTaskDAO struct {
	Store []*model.Task
}

func (o *MemoryTaskDAO) List(ctx context.Context) ([]*model.Task, error) {
	return o.Store, nil
}

func (o *MemoryTaskDAO) Create(ctx context.Context, task *model.Task) error {
	task.ID = o.calcID()
	o.Store = append(o.Store, task)

	return nil
}

func (o *MemoryTaskDAO) Update(ctx context.Context, task *model.Task) error {
	_, index, found := lo.FindIndexOf(o.Store, func(item *model.Task) bool {
		return item.ID == task.ID
	})

	if !found {
		return ErrTaskNotFound
	}

	o.Store[index] = task

	return nil
}

func (o *MemoryTaskDAO) Delete(ctx context.Context, id int) error {
	o.Store = lo.Filter(o.Store, func(item *model.Task, _ int) bool {
		return item.ID != id
	})

	return nil
}

func (o *MemoryTaskDAO) calcID() int {
	if len(o.Store) == 0 {
		return 1
	}

	return o.Store[len(o.Store)-1].ID + 1
}
