package logic

import (
	"auth/config"
	"auth/models"
	"context"
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GoogleAuthen(ctx fiber.Ctx, token string) (string, error) {
	// Decode TokenID
	decodedToken, err := config.AuthClient.VerifyIDToken(
		context.Background(),
		token,
	)
	if err != nil {
		return "", errors.New("Error verify token -> " + err.Error());
	}

	googleUID := decodedToken.UID;
	email := decodedToken.Claims["email"].(string);
	username := decodedToken.Claims["name"].(string);
	photoURL := decodedToken.Claims["picture"].(string);

	// Check user exist in table
	user, err := gorm.G[models.User](config.DB).Where("email = ?", email).First(ctx);
	if err != nil {
		// Create new user
		if err == gorm.ErrRecordNotFound {
			uid := uuid.New();
			user = models.User{UID: uid.String(), Google_UID: googleUID, Email: email, Username: username, PhotoURL: photoURL, Role: "User", Version: 1};
			err := gorm.G[models.User](config.DB).Create(ctx, &user);
			if err != nil {
				return "", errors.New("Error create new user to database -> " + err.Error());
			}
		} else {
			return "", errors.New("Error query user in database -> " + err.Error());
		}
	} else if user.Google_UID == ""{
		// If user exist but never login with google
		user.Google_UID = googleUID;
		user.Version ++;
		rows, err := gorm.G[models.User](config.DB).Where("uid = ? AND version = ?", user.UID, user.Version - 1).Select("google_uid", "version").Updates(ctx, user);
		if err != nil {
			return "", errors.New("Error add google uid to database -> " + err.Error());
		}
		if rows == 0 {
			return "", errors.New("Error add google uid to database -> Race condition"); 
		}
	}

	// Generate JSON web token
	jwtToken, err := GenerateToken(user.UID, user.Role);
	if err != nil {
		return "", errors.New("Error generate JWT -> " + err.Error());
	}

	return jwtToken, nil;
}