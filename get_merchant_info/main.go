package main

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"

	openapidto "github.com/rezaig/snap_example/openapi/dto"
)

const (
	baseURL = "https://tpartner-api.digitalbanking.id" // dev
	// baseURL = "https://wallet-api.tcash.telkomsel.co.id" // stg

	// emoney stg
	// clientID          = "emoney"
	// clientSecret      = "r8bIZiMqnewIhgSna1WiRF6hnjFLtuk3"
	// merchantID        = "a4d4b292-b49a-47d4-a991-5b088b030589"
	// partnerPrivateKey = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCnEu5uAqkf1o1WX2Ps21Q2DnGk4/GZ/9/CmNqkdx4e5LD3VchcKaCZ+A57rqettOVaRCfUhD5t3ZployBL5Ca4yGcA1gx9xKYa8fcJPBNCAF9SUtBNmTNEbFCPuarR/D0QdPO8XHaBt3rqsUlCIgljszZchbjlBpAurbq/lveY+ZTDSjCrnoVVQ14UttqD4yYu0lX9+y75zN51UaQ3nAWbkePzCucOO0NZ3vvCqeBH2tTiK2p6uf6SLzFpPGUY27Mj+nq3CCu3TcJ9GvN5EZ7cluyKBQVT6WB6L47ueMkIBKi7pcP+98pGbneCZgIoVDOCYOkLgzaVL1Npi4wj8g/pAgMBAAECggEAQEWYoop0fFKByb3lh3sykJ8K2ed0jI0yC77YsZc/Z5wLdgKykr3S0hhqnerpB3qFFq4McoUzLPmoiTvxUzCgMsqpUVmGsaVCTzBRG+TX9baF9Pn1tzxdzA8tCLVgwOobpaaQEyCw2CL47qsn5wCIAyowmfQ5ViWptsNPiZ+ufVK6p3/OuG/JytAPmo+ufZEGMd6n8i7J0XcwpKugCc1mMSDgivNE7L8vNNJ4dLS9YWrb32xVDdSsWN7Y76D/pE2l8idR1UXe+sxviU2wKe8YLckIQNFiesK1VQ45gFKgR9ayIo7DkBD14GbrOa9lrlH2IsKSwLeGCldvK78vWnjnBwKBgQDa2fqSTcQENU5A8FZHvHnlv89Vy3pVgWdJnZw5DWP1B5vKBNuKR5xar8jOaPCHThP3ijIdMTwPvlCFeEvfKY05DrNBW52CsS1ANxSBgLs4a/atCj1pO9TW4ZraU0qdS8ROC32IgjROr+cnoo6ff+DmoKJj8E4P/9fQYFt49k5wBwKBgQDDbwAPHX8tFqhmhkZqdvR7ZPxu1Tca+MewbsYEuJHl6k5VZQFMs6ba2tHFcRteh0RCewyQgB5gQOW+9cS7G9Wa5elnySerze9hRTIBsJ3SqymFWHoV1H4ebAOT1qLpU/Lg9pt4UTPJo4xWRTlDF3/M6uqUFKPjgG9QHHtm7XukjwKBgAU/Am2tZYyARp7x4++WSgnC3lJ9LTKNho9SMuN/Oa0vAIIIOccHzmyyGAsyoslrirj9XBQtEPaDpmR8rLztvw/mFU/0xULTwnTunRQ5pMNGe0RMoYo7P+/iupaPNpOstEj2p4y1KlHUj6L4l5ilNRvyL9JbeVOS23aISMkOhnTTAoGAJpQd1d4DwdnMKljtQ4zx2/3mWtaaByf++1QSoNHycwlapz7GD+cS0/cIG8qlFXbsQZdatpej52pIL/cB+9GVy/sApS0vOJnxXCk1ouHDdde13Y3Go9KLhuPZnPBsvlSFCGWF8S1OZMp1JH6LvDPCVag14D9mzr0GvDCzJ3FPy/8CgYEAyN7s901aq0iSjTlN2G0zRoUNucsx5dChcdpL6t+97Ugier29fZycHs94d7FG7Th9UUo/+GmnCRJPxYIA1Qk4oJ2HcMIIsDR+P8f4zwvg1llv9X6EnhCxnAS7FSLjajFWVEHi0VwLHWa73Uxg2X3dyBA4GagIVsjZE4qPecNlh40=" // ASTRAPAY_EMONEY_PRIVATE_KEY

	// emoney dev
	clientID          = "emoney"
	clientSecret      = "TXBm0So2a5MhP8rzaIfRpemY6ueqPARG"
	merchantID        = "f1a44c39-ec95-4659-8475-9eedd0e354fb"
	partnerPrivateKey = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCnEu5uAqkf1o1WX2Ps21Q2DnGk4/GZ/9/CmNqkdx4e5LD3VchcKaCZ+A57rqettOVaRCfUhD5t3ZployBL5Ca4yGcA1gx9xKYa8fcJPBNCAF9SUtBNmTNEbFCPuarR/D0QdPO8XHaBt3rqsUlCIgljszZchbjlBpAurbq/lveY+ZTDSjCrnoVVQ14UttqD4yYu0lX9+y75zN51UaQ3nAWbkePzCucOO0NZ3vvCqeBH2tTiK2p6uf6SLzFpPGUY27Mj+nq3CCu3TcJ9GvN5EZ7cluyKBQVT6WB6L47ueMkIBKi7pcP+98pGbneCZgIoVDOCYOkLgzaVL1Npi4wj8g/pAgMBAAECggEAQEWYoop0fFKByb3lh3sykJ8K2ed0jI0yC77YsZc/Z5wLdgKykr3S0hhqnerpB3qFFq4McoUzLPmoiTvxUzCgMsqpUVmGsaVCTzBRG+TX9baF9Pn1tzxdzA8tCLVgwOobpaaQEyCw2CL47qsn5wCIAyowmfQ5ViWptsNPiZ+ufVK6p3/OuG/JytAPmo+ufZEGMd6n8i7J0XcwpKugCc1mMSDgivNE7L8vNNJ4dLS9YWrb32xVDdSsWN7Y76D/pE2l8idR1UXe+sxviU2wKe8YLckIQNFiesK1VQ45gFKgR9ayIo7DkBD14GbrOa9lrlH2IsKSwLeGCldvK78vWnjnBwKBgQDa2fqSTcQENU5A8FZHvHnlv89Vy3pVgWdJnZw5DWP1B5vKBNuKR5xar8jOaPCHThP3ijIdMTwPvlCFeEvfKY05DrNBW52CsS1ANxSBgLs4a/atCj1pO9TW4ZraU0qdS8ROC32IgjROr+cnoo6ff+DmoKJj8E4P/9fQYFt49k5wBwKBgQDDbwAPHX8tFqhmhkZqdvR7ZPxu1Tca+MewbsYEuJHl6k5VZQFMs6ba2tHFcRteh0RCewyQgB5gQOW+9cS7G9Wa5elnySerze9hRTIBsJ3SqymFWHoV1H4ebAOT1qLpU/Lg9pt4UTPJo4xWRTlDF3/M6uqUFKPjgG9QHHtm7XukjwKBgAU/Am2tZYyARp7x4++WSgnC3lJ9LTKNho9SMuN/Oa0vAIIIOccHzmyyGAsyoslrirj9XBQtEPaDpmR8rLztvw/mFU/0xULTwnTunRQ5pMNGe0RMoYo7P+/iupaPNpOstEj2p4y1KlHUj6L4l5ilNRvyL9JbeVOS23aISMkOhnTTAoGAJpQd1d4DwdnMKljtQ4zx2/3mWtaaByf++1QSoNHycwlapz7GD+cS0/cIG8qlFXbsQZdatpej52pIL/cB+9GVy/sApS0vOJnxXCk1ouHDdde13Y3Go9KLhuPZnPBsvlSFCGWF8S1OZMp1JH6LvDPCVag14D9mzr0GvDCzJ3FPy/8CgYEAyN7s901aq0iSjTlN2G0zRoUNucsx5dChcdpL6t+97Ugier29fZycHs94d7FG7Th9UUo/+GmnCRJPxYIA1Qk4oJ2HcMIIsDR+P8f4zwvg1llv9X6EnhCxnAS7FSLjajFWVEHi0VwLHWa73Uxg2X3dyBA4GagIVsjZE4qPecNlh40=" // ASTRAPAY_EMONEY_PRIVATE_KEY

	privPKCS1 = "RSA PRIVATE KEY"
	privPKCS8 = "PRIVATE KEY"
)

