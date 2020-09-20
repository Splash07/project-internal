package handler

import (
	"net/http"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
)

// SSO handler
var SSO ssoHandler

func init() {
	SSO = ssoHandler{}
}

type ssoHandler struct {
}

// Check func
func (o ssoHandler) Check(c echo.Context) (err error) {
	type myRequest struct {
		Token string `json:"token" query:"token" validate:"required"`
	}

	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}

	redisClient := cfg.Redis["online"].GetClient()

	value, err := redisClient.Get(request.Token).Result()

	if err == redis.Nil {
		return echo.NewHTTPError(http.StatusNotFound, "Không tìm thấy hoặc đường dẫn đã hết hạn")
	}

	clientID, errParse := strconv.Atoi(value)
	if errParse != nil {
		err = errParse
		return
	}

	return c.JSON(success(map[string]interface{}{
		"clientID": clientID,
	}))
}
