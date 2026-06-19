package handlers

import (
	"context"
	"os"
	"time"

	"be_dashboard/database"
	"be_dashboard/dto/requests"
	"be_dashboard/mappers"
	"be_dashboard/models"
	"be_dashboard/services"
	"be_dashboard/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/idtoken"
)

func CreateAuthHandler(c *fiber.Ctx) error {
	var req requests.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid payload",
		})
	}

	// validasi
	if req.Email == "" || req.Password == "" || req.Name == "" || req.NomorHP == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "all fields are required, except bio",
		})
	}

	// cek email
	var exist models.Users
	err := database.DB.Where("email = ?", req.Email).First(&exist).Error
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "email already used!",
		})
	}

	// hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to hash password",
		})
	}

	// generate verification code
	verificationCode := utils.GenerateVerificationCode()
	expireAt := time.Now().Add(24 * time.Hour)

	user := models.Users{
		Name:                 req.Name,
		Email:                req.Email,
		Password:             string(passwordHash),
		NomorHP:              req.NomorHP,
		Bio:                  &req.Bio,
		EmailVerified:        false,
		VerificationCode:     verificationCode,
		VerificationExpireAt: &expireAt,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create user",
		})
	}

	// send verification email
	if err := services.SendVerificationEmail(user.Email, verificationCode); err != nil {
		// log error but don't fail the registration
		// in production, you might want to retry or handle this differently
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "register success but failed to send verification email. Please try again later.",
			"email":   user.Email,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "register success! Please check your email for verification code.",
		"email":   user.Email,
	})
}

func LoginAuthHandler(c *fiber.Ctx) error {
	// parsing payload
	var req requests.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid payload",
		})
	}
	// validasi
	if req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "email and password are required",
		})
	}
	// cek email
	var user models.Users
	err := database.DB.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid email or password",
		})
	}

	// check if email is verified
	if !user.EmailVerified {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "email not verified. Please check your email for verification code.",
		})
	}

	// cek password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid email or password",
		})
	}
	// generate token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to generate token",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Successfully logged in",
		"token":   token,
	})
}

func MeAuthHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var user models.Users
	err := database.DB.First(&user, "id = ?", userID).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch user",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "User data fetched successfully",
		"data":    mappers.ToUserResponse(user),
	})
}

func EditMeAuthHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var req requests.EditRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid payload",
		})
	}
	if req.Email == "" || req.Name == "" || req.NomorHP == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "email, name, and nomor_hp are required, bio is optional",
		})
	}
	if err := services.EditAuthService(userID, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "User data updated successfully",
	})
}

func ChangePasswordAuthHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var req requests.ChangePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid payload",
		})
	}
	if req.CurrentPassword == "" || req.NewPassword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "current password and new password are required",
		})
	}
	if err := services.ChangePasswordAuthService(userID, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully to change password !"})
}

// GoogleAuthHandler accepts a Google ID token (from client) and verifies it,
// then finds or creates the user and returns a JWT.
func GoogleAuthHandler(c *fiber.Ctx) error {
	var req requests.GoogleAuthRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid payload",
		})
	}

	if req.IDToken == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id_token is required",
		})
	}

	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	if clientID == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "GOOGLE_CLIENT_ID not configured",
		})
	}

	// Verify Google token
	payload, err := idtoken.Validate(context.Background(), req.IDToken, clientID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid google id token",
		})
	}

	// Ambil data dari token Google
	emailI, _ := payload.Claims["email"]
	nameI, _ := payload.Claims["name"]

	email, _ := emailI.(string)
	name, _ := nameI.(string)

	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "email not found in token",
		})
	}

	// Cari user
	var user models.Users

	err = database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		// User belum ada -> buat baru
		user = models.Users{
			Name:          name,
			Email:         email,
			EmailVerified: true,
		}

		if err := database.DB.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to create user",
			})
		}
	} else {
		// User sudah ada -> langsung verifikasi
		if !user.EmailVerified {
			user.EmailVerified = true
			user.VerificationCode = ""
			user.VerificationExpireAt = nil

			if err := database.DB.Save(&user).Error; err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "failed to update user verification status",
				})
			}
		}
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to generate token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"token":   token,
		"data":    mappers.ToUserResponse(user),
	})
}


// VerifyEmailHandler verifies the user's email with the provided verification code
func VerifyEmailHandler(c *fiber.Ctx) error {
	var req requests.VerifyEmailRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid payload",
		})
	}

	if req.Email == "" || req.VerificationCode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "email and verification_code are required",
		})
	}

	if err := services.VerifyEmailService(req.Email, req.VerificationCode); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "email verified successfully",
	})
}

// ForgotPasswordHandler handles requesting a password reset code
func ForgotPasswordHandler(c *fiber.Ctx) error {
	var req requests.ForgotPasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid payload",
		})
	}

	if req.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "email is required",
		})
	}

	code, err := services.ForgotPasswordService(req.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := services.SendResetPasswordEmail(req.Email, code); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "reset code generated, but failed to send email. Please try again later.",
			"email":   req.Email,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "reset password code has been sent to your email",
		"email":   req.Email,
	})
}

// ResetPasswordHandler handles resetting the password with the OTP code
func ResetPasswordHandler(c *fiber.Ctx) error {
	var req requests.ResetPasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid payload",
		})
	}

	if req.Email == "" || req.Code == "" || req.NewPassword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "email, code, and new_password are required",
		})
	}

	if len(req.NewPassword) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "new_password must be at least 6 characters",
		})
	}

	if err := services.ResetPasswordService(req.Email, req.Code, req.NewPassword); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "password has been reset successfully",
	})
}



