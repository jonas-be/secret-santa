package email

type Content struct {
	Subject string
	Mime    string
	Body    string
}

func (content Content) toByteArray() []byte {
	return []byte(content.Subject + content.Mime + content.Body)
}
