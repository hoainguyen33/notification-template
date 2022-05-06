package keycloak

import (
	"context"
	"os"

	"github.com/Nerzal/gocloak/v10"
	"github.com/gin-gonic/gin"
)

func AuthenticateWithKeyCloak(ctx *gin.Context, username string, password string) (*gocloak.JWT, error) {
	//ctx := context.Background()
	var hostname = os.Getenv("KEY_CLOAK_HOST")
	var clientId = os.Getenv("KEY_CLOAK_CLIENT_ID")
	var clientSecret = os.Getenv("KEY_CLOAK_CLIENT_SECRET")
	var realm = os.Getenv("KEY_CLOAK_REALM")
	client := gocloak.NewClient(hostname)
	return client.Login(ctx, clientId, clientSecret, realm, username, password)
}

func GetKeyCloakUserInfo(ctx *gin.Context, accessToken string) (*gocloak.UserInfo, error) {
	var hostname = os.Getenv("KEY_CLOAK_HOST")
	var realm = os.Getenv("KEY_CLOAK_REALM")
	client := gocloak.NewClient(hostname)

	return client.GetUserInfo(ctx, accessToken, realm)
}

func GetAccessTokenAdmin() (*gocloak.JWT, error) {
	ctx := context.Background()
	var hostname = os.Getenv("KEY_CLOAK_HOST")
	var userAdmin = os.Getenv("KEY_CLOAK_USER_NAME")
	var passwordAdmin = os.Getenv("KEY_CLOAK_PASSWORD")
	var realm = os.Getenv("KEY_CLOAK_REALM")

	client := gocloak.NewClient(hostname)

	return client.LoginAdmin(ctx, userAdmin, passwordAdmin, realm)
}

func GetUsersById(ctx context.Context, jwt *gocloak.JWT, userID string) (*gocloak.User, error) {
	var hostname = os.Getenv("KEY_CLOAK_HOST")
	var realm = os.Getenv("KEY_CLOAK_REALM")
	client := gocloak.NewClient(hostname)
	return client.GetUserByID(ctx, jwt.AccessToken, realm, userID)
}

type GetcareUser struct {
	Name  string
	Phone string
	Email string
}
