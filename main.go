package main

import (
	"readitlater/cmd"
	"readitlater/data"
)

func main() {
	data.InitializeDB("./data.db")
	cmd.Execute()
}
