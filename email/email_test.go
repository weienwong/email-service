package email

import (
	"os"
	"testing"
)

func TestSendEmail(t *testing.T) {
	to := []string{
		"weienwong.93@gmail.com",
	}
	from := "weienwong.93@gmail.com"

	e := &EmailClient{MailgunAPIKey: os.Getenv("MAILGUN_API_KEY")}

	err := e.SendEmail(to, from, "subject", "body")

	if err != nil {
		t.Fatal(err)
	}
}
