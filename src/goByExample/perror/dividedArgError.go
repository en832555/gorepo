package perror

import "fmt"

type DividedArgError struct {
	Divisor int
	Dividend int
	Msg string
}
func (e *DividedArgError)Error () string{
	return  fmt.Sprintf("%d不能被%d整除",e.Dividend,e.Divisor,e.Msg)
}