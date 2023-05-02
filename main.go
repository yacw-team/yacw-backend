package main

import (
	"github.com/yacw-team/yacw/routes"
	"github.com/yacw-team/yacw/utils"
)

func main() {
	utils.InitDB()

	r := routes.SetupRouter()
	err := r.Run()
	if err != nil {
		print("Something went wrong!")
		panic(err)
	}
}
