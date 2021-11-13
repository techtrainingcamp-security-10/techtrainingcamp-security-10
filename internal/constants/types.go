package constants

type EnvironmentType struct {
	IP       string `json:"IP" binding:"required"`
	DeviceID string `json:"DeviceID" binding:"required"`
}

type ApplyCodeType struct {
	PhoneNumber uint `json:"PhoneNumber" binding:"required"`
	//Environment EnvironmentType `json:"Environment" binding:"required"`
}

type RegisterType struct {
	UserName    string `json:"UserName" binding:"required"`
	Password    string `json:"Password" binding:"required"`
	PhoneNumber uint   `json:"PhoneNumber" binding:"required"`
	VerifyCode  string `json:"VerifyCode" binding:"required"`
	//Environment EnvironmentType `json:"Environment" binding:"required"`
}

type LoginUIDType struct {
	UserName string `json:"UserName" binding:"required"`
	Password string `json:"Password" binding:"required"`
	//Environment EnvironmentType `json:"Environment" binding:"required"`
}

type LoginPhoneType struct {
	PhoneNumber uint   `json:"PhoneNumber" binding:"required"`
	VerifyCode  string `json:"VerifyCode" binding:"required"`
	//Environment EnvironmentType `json:"Environment" binding:"required"`
}

type LogOutType struct {
	SessionID  string `json:"SessionID" binding:"required"`
	ActionType uint   `json:"ActionType" binding:"required"`
	//Environment EnvironmentType `json:"Environment" binding:"required"`
}

const (
	UserNameAlreadyExists    = "相同的用户名已经被注册过了，请更换用户名试试"
	PhoneNumberAlreadyExists = "相同的手机号已经被注册过了，请更换用户名试试"
	RegisterSuccess          = "注册成功"
	VerifyCodeInvalid        = "验证码无效"
	VerifyCodeError          = "验证码错误"
	UserNameNotRegister      = "用户名不存在"
	LoginFailed              = "用户名或密码错误"
	LoginSuccess             = "登录成功"
	PhoneNumberNotRegister   = "手机号未注册"
	UserLoginStateInvalid    = "用户登录状态失效"
	LogOutSuccess            = "登出成功"
	CancellationSuccess      = "注销成功"
	RequestSuccess           = "请求成功"
	LogOutFailed             = "登出失败"
	PhoneNumberStateErr      = "请换个手机号试试"
	UserNameErr              = "请换个用户名试试"
	FrequencyLimit           = "请稍后再试"
	Lock                     = "你已被封"
)

const (
	GETSuccessCode    = 200
	GETFailedCode     = 400
	POSTSuccessCode   = 201
	POSTFailedCode    = 400
	DELETESuccessCode = 201
	DELETEFailedCode  = 400
	SuccessCode       = 0
	FailedCode        = 1
	Normal            = 0
	SlideBar          = 1
	FrequentLimit     = 2
	Locked            = 3
)
