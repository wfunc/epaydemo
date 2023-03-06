package epaydemo

import (
	"crypto/sha256"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xhttp"
	"github.com/codingeasygo/util/xmap"
	"github.com/shopspring/decimal"
	"github.com/wfunc/go/xlog"
)

var (
	MerchantID                = 100000
	AccessToken               = "48c3cf04d96841986b073a1c98a45cec2e4adfa9"
	ApiURL                    = "https://example.com"
	PayCreateNotifyURL        = "https://example.com/notify/testNofity"
	ApplyWithdrawNotifyURL    = "https://example.com/notify/testNofity"
	ApplyOpenAcctNotifyURL    = "https://example.com/notify/testNofity"
	ApplyScanPayNotifyURL     = "https://example.com/notify/testNofity"
	HCApiSendMessageNotifyURL = "https://example.com/notify/testNofity"
	Debug                     = false
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

	signBefore := fmt.Sprintf("%v&access_token=%v", args.Encode(), AccessToken)
	debugf("签名前字符串：%s", signBefore)
	h := sha256.New()
	h.Write([]byte(signBefore))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 汇潮扫码支付
func ApplyScanPay(outOrderID, amount, fromIPAddr string) (data xmap.M, err error) {
	method := "applyScanPay"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)
	body.SetValue("sign", sign)
	body.SetValue("pay_type", "alipay_qrcode")
	body.SetValue("out_order_id", outOrderID)
	body.SetValue("notify_url", ApplyScanPayNotifyURL)
	body.SetValue("amount", amount)
	body.SetValue("from_ip_addr", fromIPAddr)
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}

// 汇潮用户进件
func HCNetwork(account, name, mobile, idNo, idValidity, cardNo string) (data xmap.M, err error) {
	method := "hcNetwork"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)
	body.SetValue("sign", sign)
	body.SetValue("account", account)
	body.SetValue("name", name)
	body.SetValue("mobile", mobile)
	body.SetValue("id_no", idNo)
	body.SetValue("id_validity", idValidity)
	body.SetValue("card_no", cardNo)
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}

func HCQueryNetwork(account string) (data xmap.M, err error) {
	method := "hcQueryNetwork"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)
	body.SetValue("sign", sign)
	body.SetValue("account", account)
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}

func HCApiSendMessage(account, outOrderID, cardNo, cardPhone, cvn2, expired string) (data xmap.M, err error) {
	method := "hcApiSendMessage"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)
	body.SetValue("sign", sign)
	body.SetValue("account", account)
	body.SetValue("out_order_id", outOrderID)
	body.SetValue("card_no", cardNo)
	body.SetValue("card_phone", cardPhone)
	body.SetValue("cvn2", cvn2)
	body.SetValue("expired", expired)
	body.SetValue("notify_url", HCApiSendMessageNotifyURL)
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}

func HCApiVerifyCard(outOrderID, code string) (data xmap.M, err error) {
	method := "hcApiVerifyCard"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)
	body.SetValue("sign", sign)
	body.SetValue("out_order_id", outOrderID)
	body.SetValue("code", code)
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}

func ApplyMobilePay(orderID, payType, amount, fromIPAddr string) (data xmap.M, err error) {
	method := "applyMobilePay"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)
	body.SetValue("sign", sign)
	body.SetValue("out_order_id", orderID)
	body.SetValue("pay_type", payType)
	body.SetValue("notify_url", "https://example.com/notify/testNofity")
	body.SetValue("amount", amount)
	body.SetValue("from_ip_addr", fromIPAddr)
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}

func PayCreate(orderID, account, phone, amount, fromIPAddr string) (data xmap.M, err error) {
	method := "payCreate"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)
	body.SetValue("sign", sign)
	body.SetValue("out_order_id", orderID)
	body.SetValue("account", account)
	body.SetValue("phone", phone)
	body.SetValue("notify_url", PayCreateNotifyURL)
	body.SetValue("amount", amount)
	body.SetValue("from_ip_addr", fromIPAddr)
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}

func ApplyPhoneCode(account, phone string) (data xmap.M, err error) {
	method := "applyPhoneCode"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)
	body.SetValue("sign", sign)
	body.SetValue("account", account)
	body.SetValue("phone", phone)
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}

func VerifyPhoneCode(account, phone, code string) (data xmap.M, err error) {
	method := "verifyPhoneCode"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)
	body.SetValue("sign", sign)
	body.SetValue("account", account)
	body.SetValue("phone", phone)
	body.SetValue("code", code)
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}

func ApplyPasswordToken(account, phone, passwordScene string) (data xmap.M, err error) {
	method := "applyPasswordToken"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)
	body.SetValue("sign", sign)
	body.SetValue("account", account)
	body.SetValue("phone", phone)
	body.SetValue("password_scene", passwordScene)
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}

