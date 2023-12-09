package epaydemo

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xhttp"
)

func init() {
	Debug = true
}

func TestBindCardApply(t *testing.T) {
	outOrderID := time.Now().Format("20060102150405")
	bindCardApply, err := BindCardApply(outOrderID, "张三", "445xxx", "6214xxx", "182xx", "xx银行")
	if err != nil {
		t.Errorf("TestBindCardApply fail with %v", err)
		return
	}
	fmt.Println(converter.JSON(bindCardApply))
}

func TestBindCardConfirm(t *testing.T) {
	outOrderID := ""
	customerID := int64(0)
	code := ""
	bindCardConfirm, err := BindCardConfirm(outOrderID, code, customerID)
	if err != nil {
		t.Errorf("TestBindCardConfirm fail with %v", err)
		return
	}
	fmt.Println(converter.JSON(bindCardConfirm))
}

func TestQuickPayApply(t *testing.T) {
	outOrderID := time.Now().Format("20060102150405")
	customerOutOrderID := ""
	customerID := int64(0)
	quickPayApply, err := QuickPayApply(outOrderID, "1", "127.0.0.1", "test", customerOutOrderID, customerID)
	if err != nil {
		t.Errorf("TestQuickPayApply fail with %v", err)
		return
	}
	fmt.Println(converter.JSON(quickPayApply))
}

func TestQuickPayConfirm(t *testing.T) {
	outOrderID := ""
	code := ""
	quickPayConfirm, err := QuickPayConfirm(outOrderID, code)
	if err != nil {
		t.Errorf("TestQuickPayConfirm fail with %v", err)
		return
	}
	fmt.Println(converter.JSON(quickPayConfirm))
}

func TestWithdraw(t *testing.T) {
	outOrderID := time.Now().Format("20060102150405")
	customerOutOrderID := ""
	withdraw, err := Withdraw(outOrderID, "", "", "", "1", customerOutOrderID)
	if err != nil {
		t.Errorf("TestWithdraw fail with %v", err)
		return
	}
	fmt.Println(converter.JSON(withdraw))
}

func TestNoBindWKPayment(t *testing.T) {
	orderID := NewOrderID()

	noBindWKPayment, err := NoBindWKPayment("test", orderID, "10", "6258xx", "张三", "182xx", "xxx", "xxxx", "xx银行", "62148xx", "182xx", "xx银行", "445xxx", "0.0045", "0.5")
	if err != nil {
		t.Errorf("TestNoBindWKPayment fail with %v", err)
		return
	}
	fmt.Println(converter.JSON(noBindWKPayment))
}

func TestSandpayTradeCrate(t *testing.T) {
	orderID := NewOrderID()
	sandpayTradeCrate, err := SandpayTradeCreate(orderID, "1", "1")
	if err != nil {
		t.Errorf("TestSandpayTradeCrate fail with %v", err)
		return
	}
	fmt.Println(converter.JSON(sandpayTradeCrate))
}

func TestSandpayTradePreCreate(t *testing.T) {
	proxyStr := "http://127.0.0.1:1105"
	proxyURL, _ := url.Parse(proxyStr)
	xhttp.DefaultClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
	orderID := NewOrderID()
	sandpayTradeCrate, err := SandpayTradePreCreate(orderID, "1", "1")
	if err != nil {
		t.Errorf("TestSandpayTradePreCreate fail with %v", err)
		return
	}
	fmt.Println(converter.JSON(sandpayTradeCrate))
}

func TestSandpayAgentPayCreate(t *testing.T) {
	orderID := NewOrderID()
	sandpayAgentPayCreate, err := SandpayAgentPayCreate(orderID, "1", "1", "1")
	if err != nil {
		t.Errorf("TestSandpayAgentPayCreate fail with %v", err)
		return
	}
	fmt.Println(converter.JSON(sandpayAgentPayCreate))
}

//

func TestEfpsBindCard(t *testing.T) {
	efpsBindCard, err := EfpsBindCard("test", "张三", "230xxx", "6212xx", "156xx", "debit")
	if err != nil {
		t.Errorf("TestEfpsBindCard fail with %v", err)
		return
	}
	fmt.Println(converter.JSON(efpsBindCard))
}

func TestEfpsBindCardConfirm(t *testing.T) {
	efpsBindCardConfirm, err := EfpsBindCardConfirm("test", "737083")
	if err != nil {
		t.Errorf("TestEfpsBindCardConfirm fail with %v", err)
		return
	}
	fmt.Println(converter.JSON(efpsBindCardConfirm))
}

