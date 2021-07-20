package service

import (
	"fmt"
	"io/ioutil"
	"net/mail"
	"os"
	"strconv"

	"github.com/matcornic/hermes/v2"
	"github.com/spf13/viper"
	"gitlab.com/renodesper/gokit-microservices/util/errors"
	"gitlab.com/renodesper/gokit-microservices/util/logger"
	"gopkg.in/gomail.v2"
)

type (
	EmailService interface {
		SendMail(userID, userEmail, subject string, email hermes.Email, emailType string) error
		Welcome(userName, token string) hermes.Email
		ResetPassword(userName, token string) hermes.Email
		Maintenance(userName string) hermes.Email
		GenerateEmail(email hermes.Email, emailType string, userID string) error
		Send(options SendOptions, htmlBody string, txtBody string) error
	}

	EmailSvc struct {
		Log logger.Logger
		SmtpConfig
		Hermes hermes.Hermes
	}

	SmtpConfig struct {
		Server         string
		Port           int
		SenderEmail    string
		SenderIdentity string
		SMTPUser       string
		SMTPPassword   string
	}

	SendOptions struct {
		To      string
		Subject string
	}

	UserData struct {
		Name string
	}
)

var HTML_PATH_FORMAT = "static/%v_%v.html"
var TXT_PATH_FORMAT = "static/%v_%v.txt"

// NewEmailSvc creates email service
func NewEmailSvc(log logger.Logger) *EmailSvc {
	port, _ := strconv.Atoi(viper.GetString("mail.smtp_port"))
	smtpConfig := SmtpConfig{
		Port:           port,
		Server:         viper.GetString("mail.smtp_server"),
		SenderEmail:    viper.GetString("mail.sender_email"),
		SenderIdentity: viper.GetString("mail.sender_identity"),
		SMTPUser:       viper.GetString("mail.smtp_user"),
		SMTPPassword:   viper.GetString("mail.smtp_password"),
	}
	h := hermes.Hermes{
		Theme: new(hermes.Flat),
		Product: hermes.Product{
			Name: viper.GetString("app.name"),
			Link: viper.GetString("app.url"),
			Logo: viper.GetString("app.logo"),
		},
	}

	return &EmailSvc{
		Log:        log,
		SmtpConfig: smtpConfig,
		Hermes:     h,
	}
}

func (e *EmailSvc) SendMail(userID, userEmail, subject string, email hermes.Email, emailType string) error {
	options := SendOptions{
		To:      userEmail,
		Subject: fmt.Sprintf(subject, e.Hermes.Product.Name),
	}

	err := e.GenerateEmail(email, emailType, userID)
	if err != nil {
		return errors.FailedGenerateEmail.AppendError(err)
	}

	htmlBytes, err := ioutil.ReadFile(fmt.Sprintf(HTML_PATH_FORMAT, emailType, userID))
	if err != nil {
		return errors.FailedReadFile.AppendError(err)
	}

	txtBytes, err := ioutil.ReadFile(fmt.Sprintf(TXT_PATH_FORMAT, emailType, userID))
	if err != nil {
		return errors.FailedReadFile.AppendError(err)
	}

	err = e.Send(options, string(htmlBytes), string(txtBytes))
	if err != nil {
		return errors.FailedSendEmail.AppendError(err)
	}

	return nil
}

func (e *EmailSvc) GenerateEmail(email hermes.Email, emailType string, userID string) error {
	// Generate the HTML template and save it
	res, err := e.Hermes.GenerateHTML(email)
	if err != nil {
		return errors.FailedGenerateEmailBody.AppendError(err)
	}
	err = os.MkdirAll(e.Hermes.Theme.Name(), 0744)
	if err != nil {
		return errors.FailedCreateEmailDirectory.AppendError(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf(HTML_PATH_FORMAT, emailType, userID), []byte(res), 0644)
	if err != nil {
		return errors.FailedWriteEmail.AppendError(err)
	}

	// Generate the plaintext template and save it
	res, err = e.Hermes.GeneratePlainText(email)
	if err != nil {
		return errors.FailedGeneratePlainText.AppendError(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf(TXT_PATH_FORMAT, emailType, userID), []byte(res), 0644)
	if err != nil {
		return errors.FailedWritePlainText.AppendError(err)
	}

	return nil
}

func (e *EmailSvc) Send(options SendOptions, htmlBody string, txtBody string) error {
	if e.SmtpConfig.Server == "" {
		return errors.FailedMissingSMTPServerConfig
	}

	if e.SmtpConfig.Port == 0 {
		return errors.FailedMissingSMTPPortConfig
	}

	if e.SmtpConfig.SMTPUser == "" {
		return errors.FailedMissingSMTPUser
	}

	if e.SmtpConfig.SenderIdentity == "" {
		return errors.FailedMissingSMTPSenderIdentity
	}

	if e.SmtpConfig.SenderEmail == "" {
		return errors.FailedMissingSMTPSenderEmail
	}

	if options.To == "" {
		return errors.FailedMissingSMTPReceiverEmail
	}

	from := mail.Address{
		Name:    e.SmtpConfig.SenderIdentity,
		Address: e.SmtpConfig.SenderEmail,
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from.String())
	m.SetHeader("To", options.To)
	m.SetHeader("Subject", options.Subject)

	m.SetBody("text/plain", txtBody)
	m.AddAlternative("text/html", htmlBody)

	d := gomail.NewDialer(e.SmtpConfig.Server, e.SmtpConfig.Port, e.SmtpConfig.SMTPUser, e.SmtpConfig.SMTPPassword)

	return d.DialAndSend(m)
}

func (e *EmailSvc) Welcome(userName, token string) hermes.Email {
	return hermes.Email{
		Body: hermes.Body{
			Name: userName,
			Intros: []string{
				fmt.Sprintf("Welcome to %s! We're very excited to have you on board.", e.Hermes.Product.Name),
			},
			Actions: []hermes.Action{
				{
					Instructions: fmt.Sprintf("To get started with %s, please confirm your account by clicking the button below:", e.Hermes.Product.Name),
					Button: hermes.Button{
						Text: "Confirm your account",
						Link: fmt.Sprintf("%s/confirm/%s", e.Hermes.Product.Link, token),
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help you.",
			},
			Signature: "Best regards",
		},
	}
}

func (e *EmailSvc) ResetPassword(userName, token string) hermes.Email {
	return hermes.Email{
		Body: hermes.Body{
			Name: userName,
			Intros: []string{
				fmt.Sprintf("You have received this email because a password reset request for %s account was received.", e.Hermes.Product.Name),
			},
			Actions: []hermes.Action{
				{
					Instructions: "Click the button below to reset your password:",
					Button: hermes.Button{
						Text: "Reset your password",
						Link: fmt.Sprintf("%s/reset-password/%s", e.Hermes.Product.Link, token),
					},
				},
			},
			Outros: []string{
				"If you did not request a password reset, no further action is required on your part.",
			},
			Signature: "Best regards",
		},
	}
}

// TODO: Unfinished
func (e *EmailSvc) Maintenance(userName string) hermes.Email {
	return hermes.Email{
		Body: hermes.Body{
			Name: userName,
			FreeMarkdown: `
> _Hermes_ service will shutdown the **1st August 2017** for maintenance operations.
Services will be unavailable based on the following schedule:
| Services | Downtime |
| :------:| :-----------: |
| Service A | 2AM to 3AM |
| Service B | 4AM to 5AM |
| Service C | 5AM to 6AM |
Feel free to contact us for any question regarding this matter at [support@hermes-example.com](mailto:support@hermes-example.com) or in our [Gitter](https://gitter.im/)
`,
		},
	}
}
