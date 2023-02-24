package epaydemo

import (
	"fmt"
	"testing"
)

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
	data, err := ApplyOpenAcct("100010", "", "https://www.baidu.com", "https://www.baidu.com")
	if err != nil {
		t.Errorf("TestApplyOpenAcct fail with %v", err)
		return
	}
	fmt.Printf("TestApplyOpenAcct data is %v", data)
}

func TestWithdrawal(t *testing.T) {
	orderID := NewOrderID()
	data, err := Withdrawal("1", "1", orderID, "1")
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
