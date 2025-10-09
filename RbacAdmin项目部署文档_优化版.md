# RbacAdmin项目部署文档

## 版本信息
**版本号**: v1.0
**更新日期**: 2024-07-05
**文档状态**: 正式版
**Go版本**: 1.25.1

## 1. 项目概况
RbacAdmin是一个企业级轻量级RBAC权限管理系统，专为中小型应用提供灵活、高效的权限控制解决方案。系统基于Go语言开发，采用现代化的架构设计，支持多种数据库，并提供简洁易用的API接口。

### 核心功能特性
- **完善的RBAC权限管理**：支持用户、角色、菜单、API的精细化权限控制
- **多数据库支持**：兼容MySQL、SQLite、PostgreSQL等主流数据库
- **高性能架构**：基于Go语言的高并发特性，提供卓越的性能表现
- **日志系统**：集成结构化日志管理，便于问题排查和系统监控
- **灵活配置**：基于YAML的配置管理，支持多环境配置切换
- **简洁API设计**：提供清晰规范的RESTful API接口
- **验证码系统**：支持图片验证码和邮件验证码
- **邮箱服务**：集成邮件发送功能，支持账户验证和通知

### 典型应用场景
- 企业内部管理系统的权限控制
- SaaS平台的多租户权限隔离
- 需要精细权限管理的Web应用
- 微服务架构中的认证授权中心

## 2. 项目技术栈

### 基础技术
- **编程语言**：Go 1.25.1
- **构建工具**：Go Module
- **Web框架**: Gin Web Framework v1.10.1
- **ORM框架**: GORM v1.31.0
- **缓存**: Redis v6.15.9
- **权限控制**: Casbin v2.123.0
- **日志系统**: Logrus v1.9.3

### 核心框架与依赖
| 技术/框架 | 版本 | 用途 | 来源 |
|---------|------|------|------|
| Gin Web Framework | 1.10.1 | HTTP Web框架 | <mcfile name="go.mod" path="e:\myblog\Go项目学习\rbacAdmin\go.mod"></mcfile> |
| GORM | 1.31.0 | 数据库ORM框架 | <mcfile name="go.mod" path="e:\myblog\Go项目学习\rbacAdmin\go.mod"></mcfile> |
| go-redis/redis | v6.15.9+incompatible | Redis缓存客户端 | <mcfile name="go.mod" path="e:\myblog\Go项目学习\rbacAdmin\go.mod"></mcfile> |
| sirupsen/logrus | 1.9.3 | 结构化日志系统 | <mcfile name="go.mod" path="e:\myblog\Go项目学习\rbacAdmin\go.mod"></mcfile> |
| gopkg.in/yaml.v3 | v3.0.1 | YAML配置文件解析 | <mcfile name="go.mod" path="e:\myblog\Go项目学习\rbacAdmin\go.mod"></mcfile> |
| casbin | v2.123.0 | 权限控制框架 | <mcfile name="go.mod" path="e:\myblog\Go项目学习\rbacAdmin\go.mod"></mcfile> |
| gorm-adapter | v3.37.0 | Casbin的GORM适配器 | <mcfile name="go.mod" path="e:\myblog\Go项目学习\rbacAdmin\go.mod"></mcfile> |
| base64Captcha | v1.3.8 | 验证码生成库 | <mcfile name="go.mod" path="e:\myblog\Go项目学习\rbacAdmin\go.mod"></mcfile> |

### 运行环境依赖
- **操作系统**：支持Windows、Linux、macOS等主流操作系统
- **数据库**：MySQL 5.7+ / SQLite 3+ / PostgreSQL 9.6+
- **可选组件**：Redis 4.0+（用于缓存）
- **运行内存**：建议2GB以上
- **磁盘空间**：建议10GB以上

## 3. 项目目录结构
RbacAdmin项目采用清晰的模块化结构设计，遵循Go语言的标准项目布局规范。核心业务逻辑与配置、模型、路由等组件分离，便于维护和扩展。

