package dto

type ReqUserListHandler struct {
	UID             int    `json:"uid" form:"uid"`
	Phone           string `json:"phone" form:"phone"`
	Email           string `json:"email" form:"email"`
	Nickname        string `json:"nickname" form:"nickname"`
	Status          int    `json:"status" form:"status"`
	StartDate       string `json:"start_date" form:"start_date"`
	EndDate         string `json:"end_date" form:"end_date"`
	Page            int    `json:"page" form:"page"`
	PageSize        int    `json:"page_size" form:"page_size"`
	Ids             []int  `json:"ids" form:"ids"`
	ActiveStartDate string `json:"active_start_date" form:"active_start_date"`
	ActiveEndDate   string `json:"active_end_date" form:"active_end_date"`
}

type ReqLoginHandler struct {
	Nickname        string `json:"nickname" form:"nickname"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
	Action          string `json:"action" form:"action"`
}

type ReqAuthCodeHandler struct {
	Phone string `json:"phone" form:"phone"`
}

type ReqAuthPhoneHandler struct {
	Phone string `json:"phone" form:"phone"`
	Code  string `json:"code" form:"code"`
}

type Token struct {
	UserId    int    `json:"id"`
	LoginTime string `json:"login_time"`
	Expire    int64  `json:"expire"`
}

type ReqModifyUserInfoHandler struct {
	Nickname string `json:"nickname" form:"nickname"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
}

type ReqModifyUserPasswordHandler struct {
	OldPassword     string `json:"old_password" form:"old_password"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type ReqResetUserPasswordHandler struct {
	Email string `json:"email" form:"email"`
}

type ResetPasswordToken struct {
	Email     string `json:"email"`
	ResetTime string `json:"reset_time"`
	Expire    int64  `json:"expire"`
}

type ReqResetPasswordSetNewHandler struct {
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
	Token           string `json:"token" form:"token"`
}
