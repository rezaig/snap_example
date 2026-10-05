package openapidto

const (
	QrMpmStatusSuccess QrMpmStatus = "00" // Success
	// QrMpmStatusInitiated QrMpmStatus = "01" // Initiated
	// QrMpmStatusPaying    QrMpmStatus = "02" // Paying
	QrMpmStatusPending  QrMpmStatus = "03" // Pending
	QrMpmStatusRefunded QrMpmStatus = "04" // Refunded
	QrMpmStatusCanceled QrMpmStatus = "05" // Canceled
	QrMpmStatusFailed   QrMpmStatus = "06" // Failed
	// QrMpmStatusNotFound  QrMpmStatus = "07" // Not Found
	QrMpmStatusExpired QrMpmStatus = "08" // Expired
	// QrMpmStatusRejected  QrMpmStatus = "09" // Rejected
)

type (
	QrMpmStatus string

	GenerateDynamicQrMpmRequest struct {
		PartnerReferenceNo string `json:"partnerReferenceNo" validate:"required"`
		MerchantId         string `json:"merchantId" validate:"required"`
		Amount             Amount `json:"amount"`
	}
	GenerateDynamicQRMPResponse struct {
		ResponseCode       string `json:"responseCode"`
		ResponseMessage    string `json:"responseMessage"`
		ReferenceNo        string `json:"referenceNo"`
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		QrContent          string `json:"qrContent"`
		QrURL              string `json:"qrUrl"`
		QrImage            string `json:"qrImage"`
		AdditionalInfo     any    `json:"additionalInfo,omitempty"`
	}

	CheckQrMpmStatusRequest struct {
		OriginalReferenceNo        string `json:"originalReferenceNo" validate:"required"`
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo" validate:"required"`
		ServiceCode                string `json:"serviceCode" validate:"required"` // service code of original transaction, it should be service code of generate dynamic QR (47)
		MerchantId                 string `json:"merchantId" validate:"required"`
	}
	CheckQrMpmStatusResponse struct {
		ResponseCode               string      `json:"responseCode"`
		ResponseMessage            string      `json:"responseMessage"`
		OriginalReferenceNo        string      `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string      `json:"originalPartnerReferenceNo"`
		ServiceCode                string      `json:"serviceCode"`
		LatestTransactionStatus    QrMpmStatus `json:"latestTransactionStatus"`
		TransactionStatusDesc      string      `json:"transactionStatusDesc"`
		PaidTime                   *string     `json:"paidTime"`
		Amount                     Amount      `json:"amount"`
		AdditionalInfo             any         `json:"additionalInfo,omitempty"`
	}

	RefundQrMpmTrxRequest struct {
		OriginalReferenceNo        string `json:"originalReferenceNo" validate:"required"`
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo" validate:"required"`
		PartnerRefundNo            string `json:"partnerRefundNo" validate:"required"`
		MerchantId                 string `json:"merchantId" validate:"required"`
		Reason                     string `json:"reason" validate:"required"`
	}
	RefundQrMpmTrxResponse struct {
		ResponseCode               string `json:"responseCode"`
		ResponseMessage            string `json:"responseMessage"`
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		PartnerRefundNo            string `json:"partnerRefundNo"`
		RefundNo                   string `json:"refundNo"`
		RefundTime                 string `json:"refundTime"`
		RefundAmount               Amount `json:"refundAmount"`
		AdditionalInfo             any    `json:"additionalInfo,omitempty"`
	}

	QrMpmPaymentNotifyRequest struct {
		OriginalReferenceNo        string      `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string      `json:"originalPartnerReferenceNo"`
		LatestTransactionStatus    QrMpmStatus `json:"latestTransactionStatus"`
		TransactionStatusDesc      string      `json:"transactionStatusDesc"`
		PaidTime                   *string     `json:"paidTime"`
		Amount                     Amount      `json:"amount"`
		AdditionalInfo             any         `json:"additionalInfo,omitempty"`
	}
)

func (s QrMpmStatus) Label() string {
	switch s {
	case QrMpmStatusSuccess:
		return "Success"
	// case QrMpmStatusInitiated:
	// 	return "Initiated"
	// case QrMpmStatusPaying:
	// 	return "Paying"
	case QrMpmStatusPending:
		return "Pending"
	case QrMpmStatusRefunded:
		return "Refunded"
	case QrMpmStatusCanceled:
		return "Canceled"
	// case QrMpmStatusNotFound:
	// 	return "Not Found"
	case QrMpmStatusExpired:
		return "Expired"
	// case QrMpmStatusRejected:
	// 	return "Rejected"
	default:
		return "Failed"
	}
}
