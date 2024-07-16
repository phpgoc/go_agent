package utils

import (
	pb "go-agent/agent_proto"
	"reflect"
)

// SetResponseErrorAndLogMessageGeneric sets the message and code fields of a response struct
// 这里如果出错,直接退出程序,一定是代码有问题 传入的必须是指针并且必须包含Message和Code字段
// 如果有setter方法,可以优雅的使用接口,但是proto生成的代码没有setter方法
// 这个error根本就没有用,因为proto实现要求返回,真有错都panic了
func SetResponseErrorAndLogMessageGeneric[T any](responsePtr T, message string, code pb.ResponseCode) (T, error) {
	val := reflect.ValueOf(&responsePtr).Elem()
	if val.Kind() != reflect.Ptr {
		//这里基本不会进,因为要求返回的是指针,所以传入的一定是指针
		LogErrorWithCallerLevel("responsePtr is not a pointer", 3)
		panic("responsePtr is not a pointer")
	}
	val = val.Elem()

	messageField := val.FieldByName("Message")
	if messageField.IsValid() && messageField.CanSet() && messageField.Kind() == reflect.String {
		LogErrorWithCallerLevel(message, 3)
		messageField.SetString(message)
	} else {
		LogErrorWithCallerLevel("Message field not found or cannot be set", 3)
		panic("message field not found or cannot be set")
	}

	codeField := val.FieldByName("Code")
	if codeField.IsValid() && codeField.CanSet() && codeField.Kind() == reflect.Int32 {
		codeField.SetInt(int64(code))
	} else {
		LogErrorWithCallerLevel("Code field not found or cannot be set", 3)
		panic("code field not found or cannot be set")
	}

	return responsePtr, nil
}
