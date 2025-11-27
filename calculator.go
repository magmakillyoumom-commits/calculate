package calculator_

type Calculator interface {
	Calculate(a, b int) int
}

type Plus struct{}

func (p Plus) Calculate(a, b int) int { return a + b }

type Minus struct{}

func (m Minus) Calculate(a, b int) int { return a - b }

type Multiply struct{}

func (m Multiply) Calculate(a, b int) int { return a * b }

type Divide struct{}

func (d Divide) Calculate(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func New(operator string) Calculator { // эта хуйня как раз дает выполнить тот или иной метод если находит оператора в свич
	// соотвественно если выолнятся какой то из методов то он проверяется интерфейсом и возращает резутат
	switch operator {
	case "+":
		return Plus{}
	case "-":
		return Minus{}
	case "*":
		return Multiply{}
	case "/":
		return Divide{}
	default:
		return nil
	}
}
