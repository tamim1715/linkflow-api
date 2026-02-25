package email

import "log"

type MockSender struct{}

func NewMockSender() *MockSender {
	return &MockSender{}
}

func (m *MockSender) SendMagicLink(email string, link string) error {
	log.Println("📩 MAGIC LINK EMAIL TO:", email)
	log.Println("🔗 LOGIN LINK:", link)
	return nil
}
