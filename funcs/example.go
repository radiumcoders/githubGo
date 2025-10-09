package funcs

import "errors"

type ReturnValue struct {
	Result int
	Error  error
}

func Multiplication(a, b int) ReturnValue {
	if a == 0 || b == 0 {
		return ReturnValue{0, errors.New("Multiplication by 0 — who wants to do that?!")}
	}
	return ReturnValue{a * b, nil}
}

func Division(a, b int) ReturnValue {
	if b == 0 {
		return ReturnValue{0, errors.New("Division by 0 — who wants to do that?!")}
	}
	return ReturnValue{a / b, nil}
}
