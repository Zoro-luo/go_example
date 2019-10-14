格式化：
	gofmt -w hello.go
输入输出fmt:
		var name string 
		fmt.Println("请输入姓名")	//输出
		fmt.Scanln(&name)			//输入
		Sprintf:	接受打印后的字符串赋给一个变量存储
			dataStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d \n",time.Now().Year(),.....)
			fmt.Println("dataStr=%v",dataStr)
进制*：
	(1) 其他进制转十进制：
		a. 二进制转十进制：
			:: 从最低位开始(右边),将每个位上的数据提取出来,乘以2的(位数-1)次方,然后求和
				1011 => 1*2零次方 + 1*2一次方 + 0*2二次方 + 1*2三次方 => 1+2+0+8 = 11
		b. 八进制转十进制:
			:: 从最低位开始(右边),将每个位上的数据提取出来,乘以8的(位数-1)次方,然后求和
				0123 => 3*8零次方 + 2*8一次方 + 1*8二次方 + 0*8三次方 => 3+16+64+0 = 83
		c. 十六进制转十进制：
			:: 从最低位开始(右边),将每个位上的数据提取出来,乘以16的(位数-1)次方,然后求和
				0x34A => 10*16零次方 + 4*16一次方 + 3*16二	次方 => 10 + 4*16 + 3*16*16 = 842
	(2) 十进制转其他进制：
		a. 十进制转二进制：
			:: 将该数不断除以2,直到商为0为止,然后将每部得到的余数倒过来,就是对应的二进制
				56 => 56/2余0 28/2余0 14/2余0 7/2余1 3/2余1 余1 => 111000
		b. 十进制转八进制 (八进制以0开头)
			:: 将该数不断除以8,直到商为0为止,然后将每部得到的余数倒过来,就是对应的八进制
				156 => 156/8余
		c. 十进制转十六进制： (十六进制以0x开头)
			:: 将该数不断除以16,直到商为0为止,然后将每部得到的余数倒过来,就是对应的十六进制
				356 => 356/16余4 22/16余6 余1 => 0x164
	(3) 二进制转八进制和十六进制：
		a. 二进制转八进制：
			:: 将二进制数每三位一组(从低位开始组合),转换成对应的八进制数即可 0开头倒转
				11010101 => 101转8#5 010转8#2 11转8#3 =>0325
		b. 二进制转十六进制：
			:: 将二进制数每四位一组(从低位开始组合),转换成对应的十六进制数即可 0x开头倒转
				11010101 => 0101转16#5 1101转16#13(D) => 0xD5
	(4) 八进制和十六进制转二进制:
		a. 八进制转二进制:
			:: 将八进制的每1位,转为对应的一个3位的二进制即可 倒转
				0237 => 7转2#111 3转2#011 2转2#10 => 10011111
		b. 十六进制转二进制:
			:: 将十六进制的每1位,转为对应的一个4位的二进制即可 倒转
				0x237 => 7转2#0111 3转2#	0011 2转2#10 => 1000110111
位运算*：
	(1) 二进制的三个重要概念：原码，反码，补码
		a. 对于有符号的而言:
			1) 二进制的最高位是符号位： 0 表示正数, 1 表示负数 
				1 ==> [0000 0001]	-1 ==> [1000 0001]
			2) 正数的原码,反码,补码都一样
			3) 负数的反码 = 它的源码符号位不变,其他位取反(0->1,1->0)
				1 ==> 原码[0000 0001] 反码[0000 0001] 补码[0000 0001]
				-1 ==> 原码[1000 0001] 反码[1111 1110] 补码[1111 1111] 
			4) 负数的补码 = 它的反码+1
			5) 0的反码,补码都是0
			6) *在计算机运算的时候,都是以补码的方式来运算  
					1-1 => 1 + (-1)
	(2) go中有三个位运算 按位与(&)，按位或(|)，按位异或(^)
			a. 按位与& 	 ：	 两位全部为1,结果为1,否则为0
				2&3 ==> 2的补码[0000 0010] 3的补码[0000 0011]		2&3 => [0000 0010] ->2
			b. 按位或|	 :	 两位只要有一个为1,结果为1,否则为0
				2|3 ==> 2的补码[0000 0010] 3的补码[0000 0011] 	2|3 => [0000 0011] ->3
			c. 按位异或^  :   两位相同取0,不相同取1
				2^3 ==>	2的补码[0000 0010] 3的补码[0000 0011]		2^3 => [0000 0001] ->1
				-2^2 ==> -2的补码->原码[1000 0010]->反码[1111 1101]->补码[1111 1110]
						  2的补码[0000 0010]		
						  	-2^2 => [1111 1100](补码)->反码[1111 1011](补码-1)->原码[1000 0100](符号位不变其他取反) ->-4
	(3) go中的2个移位运算符
			a. >> 右移运算符		:: 低位溢出,符号位不变,并用符号位补溢出的高位
				1>>2 ==> 1的补码[0000 0001] =>[0000 0000] = 0
			b. << 左移运算符		:: 符号位不变,低位补0
				1<<2 ==> 1补码[0000 0001] => [0000 0100] = 4

			
变量*的使用方法：
	(1) 变量作用域：
		a. 函数内申明/定义的变量交局部变量,作用域仅限于函数内部
		b. 函数外部申明/定义的变量交全局变量,作用域在整个包都使用,如果其首字母大写,则作用域在整个程序都有效
		c. 如果变量是在一个代码块,比如for/if中,那么该变量的作用域就是在该代码块中
	(2) 变量的使用方法：
		var name = "tom"
		Name := "jack"	//会报错,因为函数外部不能做赋值操作
		func main(){
			//golang 变量使用
			//1. 指定变量类型 声明后不赋值 使用默认值 
			var i int
			fmt.Println("i=",i)
			//2. 根据数值类型判断（类型推导）
			var num = 10.22
			fmt.Println("num=",num)
			//3. 省略var 使用 :=  左侧的变量不应该是已经申明过的 否则编译会导致错误
			// 等价于  var name string     name := "tom"
			name := "tom"
			fmt.Println("name=",name)
		}
	(3) 多变量申明：
		func main(){
			//方式1
			var n1, n2, n3 int
			fmt.Println("n1=",n1,"n2=",n2,"n3=",n3)
			//方式2
			var n1, name ,ne = 100,"tom",888
			fmt.Println("n1=",n1,"name=",name,"ne=",ne)
		}
	(4) 判断类型：
		func main(){
			var num = 10
			//fmt.Printf()可以做格式化输出
			fmt.Printf("num 的类型 %T", num)
			// 如何在程序查看某个变量的占用字节大小和数据类型(使用较多)
			var n1 int64 = 10
			// unsafe.Sizeof(n1) 是unsafe包的一个函数 可以返回n1变量占用的字节数
			fmt.Printf("n1 的类型 T%  n1占用的字节数是 d%"，n1, unsafe.Sizeof(n1))
		}
	(5) 基本数据类型转换：
		func main(){
			// 将i => float
			var i int32 = 100
			var n1 float32 = float32(i)
			// %T 查看类型
			fmt.Printf("n1 type is %T\n",n1)
			//被转换的是变量存储的数值 变量本身的数据类型是不会改变的
			fmt.Printf("i type is %T\n",i)
			//在转换中 如将 int64 转成int8【-128---127】编译时不会报错
			//只是转换的结果是按溢出处理, 和我们希望的结果不一样 
			var num1 int64 = 999999
			var num2 int8 = int8(num1)
			fmt.Println("num2 =",num2)
		}
常量*：
	常量的介绍：
		a. 常量必须在申明时候初始化
		b. 常量不能修改
		c. 常量只能修饰bool，数值类型(int，float系列)，string类型
		d. 常量没有硬性规定常量必须大写字母
		e. 依旧通过首字母大小写来控制常量的访问范围
	常量的注意事项：
		a. 比较简洁的写法：
			const (
				a = 1
				b = 2
			)
			fmt.Println(a,b)
		b. 还有一种专业的写法
			const (
				a = iota 	// 0
				b 			// 1
				c 			// 2
			)
字符串*：
	(1) 改变字符串：
		func modifystring(){
			s := "我hello word"
			s1 := []rune(s)		//字符串强转为数组
			s1[0] = 'a'
			str := string(s1)	//数组转为字符串
			fmt.Println(str)
		}
	(2) 字符串常用的系统函数
		1. len()	=>	func len(v type) int
			example：
				len("aaa北")		//返回6 一个汉字占3个字节
		2. []rune(str)：字符串遍历,同时处理有中文的问题 res := []rune(str)	//吧字符串转rune的切片、
			example：
				str := "hello北京"
				r := []rune(str)	
				for i=0; i<len(r);i++{ fmt.Printf("字符=%c\n",r[i]) }
		3. strconv.Atoi：字符串转整数：		n,err := strconv.Atoi("12")
			example：
				n,err := strconv.Atoi("11")
				if err != nil {
					fmt.Println("转换错误：",err)
				}else{
					fmt.Println("转换结果是：",n)
				}
		4. strconv.Itoa： 整数转字符串：		str := strconv.Itoa(1234)
				fmt.Printf("str=%v, 类型是 %T",str,str)
		5. 字符串转[]byte ：	 var bytes = []byte("abc")//
				fmt.Printf("bytes=%v\n",bytes)		//打印出 bytes=[97 98 99]
		6. []byte 转字符串：	str := string([]byte{97,98,99})
				fmt.Printf("str=%v\n",str)			//打印出 str=abc
		7. strconv.FormatInt：1O进制转 2,8,16进制： str := strconv.FormatInt(132,2)	//第一个参数是传入的整数, 第二个参数是要转的进制(2,8,16) 返回字符串
		8. strings.Contains：查找子串是否在指定字符串中:
				strings.Contains("abc","b")	//查找在第一个参数中是否存在第二个参数的子串 返回bool型
		9. strings.Count：统计一个字符串有几个指定的子串：
				res := strings.Count("aabbbcc","b")	//3
		10. strings.EqualFold：不区分大小写的字符串比较(==是区分字母大小写)：
				strings.EqualFold("abc","ABC")	//true
		11. strings.Index：返回子串在字符串第一次出现的index值,如果没有返回-1：
				strings.Index("NLT_abc","abc")	//4
		12.	strings.LastIndex： 返回子串在字符串中最后一次出现的index,没有返回-1
				strings.LastIndex:("ttt_abc","abc") //
		13. strings.Replace("go go llo","go","go语言",n) //n可以指定你希望替换几个,如果n=-1表示全部替换
		14. strings.Split("hello,word,ok",",")		//按照指定的某个字符,为分割标识,将一个字符串分割返回一个切片
		15. 字符串大小写转换:
				strings.ToLower("Go")	//go
				strings.ToUpper("Go") 	//GO
		16.	a.将字符串左右两边的空格去掉：	strinngs.TrimSpace("yyy hee ")
				res := strings.TrimSpace("yyy hhh   ")  
				fmt.Printf("res=%q\n",res)	//%q 字符串格式 双引号格式
			b.将字符串左右两边的指定字符去掉 ： strings.Trim("!hello!"," !")	//[hello] 去掉了左右两边的！和" "
			c.将字符串左边指定字符去掉： strings.TrimLeft("!hello!","!")
		17 a.判断字符串以指定字符串开头：  strings.HasPrefix("ftp://192.128.1.1","ftp")	//true
		   b.判断字符串以指定	字符串结束： strings.HasSuffix("abc.jpg","abc")	//false


