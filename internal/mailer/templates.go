package mailer

import (
	"github.com/dreezy305/library-core-service/internal/renderer"
	"github.com/dreezy305/library-core-service/internal/types"
)

func ForgotPasswordtemplate(payload *types.ResetTokenHtmlBodyStruct) (string, error) {
	return renderer.RenderTemplate("forgot_password.html", payload)
}
