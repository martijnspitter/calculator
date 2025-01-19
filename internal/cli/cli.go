package cli

import (
	"github.com/spf13/cobra"
)

var (
	input string
)

// Config will hold our parsed command line arguments
type Config struct {
	Input string
}

type Cli interface {
	Execute() (Config, error)
}

type Cmd struct {
	rootCmd *cobra.Command
}

func NewConfig(input string) Config {
	return Config{
		Input: input,
	}
}

func NewCmd() *Cmd {
	return &Cmd{
		rootCmd: &cobra.Command{
			Use:   "calc",
			Short: "Calculator",
			Long:  `Calculator is a simple calculator that can perform basic arithmetic operations.`,
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) > 0 {
					input = args[0]
				}
			},
		},
	}
}

func (c *Cmd) Execute() (Config, error) {
	if err := c.rootCmd.Execute(); err != nil {
		return Config{}, err
	}

	return NewConfig(input), nil
}
