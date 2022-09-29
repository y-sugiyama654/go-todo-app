package main

import (
	"fmt"
	"main/app/models"
)

func main() {
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		// ログ出力(動作確認用)
		log.Println("test")

		fmt.Println(models.Db)
	*/

	/*
		u := &models.User{}
		u.Name = "Yuta"
		u.Email = "yuta@gmail.com"
		u.PassWord = "Test2test"
		fmt.Println(u)
		u.CreateUser()
	*/

	u, _ := models.GetUser(1)
	fmt.Println(u)

	u.Name = "Test2"
	u.Email = "test2@example.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)

	u.DeleteUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)
}
