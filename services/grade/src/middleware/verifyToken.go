package middleware

import (
	"log"
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
		log.Println("Error jwt parse -> " + err.Error());
		return ctx.SendStatus(401);
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uid, ok1 := claims["uid"].(string);
		role, ok2 := claims["role"].(string);
		if (ok1 && ok2){
			ctx.Locals("uid", uid);
			ctx.Locals("role", role);
			return ctx.Next();
		}
	}

	log.Println("Error get token claims");
	return ctx.SendStatus(401);
}