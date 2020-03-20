/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2020/1/19 16:40
*/
package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func Code(length int) string {
	if length <= 0 {
		length = 4
	}
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := ""
	for i := 0; i < length; i++ {
		s += strconv.Itoa(rd.Intn(9))
	}
	return s
}
