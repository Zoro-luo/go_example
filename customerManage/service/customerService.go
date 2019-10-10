package service
import (
	"go_code/customerManage/model"
)

//该CustomerService 完成对Customer的操作 包括赠删改查
type CustomerService struct {
	customers []model.Customer
	//申明一个字段 表示当前切片含有多少个客户
	customerNum int
}

//编写一个函数 返回一个 *CustomerService
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	cucustomer := model.NewCustomer(1,"张三","nan",22,"112","zs@163.com")
	customerService.customers = append(customerService.customers , cucustomer)
	return customerService
}
//根据id找到对应切片的下标 没有则返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i<len(this.customers);i++{
		if this.customers[i].Id == id {	//找到
			index = i
		}
	}
	return index
}


//返回客户列表客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

//添加客户到Customer切片
func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers,customer)
	return true
}

//更改
// func (this *CustomerService) Update(id int) bool {
// 	index := this.FindById(id)
// 	if index == -1 {
// 		return false
// 	}
// }

//删除
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}