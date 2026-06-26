package services

import (
	"be_dashboard/database"
	"be_dashboard/dto/requests"
	"be_dashboard/models"
	"errors"

	"gorm.io/gorm"
)

func SaveFCMToken(userID string, req requests.UserDeviceRequest) error {
	if req.FMCToken == "" {
		return errors.New("fcm token is required")
	}

	var userDevice models.UserDevice

	err := database.DB.
		Where("fcm_token = ?", req.FMCToken).
		First(&userDevice).
		Error

	if err != nil {
		// Token belum ada -> insert baru
		if errors.Is(err, gorm.ErrRecordNotFound) {

			userDevice = models.UserDevice{
				UserID:     userID,
				FMCToken:   req.FMCToken,
				DeviceType: req.DeviceType,
			}

			return database.DB.Create(&userDevice).Error
		}

		// Error database
		return err
	}

	// Token sudah ada -> update pemilik/device
	userDevice.UserID = userID
	userDevice.DeviceType = req.DeviceType

	return database.DB.
		Model(&userDevice).
		Updates(map[string]interface{}{
			"user_id":     userDevice.UserID,
			"device_type": userDevice.DeviceType,
		}).
		Error
}

func DeleteFCMToken(userID, fcmToken string) error {
	if fcmToken == "" {
		return errors.New("fcm token is required")
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