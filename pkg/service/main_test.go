package service

import (
	"testing"

	"github.com/Yapcheekian/task/internal/daomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var (
	svc     *Service
	taskDAO *daomock.MockTaskDAO
)

var _ = BeforeEach(func() {
	ctrl := gomock.NewController(GinkgoT())
	taskDAO = daomock.NewMockTaskDAO(ctrl)
	svc = &Service{
		TaskDAO: taskDAO,
	}
})

func TestDAO(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suite")
}
