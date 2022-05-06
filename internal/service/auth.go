package service

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"time"

	"github.com/Nerzal/gocloak/v10"
	"github.com/gin-gonic/gin"
)

const charsetString = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const charsetStringUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const charsetNumber = "0123456789"

func MySha1(input string) string {
	bv := []byte(input + "_" + os.Getenv("TOKEN_SECRET"))
	hasher := sha1.New()
	hasher.Write(bv)

	return hex.EncodeToString(hasher.Sum(nil))
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	len := len(charset)
	for i := range b {
		b[i] = charset[seededRand.Intn(len)]
	}
	return string(b)
}

func GeneratorStringUpper(length int) string {
	return StringWithCharset(length, charsetStringUpper)
}

func GeneratorString(length int) string {
	return StringWithCharset(length, charsetString)
}

func GeneratorNumber(length int) string {
	return StringWithCharset(length, charsetNumber)
}

func GetTableNameFromPath(pathURL string) string {
	tableName := ""
	re := regexp.MustCompile("api/v1/(getcare_\\w+)")
	arrStr := re.FindStringSubmatch(pathURL)

	if len(arrStr) > 1 {
		re = regexp.MustCompile("(getcare_\\w+)_action$")
		arrStr2 := re.FindStringSubmatch(arrStr[1])
		if len(arrStr2) > 1 {
			tableName = arrStr2[1]
		} else {
			tableName = arrStr[1]
		}
	} else {
		re = regexp.MustCompile("api/v1/(\\w+)")
		arrStr = re.FindStringSubmatch(pathURL)
		if len(arrStr) > 1 {
			tableName = fmt.Sprintf("getcare_%s", arrStr[1])
		}
	}

	return tableName
}

func GetKeyCloakUserInfo(ctx *gin.Context, accessToken string) (*gocloak.UserInfo, error) {
	var hostname = os.Getenv("KEY_CLOAK_HOST")
	var realm = os.Getenv("KEY_CLOAK_REALM")
	client := gocloak.NewClient(hostname)

	return client.GetUserInfo(ctx, accessToken, realm)
}

func GetKeyCloakUserByID(ctx *gin.Context, userID string) (*gocloak.User, error) {
	var hostname = os.Getenv("KEY_CLOAK_HOST")
	var realm = os.Getenv("KEY_CLOAK_REALM")
	client := gocloak.NewClient(hostname)

	jwt, err := GetKeyCloakAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	return client.GetUserByID(ctx, jwt.AccessToken, realm, userID)
}

func GetKeyCloakUsers(ctx *gin.Context, phone string, keyword string, page int, pageSize int) ([]*gocloak.User, error) {
	var hostname = os.Getenv("KEY_CLOAK_HOST")
	var realm = os.Getenv("KEY_CLOAK_REALM")
	client := gocloak.NewClient(hostname)

	_jwt, err := GetKeyCloakAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}

	param := &gocloak.GetUsersParams{
		Max:   &pageSize,
		First: &offset,
	}
	if keyword != "" {
		param.Search = &keyword
	}
	if phone != "" {
		param.Username = &phone
	}

	return client.GetUsers(ctx, _jwt.AccessToken, realm, *param)
}

func GetKeyCloakUsersCount(ctx *gin.Context, keyword string) (int, error) {
	var hostname = os.Getenv("KEY_CLOAK_HOST")
	var realm = os.Getenv("KEY_CLOAK_REALM")
	client := gocloak.NewClient(hostname)

	_jwt, err := GetKeyCloakAccessToken(ctx)
	if err != nil {
		return 0, err
	}

	return client.GetUserCount(ctx, _jwt.AccessToken, realm, gocloak.GetUsersParams{
		Search: &keyword,
	})
}

func GetKeyCloakAccessToken(ctx *gin.Context) (*gocloak.JWT, error) {
	var hostname = os.Getenv("KEY_CLOAK_HOST")
	var userAdmin = os.Getenv("KEY_CLOAK_USER_NAME")
	var passwordAdmin = os.Getenv("KEY_CLOAK_PASSWORD")
	var realm = os.Getenv("KEY_CLOAK_REALM")

	client := gocloak.NewClient(hostname)

	return client.LoginAdmin(ctx, userAdmin, passwordAdmin, realm)
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
