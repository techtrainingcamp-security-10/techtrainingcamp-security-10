package service

import (
	"crypto/rand"
	"fmt"
	"testing"
)

// Case1: 随机测试
func TestRandomPassword(t *testing.T) {
	const T = 1000 // 随机测试次数
	length, _ := randomBytes(T)
	for i := range length {
		length[i] = length[i]%13 + 6 // 设密码长度为6~18
	}

	passwords := make([]string, T)
	PasswordsAddSalt := make([]Password, T)
	for i := 0; i < T; i++ {
		passwords[i] = randomString(int(length[i]))
		PasswordsAddSalt[i] = NewPassword(passwords[i])
	}

	for i := 0; i < T; i++ {
		if !PasswordsAddSalt[i].Verify(passwords[i]) {
			t.Error("密码", passwords[i], "加密解密出错")
		}
	}
}

// Case2: 常见(?)密码测试
func TestUsualPassword(t *testing.T) {
	passwords := []string{
		"123123",
		"password",
		"qwerty",
		"/.,/.,",
		"p/q2-q4!",
		"qwqwqw",
		"3.1415926",
		"testtest",
		"qwerty",
		"qazwsxedc",
		"7894561230",
	}
	T := len(passwords)

	PasswordsAddSalt := make([]Password, T)
	for i := 0; i < T; i++ {
		PasswordsAddSalt[i] = NewPassword(passwords[i])
	}

	for i := 0; i < T; i++ {
		if !PasswordsAddSalt[i].Verify(passwords[i]) {
			t.Error("密码", passwords[i], "加密解密出错")
		}
	}
}

// 生成随机byte序列
func randomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return b, nil
}

// 生成指定字符集的随机字符串
var defaultLetters = []byte(",-.@$!%*#_~?&^abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var defaultMod = byte(len(defaultLetters))

func randomString(n int) string {
	b, _ := randomBytes(n)
	for i := range b {
		b[i] = defaultLetters[b[i]%defaultMod]
	}
	return string(b)
}