指针*：
	（1）普通类型,变量存的就是值,也叫值类型
	 (2) 获取变量地址 用&, 比如： var a int , 获取a的地址：&a
	 (3) 指针类型 （*int）, 变量存的是一个地址 这个地址存的才是值
	 (4) 获取指针类型所指向的值,使用: *, 比如: var *p int, 使用*p 获取P指向的值
		func main(){
			var a int = 5
			var p *int =&a
			fmt.Println(p)	//地址：0x12454078
			fmt.Println(*p)	//指针所指向的值： 5
		}
时间*日期*函数：
	(1) 时间和日期函数,需要导入time包 	  import "time"
		a. 格式化的第一种方式：
			now := time.Now() 		//返回当前时间戳
			fmt.Printf("当前年月日时分秒：%d-%d-%d %d:%d:%d\n",now.Year(),now.Month(),now.Day(),
				now.Hour(),now.Minute(),now.Sencond())
		b. 格式化的第二种方式
			fmt.Printf(now.Format("2006/01/02 15:04:05"))
			fmt.Printf(now.Format("2006-01-02"))	//返回当前的年月日
			fmt.Printf(now.Format("15:04:05"))		//返回当前的时分秒
			fmt.Printf(now.Format("15"))			//返回当前的小时数值
			* "2006/01/02 15:04:05" 这个字符串的各个数字是固定的,必须是这样写
		(2) 时间常量：
			const(
				    Nanosecond  Duration = 1					//纳秒
				    Microsecond          = 1000 * Nanosecond	//微妙
				    Millisecond          = 1000 * Microsecond	//毫秒
				    Second               = 1000 * Millisecond	//秒
				    Minute               = 60 * Sencond			//分钟
				    Hour                 = 60 * Minute			//小时
				)
			常量的作用：在程序中可用于获取指定时间单位的时间，比如像得到100毫秒 100*time.Microsecond
流程控制*：
	（1） 从终端出入字符串 转为整形 否则输出错误信息
		func main(){
			var str string 
			fmt.Scanf("%s",&str)
			number,err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("convert faild,err:",err)
				return
			}
			fmt.Println(number)
		}	
	（2）Switch：		//没有break
		func main(){
			a := 10
			switch a {
			case 0,1,2,3,4,5:	//满足其中一个都走此分支
				fmt.Println("a is equal 0")
				fallthrough		//继续走下一个分支
			case 10:
				fmt.Println("a is equal 10")
			default:
				fmt.Println("a is equal default")		
			}
		}
		//猜字游戏：
			import("fmt" "math/rand")
			func main(){
				n ：= rand.Intn(100)
				for{
					var input int
					fmt.Scanf("%d\n",&input)	///Scanf 可变需传入地址
					flag := false
					switch  {
						case input == n:
							fmt.Println("you are ok")
							flag = true
						case input > n:
							fmt.Println("you are bigger")
						case input < n:
							fmt.Println("you are less")
					}
					if flag { break }
				}
			}	
	 (3) For 循环：
		a. 常规用法：
			func Printa(n int){
				for i := 0; i <=n; i++ {
					for j := 0; j < i; j++ {
						fmt.Printf("A")		//Printf 格式化打印A (横排)
					}	
					fmt.Println()	//换行
				}
			}
		b. For range 语句：
			str := "hello world 中国"	
			for i, v := range str {	//i=>key   v=>value
				fmt.Printf("index[%d] val[%c] len[%d]\n",i,v,len([]byte(v)))	
			}
		c. For true语句（死循环）：
			for true { fmt.Println(i) }	
			for { fmt.Println(i) }	// 条件true 可省略（等效于上面）
		d. goto 和 label语句
				func main(){
					LABEL:
					HERE:
					for i := 0; i < 6; i++ {
						if i ==2 {
							continue LABEL 		// 当满足条件后 跳到循环外部的LABEL 位置 作用域是一个函数内	
							goto HERE			// 当满足条件后 跳到循环外部的HERE 位置 作用域是一个函数内
						}
					}
				}
包*：go的每一个文件都属于一个包,也就是说Go是以包的形式来管理文件和项目目录结构的
	(1) 包的三大作用：
		a. 区分相同名字的函数,变量等标识符
		b. 当程序文件过多时,可以很好的管理项目
		c. 控制函数 变量等访问范围,即作用域
			打包基本语法：	package 包名
			引入包语法：		import "包名"
	(2) 包的注意事项：
		a. 在给一个文件打包时 该包对应一个文件夹.比如utils文件夹对应的包名就是utils,文件的包名通常和文件所在的文件夹名一致，一般为小写字母
		b. 当一个文件要使用其他包的函数或变量时,需要引入对应的包
		c. package在文件第一行,然后是import指令
		d. 在inport包时,路径从$GOPATH的src下开始,不用带src,编译器会自动从src下开始引入
		e. 为了让其他包的文件,可以访问到本包的函数,则改函数的首字母必须大写,类似其他语言的Public,这样才能跨包访问
		f. 在访问其他包函数时,其语法是： 包名.函数名. 比如：Utils.Cal()
		g. 如果包名过长,Go支持给包取别名. 注意细节：取别名后,原来的包名就不能用了
		h. *如果需要编译成一个可执行的文件，就需要将这个包声明为main,即package main. 这就是语法规范.如果写一个库文件,包名可以自定义
		i. *项目目录结构在$GOPATH的src下. 编译需进入在main包的目录下,
				编译指令：比如：E:\code\Go> go build -o bin/my.exe go_code/day2/code/demo/main //在GOPATH目录下生成一个存编译文件的目录
	(3) 调用外部包和函数
			package main
			import(
				"fmt"
				"go_code/chaptes/function1/utils"		//引入包路径 (GOPATG的src下路径开始不包含src)
				//ut "go_code/chaptes/function1/utils"	//取别名
			)
			return := utils.Cal(n1,n2,operator)			//调用外部函数 需导入包名和函数名(函数名需首字母大写->该函数可导出)
函数*：
	(1) 函数注意事项和细节：
		a. 函数的形参列表可以是多个,返回值列表页可以是多个
		b. 形参列表和返回值列表的数据类型可以是值类型也可以是引用类型
		c. 函数的命名遵循标识符命名规范，首字母不能是数字,首字母是大写的函数可以被本包文件和其它包文件使用,类似Public,
			首字母小写只能被本包文件使用,类型private
		d. 函数中的变量是局部的,函数外部生效
		e. 基本数据类型和数组默认是值传递的,即进行值拷贝. 在函数内修改,不会影响到原来的值
		f. 如何希望函数内的变量能改变函数外的变量,可以传递变量的地址&,函数内以指针的方式操作变量
		g. Go函数不支持重载
	(2) 函数在内存分布说明：
		a. 在调用一个函数时,会给函数分配一个新的空间(栈区),编译器会自身处理让这个新的空间和其它的栈的空间区分
		b. 在每个函数对应的栈中,数据空间是独立的,不会混淆
		c. 当一个函数调用完毕(执行完毕)后,程序会销毁这个函数的栈空间
	   1) 函数也是一种数据类型 可以赋值给一个变量,改变量就是一个函数类型的变量，通过该变量可以对函数调用
	 		func add(a,b int) int { return a+b }	
	 		func main(){
	 			c := add			// add函数赋值变量c
	 			sum := c(100,200)
	 			fmt.Println(sum)
	 		}
	   2) 函数既是一种数据类型,因此在go中函数可以做为形参，并且调用：
	   		基本语法： type 自定义数据类型 数据类型	//理解：相当于一个别名
	   			案列： type myint int 				//这时myint等价于int来使用
	   			案列： type mySum func(int,int)int 	//这时mySum等价于一个函数类型
		  		type op_func func(int,int) int		//自定义类型 2个参数 1个返回值[函数也是一种类型] 相当于一个别名
		  		func add(a,b int) int {	return a+b }
				func operator(op op_func, a, b int) int {	//op 变量名,类型是上面自定义的op_func 参数是上面定义的两个int型
					return op(a,b)							//上面 op_func 类型 接受的2个参数
				}
				func main(){
					c := add								//add 函数赋值给变量c
					sum := operator(c,100,200)				// c函数当参数引用
					fmt.Println(sum)
				}
	   3）*函数传参的方法： 值传递和引用传递
	 		注意：a. 无论是值传递还是引用传递 传递给函数的都是变量的副本 不过,值传递是值得拷贝。引用传递是地址的拷贝。一般来说,地址拷贝更为高效,
	 			 因为数据量小,而值拷贝取决于拷贝的对象大小,对象越大,则性能越低。
	 			 b. *值类型： 基本数据类型 int系列,float系列,bool,string,数组和结构体struct 
	 			 b. *引用类型：map，slice切片，管道chan，指针,interface，默认都是引用的方式传递 
	   4) 函数返回值命名：
	  		func cale(a,b int) (sum int, avg int) {
	  			sum = a+b
	  			avg = (a+b)/2
	  			return
	  		}
	   5) _标识符 , 用来忽略返回值
	  		func main(){ sum,_ := cale(100,200) }
	   6) 支持可变参数：（可变参数要放在形参列表最后）
	  		func add(a int, arg... int) int {}	//1个或多个参数
	  		//注意： 其中 arg是一个slice 我们可以通过arg[index]依次访问所有参数,通过len(arg)来判断传递参数的个数		
	   7) defer* ：
	   		(1) 在函数中,程序员经常需要创建资源(比如数据库连接,文件句柄,锁等),为了在函数执行完毕后及时释放资源,Go的设计者提供了defer(延时机制)
	   		(2) 当函数执行defer时,暂时不会执行,会将defer后面的语句压入到独立的栈(defer栈)，当函数执行完毕后,再从栈按先入后出的方式出栈 
		  		a. 当函数返回时 执行defer语句. 因此,可以用来做资源清理
		  		b. 多个defer语句 按先进后出的方式执行
		  		c. defer语句中的变量，在defer声明时就决定了
		  			example1:
		  				func a(){
		  					i := 0
		  					defer fmt.Println(i)	//当函数返回结果之后才执行(定义一个压入一个栈里[先进后出])
		  					i= 10
		  					fmt.Println(i)
		  				}
		  	(3) 
	   8) 匿名函数*：
	   		a. 如果某个函数只是希望使用一次,可以考虑使用匿名函数,匿名函数也可以实现多次调用
	   				func test(a,b int) int {
			  			result := func(a1,b1 int) int {	return a1+b1 }(a,b)		//(a,b)体现了调用
			  			return result
			  		}
	   		b. 将匿名函数赋给一个变量(函数变量),再通过该变量来调用匿名函数
	   				a := func(n1 int, n2 int) int { return n1-n2 }
	   				res := a(10,20)
	   		c. 全局匿名函数的使用
	   				var (	//申明
	   					Fun1 = func (n1 int,n2 int) int { return n1*n2 }
	   				)
	   				res ：= Fun1(2,5)	//调用全局匿名函数

	  	9) init 函数 
	  		a. 如果一个文件同时包含全局变量定义，init函数和main函数,则执行流程顺序是： 变量定义->init函数->main函数
	  		b. init函数主要作用是完成一些初始化的工作

