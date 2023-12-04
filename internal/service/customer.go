package service

import (
	"context"
	"verify-code/internal/data"
)

type CustomerService struct {
	customerData *data.CustomerData
}

func NewCustomerService(customerData *data.CustomerData) *CustomerService {
	return &CustomerService{
		customerData: customerData,
	}
}

// 获取验证码
func (s *CustomerService) GetVerifyCode(ctx context.Context) {

}