func main() {
	timestamp := time.Now().Format(time.RFC3339)
	client := newHTTPCient()

	// Step 1: Generate Asymmetric Signature for Access Token Request

	signature, err := generateAsymSign(fmt.Sprintf("%s|%s", clientID, timestamp), partnerPrivateKey)
	if err != nil {
		log.Fatalf("Failed to generate signature: %v", err)
		os.Exit(1)
	}
	if signature == nil {
		log.Fatalf("Signature is nil")
		os.Exit(1)
	}

	url := baseURL + "/snap/v1.0/access-token/b2b"
	method := "POST"

	payload := strings.NewReader(`{
    "grantType": "client_credentials"
}`)

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
		os.Exit(1)
	}
	req.Header.Add("X-CLIENT-KEY", clientID)
	req.Header.Add("X-TIMESTAMP", timestamp)
	req.Header.Add("X-SIGNATURE", *signature)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	var responseMap map[string]interface{}
	err = json.Unmarshal(body, &responseMap)
	if err != nil {
		log.Fatal("Failed to parse JSON response:", err)
		os.Exit(1)
	}
	accessToken, ok := responseMap["accessToken"].(string)
	if !ok {
		log.Fatalf("accessToken not found in response: %s", string(body))
		os.Exit(1)
	}
	getB2bTokenResponse, _ := json.MarshalIndent(responseMap, "", "  ")
	fmt.Println("===================== Get B2B Token Response =====================")
	fmt.Println(string(getB2bTokenResponse))
	fmt.Println("===================== Get B2B Token Response =====================")

	// Step 2: Generate Symmetric Signature for Get Merchant Info

	bodyReq := map[string]any{
		"merchantId": merchantID,
	}
	bodyReqStr, err := json.Marshal(bodyReq)
	if err != nil {
		log.Fatalf("Failed to marshal request body: %v", err)
		os.Exit(1)

	}

	// string minify body
	minify := func(s string) string {
		return strings.Join(strings.Fields(strings.TrimSpace(s)), "")
	}

	bodyReqHash, err := hashStr(minify(string(bodyReqStr)), crypto.SHA256)
	if err != nil {
		log.Fatalf("Failed to hash request body: %v", err)
		os.Exit(1)
	}
	strToSign := fmt.Sprintf("%s:%s:%s:%s:%s", "POST", "/snap/v1.0/merchant/info", accessToken, strings.ToLower(hex.EncodeToString(bodyReqHash)), timestamp)
	signature, err = generateSymSign(strToSign, clientSecret)
	if err != nil {
		log.Fatalf("Failed to generate signature: %v", err)
		os.Exit(1)
	}
	if signature == nil {
		log.Fatalf("Signature is nil")
		os.Exit(1)
	}

	url = baseURL + "/snap/v1.0/merchant/info"
	method = "POST"

	payload = strings.NewReader(string(bodyReqStr))

	req, err = http.NewRequest(method, url, payload)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
		os.Exit(1)
	}
	req.Header.Add("X-PARTNER-ID", clientID)
	req.Header.Add("X-DEVICE-ID", "12345")
	req.Header.Add("X-TIMESTAMP", timestamp)
	req.Header.Add("X-CHANNEL-MODEL", "iPhone 13")
	req.Header.Add("X-CHANNEL-OS", "iOS 16")
	req.Header.Add("X-EXTERNAL-ID", uuid.NewString())
	req.Header.Add("CHANNEL-ID", "00313")
	req.Header.Add("X-SIGNATURE", *signature)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	res, err = client.Do(req)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
		os.Exit(1)
	}
	var resp openapidto.MerchantInfoResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Fatalf("Failed to parse JSON response: %v", err)
		os.Exit(1)
	}
	respJSON, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println("===================== Get Merchant Info Response =====================")
	fmt.Println(string(respJSON))
	fmt.Println("===================== Get Merchant Info Response =====================")
}

