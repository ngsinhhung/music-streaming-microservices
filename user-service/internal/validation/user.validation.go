package validation

type UserRegisterSchema struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Avatar   string `json:"avatar" validate:"omitempty,url"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

type VerifyOTPRequest struct {
	Email string `json:"email" validate:"required"`
	OTP   int    `json:"otp" validate:"required"`
}

type UserLoginSchema struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}