func TestEfpsProtocolPayPre(t *testing.T) {
	orderID := NewOrderID()
	efpsProtocolPayPre, err := EfpsProtocolPayPre("test", orderID, "1")
	if err != nil {
		t.Errorf("TestEfpsProtocolPayPre fail with %v", err)
		return
	}
	fmt.Println(orderID, converter.JSON(efpsProtocolPayPre))
}

func TestEfpsProtocolPayConfirm(t *testing.T) {
	efpsProtocolPayConfirm, err := EfpsProtocolPayConfirm("202307061550550100001", "")
	if err != nil {
		t.Errorf("TestEfpsProtocolPayConfirm fail with %v", err)
		return
	}
	fmt.Println(converter.JSON(efpsProtocolPayConfirm))
}

func TestUnifiedPayment(t *testing.T) {
	orderID := NewOrderID()
	unifiedPayment, err := UnifiedPayment(orderID, "12")
	if err != nil {
		t.Errorf("TestUnifiedPayment fail with %v", err)
		return
	}
	fmt.Println(converter.JSON(unifiedPayment))
}

func TestWithdrawalToCard(t *testing.T) {
	orderID := NewOrderID()
	withdrawalToCard, err := WithdrawalToCard(orderID, "101744xx", "", "", "", "4", 2)
	if err != nil {
		t.Errorf("TestWithdrawalToCard fail with %v", err)
		return
	}
	fmt.Println(converter.JSON(withdrawalToCard))
}

func TestQueryCard(t *testing.T) {
	queryCard, err := QueryCard("202304181846190004")
	if err != nil {
		t.Errorf("TestQueryCard fail with %v", err)
		return
	}
	fmt.Println(converter.JSON(queryCard))
}

func TestUpload(t *testing.T) {
	upload, err := Upload("xx.jpg")
	if err != nil {
		t.Errorf("TestUpload fail with %v", err)
		return
	}
	fmt.Println(converter.JSON(upload))

}

func TestYeepayAddCustomer(t *testing.T) {
	yeepayAddCustomer, err := YeepayAddCustomer("yepaytest001", "张三", "2323011x", "/upload/2023-04-07/642f847b4f6f810007000004.jpg", "/upload/2023-04-07/642f8ae44f6f810007000006.jpg", "622848125x", "1x9x68", "ABC")
	if err != nil {
		t.Errorf("yeepayAddCustomer fail with %v", err)
		return
	}
	fmt.Println(converter.JSON(yeepayAddCustomer))

}

// 支付宝扫码支付
func TestApplyScanPay(t *testing.T) {
	MerchantID = 100001
	orderID := NewOrderID()
	applyScanPay, err := ApplyScanPay(orderID, "0.1", "127.0.0.1")
	if err != nil {
		t.Errorf("TestApplyScanPay fail with %v", err)
		return
	}
	fmt.Println(converter.JSON(applyScanPay))
}

// 智付微信支付宝网银支付
func TestZfGatewayPay(t *testing.T) {
	orderID := NewOrderID()
	zfGatewayPay, err := ZfGatewayPay(orderID, "direct_pay", "9.95", "127.0.0.1")
	if err != nil {
		t.Errorf("TestZfGatewayPay fail with %v", err)
		return
	}
	fmt.Println(converter.JSON(zfGatewayPay))
}

// 汇潮用户进件
func TestHCNetwork(t *testing.T) {
	hcNetwork, err := HCNetwork("test1", "123", "123", "123", "123", "123", "123", "123")
	if err != nil {
		t.Errorf("TestHCNetwork fail with %v", err)
		return
	}
	fmt.Printf("TestHCNetwork data is %v\n", converter.JSON(hcNetwork))
}

func TestHCQueryNetwork(t *testing.T) {
	hcNetwork, err := HCQueryNetwork("test1")
	if err != nil {
		t.Errorf("TestHCQueryNetwork fail with %v", err)
		return
	}
	fmt.Printf("TestHCQueryNetwork data is %v\n", converter.JSON(hcNetwork))
}

func TestHCApiSendMessage(t *testing.T) {
	orderID := NewOrderID()
	hcApiSendMessage, err := HCApiSendMessage("test1", orderID, "123", "123", "123", "123")
	if err != nil {
		t.Errorf("TestHCApiSendMessage fail with %v", err)
		return
	}
	fmt.Printf("TestHCApiSendMessage data is %v\n", converter.JSON(hcApiSendMessage))
}

