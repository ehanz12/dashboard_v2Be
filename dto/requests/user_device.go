package requests

type UserDeviceRequest struct {
	FMCToken   string `json:"fcm_token" validate:"required"`
	DeviceType string `json:"device_type" validate:"required,oneof=ios android"`
}