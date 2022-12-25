package myfuzz

import "github.com/golang-module/dongle"

func Fuzz(data []byte) int {
	dongle.Decode.FromString(string(data)).ByBase16().ToString()
	return 0
}
