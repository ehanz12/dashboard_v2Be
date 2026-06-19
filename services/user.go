package services

import (
	"be_dashboard/database"
	"be_dashboard/dto/requests"
	"be_dashboard/models"
	"be_dashboard/utils"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func EditAuthService(userID string, r requests.EditRequest) error {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return errors.New("internal server error")
	}

	var user models.Users
	if err := tx.Where("id = ?", userID).First(&user).Error; err != nil {
		tx.Rollback()
		return errors.New("user not found !")
	}

	if r.Name != "" {
		user.Name = r.Name
	}
	if r.Email != "" {
		user.Email = r.Email
	}
	if r.NomorHP != "" {
		user.NomorHP = r.NomorHP
	}
	if r.Bio != "" {
		user.Bio = &r.Bio
	}

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to update user")
	}
	tx.Commit()
	return nil
}

func ChangePasswordAuthService(userID string, r requests.ChangePasswordRequest) error {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return errors.New("internal server error")
	}
	var user models.Users
	if err := tx.Where("id = ?", userID).First(&user).Error; err != nil {
		tx.Rollback()
		return errors.New("user not found !")
	}
	if r.CurrentPassword == r.NewPassword {
		tx.Rollback()
		return errors.New("new password must be different from current password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.CurrentPassword)); err != nil {
		tx.Rollback()
		return errors.New("current password is incorrect")
	}
	newPass, _ := bcrypt.GenerateFromPassword([]byte(r.NewPassword), bcrypt.DefaultCost)
	user.Password = string(newPass)
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to update password")
	}
	tx.Commit()
	return nil
}

// VerifyEmailService verifies the email with the provided code
func VerifyEmailService(email, verificationCode string) error {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return errors.New("internal server error")
	}

	var user models.Users
	if err := tx.Where("email = ?", email).First(&user).Error; err != nil {
		tx.Rollback()
		return errors.New("user not found")
	}

	// Check if verification code matches
	if user.VerificationCode != verificationCode {
		tx.Rollback()
		return errors.New("invalid verification code")
	}

	// Check if code is expired
	if user.VerificationExpireAt != nil && time.Now().After(*user.VerificationExpireAt) {
		tx.Rollback()
		return errors.New("verification code has expired")
	}

	// Mark email as verified
	user.EmailVerified = true
	user.VerificationCode = ""
	user.VerificationExpireAt = nil

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to verify email")
	}

	tx.Commit()
	return nil
}

// ForgotPasswordService generates a reset code and saves it to the user
func ForgotPasswordService(email string) (string, error) {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return "", errors.New("internal server error")
	}

	var user models.Users
	if err := tx.Where("email = ?", email).First(&user).Error; err != nil {
		tx.Rollback()
		return "", errors.New("user not found")
	}

	code := utils.GenerateVerificationCode()
	expireAt := time.Now().Add(1 * time.Hour)

	user.ResetPasswordCode = code
	user.ResetPasswordExpireAt = &expireAt

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return "", errors.New("failed to generate reset code")
	}

	tx.Commit()
	return code, nil
}

// ResetPasswordService verifies the code and resets the user's password
func ResetPasswordService(email, code, newPassword string) error {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return errors.New("internal server error")
	}

	var user models.Users
	if err := tx.Where("email = ?", email).First(&user).Error; err != nil {
		tx.Rollback()
		return errors.New("user not found")
	}

	if user.ResetPasswordCode == "" || user.ResetPasswordCode != code {
		tx.Rollback()
		return errors.New("invalid reset password code")
	}

	if user.ResetPasswordExpireAt != nil && time.Now().After(*user.ResetPasswordExpireAt) {
		tx.Rollback()
		return errors.New("reset password code has expired")
	}

	// Hash password baru
	newPass, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		return errors.New("failed to hash password")
	}

	user.Password = string(newPass)
	user.ResetPasswordCode = ""
	user.ResetPasswordExpireAt = nil

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to update password")
	}

	tx.Commit()
	return nil
}

