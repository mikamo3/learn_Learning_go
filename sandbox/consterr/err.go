package consterr

import "reflect"

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func (se StatusErr) Is(target error) bool {
	if me2, ok := target.(StatusErr); ok {
		return reflect.DeepEqual(se, me2)
	}
	return false
}
func LoginAndGGetData(uid, pwd, file string) error {
	return StatusErr{
		Status:  InvalidLogin,
		Message: "invalid",
	}
}