### 整体架构
```text
rbacAdmin/
├── api/             # API接口定义及处理
├── cmd/             # 命令行工具和子命令
├── config/          # 配置结构体定义
├── core/            # 核心功能实现（日志、配置读取、数据库初始化等）
├── flags/           # 命令行参数处理
├── global/          # 全局变量定义
├── logs/            # 日志文件存储
├── middleware/      # HTTP中间件
├── models/          # 数据模型定义
├── routes/          # 路由配置
├── service/         # 业务逻辑层
├── uploads/         # 上传文件存储
├── utils/           # 工具函数
├── go.mod           # Go模块依赖
├── go.sum           # 依赖版本锁定
├── main.go          # 应用程序入口
├── rbacAdmin.exe    # Windows可执行文件
├── settings.yaml    # 主配置文件
├── settings.yaml.example  # 配置文件模板
└── RbacAdmin项目部署文档.md    # 部署文档
```

### 主要目录功能说明
| 目录/文件 | 主要职责 | 文件位置 |
|----------|---------|---------|
| api/ | API接口定义及处理逻辑 | <mcfile name="api" path="e:\myblog\Go项目学习\rbacAdmin\api"></mcfile> |
| config/ | 配置结构体定义，包括系统、数据库、Redis等配置 | <mcfile name="config" path="e:\myblog\Go项目学习\rbacAdmin\config"></mcfile> |
| core/ | 核心功能实现，包括配置读取、数据库初始化、日志初始化等 | <mcfile name="core" path="e:\myblog\Go项目学习\rbacAdmin\core"></mcfile> |
| flags/ | 命令行参数解析和处理 | <mcfile name="flags" path="e:\myblog\Go项目学习\rbacAdmin\flags"></mcfile> |
| global/ | 全局变量定义，方便跨模块访问配置、数据库等 | <mcfile name="global" path="e:\myblog\Go项目学习\rbacAdmin\global"></mcfile> |
| models/ | 数据模型定义，映射数据库表结构 | <mcfile name="models" path="e:\myblog\Go项目学习\rbacAdmin\models"></mcfile> |
| main.go | 应用程序入口文件，定义启动流程 | <mcfile name="main.go" path="e:\myblog\Go项目学习\rbacAdmin\main.go"></mcfile> |
| settings.yaml | 系统配置文件，包含数据库连接、服务器设置等 | <mcfile name="settings.yaml" path="e:\myblog\Go项目学习\rbacAdmin\settings.yaml"></mcfile> |

## 4. 项目部署步骤

### 4.1 环境准备
在部署RbacAdmin之前，请确保目标服务器满足以下环境要求：

#### 4.1.1 安装Go环境（开发环境需要）
```bash
# Linux系统安装示例（以Ubuntu为例）
sudo apt update
sudo apt install golang-go

# 验证安装
go version
# 应显示类似: go version go1.25.1 linux/amd64
```

#### 4.1.2 安装数据库
根据项目需求选择并安装相应的数据库：

**MySQL安装示例**：
```bash
# Ubuntu系统安装MySQL
sudo apt update
sudo apt install mysql-server

sudo systemctl start mysql
sudo systemctl enable mysql

sudo mysql_secure_installation
```

**创建数据库**：
```bash
mysql -u root -p
CREATE DATABASE rbacAdmin DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
GRANT ALL PRIVILEGES ON rbacAdmin.* TO 'root'@'localhost' IDENTIFIED BY 'your_password';
FLUSH PRIVILEGES;
EXIT;
```

#### 4.1.3 安装Redis（可选）
```bash
# Ubuntu系统安装Redis
sudo apt update
sudo apt install redis-server

sudo systemctl start redis
sudo systemctl enable redis
```

### 4.2 项目获取与编译

#### 4.2.1 获取项目源码
```bash
# 克隆项目仓库
git clone https://github.com/zhu6743767/rbacAdmin.git
cd rbacAdmin
```

#### 4.2.2 安装依赖
```bash
go mod download
```

#### 4.2.3 编译项目
```bash
# 编译为当前系统可执行文件
go build -o rbacAdmin

# 交叉编译为其他系统（例如Windows）
GOOS=windows GOARCH=amd64 go build -o rbacAdmin.exe
```

### 4.3 配置文件设置

#### 4.3.1 配置文件模板
项目提供了`settings.yaml.example`作为配置模板，请根据实际环境创建并修改配置文件：

```bash
# 复制配置模板
cp settings.yaml.example settings.yaml
```

#### 4.3.2 配置文件内容详解
配置文件主要包含以下几个部分：

