package services

import (
	"be_dashboard/database"
	"be_dashboard/dto/requests"
	"be_dashboard/models"
	"errors"
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