func TestHCApiVerifyCard(t *testing.T) {
	hcApiVerifyCard, err := HCApiVerifyCard("test1", "123")
	if err != nil {
		t.Errorf("TestHCApiVerifyCard fail with %v", err)
		return
	}
	fmt.Printf("TestHCApiVerifyCard data is %v\n", converter.JSON(hcApiVerifyCard))
}

func TestHCTradeForCard(t *testing.T) {
	orderID := NewOrderID()
	hcTradeForCard, err := HCTradeForCard(orderID, "test", "6214xxx", "13800138000", "1000", "0")
	if err != nil {
		t.Errorf("TestHCTradeForCard fail with %v", err)
		return
	}
	fmt.Printf("TestHCTradeForCard data is %v\n", converter.JSON(hcTradeForCard))
}

func TestHCWithdrawal(t *testing.T) {
	orderID := NewOrderID()
	hcWithdrawal, err := HCWithdrawal(orderID, "test", "6214xxx", "13800138000", "1000")
	if err != nil {
		t.Errorf("TestHCWithdrawal fail with %v", err)
		return
	}
	fmt.Printf("TestHCWithdrawal data is %v\n", converter.JSON(hcWithdrawal))
}

func TestHCRecycle(t *testing.T) {
	// HCRecycle
	orderID := NewOrderID()
	hcRecycle, err := HCRecycle(orderID, "xx", "6226xxx", "138xx", "")
	if err != nil {
		t.Errorf("TestHCRecycle fail with %v", err)
		return
	}
	fmt.Printf("TestHCRecycle data is %v\n", converter.JSON(hcRecycle))
}

func TestHCPayOut(t *testing.T) {
	orderID := NewOrderID()
	hcPayOut, err := HCPayOut(orderID, "test", "6214xxx", "138xx", "xx银行", "张三", "11")
	if err != nil {
		t.Errorf("TestHCPayOut fail with %v", err)
		return
	}
	fmt.Printf("TestHCPayOut data is %v\n", converter.JSON(hcPayOut))
}

// HCPayOut

func TestHCQueryBalance(t *testing.T) {
	hcQueryBalance, err := HCQueryBalance("test1")
	if err != nil {
		t.Errorf("TestHCQueryBalance fail with %v", err)
		return
	}
	fmt.Printf("TestHCQueryBalance data is %v\n", converter.JSON(hcQueryBalance))
}

func TestHCSubMerBindCard(t *testing.T) {
	hcSubMerBindCard, err := HCSubMerBindCard("test", "test", "test", "xx银行")
	if err != nil {
		t.Errorf("TestHCQueryBalance fail with %v", err)
		return
	}
	fmt.Printf("TestHCQueryBalance data is %v\n", converter.JSON(hcSubMerBindCard))
}

// 查询订单接口
func TestQueryOrder(t *testing.T) {
	_, err := QueryOrder("20230xx")
	if err != nil {
		t.Errorf("TestQueryOrder fail with %v", err)
		return
	}
}

// 微信支付宝下单接口
func TestApplyMobilePay(t *testing.T) {
	orderID := NewOrderID()
	_, err := ApplyMobilePay(orderID, "alipay_qrcode", "0.01", "127.0.0.1")
	if err != nil {
		t.Errorf("TestApplyMobilePay fail with %v", err)
		return
	}
}

// 收银台支付接口
func TestPayCreate(t *testing.T) {
	orderID := NewOrderID()
	_, err := PayCreate(orderID, "1000102", "12312341234", "1", "127.0.0.1")
	if err != nil {
		t.Errorf("TestPayCreate fail with %v", err)
		return
	}
}

// 银行卡快捷
func TestPaymentBankCard(t *testing.T) {
	orderID := NewOrderID()
	_, err := PaymentBankCard(orderID, "1000102", "", "1", "127.0.0.1", "6214", "12345", "张三", "44222", "AGRT_DEBIT_CARD")
	if err != nil {
		t.Errorf("TestPaymentBankCard fail with %v", err)
		return
	}
}

func TestQueryLinkedAcct(t *testing.T) {
	_, err := QueryLinkedAcct("1000102")
	if err != nil {
		t.Errorf("TestQueryLinkedAcct fail with %v", err)
		return
	}
}

func TestQueryUserInfo(t *testing.T) {
	queryUserInfo, err := QueryUserInfo("63155")
	if err != nil {
		t.Errorf("TestQueryUserInfo fail with %v", err)
		return
	}
	fmt.Println(converter.JSON(queryUserInfo))
}

