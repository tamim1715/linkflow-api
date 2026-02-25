package slack

import "log"

type MockClient struct{}

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (m *MockClient) Publish(message string) error {
	log.Println("[MOCK SLACK]", message)
	return nil
}
