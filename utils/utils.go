package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/Nerzal/gocloak/v10"
	"github.com/guregu/null"
)

const charsetString = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ"

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

func GetNameFromUser(user gocloak.User) string {
	name := ""
	if user.FirstName != nil {
		name += fmt.Sprintf("%s", *user.FirstName)
	}

	if user.LastName != nil {
		if name != "" {
			name += " "
		}
		name += fmt.Sprintf("%s", *user.LastName)
	}

	return name
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

func NullIntToInt32(value null.Int) int32 {
	return int32(value.Int64)
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
