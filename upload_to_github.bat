@echo off
echo 🚀 开始上传 rbacAdmin 项目到 GitHub...
echo.

REM 进入项目目录
cd /d "e:\myblog\Go项目学习\rbacAdmin"

echo 📁 当前目录: %cd%
echo.

echo 🔍 检查Git状态...
git status
echo.

echo 📝 添加所有文件到暂存区...
git add .
echo.

echo 📊 查看添加的文件...
git status
echo.

echo 💾 提交更改...
git commit -m "Update rbacAdmin project files - Add complete project structure with config, core modules, and documentation"
echo.

echo 🔄 推送到GitHub...
git push -u origin main

echo.
echo ✅ 上传完成！
echo 🌐 请访问: https://github.com/zhu6743767/rbacAdmin 查看结果
echo.
pause