```yaml
# 系统配置
system:
    mode: debug       # 运行模式: debug, release, test
    ip: 127.0.0.1     # 监听IP地址
    port: 8080        # 监听端口

# 数据库配置
db:
    # 数据库类型：mysql, sqlite, postgres
    mode: mysql
    host: localhost    # 数据库主机地址
    port: 3306         # 数据库端口
    user: root         # 数据库用户名
    password: "your_secure_password_here"  # 数据库密码
    db_name: rbacAdmin # 数据库名称

# Redis配置（可选）
redis:
    addr: localhost:6379  # Redis服务器地址
    password: ""            # Redis密码（如无密码留空）
    db: 0                   # Redis数据库编号（0-15）

# JWT配置
jwt:
    secret: "your_jwt_secret_key"  # JWT签名密钥
    expire: 24                      # 令牌有效期(小时) - 注意：使用expire而不是expires
    issuer: "rbacAdmin"            # 令牌签发者

# 验证码配置
captcha:
    enable: true  # 是否启用验证码

# 邮箱配置
email:
    user: "your_email@example.com"  # 邮箱账号
    password: "your_email_password"  # 邮箱密码/授权码
    Host: "smtp.example.com"         # SMTP服务器地址
    Port: 587                       # SMTP服务器端口
```

根据实际环境修改相应的配置项，如数据库连接信息、Redis连接信息、JWT密钥、邮箱配置等。

**重要说明：**
1. JWT配置中的令牌有效期字段为`expire`（单数形式），而不是`expires`
2. 认证中间件中的`GetAuth`函数返回的是`jwts.ClaimsUserInfo`类型，用于获取当前登录用户信息
3. 对于文件上传功能，需要确保`uploads/file/{username}`目录存在且具有写入权限

### 4.4 数据库初始化
在首次部署或更新后，需要执行数据库迁移以创建或更新数据表结构：

```bash
# 执行数据库迁移
./rbacAdmin -db
# 或在Windows环境下
rbacAdmin.exe -db
```

数据库迁移将自动创建以下表：
- 用户表(users)
- 角色表(roles)
- 菜单表(menus)
- API权限表(apis)
- 用户角色关联表(user_role_models)
- 角色菜单关联表(role_menu_models)
- Casbin规则表(casbin_rules)

### 4.5 启动应用

#### 4.5.1 直接启动
```bash
# Linux/macOS
./rbacAdmin

# Windows
rbacAdmin.exe
```

#### 4.5.2 指定配置文件启动
```bash
# Linux/macOS
./rbacAdmin -f settings_dev.yaml

# Windows
rbacAdmin.exe -f settings_dev.yaml
```

#### 4.5.3 作为系统服务运行（Linux系统）

**创建systemd服务文件**：
```bash
sudo nano /etc/systemd/system/rbacAdmin.service
```

**服务文件内容**：
```ini
[Unit]
Description=RbacAdmin RBAC Permission Management System
After=network.target mysql.service redis.service

[Service]
Type=simple
WorkingDirectory=/path/to/rbacAdmin
ExecStart=/path/to/rbacAdmin/rbacAdmin
Restart=on-failure
RestartSec=5s
User=www-data
Group=www-data

[Install]
WantedBy=multi-user.target
```

**启用并启动服务**：
```bash
sudo systemctl daemon-reload
sudo systemctl enable rbacAdmin
sudo systemctl start rbacAdmin
sudo systemctl status rbacAdmin
```

### 4.6 验证部署
应用启动后，可以通过以下方式验证部署是否成功：

1. 查看应用日志：
   ```bash
tail -f logs/app.log
```

2. 访问API接口：
   ```bash
curl http://服务器IP:8080/api/v1/system/info
```

## 5. 项目配置详解

### 5.1 配置文件结构
RbacAdmin项目使用YAML格式的配置文件，主要包含以下核心部分：

```go
type Config struct {
    System  SystemConfig `yaml:"system"`
    DB      DB           `yaml:"db"`
    Redis   Redis        `yaml:"redis"`
    JWT     JWT          `yaml:"jwt"`
    Captcha Captcha      `yaml:"captcha"`
    Email   Email        `yaml:"email"`
}
```
<mcfile name="config\enter.go" path="e:\myblog\Go项目学习\rbacAdmin\config\enter.go"></mcfile>

