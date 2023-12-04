package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"verify-code/internal/biz"
)

type studentRepo struct {
	data *Data
	log  *log.Helper
}

func (s studentRepo) Save(ctx context.Context, student *biz.Student) (*biz.Student, error) {
	//TODO implement me
	return student, nil
}

// NewStudentRepo .
func NewStudentRepo(data *Data, logger log.Logger) biz.StudentRepo {
	return &studentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
