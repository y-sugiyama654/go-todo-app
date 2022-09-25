package main

import (
	"fmt"
	"go-todo-app/app/models"
	"go-todo-app/config"
	"log"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	// ログ出力(動作確認用)
	log.Println("test")

	fmt.Println(models.Db)
}
