package services

import (
	"be_dashboard/database"
	"be_dashboard/dto/requests"
	"be_dashboard/models"
	"errors"
	"log"

	"gorm.io/gorm"
)

func SaveFCMToken(userID string, req requests.UserDeviceRequest) error {
	log.Println("========== SAVE FCM TOKEN ==========")
	log.Println("UserID     :", userID)
	log.Println("FCM Token  :", req.FMCToken)
	log.Println("DeviceType :", req.DeviceType)

	if req.FMCToken == "" {
		return errors.New("fcm token is required")
	}

	var device models.UserDevice

	// Cari apakah token sudah pernah tersimpan
	err := database.DB.
		Where("fmc_token = ?", req.FMCToken). // sesuaikan dengan nama kolom database
		First(&device).Error

	// Token belum ada → INSERT
	if errors.Is(err, gorm.ErrRecordNotFound) {

		newDevice := models.UserDevice{
			UserID:     userID,
			FMCToken:   req.FMCToken,
			DeviceType: req.DeviceType,
		}

		if err := database.DB.Create(&newDevice).Error; err != nil {
			log.Println("Insert Device Error :", err)
			return err
		}

		log.Println("Device inserted successfully")
		return nil
	}

	// Error selain record not found
	if err != nil {
		log.Println("Database Error :", err)
		return err
	}

	// Token sudah ada → UPDATE
	device.UserID = userID
	device.DeviceType = req.DeviceType

	if err := database.DB.Save(&device).Error; err != nil {
		log.Println("Update Device Error :", err)
		return err
	}

	log.Println("Device updated successfully")
	return nil
}

func DeleteFCMToken(userID, fcmToken string) error {
	if fcmToken == "" {
		return errors.New("fmc token is required")
	}

	result := database.DB.
		Where("user_id = ? AND fmc_token = ?", userID, fcmToken).
		Delete(&models.UserDevice{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("device not found")
	}

	return nil
}
