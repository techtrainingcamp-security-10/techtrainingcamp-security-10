package service

type UserTable struct {
	UserName    string `json:"UserName" binding:"required"`
	Password    string `json:"Password" binding:"required"`
	PhoneNumber string `json:"PhoneNumber" binding:"required"`
	Salt        string `json:"Salt" binding:"required"`
}

const (
	UserNameAlreadyExists    = "用户名已存在"
	PhoneNumberAlreadyExists = "电话号码已被注册"
	RegisterSuccess          = "注册成功"
	VerifyCodeInvalid        = "验证码无效"
	VerifyCodeError          = "验证码错误"
	LoginFailed              = "用户名或密码错误"
	LoginSuccess             = "登录成功"
	PhoneNumberNotRegister   = "手机号未注册"
	UserLoginStateInvalid    = "用户登录状态失效"
	LogOutSuccess            = "登出成功"
	CancellationSuccess      = "注销成功"
)

const (
	VerifyCodeExpireTime = 3 * 60
	SessionIdExpireTime  = 3 * 60 * 60
	SplitChar            = "|"
)