内置函数*：
	(1) close:用来关闭channel
	(2) len:用来求长度, 比如string,array,slice,map,channel
	(3) new*: 用来分配内存,主要用来分配值类型,比如int,struct,返回的是指针
			j := new(int) 	
				fmt.Printf("j的类型%T, j的值=%v，j的地址%v,j这个指针指向的值=%v\n",j,j,&j,*j)	
				//j的类型*int, j的值=0xc000062090,j的地址0xc00008c018,j这个指针指向的值=0
				*j = 100	
			fmt.Println(*j)	//地址内的值 100
	(4) make*: 用来分配内存,主要用来分配引用类型,比如chan,map,slice,返回slice本身
			s2 := make([]int,5)
			fmt.Println(s2)			// [0 0 0 0 0]
	(5) append： 用来追加元素到数组,slice中
			var a []int	// [5] --> 数组  [] -->切片slice
			a = append(a,10,20,33)
			a = append(a,a...)	//a... 相当于 展开上面的数组[10,20,33]
			fmt.Println(a)
	(6) panic和recover: 用来做错误处理
			func test(){
				defer func(){							//defer return 结束后调用
					err := recover(); 					//recover() 内置函数,可以捕获异常
					if err ！= nil {						//nil 表示没有错误
						fmt.Println("err=",err)		
						fmt.Println("发送邮箱给admin@163.com")
					}
				}()										//申明匿名函数并调用
				b := 0
				a := 100/b
				fmt.Println(a)
				return
			}
			func main(){
				for {
					test()
					time.Sleep(time.Second)
				}
			}
递归函数*： =一个函数在体内调用了本身,则称为递归调用
		(1) 递归函数的重要原则：
			a. 执行一个函数时,就会创建一个新的受保护的独立空间
			b. 函数的局部变量是独立的，不会相互影响
			c. 递归必须向退出递归的条件逼近，否则就是无限递归
		func test(n int){
			fmt.Println("hello word")
			time.Sleep(time.Second)
			if n>10 {	//出口条件 跳出循环
				return 
			}
			test(n+1)	
		}
闭包*： 一个函数和与其相关的引用环境组合而成的实体
			//累加器
			//a.返回的是一个匿名函数,但是这个匿名函数引用到函数外的x,因此这个匿名函数和x形成一个整体,构成闭包
			//b.可以这样理解：闭包是类,匿名函数是操作(方法),x是字段(属性)
			//c.我们反复调用f函数时,因为a是初始化一次,故每次调用一次就进行累计
			//d.我们要搞清楚闭包的关键,就是分析出返回函数它使用(引用)到哪些变量,因为函数和它引用到的变量共同构成闭包
			func Adder() func(int) int {	//返回值是一个函数类型
				var x int 					//定义局部变量x
				return func(d int) int {	//返回的闭包函数 和上面的Adder函数返回函数一一对应
					x += d 					//外部引用变量x会保留上次的值
					return x 
				}
			}
			func main(){
				var f := Adder()
				fmt.Print(f(1),"--")
				fmt.Print(f(20),"--")
			}
			//输入一个文件名,有后缀名就返回原文件名,没有后缀则返回带后缀名的文件名
			func makeSuffix(suffix string ) func (string) string {
				return func (name string) string {
					if !strings.HasSuffix(name,suffix) {
						return name+suffix
					}
					return name
				}
			}
			f := makeSuffix()
			fmt.Print(f("aa.jpg"),"--")
数组*：
	(1) 数组内存中分配：
		a. 数组的地址可以通过数组名来获取&Arr 
		b. 数组的第一个元素的地址就是数组的首地址  
		c. 数组的各个元素的地址间隔是依据数组类型所占字节数递增而列 比如int 8字节 int32 4字节
	(2) 数组使用注意事项和细节(重要)：
		1. 使用数组步骤：1 申明数组并开辟空间 2 给数组各个元素赋值 3 使用数组
		1. 数组是多个相同类型数据的组合,一个数组一旦申明/定义了 其长度是固定的,不能动态变化
		2. var arr[]int 这时arr 是一个slice切片 
		3. 数组中的元素可以是任意数据类型,包括值类型和引用类型,但是不能混用
		4. 创建数组后 如果没有赋值 则有默认值 数组类型默认值0，字符串数组默认值"", bool数组默认值false
		5. 数组属于值类型,默认是值传递,因此会进行值拷贝.(数组间不会相互影响)
	(3) 数组定义：   数组是值类型
		a. var a[len] int  比如：var a[5] int   * 一旦定义,长度不能变(默认值是0)
		b. 长度是数组类型的一部分,因此 var a[6] int 和 var a[10] int 是不同的类型
		c. 数组可以通过下标进行访问,下标从0开始,最后一个元素下标是: len-1
			for i := 0; i < len(a); i++ {}
			for index,value := range a {}
				1. index是数组的下标,value是该下标位置的值
				2. 它们都仅是在for循环内部可见的局部变量
				3. 遍历数组元素的时候,如果不想使用下标index,可以直接把下标index标为下划线_
		d. 访问越界,如果下标在数组合法范围外,则会触发访问越界,会panic
		e. 数组是值的类型，因此改变副本的值,不会改变本身的值
			var a [5] int
			b := a 				//b 是 a 的副本
			b[0] = 100			//改变数组b的值
			fmt.Println(a)		//a的值不变 [0 0 0 0 0]
			fmt.Println(b)		//b的值会变 [100 0 0 0 0]
	(4) 初始化：
		a. var arr1 [5]int = [5]int{1,2,3}
		b. var arr2 = [5]int{1,2,3,4,5}
		c. var arr3 = [...]int{1,2,3,4,5,6}			//[...] 让系统自己判断数组的列个数
		d. var str = [5]string{3:"hello",4:"tom"}	// 可以指定元素数组的下表来初始化数组的值 第4个元素值是hello
	(5) 多维数组：
		a. var age [5][3]int
		b. var f [2][3]int = [...][3]int{{1,2,3},{7,8,9}}	//2个子数组,每个子数组有3个值
切片*：   
	(1) 切片是数组的一个引用,因此切片是引用类型
	(2) 切片的长度可以改变,因此切片是一个可变的数组
	(3) 切片遍历方式和数组一样,长度用len()求长度
	(4) cap可以求出slice的最大容量, 0<=len(slice)<=cap(array) 其中array是slice引用的数组
	(4.1) 从slice底层来说,其实切片就是一个数据结构(struct结构体)
		a. 切边内存包含3个空间 引用到数据的那个下标的指针，len ， cap

	(5) 切片的定义： var 变量名 []类型, 比如：  var str []string  var arr []int
		*切片使用的三种方式：
			1. 定义一个切片,然后让切片去引用一个创建好的数组
				var arr[5]int = [...]int{1,2,3,4,5}	
				var slice []int
				slice = arr[2:3]		//slice是切片名  arr[2:3] 表示slice引用到arr这个数组 从下标2到3结束但不包含3
				fmt.Printf("切片的地址%p\n",slice)			//输出：切片的地址0x120100bc
				fmt.Printf("数组第三个元素地址%p\n",&a[2])	//输出：数组第三个元素地址0x120100bc		
			2. 通过make来创建切片
				 var slice []int = make([]int,4,10)			//make来创建一个切片	（slice := make([]type,len,cap)）.
				 fmt.Println(slice)				//默认值是0 
				 fmt.Printf("len=%v cap=%v 地址=%p",len(slice),cap(slice),&slice) //len=4 cap=10 地址=0xc00005a2c0
				 slice[2] = 200
			总结:	a. 通过make方式创建切片可以指定切片的大小和容量
				  	b. 如果没有给切片赋值,那么就胡使用默认值[int float =>0 string=>"" bool=>false]
				  	c. 通过make方式创建的切片对应的数组是由make底层维护,对外不可见,即只能通过slice去访问各个元素
			3. 定义一个切片,直接就指定具体数组,使用原理类型make的方式
				var slice []int = []int {1,3,5}
				fmt.Println(slice)		
	(6) 切片扩容(追加)
			用append内置函数,可以对切片进行动态追加
				1. 通过append 直接给slice追加具体元素
					var slice []int = []int{100,200}
					slice = append(slice,300,400)
					fmt.Println(slice)	//[100,200,300,400]
				2. 通过append将切片slice追加给slice
					slice = append(slice,slice...)	//固定写法
			切片append操作的底层原理分析： 	
				1. 切片append操作本质就是对数组扩容
				2. go底层会创建一个新的数组newArr(扩容后的大小)
				3. 将slice原来包含的元素拷贝到新的数组newArr
				4. slice重新引用到newArr
				5. 注意newArr 是在底层维护的,程序员不可见

	(7) 切片拷贝
		var a[]int  = []int{1,2,3,4,5}	
		b := make([]int, 10)		//make 一个切片
		copy(b,a)
		fmt.Println(a)				//[1 2 3 4 5]
		fmt.Println(b)				//[1 2 3 4 5 0 0 0 0 0]
		说明： copy(para1,para2)	：para1和para2都是切片类型
	(8)string与切片
		str := "hello go"
		s1 := str[0:5]
		s2 := str[5：]

排序*：
	(1) 引入包 import "sort"
	(2) 数组排序：
		a.func arrSort(){
			var arr = [...]int{1,33,5,99,2}
			sort.Ints(a[:])			//因为数组是值类型 所以要传入切片 
			fmt.Println(a)
		}
		b.  var s = [...]string{"as","eee","rrt"}	
			sort.Strings(s[:])		//对字符串进行排序
		c.  var f = [...]float64{0.8,0.2,99,8.2}
			sort.Float64s(f[:])		//对浮点数进行排序
	(3) search 检索
			s := [...]int{3,1,6,5}
			sort.Ints(s[:])						//检索key前先要排序 不然search会不准
			index := sort.SearchInts(s[:],3)	//返回值3的key
