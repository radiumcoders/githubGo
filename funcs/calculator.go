package funcs

import "errors"

// its recommended to use a struct for error handling
type ReturnValue struct {
	Result float32
	Error  error
}

func Multiplication(a, b *float32) ReturnValue {
	if a == nil || b == nil {
		return ReturnValue{0, errors.New("Multiplication by nil — who wants to do that?!")}
	}
	if *a == 0 || *b == 0 {
		return ReturnValue{0, errors.New("Multiplication by 0 — who wants to do that?!")}
	}
	return ReturnValue{*a * *b, nil}
}

func Division(a, b *float32) ReturnValue {
	if *b == 0 {
		return ReturnValue{0, errors.New("Division by 0 — who wants to do that?!")}
	}
	return ReturnValue{*a / *b, nil}
}

func Addition(a, b *float32) ReturnValue {
	if a == nil || b == nil {
		return ReturnValue{0, errors.New("Addition by nil — who wants to do that?!")}
	}
	if *a == 0 || *b == 0 {
		return ReturnValue{0, errors.New("Addition by 0 — who wants to do that?!")}
	}
	return ReturnValue{*a + *b, nil}
}

func Subtraction(a, b *float32) ReturnValue {
	if a == nil || b == nil {
		return ReturnValue{0, errors.New("Subtraction by nil — who wants to do that?!")}
	}
	if *a == 0 || *b == 0 {
		return ReturnValue{0, errors.New("Subtraction by 0 — who wants to do that?!")}
	}
	return ReturnValue{*a - *b, nil}
}