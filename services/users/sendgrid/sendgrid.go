package sendgrid

import (
	"context"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"go.uber.org/zap"
)

const (
	FROM_ADDRESS = "matthieu@heybutler.io"
	USER_NAME    = "Matthieu"
)

type EmailClient struct {
	apiKey string
}

func NewEmailClient(apiKey string) *EmailClient {
	return &EmailClient{apiKey}
}

func (c *EmailClient) SendEmail(ctx context.Context, email *mail.SGMailV3) error {
	request := sendgrid.GetRequest(c.apiKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = mail.GetRequestBody(email)
	request.Body = Body
	response, err := sendgrid.API(request)
	if err != nil {
		logger.Error(ctx, "Failed to send email", zap.Error(err))
	}
	logger.Debug(ctx, "Successfully send email", zap.Any("respponse", response))
	return err
}
