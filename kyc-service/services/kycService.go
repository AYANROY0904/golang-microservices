package services

import (
	"log"
	"fmt"
	"shared/utils"

	"github.com/getsentry/sentry-go"
)

func StoreKYCData(phoneNumber, aadharNumber, panNumber string) (string, error) {
	var userID int
	var resultMessage string

	err := utils.DB.Raw(`
		SELECT id FROM users WHERE phone_number = ?
	`, phoneNumber).Scan(&userID).Error

	if err != nil {
		sentry.CaptureException(err)
		log.Println("Error fetching user ID:", err)
		return "", err
	}

	if userID == 0 {
		return "", fmt.Errorf("user does not exist with the given phone number")
	}

	err = utils.DB.Raw(`
		SELECT store_or_update_kyc_data(?, ?, ?, ?)
	`, userID, phoneNumber, aadharNumber, panNumber).Scan(&resultMessage).Error

	if err != nil {
		sentry.CaptureException(err)
		log.Println("Error executing stored procedure:", err)
		return "", err
	}

	log.Println(resultMessage)
	return resultMessage, nil
}