func generateSymSign(stringToSign, signKey string) (*string, error) {
	// log.Debugf("stringToSign: %v, signKey: %v", stringToSign, signKey)
	hashed, err := hmacStr(signKey, stringToSign, crypto.SHA512)
	if err != nil {
		return nil, err
	}
	strEncoeded := base64.StdEncoding.EncodeToString(hashed)
	return &strEncoeded, nil
}

func generateAsymSign(stringToSign, signKey string) (*string, error) {
	rsaPrivateKey := getPrivateKey(signKey)
	hashed, err := hashStr(stringToSign, crypto.SHA256)
	if err != nil {
		return nil, err
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.SHA256, hashed)
	if err != nil {
		return nil, err
	}
	strEncoded := base64.StdEncoding.EncodeToString(signature)
	return &strEncoded, nil
}

func getPrivateKey(pKeyStr string) *rsa.PrivateKey {
	log.Debugf("Private Key: %s", pKeyStr)

	var err error
	if !strings.Contains(pKeyStr, "-----BEGIN ") {
		pKeyStr, err = convertBase64ToPEM(pKeyStr, privPKCS8)
		if err != nil {
			log.Fatalf("Failed to parse private key: %v", err)
			os.Exit(1)
		}
	}

	pemPrivateKey, err := parsePrivateKey(pKeyStr)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
		os.Exit(1)
	}
	return pemPrivateKey
}

