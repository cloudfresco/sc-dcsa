package test

import (
	"github.com/cloudfresco/sc-dcsa/internal/common"
	"go.uber.org/zap"
)

// MailerServiceTest Pointer to mailer
type MailerServiceTest struct{}

// NewMailerServiceTest get connection to mailer and create a MailerService struct
func NewMailerServiceTest() (*MailerServiceTest, error) {
	mailerService := &MailerServiceTest{}
	return mailerService, nil
}

// SendConfirmationMail - used for sending email
func (mailerService *MailerServiceTest) SendConfirmationMail(hostURL string, userEmail string, emailConfirmationToken string, requestID string) error {
	return nil
}

// SendResetPasswordMail - used for sending email
func (mailerService *MailerServiceTest) SendResetPasswordMail(hostURL string, userEmail string, resetToken string, requestID string) error {
	return nil
}

// SendChangeMail - used for sending email
func (mailerService *MailerServiceTest) SendChangeMail(hostURL string, userEmail string, newEmail string, resetToken string, requestID string) error {
	return nil
}

// SendMail - used for sending email
func (mailerService *MailerServiceTest) SendMail(msg common.Email) error {
	return nil
}

// CreateMailerServiceTest - init mailer
func CreateMailerServiceTest(log *zap.Logger) (*MailerServiceTest, error) {
	mailerService, err := NewMailerServiceTest()
	if err != nil {
		log.Error("Error",
			zap.Error(err))
		return nil, err
	}
	return mailerService, nil
}
