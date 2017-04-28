package main

import (
	"MyUtils/dao"
	"fmt"
)

type User struct {
	Id int
	Name string
}
var mysqler dao.MySqler
func init(){
	mysqler.CharSet="utf8"
	mysqler.DbName="test_user"
	mysqler.Host="localhost"
	mysqler.Port= "3306"
	mysqler.PassWord="root"
	mysqler.User="root"
}
func main() {
	//插入数据
	userInsert :=&User{1,"大帅气"}
	mysqler.Insert(userInsert)

	//查询数据
	result :=mysqler.Query("user",`name ="dashuaiqi"`)

	//解析数据结果
	var id int
	var name string
	for result.Next(){
		result.Scan(&id,&name)
		fmt.Println(id)
		fmt.Println(name)
	}
	defer result.Close()

	//更新数据
	mysqler.Update("user",`name = "lwc"`,`name="dashuaiqi"`)

	//删除数据
	userDelete :=&User{1,"大帅气"}
	mysqler.Delete(userDelete)

}

