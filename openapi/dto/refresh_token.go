package openapidto

type (
	GrantTypeB2B2C string

	TokenB2B2CRequest struct {
		GrantType      string `json:"grantType" validate:"required"`
		AuthCode       string `json:"authCode"`
		RefreshToken   string `json:"refreshToken"`
	}
	TokenB2B2CResponse struct {
		ResponseCode           string `json:"responseCode"`
		ResponseMessage        string `json:"responseMessage"`
		AccessToken            string `json:"accessToken"`
		TokenType              string `json:"tokenType"`
		AccessTokenExpiryTime  string `json:"accessTokenExpiryTime"`
		RefreshToken           string `json:"refreshToken"`
		RefreshTokenExpiryTime string `json:"refreshTokenExpiryTime"`
	}
)

var (
	GrantTypeAuthorizationCode GrantTypeB2B2C = "AUTHORIZATION_CODE"
	GrantTypeRefreshToken      GrantTypeB2B2C = "REFRESH_TOKEN"
)
