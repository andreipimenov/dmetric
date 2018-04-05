package mail

import gomail "gopkg.in/gomail.v2"

type Mail struct {
	SMTPHost     string
	SMTPPort     int
	SMTPLogin    string
	SMTPPassword string
}

func NewMail(SMTPHost string, SMTPPort int, SMTPLogin string, SMTPPassword string) *Mail {
	return &Mail{
		SMTPHost:     SMTPHost,
		SMTPPort:     SMTPPort,
		SMTPLogin:    SMTPLogin,
		SMTPPassword: SMTPPassword,
	}
}

func (m *Mail) Send(to string, subject string, body string) error {
	l := gomail.NewMessage()
	l.SetHeader("From", m.SMTPLogin)
	l.SetHeader("To", to)
	l.SetHeader("Subject", subject)
	l.SetBody("text/html", body)
	d := gomail.NewDialer(m.SMTPHost, m.SMTPPort, m.SMTPLogin, m.SMTPPassword)
	return d.DialAndSend(l)
}
