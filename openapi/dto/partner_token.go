package openapidto

type (
	TokenB2BRequest struct {
		GrantType      string `json:"grantType" validate:"required" mapstructure:"grantType"`
		AdditionalInfo struct {
		} `json:"additionalInfo" mapstructure:"additionalInfo"`
	}
	TokenB2BResponse struct {
		ResponseCode    string `json:"responseCode" mapstructure:"responseCode"`
		ResponseMessage string `json:"responseMessage" mapstructure:"responseMessage"`
		AccessToken     string `json:"accessToken" mapstructure:"accessToken"`
		TokenType       string `json:"tokenType" mapstructure:"tokenType"`
		ExpiresIn       int    `json:"expiresIn" mapstructure:"expiresIn"`
		AdditionalInfo  struct {
		} `json:"additionalInfo" mapstructure:"additionalInfo"`
	}
)
