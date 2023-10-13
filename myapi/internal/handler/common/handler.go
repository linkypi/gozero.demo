package common

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
	xhttp "github.com/zeromicro/x/http"
	"myapi/internal/logic/auth"
	"myapi/internal/svc"
	"net/http"
	"reflect"
	"strings"
)

type stubMapping map[string]interface{}

var StubStorage = stubMapping{}

//	func Call(funcName string, params ...interface{}) (result interface{}, err error) {
//		f := reflect.ValueOf(StubStorage[funcName])
//		if len(params) != f.Type().NumIn() {
//			//err = errors.New("The number of params is out of index.")
//			logx.Error("The number of params is out of index.")
//			return
//		}
//		in := make([]reflect.Value, len(params))
//		for k, param := range params {
//			in[k] = reflect.ValueOf(param)
//		}
//		var res []reflect.Value
//		res = f.Call(in)
//		result = res[0].Interface()
//		return
//	}
var am auth.LoginLogic
var pm auth.LoginLogic
var registry = map[string]reflect.Type{
	"auth":  reflect.TypeOf(am),
	"party": reflect.TypeOf(pm),
}

func Call(module string, funcName string,
	w http.ResponseWriter, r *http.Request) (result interface{}, err error) {

	f, ok := registry[module].MethodByName("Login")
	if !ok {
		logx.Error("function name not found in module.")
		return nil, errors.New(500, "function name not found in module")
	}
	for i := 0; i < f.Type.NumIn(); i++ {
		fmt.Println("参数:", f.Type.In(i))
	}

	var req = reflect.New(f.Type.In(1))
	if e := httpx.Parse(r, &req); e != nil {
		xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New(505, "参数异常"))
		//httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	//f := reflect.ValueOf(StubStorage[funcName])
	//if len(params) != f.Type().NumIn() {
	//	//err = errors.New("The number of params is out of index.")
	//	logx.Error("The number of params is out of index.")
	//	return
	//}
	//pa := req.Interface().(*types.LoginReq)
	//fmt.Println("params: ", pa)
	inParams := []reflect.Value{req}
	//for k, param := range params {
	//	in[k] = reflect.ValueOf(param)
	//}
	var res []reflect.Value
	fun := reflect.New(f.Type)

	res = fun.Call(inParams)
	result = res[0].Interface()
	return result, nil
}

func Handle(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		arr := strings.Split(r.RequestURI, "/")
		module := arr[2]
		method := arr[3]

		resp, err := Call(module, method, w, r)

		//l := auth.NewLogic(r.Context(), svcCtx)
		//
		//resp, err := l.Login(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			//httpx.ErrorCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
			//httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
