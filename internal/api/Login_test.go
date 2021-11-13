package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http/httptest"
	"techtrainingcamp-security-10/internal/resource"
	"techtrainingcamp-security-10/internal/service"
	"testing"
)

func TestGetSessionId(t *testing.T) {
	if getSessionId() == getSessionId() {
		t.Error("SessionID应当随机")
	}
}

type loginRespond struct {
	Code      int    `json:"Code"`
	Message   string `json:"Message"`
	SessionID string `json:"SessionID"`
}

// Case1：用户名不存在
func TestLoginByUID1(t *testing.T) {
	server, _ := resource.NewServer()
	s := server.Service
	router := gin.Default()
	method := "POST"
	uri := "/api/login_uid"
	router.POST(uri, LoginByUID(s))
	contextTest := make(map[string]string)
	contextTest["UserName"] = "test"
	contextTest["Password"] = "123456"

	contextByte, _ := json.Marshal(contextTest)
	req := httptest.NewRequest(method, uri, bytes.NewReader(contextByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result, _ := ioutil.ReadAll(w.Result().Body)
	resultJSON := &loginRespond{}
	_ = json.Unmarshal(result, resultJSON)
	if resultJSON.Code != 1 || resultJSON.Message != "用户名不存在" {
		t.Error("test用户不存在，但请求成功了")
	}
}

// Case2：正常流程用户名登录
func TestLoginByUID2(t *testing.T) {
	server, _ := resource.NewServer()
	s := server.Service
	router := gin.Default()
	method := "POST"
	uri := "/api/login_uid"
	router.POST(uri, LoginByUID(s))

	PasswordAddSalt := service.NewPassword("123456")
	testUser := service.UserTable{
		UserName:    "test",
		Password:    PasswordAddSalt,
		PhoneNumber: "13812345678",
	}
	err := s.InsertUser(testUser)
	if err != nil {
		t.Error("添加test用户失败了")
	}
	defer s.DeleteUserByPhoneNumber("13812345678")

	contextTest := make(map[string]string)
	contextTest["UserName"] = "test"
	contextTest["Password"] = "123456"

	contextByte, _ := json.Marshal(contextTest)
	req := httptest.NewRequest(method, uri, bytes.NewReader(contextByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result, _ := ioutil.ReadAll(w.Result().Body)
	resultJSON := &loginRespond{}
	_ = json.Unmarshal(result, resultJSON)
	if resultJSON.Code != 0 || resultJSON.Message != "登录成功" {
		t.Error("test登录失败了")
	}

}

// Case3：密码错误
func TestLoginByUID3(t *testing.T) {
	server, _ := resource.NewServer()
	s := server.Service
	router := gin.Default()
	method := "POST"
	uri := "/api/login_uid"
	router.POST(uri, LoginByUID(s))

	PasswordAddSalt := service.NewPassword("123456")
	testUser := service.UserTable{
		UserName:    "test",
		Password:    PasswordAddSalt,
		PhoneNumber: "13812345678",
	}
	err := s.InsertUser(testUser)
	if err != nil {
		t.Error("添加test用户失败了")
	}
	defer s.DeleteUserByPhoneNumber("13812345678")

	contextTest := make(map[string]string)
	contextTest["UserName"] = "test"
	contextTest["Password"] = "1234567"

	contextByte, _ := json.Marshal(contextTest)
	req := httptest.NewRequest(method, uri, bytes.NewReader(contextByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result, _ := ioutil.ReadAll(w.Result().Body)
	resultJSON := &loginRespond{}
	_ = json.Unmarshal(result, resultJSON)
	if resultJSON.Code != 1 || resultJSON.Message != "用户名或密码错误" {
		t.Error("test密码错误，但是登录成功了")
	}

}

// Case1：验证码无效
func TestLoginByPhone1(t *testing.T) {
	server, _ := resource.NewServer()
	s := server.Service
	router := gin.Default()
	method := "POST"
	uri := "/api/login_phone"
	router.POST(uri, LoginByPhone(s))
	s.DeleteVerifyCode("13812345678")
	contextTest := make(map[string]interface{})
	contextTest["PhoneNumber"] = 13812345678
	contextTest["VerifyCode"] = "123456"

	contextByte, _ := json.Marshal(contextTest)
	req := httptest.NewRequest(method, uri, bytes.NewReader(contextByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result, _ := ioutil.ReadAll(w.Result().Body)
	resultJSON := &loginRespond{}
	_ = json.Unmarshal(result, resultJSON)
	if resultJSON.Code != 1 || resultJSON.Message != "验证码无效" {
		fmt.Println(resultJSON)
		t.Error("验证码无效但登录成功了")
	}

}

// Case2： 正常流程手机号登录
func TestLoginByPhone2(t *testing.T) {
	server, _ := resource.NewServer()
	s := server.Service
	router := gin.Default()
	method := "POST"
	uri := "/api/login_phone"
	router.POST(uri, LoginByPhone(s))

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
	contextTest["PhoneNumber"] = 13812345678
	contextTest["VerifyCode"] = "123456"

	contextByte, _ := json.Marshal(contextTest)
	req := httptest.NewRequest(method, uri, bytes.NewReader(contextByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result, _ := ioutil.ReadAll(w.Result().Body)
	resultJSON := &loginRespond{}
	_ = json.Unmarshal(result, resultJSON)
	if resultJSON.Code != 0 || resultJSON.Message != "登录成功" {
		t.Error("test登录失败了")
	}

}

// Case3：验证码不正确
func TestLoginByPhone3(t *testing.T) {
	server, _ := resource.NewServer()
	s := server.Service
	router := gin.Default()
	method := "POST"
	uri := "/api/login_phone"
	router.POST(uri, LoginByPhone(s))

	PasswordAddSalt := service.NewPassword("123456")
	testUser := service.UserTable{
		UserName:    "test",
		Password:    PasswordAddSalt,
		PhoneNumber: "13812345678",
	}
	s.InsertVerifyCode("13812345678", "654321")
	defer s.DeleteVerifyCode("13812345678")
	err := s.InsertUser(testUser)
	if err != nil {
		t.Error("添加test用户失败了")
	}
	defer s.DeleteUserByPhoneNumber("13812345678")

	contextTest := make(map[string]interface{})
	contextTest["PhoneNumber"] = 13812345678
	contextTest["VerifyCode"] = "123456"

	contextByte, _ := json.Marshal(contextTest)
	req := httptest.NewRequest(method, uri, bytes.NewReader(contextByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result, _ := ioutil.ReadAll(w.Result().Body)
	resultJSON := &loginRespond{}
	_ = json.Unmarshal(result, resultJSON)
	if resultJSON.Code != 1 || resultJSON.Message != "验证码错误" {
		t.Error("验证码错误但登录成功了")
	}
}

// Case4：虚拟号段或无效手机号
func TestLoginByPhone4(t *testing.T) {
	server, _ := resource.NewServer()
	s := server.Service
	router := gin.Default()
	method := "POST"
	uri := "/api/login_phone"
	router.POST(uri, LoginByPhone(s))

	contextTest := make(map[string]interface{})
	contextTest["PhoneNumber"] = 12812345678
	contextTest["VerifyCode"] = "123456"

	contextByte, _ := json.Marshal(contextTest)
	req := httptest.NewRequest(method, uri, bytes.NewReader(contextByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result, _ := ioutil.ReadAll(w.Result().Body)
	resultJSON := &loginRespond{}
	_ = json.Unmarshal(result, resultJSON)
	if resultJSON.Code != 1 || resultJSON.Message != "请换个手机号试试" {
		t.Error("不合法手机号但请求成功了")
	}
}

// Case5：用户不存在
func TestLoginByPhone5(t *testing.T) {
	server, _ := resource.NewServer()
	s := server.Service
	router := gin.Default()
	method := "POST"
	uri := "/api/login_phone"
	router.POST(uri, LoginByPhone(s))

	s.InsertVerifyCode("13812345678", "123456")
	defer s.DeleteVerifyCode("13812345678")

	contextTest := make(map[string]interface{})
	contextTest["PhoneNumber"] = 13812345678
	contextTest["VerifyCode"] = "123456"

	contextByte, _ := json.Marshal(contextTest)
	req := httptest.NewRequest(method, uri, bytes.NewReader(contextByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result, _ := ioutil.ReadAll(w.Result().Body)
	resultJSON := &loginRespond{}
	_ = json.Unmarshal(result, resultJSON)
	if resultJSON.Code != 1 || resultJSON.Message != "手机号未注册" {
		t.Error("手机号未注册但登录成功了")
	}
}
