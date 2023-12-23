package dao

import (
	"context"

	"github.com/Yapcheekian/task/pkg/model"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("MemoryTaskDAO", func() {
	var (
		dao *MemoryTaskDAO
		ctx context.Context

		err error
	)

	BeforeEach(func() {
		dao = &MemoryTaskDAO{
			Store: make([]*model.Task, 0),
		}
		ctx = context.Background()
	})

	Describe("List", func() {
		var result []*model.Task

		BeforeEach(func() {
			dao.Store = append(dao.Store, &model.Task{ID: 1})
		})

		AfterEach(func() {
			dao.Store = make([]*model.Task, 0)
		})

		JustBeforeEach(func() {
			result, err = dao.List(ctx)
		})

		Context("normal case", func() {
			It("should not have error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("return correct result", func() {
				Expect(result).To(ConsistOf(&model.Task{ID: 1}))
			})
		})
	})

	Describe("Create", func() {
		var task *model.Task

		JustBeforeEach(func() {
			err = dao.Create(ctx, task)
		})

		AfterEach(func() {
			dao.Store = make([]*model.Task, 0)
		})

		Context("normal case", func() {
			BeforeEach(func() {
				task = &model.Task{
					Name: "test",
				}
			})

			It("should not have error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("create new record", func() {
				result, err := dao.List(ctx)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(ConsistOf(&model.Task{ID: 1, Name: "test", Status: false}))
				Expect(task).To(PointTo(MatchFields(IgnoreExtras, Fields{
					"ID":     Equal(1),
					"Name":   Equal("test"),
					"Status": Equal(false),
				})))
			})
		})
	})

	Describe("Update", func() {
		var task *model.Task

		BeforeEach(func() {
			dao.Store = append(dao.Store, &model.Task{ID: 1, Name: "test", Status: false})
		})

		AfterEach(func() {
			dao.Store = make([]*model.Task, 0)
		})

		JustBeforeEach(func() {
			err = dao.Update(ctx, task)
		})

		Context("normal case", func() {
			BeforeEach(func() {
				task = &model.Task{
					ID:     1,
					Name:   "test2",
					Status: true,
				}
			})

			It("should not have error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("update record", func() {
				result, err := dao.List(ctx)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(ConsistOf(&model.Task{ID: 1, Name: "test2", Status: true}))
			})
		})
	})

	Describe("Delete", func() {
		var id int

		BeforeEach(func() {
			dao.Store = append(dao.Store, &model.Task{ID: 1, Name: "test", Status: false})
		})

		AfterEach(func() {
			dao.Store = make([]*model.Task, 0)
		})

		JustBeforeEach(func() {
			err = dao.Delete(ctx, id)
		})

		Context("normal case", func() {
			BeforeEach(func() {
				id = 1
			})

			It("should not have error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("delete record", func() {
				result, err := dao.List(ctx)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(HaveLen(0))
			})
		})
	})
})
