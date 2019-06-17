package response

import (
	"reflect"

	"myiris/library/logger"
)

// 应答统一结构
type Response struct {
	Code   int32       `json:"code"`
	ErrMsg string      `json:"errMsg"`
	Body   interface{} `json:"body"`
}

func NewResponse(res ...interface{}) *Response {
	response := &Response{
		Code:   400,
		ErrMsg: "visit to failed",
	}

	if len(res) > 3 {
		logger.Error("Unsupported parameter length, it's too long")
		return &Response{}
	}

	for _, v := range res {
		if v != nil {
			t := reflect.TypeOf(v)
			switch t.Kind() {
			case reflect.Int, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint32, reflect.Uint64:
				response.Code = v.(int32)
			case reflect.String:
				response.ErrMsg = v.(string)
			case reflect.Ptr:
				t = t.Elem()
				if t.Kind() == reflect.Struct {
					response.Body = v
				} else {
					logger.Error("Invalid data type", t.Kind())
				}
			case reflect.Struct:
				logger.Error("Invalid body type, Body should be a struct pointer")
				return &Response{}

			default:
				logger.Error("Unsupported data types")
				return &Response{}
			}
		} else {
			response.ErrMsg = ""
		}
	}
	return response
}
