package handlers

import (
	"context"
	"os"

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

	user := models.Users{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(passwordHash),
		NomorHP:  req.NomorHP,
		Bio:      &req.Bio,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "register success",
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
	}
	if req.IDToken == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "id_token is required"})
	}

	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	if clientID == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "GOOGLE_CLIENT_ID not configured"})
	}

	// verify id token
	payload, err := idtoken.Validate(context.Background(), req.IDToken, clientID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid google id token"})
	}

	// extract email and name from payload
	emailI, _ := payload.Claims["email"]
	nameI, _ := payload.Claims["name"]
	email, _ := emailI.(string)
	name, _ := nameI.(string)
	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "email not found in token"})
	}

	// find or create user
	var user models.Users
	err = database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		// create new user
		user = models.Users{
			Name:  name,
			Email: email,
		}
		if err := database.DB.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create user"})
		}
	}

	// generate jwt
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to generate token"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"token":   token,
		"data":    mappers.ToUserResponse(user),
	})
}
