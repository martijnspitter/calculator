package service

import (
	"calculator/internal/cli"
	"fmt"
)

type Service interface {
	Execute() error
}

type SortService struct {
	cli cli.Cli
}

func NewSortService(cli cli.Cli) Service {
	return &SortService{
		cli: cli,
	}
}

func (s *SortService) Execute() error {
	config, err := s.cli.Execute()
	fmt.Println(config)

	if err != nil {
		return err
	}

	return nil
}
