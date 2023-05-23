package main

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	orm.RegisterDataBase("default", "mysql", "beego_queries:tmp_pwd@tcp(127.0.0.1:3306)/beego_queries")
	orm.RegisterModel(new(User))

	o := orm.NewOrm()

	// Insert
	user := User{Name: "MountainMaverick", Age: 34}
	_, err := o.Insert(&user)
	fmt.Println(user, err)

	// Update the user we just created
	user.Name = "RainbowRanger"
	user.Age = 35
	num, err := o.Update(&user)
	fmt.Println(user)
	fmt.Println(num, err)

	// Delete the user
	num, err = o.Delete(&user)
	fmt.Println(num, err)
}
