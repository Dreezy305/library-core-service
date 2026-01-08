package handler

import (
	"fmt"

	"github.com/dreezy305/library-core-service/internal/auth/service"
	"github.com/dreezy305/library-core-service/internal/types"
	"github.com/dreezy305/library-core-service/internal/validators"
	"github.com/gofiber/fiber/v3"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{Service: s}
}

func (h *AuthHandler) RegisterUserHandler(c fiber.Ctx) error {
	// Implementation of user registration handler goes here
	var payload types.UserType
	if err := c.Bind().Body(&payload); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusForbidden).JSON(validators.FormatValidationError(err))
	}
	fmt.Println(payload)
	errs := validators.ValidateStruct(payload)
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validators.FormatValidationError(errs))
	}
	error := h.Service.RegisterUserService(&payload)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validators.FormatValidationError(errs))
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Account activated successfully"})
}

func (h *AuthHandler) LoginUserHandler(c fiber.Ctx) error {
	var payload types.LoginUserType
	if err := c.Bind().Body(&payload); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusForbidden).JSON(validators.FormatValidationError(err))
	}
	fmt.Println(payload)
	errs := validators.ValidateStruct(payload)
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validators.FormatValidationError(errs))
	}
	// Implementation of user login handler goes here
	token, err := h.Service.LoginUserService(&payload)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Login successful", "token": token})
}

func (h *AuthHandler) VerifyEmailHandler(c fiber.Ctx) error {
	// Implementation of email verification handler goes here
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Email verified successfully"})
}

func (h *AuthHandler) ForgotPasswordHandler(c fiber.Ctx) error {
	// Implementation of forgot password handler goes here
	var payload types.ForgotPassword
	if err := c.Bind().Body(&payload); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusForbidden).JSON(validators.FormatValidationError(err))
	}
	fmt.Println(payload)
	errs := validators.ValidateStruct(payload)
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validators.FormatValidationError(errs))
	}
	err := h.Service.ForgotPasswordService(&payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validators.FormatValidationError(errs))
	}
	// send email coontaiing OTP
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("OTP has been sent to your email: %s", payload.Email)})
}

func (h *AuthHandler) ResetPasswordHandler(c fiber.Ctx) error {
	// Implementation of reset password handler goes here
	var payload types.ResetPassword
	if err := c.Bind().Body(&payload); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusForbidden).JSON(validators.FormatValidationError(err))
	}
	errs := validators.ValidateStruct(payload)
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validators.FormatValidationError(errs))
	}

	err := h.Service.CompareOldNewPasswordsAndResetToken(&payload)

	if err != nil {
		fmt.Println(err, "err")
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Password reset successfully"})
}
