package service

import (
	"calculator/internal/calculator"
	"calculator/internal/cli"
	"calculator/internal/shunting_yard"
)

type Service interface {
	Execute() (float64, error)
}

type SortService struct {
	cli cli.Cli
}

func NewSortService(cli cli.Cli) Service {
	return &SortService{
		cli: cli,
	}
}

func (s *SortService) Execute() (float64, error) {
	config, err := s.cli.Execute()
	if err != nil {
		return 0, err
	}

	parser := shuntingyard.NewParser(config.Input)
	parsed := parser.Parse()
	calc := calculator.NewCalculator(parsed)
	result, err := calc.Calculate()

	if err != nil {
		return 0, err
	}

	return result, nil
}
