package config

type Captcha struct {
	WIDTH      int     `mapstructure:"CAPTCAH_WIDTH"`
	HEIGHT     int     `mapstructure:"CAPTCAH_HEIGHT"`
	LENGTH     int     `mapstructure:"CAPTCAH_LENGTH"`
	MAKSKEW    float64 `mapstructure:"CAPTCAH_MAKSKEW"`
	DOTCOUNT   int     `mapstructure:"CAPTCAH_DOTCOUNT"`
	EXPIRETIME string  `mapstructure:"CAPTCAH_EXPIRE_TIME"` // unit: minutes
}
