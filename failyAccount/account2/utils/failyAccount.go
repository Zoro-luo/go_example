package utils

import (
	"fmt"
)

type FamilyAccount struct {
	//控制是否退出for
	loop bool
	//保存接受用户输入的选项
	key string
	//余额
	balace float64
	//每次收支的金额
	money float64
	//否有收支的行为
	flag bool
	//每次收支说明
	note string
	//收支的详情
	details string
}

//编写一个工厂模式的构造方法, 返回一个*FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balace:  10000.0,
		money:   0.0,
		flag:    false,
		note:    "",
		details: "收支\t账户金额\t收支金额\t说  明",
	}
}

//收支明细
func (this *FamilyAccount) showDetails() {
	fmt.Println("------------当前收支明细--------------")
	if this.flag {
		fmt.Println(this.details)

	} else {
		fmt.Println("  当前没有收支 请来一笔吧!")
	}
}

//登记收入
func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	this.balace += this.money
	this.flag = true
	fmt.Println("本次收入说明")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收支\t%v\t\t+%v\t\t%v", this.balace, this.money, this.note)
}

//3登记支出
func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	if this.money > this.balace {
		fmt.Println("余额不足,请重新输入!")
		return
	}
	this.balace -= this.money
	this.flag = true
	fmt.Println("本次支出说明")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收支\t%v\t\t-%v\t\t%v", this.balace, this.money, this.note)
}

//退出软件
func (this *FamilyAccount) exit() {
	fmt.Println("你确定要退出吗? y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "Y" || choice == "n" || choice == "N" {
			break
		}
		fmt.Println("你的输入有误,请重新输入? y/n")
	}
	if choice == "y" || choice == "Y" {
		this.loop = false
	}
}

//显示主菜单
func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n---------家庭收支记账软件-----------")
		fmt.Println("           1 收支明细")
		fmt.Println("           2 登记收入")
		fmt.Println("           3 登记支出")
		fmt.Println("           4 退出软件")
		fmt.Print("请选择(1-4):")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1": //收支明细
			this.showDetails()
		case "2": //登记收入
			this.income()
		case "3": //登记支出
			this.pay()
		case "4": //退出软件
			this.exit()
		default:
			fmt.Println("请输入正确的现象")
		}
		if !this.loop {
			break
		}
	}
}
