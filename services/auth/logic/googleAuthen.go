package logic

import (
	"auth/config"
	"auth/models"
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GoogleAuthen(ctx fiber.Ctx, token string) (string, error) {
	fmt.Println(token);
	decodedToken, err := config.AuthClient.VerifyIDToken(
		context.Background(),
		token,
	)
	if err != nil {
		return "", err;
	}

	googleUID := decodedToken.UID;
	email := decodedToken.Claims["email"].(string);
	username := decodedToken.Claims["name"].(string);

	user, err := gorm.G[models.User](config.DB).Where("email = ?", email).First(ctx);
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			uid := uuid.New();
			user = models.User{UID: uid.String(), Google_UID: googleUID, Email: email, Username: username, Version: 1};
			err := gorm.G[models.User](config.DB).Create(ctx, &user);
			if err != nil {
				return "", err;
			}
		} else {
			return "", err;
		}
	} else if user.Google_UID == ""{
		user.Google_UID = googleUID;
		user.Version ++;
		rows, err := gorm.G[models.User](config.DB).Where("uid = ? AND version = ?", user.UID, user.Version - 1).Select("google_uid", "version").Updates(ctx, user);
		if err != nil {
			panic(err);
		}
		if rows == 0 {
			return "", errors.New("User is just updated"); 
		}
	}

	jwtToken, err := GenerateToken(user.UID);
	if err != nil {
		return "", err;
	}

	return jwtToken, nil;
}