### 5.2 系统配置
系统配置定义在`config/system.go`中，包含服务器监听的IP地址、端口和运行模式：

```go
type SystemConfig struct {
    Mode string `yaml:"mode"` // 运行模式: debug, release, test
    Ip   string `yaml:"ip"`   // 监听IP地址
    Port int    `yaml:"port"` // 监听端口
}

func (s *SystemConfig) Addr() string {
    return fmt.Sprintf("%s:%d", s.Ip, s.Port)
}
```
<mcfile name="config\system.go" path="e:\myblog\Go项目学习\rbacAdmin\config\system.go"></mcfile>

### 5.3 数据库配置
数据库配置定义在`config/db.go`中，支持MySQL、SQLite和PostgreSQL三种数据库类型：

```go
type DB struct {
    MODE     string `yaml:"mode"`
    HOST     string `yaml:"host"`
    PORT     int    `yaml:"port"`
    USER     string `yaml:"user"`
    PASSWORD string `yaml:"password"`
    DbNAME   string `yaml:"db_name"`
}
```
<mcfile name="config\db.go" path="e:\myblog\Go项目学习\rbacAdmin\config\db.go"></mcfile>

### 5.4 Redis配置
Redis配置定义在`config/redis.go`中，用于配置缓存服务：

```go
type Redis struct {
    Addr     string `yaml:"addr"`     // Redis服务器地址，格式为host:port
    Password string `yaml:"password"` // Redis密码（如无密码留空）
    DB       int    `yaml:"db"`       // Redis数据库编号（0-15）
}
```
<mcfile name="config\redis.go" path="e:\myblog\Go项目学习\rbacAdmin\config\redis.go"></mcfile>

### 5.5 JWT配置
JWT配置定义了系统的认证令牌相关设置：

```go
type JWT struct {
    Secret string `yaml:"secret"` // JWT签名密钥
    Expire int    `yaml:"expire"` // 令牌有效期(小时)
    Issuer string `yaml:"issuer"` // 令牌签发者
}
```
<mcfile name="config\jwt.go" path="e:\myblog\Go项目学习\rbacAdmin\config\jwt.go"></mcfile>

### 5.6 验证码配置
验证码配置定义了是否启用验证码功能：

```go
type Captcha struct {
    Enable bool `yaml:"enable"`
}
```
<mcfile name="config\captcha.go" path="e:\myblog\Go项目学习\rbacAdmin\config\captcha.go"></mcfile>

### 5.7 邮箱配置
邮箱配置定义了邮件发送相关的设置：

```go
type Email struct {
    User     string `yaml:"user"`
    Password string `yaml:"password"`
    Host     string `yaml:"Host"`
    Port     int    `yaml:"Port"`
}

func (e Email) Verify() bool {
    if e.User == "" || e.Password == "" || e.Host == "" || e.Port == 0 {
        return false
    }
    return true
}
```
<mcfile name="config\email.go" path="e:\myblog\Go项目学习\rbacAdmin\config\email.go"></mcfile>

### 5.8 命令行参数
命令行参数的定义在`flags/enter.go`中：

```go
func init() {
    flag.StringVar(&FlagOptions.File, "f", "settings.yaml", "配置文件路径")
    flag.StringVar(&FlagOptions.Menu, "m", "menu", "菜单")
    flag.StringVar(&FlagOptions.Type, "t", "", "类型")
    flag.BoolVar(&FlagOptions.DB, "db", false, "数据库迁移")
    flag.Parse()
}
```
<mcfile name="flags\enter.go" path="e:\myblog\Go项目学习\rbacAdmin\flags\enter.go"></mcfile>

## 6. 项目代码运行与加载流程

RbacAdmin项目的启动和运行遵循明确的流程，从入口文件到各个组件的初始化都有清晰的调用链。

### 6.1 应用程序启动流程
应用程序的入口点位于`main.go`文件中，定义了应用启动时的主要流程：

```go
func main() {
    core.InitLogger("logs")           // 初始化日志
    global.Config = core.ReadConfig() // 读取配置
    global.DB = core.InitGorm()       // 初始化数据库
    global.Casbin = core.InitCasbin() // 初始化casbin
    global.Redis = core.InitRedis()   // 初始化redis

    // 启动邮件验证码清理定时器
    captcha.EmailStore.StartCleanupTimer()

    flags.Run() // 运行应用
    route.Run() // 运行路由
}
```
<mcfile name="main.go" path="e:\myblog\Go项目学习\rbacAdmin\main.go"></mcfile>

