package email

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Sender interface {
	SendEmail(to []string, from string, subject string, body string) error
	SendEmailAttachment(to []string, from string, subject string, body string, filename string) error
}

type EmailClient struct {
	MailgunAPIKey string
}

func (e *EmailClient) SendEmail(to []string, from string, subject string, body string) error {
	recipients := strings.Join(to, ",")
	mailgunURL := "https://api.mailgun.net/v3/sandbox3af273bc6e66458f9b9a4015b7d3a28e.mailgun.org/messages"
	// Need to build request body using url.Values
	v := url.Values{}
	v.Add("from", from)
	v.Add("to", recipients)
	v.Add("subject", subject)
	v.Add("text", body)

	client := &http.Client{}
	req, err := http.NewRequest("POST", mailgunURL, strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %v", e.MailgunAPIKey))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Println(string(respBody))
		return fmt.Errorf("error: unable to send email request")
	}

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (e *EmailClient) SendEmailAttachment(to []string, from string, subject string, body string, filename string) error {
	return nil
}
