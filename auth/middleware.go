package auth

import (
	"fmt"

	"github.com/tphan267/common/api"
	"github.com/tphan267/common/system"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler:   authError,
		SuccessHandler: authSuccess,
		SigningKey:     []byte(system.Env("JWT_SECRET")),
		SigningMethod:  "HS256",
		ContextKey:     "authToken",
		TokenLookup:    fmt.Sprintf("header:%s,cookie:auth_token,query:auth_token", fiber.HeaderAuthorization),
	})
}

func authError(ctx *fiber.Ctx, err error) error {
	accept := ctx.Get("Accept")
	uri := ctx.Query("redirect_uri")

	if accept != "application/json" && uri != "" {
		return ctx.Redirect(uri)
	}

	return api.ErrorUnauthorizedResp(ctx, err.Error())
}

func authSuccess(ctx *fiber.Ctx) error {
	jwtData := ctx.Locals("authToken").(*jwt.Token)
	claims := jwtData.Claims.(jwt.MapClaims)
	uiID := uint64(claims["id"].(float64))
	usID := fmt.Sprintf("%d", uiID)

	ctx.Locals("uiID", uiID)
	ctx.Locals("usID", usID)

	// // validate if user exist and to context
	// account := &models.Account{}
	// database.DB.Find(account, "id=?", id)
	// if account.ID == 0 {
	// 	return api.ErrorUnauthorizedResp(ctx, fmt.Sprintf("Account #%v does not exist!", id))
	// }
	// ctx.Locals("account", account)

	return ctx.Next()
}
