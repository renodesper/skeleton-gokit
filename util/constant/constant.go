package constant

var (
	VerificationTypeRegistration  = "registration"
	VerificationTypeResetPassword = "reset-password"

	EmailTypeWelcome                   = "welcome"
	EmailTypeResetPassword             = "reset-password"
	EmailTypeResetPasswordNotification = "reset-password-notification"
	EmailTypeMaintenance               = "maintenance"

	EmailSubjectWelcome                   = "Welcome to %s!"
	EmailSubjectResetPassword             = "Reset Password Request for your account in %s"
	EmailSubjectResetPasswordNotification = "Your Password Has Been Reset"
	EmailSubjectMaintenance               = "maintenance"
)
