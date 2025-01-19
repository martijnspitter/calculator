package main

import (
	"calculator/internal/cli"
	"calculator/internal/service"
	"fmt"
	"os"
)

func main() {
	cli := cli.NewCmd()

	service := service.NewSortService(cli)
	result, err := service.Execute()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, result)
	os.Exit(0)
}
