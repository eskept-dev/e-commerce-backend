package services

import (
	"bytes"
	"eskept/internal/app/context"
	"eskept/internal/repositories"
	jwt "eskept/internal/utils/auth"
	"fmt"
	"html/template"
	"net/smtp"
)

type EmailService struct {
	repo   *repositories.UserRepository
	appCtx *context.AppContext
}

func NewEmailService(
	repo *repositories.UserRepository,
	appCtx *context.AppContext,
) *EmailService {
	return &EmailService{repo: repo, appCtx: appCtx}
}

func (s *EmailService) GenerateActivationLink(email, role string) (string, error) {
	activationToken, err := jwt.GenerateActivationToken(email, role, s.appCtx)
	if err != nil {
		return "", err
	}

	activationLink := s.appCtx.Cfg.App.ActivationURL + "?activationToken=" + activationToken
	return activationLink, nil
}

func (s *EmailService) SendActivationLink(email, role string) error {
	activationLink, err := s.GenerateActivationLink(email, role)
	if err != nil {
		return err
	}

	body, err := loadEmailTemplate(
		s.appCtx.Cfg.Template.EmailActivation,
		map[string]string{
			"ActivationLink": activationLink,
		},
	)
	if err != nil {
		return err
	}

	err = s.SendEmail(email, "Activation Link", body)
	if err != nil {
		return err
	}
	return nil
}

func (s *EmailService) SendEmail(toEmail, subject, body string) error {
	fromEmail := s.appCtx.Cfg.SMTP.Email
	password := s.appCtx.Cfg.SMTP.Password
	SMTPServer := s.appCtx.Cfg.SMTP.Host
	addr := fmt.Sprintf("%s:%d", SMTPServer, 587)

	auth := smtp.PlainAuth(
		"",
		fromEmail,
		password,
		SMTPServer,
	)
	toEmails := []string{toEmail}
	var emailContent bytes.Buffer

	// Write email headers
	emailContent.WriteString("From: " + fromEmail + "\r\n")
	emailContent.WriteString("To: " + toEmail + "\r\n")
	emailContent.WriteString("Subject: " + subject + "\r\n")
	emailContent.WriteString("MIME-Version: 1.0\r\n")
	emailContent.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n")
	emailContent.WriteString(body)

	err := smtp.SendMail(
		addr,
		auth,
		fromEmail,
		toEmails,
		emailContent.Bytes(),
	)
	if err != nil {
		return err
	}

	return nil
}

func loadEmailTemplate(templatePath string, data interface{}) (string, error) {
	// Parse the HTML file
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	// Buffer to hold the parsed template
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}
