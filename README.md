# 🛡️ RBAC Admin System

基于Go语言的轻量级RBAC（基于角色的访问控制）权限管理系统，支持MySQL、PostgreSQL和SQLite多种数据库。

## ✨ 功能特性

- 🔐 **完善的RBAC权限管理**：角色、权限、用户三位一体管理，支持细粒度权限控制
- 🗄️ **多数据库支持**：内置支持MySQL、SQLite、PostgreSQL三种数据库
- 🚀 **高性能**：基于Gin框架，路由设计清晰，性能卓越
- 📊 **完整的日志系统**：使用logrus实现完善的日志记录和查询
- 🔧 **灵活的配置系统**：YAML配置文件，支持多环境配置切换
- 🔑 **JWT认证**：基于Token的认证机制，支持会话管理
- 📧 **邮件服务**：支持邮件验证码发送和用户验证
- 🖼️ **图片上传**：支持文件上传和静态资源管理
- 🔒 **安全防护**：包含密码加密、输入验证等安全机制
- 🎯 **简单易用**：简洁的API设计，快速上手开发

## 🚀 快速开始

### 环境要求

- Go 1.21+ （推荐1.25.1或更高版本）
- MySQL 5.7+ 或 PostgreSQL 9.6+ 或 SQLite 3
- Redis（可选，用于缓存和验证码存储）

### 安装依赖

```bash
# 克隆项目后，进入项目目录
cd rbacAdmin

# 安装依赖
go mod download
```

### 配置数据库

复制并编辑配置文件：

```bash
cp settings.yaml.example settings.yaml
```

根据您的环境修改 `settings.yaml` 文件：

```yaml
# 系统配置
system:
    mode: debug  # debug, release
    ip: 127.0.0.1
    port: 8080

# 数据库配置
db:
    mode: mysql        # mysql, sqlite, postgres
    host: 127.0.0.1
    port: 3306
    user: root
    password: your_password
    db_name: rbacAdmin

# Redis配置
redis:
    addr: 127.0.0.1:6379
    password: 
    db: 3

# JWT配置
jwt:
    secret: your_jwt_secret
    expire: 24  # 单位小时
    issuer: rbacAdmin

# 验证码配置
captcha:
    # 验证码相关配置

# 邮件配置
email:
    # 邮件服务相关配置

# 上传配置
upload:
    # 文件上传相关配置
```

### 运行项目

```bash
# 开发环境运行
go run main.go

# 编译后运行
go build -o rbacAdmin.exe
./rbacAdmin.exe
```

## 📁 项目结构

```
rbacAdmin/
├── api/              # API接口层，处理HTTP请求
│   ├── captcha_api/  # 验证码相关API
│   ├── email_api/    # 邮件相关API
│   ├── image_api/    # 图片上传相关API
│   └── user_api/     # 用户相关API
├── config/           # 配置结构体定义
├── core/             # 核心功能实现
│   ├── db.go         # 数据库初始化与连接
│   ├── logger.go     # 日志系统配置
│   ├── casbin.go     # 权限控制初始化
│   ├── read_config.go # 配置文件读取
│   └── redis.go      # Redis连接配置
├── flags/            # 命令行参数处理
├── global/           # 全局变量定义
├── middleware/       # 中间件实现
│   ├── auth_middleware.go # JWT认证中间件
│   └── band_middleware.go # 绑定中间件
├── models/           # 数据模型定义
├── routes/           # 路由定义与注册
├── utils/            # 工具函数
│   ├── captcha/      # 验证码工具
│   ├── email/        # 邮件工具
│   ├── jwts/         # JWT工具
│   └── pwd/          # 密码加密工具
├── settings.yaml     # 主配置文件
├── settings_dev.yaml # 开发环境配置
└── main.go          # 程序主入口
```

## 🔧 核心功能模块

### 用户管理模块
- 用户注册、登录、注销
- 用户信息管理（修改昵称、头像、密码等）
- 用户列表查询与分页
- 管理员特权操作

### 权限管理模块
- 基于Casbin的RBAC权限控制
- 角色分配与管理
- 菜单权限配置
- API权限控制

### 数据模型

系统包含以下主要数据模型：
- **UserModel**: 用户信息模型
- **RoleModel**: 角色信息模型
- **MenuModel**: 菜单信息模型
- **APIModel**: API信息模型
- **UserRoleModel**: 用户角色关联模型
- **RoleMenuModel**: 角色菜单关联模型

### 日志系统

使用logrus实现的日志系统，支持多级别日志记录：
- 信息日志：记录系统正常运行状态
- 错误日志：记录系统异常和错误
- 调试日志：开发环境下的详细调试信息

## 🎯 API接口文档

### 用户相关接口

- `POST /api/login` - 用户登录
- `POST /api/register` - 用户注册
- `PUT /api/user/password` - 更新密码
- `PUT /api/users` - 更新用户信息
- `GET /api/users/info` - 获取用户信息
- `GET /api/users/list` - 获取用户列表

### 验证码接口

- 验证码生成与校验接口

### 邮件接口

- 邮件验证码发送接口

### 图片接口

- 图片上传接口

## 🔐 安全建议

请参考项目中的 [SECURITY.md](SECURITY.md) 文件，获取详细的安全配置指南和最佳实践。

## 📝 开发指南

### 添加新功能

1. 在 `models/` 目录定义数据模型
2. 在 `api/` 目录创建API处理器
3. 在 `routes/` 目录注册路由
4. 更新配置文件（如有必要）

### 中间件使用

目前系统实现了以下中间件：
- `AuthMiddleware`: JWT认证中间件，用于保护需要登录的接口
- `BindJson`: 请求参数绑定中间件
- `BindQuery`: URL查询参数绑定中间件

### 日志使用

```go
import "rbacAdmin/global"

// 使用全局日志器
global.Logger.Info("这是一条信息日志")
global.Logger.Error("这是一条错误日志")
```

## 🐛 常见问题

### Q: 数据库连接失败？
A: 检查数据库配置是否正确，确保数据库服务已启动，并验证用户名和密码是否匹配

### Q: Redis连接失败？
A: 检查Redis配置，确保Redis服务已启动，端口和密码配置正确

### Q: 端口被占用？
A: 修改 `settings.yaml` 中的端口配置，选择一个未被占用的端口

### Q: JWT认证失败？
A: 检查JWT配置中的secret是否正确，确保token未过期

## 🤝 贡献指南

欢迎提交Issue和Pull Request来改进这个项目。在提交PR前，请确保：
1. 代码风格与项目保持一致
2. 添加必要的注释和文档
3. 测试代码功能正常

## 📄 许可证

MIT License - 详见LICENSE文件

## 📞 联系方式

如有问题，请在GitHub上提交Issue。

---

⭐ 如果这个项目对你有帮助，请给个Star支持一下！
