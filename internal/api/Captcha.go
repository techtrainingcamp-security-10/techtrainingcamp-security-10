package api

import (
	"bytes"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"

	"techtrainingcamp-security-10/internal/constants"
)

type CaptchaType struct {
	CaptchaId string `json:"captchaId"` //图形验证码Id
	ImageUrl  string `json:"imageUrl"`  //图形验证码图片url
}

// GetCaptcha
// @Description 获取图形验证码（验证码ID和图片url）
// @Router /api/get_captcha [get]
func GetCaptcha(context *gin.Context) {
	length := captcha.DefaultLen // 6
	captchaId := captcha.NewLen(length)
	var captchaServer CaptchaType
	captchaServer.CaptchaId = captchaId
	captchaServer.ImageUrl = "/captchaServer/" + captchaId + ".png"
	if captchaId == "" {
		context.JSON(constants.GETFailedCode, captchaServer)
	} else {
		context.JSON(constants.GETSuccessCode, captchaServer)
	}
}

// GetCaptchaImage
// @Description 获取验证码图片
// @Router /api/captcha/:captchaId [get]
func GetCaptchaImage(context *gin.Context) {
	ServeHTTP(context.Writer, context.Request)
}

// VerifyCaptcha
// @Description 验证验证码输入
// @Router /api/captcha/verify/:captchaId/:value [get]
func VerifyCaptcha(context *gin.Context) {
	captchaId := context.Param("captchaId")
	value := context.Param("value")
	if captchaId == "" || value == "" {
		context.String(constants.GETFailedCode, constants.VerifyCodeError)
	}
	if captcha.VerifyString(captchaId, value) {
		context.JSON(constants.GETSuccessCode, constants.SuccessCode)
	} else {
		context.JSON(constants.GETSuccessCode, constants.FailedCode)
	}
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dir, file := path.Split(r.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	if ext == "" || id == "" {
		http.NotFound(w, r)
		return
	}
	if r.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	lang := strings.ToLower(r.FormValue("lang"))
	download := path.Base(dir) == "download"
	if Serve(w, r, id, ext, lang, download, captcha.StdWidth, captcha.StdHeight) == captcha.ErrNotFound {
		http.NotFound(w, r)
	}
}

func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		err := captcha.WriteImage(&content, id, width, height)
		if err != nil {
			return err
		}
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		err := captcha.WriteAudio(&content, id, lang)
		if err != nil {
			return err
		}
	default:
		return captcha.ErrNotFound
	}
	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
