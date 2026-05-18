package middleware

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v4"
)

func VerifyToken(ctx fiber.Ctx) error {
	tokenString := ctx.Cookies("Bearer");
	token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid;
		}
		return []byte(os.Getenv("JWT_SECRET")), nil;
	})

	if err != nil {
		return ctx.SendStatus(401);
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uid := claims["uid"].(string);
		ctx.Locals("uid", uid);
		return ctx.Next();
	}

	return ctx.SendStatus(401);
}