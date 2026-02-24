package slack

import "log"

type MockClient struct{}

func (m *MockClient) Publish(message string) error {
	log.Println("[MOCK SLACK]", message)
	return nil
}
