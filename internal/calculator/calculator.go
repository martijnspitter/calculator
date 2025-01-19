package calculator

type Calculator struct {
	input string
}

func NewCalculator(input string) Calculator {
	return Calculator{
		input: input,
	}
}
