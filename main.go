package main

import (
	"github.com/yacw-team/yacw/routes"
	"github.com/yacw-team/yacw/utils"
	"os"
)

func main() {

	_, err := utils.GetSalt()
	if err != nil {
		os.Exit(1)
	}

	utils.InitDB()
	if !utils.DatabaseCheck() {
		os.Exit(1)
	}
	r := routes.SetupRouter()
	err = r.Run()
	if err != nil {
		print("Something went wrong!")
		panic(err)
	}
}
