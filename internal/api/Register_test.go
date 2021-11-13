package api

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http/httptest"
	"techtrainingcamp-security-10/internal/resource"
	"techtrainingcamp-security-10/internal/service"
	"testing"
)

type registerRespond struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
}

// Case1: 用户名已注册
func TestRegister1(t *testing.T) {
	server, _ := resource.NewServer()
	s := server.Service
	router := gin.Default()
	method := "POST"
	uri := "/api/register"
	router.POST(uri, Register(s))

	PasswordAddSalt := service.NewPassword("123456")
	testUser := service.UserTable{
		UserName:    "test",
		Password:    PasswordAddSalt,
		PhoneNumber: "13887654321",
	}

	s.InsertVerifyCode("13812345678", "123456")
	defer s.DeleteVerifyCode("13812345678")
	err := s.InsertUser(testUser)
	if err != nil {
		t.Error("添加test用户失败了")
	}
	defer s.DeleteUserByPhoneNumber("13887654321")

	contextTest := make(map[string]interface{})
	contextTest["UserName"] = "test"
	contextTest["Password"] = "123456"
	contextTest["PhoneNumber"] = 13812345678
	contextTest["VerifyCode"] = "123456"

	contextByte, _ := json.Marshal(contextTest)
	req := httptest.NewRequest(method, uri, bytes.NewReader(contextByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result, _ := ioutil.ReadAll(w.Result().Body)
	resultJSON := &registerRespond{}
	_ = json.Unmarshal(result, resultJSON)
	if resultJSON.Code != 1 || resultJSON.Message != "相同的用户名已经被注册过了，请更换用户名试试" {
		t.Error("用户名重复但注册成功了")
	}
}

// Case2: 手机号已注册
func TestRegister2(t *testing.T) {
	server, _ := resource.NewServer()
	s := server.Service
	router := gin.Default()
	method := "POST"
	uri := "/api/register"
	router.POST(uri, Register(s))

	PasswordAddSalt := service.NewPassword("123456")
	testUser := service.UserTable{
		UserName:    "test",
		Password:    PasswordAddSalt,
		PhoneNumber: "13812345678",
	}
	s.InsertVerifyCode("13812345678", "123456")
	defer s.DeleteVerifyCode("13812345678")
	err := s.InsertUser(testUser)
	if err != nil {
		t.Error("添加test用户失败了")
	}
	defer s.DeleteUserByPhoneNumber("13812345678")

	contextTest := make(map[string]interface{})
	contextTest["UserName"] = "test2"
	contextTest["Password"] = "123456"
	contextTest["PhoneNumber"] = 13812345678
	contextTest["VerifyCode"] = "123456"

	contextByte, _ := json.Marshal(contextTest)
	req := httptest.NewRequest(method, uri, bytes.NewReader(contextByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result, _ := ioutil.ReadAll(w.Result().Body)
	resultJSON := &registerRespond{}
	_ = json.Unmarshal(result, resultJSON)
	if resultJSON.Code != 1 || resultJSON.Message != "相同的手机号已经被注册过了，请更换用户名试试" {
		t.Error("手机号重复但注册成功了")
	}
}

// Case3: 手机号不合法
func TestRegister3(t *testing.T) {
	server, _ := resource.NewServer()
	s := server.Service
	router := gin.Default()
	method := "POST"
	uri := "/api/register"
	router.POST(uri, Register(s))

	contextTest := make(map[string]interface{})
	contextTest["UserName"] = "test"
	contextTest["Password"] = "123456"
	contextTest["PhoneNumber"] = 12812345678
	contextTest["VerifyCode"] = "123456"

	contextByte, _ := json.Marshal(contextTest)
	req := httptest.NewRequest(method, uri, bytes.NewReader(contextByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result, _ := ioutil.ReadAll(w.Result().Body)
	resultJSON := &registerRespond{}
	_ = json.Unmarshal(result, resultJSON)
	if resultJSON.Code != 1 || resultJSON.Message != "请换个手机号试试" {
		t.Error("手机号不合法但请求成功了")
	}
}

// Case4: 验证码不正确
func TestRegister4(t *testing.T) {
	server, _ := resource.NewServer()
	s := server.Service
	router := gin.Default()
	method := "POST"
	uri := "/api/register"
	router.POST(uri, Register(s))

	s.InsertVerifyCode("13812345678", "654321")
	defer s.DeleteVerifyCode("13812345678")

	contextTest := make(map[string]interface{})
	contextTest["UserName"] = "test2"
	contextTest["Password"] = "123456"
	contextTest["PhoneNumber"] = 13812345678
	contextTest["VerifyCode"] = "123456"

	contextByte, _ := json.Marshal(contextTest)
	req := httptest.NewRequest(method, uri, bytes.NewReader(contextByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result, _ := ioutil.ReadAll(w.Result().Body)
	resultJSON := &registerRespond{}
	_ = json.Unmarshal(result, resultJSON)
	if resultJSON.Code != 1 || resultJSON.Message != "验证码错误" {
		t.Error("验证码错误但注册成功了")
	}
}

// Case5: 验证码不正确
func TestRegister5(t *testing.T) {
	server, _ := resource.NewServer()
	s := server.Service
	router := gin.Default()
	method := "POST"
	uri := "/api/register"
	router.POST(uri, Register(s))

	s.InsertVerifyCode("13812345678", "123456")
	defer s.DeleteVerifyCode("13812345678")

	contextTest := make(map[string]interface{})
	contextTest["UserName"] = "test2"
	contextTest["Password"] = "123456"
	contextTest["PhoneNumber"] = 13812345678
	contextTest["VerifyCode"] = "123456"

	contextByte, _ := json.Marshal(contextTest)
	req := httptest.NewRequest(method, uri, bytes.NewReader(contextByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result, _ := ioutil.ReadAll(w.Result().Body)
	resultJSON := &registerRespond{}
	_ = json.Unmarshal(result, resultJSON)
	if resultJSON.Code != 0 || resultJSON.Message != "注册成功" {
		t.Error("验证码错误但注册成功了")
	}
	if s.DeleteUserByPhoneNumber("13812345678") != true {
		t.Error("未正确注册")
	}
}
