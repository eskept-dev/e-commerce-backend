package services

import (
	"bytes"
	"eskept/internal/app/context"
	"eskept/internal/repositories"
	jwt "eskept/internal/utils/auth"
	"fmt"
	"html/template"
	"log"
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

func (s *EmailService) GenerateAuthorizationLink(
	email, role, rootURL string,
	expirationTime int,
) (string, error) {
	authorizationToken, err := jwt.GenerateToken(email, role, expirationTime, s.appCtx)
	if err != nil {
		return "", err
	}

	authorizationLink := rootURL + "?token=" + authorizationToken
	return authorizationLink, nil
}

func (s *EmailService) SendActivationEmail(email, role string) error {
	activationLink, err := s.GenerateAuthorizationLink(
		email,
		role,
		s.appCtx.Cfg.App.ActivationURL,
		s.appCtx.Cfg.JWT.ActivationTokenExpirationTime,
	)
	if err != nil {
		return err
	}

	log.Println("------------------- Send activation email -------------------")
	log.Println("Email:", email)
	log.Println("Role:", role)
	log.Println("Activation link:", activationLink)
	log.Println("------------------------------------------------------------")

	body, err := loadEmailTemplate(
		s.appCtx.Cfg.Template.EmailActivation,
		map[string]string{
			"ActivationLink": activationLink,
		},
	)
	if err != nil {
		return err
	}

	err = s.SendEmail(email, "Eskept Account Activation", body)
	if err != nil {
		return err
	}
	return nil
}

func (s *EmailService) SendAuthenticationEmail(email, role string) error {
	authenticationLink, err := s.GenerateAuthorizationLink(
		email,
		role,
		s.appCtx.Cfg.App.AuthenticationURL,
		s.appCtx.Cfg.JWT.AuthenticationTokenExpirationTime,
	)
	if err != nil {
		return err
	}
	log.Println("------------------- Send authentication email -------------------")
	log.Println("Email:", email)
	log.Println("Role:", role)
	log.Println("Authentication link:", authenticationLink)
	log.Println("------------------------------------------------------------")

	body, err := loadEmailTemplate(
		s.appCtx.Cfg.Template.EmailAuthentication,
		map[string]string{
			"AuthenticationLink": authenticationLink,
		},
	)
	if err != nil {
		return err
	}

	err = s.SendEmail(email, "Eskept Authentication", body)
	if err != nil {
		return err
	}
	return nil
}

func (s *EmailService) SendVerificationEmail(email string) error {
	verificationLink, err := s.GenerateAuthorizationLink(
		email,
		"",
		s.appCtx.Cfg.App.RegistrationURL,
		s.appCtx.Cfg.JWT.RegistrationTokenExpirationTime,
	)
	if err != nil {
		return err
	}

	// Add data to the verification link
	verificationLink = fmt.Sprintf("%s&email=%s", verificationLink, email)

	log.Println("------------------- Send verification email -------------------")
	log.Println("Email:", email)
	log.Println("Verification link:", verificationLink)
	log.Println("------------------------------------------------------------")

	body, err := loadEmailTemplate(
		s.appCtx.Cfg.Template.EmailRegistration,
		map[string]string{
			"VerificationLink": verificationLink,
		},
	)
	if err != nil {
		return err
	}

	err = s.SendEmail(email, "Eskept Registration", body)
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
