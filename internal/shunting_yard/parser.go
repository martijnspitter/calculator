package shuntingyard

type Parser struct {
	input string
}

func NewParser(input string) Parser {
	return Parser{
		input: input,
	}
}

func (p *Parser) Parse() (string, error) {
	return "", nil
}
