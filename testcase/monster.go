package monster
//单元测试
import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Monster struct {
	Name string 
	Age int
	Skill string 
}

// 给monster 绑定一个store方法 将monster对象 序列化保存到文件中
 func (this *Monster) Store() bool {
	//1序列化
	data,err := json.Marshal(this)
	if err != nil {
		fmt.Print("Marshal err= ",err)
		return false
	}
	//2保存到文件
	filePath := "d:/monster.txt"
	err = ioutil.WriteFile(filePath,data,0666)
	if err != nil {
		fmt.Print("Write err= ",err)
		return false
	}
	return true
 }

 //给monster 绑定一个ReStore方法 将序列化的Monster从文件中读取 并反序列化为Monster对象 检测反序列化
 func (this *Monster) ReStore() bool {
	//1 从文件中读取序列化的字符串
	filePath := "d:/monster.txt"
	data,err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Print("ReadFile err= ",err)
		return false
	}
	//2 读取到的data []byte 进行反序列列化 存在文件
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Print("Unmarshal err= ",err)
		return false
	}
	return true
 }