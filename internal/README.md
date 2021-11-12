### 项目文件结构
```
.
├── README.md
├── docker-compose.yml      数据库、缓存 docker 配置
├── go.mod                  go module
├── go.sum
├── internal
│   ├── README.md
│   ├── api                     APIs
│   │   ├── ApplyCode.go            请求手机验证码 API
│   │   ├── Captcha.go              请求图片验证码 API(未启用)
│   │   ├── LogOut.go               登出/注销 API
│   │   ├── Login.go                用户名/手机号登录 API
│   │   └── Register.go             注册 API
│   ├── constants               参数
│   │   └── types.go                接口结构体定义
│   ├── resource                资源
│   │   ├── config.example.yml      数据库、缓存连接配置(样例)
│   │   ├── config.go               数据库、缓存连接配置解析
│   │   ├── config.yml              数据库、缓存连接配置(已隐藏)
│   │   ├── db.go                   数据库初始化
│   │   ├── redis.go                缓存初始化
│   │   └── resource.go             资源初始化
│   ├── route                   路由
│   │   └── route.go                路由
│   ├── service                 封装的服务
│   │   ├── messages.go             数据库表格定义
│   │   ├── passwords.go            密码加密校验类
│   │   ├── service.go              封装的服务
│   │   ├── service_db.go           数据库增删改查接口
│   │   └── service_redis.go        缓存增删改查接口
│   └── utils                   工具
│       ├── EnvCheck.go             环境检测
│       ├── FailRecordsCheck.go     失败次数检测
│       ├── sensitiveWords.go       敏感词检测
│       ├── virtualPhoneNumber.go   虚拟号段检测
│       └── virtualPhoneNumber_test.go
├── keywords.txt            敏感词列表(已隐藏)
└── main.go                 主程序入口
```
层级关系：  
```
 route
   |
  API
   |
service
   |
resource
```
### 请求返回 JSON 样例（成功）
1. 请求验证码  
[get] http://localhost:8080/api/apply_code 
```json
{
  "PhoneNumber": 13812345678,
  "Environment": {
    "IP": "127.0.0.1",
    "DeviceID": "123"
  }
}
```
```json
{
  "Code": 0,
  "Data": {
    "DecisionType": 0,
    "ExpireTime": 10800,
    "VerifyCode": "QnzJRm"
  },
  "Message": "请求成功"
}
```

2. 请求注册  
[post] http://localhost:8080/api/register
```json
{
  "UserName": "user_test",
  "Password": "123456",
  "PhoneNumber": 13812345678,
  "VerifyCode": "QnzJRm",
  "Environment": {
    "IP": "127.0.0.1",
    "DeviceID": "123"
  }
}
```
```json
{
  "Code": 0,
  "Data": {
    "DecisionType": 0,
    "ExpireTime": 10800,
    "SessionID": "fe985458-11ec-49e9-90cc-4929dbeb6ef4"
  },
  "Message": "注册成功"
}
```
3. 请求登录(用户名)  
[post] http://localhost:8080/api/login_uid 
```json
{
  "UserName": "user_test",
  "Password": "123456",
  "Environment": {
    "IP": "127.0.0.1",
    "DeviceID": "123"
  }
}
```
```json
{
  "Code": 0,
  "Data": {
    "DecisionType": 0,
    "ExpireTime": 10800,
    "SessionID": "ab4e1129-9069-4077-a68d-d584e3b0c78b"
  },
  "Message": "登录成功",
  "SessionID": "ab4e1129-9069-4077-a68d-d584e3b0c78b"
}
```
4. 请求登录(手机号)  
[post] http://localhost:8080/api/login_phone
```json
{
  "PhoneNumber": 13812345678,
  "VerifyCode": "wOYeHr",
  "Environment": {
    "IP": "127.0.0.1",
    "DeviceID": "123"
  }
}
```
```json
{
  "Code": 0,
  "Data": {
    "DecisionType": 0,
    "ExpireTime": 10800,
    "SessionID": "1aafcb61-5eba-4d72-85d0-f038c2f29108"
  },
  "Message": "登录成功",
  "SessionID": "1aafcb61-5eba-4d72-85d0-f038c2f29108"
}
```
5. 请求登出  
[delete] http://localhost:8080/api/logout
```json
{
  "SessionID": "1aafcb61-5eba-4d72-85d0-f038c2f29108",
  "ActionType": 1,
  "Environment": {
    "IP": "127.0.0.1",
    "DeviceID": "123"
  }
}
```
```json
{
  "Code": 0,
  "Message": "登出成功"
}
```