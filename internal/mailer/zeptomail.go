package mailer

import (
	"fmt"
	"net/http"

	"github.com/dreezy305/library-core-service/internal/config"
	"github.com/dreezy305/library-core-service/internal/types"
	"github.com/jiyamathias/zeptomail"
)

func SendEmail(SendEmailPayload *types.SendEmail) {
	zeptomailToken := config.Load().ZeptoApiKey
	baseURL := config.Load().ZeptobBaseURL
	senderAddress := config.Load().SMTPSenderEmail
	senderName := config.Load().SMTPSenderName

	client := zeptomail.New(http.DefaultClient, zeptomailToken, baseURL)

	req := zeptomail.SendHTMLEmailReq{
		To: []zeptomail.SendEmailTo{
			{
				EmailAddress: zeptomail.EmailAddress{
					Address: SendEmailPayload.EmailAddress,
					Name:    SendEmailPayload.Name,
				},
			},
		},
		From: zeptomail.EmailAddress{
			Address: senderAddress,
			Name:    senderName,
		},
		Subject:  SendEmailPayload.Subject,
		Htmlbody: SendEmailPayload.HtmlBody,
	}

	res, err := client.SendHTMLEmail(req)

	if err != nil {
		fmt.Printf("This is the error: %v", err.Error())
	}

	for _, e := range res.Data {
		fmt.Printf("response message: %v\n", e.Message)
	}
}
