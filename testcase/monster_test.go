package monster
//单元测试
/*	指令：
	go test -v							测试所有
	go test -v -test.run TestReStore	只测试TestReStore 
*/
import (
	"testing"
)

//测试 Store 
func TestStore(t *testing.T) {
	//先创建一个monster对象实例
	monster := &Monster{
		Name : "红孩儿",
		Age : 11,
		Skill : "三位真火",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store() 错误,希望为=%v 实例=%v",true,res)
	}
	t.Logf("monster.Store() 测试成功")
}

//测试 ReStore 
func TestReStore(t *testing.T) {
	//先创建一个monster对象实例 
	var monster Monster
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore() 错误,希望为=%v 实例=%v",true,res)
	}
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.ReStore() 错误,希望为=%v 实例=%v","红孩儿",monster.Name)
	}
	t.Logf("monster.ReStore() 测试成功")
}