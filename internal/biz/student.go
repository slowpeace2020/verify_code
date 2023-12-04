package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Student struct {
	Name string
	Age  int
}

type StudentRepo interface {
	Save(ctx context.Context, student *Student) (*Student, error)
}

type StudentUsecase struct {
	repo StudentRepo
	log  *log.Helper
}

func NewStudentUsecase(repo StudentRepo, logger log.Logger) *StudentUsecase {
	return &StudentUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateStudent creates a Student, and returns the new Student.
func (uc *StudentUsecase) SaveStudent(ctx context.Context, g *Student) (*Student, error) {
	uc.log.WithContext(ctx).Infof("CreateStudent: %v", g.Name)
	return uc.repo.Save(ctx, g)
}
