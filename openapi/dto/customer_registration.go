package openapidto

type (
	BindingDataResponse struct {
		PartnerId   string `json:"partner_id"`
		PartnerName string `json:"partner_name"`
		PhoneNo     string `json:"phone_no"`
		FallbackUrl string `json:"fallback_url"`
	}

	CheckRegisteredAccountRequest struct {
		PhoneNo     string `json:"phone_no" validate:"required"`
		ReferenceNo string `json:"-"`
	}
	CheckRegisteredAccountResponse struct {
		PhoneNo      string `json:"phone_no"`
		IsRegistered bool   `json:"is_registered"`
	}

	SendOTPRequest struct {
		PhoneNo     string `json:"phone_no" validate:"required"`
		ReferenceNo string `json:"-"`
	}
	VerifyOTPRequest struct {
		PhoneNo     string `json:"phone_no" validate:"required"`
		OTP         string `json:"otp" validate:"required"`
		ReferenceNo string `json:"-"`
	}
	CheckPINRequest struct {
		PIN         string `json:"pin" validate:"required"`
		PhoneNo     string `json:"phone_no" validate:"required"`
		ReferenceNo string `json:"-"`
	}

	CreateAccountRequest struct {
		Email          string `json:"email" validate:"omitempty,email"`
		Name           string `json:"name" validate:"required"`
		PhoneNo        string `json:"phone_no" validate:"required"`
		PIN            string `json:"pin" validate:"required"`
		ConfirmPIN     string `json:"confirm_pin" validate:"required"`
		AdditionalInfo struct {
			ReferralCode     string `json:"referral_code"`
			SecurityQuestion string `json:"security_question" validate:"required"`
			SecurityAnswer   string `json:"security_answer" validate:"required"`
		} `json:"additional_info"`
		ReferenceNo string `json:"-"`
	}
	BindAccountRequest struct {
		PhoneNo     string `json:"phone_no" validate:"required"`
		PIN         string `json:"pin" validate:"required"`
		ConfirmPIN  string `json:"confirm_pin" validate:"required"`
		ReferenceNo string `json:"-"`
	}
	BindAccountResponse struct {
		ReferenceNo     string                 `json:"reference_no"`
		AccessToken     string                 `json:"access_token"`
		RefreshToken    string                 `json:"refresh_token"`
		AccessTokenInfo BindingAccessTokenInfo `json:"access_token_info"`
		UserInfo        BindingUserInfo        `json:"user_info"`
		AdditionalInfo  any                    `json:"additional_info"`
	}
	BindingAccessTokenInfo struct {
		ExpiresIn        string `json:"expires_in"`
		RefreshExpiresIn string `json:"refresh_expires_in"`
		TokenStatus      string `json:"token_status"`
	}
	BindingUserInfo struct {
		PublicUserId  string `json:"public_user_id"`
		AccountNumber string `json:"account_number"`
	}

	CheckCreateAccountRequest struct {
		Email       string `json:"email" validate:"omitempty,email"`
		PhoneNo     string `json:"phone_no" validate:"required"`
		ReferenceNo string `json:"-"`
	}
)
