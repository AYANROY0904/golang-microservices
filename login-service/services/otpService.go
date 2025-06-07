package services

import (
	"fmt"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
	"shared/utils"
	"github.com/getsentry/sentry-go"
)

var TWILIO_ACCOUNT_SID string = utils.LoadConfig().TWILIO_ACCOUNT_SID
var TWILIO_AUTH_TOKEN string = utils.LoadConfig().TWILIO_AUTH_TOKEN
var VERIFY_SERVICE_SID string = utils.LoadConfig().TWILIO_SERVICE_SID
var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: TWILIO_ACCOUNT_SID,
	Password: TWILIO_AUTH_TOKEN,
})

func SendOTP(phoneNumber string) error {
	params := &openapi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(VERIFY_SERVICE_SID, params)
	if err != nil {
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetTag("phone_number", phoneNumber)
			sentry.CaptureException(err)
		})
		return err
	}
	fmt.Printf("Sent verification '%s'\n", *resp.Sid)
	return nil
}

func VerifyOTP(phoneNumber, otp string) (bool, error) {
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(otp)

	resp, err := client.VerifyV2.CreateVerificationCheck(VERIFY_SERVICE_SID, params)
	if err != nil {
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetTag("phone_number", phoneNumber)
			sentry.CaptureException(err)
		})
		return false, err
	} else if *resp.Status == "approved" {
		return true, nil
	}
	return false, fmt.Errorf("incorrect OTP")
}