map*： key-value的数据结构,又叫字典或关联数组
	(1) 注意： slice,map,还有function不可以用做key， 因为这几个没法用==判断 
	(1) 声明：	(申明是不会分配内存的 初始化需要make)
			var a map[string]string 					//key是字符串value是字符串
			var a map[string]int 						//key是字符串value是int
			var a map[string]map[string]string 			//key是字符串value是map
			注意：map申明是不会分配内存的,初始化需要make,分配内存后才能赋值和使用
			总结：	a. map在使用前一定要make
					b. map的key是不能重复的.map的value可以重复
					c. map的key-value是无序的
	(2) map的使用：
		1. 方式1： 
			var a map[string]string
			a = make(map[string]string,10)
			a["one"] = "张三"
			a["two"] = "李四"
		2. 方式2：
			m := make(map[string]string,10)
			m["no1"] = "北京"
			m["no2"] = "上海"
		3. 方式3：
			a := map[string]string{	//申明初始化赋值
				"a":"qq",
				"b":"ww",
			}
	(3) map的增删改查
		a. map的增加和更新
			map["key"] = value 	//如果key不存在,就是新增,如果key不存在就是修改
		b. map的删除
			m := make(map[string]string,10)
			m["no1"] = "北京"
			m["no2"] = "上海"
			delete(m,"no1")	//删除key是no2的值 delete是内置函数 当指定key不存在,删除操作也不会报错
			1) 如果希望一次性删除所有的key
				a. 遍历所有的key,逐一删除
					*map遍历只能用for-range的结构遍历
					for k,v := range m {
						fmt.Printf("k =%v v=%v\n",k,v)
					}
				b. 直接make一个新的空间
					citis = make(map[string]string)
					fmt.Println(citis)
		c. map的查找
			val, ok := m["no1"]		//val 为map对应key的value , ok 为查找返回结果bool类型
			if ok { fmt.Printf("有no1的key 值为%v\n",val) }
	(4) map切片		[]map
		基本介绍： 切片数据类型如果是map,则我们称为slice of map切片,这样使用则map个数就可以动态变化了
			案例：
				var mons []map[string]string 			//申明一个map切片
				mons = make([]map[string]string,2)		//切片也要make分配内存
				// 给map切片赋值2个
				if mons[0] == nil {
					mons[0] = make(map[string]string,2)	//map赋值钱需要make
					mons[0]["name"] = "张三"
					mons[0]["age"] = "20"
				}
				if mons[1] == nil {
					mons[1] = make(map[string]string,2)
					mons[1]["name"] = "李四"
					mons[1]["age"] = "24"
				}
				// 定义一个新的map
				newMons := map[string]string{ "name" = "王武","age" = "30",}
				mons = append(mons,newMons)
				fmt.Println(mons)
	(5) map排序：
		a. golang中的map默认是无序的,每次遍历得到的输出可能不一样
		b. golang中map的排序,是现将key进行排序,然后根据key值遍历输出即可
			1. 现将map的key放入到切片中,再对切片排序,遍历切片,然后再按key输出map
				map1 := map[int]int{
					10:100,
					3:44,
					7:77,
				}
				var keys []int
				for k,_  := range map1{
					keys = append(keys,k)
				}
				sort.Ints(keys)
				for _,k := range keys{
					fmt.Sprintf("map1[%v]=%v\n",k,map1[k])
				}
	(6)	map的使用细节：
		a. map是引用类型，遵循引用类型的机制,在一个函数接收map，修改后会直接修改原来的map
		b. map的容量达到后,想要map增加元素,会自动扩容,并不会发送panic,也就是说map能自动增长键值对(key-value)
		c. map的value也经常使用struct类型,更适用于管理复杂的数据
锁*：
	（1）线程同步 （sync锁）
			a. 互斥锁, var mu sync.Mutex
			b. 读写锁, var mu sync.RWMutex
 	（2）go get 安装第三方包
 			go get github：//。。。。
结构体*：(和 java/PHP 语言的对象class 类似)
	(1) 结构体特点：
		a. 去掉了传统oop语言的继承,方法重载,构造函数和析构函数,隐藏的this指针等等
		a. 用来定义复杂数据结构
		b. struct里面可以包含多个字段
		c. struct类型可以定义方法，注意和函数的区分
		d. struct类型是值类型,也可以嵌套
		e. Go语言没有class类型,只有struct类型
	(2) 结构体的内存分配机制:
		a. 变量总数分配在内存中的
	(3) 结构体定义：
		a.  结构体申明：
			type 结构体名称(自定义) struct {
				field1 type
				field2 type
			}
			例如： type Student struct{		// 结构体名称如果首字母大写则在其他包和本包使用
						Name string 		// 属性首字母大写也一样可以在其他包使用 （public）
						Age int 			// 属性一般是基本数据类型,数组,也可以是引用类型
				   }
		b. 创建结构体变量和访问结构体字段
			1) 	var stu Student
			2) 	var stu1 Student = Student{"张三",20}
				stu2 := Student{"张三",20}

			2) 	var stu3 *Student = new (Student)	//申明
					//因为stu3是一个指针,因此标准的字段赋值的方式是：
						(*stu3).Name = "小王"		//赋值   [标准的写法]
						(*stu3).Age = 30			//(*stu3).Age = 30 等价于 stu3.Age = 30	原因： go的设计者为了使用方法在底层对
						stu3.Age = 30				// stu3.Age = 30进行处理 会对stu3加上取值运算(*stu3).Age = 30 
			5) 	var stu5 *Student = &Student{"mary",45}	//申明并给属性赋值
			4) 	var stu4 *Student = &Student{}
					// stu4是一个指针
					(*stu4).Age = 11		//标准写法
					stu4.Name = "tom"		//简单写法
					*stu4.Age = 11 			//会报错！ 原因是.的运算优先级会高于*
		c. 结构体初始化：
			1) 		var stu Student
					stu.Name = "jack"
					stu.Age = 20
			2)		var stu2 = Student{
						Age:20,
						Name:"jack",
					}
					fmt.Printf("Name:%p\n",&stu2.Name)			//打印Name的地址
		d. 结构体的使用注意事项和细节：
			1. 结构体type重新定义(相当于取别名),golang认为是新的数据类型，但是相互间可以强转
				type integer init
				func main(){
					var i integer =10
					var j int =20
					//j=i //会报错
					j = integer(i)
				}
			2. 在结构体的每个字段上，可以写一个tag,该tag可以通过反射机制获取,常见的使用场景是序列化和反序列化
					// 序列化 是返回一个字符串json到客户端 
					type Monster struct{			//申明一个结构体 
						Name string `json:"name"`	//`json:"name"` 就是结构体标签 struct tag 作用在反射结果json后可以使Name小写成name
						Age int	`json:"age"`
						Skill string `json:"skill"`
					}
					func main(){
						// 1 创建一个monster变量
						monster := Monster{"牛魔王",200,"芭蕉扇"}
						//2 将moster变量序列化为一个josn字符串
						jsonStr,err := json.Marshal(monster)		//json.Marshal 函数中用到反射
						if err != nil{
							fmt.Println("json 错误处理，",err)
						}
						fmt.Println("jsonStr:",string(jsonStr))
					}
