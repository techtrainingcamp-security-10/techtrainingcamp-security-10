package api

type EnvironmentType struct {
	IP       string `json:"IP" binding:"required"`
	DeviceID string `json:"DeviceID" binding:"required"`
}

type ApplyCodeType struct {
	PhoneNumber uint            `json:"PhoneNumber" binding:"required"`
	Environment EnvironmentType `json:"Environment" binding:"required"`
}

type RegisterType struct {
	UserName    string          `json:"UserName" binding:"required"`
	Password    string          `json:"Password" binding:"required"`
	PhoneNumber uint            `json:"PhoneNumber" binding:"required"`
	VerifyCode  string          `json:"VerifyCode" binding:"required"`
	Environment EnvironmentType `json:"Environment" binding:"required"`
}

type LoginUIDType struct {
	UserName    string          `json:"UserName" binding:"required"`
	Password    string          `json:"Password" binding:"required"`
	Environment EnvironmentType `json:"Environment" binding:"required"`
}

type LoginPhoneType struct {
	PhoneNumber uint            `json:"PhoneNumber" binding:"required"`
	VerifyCode  string          `json:"VerifyCode" binding:"required"`
	Environment EnvironmentType `json:"Environment" binding:"required"`
}

type LogOutType struct {
	SessionID   string          `json:"SessionID" binding:"required"`
	ActionType  uint            `json:"ActionType" binding:"required"`
	Environment EnvironmentType `json:"Environment" binding:"required"`
}

type UserTable struct {
	UserName    string `json:"UserName" binding:"required"`
	Password    string `json:"Password" binding:"required"`
	PhoneNumber string `json:"PhoneNumber" binding:"required"`
	Salt        string `json:"Salt" binding:"required"`
}
