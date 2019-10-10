package main
//json 序列化和反序列化
import (
	"fmt"
	"encoding/json"
)

type Monster struct {
	Name string `json:"name"`
	Age int
	Birthday string 
	Sal float64
	skill string 
}

//结构体序列化
func testStruct(){
	monster := Monster{
		Name : "牛魔王",
		Age : 500,
		Birthday : "1880-11-22",
		Sal : 8000.0,
		skill : "风沙走石",
	}
	// monster 序列化
	data,err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化失败",err)
	}
	fmt.Printf("序列化后的结构%v\n",string(data))
}
//Map序列化
func testMap(){
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"]="红孩儿"
	a["age"]=15
	a["address"]="火云洞"

	data,err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化错误",err)
	}
	fmt.Printf("序列化后的结构%v\n",string(data))
}
//切片序列化
func testSlice(){
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 10
	m1["address"] = "上海"
	slice = append(slice,m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 19
	m2["address"] = [2]string{"长沙","武汉"}
	slice = append(slice,m2)

	data,err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化错误",err)
	}
	fmt.Printf("序列化后的结构%v\n",string(data))
}

func testFloat64(){
	var num1 float64 = 123.54
	data,err := json.Marshal(num1)
	if err != nil{
		fmt.Println("序列化错误",err)
	}
	fmt.Printf("序列化后的结构%v\n",string(data))
}

//反序列化Struct
func unmarshalStruct() {
	str := "{\"name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"1880-11-22\",\"Sal\":8000}"
	var monster Monster
	err := json.Unmarshal([]byte(str),&monster)
	if err != nil {
		fmt.Println("反序列化错误",err)
	}
	fmt.Printf("反序列化后 monster=%v monster.Name=%v\n",monster,monster.Name)
}
//反序列化Map
func unmarshalMap() {
	str := "{\"address\":\"火云洞\",\"age\":15,\"name\":\"红孩儿\"}"
	var map1 map[string]interface{}
	//反序列化map的时候不需要make,因为make操作封装到Unmarshal函数
	err := json.Unmarshal([]byte(str),&map1)
	if err != nil {
		fmt.Println("反序列化错误",err)
	}
	fmt.Printf("反序列化后 monster=%v \n",map1)

}
//反序列化Slice
func unmarshalSlice() {
	str := "[{\"address\":\"上海\",\"age\":10,\"name\":\"jack\"},"+
	"{\"address\":[\"长沙\",\"武汉\"],\"age\":19,\"name\":\"tom\"}]"
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str),&slice)
	if err != nil {
		fmt.Println("反序列化错误",err)
	}
	fmt.Printf("反序列化后 monster=%v \n",slice)
}
func main(){
	// testStruct()
	// testMap()
	// testSlice()
	// testFloat64()
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
