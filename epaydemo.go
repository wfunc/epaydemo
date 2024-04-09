package epaydemo

import (
	"crypto/sha256"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xhttp"
	"github.com/codingeasygo/util/xmap"
	"github.com/shopspring/decimal"
)

var (
	MerchantID                  = 0
	AccessToken                 = "xx"
	ApiURL                      = "https://example.com"
	PayCreateNotifyURL          = "https://example.com/notify/testNofity"
	ApplyWithdrawNotifyURL      = "https://example.com/notify/testNofity"
	ApplyOpenAcctNotifyURL      = "https://example.com/notify/testNofity"
	ApplyScanPayNotifyURL       = "https://example.com/notify/testNofity"
	ApplyScanWithdrawNotifyURL  = "https://example.com/notify/testNofity"
	HCApiSendMessageNotifyURL   = "https://example.com/notify/testNofity"
	HCTradeForCardNotifyURL     = "https://example.com/notify/testNofity"
	HCWithdrawalNotifyURL       = "https://example.com/notify/testNofity"
	HCPayOutNotifyURL           = "https://example.com/notify/testNofity"
	YeepayAddCustomerNotifyURL  = "https://example.com/notify/testNofity"
	ApplyAddCardVerifyNotifyURL = "https://example.com/notify/testNofity"
	ApplyBankTransferNotifyURL  = "https://example.com/notify/testNofity"
	UnifiedPaymentNotifyURL     = "https://example.com/notify/testNofity"
	WithdrawalToCardNotifyURL   = "https://example.com/notify/testNofity"
	EfpsBindCardNotifyURL       = "https://example.com/notify/testNofity"
	EfpsProtocolPayPreNotifyURL = "https://example.com/notify/testNofity"
	NoBindWKPaymentNotifyURL    = "https://example.com/notify/testNofity"

	Debug = false
)

var MarchineID = 1
var seqOrderID uint16
var lckOrderID = sync.RWMutex{}

func NewOrderID() (orderID string) {
	lckOrderID.Lock()
	defer lckOrderID.Unlock()
	seqOrderID++
	timeStr := time.Now().Format("20060102150405")
	return fmt.Sprintf("%v%02d%05d", timeStr, MarchineID, seqOrderID)
}

