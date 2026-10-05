package openapidto

type BalanceType string

var (
	BalanceTypeBalance BalanceType = "BALANCE"
	BalanceTypePoint   BalanceType = "POINT"
)

type (
	PinRequest struct {
		LinkageToken string `json:"linkage_token"`
		Pin          string `json:"pin"`
	}

	BindingRequest struct {
		MerchantId     string                   `json:"merchantId" validate:"required"`
		PhoneNo        string                   `json:"phoneNo" validate:"required"`
		AdditionalInfo BindingReqAdditionalInfo `json:"additionalInfo"`
	}
	BindingReqAdditionalInfo struct {
		FallbackUrl      string `json:"fallbackUrl" validate:"required"`
		FinishBindingUrl string `json:"finishBindingUrl" validate:"required"`
		ExternalUid      string `json:"externalUid" validate:"required"`
	}
	BindingResponse struct {
		ResponseCode    string                    `json:"responseCode"`
		ResponseMessage string                    `json:"responseMessage"`
		ReferenceNo     string                    `json:"referenceNo"`
		RedirectUrl     string                    `json:"redirectUrl"`
		AdditionalInfo  BindingRespAdditionalInfo `json:"additionalInfo"`
	}
	BindingRespAdditionalInfo struct {
		AuthCode string `json:"authCode"`
	}

	BindingInquiryRequest struct {
		AdditionalInfo AdditionalInfoRefreshToken `json:"additionalInfo"`
	}
	BindingInquiryResponse struct {
		ResponseCode    string `json:"responseCode"`
		ResponseMessage string `json:"responseMessage"`
		ReferenceNo     string `json:"referenceNo"`
	}

	UnbindingRequest struct {
		MerchantId     string                     `json:"merchantId" validate:"required"`
		AdditionalInfo AdditionalInfoRefreshToken `json:"additionalInfo"`
	}
	UnbindingResponse struct {
		ResponseCode    string `json:"responseCode"`
		ResponseMessage string `json:"responseMessage"`
		ReferenceNo     string `json:"referenceNo"`
		UnlinkResult    string `json:"unlinkResult"`
	}

	AdditionalInfoRefreshToken struct {
		RefreshToken string `json:"refreshToken" validate:"required"`
	}

	BalanceInquiryRequest struct {
		BalanceTypes []string `json:"balanceTypes" validate:"required"`
		UserId       string   `json:"-"`
	}
	BalanceInquiryResponse struct {
		ResponseCode    string                       `json:"responseCode"`
		ResponseMessage string                       `json:"responseMessage"`
		AccountInfos    []BalanceInquirAccountInfo   `json:"accountInfos"`
		AdditionalInfo  AdditionalInfoBalanceInquiry `json:"additionalInfo"`
	}
	BalanceInquirAccountInfo struct {
		BalanceType BalanceType `json:"balanceType"`
		Amount      Amount      `json:"amount"`
	}
	AdditionalInfoBalanceInquiry struct {
		UserStatus string `json:"userStatus"`
	}
)
