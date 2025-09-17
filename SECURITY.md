# 🔐 RBAC Admin System - 安全配置指南

## 🚨 重要安全提醒

在部署和使用本系统时，请务必遵循以下安全最佳实践，保护您的系统和数据安全。

## 🛡️ 敏感信息保护

### 1. 配置文件安全

#### ⚠️ 绝不要提交到Git的文件：
- `settings.yaml` - 包含数据库密码
- `settings_prod.yaml` - 生产环境配置
- `settings_test.yaml` - 测试环境配置
- `.env` - 环境变量文件
- `*.key`, `*.pem` - 密钥和证书文件
- `logs/` - 日志文件目录

#### ✅ 安全的配置模板：
- `settings.yaml.example` - 配置模板文件
- `.env.example` - 环境变量模板

### 2. 数据库安全配置

#### 数据库连接配置：
```yaml
db:
    mode: mysql
    host: localhost      # 使用localhost而非公网IP
    port: 3306
    user: rbac_user      # 使用专用用户，而非root
    password: YOUR_STRONG_PASSWORD_HERE  # 强密码
    db_name: rbacAdmin
```

#### 数据库安全建议：
- 🔑 **使用强密码**：至少12位，包含大小写字母、数字和特殊字符
- 👤 **创建专用用户**：不要使用root用户连接应用
- 🔒 **限制访问IP**：配置数据库只允许特定IP访问
- 🛡️ **启用SSL连接**：生产环境启用数据库SSL连接
- 📋 **定期备份**：设置自动备份策略

### 3. Redis安全配置

```yaml
redis:
    addr: localhost:6379
    password: YOUR_REDIS_PASSWORD  # 设置强密码
    db: 0                          # 使用专用数据库
```

#### Redis安全建议：
- 🔑 **设置密码**：为Redis设置强密码
- 🔒 **绑定本地IP**：只允许本地访问
- 🛡️ **禁用危险命令**：禁用FLUSHDB、FLUSHALL等命令
- 📊 **监控访问**：启用Redis访问日志

## 🔧 环境配置指南

### 开发环境
```yaml
# settings_dev.yaml
system:
    ip: 127.0.0.1
    port: 8080

db:
    mode: sqlite  # 开发环境使用SQLite
    db_name: rbac_dev.db
```

### 测试环境
```yaml
# settings_test.yaml
system:
    ip: 127.0.0.1
    port: 8081

db:
    mode: mysql
    host: localhost
    user: test_user
    password: TEST_PASSWORD
    db_name: rbac_test
```

### 生产环境
```yaml
# settings_prod.yaml
system:
    ip: 0.0.0.0     # 监听所有网卡
    port: 8080

db:
    mode: mysql
    host: localhost  # 或使用内网IP
    user: rbac_prod
    password: PRODUCTION_STRONG_PASSWORD
    db_name: rbac_production
```

## 🚀 部署安全建议

### 1. 服务器安全
- 🔄 **定期更新**：保持操作系统和软件更新
- 🔥 **配置防火墙**：只开放必要的端口
- 👥 **使用非root用户**：运行应用的用户权限最小化
- 🔍 **监控日志**：定期检查系统和应用日志

### 2. 网络安全
- 🔒 **使用HTTPS**：生产环境必须启用HTTPS
- 🛡️ **配置CORS**：合理配置跨域请求
- 📍 **IP白名单**：限制管理后台访问IP
- 🕐 **会话超时**：设置合理的会话超时时间

### 3. 应用安全
- 🔑 **JWT密钥**：使用强密钥，定期更换
- 📝 **输入验证**：严格验证所有用户输入
- 🛡️ **SQL注入防护**：使用参数化查询
- 📊 **错误处理**：不要暴露敏感错误信息

## 📝 安全检查清单

在部署前，请确认：

- [ ] 已修改所有默认密码
- [ ] 已配置.gitignore忽略敏感文件
- [ ] 已创建专用数据库用户
- [ ] 已设置数据库访问IP限制
- [ ] 已配置Redis密码
- [ ] 已启用HTTPS（生产环境）
- [ ] 已配置防火墙规则
- [ ] 已设置日志监控
- [ ] 已创建备份策略
- [ ] 已进行安全测试

## 🚨 紧急情况处理

### 密码泄露
1. 立即更改所有相关密码
2. 检查访问日志
3. 通知相关人员
4. 更新配置并重新部署

### 系统入侵
1. 立即断开网络连接
2. 保存日志证据
3. 检查系统完整性
4. 恢复安全备份

## 📞 安全支持

如发现安全漏洞或有安全建议：
1. 不要公开披露漏洞详情
2. 通过私密渠道报告
3. 等待修复后再公开

---

**记住：安全是一个持续的过程，需要定期审查和更新！**