package calc

import "fmt"

const (
	operatorplus  = "+"
	operatorminus = "-"
	operatormt    = "*"
	operatordiv   = "/"
)

type calculator struct{}

func NewCalculator() calculator {
	return calculator{}
}

func (c *calculator) Calculate(oprnd1, oprnd2 float64, operator string) float64 {

	switch operator {
	case operatorplus:
		return c.plus(oprnd1, oprnd2)
	case operatorminus:
		return c.minus(oprnd1, oprnd2)
	case operatormt:
		return c.mt(oprnd1, oprnd2)
	case operatordiv:
		return c.div(oprnd1, oprnd2)
	default:
		fmt.Printf("несуществующий оператор: %s\n", operator)
		return 0
	}
}
func (c *calculator) plus(oprnd1, oprnd2 float64) float64 {
	return oprnd1 + oprnd2
}
func (c *calculator) minus(oprnd1, oprnd2 float64) float64 {
	return oprnd1 - oprnd2
}
func (c *calculator) mt(oprnd1, oprnd2 float64) float64 {
	return oprnd1 * oprnd2
}
func (c *calculator) div(oprnd1, oprnd2 float64) float64 {
	if oprnd2 == 0 {
		fmt.Println("ошибка, деление на 0")
		return 0
	}

	return oprnd1 / oprnd2
}
