参考：https://blog.zxysilent.com/archives/ 	#2018年07月 (17)

我的scdn博客地址：
	https://blog.csdn.net/u011620855

beego框架：		// 官网： https://beego.me/  
	下载：	go get github.com/astaxie/beego
	安装bee 工具：	go get github.com/beego/bee
	//设置环境变量path  Golang安装目录的bin目录和 GOPATH --src的同级目录 以及E:\code\Go\bin
	创建项目：
		$GOPATH/src/ 	bee new liteblog
		cd liteblog
		bee run 	//启动项目  浏览器访问 http://127.0.0.1:8080/
	路由：	
		https://beego.me/docs/mvc/controller/router.md 
		# 注解路由：
	MVC：
	orm：
		go get github.com/astaxie/beego/orm
		orm生成表：
			package models
			import (
				"github.com/astaxie/beego/orm"
				_"github.com/go-sql-driver/mysql"	//mysql 驱动
			)
			type User struct {
				Id int
				Name string
				Pwd string
			}
			func init()  {
				//设置数据库基本信息
				orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/liteblog?charset=utf8")
				//指定model映射表
				orm.RegisterModel(new(User))
				//生成表	//name=>default [别名] force=>false[是否强制更新表结构] verbose=>true[是否可见创建sql过程]
				orm.RunSyncdb("default",false,true)
			}



