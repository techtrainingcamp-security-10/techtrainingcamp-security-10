项目后端结构如下：
```
.
├── README.md
├── api                     api定义
│   ├── ApplyCode.go            获取验证码请求api
│   ├── LogOut.go               登出请求api
│   ├── Login.go                登录请求api
│   ├── Register.go             注册请求api
│   └── types.go                结构体类型定义
├── resource                资源相关
│   ├── db.go                   数据库，mysql
│   ├── redis.go                缓存，redis
│   └── resource.go             资源初始化
└── route                   路由
    ├── middleware              中间件
    │   ├── EnvCheck.go             环境检测
    │   ├── middleware.go           中间件注册
    │   └── service.go              数据库增删改查
    └── route.go                路由
```
依赖关系：  
```
resource <- route  
resource <- api  
resource <- middleware  
middleware <- route  
api <- route  
api <- middleware
```
请求返回参数：
1. 请求验证码
```json
// http://localhost:8080/api/apply_code [get]
{
    "PhoneNumber": 12345678901,
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
        "ExpireTime": 180,
        "VerifyCode": 1234
    },
    "Message": "请求成功"
}
```

2.请求注册
```json
// http://localhost:8080/api/register [post]
{
    "UserName": "user1",
    "Password": "123456",
    "PhoneNumber": 12345678901,
    "VerifyCode": "1234",
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
        "ExpireTime": 180,
        "SessionID": 123456
    },
    "Message": "注册成功",
    "SessionID": 123456
}
```
3. 请求登录(用户名)
```json
// http://localhost:8080/api/login_uid [post]
{
    "UserName": "user1",
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
        "ExpireTime": 180,
        "SessionID": 123456
    },
    "Message": "登录成功",
    "SessionID": 123456
}
```
4.请求登录(手机号)
```json
// http://localhost:8080/api/login_phone [post]
{
    "PhoneNumber": 12345678901,
    "VerifyCode": "1234",
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
        "ExpireTime": 180,
        "SessionID": 123456
    },
    "Message": "登录成功",
    "SessionID": 123456
}
```
5.请求登出
```json
// http://localhost:8080/api/logout [delete]
{
    "SessionID": "123456",
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