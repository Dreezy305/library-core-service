package scripts

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/dreezy305/library-core-service/internal/auth/service"
	"github.com/dreezy305/library-core-service/internal/types"
)

func SeedUsers(authService *service.AuthService, count int) {
	gofakeit.Seed(time.Now().UnixNano())
	for i := 0; i < count; {
		fmt.Println(gofakeit.Email())
		payload := &types.UserType{
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Email:     gofakeit.Email(),
			Password:  gofakeit.Password(true, true, true, true, false, 12),
			Role:      "member",
		}

		err := authService.RegisterUserService(payload)
		if err != nil {
			fmt.Printf("Retrying user %d: %v\n", i+1, err)
			continue
		}

	}
}
