package dto

type CaptchaResponse struct {
	CaptchaId     string `json:"captcha_id"`
	PicPath       string `json:"picture_path"`
	CaptchaLength int    `json:"captcha_length""`
}

type LoginRequest struct {
	Mobile    string `json:"mobile"`
	Password  string `json:"password"`
	CaptchaId string `json:"captcha_id"`
	Captcha   string `json:"captcha"`
}

type LoginResponse struct {
	Mobile       string `json:"mobile"`
	UserId       string `json:"user_id"`
	UserName     string `json:"user_name"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RegisterRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
