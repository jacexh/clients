package cctp

import "fmt"

type BadRequestResponse struct {
	Code         int
	ErrorMessage string
}

type UndefinedError struct {
	Code         int
	ErrorMessage string
}

func (err *BadRequestResponse) Error() string {
	return fmt.Sprintf("[%d] %s", err.Code, err.ErrorMessage)
}

func (err *UndefinedError) Error() string {
	return fmt.Sprintf("[%d] %s", err.Code, err.ErrorMessage)
}