方法*：
	(1) 基本介绍：
		Golang 中的方法是作用在指定的数据结构上的(即：和指定的数据类型绑定),因此自定义类型都可以有方法,而不仅仅是struct
	(2) 方法的申明和调用
			type Person struct{ Name string }
			//给Person这个结构体绑定一个test方法
			func (p Person) test(){		//表示Person这个结构体有一个方法,方法名是test (类型php里的class类里有一个方法)
				fmt.Println(a.Num)
			}
			func main(){
				var p Person 		//申明一个结构体变量
				p.Name = "tom"		//给这个结构体变量赋值
				p.test()			//调用这个结构体里绑定的一个方法
			}
		说明：	a. 	func (p Person) test(){	//表示A这个结构体有一方法,方法名是test
				b.	(p Person) 				//体现 test方法是A类型绑定的	，表示哪个Person变量调用,这个p就是它的副本(类型函数传参)
				c.  p.test()				//不能直接test()调用和用其他类型结构体调用,只能本所绑定的结构体来调用
				d. p是可以自定义的
	(3) *方法的调用和传参机制原理：
		a. 通过一个变量取调用方法时,其调用机制和函数一样
		b. 不一样的是 变量调用方法时, 该变量本身也会作为一个参数传递到方法(如果变量是值类型 则是值拷贝,如果变量是引用类型 则进行地址拷贝)
	(4) 方法注意事项和细节：
		a. 结构体类型是值类型,在方法调用中，遵循值类型的传递机制，是只拷贝传递方式
		b. 如果想在方法中修改结构体变量的值，可以通过结构体指针的方式来处理  
				type Circle struct{ radius float64 }
				func (c *Circle) area() float64 {	// 为了提高效率 通常结构体方法定义是指针类型 
					c.radius = 10.0					//c.radius 等价于 (*c).radius
					return 3.14 * c.radius * c.radius
				}
				func main(){
					var c Circle
					c.radius = 7.0
					res := c.area()			
					fmt.Println(c.radius)	//10
					fmt.Println(res)		
				}
		c. Golang中的方法作用在指定的数据类型上(即:和指定的数据类型绑定)，因此自定义类型都可以有方法，而不仅仅是struct，比如int，float64等都可以有
				type integer int
				func (i *integer) printt() {
					*i = *i + 1
				}
				func main(){
					var i integer = 55
					i.printt()
					fmt.Println(i)		//56	
				}
		d. 方法的访问范围控制规则和函数一样,方法名首字母小写,只能在本包访问,方法名首字母大写,可以在本包和其他包访问
		e. 如果一个类型实现了String()这个方法,那么fmt.Println默认会调用这个变量的String()进行输出
				type Student struct {
					Name string
					Age int
				}
				func (stu *Student) string() string {
					str := fmt.Sprintf("Name=[%v] Age=[%v]",stu.Name,stu.Age)
					return str
				}
				func main(){
					stu := Student{
						Name : "TOM",
						Age : 22,
					}
					fmt.Println(&stu)
				}
面向对象*：
	(1) 面向对象编程实例
		步骤：	申明定义结构体 编写结构体字段 编写结构体方法
			type Student struct{
				name string 
				gender string
				age int
				id int
				score float64
			}
			func (s *Student) say() string {
				inforStr := fmt.Sprintf("student的信息：name=[%v],gender=[%v],age=[%v],id=[%v],score=[%v]",
				s.name,s.gender,s.age,s.id,s.score)
				return inforStr
			}
			func main(){
				//创建一个实例
				var stu = Student{
					name : "tom",
					gender : "male",
					age : 18,
					id : 1000,
					score : 99.0,
				}
				str := stu.say()
				fmt.Println(str)
			}
	(2) 工厂模式：
		说明：Golang的结构体没有构造函数,通常可以使用工厂模式解决这个问题
		看一个需求：
			package model
			type Student struct{
				Name string ...
			}
			详解：	因为这里的Student的首字母S是大写的,如果我们想在其他包创建Student的实例(比如main包),引入model包后，就可以直接创建Student结构
				 	体的实例。但是如果首字母是小写,比如是type student struct{...}就不行了 怎么办--->工厂模式来解决。
		例子：
			student.go:
				package model
				//定义一个结构ti
				type student struct{	//student首字母小写 外部包不能直接访问
					Name string 		//首字母大写 外部包可以直接访问
					score float64 		//首字母小写 外部包不能直接访问
				}
				//定义一个工厂可外部访问的函数
				func NewStudent(n string, s float64) *student {
					return &student{
						Name : n,
						score : s,
					}
				}
				//如果score字段首字母小写,则在其他包不可以直接访问,我们可以提供一个方法在其他包访问小写字母的字段
				func (s *student) GetScore() float64 {
					return s.score
				}
			main.go :
				package main
				import "go_code/day6/code/factory/model"
				func main(){
					//定义的student结构体是首字母小写,我们可以通过工厂模式解决
					var stu = model.NewStudent("tom~~",88.8)
					fmt.Println("Name=",stu.Name , "score=",stu.GetScore())
				}
	(3) 面向对象的三大特性：
		说明：Golang语言也又面向对象的继承,封装，多态的特性，只是实现方式和其他oop语言不一样
		如何理解抽象：定义一个结构体的时候，实际上是把一类事物的共有的属性和方法提出出来，形成一个屋里模型(结构体),这种研究问题的方法称为抽象
		a.封装：就是吧抽象出来的字段和对字段的操作封装在一起,数据被保护在内部，程序的其他包只能通过被授权的方法才能对字段进行操作
			封装的实现步骤：
			1. 将结构体,字段的首字母小写(其他包不能访问，类似private)
			2. 将结构体所在包提供一个工厂模式的函数,首字母大写,类似一个构造函数
			3. 提供一个首字母大写的Set方法(类似其他语言的Public),用于对属性赋值
				func (var 结构体类型名) SetXxx(参数列表) (返回值列表) {
					//数据验证逻辑
					var 字段 = 参数
				}
			4. 提供一个首字母大写的GetXxx方法(类似public) 用于获取属性的值
				func (s *student) GetScore() float64 {
					return s.score
				}
		b. 继承：在Golang中,如果一个struct嵌套了另一个匿名结构体,那么这个结构体可以直接访问匿名结构体的字段和方法,从而实现了继承特性
			基本语法：
				type Goods struct{
					Name string 
					Price int
				}
				type Book struct{	
					Goods			//这里嵌套匿名结构体Goods 相当于继承了Goods结构体的属性和方法
					Writer string 
				}
			继承的深入讨论：
				1) 结构体可以使用嵌套结构体所有的字段和方法. 即:首字母大写或者小写的字段,方法，都可以使用
					type A struct{
						Name string
						age int
					}
					func (a *A) SayOk(){	//A 结构体的方法	SayOK
						fmt.Println("A SayOk",a.Name)
					}
					func (a *A) hello(){	//A 结构体的方法	hello	
						fmt.Println("A hello",a.Name)
					}
					type B struct{			//相当于B 是 A 的子类
						A 			
					}
					func main(){
						var b B 
						b.A.Name = "tom"
						b.A.age = 22
						b.A.SayOk()		//A SayOk tom 	// 简化： b.SayOk()
						b.A.hello()		//A hello tom   //小写的方法也可以调用
					}
				2) 匿名结构体访问可以简化：
					var b B 
					b.a.name = "tom"	===>	b.name= "tom"
					b.a.say()			===>	b.say()
				3) 当结构体和嵌套的匿名结构体有相同的字段或方法时,编译器采用就近访问原则，如果希望访问匿名结构体的字段和方法,
				   可以通过匿名结构体名来区分 即：b.A.hello() 明确是A结构体的方法 b.hello() 则优先访问B结构体的hello方法
				4) 结构体嵌入两个(或多个)匿名结构体,如两个匿名结构体有相同的字段或方法(同时结构体本身没有同名的字段和方法)，在访问时，
					就必须明确指定匿名结构体名字,否则编译报错	
						type A struct{
							Name string 
							Age int
						}
						type B struct{
							Name string
							score int
						}
						type C struct{
							A
							B
							// Name string
						}
						func main(){
							var c C 
							c.Name = "tom" 		//报错 必须 c.A.Name  或 c.B.Name 
						}
				5) 如果一个结构体嵌套了一个有名结构体,这种模式就是组合。如果是组合关系,那么在访问组合的结构体的字段或方法时，
					必须带上结构体名字
						type A struct{
							Name string
						}
						type C struct{
							a A 	//相当于给嵌套的A 结构体取了一个别名
						}
						func main(){
							var c C
							c.a.Name = "jack"	//访问有名结构体字段时 必须带上有名结构体的名字
						}
				6)嵌套匿名结构体后,也可以在创建结构体变量(实例)时,直接指定各个匿名结构体字段的值
						type Good struct{
							Name string 
							Price float64
						}
						type Brand struct{
							Name string
							Address string
						}
						type TV struct{
							Goods
							Brand
						}
						type TV2 struct{
							*Goods 				//也可以嵌套指针类型
							*Brand
						}
						func main(){
							//嵌套匿名结构体后,也可在创建结构体实例时，直接指定各个匿名结构体字段的值
							tv := TV{ Goods{"电视机01",4999.00},Brand{"海尔","山东"}, }
							tv2 := TV{ 
								Goods{
									Name : "电视机02",
									Price : 5999.00,
								},
								Brand{
									Name : "夏普",
									Address : "北京",
								}, 
							}
							tv3 := TV2{ &Goods{"电视机03",6999.00}, &Brand{"创维","湖北"}, }	//指针类型用&
							fmt.Println(tv)
							fmt.Println(tv2)
							fmt.Println("tv3",*tv3.Goods, *tv3.Brand)
						}
接口*(interface)：  多态性主要通过接口来体现
	基本介绍：	interface类型可以定义一组方法,但这些方法不需要实现.且interface不能包含任何变量。到某个自定义类型(比如结构体Phone)要
				使用的时候,根据具体情况把这些方法实现
	基本语法：
				type 接口名 interface{
					method1(参数列表) 返回值列表
					method2(参数列表) 返回值列表
				}
	小结说明：
				1)	接口里的所有方法都没有方法体,即接口的方法都是没有实现的方法。接口体现了程序设计的多态和高内聚低耦合的思想
				2)	Golang中的接口,不需要显式的实现。只要一个变量,含有接口类型的所有方法，那么这个变量就实现了这个接口。因此
					Golang中没有implement这样的关键字
	注意事项和细节：
				1)	接口本身不能创建实例,但是可以指向一个实现了该接口的自定义类型的实例
						type Usb interface{ Start();Stop(); }
						type Phone struct{}
						type Camera struct{}
						func (p Phone) Start(){fmt.Println("手机开始工作..")}
						func (p Phone) Stop(){fmt.Println("手机停止工作..")}
						func (p Camera) Start(){fmt.Println("相机开始工作..")}
						func (p Camera) Stop(){fmt.Println("相机停止工作..")}
						type Computer struct{}
						//编写一个work方法,接受一个Usb接口类型变量 实现了Usb接口
						func (c Computer) Working(usb Usb){
							//通过usb这个接口变量来调用Start和Stop方法
							usb.Start()
							usb.Stop()
						}
						func main(){
							//先创建结构体实例
							computer := Computer{}
							phone := Phone{}
							camera := Camera{}
							computer.Working(phone)	
							computer.Working(camera)
						}
				2)	接口中所有的方法都没有方法体,即都是没有实现的方法 
				3)	在Golang中,一个自定义类型需要将某个接口的所有方法都实现,我们说这个自定义类型实现了该接口
				4)	只要是自定义数据类型,就可以实现接口,不仅仅是结构体类型
				5)	一个自定义类型可以实现多个接口
						type AInterface interface{ Say() }
						type BInterface interface{ Hello() }
						type Monster struct{}		//Monster 结构体实现了AInterface和BInterface
						func (m Monster) Say(){ fmt.Println("Monster Hello()") }
						func (m Monster) Hello(){ fmt.Println("Monster Hello()") }
						func main(){
							var monster Monster
							var a2 AInterface = monster
							var b2 BInterface = monster
							a2.Say()
							b2.Hello()
						}

				6)	Golang接口中不能有任何变量
				7)	一个接口(比如A接口)可以继承多个别的接口(比如B,C接口),这时如果要实现A接口,也必须将B,C接口的方法也全部实现
						type BInterface interface{ test01() }
						type CInterface interface{ test02() }
						type AInterface interface{	//AInterface 继承了 BInterface 和 CInterface
							BInterface
							CInterface
							test03()
						}
						//如果需要实现AInterface，就需要将BInterface和CInterface的方法都实现
						type Stu struct{ }
						//必须要实现所有方法
						func (stu Stu) test01(){}
						func (stu Stu) test02(){}
						func (stu Stu) test03(){}
						func main(){
							var stu Stu
							var a AInterface = stu
							var b BInterface = stu
							var c CInterface = stu
							a.test01()	//a可以调 test01 test02 test03 三个方法
							b.test01()	//b只能调 本身的test01方法
							c.test02()	//c只能调 本身的test02方法
						}
				8)	interface类型默认是一个指针(引用类型)，如果没有interface初始化就使用,那么会输出nil
				9)	空接口interface{} 没有任何方法,所以所有类型都实现了空接口,即我们可以把任何一个变量赋给空接口
	接口编程的最佳实践：
			import (
				"fmt"
				"sort"
				"math/rand"
			)
			//1申明Hero结构体
			type Hero struct{
				Name string
				Age int
			}
			//2申明Hero结构体切片类型
			type HeroSlice []Hero
			//3 实现Interface 接口
			func (hs HeroSlice) Len() int { return len(hs) }
			// Less 方法决定使用什么标准进行排序
			func (hs HeroSlice) Less(i,j int) bool {
				return hs[i].Age < hs[j].Age
				//也可以改为对Name排序
				//return hs[i].Name < hs[j].Name
			}
			func (hs HeroSlice) Swap(i,j int)  { hs[i], hs[j] = hs[j], hs[i] } 	//交换
			func main(){
				var intSlice = []int{0,-1,22,7,99,66}
				sort.Ints(intSlice)		//对切片进行排序
				fmt.Println(intSlice)

				//对结构体切片排序
				var heroes HeroSlice
				for i:=0;i<10;i++{
					hero := Hero{
						Name : fmt.Sprintf("英雄|%d", rand.Intn(100)),
						Age : rand.Intn(100),
					}
					//将hero append 到heroes切片
					heroes = append(heroes,hero)
				}
				//看看排序前的顺序
				for _,v := range heroes {
					fmt.Println(v)
				}
				//调用sort.Sort
				sort.Sort(heroes)
				fmt.Println("排序后~~~~~~~")
				//看看排序后的顺序
				for _,v := range heroes {
					fmt.Println(v)
				}

			}
	实现接口VS 继承：
			继承的价值主要在于：解决代码的复用性和可维护性
			接口的价值主要在于：设计好各种规范,让其他自定义类型取实现这些方法
		接口比继承更灵活
			接口比继承更加灵活,继承是满足 is - a 的关系， 而接口只需满足like - a 的关系
		接口在一定程度上实现代码解耦
	多态：
		基本介绍：实例具有多种形态,面向对象的第三大特征,在Go语言，多态特征是通过接口实现的。可以按统一的接口来调用不同
				 的实现。这时接口实例就呈现不同的形态
	类型断言： 由于接口是一般类型,不知道具体类型,如果要转成具体类型,就需要使用类型断言。
		类型断言判断案例： 如何在进行断言时,带上检测机制,如果成功ok，否则也不要抱panic
				var x interface{}
				var b float32 = 1.1
				x = b //空接口可以接收任意类型
				y,ok := x.(float64)	//x ==>float32 [使用类型断言]
				if ok {
					fmt.Println("convert success...")
					fmt.Printf("y 的类型是%T 值是%v\n",y,y)
				}else{
					fmt.Println("convert fail...")
				}
				fmt.Println("go........")
		类型断言最佳案例： [！！！注意体会]
				type Usb interface{	Start(); Stop();}
				type Phone struct{name string}
				//让Phone 实现 Usb接口的方法 并额外实现一个Call方法
				func (p Phone) Start(){fmt.Println("手机开始工作..")}
				func (p Phone) Stop(){fmt.Println("手机停止工作..")}
				func (p Phone) Call(){fmt.Println("手机开始打电话..")}
				type Camera struct{name string}	
				func (p Camera) Start(){fmt.Println("相机开始工作..")}
				func (p Camera) Stop(){fmt.Println("相机停止工作..")}
				type Computer struct{}			
				//编写一个Working方法,接受一个Usb接口类型变量 通过usb这个接口变量来调用Start和Stop方法
				func (c Computer) Working(usb Usb){
					usb.Start()
					//如果usb是指向Phone的结构体变量,则还需要调用Call方法
					if phone,ok := usb.(Phone); ok {	//类型断言 [!!!注意体会]
						phone.Call()
					}
					usb.Stop()
				}
				func main(){
					var usbArr [3]Usb
					usbArr[0] = Phone{"vivo"}
					usbArr[1] = Phone{"小米"}
					usbArr[2] = Camera{"尼康"}
					var computer Computer
					for _,v := range usbArr {
						computer.Working(v)
						fmt.Println()
					}
				}
