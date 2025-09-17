# 🚀 RBAC Admin System - 仓库提交指南

## 📋 提交前检查清单

在提交代码到Git仓库之前，请务必完成以下检查：

### ✅ 安全检查
- [ ] 已检查 `.gitignore` 文件是否包含所有敏感文件
- [ ] 已确认没有提交 `settings.yaml`（包含数据库密码）
- [ ] 已确认没有提交 `.env` 文件
- [ ] 已确认没有提交日志文件 (`logs/` 目录)
- [ ] 已确认没有提交可执行文件 (`*.exe`)

### ✅ 代码质量
- [ ] 代码已通过编译测试：`go build`
- [ ] 没有明显的语法错误
- [ ] 代码格式整洁（建议格式化）

### ✅ 文档完整性
- [ ] `README.md` 已更新（如功能有变更）
- [ ] `SECURITY.md` 已包含必要的安全说明
- [ ] 配置模板文件 `settings.yaml.example` 已更新

## 📝 推荐的提交信息格式

### 功能添加
```
feat: 添加用户角色管理功能

- 实现用户角色的增删改查
- 添加角色权限关联
- 优化角色查询性能
```

### 问题修复
```
fix: 修复数据库连接池泄漏问题

- 修复连接未正确释放的问题
- 添加连接超时处理
- 优化连接池配置
```

### 文档更新
```
docs: 更新README和安全配置说明

- 添加部署步骤说明
- 更新安全配置指南
- 完善API文档
```

### 安全修复
```
security: 修复JWT令牌验证漏洞

- 加强令牌签名验证
- 添加令牌过期检查
- 更新相关依赖包
```

## 🚫 禁止提交的内容

### ❌ 绝对禁止提交
- `settings.yaml` - 包含数据库密码
- `.env` - 环境变量文件
- `*.key`, `*.pem` - 密钥和证书文件
- `logs/` - 日志文件目录
- `*.exe` - 可执行文件

### ⚠️ 谨慎提交
- 大型二进制文件
- 临时测试文件
- 个人IDE配置文件

## 🔍 提交前验证命令

```bash
# 1. 检查待提交文件
git status

# 2. 查看具体变更
git diff

# 3. 测试编译
go build

# 4. 检查敏感文件是否被忽略
git check-ignore settings.yaml
```

## 🎯 首次提交建议

### 第一步：创建安全的基础提交
```bash
# 添加基础文件
git add README.md SECURITY.md settings.yaml.example .gitignore
git commit -m "initial: 添加项目基础文档和配置模板"
```

### 第二步：添加代码文件
```bash
# 添加源代码（不包含敏感文件）
git add api/ config/ core/ flags/ global/ middleware/ models/ routes/ service/ utils/ mian.go go.mod go.sum
git commit -m "feat: 添加RBAC权限管理系统核心代码"
```

### 第三步：添加开发配置
```bash
# 添加开发环境配置（已清理敏感信息）
git add settings_dev.yaml
git commit -m "config: 添加开发环境配置"
```

## 🔔 重要提醒

1. **定期推送**：建议定期将本地提交推送到远程仓库
2. **分支管理**：新功能开发请创建feature分支
3. **代码审查**：重要变更建议进行代码审查
4. **备份重要分支**：保护main/master分支

## 📞 遇到问题？

如果在提交过程中遇到问题：
1. 检查 `.gitignore` 配置是否正确
2. 确认敏感文件没有被意外添加
3. 查看Git状态：`git status`
4. 必要时可以回退：`git reset HEAD <file>`

---

**记住：安全第一，提交第二！保护好敏感信息比提交代码更重要。**