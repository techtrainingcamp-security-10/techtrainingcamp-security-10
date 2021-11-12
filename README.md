# 《抓到你了——具备安全防护能力的账号系统》
## 项目规划
### 后端技术栈
- 服务：单机（多端口模拟分布式）
- 缓存：单机（Redis）
- 库：MySQL（主从实现读写分离）
- 接口：Restful
- 风控：
  - 滑块：
    - 连续错误次数大于2 概率返回
  - 频控：
    - 5分钟内错误次数大于5 后端概率返回
    - 注册用户名不符合规范
  - 拦截：
    - 1分钟内同IP/ENV请求次数大于15
    - IP非国内
    - 手机号非国内
    - 黑名单用户
    - 错误次数大于20
    - 已注销用户
    - 连续频控多次

### 需要实现的功能及分工
请访问 [Projects](https://github.com/techtrainingcamp-security-10/techtrainingcamp-security-10/projects)

## 第三方依赖
| 名称         | 版本  | 主页                                       |
| ----------- | ----- | ----------------------------------------- |
| gin         | 1.17  | https://github.com/gin-gonic/gin          |
| gorm        | 1.9.16| https://github.com/jinzhu/gorm            |
| redigo      | 1.8.5 | https://github.com/gomodule/redigo        |
| yaml        | 2.4.0 | https://github.com/go-yaml/yaml           |
| iploc       | 1.0.2 | https://github.com/phuslu/iploc           |
| uuid        | 4.1.0 | https://github.com/gofrs/uuid             |