文件* ： 文件是数据源(保存数据的地方)	----->  os.File
	基本介绍： 文件在程序中是以流的形式来操作的
		流：		数据在数据源(文件)和程序(内存)之间经历的路径
		输入流:	数据从数据源(文件)到程序(内存)的路径	[读文件]
		输出流:	数据从程序(内存)到数据源(文件)的路径	[写文件]
	a.读文件操作：
		(1) 打开文件和关闭文件：
			func Open(name string) (file *File, err error)
			func (f *File) Close() error
			example：
				import ( "fmt";"os";)
				func main(){
					file , err := os.Open("d:/go.go")	//open file
					if err != nil { fmt.Println("open file err=",err) }
					fmt.Printf("file=%v",file)
					err = file.Close()	//close file
					if err != nil { fmt.Println("close file err=",err)}
				}
		(2) 读取文件内容并显示在终端(带缓冲的方式 [适合大文件] )： 使用os.Open , file.Close ,bufio.NewReader() , reader.ReadString  函数和方法
				import ( "fmt";"os";"bufio";"io"; )
				func main(){
					file , err := os.Open("d:/go.go")			//open file
					if err != nil { fmt.Println("open file err=",err) }
					defer file.Close()							//timeliness close file
					reader := bufio.NewReader(file)				//创建一个Reader 带缓冲的 适合大文件
					for {
						str,err := reader.ReadString('\n')		//每读到一个换行符就结束 
						if err == io.EOF { break }				//io.EOF 代表文件的末尾  //读取到文件末尾不再继续读文件
						fmt.Print(str)
					}
					fmt.Println("文件读取结束...") 
				}
		(3) 读取文件内容并显示在终端(使用ioutil一次性将整个文件读取到内存中) [适合文件不大的情况].相关方法和函数(ioutil.ReadFile)
				import ( "fmt";"io/ioutil";)
				func main(){
					//使用ioutil.ReadFile 一次性将文件读取到位 不需要显示Open 和 Close 文件
					file := "d:/go.go"
					content,err := ioutil.ReadFile(file)
					if err != nil { fmt.Printf("read file err=%v",err)}
					fmt.Printf("%v",string(content))	//[]byte
				}
	b.写文件操作：
		基本介绍：
			 func OpenFile(name string,flag int,perm FileModel)(file *file ,err error) //name打开文件路径 flag文件打开模式 perm权限控制
		(1) 创建一个新文件 写入5句 "hello go" 
				import ( "fmt";"bufio";"os";)
				func main(){
					filePath := "d:/abc.txt"
					file,err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE,0666)
					if err != nil { fmt.Printf("open file err=%v\n",err);return;}
					defer file.Close()			//及时关闭句柄
					str := "hello go\r\n"
					//写入到缓冲 使用带缓冲的*writer
					writer := bufio.NewWriter(file)
					for i := 0;i<5;i++{ writer.WriteString(str) }
					//因为writer是带缓冲的,因此在调用WriteString方法时 其实内容是线写入到缓冲的 所以需要调用Flush方法
					//将缓冲的数据真正写入到文件中,否则文件中会没有数据
					writer.Flush()
				}
		(2)	打开一个存在的文件 将原来的内容覆盖成新的内容 10 句 "你好 祖国！"
				file,err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE,0666) 替换如下，
				file,err := os.OpenFile(filePath, os.O_WRONLY | os.O_TRUNC ,0666)
		(3) 打开一个存在的文件，将原来的内容追加内容 "追加内容"
				file,err := os.OpenFile(filePath, os.O_WRONLY | os.O_TRUNC ,0666) 替换如下，
				file,err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND ,0666)
		(4) 将一个文件的内存写入到另一个文件。注：两个文件都存在		使用 ioutil.ReadFile / ioutil.WriteFile
				import ( "fmt";"io/ioutil";)
				func main(){	//file1Path 数据写到file2Path
					file1Path := "d:/abc.txt"	
					file2Path := "d:/kkk.txt"
					data,err := ioutil.ReadFile(file1Path)
					if err != nil { fmt.Printf("read file err=%v\n",err) return }
					err = ioutil.WriteFile(file2Path,data,0666)
					if err != nil {	fmt.Printf("Write file err=%v",err) }
				}
		(5) 统计一个文件中含有的英文,数字,空格以及其他字符数量
				import ( "fmt";"os";"io";"bufio";)
				type CharCount struct{
					ChCount int 	//英文个数
					NumCount int 	//数字个数
					SpaceCount int 	//空格个数
					OtherCcunt int 	//其他字符个数
				}
				func main(){
					fileName := "d:/abc.txt"
					file,err := os.Open(fileName)
					if err != nil {fmt.Printf("open file err=%v\n",err);return ;}
					defer file.Close()
					var count CharCount
					reader := bufio.NewReader(file)
					for{		//循环读取fileName 内容
						str, err := reader.ReadString('\n')
						if err == io.EOF{ break }
						for _,v := range str{	// 遍历 str 进行统计
							switch {
							case v >='a' && v <= 'z':
								fallthrough
							case v >='A' && v <= 'Z':
								count.ChCount++
							case v ==' ' || v== '\t':
								count.SpaceCount++
							case v =='0' || v== '9':
								count.NumCount++
							default:
								count.OtherCcunt++
							}
						}
					}
					fmt.Printf("英文个数=%v 数字个数=%v 空格个数=%v 其他字符个数=%v",
						count.ChCount,count.NumCount,count.SpaceCount,count.OtherCcunt)

				}
		(6) 编写一段程序,可以获得命令行各个参数  os.Args
			func main(){
				fmt.Println("命令行的参数",len(os.Args))
				for i,v := range os.Args {
					fmt.Printf("Args[%v]=%v\n",i,v)
				}
			}
			go build -o test.exe
			test.exe tom d:/aaa/bbb/init.log 999		//参数以空格分隔放在切片里
				// Args[0]=test.exe
				// Args[1]=tom
				// Args[2]=d:/aaa/bbb/init.log
				// Args[3]=999
		(7) flag* 包来解析命令行参数：
				func main(){
					var user string;var pwd string;var host string;var port int;
					flag.StringVar(&user,"u","","用户名默认为空")
					flag.StringVar(&pwd,"pwd","","密码默认为空")
					flag.StringVar(&host,"h","localhost","主机名默认为localhost")
					flag.IntVar(&port,"port",3306,"端口号默认为3306")
					flag.Parse()	//转换解析
					fmt.Printf("user=%v pwd=%v host=%v port=%v",user,pwd,host,port)
				}
				go build -o test.exe main.go 					//编译文件
				test.exe -pwd 123456 -h 127.0.0.1 -u jack		//user=jack pwd=123456 host=127.0.0.1 port=3306
	c.判断文件是否存在：
		基本介绍：os.Stat() (1)返回错误为nil,说明文件或文件夹存在 (2)返回错误类型使用os.isNotExist()判断为true,说明文件或文件夹不存在
						   (3)返回错误为其他类型,则不确定是否存在
				// 基于以上三个错误类型  自己定义了一个函数
					func PathExists(path string)(bool,error){
						_,err := os.Stat(path)
						if err == nil {	//文件或目录存在
							return true,nil
						}
						if os.isNotExist(err){
							return false,nil
						}
						return false,err
					}