// SHA256 sign
func Sign(AccessToken string, m xmap.M) string {
	args := url.Values{}
	// format timestamp
	timestamp, err := decimal.NewFromString(m.Str("timestamp"))
	if err != nil {
		return ""
	}
	args.Set("merchant_id", m.Str("merchant_id"))
	args.Set("timestamp", timestamp.String())
	args.Set("method", m.Str("method"))

	signStr := fmt.Sprintf("%v&access_token=%v", args.Encode(), AccessToken)
	debugf("the string before sign：%s", signStr)
	h := sha256.New()
	h.Write([]byte(signStr))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// verify sign
func VerifySign(AccessToken string, m xmap.M) bool {
	sign := m.Str("sign")
	calcSign := Sign(AccessToken, m)
	return strings.EqualFold(strings.ToLower(sign), strings.ToLower(calcSign))
}

// create new params
func newParams(method string) xmap.M {
	return xmap.M{
		"merchant_id": MerchantID,
		"method":      method,
		"timestamp":   time.Now().UnixMilli(),
	}
}

type APIUserRegist struct {
	MerchantID            int64  `json:"merchant_id" valid:"merchant_id,r|i,r:0;"`
	Method                string `json:"method" valid:"method,o|s,O:userRegist;"`
	Timestamp             int64  `json:"timestamp" valid:"timestamp,r|i,r:0;"`
	Sign                  string `json:"sign" valid:"sign,o|s,l:0;"`
	OutOrderID            string `json:"out_order_id" valid:"out_order_id,r|s,l:0;"`
	UserName              string `json:"user_name" valid:"user_name,r|s,l:0;"`
	Phone                 string `json:"phone" valid:"phone,r|s,l:0;"`
	IDCardNo              string `json:"id_card_no" valid:"id_card_no,r|s,l:0;"`
	IDCardGrantDate       string `json:"id_card_grant_date" valid:"id_card_grant_date,r|s,l:0;"`
	IDCardExpire          string `json:"id_card_expire" valid:"id_card_expire,r|s,l:0;"`
	IDCardFront           string `json:"id_card_front" valid:"id_card_front,r|s,l:0;"`
	IDCardBack            string `json:"id_card_back" valid:"id_card_back,r|s,l:0;"`
	HandIdCard            string `json:"hand_id_card" valid:"hand_id_card,r|s,l:0;"`
	DebitCardAccount      string `json:"debit_card_account" valid:"debit_card_account,r|s,l:0;"`
	DebitCardAccountPhone string `json:"debit_card_account_phone" valid:"debit_card_account_phone,r|s,l:0;"`
	DebitCardBankName     string `json:"debit_card_bank_name" valid:"debit_card_bank_name,r|s,l:0;"`
	DebitCardBankCode     string `json:"debit_card_bank_code" valid:"debit_card_bank_code,r|s,l:0;"`
	BankUnitedCode        string `json:"bank_united_code" valid:"bank_united_code,r|s,l:0;"`
	BankSubName           string `json:"bank_sub_name" valid:"bank_sub_name,r|s,l:0;"`
	BankProvince          string `json:"bank_province" valid:"bank_province,r|s,l:0;"`
	BankCity              string `json:"bank_city" valid:"bank_city,r|s,l:0;"`
	NotifyURL             string `json:"notify_url" valid:"notify_url,r|s,l:0;"`
	SuccessTurnUrl        string `json:"success_turn_url" valid:"success_turn_url,r|s,l:0;"`
	ProvinceCode          string `json:"province_code" valid:"province_code,r|s,l:0;"`
	CityCode              string `json:"city_code" valid:"city_code,r|s,l:0;"`
	AreaCountyCode        string `json:"area_county_code" valid:"area_county_code,r|s,l:0;"`
	MerchantAddress       string `json:"merchant_address" valid:"merchant_address,r|s,l:0;"`
	Memo                  string `json:"memo" valid:"memo,o|s,l:0;"` //可选
}

func AlipayTradePreCreate(outOrderID, amount, subject, memo string) (resp xmap.M, err error) {
	method := "alipayTradePreCreate"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("amount", amount)
	p.SetValue("subject", subject)
	if len(memo) > 0 {
		p.SetValue("memo", memo)
	}
	resp, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(resp))
	return
}

func SandpayTransOrderCreate(outOrderID, amount, fromIpAddr, memo, subject string) (resp xmap.M, err error) {
	method := "sandpayTransOrderCreate"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("amount", amount)
	p.SetValue("from_ip_addr", fromIpAddr)
	if len(memo) > 0 {
		p.SetValue("memo", memo)
	}
	if len(subject) > 0 {
		p.SetValue("subject", subject)
	}
	resp, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(resp))
	return
}

func UserRegist(req APIUserRegist) (resp xmap.M, err error) {
	method := "userRegist"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	req.Sign = sign
	req.MerchantID = p.Int64("merchant_id")
	req.Timestamp = p.Int64("timestamp")
	req.Method = method
	resp, err = xhttp.PostJSONMap(req, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(resp))
	return
}

type APILivingIdentify struct {
	MerchantID int64  `json:"merchant_id" valid:"merchant_id,r|i,r:0;"`
	Method     string `json:"method" valid:"method,o|s,O:livingIdentify;"`
	Timestamp  int64  `json:"timestamp" valid:"timestamp,r|i,r:0;"`
	Sign       string `json:"sign" valid:"sign,o|s,l:0;"`
	OutOrderID string `json:"out_order_id" valid:"out_order_id,r|s,l:0;"`
}

func LivingIdentify(req APILivingIdentify) (resp xmap.M, err error) {
	method := "livingIdentify"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	req.Sign = sign
	req.MerchantID = p.Int64("merchant_id")
	req.Timestamp = p.Int64("timestamp")
	req.Method = method
	resp, err = xhttp.PostJSONMap(req, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(resp))
	return
}

