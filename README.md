# 《抓到你了——具备安全防护能力的账号系统》
## 项目规划
### 后端技术栈
- 服务：单机（多端口模拟分布式）
- 缓存：单机（Redis）
- 库：MySQL（主从实现读写分离）
- 接口：Restful

### 风险控制策略

风险控制策略分为如下四种：
- 拦截请求：本次请求返回失败
- 滑块验证：接下来的请求要求滑块验证
- 频率控制：接下来一段时间内的请求返回失败
- 帐号锁定：永久拒绝请求

<table>
<thead>
  <tr>
    <th>分类<br></th>
    <th>接口</th>
    <th>行为</th>
    <th>策略</th>
    <th>应对场景</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td rowspan="4">登录</td>
    <td rowspan="3">/api/login_uid<br>/api/login_phone<br></td>
    <td>一分钟内有失败且不多于3次<br></td>
    <td>滑块验证</td>
    <td rowspan="3">防止用户试图穷举密码</td>
  </tr>
  <tr>
    <td>一分钟内有失败且不多于10次，或5秒内多于或等于5次</td>
    <td>频率控制<br></td>
  </tr>
  <tr>
    <td>一分钟内失败次数多于10次</td>
    <td>帐号锁定</td>
  </tr>
  <tr>
    <td>/api/login_phone</td>
    <td>登录手机号为虚拟号段</td>
    <td>拦截请求</td>
    <td rowspan="2">确保手机号的真实性</td>
  </tr>
  <tr>
    <td rowspan="2">注册</td>
    <td rowspan="2">/api/register</td>
    <td>注册手机号为虚拟号段</td>
    <td>拦截请求</td>
  </tr>
  <tr>
    <td>用户名含敏感词</td>
    <td>拦截请求</td>
    <td>维护社区风气<br></td>
  </tr>
</tbody>
</table>

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