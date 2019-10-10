package main

import (
	"fmt"
	"go_code/customerManage/service"
	"go_code/customerManage/model"
)

type customerView struct {
	key string 	//接受用户的输入
	loop bool	//控制循环
	customerService *service.CustomerService
}

//显示所有的客户信息
func (this *customerView) list() {
	//获取当前所有的客户信息 在切片中
	customers := this.customerService.List()
	//显示：
	fmt.Println("------------------客户列表-------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t手机\t邮箱")
	for i := 0; i < len(customers) ; i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("----------------客户列表完成-------------------")
}

func (this *customerView) add() {
	fmt.Println("----------------修改客户-------------------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gander := ""
	fmt.Scanln(&gander)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮")
	email := ""
	fmt.Scanln(&email)
	//构建一个新的Customer实例
	customer := model.NewCustomer2(name,gander,age,phone,email)
	//调用
	if this.customerService.Add(customer) {
		fmt.Println("-------------添加完成--------------------")
	}else{
		fmt.Println("-------------添加失败--------------------")
	}
}

//得到用户的输入id 删除id对应的客户
func (this *customerView) delete() {
	fmt.Println("----------------删除客户-------------------------")
	fmt.Println("请删除待删除客户编号(-1退出)：")
	id := -1 
	fmt.Scanln(&id)
	if id == -1 {
		return 	//放弃删除
	}
	fmt.Println("是否确认删除(Y/N):")
	choice := ""
	fmt.Scanln(&choice)
	if choice =="y" || choice == "Y" {
		//调用CustomerService的Delete方法
		if this.customerService.Delete(id) {
			fmt.Println("----------------删除成功-------------------------")
		}else{
			fmt.Println("----------------删除失败,该id号不存在--------------")
		}
	}
}
//退出软件
func (this *customerView) exit() {
	fmt.Println("确认是否退出(Y/N):")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}
		fmt.Println("你的输入有误,确认是否退出(Y/N):")
	}

	if this.key == "Y" || this.key == "y"{
		this.loop = false
		fmt.Println("你退出了客户关系管理系统!...")
	}

}


func (this *customerView) mainMenu() {
	for {
		fmt.Println("---------------客户信息管理软件---------------")
		fmt.Println("		1 添加客户")
		fmt.Println("		2 修改客户")
		fmt.Println("		3 删除客户")
		fmt.Println("		4 客户列表")
		fmt.Println("		5 退    出")
		fmt.Print("请选择(1-5)：")

		fmt.Scanln(&this.key)
		switch this.key {
			case "1":		//添加客户
				this.add()
			case "2":		//修改客户
				fmt.Println("修改客户")
			case "3":		//删除客户
				this.delete()
			case "4":		//客户列表
				this.list()
			case "5":		//退    出
				this.exit()
			default :
			fmt.Println("输入有误,请重新输入....")
		}

		if !this.loop {
			break
		}

	}
}

func main() {
	//创建一个customerView
	customerView := customerView{
		key : "",
		loop : true,
	}
	//完成对customerView结构体的customerService字段的初始化工作
	customerView.customerService = service.NewCustomerService()

	//显示主菜单
	customerView.mainMenu()
}

