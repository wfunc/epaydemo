package epaydemo

import (
	"fmt"
	"testing"

	"github.com/codingeasygo/util/converter"
)

// 支付宝扫码支付
func TestApplyScanPay(t *testing.T) {
	orderID := NewOrderID()
	_, err := ApplyScanPay(orderID, "0.01", "127.0.0.1")
	if err != nil {
		t.Errorf("TestApplyScanPay fail with %v", err)
		return
	}
}

// 汇潮用户进件
func TestHCNetwork(t *testing.T) {
	hcNetwork, err := HCNetwork("test1", "123", "123", "123", "123", "123")
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

func TestHCPayOut(t *testing.T) {
	orderID := NewOrderID()
	hcPayOut, err := HCPayOut(orderID, "18200892266_18200892266", "6214830207649988", "18200892266", "招商银行", "吴辉南", "11")
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
	hcSubMerBindCard, err := HCSubMerBindCard("test", "test", "test", "招商银行")
	if err != nil {
		t.Errorf("TestHCQueryBalance fail with %v", err)
		return
	}
	fmt.Printf("TestHCQueryBalance data is %v\n", converter.JSON(hcSubMerBindCard))
}

// 查询订单接口
func TestQueryOrder(t *testing.T) {
	_, err := QueryOrder("202302041730530100001")
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
	_, err := PayCreate(orderID, "test", "12312341234", "1", "47.96.26.214")
	if err != nil {
		t.Errorf("TestPayCreate fail with %v", err)
		return
	}
}

func TestQueryLinkedAcct(t *testing.T) {
	_, err := QueryLinkedAcct("test")
	if err != nil {
		t.Errorf("TestQueryLinkedAcct fail with %v", err)
		return
	}
}

func TestQueryUserInfo(t *testing.T) {
	ApiURL = "https://epay.michongfun.com"
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
	// ApiURL = "http://127.0.0.1:3838"
	_, err := ApplyPasswordToken("test", "test", "setting_password")
	if err != nil {
		t.Errorf("TestApplyPasswordToken fail with %v", err)
		return
	}
}

func TestApplyIndividual(t *testing.T) {
	_, err := ApplyIndividual("test", "test", "张三", "123", "20500822", "广东省广州市", "19", "", "")
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

// 请求开户页面
func TestApplyOpenAcct(t *testing.T) {
	// ApiURL = "http://127.0.0.1:3838"
	data, err := ApplyOpenAcct("100010", "", "https://www.baidu.com")
	if err != nil {
		t.Errorf("TestApplyOpenAcct fail with %v", err)
		return
	}
	fmt.Printf("TestApplyOpenAcct data is %v", data)
}

func TestWithdrawal(t *testing.T) {
	Debug = true
	orderID := NewOrderID()
	data, err := Withdrawal("1", "1", orderID, "1", "1")
	if err != nil {
		t.Errorf("TestWithdrawal fail with %v", err)
		return
	}
	fmt.Printf("TestWithdrawal data is %v", data)
}

func TestPreAddCard(t *testing.T) {
	// ApiURL = "http://127.0.0.1:3838"
	orderID := NewOrderID()
	data, err := PreAddCard(orderID, "100002", "https://example.com/notify/testNofity", "https://example.com/notify/testNofity", "", "")
	if err != nil {
		t.Errorf("TestPreAddCard fail with %v", err)
		return
	}
	fmt.Printf("TestPreAddCard data is %v", data)
}
