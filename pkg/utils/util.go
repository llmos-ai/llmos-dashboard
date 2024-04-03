package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"

	entv1 "github.com/llmos/llmos-dashboard/pkg/generated/ent"
)

func GetSessionUser(c *gin.Context) (*entv1.User, error) {
	userObj, exist := c.Get("user")
	if !exist || userObj == nil {
		return nil, fmt.Errorf("empty user")
	}

	user, ok := userObj.(*entv1.User)
	if !ok {
		return nil, fmt.Errorf("failed to parse user obj: %v", userObj)
	}

	return user, nil
}
