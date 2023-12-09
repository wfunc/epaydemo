# EPAY DEMO

## SHA256 online for check
https://emn178.github.io/online-tools/sha256.html

## epaydemo_test.go for example test
## should set MerchantID & AccessToken

## epaydemo.go for example
### about sign
* sign code
```
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
```
* verify sign
```
func VerifySign(AccessToken string, m xmap.M) bool {
	sign := m.Str("sign")
	calcSign := Sign(AccessToken, m)
	return strings.EqualFold(strings.ToLower(sign), strings.ToLower(calcSign))
}
```

## about api
### example for BindCardApply
```
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
	debugf("接口响应：%v", converter.JSON(data))
	return
}
```
