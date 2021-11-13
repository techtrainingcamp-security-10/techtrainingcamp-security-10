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

type logoutRespond struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
}

// Case1: 正常流程登出
func TestLogOut1(t *testing.T) {
	server, _ := resource.NewServer()
	s := server.Service
	router := gin.Default()
	method := "DELETE"
	uri := "/api/logout"
	router.DELETE(uri, LogOut(s))

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

	s.InsertSessionId("13812345678", "test")
	defer s.DeleteSessionId("test")

	contextTest := make(map[string]interface{})
	contextTest["SessionID"] = "test"
	contextTest["ActionType"] = 1

	contextByte, _ := json.Marshal(contextTest)
	req := httptest.NewRequest(method, uri, bytes.NewReader(contextByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result, _ := ioutil.ReadAll(w.Result().Body)
	resultJSON := &logoutRespond{}
	_ = json.Unmarshal(result, resultJSON)
	if resultJSON.Code != 0 || resultJSON.Message != "登出成功" {
		t.Error("test登出失败了")
	}
}

// Case2: SessionID 不存在
func TestLogOut2(t *testing.T) {
	server, _ := resource.NewServer()
	s := server.Service
	router := gin.Default()
	method := "DELETE"
	uri := "/api/logout"
	router.DELETE(uri, LogOut(s))

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

	//s.InsertSessionId("13812345678", "test")
	//defer s.DeleteSessionId("test")

	contextTest := make(map[string]interface{})
	contextTest["SessionID"] = "test"
	contextTest["ActionType"] = 1

	contextByte, _ := json.Marshal(contextTest)
	req := httptest.NewRequest(method, uri, bytes.NewReader(contextByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result, _ := ioutil.ReadAll(w.Result().Body)
	resultJSON := &logoutRespond{}
	_ = json.Unmarshal(result, resultJSON)
	if resultJSON.Code != 1 || resultJSON.Message != "登出失败" {
		t.Error("SessionID不存在登出成功了")
	}
}

// Case3: 正常流程注销
func TestLogOut3(t *testing.T) {
	server, _ := resource.NewServer()
	s := server.Service
	router := gin.Default()
	method := "DELETE"
	uri := "/api/logout"
	router.DELETE(uri, LogOut(s))

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

	s.InsertSessionId("13812345678", "test")
	defer s.DeleteSessionId("test")

	contextTest := make(map[string]interface{})
	contextTest["SessionID"] = "test"
	contextTest["ActionType"] = 2

	contextByte, _ := json.Marshal(contextTest)
	req := httptest.NewRequest(method, uri, bytes.NewReader(contextByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result, _ := ioutil.ReadAll(w.Result().Body)
	resultJSON := &logoutRespond{}
	_ = json.Unmarshal(result, resultJSON)
	if resultJSON.Code != 0 || resultJSON.Message != "注销成功" {
		t.Error("test注销失败了")
	}
}
