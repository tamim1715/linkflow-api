package email

type Sender interface {
	SendMagicLink(email string, link string) error
}
