package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http/httptest"
	"techtrainingcamp-security-10/internal/resource"
	"testing"
)

func TestRandomString(t *testing.T) {
	var testLetters = []byte("0123456789")
	if RandomString(6, testLetters) == RandomString(6, testLetters) {
		t.Error("验证码应当随机")
	}
	testLetters = []byte("0")
	if RandomString(10, testLetters) != "0000000000" {
		t.Error("单个字符是验证码应当不随机")
	}
}

type applyCodeRespond struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
}

// 正常流程获取验证码
func TestApplyCode1(t *testing.T) {
	server, _ := resource.NewServer()
	s := server.Service
	router := gin.Default()
	method := "GET"
	uri := "/api/apply_code"
	router.GET(uri, ApplyCode(s))

	contextTest := make(map[string]uint)
	contextTest["PhoneNumber"] = 13812345678
	contextByte, _ := json.Marshal(contextTest)
	req := httptest.NewRequest(method, uri, bytes.NewReader(contextByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result, _ := ioutil.ReadAll(w.Result().Body)
	fmt.Println(string(result))
	resultJSON := &applyCodeRespond{}
	_ = json.Unmarshal(result, resultJSON)
	if resultJSON.Code != 0 || resultJSON.Message != "请求成功" {
		t.Error("请求失败了")
	}
}

// 不合法手机号
func TestApplyCode2(t *testing.T) {
	server, _ := resource.NewServer()
	s := server.Service
	router := gin.Default()
	method := "GET"
	uri := "/api/apply_code"
	router.GET(uri, ApplyCode(s))

	contextTest := make(map[string]uint)
	contextTest["PhoneNumber"] = 12812345678
	contextByte, _ := json.Marshal(contextTest)
	req := httptest.NewRequest(method, uri, bytes.NewReader(contextByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result, _ := ioutil.ReadAll(w.Result().Body)
	fmt.Println(string(result))
	resultJSON := &applyCodeRespond{}
	_ = json.Unmarshal(result, resultJSON)
	if resultJSON.Code != 1 || resultJSON.Message != "请换个手机号试试" {
		t.Error("不合法手机号请求成功了")
	}
}
