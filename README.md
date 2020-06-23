### 简介
> 用于监控各个dpi和分流器实时情况的后台
### 技术选型
- 前端：用基于vue的Element-UI构建基础页面。
- 后端：用Gin快速搭建基础restful风格API，Gin是一个go语言编写的Web框架。
- 数据库：采用MySql(5.6.44)版本，使用gorm实现对数据库的基本操作,已添加对sqlite数据库的支持。
- 缓存：使用Redis实现记录当前活跃用户的jwt令牌并实现多点登录限制。
- API文档：使用Swagger构建自动化文档。
- 配置文件：使用fsnotify和viper实现yaml格式的配置文件。
- 日志：使用go-logging实现日志记录。

### 基础功能
```shell script
1.权限+登录注册
2.操作日志
3.api文档
```

### 目录
```shell 
├─server
│  ├─api        (API)
│  ├─config     (配置包)
│  ├─core       (内核)
│  ├─db         (数据库脚本)
│  ├─docs       （swagger文档目录）
│  ├─global     （全局对象）
│  ├─initialiaze（初始化）
│  ├─middleware （中间件）
│  ├─model      （结构体层）
│  ├─resource   （资源）
│  ├─router     （路由）
│  ├─service    (服务)
│  └─utils      （公共功能）
└─web
```
