package slack

type Client interface {
	Publish(message string) error
}
