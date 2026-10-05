package openapidto

type (
	QrisTransactionPaidPayload struct {
		RequestID     string `json:"requestId"`
		TransactionID string `json:"transactionId"`
		MerchantID    string `json:"merchantId"`
		QrID          string `json:"qrId"`
	}
)
