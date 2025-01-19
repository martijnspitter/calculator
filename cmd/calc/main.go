package main

import (
	"calculator/internal/cli"
	"calculator/internal/service"
	"fmt"
)

func main() {
	cli := cli.NewCmd()

	service := service.NewSortService(cli)
	err := service.Execute()

	if err != nil {
		fmt.Println(err)
	}
}
