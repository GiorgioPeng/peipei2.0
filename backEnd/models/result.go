package models

import (
	"encoding/json"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResult(data interface{}, c int, m ...string) *Result {
	r := &Result{Data: data, Code: c}

	if e, ok := data.(error); ok {
		if m == nil {
			r.Msg = e.Error()
		}
	} else {
		r.Msg = "SUCCESS"
	}
	if len(m) > 0 {
		r.Msg = m[0]
	}

	return r
}

func (res Result) JsonEncode() ([]byte, error) {
	data, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return data, nil
}
