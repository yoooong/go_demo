package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpBaseRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CommonHttpBaseRes struct {
	HttpBaseRes
	Data interface{} `json:"data,omitempty"`
}

// HTTP HEAD相关字段定义
const (
	HTTP_HD_APP_TOKEN        = "X-App-Token"
	HTTP_HD_APP_UID          = "X-App-Uid"
	HTTP_HD_CLIENT_TYPE      = "X-Client-Type"
	HTTP_HD_REQUEST_ID       = "X-App-Request-Id" // 用来进行bug查找和问题跟踪用的
	HTTP_HD_APP_VERSION_CODE = "X-App-Version-Code"
)

type SelfHttpHeader struct {
	Token        string
	Uid          uint64
	ClientType   uint32
	SysInfo      map[string]interface{}
	SysInfoStr   string
	AppMajorVer  uint32
	AppMinorVer  uint32
	AppReviseVer uint32
	VersionCode  uint32
	RequestId    string
	ProtoType    string
}

func (this *SelfHttpHeader) String() string {
	ret := HTTP_HD_APP_UID + "=" + fmt.Sprint(this.Uid) + ","
	ret += HTTP_HD_APP_TOKEN + "=" + this.Token + ","
	ret += HTTP_HD_CLIENT_TYPE + "=" + fmt.Sprint(this.ClientType) + ","
	ret += HTTP_HD_APP_VERSION_CODE + "=" + fmt.Sprint(this.VersionCode) + ","

	return ret
}

func (this *SelfHttpHeader) CheckArgs() bool {
	return this.Uid != 0 && this.ClientType != 0 && this.Token != ""
}

func WriteHttpJsonRes(resWriter http.ResponseWriter, obj interface{}) (err error) {
	var data []byte
	if data, err = json.Marshal(obj); err != nil {
		return
	} else {
		resWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, err = resWriter.Write(data)
		//if err != nil {
		//	log.Fatalf("http write err:%s", err.Error())
		//} else {
		//	log.Fatalf("http write success")
		//}
	}
	return
}

// 在独立的协程安全的执行给定的函数，执行过程中会捕获任何的panic
func SafeExecFunc(fn func(...interface{}), args ...interface{}) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
			}
		}()

		fn(args...)
	}()
}
