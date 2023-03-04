package email

type Content struct {
	Subject string
	Body    string
}

const Mime = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

func (content Content) toByteArray() []byte {
	return []byte(content.Subject + Mime + content.Body)
}
