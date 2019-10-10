package main

import (
	"fmt"
	"go_code/failyAccount/account2/utils"
)

func main() {
	fmt.Println("这是面向对象的方式实例完成~~")
	utils.NewFamilyAccount().MainMenu()
}
