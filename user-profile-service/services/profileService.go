package services

import (
	"log"
	"shared/utils"
)

func FetchUserProfileWithKYC(phoneNumber string) (map[string]interface{}, error) {
	var user struct {
		UserID       int    `json:"user_id"`
		PhoneNumber  string `json:"phone_number"`
		AadharNumber string `json:"aadhar_number"`
		PanNumber    string `json:"pan_number"`
		KYCStatus    string `json:"kyc_status"`
	}

	err := utils.DB.Raw(`
		SELECT * FROM fetch_user_profile_with_kyc(?)
	`, phoneNumber).Scan(&user).Error

	if err != nil {
		log.Println("Error fetching user profile with KYC:", err)
		return nil, err
	}

	profile := make(map[string]interface{})
	profile["user_id"] = user.UserID
	profile["phone_number"] = user.PhoneNumber
	profile["aadhar_number"] = user.AadharNumber
	profile["pan_number"] = user.PanNumber
	profile["kyc_status"] = user.KYCStatus

	return profile, nil
}