启动流程主要包括以下几个步骤：

1. **初始化日志系统**：设置日志输出目录和格式
2. **读取配置文件**：解析YAML配置文件，加载系统、数据库和Redis配置
3. **初始化数据库连接**：根据配置创建GORM数据库连接
4. **初始化Casbin权限控制**：加载权限策略
5. **初始化Redis连接**：创建Redis客户端连接
6. **启动邮件验证码清理定时器**：定期清理过期的邮件验证码
7. **执行命令行操作**：根据命令行参数执行相应操作，如数据库迁移
8. **启动HTTP服务**：配置路由并启动Web服务器

### 6.2 配置文件读取流程
配置文件的读取逻辑位于`core/read_config.go`中：

```go
func ReadConfig() *config.Config {
    byteData, err := os.ReadFile(flags.FlagOptions.File)
    if err != nil {
        logrus.Fatalf("❌ 配置文件读取失败: %v", err.Error())
        return nil
    }
    var c *config.Config
    err = yaml.Unmarshal(byteData, &c)
    if err != nil {
        logrus.Fatalf("❌ 配置文件格式解析失败: %v", err.Error())
        return nil
    }
    logrus.Infof("✅ 配置文件加载成功: %s", flags.FlagOptions.File)
    return c
}
```
<mcfile name="core\read_config.go" path="e:\myblog\Go项目学习\rbacAdmin\core\read_config.go"></mcfile>

配置读取流程：
1. 根据命令行参数指定的配置文件路径读取文件内容
2. 使用yaml.v3库解析配置文件内容到Config结构体
3. 将解析后的配置对象返回给调用者
4. 如果读取或解析失败，记录错误日志并终止程序

### 6.3 数据库初始化流程
数据库初始化逻辑位于`core/db.go`中，支持多种数据库类型：

```go
func InitGorm() (database *gorm.DB) {
    var db = global.Config.DB
    var dialector gorm.Dialector

    // 根据数据库类型构建连接字符串
    switch db.MODE {
    case "mysql":
        dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
            db.USER, db.PASSWORD, db.HOST, db.PORT, db.DbNAME)
        dialector = mysql.Open(dsn)
    case "pgsql", "postgres", "postgresql":
        dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
            db.HOST, db.USER, db.PASSWORD, db.DbNAME, db.PORT)
        dialector = postgres.Open(dsn)
    case "sqlite":
        dialector = sqlite.Open(db.HOST)
    case "":
        logrus.Fatalf("未配置数据库类型")
        return nil

    default:
        logrus.Fatalf("不支持的数据库类型: %s", db.MODE)
    }

    // 打开数据库连接
    database, err := gorm.Open(dialector, &gorm.Config{
        DisableForeignKeyConstraintWhenMigrating: true, // 不生成外键约束
    })
    if err != nil {
        logrus.Fatalf("❌ 数据库连接失败: %v", err)
        return nil
    }

    // 配置连接池
    sqlDB, err := database.DB()
    if err != nil {
        logrus.Fatalf("获取数据库连接池失败: %v", err)
        return nil
    }

    // 测试数据库连接
    err = sqlDB.Ping()
    if err != nil {
        logrus.Fatalf("❌ 数据库连接测试失败: %v", err)
        return nil
    }

    // 设置连接池参数
    sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
    sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
    sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大生命周期

    logrus.Infof("✅ 数据库连接成功: %s", db.MODE)
    return database
}
```
<mcfile name="core\db.go" path="e:\myblog\Go项目学习\rbacAdmin\core\db.go"></mcfile>

数据库初始化流程：
1. 获取全局配置中的数据库配置信息
2. 根据数据库类型（MySQL、PostgreSQL或SQLite）构建相应的连接字符串
3. 使用GORM库连接数据库并进行配置
4. 配置数据库连接池参数（最大连接数、空闲连接数等）
5. 返回数据库连接对象

### 6.4 数据库迁移流程
数据库迁移功能通过命令行参数`-db`触发，实现在`flags/db.go`中：