type APIQueryLiving struct {
	MerchantID int64  `json:"merchant_id" valid:"merchant_id,r|i,r:0;"`
	Method     string `json:"method" valid:"method,o|s,O:queryLiving;"`
	Timestamp  int64  `json:"timestamp" valid:"timestamp,r|i,r:0;"`
	Sign       string `json:"sign" valid:"sign,o|s,l:0;"`
	OutOrderID string `json:"out_order_id" valid:"out_order_id,r|s,l:0;"`
}

func QueryLiving(req APIQueryLiving) (resp xmap.M, err error) {
	method := "queryLiving"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	req.Sign = sign
	req.MerchantID = p.Int64("merchant_id")
	req.Timestamp = p.Int64("timestamp")
	req.Method = method
	resp, err = xhttp.PostJSONMap(req, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(resp))
	return
}

type APISendWKPayment struct {
	MerchantID         int64  `json:"merchant_id" valid:"merchant_id,r|i,r:0;"`
	Method             string `json:"method" valid:"method,o|s,O:sendWKPayment;"`
	Timestamp          int64  `json:"timestamp" valid:"timestamp,o|i,r:0;"`
	Sign               string `json:"sign" valid:"sign,o|s,l:0;"`
	CustomerID         int64  `json:"customer_id" valid:"customer_id,o|i,r:0;"`
	CustomerOutOrderID string `json:"customer_out_order_id" valid:"customer_out_order_id,o|s,l:0;"` //可选
	OutOrderID         string `json:"out_order_id" valid:"out_order_id,r|s,l:0;"`
	Amount             string `json:"amount" valid:"amount,r|s,l:0;"`
	NotifyURL          string `json:"notify_url" valid:"notify_url,o|s,l:0;"`
	FromIPAddr         string `json:"from_ip_addr" valid:"from_ip_addr,o|s,l:0;"`
	Memo               string `json:"memo" valid:"memo,o|s,l:0;"` //可选
	GoodsTitle         string `json:"goods_title" valid:"goods_title,o|s,l:0;"`
	GoodsDescription   string `json:"goods_description" valid:"goods_description,o|s,l:0;"`
	SuccessTurnUrl     string `json:"success_turn_url" valid:"success_turn_url,o|s,l:0;"`
	Rate               string `json:"rate" valid:"rate,r|s,l:0;"`
	Pro                string `json:"pro" valid:"pro,r|s,l:0;"`
}

func SendWKPayment(req APISendWKPayment) (resp xmap.M, err error) {
	method := "sendWKPayment"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	req.Sign = sign
	req.MerchantID = p.Int64("merchant_id")
	req.Timestamp = p.Int64("timestamp")
	req.Method = method
	resp, err = xhttp.PostJSONMap(req, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(resp))
	return
}

