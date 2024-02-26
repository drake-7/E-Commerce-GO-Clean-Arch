package interfaces

import "github.com/msazad/go-Ecommerce/pkg/utils/models"

type OtpUsecase interface {
	VerifyOTP(code models.VerifyData) (models.UserToken, error)
	SendOTP(phone string) error
}
