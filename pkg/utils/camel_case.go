/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2020/1/19 15:43
*/
package utils

import (
	"fmt"
	"strings"
)

func CamelCase(s string) string {
	s1 := make([]string, len(s))
	for i, v := range s {
		if i == 0 {
			fmt.Println(strings.ToLower(string(v)))
			s1[i] = strings.ToLower(string(v))
		} else {
			s1[i] = string(s[i])
		}
	}
	return strings.Join(s1, "")
}