json*序列化：
	(1) 将结构体，map，切片 进行序列化 	 ---> import "encoding/json"
		a.//结构体序列化
			type Monster struct { Name string `json:"name"`;Age int;Birthday string;Sal float64;skill string; }
			func testStruct(){
				monster := Monster{
					Name : "牛魔王",
					Age : 500,
					Birthday : "1880-11-22",
					Sal : 8000.0,
					skill : "风沙走石",
				}
				data,err := json.Marshal(&monster)	// monster 序列化
				if err != nil { fmt.Println("序列化失败",err) }
				fmt.Printf("序列化后的结构%v\n",string(data))
			}
		b. //Map 序列化
			func testMap(){
				var a map[string]interface{}
				a = make(map[string]interface{})
				a["name"]="红孩儿";a["age"]=15;a["address"]="火云洞";
				data,err := json.Marshal(a)
				if err != nil { fmt.Println("序列化错误",err) }
				fmt.Printf("序列化后的结构%v\n",string(data))
			}
		c. //切片序列化
			func testSlice(){
				var slice []map[string]interface{}
				var m1 map[string]interface{}
				m1 = make(map[string]interface{})
				m1["name"] = "jack"; m1["age"] = 10; m1["address"] = "上海";
				slice = append(slice,m1)
				var m2 map[string]interface{}
				m2 = make(map[string]interface{})
				m2["name"] = "tom"; m2["age"] = 19; m2["address"] = [2]string{"长沙","武汉"}
				slice = append(slice,m2)
				data,err := json.Marshal(slice)
				if err != nil { fmt.Println("序列化错误",err) }
				fmt.Printf("序列化后的结构%v\n",string(data))
			}	
	(2) 反序列化：	---->  json.Unmarshal([]byte(str),&monster)
		a. //反序列化为结构体	
			type Monster struct { Name string `json:"name"`;Age int;Birthday string;Sal float64;skill string; }
			func unmarshalStruct() {
				str := "{\"name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"1880-11-22\",\"Sal\":8000}"
				var monster Monster
				err := json.Unmarshal([]byte(str),&monster)
				if err != nil { fmt.Println("反序列化错误",err) }
				fmt.Printf("反序列化后 monster=%v monster.Name=%v\n",monster,monster.Name)
			}
		b. //反序列化为map
			func unmarshalMap() {
				str := "{\"address\":\"火云洞\",\"age\":15,\"name\":\"红孩儿\"}"
				var map1 map[string]interface{}
				//反序列化map的时候不需要make,因为make操作封装到Unmarshal函数
				err := json.Unmarshal([]byte(str),&map1)
				if err != nil { fmt.Println("反序列化错误",err) }
				fmt.Printf("反序列化后 monster=%v \n",map1)
			}
		c. //反序列化为Slice
			func unmarshalSlice() {
				str := "[{\"address\":\"上海\",\"age\":10,\"name\":\"jack\"},"+
				"{\"address\":[\"长沙\",\"武汉\"],\"age\":19,\"name\":\"tom\"}]"
				var slice []map[string]interface{}
				//反序列化Slice的时候不需要make,因为make操作封装到Unmarshal函数
				err := json.Unmarshal([]byte(str),&slice)
				if err != nil { fmt.Println("反序列化错误",err) 
				fmt.Printf("反序列化后 monster=%v \n",slice)
			}
单元测试*
	基本介绍： 	go语言自带一个轻量级的测试框架testing和自带的go test命令来实现单元测试和性能测试。testing框架和其他语言中的测试框架类似,
				可以基于这个框架写针对相应的测试应用,也可以基于框架写相应的压力测试用例,通过单元可是可以解决如下问题：
					1)	确保每个函数都是可运行的,并且结果是正确的
					2)  确保写出来的代码性能是好的
					3)	单元测试能及时的发现程序设计或实现的逻辑错误,是问题早暴露,便于问题的定位解决。而性能测试的重点在于程序设计上的一些问题,让程序在高并发的情况下还能稳定
	单元测试细节：
		1 测试用例文件名必须以 _test.go结尾.比如cal_test.go ,
		2 测试用例函数必须以Test开头,一般来说就是Test+被测试的函数名。比如TestAddUpper
		3 TestAddUpper(t *testing.T)的形参类型必须是 *testing.T
		4 一个测试用例文件中,可以有多个测试用例函数.比如 TestAddUpper TestSub
		5 运行测试用例指令
			a. cmd>go test [如果运行正确,无日志,错误时,会输出日志]
			b. cmd>go test -v [运行正确或是错误,都输出日志]
		6 当出现错误时,可以使用t.Fatalf来格式化输出错误信息,并退出程序
		7 t.Logf 可以输出相应的日志
		8 测试用例函数,并没有放在main函数中,但也执行了,这就是测试用例的方便之处
		9 PASS表示测试用例运行成功,FALL表示测试用例失败
		10 测试单个文件,一定要带上被测试的源文件   go test -v cal_test.go cal.go 		//cal被测试的源文件
		11 测试单个方法							go test -v -test.run TetsAddupper



				
goroutine*
	(1) goroutine基本介绍：
		a. 进程和线程说明：
			1) 进程是程序在操作系统中的一次执行过程,是系统进行资源分配和调度的基本单位
			2) 线程是进程的一个执行实例,是程序执行的最小单元,它是比进程更小的能独立运行的基本单元
			3) 一个进程可以创建和销毁多个线程,同一个进程中的多个线程可以并发执行
			4) 一个程序至少有一个进程 一个进程至少有一个线程
		b. 并发和并行：
			1) 多线程程序在单核上运行,就是并发	---> 多个任务作用在一个cpu,在微观角度,在同一个时间点上其实只有一个任务在执行
			2) 多线程程序在多核上运行,就是并行 ---> 多个任务作用在多个cpu同时进行
	(2) go协程和go的主线程:
		1) go主线程(有程序员称为线程/也可以理解位就是进程):一个Go线程上,可以起多个协程,可以理解,协程就是轻量级线程。
		2) Go协程的特点*：有独立的栈空间，共享程序堆程序，调度由用户控制，协程(goroutine)是轻量级线程 
	(3) goroutine案例：
		1) 主线程是一个物理线程,直接作用在cpu上.是重量级的 非常耗费cpu资源
		2) 协程从主线程开启的,是轻量级的线程,是逻辑态。对资源消耗相对小
		3 Golang的协程机制是重要的特点,开启过多的线程 资源消耗大,这里就凸显Golang在并发上的优势
	(4)	goroutine的调度模型：	MPG模式	
			MPG模式基本介绍：	 M:操作系统主线程(是物理线程)  P: 协程执行需要上下文环境(需要的资源,内存空间等)  G：协程
	(5) 设置Golang 运行的cpu数
			import ( "fmt"; "runtime"; )
			func main(){
				cpuNum := runtime.NumCPU()			// 返回本地机器的逻辑CPU个数
				fmt.Println("cpuNum=",cpuNum)
				runtime.GOMAXPROCS(cpuNum-1)		// 设置可同时执行的最大CPU数 1.8版本以后不需要设置
			}
chanel*
	(1) channel(管道)-基本介绍：
		1)	channel本质上就是一个数据结构-队列
		2)	数据是先进先出 【FIFO】
		3)	线程安全,多goroutine访问时 不需要加锁,就是说channel本身是线程安全的[多个协程操作同一个管道时,不会发生资源竞争问题]
		4)	channel时有类型的，一个string的channel只能存放string类型数据
	(2) 定义/声明 channel
		1）var 变量名 chan 数据类型 
			var intChan chan int (intChan 存放int数据的管道变量名)
		2) 说明： 
			a. channel是引用类型
			b. channel必须初始化才能写入数据,即make后才能使用
			c. 管道是有类型的,intChan 只能写入int
	(3) 管道的初始化 读 写 
			func main() {
				var intChan chan int 			//申明
				intChan = make(chan int, 3)		//初始化 容量是3 写入不能超过3
				fmt.Printf("值=%v 地址=%p\n",intChan,&intChan) //管道是引用类型 
				intChan <- 10				//向管道写入数据
				num := 10					
				intChan <- num
				intChan <- 11
				fmt.Printf("管道长度=%v 容量=%v\n",len(intChan),cap(intChan))
				//从管道取取数据
				num2 := <-intChan 	// 给管道的第一个数据取出赋值给变量num2
				intChan <- 14		// 先进先出 只能取出一个后再才能像写入一个
				fmt.Printf("管道长度=%v 容量=%v num2=%v\n",len(intChan),cap(intChan),num2)
			}
		 注意事项和细节：	a. channel只能存放指定类型的数据类型	
		 				b. channel数据放满后就不能再放入了,如果再从channel取出数据后,可以继续放入
		 				c. 再没有使用协程的情况下，如果channel数据取完了，再取，就会报deadlock
	(4) 存放任意数据类型的管道
			func main() {
				//定义一个存放任意数据类型的管道 3个数据
				//var allChan chan interface{}
				allChan := make(chan interface{},3)
				allChan <- 10
				allChan <- "tom"
				cat := Cat{"小白猫",4}
				allChan <- cat 
				//我们希望获得管道中的第三个元素 则先将前2个推出
				<-allChan
				<-allChan
				newCat := <-allChan
				fmt.Printf("newCat=%T , newCat=%v\n",newCat,newCat)
				a := newCat.(Cat)							//类型断言
				// fmt.Printf("newCat.Name=%v",newCat.Name)	//错误写法 需要类型断言
				fmt.Printf("newCat.Name=%v",a.Name)			
			}
	(5) channel的遍历和关闭：
			关闭：使用内置函数close可以关闭channel，当channel关闭后，就不能再向channel写数据了，但仍然可以从该channel读数据
			遍历：	a. 使用for-range 方式进行遍历 
					b. 在遍历时，如果channel没有关闭 则会出现deadlock错误
					c. 在遍历时，如果channel已经关闭 则会正常遍历数据，遍历完后就会推出遍历
				func main() {
					intChan := make(chan int,100)
					for i := 0; i < 100; i++ {	//遍历管道
						intChan <- i*2	//放入100个数据到管道
					}
					close(intChan)	//遍历管道前关闭管道 避免出现deadlock
					for v := range intChan{	//取的时候 变量管道只能for-range 
						fmt.Println("v=",v)
					}
				}
	(6) 控制协程在主进程结束之后全部读写完成、
			func writeData(intChan chan int) {	//write Data
				for i := 0; i < 50; i++ {
					intChan <- i
					fmt.Println("write v=",i)
				}
				close(intChan)
			}
			func readData(intChan chan int ,exitChan chan bool) {	//read data
				for { v,ok := <- intChan
					if !ok { break } 
					fmt.Printf("readData 读到数据=%v\n",v)
				}
				exitChan <- true //读完管道数据
				close(exitChan)
			}
			func main() {
				intChan := make(chan int,50)
				exitChan := make(chan bool,1)
				go writeData(intChan)
				go readData(intChan,exitChan)
				for { _,ok := <- exitChan 	//读出是true则说明协程已经读取完毕 
					if !ok { break } 
				}
			}
	(7) channel 可以申明为只读或者只写 性质	
			func main() {
				var chan1 chan<- int 		//申明为只写
				chan1 = make(chan int,3)
				chan1 <- 100
				//num := <- chan1	// error
				var chan2 <-chan int 		//申明为只读
				num2 := <-chan2
				//chan2 <- 20	//error
				fmt.Println(num2)
			}
	(8) 使用*select 解决从管道取数据的阻塞问题
			func main() {
				intChan := make(chan int,10)	//定义一个管道 10个数据 int
				for i := 0; i < 10; i++ { intChan <- i }
				//定义一个管道 5个数据string
				stringChan := make(chan string,5)
				for i := 0; i < 5; i++ { stringChan <- "hello" + fmt.Sprintf("%d",i) }
				//传统的方法在遍历管道时,如果不关闭会阻塞而导致deadlock
				//问题是 实际开发中,我们不好确定在哪关闭该管道 可以使用select 方式解决
				for {
					select {
						//注意：这里,如果intChan一直没有关闭,不会一直阻塞而deadlock 会自动到下一个case匹配
						case v := <- intChan :
							fmt.Printf("从intChan读取数据%d\n",v)
						case v := <- stringChan :
							fmt.Printf("从stringChan读取数据%s\n",v)
						default:
							fmt.Println("都找不到了! 开发者可以加入自己的逻辑")
							return  
					}
				}
			}	
	(9) 使用 *recover 捕获管道错误
			func sayHello() {
				for i := 0; i < 10; i++ { fmt.Println("hello word") }
			}
			func test() {
				defer func(){
					//捕获test抛出的panic
					if err := recover(); err != nil { fmt.Println("test() 发生错误\n",err) }
				}()
				var myMap map[int]string
				myMap[0] = "golang"		//没有make 会panic
			}
			func main() {
				go sayHello()
				go test()
				for i := 0; i < 10; i++ {
					fmt.Println("main() ok=",i)
					time.Sleep(time.Second)
				}
			}	