// Adapay
func BindCardApply(outOrderID, cardName, certID, cardNo, cardPhone, bankName string) (data xmap.M, err error) {
	method := "bindCardApply"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("card_name", cardName)
	p.SetValue("cert_id", certID)
	p.SetValue("card_no", cardNo)
	p.SetValue("card_phone", cardPhone)
	p.SetValue("bank_card_type", "debit")
	p.SetValue("bank_name", bankName)
	p.SetValue("notify_url", ApplyAddCardVerifyNotifyURL)
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func BindCardConfirm(outOrderID, code string, customerID int64) (data xmap.M, err error) {
	method := "bindCardConfirm"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	p.SetValue("sign", sign)
	p.SetValue("code", code)
	if len(outOrderID) > 0 {
		p.SetValue("out_order_id", outOrderID)
	}
	if customerID > 0 {
		p.SetValue("customer_id", customerID)
	}
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func QuickPayApply(outOrderID, amount, fromIPAddr, memo, customerOutOrderID string, customerID int64) (data xmap.M, err error) {
	method := "quickPayApply"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("amount", amount)
	p.SetValue("from_ip_addr", fromIPAddr)
	if len(memo) > 0 {
		p.SetValue("memo", memo)
	}
	if len(customerOutOrderID) > 0 {
		p.SetValue("customer_out_order_id", customerOutOrderID)
	}
	if customerID > 0 {
		p.SetValue("customer_id", customerID)
	}
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func QuickPayConfirm(outOrderID, code string) (data xmap.M, err error) {
	method := "quickPayConfirm"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("code", code)
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func Withdraw(outOrdeID, cardNo, name, bankName, amount, customerOutOrderID string) (data xmap.M, err error) {
	method := "withdraw"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrdeID)
	p.SetValue("amount", amount)
	if len(customerOutOrderID) > 0 {
		p.SetValue("customer_out_order_id", customerOutOrderID)
	} else {
		p.SetValue("card_no", cardNo)
		p.SetValue("name", name)
		p.SetValue("bank_name", bankName)
	}
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	return
}

func SandpayTradeCreate(outOrderID, account, amount string) (data xmap.M, err error) {
	method := "sandpayTradeCreate"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("notify_url", ApplyScanPayNotifyURL)
	p.SetValue("amount", amount)
	p.SetValue("account", account)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func SandpayTradePreCreate(outOrderID, account, amount string) (data xmap.M, err error) {
	method := "sandpayTradePreCreate"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("notify_url", ApplyScanPayNotifyURL)
	p.SetValue("amount", amount)
	p.SetValue("account", account)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func SandpayAgentPayCreate(outOrderID, name, cardNo, amount string) (data xmap.M, err error) {
	method := "sandpayAgentPayCreate"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("name", name)
	p.SetValue("card_no", cardNo)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("notify_url", ApplyScanWithdrawNotifyURL)
	p.SetValue("amount", amount)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：", converter.JSON(data))
	return
}

func QueryCard(outOrderID string) (data xmap.M, err error) {
	method := "queryCard"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func Upload(filePath string) (data xmap.M, err error) {
	method := "upload"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	p.SetValue("sign", sign)
	args := url.Values{}
	for k, v := range p {
		args.Set(k, converter.String(v))
	}
	debugf("the submit GET data：%v", args.Encode())
	return xhttp.UploadMap(nil, "file", filePath, ApiURL+"/easyapi/"+method+"?%v", args.Encode())
}

func YeepayAddCustomer(account, name, idNumber, idFrontPath, idBackPath, cardNo, phone, bankCode string) (data xmap.M, err error) {
	method := "yeepayAddCustomer"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	p.SetValue("name", name)
	p.SetValue("id_number", idNumber)
	p.SetValue("phone", phone)
	p.SetValue("bank_code", bankCode)
	p.SetValue("notify_url", YeepayAddCustomerNotifyURL)
	p.SetValue("id_front_path", idFrontPath)
	p.SetValue("id_back_path", idBackPath)
	p.SetValue("card_no", cardNo)

	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

// gateway pay
func ZfGatewayPay(outOrderID, payType, amount, fromIPAddr string) (data xmap.M, err error) {
	method := "zfGatewayPay"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("pay_type", payType)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("notify_url", ApplyScanPayNotifyURL)
	p.SetValue("amount", amount)
	p.SetValue("from_ip_addr", fromIPAddr)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

// scan pay
func ApplyScanPay(outOrderID, amount, fromIPAddr string) (data xmap.M, err error) {
	method := "applyScanPay"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("pay_type", "alipay_qrcode")
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("notify_url", ApplyScanPayNotifyURL)
	p.SetValue("amount", amount)
	p.SetValue("from_ip_addr", fromIPAddr)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

// open account
func HCNetwork(account, name, mobile, idNo, idValidity, cardNo, bankName, bankBracnh string) (data xmap.M, err error) {
	method := "hcNetwork"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	p.SetValue("name", name)
	p.SetValue("mobile", mobile)
	p.SetValue("id_no", idNo)
	p.SetValue("id_validity", idValidity)
	p.SetValue("card_no", cardNo)
	p.SetValue("bank_name", bankName)
	p.SetValue("bank_branch", bankBracnh)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func HCQueryNetwork(account string) (data xmap.M, err error) {
	method := "hcQueryNetwork"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func HCSubMerBindCard(account, cardNo, cardPhone, bankName string) (data xmap.M, err error) {
	method := "hcSubMerBindCard"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	p.SetValue("card_no", cardNo)
	p.SetValue("card_phone", cardPhone)
	p.SetValue("bank_name", bankName)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	return
}

func HCApiSendMessage(account, outOrderID, cardNo, cardPhone, cvn2, expired string) (data xmap.M, err error) {
	method := "hcApiSendMessage"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("card_no", cardNo)
	p.SetValue("card_phone", cardPhone)
	p.SetValue("cvn2", cvn2)
	p.SetValue("expired", expired)
	p.SetValue("notify_url", HCApiSendMessageNotifyURL)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func HCApiVerifyCard(outOrderID, code string) (data xmap.M, err error) {
	method := "hcApiVerifyCard"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("code", code)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func HCTradeForCard(outOrderID, account, cardNo, cardPhone, amount, fee string) (data xmap.M, err error) {
	method := "hcTradeForCard"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("account", account)
	p.SetValue("card_no", cardNo)
	p.SetValue("card_phone", cardPhone)
	p.SetValue("amount", amount)
	p.SetValue("fee", fee)
	p.SetValue("notify_url", HCTradeForCardNotifyURL)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func HCWithdrawal(outOrderID, account, cardNo, cardPhone, amount string) (data xmap.M, err error) {
	method := "hcWithdrawal"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("account", account)
	p.SetValue("card_no", cardNo)
	p.SetValue("card_phone", cardPhone)
	p.SetValue("amount", amount)
	p.SetValue("notify_url", HCWithdrawalNotifyURL)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func HCRecycle(outOrderID, account, cardNo, cardPhone, amount string) (data xmap.M, err error) {
	method := "hcRecycle"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("account", account)
	p.SetValue("card_no", cardNo)
	p.SetValue("card_phone", cardPhone)
	p.SetValue("amount", amount)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func HCPayOut(outOrderID, account, cardNo, cardPhone, bankName, name, amount string) (data xmap.M, err error) {
	method := "hcPayOut"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("account", account)
	p.SetValue("card_no", cardNo)
	p.SetValue("card_phone", cardPhone)
	p.SetValue("bank_name", bankName)
	p.SetValue("name", name)
	p.SetValue("amount", amount)
	p.SetValue("notify_url", HCPayOutNotifyURL)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func HCQueryBalance(account string) (data xmap.M, err error) {
	method := "hcQueryBalance"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func ApplyMobilePay(orderID, payType, amount, fromIPAddr string) (data xmap.M, err error) {
	method := "applyMobilePay"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", orderID)
	p.SetValue("pay_type", payType)
	p.SetValue("notify_url", "https://example.com/notify/testNofity")
	p.SetValue("amount", amount)
	p.SetValue("from_ip_addr", fromIPAddr)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func PayCreate(orderID, account, phone, amount, fromIPAddr string) (data xmap.M, err error) {
	method := "payCreate"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", orderID)
	p.SetValue("account", account)
	p.SetValue("phone", phone)
	p.SetValue("notify_url", PayCreateNotifyURL)
	p.SetValue("amount", amount)
	p.SetValue("from_ip_addr", fromIPAddr)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

// payMethods --> AGRT_CREDIT_CARD,AGRT_DEBIT_CARD
func PaymentBankCard(orderID, account, phone, amount, fromIPAddr, linkedAcctno, linkedPhone, linkedAcctname, idNo, payMethods string) (data xmap.M, err error) {
	method := "paymentBankCard"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", orderID)
	p.SetValue("account", account)
	p.SetValue("phone", phone)
	p.SetValue("notify_url", PayCreateNotifyURL)
	p.SetValue("amount", amount)
	p.SetValue("from_ip_addr", fromIPAddr)
	p.SetValue("linked_acctno", linkedAcctno)
	p.SetValue("linked_phone", linkedPhone)
	p.SetValue("linked_acctname", linkedAcctname)
	p.SetValue("id_no", idNo)
	p.SetValue("pay_methods", payMethods)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func ApplyPhoneCode(account, phone string) (data xmap.M, err error) {
	method := "applyPhoneCode"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	p.SetValue("phone", phone)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func VerifyPhoneCode(account, phone, code string) (data xmap.M, err error) {
	method := "verifyPhoneCode"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	p.SetValue("phone", phone)
	p.SetValue("code", code)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func ApplyPasswordToken(account, phone, passwordScene string) (data xmap.M, err error) {
	method := "applyPasswordToken"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	p.SetValue("phone", phone)
	p.SetValue("password_scene", passwordScene)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func ApplyIndividual(account, phone, userName, idNo, idExp, address, occupation, cardNo, cardPhone string) (data xmap.M, err error) {
	method := "applyIndividual"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	p.SetValue("phone", phone)
	p.SetValue("user_name", userName)
	p.SetValue("id_no", idNo)
	p.SetValue("id_exp", idExp)
	p.SetValue("address", address)
	p.SetValue("occupation", occupation)
	p.SetValue("card_no", cardNo)
	p.SetValue("card_phone", cardPhone)
	p.SetValue("notify_url", ApplyOpenAcctNotifyURL)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func VerifyIndividual(account, phone, verifyCode, password, randomKey string) (data xmap.M, err error) {
	method := "verifyIndividual"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	p.SetValue("phone", phone)
	p.SetValue("password", password)
	p.SetValue("random_key", randomKey)
	p.SetValue("verify_code", verifyCode)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func QueryOrder(orderID string) (data xmap.M, err error) {
	method := "queryOrder"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("order_id", orderID)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func ApplyWithdraw(orderID, cardNo, acctName, bankName, amount string) (data xmap.M, err error) {
	method := "applyWithdraw"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)

	p.SetValue("sign", sign)
	p.SetValue("out_order_id", orderID)
	p.SetValue("notify_url", ApplyWithdrawNotifyURL)
	p.SetValue("amount", amount)
	p.SetValue("card_no", cardNo)
	p.SetValue("acct_name", acctName)
	p.SetValue("bank_name", bankName)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func ApplyOpenAcct(account, phone, returnURL string) (data xmap.M, err error) {
	method := "applyOpenAcct"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	p.SetValue("phone", phone)
	p.SetValue("return_url", returnURL)
	p.SetValue("notify_url", ApplyOpenAcctNotifyURL)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func QueryLinkedAcct(account string) (data xmap.M, err error) {
	method := "queryLinkedAcct"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func QueryAcctInfo(account string) (data xmap.M, err error) {
	method := "queryAcctInfo"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func QueryUserInfo(account string) (data xmap.M, err error) {
	method := "queryUserInfo"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func Withdrawal(account, amount, outOrderID, cardNo, agreeNo string, innerTransferDone int) (data xmap.M, err error) {
	method := "withdrawal"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	p.SetValue("amount", amount)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("linked_acctno", cardNo)
	p.SetValue("linked_agrtno", agreeNo)
	p.SetValue("notify_url", ApplyWithdrawNotifyURL)
	p.SetValue("inner_transfer_done", innerTransferDone)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func ValidationSMS(outOrderID, code string) (data xmap.M, err error) {
	method := "validationSMS"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("code", code)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func PreAddCard(outOrderID, account, notifyURL, returnURL, cardNo, cardPhone string) (data xmap.M, err error) {
	method := "preAddCard"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("source", "ACCP")
	p.SetValue("notify_url", notifyURL)
	p.SetValue("return_url", returnURL)
	if cardNo != "" {
		p.SetValue("card_no", cardNo)
	}
	if cardPhone != "" {
		p.SetValue("card_phone", cardPhone)
	}
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func ApplyAddCardVerify(orderID, cardNo, cardPhone, password, randomKey string) (data xmap.M, err error) {
	method := "applyAddCardVerify"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("order_id", orderID)
	p.SetValue("card_no", cardNo)
	p.SetValue("card_phone", cardPhone)
	p.SetValue("password", password)
	p.SetValue("random_key", randomKey)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func ApplyAddCardVerifyNew(outOrderID, account, cardNo, cardPhone, password, randomKey string) (data xmap.M, err error) {
	method := "applyAddCardVerify"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("account", account)
	p.SetValue("source", "ACCP")
	p.SetValue("card_no", cardNo)
	p.SetValue("card_phone", cardPhone)
	p.SetValue("password", password)
	p.SetValue("random_key", randomKey)
	p.SetValue("notify_url", ApplyAddCardVerifyNotifyURL)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func VerifyCard(orderID, code string) (data xmap.M, err error) {
	method := "verifyCard"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("order_id", orderID)
	p.SetValue("code", code)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

// 易宝相关接口
func ApplyBankTransfer(outOrderID, amount string) (data xmap.M, err error) {
	method := "applyBankTransfer"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("amount", amount)
	p.SetValue("notify_url", ApplyBankTransferNotifyURL)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func CashierUnifiedOrder(outOrderID, amount string) (data xmap.M, err error) {
	method := "cashierUnifiedOrder"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("amount", amount)
	p.SetValue("notify_url", ApplyBankTransferNotifyURL)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func UnifiedPayment(outOrderID, amount string) (data xmap.M, err error) {
	method := "unifiedPayment"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("amount", amount)
	p.SetValue("notify_url", UnifiedPaymentNotifyURL)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func WithdrawalToCard(outOrderID, account, cardNo, acctName, bankName, amount string, sync int) (data xmap.M, err error) {
	method := "withdrawalToCard"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("account", account)
	p.SetValue("card_no", cardNo)
	p.SetValue("acct_name", acctName)
	p.SetValue("bank_name", bankName)
	p.SetValue("amount", amount)
	p.SetValue("sync", sync)
	p.SetValue("notify_url", WithdrawalToCardNotifyURL)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func EfpsBindCard(account, name, idNumber, cardNo, phone, bankCardType string) (data xmap.M, err error) {
	method := "efpsBindCard"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	p.SetValue("name", name)
	p.SetValue("id_number", idNumber)
	p.SetValue("card_no", cardNo)
	p.SetValue("phone", phone)
	p.SetValue("bank_card_type", bankCardType)
	p.SetValue("notify_url", EfpsBindCardNotifyURL)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func EfpsBindCardConfirm(account, code string) (data xmap.M, err error) {
	method := "efpsBindCardConfirm"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	p.SetValue("code", code)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func EfpsProtocolPayPre(account, outOrderID, amount string) (data xmap.M, err error) {
	method := "efpsProtocolPayPre"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("amount", amount)
	p.SetValue("notify_url", EfpsProtocolPayPreNotifyURL)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func EfpsProtocolPayConfirm(orderID, code string) (data xmap.M, err error) {
	method := "efpsProtocolPayConfirm"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("order_id", orderID)
	p.SetValue("code", code)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func NoBindWKPayment(account, outOrderID, amount, creditCardNo, creditCardAccountName, creditCardPhone, creditCardCvv, creditCardExpire, creditCardBankName, debitCardNo, debitCardPhone, debitCardBankName, idNumber, rate, pro string) (data xmap.M, err error) {
	method := "noBindWKPayment"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("account", account)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("amount", amount)
	p.SetValue("credit_card_no", creditCardNo)
	p.SetValue("credit_card_account_name", creditCardAccountName)
	p.SetValue("credit_card_phone", creditCardPhone)
	p.SetValue("credit_card_cvv", creditCardCvv)
	p.SetValue("credit_card_expire", creditCardExpire)
	p.SetValue("credit_card_bank_name", creditCardBankName)
	p.SetValue("debit_card_no", debitCardNo)
	p.SetValue("debit_card_phone", debitCardPhone)
	p.SetValue("debit_card_bank_name", debitCardBankName)
	p.SetValue("id_number", idNumber)
	p.SetValue("rate", rate)
	p.SetValue("pro", pro)
	p.SetValue("notify_url", NoBindWKPaymentNotifyURL)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func NoBindWKMsgSubmit(outOrderID, code string) (data xmap.M, err error) {
	method := "noBindWKMsgSubmit"
	p := newParams(method)
	sign := Sign(AccessToken, p)
	debugf("the string after sign：%s", sign)
	p.SetValue("sign", sign)
	p.SetValue("out_order_id", outOrderID)
	p.SetValue("code", code)
	debugf("the submit JSON data：%v", converter.JSON(p))
	data, err = xhttp.PostJSONMap(p, ApiURL+"/easyapi/"+method)
	debugf("response：%v", converter.JSON(data))
	return
}

func debugf(template string, args ...interface{}) {
	if Debug {
		Infof(template, args...)
	}
}
