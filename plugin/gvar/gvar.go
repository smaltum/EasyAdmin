package gvar

import (
	"github.com/spf13/cast"
)

type Var struct {
	value interface{}
	safe  bool
}

func New(value interface{}, safe ...bool) *Var {
	if len(safe) > 0 && !safe[0] {
		return &Var{
			value: value,
			safe:  true,
		}
	}
	return &Var{
		value: value,
	}
}

func (r *Var) String() string {
	return cast.ToString(r.value)
}

func (r *Var) Int() int {
	return cast.ToInt(r.value)
}

func (r *Var) Int64() int64 {
	return cast.ToInt64(r.value)
}

func (r *Var) Bool() bool {
	return cast.ToBool(r.value)
}
