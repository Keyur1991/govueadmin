package locale

type entry struct {
	tag, key string
	msg      interface{}
}

var entries = [...]entry{
	{"en", "login_success", "You have successfully logged in"},
	{"en", "login_failed", "Email or password is incorrect"},
	{"en", "login_internal_err_json_decode", "Invalid request body"},
	{"en", "login_email_dont_exist", "Account doesn't exist with the email specfied, please create an account then try again to login."},
	{"en", "login_wrong_password", "You have provided wrong password"},
	{"en", "login_cant_gen_token", "Unable to generate jwt token"},
	{"en", "logout_success", "You have successfully logged out"},
}
