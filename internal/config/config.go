package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &AppConfig{
		Port:            os.Getenv("PORT"),
		DatabaseURL:     os.Getenv("DATABASE_URL"),
		JWTSecret:       os.Getenv("JWT_SECRET"),
		Path:            os.Getenv("API_VERSION"),
		ZeptoApiKey:     os.Getenv("ZEPTO_MAIL_TOKEN"),
		SMTPSenderName:  os.Getenv("SMTP_SENDER_NAME"),
		SMTPSenderEmail: os.Getenv("SMTP_SENDER_EMAIL"),
		SMTPNoReply:     os.Getenv("SMTP_NO_REPLY"),
		ZeptobBaseURL:   os.Getenv("ZEPTO_BASEURL"),
	}

	// Optional: validate required fields
	if cfg.Port == "" || cfg.DatabaseURL == "" || cfg.JWTSecret == "" || cfg.Path == "" || cfg.ZeptoApiKey == "" || cfg.ZeptobBaseURL == "" {
		log.Fatal("Missing required environment variables")
	}

	return cfg
}