func ApplyIndividual(account, phone, userName, idNo, idExp, address, occupation, cardNo, cardPhone string) (data xmap.M, err error) {
	method := "applyIndividual"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)
	body.SetValue("sign", sign)
	body.SetValue("account", account)
	body.SetValue("phone", phone)
	body.SetValue("user_name", userName)
	body.SetValue("id_no", idNo)
	body.SetValue("id_exp", idExp)
	body.SetValue("address", address)
	body.SetValue("occupation", occupation)
	body.SetValue("card_no", cardNo)
	body.SetValue("card_phone", cardPhone)
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}

func VerifyIndividual(account, phone, verifyCode, password, randomKey string) (data xmap.M, err error) {
	method := "verifyIndividual"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)
	body.SetValue("sign", sign)
	body.SetValue("account", account)
	body.SetValue("phone", phone)
	body.SetValue("password", password)
	body.SetValue("random_key", randomKey)
	body.SetValue("verify_code", verifyCode)
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}

func QueryOrder(orderID string) (data xmap.M, err error) {
	method := "queryOrder"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)
	body.SetValue("sign", sign)
	body.SetValue("order_id", orderID)
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}

func ApplyWithdraw(orderID, cardNo, acctName, bankName, amount string) (data xmap.M, err error) {
	method := "applyWithdraw"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)

	body.SetValue("sign", sign)
	body.SetValue("out_order_id", orderID)
	body.SetValue("notify_url", ApplyWithdrawNotifyURL)
	body.SetValue("amount", amount)
	body.SetValue("card_no", cardNo)
	body.SetValue("acct_name", acctName)
	body.SetValue("bank_name", bankName)
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}

func ApplyOpenAcct(account, phone, returnURL string) (data xmap.M, err error) {
	method := "applyOpenAcct"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)
	body.SetValue("sign", sign)
	body.SetValue("account", account)
	body.SetValue("phone", phone)
	body.SetValue("return_url", returnURL)
	body.SetValue("notify_url", ApplyOpenAcctNotifyURL)
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}

func QueryLinkedAcct(account string) (data xmap.M, err error) {
	method := "queryLinkedAcct"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)
	body.SetValue("sign", sign)
	body.SetValue("account", account)
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}
func Withdrawal(account, amount, outOrderID, cardNo, agreeNo string) (data xmap.M, err error) {
	method := "withdrawal"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)
	body.SetValue("sign", sign)
	body.SetValue("account", account)
	body.SetValue("amount", amount)
	body.SetValue("out_order_id", outOrderID)
	body.SetValue("linked_acctno", cardNo)
	body.SetValue("linked_agrtno", agreeNo)
	body.SetValue("notify_url", ApplyWithdrawNotifyURL)
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}

func ValidationSMS(outOrderID, code string) (data xmap.M, err error) {
	method := "validationSMS"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)
	body.SetValue("sign", sign)
	body.SetValue("out_order_id", outOrderID)
	body.SetValue("code", code)
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}

func PreAddCard(outOrderID, account, notifyURL, returnURL, cardNo, cardPhone string) (data xmap.M, err error) {
	method := "preAddCard"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)
	body.SetValue("sign", sign)
	body.SetValue("account", account)
	body.SetValue("out_order_id", outOrderID)
	body.SetValue("source", "ACCP")
	body.SetValue("notify_url", notifyURL)
	body.SetValue("return_url", returnURL)
	if cardNo != "" {
		body.SetValue("card_no", cardNo)
	}
	if cardPhone != "" {
		body.SetValue("card_phone", cardPhone)
	}
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}

func ApplyAddCardVerify(orderID, cardNo, cardPhone, password, randomKey string) (data xmap.M, err error) {
	method := "applyAddCardVerify"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)
	body.SetValue("sign", sign)
	body.SetValue("order_id", orderID)
	body.SetValue("card_no", cardNo)
	body.SetValue("password", password)
	body.SetValue("random_key", randomKey)
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}

func VerifyCard(orderID, code string) (data xmap.M, err error) {
	method := "verifyCard"
	body := newBody(method)
	sign := Sign(AccessToken, body)
	debugf("签名后字符串：%s", sign)
	body.SetValue("sign", sign)
	body.SetValue("order_id", orderID)
	body.SetValue("code", code)
	debugf("提交的JSON数据：%v", converter.JSON(body))
	data, err = xhttp.PostJSONMap(body, ApiURL+"/easyapi/"+method)
	debugf("接口响应：%v", converter.JSON(data))
	return
}

func newBody(method string) xmap.M {
	return xmap.M{
		"merchant_id": MerchantID,
		"method":      method,
		"timestamp":   time.Now().UnixMilli(),
	}
}

func debugf(template string, args ...interface{}) {
	if Debug {
		xlog.Infof(template, args...)
	}
}