```go
func Run() {
    if FlagOptions.DB {
        AutoMigrate()
        os.Exit(0)
    }
    switch FlagOptions.Menu {
    case "user":
        var user User
        switch FlagOptions.Type {
        case "create":
            user.Create()
            os.Exit(0)
        }
    }
}

func AutoMigrate() {
    // 这里可以添加数据库自动迁移的逻辑
    err := global.DB.AutoMigrate(&models.UserModel{}, &models.RoleModel{}, &models.MenuModel{}, &models.APIModel{}, &models.RoleMenuModel{}, &gormadapter.CasbinRule{})
    if err != nil {
        logrus.Fatalf("数据库自动迁移失败: %v", err)
    }
    logrus.Info("数据库自动迁移成功")
}
```
<mcfile name="flags\enter.go" path="e:\myblog\Go项目学习\rbacAdmin\flags\enter.go"></mcfile>
<mcfile name="flags\db.go" path="e:\myblog\Go项目学习\rbacAdmin\flags\db.go"></mcfile>

数据库迁移流程：
1. 检查命令行参数中是否有`-db`标志
2. 如果有，调用`AutoMigrate`函数执行数据库迁移
3. 在`AutoMigrate`函数中，使用GORM的`AutoMigrate`方法根据数据模型自动创建或更新数据库表
4. 迁移完成后，程序自动退出

### 6.5 Casbin权限控制初始化
Casbin是一个强大的权限控制框架，用于实现RBAC权限模型。初始化逻辑位于`core/casbin.go`中。

### 6.6 全局变量管理
全局变量定义在`global/global.go`中，用于在不同模块间共享配置、数据库连接、Redis客户端和Casbin权限控制等资源。

## 7 数据模型详解
项目的核心数据模型定义在`models/enter.go`文件中，采用GORM进行ORM映射。以下是主要数据模型的详细说明：

### 7.1 基础模型
所有数据模型都继承自基础模型，包含ID和时间戳字段：

```go
type Model struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    CreatedAt time.Time `json:"created_at" comment:"创建时间"`
    UpdatedAt time.Time `json:"updated_at" comment:"更新时间"`
}
```
<mcfile name="models\enter.go" path="e:\myblog\Go项目学习\rbacAdmin\models\enter.go"></mcfile>

### 7.2 用户模型(UserModel)
用户模型存储系统用户信息，包含基本信息和权限相关字段：

```go
type UserModel struct {
    Model
    Username string      `gorm:"size:64;unique" json:"username" comment:"用户名"`
    Nickname string      `gorm:"size:64;" json:"nickname" comment:"昵称"`
    Avatar   string      `gorm:"size:256;" json:"avatar" comment:"头像"`
    Email    string      `gorm:"size:128;" json:"email" comment:"邮箱"`
    Password string      `gorm:"size:64" json:"password" comment:"密码"`
    IsAdmin  bool        `gorm:"default:false" json:"is_admin" comment:"是否管理员"`
    RoleList []RoleModel `gorm:"many2many:user_role_models; joinForeignKey:UserID; joinReferences:RoleID;" json:"roleList" comment:"角色列表"`
}
```
<mcfile name="models\enter.go" path="e:\myblog\Go项目学习\rbacAdmin\models\enter.go"></mcfile>

### 7.3 角色模型(RoleModel)
角色模型定义系统角色，与用户和菜单有一对多关系：

```go
type RoleModel struct {
    Model
    Title       string      `gorm:"size:16;unique" json:"title" comment:"角色名称"`
    Description string      `gorm:"size:256;" json:"description" comment:"角色描述"`
    UserList    []UserModel `gorm:"many2many:user_role_models; joinForeignKey:RoleID; joinReferences:UserID;" json:"userList" comment:"用户列表"`
    MenuList    []MenuModel `gorm:"many2many:role_menu_models; joinForeignKey:RoleID; joinReferences:MenuID;" json:"menuList" comment:"菜单列表"`
}

// 用户角色关联模型
type UserRoleModel struct {
    Model
    UserID    uint      `gorm:"index" json:"user_id" comment:"用户ID"`
    UserModel UserModel `gorm:"foreignKey:UserID;references:ID" json:"user" comment:"用户"`
    RoleID    uint      `gorm:"index" json:"role_id" comment:"角色ID"`
    RoleModel RoleModel `gorm:"foreignKey:RoleID;references:ID" json:"role" comment:"角色"`
}

// 角色菜单关联模型
type RoleMenuModel struct {
    Model
    RoleID    uint      `gorm:"index" json:"role_id" comment:"角色ID"`
    RoleModel RoleModel `gorm:"foreignKey:RoleID;references:ID" json:"role" comment:"角色"`
    MenuID    uint      `gorm:"index" json:"menu_id" comment:"菜单ID"`
    MenuModel MenuModel `gorm:"foreignKey:MenuID;references:ID" json:"menu" comment:"菜单"`
}
```
<mcfile name="models\enter.go" path="e:\myblog\Go项目学习\rbacAdmin\models\enter.go"></mcfile>

