package middleware

import (
	"github.com/fauzan264/evermos-rakamin/constants"
	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/dto/response"
	"github.com/fauzan264/evermos-rakamin/services"
	"github.com/fauzan264/evermos-rakamin/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(userService services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeaderToken := c.Get("token")

		jwtService := utils.NewJWTService()
		token, err := jwtService.ValidateToken(authHeaderToken)
		if err != nil || !token.Valid {
			return sendUnauthorizedResponse(c)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return sendUnauthorizedResponse(c)
		}

		userID, ok := claims["id"].(float64)
		if !ok {
			return sendUnauthorizedResponse(c)
		}

		requestUser := request.GetByUserIDRequest{
			ID: int(userID),
		}

		user, err := userService.GetUserByID(requestUser)
		if err != nil {
			return sendUnauthorizedResponse(c)
		}

		c.Locals("authUser", user)

		return c.Next()
	}
}

func sendUnauthorizedResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(response.Response{
		Status: false,
		Message: constants.Unauthorized,
		Errors: []string{constants.ErrUnauthorized.Error()},
		Data: nil,
	})
}
