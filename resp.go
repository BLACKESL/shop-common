package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 通过json tag进行结构体赋值
func SwapToStruct(req, target interface{}) (err error) {
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, target)
	return nil
	fmt.Println(1)
}

type H struct {
	Code                   string
	Message                string
	Traceid                string
	Data                   interface{}
	Rows                   interface{}
	Total                  interface{}
	SkyWailingDynamicField string
}

func Resp(w http.ResponseWriter, code string, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:    code,
		Message: message,
		Data:    data,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(ret)
}
func RespOK(w http.ResponseWriter, message string, data interface{}) {
	Resp(w, "SUCCESS", message, data)
}
func RespCreate(w http.ResponseWriter, message string, data interface{}) {
	Resp(w, "TOKEN_FAIL", message, data)
}
func RespList(w http.ResponseWriter, code string, message string, data interface{}, rows interface{}, total interface{}, skyWailingDynamicField string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:                   code,
		Message:                message,
		Data:                   data,
		Rows:                   rows,
		Total:                  total,
		SkyWailingDynamicField: skyWailingDynamicField}
	ret, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(ret)
}
func RespListOK(w http.ResponseWriter, message string, data interface{}, rows interface{}, total interface{}, skyWailingDynamicField string) {
	RespList(w, "SUCCESS", message, data, rows, total, skyWailingDynamicField)
}
func RespListFail(w http.ResponseWriter, code int, message string, data interface{}, rows interface{}, total interface{}, skyWailingDynamicField string) {
	RespList(w, "TOKEN_FAIL", message, data, rows, total, skyWailingDynamicField)
}