反射*
	(1)基本介绍：
		1)	反射可以运行时动态获取变量的各种信息 比如变量的类型(type) 类别(kind)
		2)	如果是结构体变量，还可以获取到结构体本身的信息(包括结构体的字段和方法)
		3)	通过反射，可以改变变量的额值，可以调用相关的方法
		4)	使用放射，需要import("reflect")
	(2)反射重要的函数和概念：
		1) 	变量，interface{} 和 reflect.Value 是可以相互转换的.这在实际开发中是经常用到的
	(3)基本数据类型，interface{}，reflect.Value 进行反射的基本操作
			func reflectTest01 (b interface{}){
				//获取reflect.Type 
				rTyp := reflect.TypeOf(b)
				//获取 reflect.Value
				rVal := reflect.ValueOf(b)
				//将 reflect.Value 转 int类型 rVal.Int()
				nu := 2 + rVal.Int()
				//将 reflect.Value 转成 interface{}
				iv := rVal.Interface()
				//将 interface{} 通过断言转换成需要的类型
				num := iv.(int)
			}
	(4) 结构体类型，interface{}，reflect.Value， 进行反射的基本类型操作
			func reflectTest01 (b interface{}){
				rVal := reflect.ValueOf(b)					//获取 reflect.Value
				iV := rVal.Interface()
				fmt.Printf("iv=%v type=%T \n",iV,iV)		//iv={tom 12} type=main.Student
				stu,ok := iV.(Student) 						//需要类型断言
				if ok {
					fmt.Printf("stu.Name=%v\n",stu.Name)	//stu.Name=tom
				}
			}
			type Student struct { Name string; Age int; }
			func main() {
				stu := Student{
					Name : "tom",
					Age : 12,
				}
				reflectTest01(stu)
			}
	(5) reflect.Value.Kind，获取变量的类别，返回的是一个常量
	(6) 通过反射修改变量的值 
		 func reflect1(b interface{}) {
		 	rVal := reflect.ValueOf(b)
		 	fmt.Printf("rVal kind=%v\n",rVal.Kind())
		 	//Elem返回v持有的接口保管的值得Value封装,或有v持有的指针的值的value封装
		 	rVal.Elem().SetInt(20) 	//反射修改变量的值
		 }
		 func main() {
		 	var num int =10
		 	reflect1(&num)
		 	fmt.Println("num=",num)	// 输出20
		 }
TCP编程：
	网络编程有两种：
		1 TCP socket编程，是网络编程的主流。之所有叫Tcp socket编程，是因为底层是基于Tcp/ip 协议的.比如QQ
		2 b/s 结构的http编程，我们使用浏览器取访问服务器时，使用的就http协议，而http协议底层也是tcp socket 实现
	端口：
		a. 0是保留端口
		b. 1-1024是固定端口，又叫做有名端口，即被某些程序固定使用，一般程序员不能使用
			22：SSH 远程登录协议	23: telent 使用 21:ftp使用 25:smtp服务使用 80:iis使用 7:echo服务
		c. 1025-65535 是动态端口，这些端口 程序员可以使用
		d. 可以使用netstat -an 查看本机有哪些端口在监听
Ridis*: (官方文档：http://redisdoc.com/)
	一. Golang 操作Redis:
		1. 使用第三方开源库 github.com/garyburd/redigo/redis
		2. 在使用Redis前 先安装第三方Redis库 在GOPATH路径下执行安装指令：
				D:\goproject>go get github.com/garyburd/redigo/redis
		注意: 在安装前 确保已经安装并配置了Git.
		(1) 添加key-value 
			import "github.com/garyburd/redigo/redis"
			func main() {
				c,err := redis.Dial("tcp","localhost:6379")		//连接redis
				if err != nil { fmt.Println("conn redis faild,",err);return; }
				defer c.Close()
				_,err = c.Do("Set","key1",888) 					//添加key-value
				if err != nil { fmt.Println(err); return; }
				r,err := redis.Int(c.Do("Get","key1")) 			//获取get key
				if err != nil { fmt.Println("get key1 faild,",err); return; }
				fmt.Println(r)
			}
		(2) 操作Hash
			_,err = c.Do("HSet","user01","name","汤姆")
			r,err := redis.String(c.Do("HGet","user01","name")) 	//如果存放的是int则使用reids.Int()
		(3) 批量Set/Get数据
			_,err = c.Do("MSet","name","小明","address","武汉")
			r,err := redis.Strings(c.Do("MGet","name","address"))
	二. Ridis的五大数据类型：String，Hash，List，Set，zSet(有序集合) 
		(1) redis 的基本使用: 默认有16个数据库 初始默认值0号库 编号是0-15
			a. 添加key-val [set]
			b. 查看当前的redis的所有key [keys *]
			c. 获取key对应的值 [get key]
			d. 切换redis数据库 [select index]
			e. 查看当前数据库key-value的数量 [dbsize]
			f. 清空当前数据库的key-val和清空所有数据库的key-val [flushdb  flushall]
		(2) String：二进制安全，除了普通字符串外，也可以存放图片等数据。redis字符串value最大值是512M
			a. CURD:
				set [存在则修改，不出存则添加] get del
				setex key1 10 helloword				设置key1 值是helloword 10秒
				mset key value [key valye...] 		一次性设置多个key value
				mget key1 key2 						同时获取多个key对应的value
		(3) Hash: 是一个键值对集合 是一个string类型的field和value的映射表
			CURD:	
				hset user1 name jack				添加user1字段 的 name jack
				hget user1 name 					获取user1字段下的name的值
				hgetall user1 						获取user1字段下的所有属性的值
				hmset user2 name mary age 21		设置user2字段下的name 和 age 属性的值
				hmget user2 name age 				获取user2 下的name和age的值
				hlen user2 							统计一个哈希有几个元素
				hexists user2 name 					查看user2哈希中是否有name这个属性
		(4) List： List本质是个链表，List元素是有序的，元素的值可以重复
			a. lindex  a按照索引下标获得元素(从0到右 编号从0开始)
			b. llen key 返回列表key的长度 如果可以不存在，则key被解释为一个空列表 返回0
			c.CURD:
				lpush city beijing shanghai 		向city列表中从左边插入beijing shanghai 
				lrange city 0 -1 					取所有的city列表中的所有值
				rpush herolist aa bb cc 			从链表的左边插入三个值
				lpop herolist 						从链表左边弹出一个 返回弹出的值
				rpop herolist						从链表右边弹出一个 返回弹出的值
				del  herolist						删除链表
	 	(5) Set: Ridis的Set是String类型的无序集合，且元素值不能重复
	 		CURD:
	 			sadd emails xx xxx 					向集合emails 集合添加2个值
	 			smembers emails						遍历该集合中的所值
	 			sismember 							判断是否是成员
	 			srem 								删除指定的值
	三. Ridis连接池 ： 事先初始化一定数量的连接,放入连接池.当Go需要操作Redis时 直接从redis连接池取出连接即可
		(1) 核心代码：
			var pool *redis.Pool
			pool = &redis.Pool{
				MaxIdle:8, 						//最大空闲链接数
				MaxActive:0,					//表示和数据库的最大链接数,0表示没有限制
				IdleTimeout:100,				//最大空闲时间
				Dial:func()(redis.Conn,error){ 	//初始化
					return redis.Dial("tcp","localhost:6379")
				},
			}
			c := pool.Get()  					//从链接池中取出一个链接
			pool.Close()						//关闭连接池 就再不能从连接池中取链接了
		(2) 案例：
			import "github.com/garyburd/redigo/redis"
			var pool *redis.Pool
			func init() {
					pool = &redis.Pool{
					MaxIdle:8, 						
					MaxActive:0,					
					IdleTimeout:100,				
					Dial:func()(redis.Conn,error){ 	
						return redis.Dial("tcp","localhost:6379")
					},
				}
			}
			func main() {
				conn := pool.Get()
				defer conn.Close()
				_,err := conn.Do("Set","name","jack")
				if err != nil{ fmt.Println(err); return; }
				r,err := redis.String(conn.Do("Get","name"))
				if err != nil{ fmt.Println(err); return; }
				fmt.Println("r=",r)
			}
数据结构*：
	a. 数据结构是一门研究算法的学科.要学数据结构就要多多思考如何将生活中遇到的问题,用程序去实现解决
	稀疏数组： 当一个数组中大部分元素是0，或者为同一个值的数组时，可以使用稀疏数组来保存该数组
		稀疏数组的处理方法：
			记录数组一共有几行几列,有多少个不同的值，吧具有不同值得元素的行列及值记录在一个小规模的数组中，从而缩小程序的规模
	队列：队列是一个有序列表，可以用数组和链表实现 遵循先入先出的原则