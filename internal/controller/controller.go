package controller

type Controller struct {
	RequestOtpController RequestOtpController
	VerifyOtpController  VerifyOtpController
	UserFcmController    UserFcmController
	LogMessageController LogMessageController
}
