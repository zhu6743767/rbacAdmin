package captcha

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

var CaptchaStore = base64Captcha.DefaultMemStore

// EmailCodeStore 邮件验证码存储
type EmailCodeStore struct {
	codes map[string]EmailCodeInfo
	mu    sync.RWMutex
}

type EmailCodeInfo struct {
	Code      string
	Email     string
	ExpiredAt time.Time
}

var EmailStore = &EmailCodeStore{
	codes: make(map[string]EmailCodeInfo),
}

// SetEmailCode 设置邮件验证码
func (s *EmailCodeStore) SetEmailCode(email, code string, duration time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.codes[email] = EmailCodeInfo{
		Code:      code,
		Email:     email,
		ExpiredAt: time.Now().Add(duration),
	}
}

// GetEmailCode 获取邮件验证码
func (s *EmailCodeStore) GetEmailCode(email string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	info, exists := s.codes[email]
	if !exists || time.Now().After(info.ExpiredAt) {
		return "", false
	}
	return info.Code, true
}

// VerifyEmailCode 验证邮件验证码
func (s *EmailCodeStore) VerifyEmailCode(email, code string) bool {
	storedCode, exists := s.GetEmailCode(email)
	if !exists {
		return false
	}
	return storedCode == code
}

// CleanExpired 清理过期验证码（可以定期调用）
func (s *EmailCodeStore) CleanExpired() {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	count := 0
	for email, info := range s.codes {
		if now.After(info.ExpiredAt) {
			delete(s.codes, email)
			count++
		}
	}
	if count > 0 {
		logrus.Infof("清理了 %d 个过期的邮件验证码", count)
	}
}

// StartCleanupTimer 启动定期清理定时器
func (s *EmailCodeStore) StartCleanupTimer() {
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			s.CleanExpired()
		}
	}()
}

// GenerateAndStoreEmailCode 生成并存储邮件验证码（统一接口）
func (s *EmailCodeStore) GenerateAndStoreEmailCode(email string, duration time.Duration) string {
	code := generateEmailCode()
	s.SetEmailCode(email, code, duration)
	return code
}

// generateEmailCode 生成6位数字验证码
func generateEmailCode() string {
	code := ""
	for i := 0; i < 6; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(10))
		code += fmt.Sprintf("%d", n.Int64())
	}
	return code
}
