package utils

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	charsetString = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsetNumber = "0123456789"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func GeneratorString(length int) string {
	return StringWithCharset(length, charsetString)
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	len := len(charset)
	for i := range b {
		b[i] = charset[seededRand.Intn(len)]
	}
	return string(b)
}

func GeneratorNumber(length int) string {
	return StringWithCharset(length, charsetNumber)
}

func FindString(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func PointerToString(pointer *string) string {
	if pointer != nil {
		return *pointer
	}
	return ""
}

func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

func ParseInt32(idStr string) (int32, error) {
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return int32(id), err
	}
	return int32(id), err
}

func RandomHex(n int) []byte {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return nil
	}
	return bytes
}

func FindIndex(arr []interface{}, key interface{}) int {
	for i, v := range arr {
		if v == key {
			return i
		}
	}
	return -1
}

func CheckIn(arr []string, v string) bool {
	for _, value := range arr {
		if value == v {
			return true
		}
	}
	return false
}

func StringToTime(str string) (time.Time, error) {
	format := "2006-01-02T15:04:05.000Z"
	return time.Parse(format, str)
}
