package config

type AppConfig struct {
	Port            string
	DatabaseURL     string
	JWTSecret       string
	Path            string
	ZeptoApiKey     string
	SMTPSenderName  string
	SMTPSenderEmail string
	SMTPNoReply     string
	ZeptobBaseURL   string
}
