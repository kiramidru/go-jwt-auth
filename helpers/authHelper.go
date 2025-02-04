package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) error {
	userType := c.GetString("user_type")
	var err error
	if userType != role {
		err = errors.New("Unauthorized access")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userID string) error {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	var err error

	if userType == "USER" && uid != userID {
		err = errors.New("Unauthorized access")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}