func parsePrivateKey(key string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return nil, errors.New("[PRIVATE KEY] failed to decode PEM block")
	}
	//pem.Encode(os.Stdout, block)

	var (
		privateKey any
		err        error
	)

	switch block.Type {
	case privPKCS1:
		privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	case privPKCS8:
		privateKey, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	default:
		return nil, errors.New("[PRIVATE KEY] Invalid Pem block type : " + block.Type)
	}

	if err != nil {
		return nil, errors.New("[PRIVATE KEY] Failed to parse private key")
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("[PRIVATE KEY] Failed to type assertion private key")
	}

	return rsaPrivateKey, nil
}

func convertBase64ToPEM(base64Str, pemType string) (string, error) {
	der, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return "", fmt.Errorf("base64 decode failed: %w", err)
	}

	block := &pem.Block{
		Type:  pemType, // for example: "RSA PRIVATE KEY" or "PRIVATE KEY"
		Bytes: der,
	}

	pemBytes := pem.EncodeToMemory(block)
	return string(pemBytes), nil
}

func hashStr(stringToHash string, hashType crypto.Hash) ([]byte, error) {
	var mac hash.Hash
	switch hashType {
	case crypto.SHA256:
		mac = sha256.New()
	case crypto.SHA512:
		mac = sha512.New()
	default:
		return nil, errors.New("unsupported hash type")
	}

	if _, err := mac.Write([]byte(stringToHash)); err != nil {
		return nil, err
	}

	resultHash := mac.Sum(nil)
	return resultHash, nil
}

func hmacStr(clientSecret string, stringToSign string, hashType crypto.Hash) ([]byte, error) {
	var mac hash.Hash
	switch hashType {
	case crypto.SHA256:
		mac = hmac.New(sha256.New, []byte(clientSecret))
	case crypto.SHA512:
		mac = hmac.New(sha512.New, []byte(clientSecret))
	default:
		return nil, errors.New("unsupported hash type")
	}

	_, err := mac.Write([]byte(stringToSign))
	if err != nil {
		return nil, err
	}

	return mac.Sum(nil), nil
}

func newHTTPCient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 50,
			DialContext: (&net.Dialer{
				Timeout:   25 * time.Second,
				KeepAlive: 20 * time.Second,
			}).DialContext,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
