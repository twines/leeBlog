/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2020/1/17 15:54
*/
package admins

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"leeblog.com/pkg/utils"
)

func Captcha(c *gin.Context) {
	utils.Captcha(c, 4)
}
func CaptchaVerify(c *gin.Context) {
	value := c.Param("value")
	fmt.Println(utils.CaptchaVerify(c, value))
}
