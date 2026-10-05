package openapidto

type (
	MerchantInfoRequest struct {
		MerchantId string `json:"merchantId" validate:"required"`
	}
	MerchantInfoResponse struct {
		ResponseCode    string       `json:"responseCode"`
		ResponseMessage string       `json:"responseMessage"`
		MerchantInfo    MerchantInfo `json:"merchantInfo"`
	}
	MerchantInfo struct {
		Name          string `json:"name"`
		Types         string `json:"types"`
		Criteria      string `json:"criteria"`
		Category      string `json:"category"`
		PostCode      string `json:"postCode"`
		Size          string `json:"size"`
		MPAN          string `json:"mpan"`
		NMID          string `json:"nmid"`
		MID           string `json:"mid"`
		MCC           string `json:"mcc"`
		AnnualIncome  string `json:"annualIncome"`
		TaxNumber     string `json:"taxNumber"`
		PhotoLink     string `json:"photoLink"`
		Location      any    `json:"location"`
		TerminalLabel string `json:"terminalLabel"`
	}

	MerchantBalanceInquiryRequest struct {
		MerchantId string `json:"merchantId" validate:"required"`
	}
	MerchantBalanceInquiryResponse struct {
		ResponseCode    string                     `json:"responseCode"`
		ResponseMessage string                     `json:"responseMessage"`
		AccountInfos    []BalanceInquirAccountInfo `json:"accountInfos"`
	}
)
