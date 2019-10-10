package main

//这是面向过程的方式实例~~
//家庭收支记账软件
import (
	"fmt"
)

func main() {
	loop := true
	key := ""
	//余额
	balace := 10000.0
	//每次收支的金额
	money := 0.0
	//定义一个变量是否有收支的行为
	flag := false
	//每次收支说明
	note := ""
	//收支的详情
	details := "收支\t账户金额\t收支金额\t说  明"
	//显示主菜单
	for {
		fmt.Println("\n---------家庭收支记账软件-----------")
		fmt.Println("           1 收支明细")
		fmt.Println("           2 登记收入")
		fmt.Println("           3 登记支出")
		fmt.Println("           4 退出软件")
		fmt.Print("请选择(1-4):")

		fmt.Scanln(&key)
		switch key {
		case "1": //收支明细
			fmt.Println("------------当前收支明细--------------")
			if flag {
				fmt.Println(details)

			} else {
				fmt.Println("  当前没有收支 请来一笔吧!")
			}

		case "2": //登记收入
			fmt.Println("本次收入金额:")
			fmt.Scanln(&money)
			balace += money
			flag = true
			fmt.Println("本次收入说明")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收支\t%v\t\t+%v\t\t%v", balace, money, note)
		case "3": //3登记支出
			fmt.Println("本次支出金额:")
			fmt.Scanln(&money)
			if money > balace {
				fmt.Println("余额不足!")
				break
			}
			balace -= money
			flag = true
			fmt.Println("本次支出说明")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收支\t%v\t\t-%v\t\t%v", balace, money, note)
		case "4": //退出软件
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
				loop = false
			}

		default:
			fmt.Println("请输入正确的现象")
		}
		if !loop {
			break
		}
	}

	fmt.Println("你退出了家庭记账软件的使用。。")
}
