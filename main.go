package main

import (
	"fmt"
	"main/app/models"
)

func main() {
	// fmt.Println(models.Db)
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		// ログ出力(動作確認用)
		log.Println("test")

		fmt.Println(models.Db)
	*/

	// u := &models.User{}
	// u.Name = "Yuta"
	// u.Email = "yuta@gmail.com"
	// u.PassWord = "Test2test"
	// // userの作成
	// u.CreateUser()

	// user, _ := models.GetUser(2)
	// // totoの作成
	// user.CreateTodo("First Todo")
	// todo, _ := models.GetTodo(1)
	// fmt.Println(todo)

	// fmt.Println(u)

	// u.Name = "Test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// todo, _ := models.GetTodo(1)
	// fmt.Println(todo)

	// user, _ := models.GetUser(2)
	// todos, _ := user.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	t, _ := models.GetTodo(2)
	t.Content = "Second Todo"
	t.UpdateTodo()
	t, _ = models.GetTodo(2)
	fmt.Println(t)
	// fmt.Println(t)
	// t.DeleteTodo()
}
