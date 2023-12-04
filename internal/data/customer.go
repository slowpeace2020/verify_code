package data

import (
	"context"
	"time"
)

// customer中与数据操作相关代码
type CustomerData struct {
	data *Data
}

// New 方法
func NewCustomerData(data *Data) *CustomerData {
	return &CustomerData{data: data}
}

// 设置验证码的方法
func (customerData CustomerData) SetVerifyCode(phone string, code string, expire int64) error {
	//设置key, curstomer-verify-code
	status := customerData.data.Rdb.Set(context.Background(), "CVC:"+phone, code, time.Duration(expire)*time.Second)
	if _, err := status.Result(); err != nil {
		return err
	}
	return nil
}
