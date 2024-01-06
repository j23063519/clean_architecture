package captcha

import (
	"sync"

	"github.com/j23063519/clean_architecture/config"
	"github.com/j23063519/clean_architecture/pkg/redis"
	"github.com/mojocn/base64Captcha"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

// only new one time (singleton pattern)
var once sync.Once

// internalCaptcha internal use Captcha
var internalCaptcha *Captcha

// new Captcha
func NewCaptcha() *Captcha {
	once.Do(func() {
		// initializing captcha
		internalCaptcha = &Captcha{}

		// initializing store
		store := RedisStore{
			RedisClient: redis.Redis,
			KeyPrefix:   config.Config.App.NAME + ":captcha:",
		}

		// initializing driver
		driver := base64Captcha.NewDriverDigit(
			config.Config.Captcha.HEIGHT,
			config.Config.Captcha.WIDTH,
			config.Config.Captcha.LENGTH,
			config.Config.Captcha.MAKSKEW,
			config.Config.Captcha.DOTCOUNT,
		)

		// initializing and resigning base64Captcha
		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)
	})

	return internalCaptcha
}

// generate captcha
func (c *Captcha) GenerateCaptcha() (id, b64s, answer string, err error) {
	return c.Base64Captcha.Generate()
}

// verify captcha
func (c *Captcha) VerifyCaptcha(id string, answer string) (match bool) {
	// It is convenient for users to submit multiple times and prevents them from having to enter the image verification code multiple times after submitting the form incorrectly.
	return c.Base64Captcha.Verify(id, answer, false)
}
