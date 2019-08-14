package sender

import (
	"github.com/mailgun/mailgun-go"
	"github.com/wexel-nath/wexel-email/pkg/config"
	"github.com/wexel-nath/wexel-email/pkg/logger"
)

var (
	client mailgun.Mailgun
)

func configure() {
	domain := config.GetMailgunDomain()
	apiKey := config.GetMailgunApiKey()

	client = mailgun.NewMailgun(domain, apiKey)
}

func Send(
	to string,
	from string,
	subject string,
	html string,
	text string,
	params map[string]interface{},
) (string, error) {
	if client == nil {
		configure()
	}

	message := client.NewMessage(from, subject, text, to)
	message.SetHtml(html)

	for key, value :=  range params {
		err := message.AddVariable(key, value)
		if err != nil {
			logger.Error(err)
			// This is not a critical error. Continue
		}
	}

	resp, id, err := client.Send(message)
	logger.Info("Email sent to Mailgun. resp=%s id=%s err=%v", resp, id, err)

	return resp, err
}
