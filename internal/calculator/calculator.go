package calculator

type Calculator struct {
	input []string
}

func NewCalculator(input []string) Calculator {
	return Calculator{
		input: input,
	}
}

func (c *Calculator) Calculate() float64 {
	return 0
}
