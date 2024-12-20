package utils

// MaskPhoneNumber 隐藏手机号中间四位
func MaskPhoneNumber(phone string) string {
	l := len(phone)
	if l == 11 {
		return phone[:3] + "****" + phone[7:]
	}
	return "****"
}
