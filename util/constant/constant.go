package constant

var (
	VerificationTypeRegistration   = "registration"
	VerificationTypeResetPassword  = "reset-password"
	VerificationTypeUpdatePassword = "update-password"

	EmailTypeWelcome       = "welcome"
	EmailTypeResetPassword = "reset-password"
	EmailTypeMaintenance   = "maintenance"

	EmailSubjectWelcome       = "Welcome to %s!"
	EmailSubjectResetPassword = "Reset Password Request for your account in %s"
	EmailSubjectMaintenance   = "maintenance"
)
