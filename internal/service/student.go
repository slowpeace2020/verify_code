package service

import (
	"context"
	"fmt"
	"verify-code/internal/biz"

	pb "verify-code/api/student/v1"
)

type StudentService struct {
	pb.UnimplementedStudentServer
	uc *biz.StudentUsecase
}

func NewStudentService(uc *biz.StudentUsecase) *StudentService {
	return &StudentService{uc: uc}
}

func (s *StudentService) CallStudent(ctx context.Context, req *pb.CallStudentRequest) (*pb.CallStudentReply, error) {
	msg := fmt.Sprintf("姓名：%s, 年龄：%d", req.Name, req.Age)
	return &pb.CallStudentReply{Message: msg}, nil
}
