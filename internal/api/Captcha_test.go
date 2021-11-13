package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

// 注：图片显示是无法测试的，但应该不会有问题

type captchaResponse struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
}

// Case1: 随机测试
func TestRandom(t *testing.T) {
	const T = 30 // 随机测试次数

	store := captcha.NewMemoryStore(captcha.CollectNum, captcha.Expiration)
	captcha.SetCustomStore(store)

	for i := 0; i < T; i++ {
		singleTest(store, t)
	}
}

var letters = []byte("0123456789")
var mod = len(defaultLetters)

func singleTest(store captcha.Store, t *testing.T) {
	router := gin.Default()
	method := "GET"
	uri := "/api/get_captcha"
	router.GET(uri, GetCaptcha)

	var captchaId string
	{
		contextTest := make(map[string]interface{})

		contextByte, _ := json.Marshal(contextTest)
		req := httptest.NewRequest(method, uri, bytes.NewReader(contextByte))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		result, _ := ioutil.ReadAll(w.Result().Body)
		resultJSON := &CaptchaType{}
		_ = json.Unmarshal(result, resultJSON)

		captchaId = resultJSON.CaptchaId
		// imageUrl = resultJSON.ImageUrl
	}
	ans := store.Get(captchaId, false) // []byte 注意不要修改此切片（会更改store的内容）

	ans0_s := ""
	for i := range ans {
		ans0_s += strconv.Itoa(int(ans[i]))
	}
	ans_s := ans0_s

	// 1/3概率使用1位错误验证码
	modified := 0
	if rand.Intn(3) >= 2 {
		modified = 1
		p := rand.Intn(len(ans_s))
		v := letters[rand.Intn(mod)]
		for v == ans_s[p] {
			v = letters[rand.Intn(mod)]
		}
		ans_s = ans_s[:p] + string(v) + ans_s[p+1:]
	}

	router = gin.Default()
	uri = "/api/captcha/verify/:captchaId/:value"
	uri_real := "/api/captcha/verify/" + captchaId + "/" + ans_s
	router.GET(uri, VerifyCaptcha)
	{
		contextTest := make(map[string]interface{})

		contextByte, _ := json.Marshal(contextTest)
		req := httptest.NewRequest(method, uri_real, bytes.NewReader(contextByte))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req) // 调用这部分要重新 router=gin.Default()，否则会panic
		result, _ := ioutil.ReadAll(w.Result().Body)
		resultJSON := &captchaResponse{}
		_ = json.Unmarshal(result, resultJSON)

		if resultJSON.Code != modified {
			t.Error("验证码测试错误：", ans0_s, ans_s, resultJSON.Code, modified)
		}
	}
}