### 7.4 菜单模型(MenuModel)
菜单模型定义系统菜单结构，支持多级菜单：

```go
type Meta struct {
    Icon  string `gorm:"size:128" json:"icon" comment:"图标"`
    Title string `gorm:"size:64" json:"title" comment:"标题"`
}

type MenuModel struct {
    Model
    Name            string `gorm:"size:64;unique" json:"name" comment:"菜单名称"`
    Path            string `gorm:"size:128" json:"path" comment:"菜单路径"`
    Component       string `gorm:"size:128" json:"component" comment:"组件路径"`
    Meta            `gorm:"embedded" json:"meta" comment:"元数据"`
    ParentMenuID    *uint        `gorm:"index" json:"ParentMenuID" comment:"父菜单ID"`
    ParentMenuModel *MenuModel   `gorm:"foreignKey:ParentMenuID;references:ID" json:"-" comment:"父菜单"`
    Children        []*MenuModel `gorm:"foreignKey:ParentMenuID;references:ID" json:"children" comment:"子菜单"`
    Sort            int          `json:"sort" comment:"排序"`
}
```
<mcfile name="models\enter.go" path="e:\myblog\Go项目学习\rbacAdmin\models\enter.go"></mcfile>

### 7.5 API模型(APIModel)
API模型定义系统API权限：

```go
type APIModel struct {
    Model
    Name        string `gorm:"size:64;unique" json:"name" comment:"API名称"`
    Path        string `gorm:"size:128" json:"path" comment:"API路径"`
    Method      string `gorm:"size:16" json:"method" comment:"请求方法"`
    Group       string `gorm:"size:64" json:"group" comment:"API分组"`
    Description string `gorm:"size:256" json:"description" comment:"API描述"`
}
```
<mcfile name="models\enter.go" path="e:\myblog\Go项目学习\rbacAdmin\models\enter.go"></mcfile>

## 8. 常见问题与解决方案

### 8.1 数据库连接失败
**问题现象**：应用启动时提示"数据库连接失败"错误

**可能原因**：
- 数据库服务未启动
- 数据库连接参数配置错误
- 数据库用户权限不足

**解决方案**：
- 确认数据库服务是否正常运行
- 检查配置文件中的数据库连接参数是否正确
- 验证数据库用户是否有足够的权限访问指定数据库

### 8.2 Redis连接失败
**问题现象**：应用启动时提示"Redis连接失败"错误

**可能原因**：
- Redis服务未启动
- Redis连接参数配置错误
- Redis密码错误

**解决方案**：
- 确认Redis服务是否正常运行
- 检查配置文件中的Redis连接参数是否正确
- 验证Redis密码是否正确

### 8.3 邮箱发送失败
**问题现象**：发送邮件验证码时提示失败

**可能原因**：
- 邮箱配置不正确
- 邮箱服务未开启SMTP服务
- 邮箱账号或密码错误

**解决方案**：
- 检查配置文件中的邮箱配置是否正确
- 确认邮箱服务已开启SMTP服务
- 验证邮箱账号和密码是否正确

### 8.4 权限验证失败
**问题现象**：用户登录后无法访问某些功能

**可能原因**：
- 用户角色权限配置不正确
- Casbin权限策略未正确加载

**解决方案**：
- 检查用户所属角色的权限配置
- 确认Casbin权限策略是否正确加载

## 9. 更新日志

### v1.0 (2024-07-05)
- 初始版本发布
- 支持完整的RBAC权限管理功能
- 支持MySQL、SQLite、PostgreSQL数据库
- 集成Redis缓存
- 实现验证码和邮箱验证功能
- 提供RESTful API接口