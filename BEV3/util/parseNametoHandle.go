package util

import (
	"fmt"
	"regexp"
	"strings"
)

func ParseNametoHandle(name string) string {
	s := strings.Split(name, " ")
	for i := 0; i < len(s); i++ {
		s[i] = strings.ToLower(s[i])
		s[i] = removeVietnameseTones(s[i])
	}
	fmt.Println(strings.Join(s, "----"))
	return strings.Join(s, "-")
}

func removeVietnameseTones(str string) string {
	str = regexp.MustCompile(`/à|á|ạ|ả|ã|â|ầ|ấ|ậ|ẩ|ẫ|ă|ằ|ắ|ặ|ẳ|ẵ/g`).ReplaceAllString(str, "a")
	str = regexp.MustCompile(`/è|é|ẹ|ẻ|ẽ|ê|ề|ế|ệ|ể|ễ/g`).ReplaceAllString(str, "e")
	str = regexp.MustCompile(`/ì|í|ị|ỉ|ĩ/g`).ReplaceAllString(str, "i")
	str = regexp.MustCompile(`/ò|ó|ọ|ỏ|õ|ô|ồ|ố|ộ|ổ|ỗ|ơ|ờ|ớ|ợ|ở|ỡ/g`).ReplaceAllString(str, "o")
	str = regexp.MustCompile(`/ù|ú|ụ|ủ|ũ|ư|ừ|ứ|ự|ữ|ử|/g`).ReplaceAllString(str, "u")
	str = regexp.MustCompile(`/ỳ|ý|ỵ|ỷ|ỹ/g`).ReplaceAllString(str, "y")
	str = regexp.MustCompile(`/đ/g`).ReplaceAllString(str, "d")
	str = regexp.MustCompile(`/À|Á|Ạ|Ả|Ã|Â|Ầ|Ấ|Ậ|Ẩ|Ẫ|Ă|Ằ|Ắ|Ặ|Ẳ|Ẵ/g`).ReplaceAllString(str, "A")
	str = regexp.MustCompile(`/È|É|Ẹ|Ẻ|Ẽ|Ê|Ề|Ế|Ệ|Ể|Ễ/g`).ReplaceAllString(str, "E")
	str = regexp.MustCompile(`/Ì|Í|Ị|Ỉ|Ĩ/g`).ReplaceAllString(str, "I")
	str = regexp.MustCompile(`/Ò|Ó|Ọ|Ỏ|Õ|Ô|Ồ|Ố|Ộ|Ổ|Ỗ|Ơ|Ờ|Ớ|Ợ|Ở|Ỡ/g`).ReplaceAllString(str, "O")
	str = regexp.MustCompile(`/Ù|Ú|Ụ|Ủ|Ũ|Ư|Ừ|Ứ|Ự|Ử|Ữ/g`).ReplaceAllString(str, "U")
	str = regexp.MustCompile(`/Ỳ|Ý|Ỵ|Ỷ|Ỹ/g`).ReplaceAllString(str, "Y")
	str = regexp.MustCompile(`/Đ/g`).ReplaceAllString(str, "D")
	str = regexp.MustCompile(`/ + /g`).ReplaceAllString(str, "")
	str = strings.Trim(str, "")
	str = regexp.MustCompile(`/!|@|%|\^|\*|\(|\)|\+|\=|\<|\>|\?|\/|,|\.|\:|\;|\'|\"|\&|\#|\[|\]|~|\$|_||-|{|}|\||\\/g`).ReplaceAllString(str, "")

	// str = strings.ReplaceAll(str, `/à|á|ạ|ả|ã|â|ầ|ấ|ậ|ẩ|ẫ|ă|ằ|ắ|ặ|ẳ|ẵ/g`, "a")
	// str = strings.ReplaceAll(str, `/è|é|ẹ|ẻ|ẽ|ê|ề|ế|ệ|ể|ễ/g`, "e")
	// str = strings.ReplaceAll(str, `/ì|í|ị|ỉ|ĩ/g`, "i")
	// str = strings.ReplaceAll(str, `/ò|ó|ọ|ỏ|õ|ô|ồ|ố|ộ|ổ|ỗ|ơ|ờ|ớ|ợ|ở|ỡ/g`, "o")
	// str = strings.ReplaceAll(str, `/ù|ú|ụ|ủ|ũ|ư|ừ|ứ|ự|ử|ữ/g`, "u")
	// str = strings.ReplaceAll(str, `/ỳ|ý|ỵ|ỷ|ỹ/g`, "y")
	// str = strings.ReplaceAll(str, `/đ/g`, "d")
	// str = strings.ReplaceAll(str, `/À|Á|Ạ|Ả|Ã|Â|Ầ|Ấ|Ậ|Ẩ|Ẫ|Ă|Ằ|Ắ|Ặ|Ẳ|Ẵ/g`, "A")
	// str = strings.ReplaceAll(str, `/È|É|Ẹ|Ẻ|Ẽ|Ê|Ề|Ế|Ệ|Ể|Ễ/g`, "E")
	// str = strings.ReplaceAll(str, `/Ì|Í|Ị|Ỉ|Ĩ/g`, "I")
	// str = strings.ReplaceAll(str, `/Ò|Ó|Ọ|Ỏ|Õ|Ô|Ồ|Ố|Ộ|Ổ|Ỗ|Ơ|Ờ|Ớ|Ợ|Ở|Ỡ/g`, "O")
	// str = strings.ReplaceAll(str, `/Ù|Ú|Ụ|Ủ|Ũ|Ư|Ừ|Ứ|Ự|Ử|Ữ/g`, "U")
	// str = strings.ReplaceAll(str, `/Ỳ|Ý|Ỵ|Ỷ|Ỹ/g`, "Y")
	// str = strings.ReplaceAll(str, `/Đ/g`, "D")
	// // Some system encode vietnamese combining accent as individual utf-8 characters
	// // Một vài bộ encode coi các dấu mũ, dấu chữ như một kí tự riêng biệt nên thêm hai dòng này
	// str = strings.ReplaceAll(str, `/\u0300|\u0301|\u0303|\u0309|\u0323/g`, "") // ̀ ́ ̃ ̉ ̣  huyền, sắc, ngã, hỏi, nặng
	// str = strings.ReplaceAll(str, `/\u02C6|\u0306|\u031B/g`, "")               // ˆ ̆ ̛  Â, Ê, Ă, Ơ, Ư
	// // Remove extra spaces
	// // Bỏ các khoảng trắng liền nhau
	// str = strings.ReplaceAll(str, `/ + /g`, " ")
	// // Remove punctuations
	// // Bỏ dấu câu, kí tự đặc biệt
	// str = strings.ReplaceAll(str, `/!|@|%|\^|\*|\(|\)|\+|\=|\<|\>|\?|\/|,|\.|\:|\;|\'|\"|\&|\#|\[|\]|~|\$|_||-|{|}|\||\\/g`, " ")
	return str
}
