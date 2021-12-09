package common

import "os"

func GetPwd() string{
	str, _ := os.Getwd()
	return str
}