// 申请验证码接口
func TestApplyPhoneCode(t *testing.T) {
	phone := "test" // 需要填写正确的手机号
	_, err := ApplyPhoneCode(phone, phone)
	if err != nil {
		t.Errorf("TestApplyPhoneCode fail with %v", err)
		return
	}
}

// 验证验证码接口
func TestVerifyPhoneCode(t *testing.T) {
	account, phone := "test", "test"
	_, err := VerifyPhoneCode(account, phone, "000000")
	if err != nil {
		t.Errorf("TestVerifyPhoneCode fail with %v", err)
		return
	}
}

// 申请密码控件Token
func TestApplyPasswordToken(t *testing.T) {
	_, err := ApplyPasswordToken("test", "test", "setting_password")
	if err != nil {
		t.Errorf("TestApplyPasswordToken fail with %v", err)
		return
	}
}

func TestApplyIndividual(t *testing.T) {
	_, err := ApplyIndividual("test", "test", "张三", "123", "20500822", "xx省xx市", "19", "", "")
	if err != nil {
		t.Errorf("TestApplyIndividual fail with %v", err)
		return
	}
}

func TestVerifyIndividual(t *testing.T) {
	_, err := VerifyIndividual("test", "test", "", "123456", "123456")
	if err != nil {
		t.Errorf("TestVerifyIndividual fail with %v", err)
		return
	}
}

// 代付下单接口
func TestApplyWithdraw(t *testing.T) {
	_, err := ApplyWithdraw(NewOrderID(), "123", "123", "123", "0.01")
	if err != nil {
		t.Errorf("TestApplyWithdraw fail with %v", err)
		return
	}
}

func TestWithdrawalSecond(t *testing.T) {
	orderID := "xx"
	data, err := Withdrawal("70457", "490", orderID, "6217xx", "", 1)
	if err != nil {
		t.Errorf("TestWithdrawal fail with %v", err)
		return
	}
	fmt.Printf("TestWithdrawal[%v] data is %v", orderID, data)
}

// 请求开户页面
func TestApplyOpenAcct(t *testing.T) {
	data, err := ApplyOpenAcct("test", "138xx", "https://www.google.com")
	if err != nil {
		t.Errorf("TestApplyOpenAcct fail with %v", err)
		return
	}
	fmt.Printf("TestApplyOpenAcct data is %v", data)
}

func TestWithdrawal(t *testing.T) {
	Debug = true
	orderID := NewOrderID()
	data, err := Withdrawal("test", "1.2", orderID, "6217xxx", "", 0)
	if err != nil {
		t.Errorf("TestWithdrawal fail with %v", err)
		return
	}
	fmt.Printf("TestWithdrawal[%v] data is %v", orderID, data)
}

//

func TestValidationSMS(t *testing.T) {
	validationSMS, err := ValidationSMS("2023050xxx", "xxx")
	fmt.Println(validationSMS, err)
}

func TestQueryAcctInfo(t *testing.T) {
	_, err := QueryAcctInfo("test")
	if err != nil {
		t.Errorf("TestQueryAcctInfo fail with %v", err)
		return
	}
	// fmt.Println(converter.JSON(queryAcctInfo))
}

func TestPreAddCard(t *testing.T) {
	orderID := NewOrderID()
	data, err := PreAddCard(orderID, "test", "https://example.com/notify/testNofity", "https://example.com/notify/testNofity", "", "")
	if err != nil {
		t.Errorf("TestPreAddCard fail with %v", err)
		return
	}
	fmt.Printf("TestPreAddCard data is %v", data)
}

func TestRound(t *testing.T) {
	// f := decimal.NewFromFloat(199.97 * 0.027)
	// fmt.Println(f.Round(2))
	// fmt.Println(f.Ceil())

	fmt.Println(time.Now().Unix())
}

func TestApplyBankTransfer(t *testing.T) {
	outOrderID := NewOrderID()
	amount := "1"
	data, err := ApplyBankTransfer(outOrderID, amount)
	if err != nil {
		t.Errorf("TestApplyBankTransfer fail with %v", err)
		return
	}
	fmt.Printf("TestApplyBankTransfer data is %v", data)
}
func TestCashierUnifiedOrder(t *testing.T) {
	outOrderID := NewOrderID()
	amount := "1"
	data, err := CashierUnifiedOrder(outOrderID, amount)
	if err != nil {
		t.Errorf("TestCashierUnifiedOrder fail with %v", err)
		return
	}
	fmt.Printf("TestCashierUnifiedOrder data is %v", data)
}
