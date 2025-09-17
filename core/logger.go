// Package core 核心功能包，提供日志系统初始化和管理功能
package core

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/sirupsen/logrus"
)

// Mylogger 自定义日志格式器结构体
type Mylogger struct {
	LogLevel string `yaml:"log_level"` // YAML配置中的日志级别
}

// MyHook 日志钩子结构体，实现logrus.Hook接口
type MyHook struct {
	logPath  string     // 日志文件存储路径
	file     *os.File   // 普通日志文件句柄
	errFile  *os.File   // 错误日志文件句柄
	fileDate string     // 当前日志文件日期（用于按日期分割）
	mu       sync.Mutex // 互斥锁，保证并发安全
}

// ANSI颜色代码常量
const (
	red    = 31 // 红色 - 用于错误、致命、恐慌级别
	yellow = 33 // 黄色 - 用于警告级别
	blue   = 36 // 蓝色 - 用于信息级别
	grau   = 37 // 灰色 - 用于调试、跟踪级别
)

// Format 实现logrus.Formatter接口，自定义日志输出格式
// 参数:
//   - entry: logrus日志条目，包含所有日志信息
//
// 返回:
//   - []byte: 格式化后的日志内容
//   - error: 格式化过程中的错误
func (Mylogger) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据日志级别设置对应的颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = grau // 调试和跟踪级别使用灰色
	case logrus.InfoLevel:
		levelColor = blue // 信息级别使用蓝色
	case logrus.WarnLevel:
		levelColor = yellow // 警告级别使用黄色
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red // 错误、致命、恐慌级别使用红色
	default:
		levelColor = blue // 默认使用蓝色
	}

	// 获取或创建缓冲区
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// 自定义日期时间格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")

	// 根据是否有调用者信息选择不同的格式
	if entry.HasCaller() {
		// 详细格式：包含函数名、文件名和行号
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(b, "\x1b[%dm[%s] [%s] %s %s: %s\x1b[0m\n",
			levelColor, timestamp, entry.Level, funcVal, fileVal, entry.Message)
	} else {
		// 简化格式：只包含时间、级别和消息
		fmt.Fprintf(b, "\x1b[%dm[%s] [%s]: %s\x1b[0m\n",
			levelColor, timestamp, entry.Level, entry.Message)
	}

	return b.Bytes(), nil
}

// Fire 实现logrus.Hook接口的核心方法
// 当日志事件触发时，此方法会被调用
// 参数:
//   - entry: 包含日志信息的条目
//
// 返回:
//   - error: 处理过程中的错误
func (hook *MyHook) Fire(entry *logrus.Entry) error {
	// 加锁保证并发安全
	hook.mu.Lock()
	defer hook.mu.Unlock()

	// 获取当前日期，用于日志文件按日期分割
	timer := entry.Time.Format("2006-01-02")

	// 将日志条目格式化为字符串
	line, err := entry.String()
	if err != nil {
		return fmt.Errorf("日志格式化失败: %v", err)
	}

	// 检查是否需要创建新的日志文件（跨天时）
	if hook.fileDate != timer {
		if err := hook.rotateFiles(timer); err != nil {
			return err
		}
	}

	// 写入普通日志文件（所有级别）
	if _, err := hook.file.Write([]byte(line)); err != nil {
		return fmt.Errorf("写入info日志文件失败: %v", err)
	}

	// 警告级别及以上额外写入错误日志文件
	if entry.Level >= logrus.WarnLevel {
		if _, err := hook.errFile.Write([]byte(line)); err != nil {
			return fmt.Errorf("写入error日志文件失败: %v", err)
		}
	}

	return nil
}

// rotateFiles 日志文件轮换函数
// 按日期创建新的日志文件，实现日志的按天分割
// 参数:
//   - timer: 当前日期字符串（格式：2006-01-02）
//
// 返回:
//   - error: 轮换过程中的错误
func (hook *MyHook) rotateFiles(timer string) error {
	// 关闭已存在的日志文件
	if hook.file != nil {
		if err := hook.file.Close(); err != nil {
			return fmt.Errorf("关闭旧日志文件失败: %v", err)
		}
	}
	if hook.errFile != nil {
		if err := hook.errFile.Close(); err != nil {
			return fmt.Errorf("关闭旧错误日志文件失败: %v", err)
		}
	}

	// 创建日期目录，如：logs/2025-09-16/
	dirName := fmt.Sprintf("%s/%s", hook.logPath, timer)
	if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
		return fmt.Errorf("创建日志目录失败: %v", err)
	}

	// 构建日志文件路径
	infoFileName := fmt.Sprintf("%s/info.log", dirName) // 普通日志文件
	errFileName := fmt.Sprintf("%s/error.log", dirName) // 错误日志文件

	// 创建或打开普通日志文件
	var err error
	hook.file, err = os.OpenFile(infoFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("创建日志文件失败: %v", err)
	}

	// 创建或打开错误日志文件
	hook.errFile, err = os.OpenFile(errFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("创建错误日志文件失败: %v", err)
	}

	// 更新当前日期
	hook.fileDate = timer
	return nil
}

// Levels 指定此Hook处理的日志级别
// 返回所有日志级别，表示处理所有类型的日志
// 返回值:
//   - []logrus.Level: 包含所有日志级别的切片
func (hook *MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// InitLogger 初始化日志系统
// 配置logrus日志库，包括格式、输出、级别等设置
// 参数:
//   - logPath: 日志文件存储路径
func InitLogger(logPath string) {
	// 设置自定义格式器，处理日志的格式化输出
	formatter := &Mylogger{}
	logrus.SetFormatter(formatter)

	// 启用日志调用者信息（显示调用日志的文件和行号）
	logrus.SetReportCaller(true)

	// 设置日志输出到控制台
	logrus.SetOutput(os.Stdout)

	// 设置日志级别为Debug，记录所有级别的日志
	logrus.SetLevel(logrus.DebugLevel)

	// 创建并添加自定义Hook，处理日志文件写入
	hook := &MyHook{
		logPath: logPath, // 日志文件根目录
	}
	logrus.AddHook(hook)
}
