package auth

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tphan267/common/api"
	"github.com/tphan267/common/cache"
	"github.com/tphan267/common/http"
	"github.com/tphan267/common/system"
	"github.com/tphan267/common/utils"
)

func RemoteAuthMiddleware() fiber.Handler {
	extractTokens := "header:Authorization,query:auth_token"
	return remoteMiddleware(extractTokens)
}

func RemoteAPIKeyMiddleware() fiber.Handler {
	extractTokens := "header:Authorization,query:apikey"
	return remoteMiddleware(extractTokens)
}

// RemoteMiddleware accept both auth_token or apikey
func RemoteMiddleware() fiber.Handler {
	extractTokens := "header:Authorization,query:auth_token,query:apikey"
	return remoteMiddleware(extractTokens)
}

func remoteMiddleware(extractTokens ...string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ExtractToken(ctx, extractTokens...)
		if token == "" {
			return api.ErrorUnauthorizedResp(ctx, "Missing auth token or apikey")
		}

		act, err := RemoteAccount(token)
		if err != nil {
			return api.ErrorUnauthorizedResp(ctx, err.Error())
		}

		ctx.Locals("authToken", token)
		ctx.Locals("account", act)
		ctx.Locals("uiID", act.ID)
		ctx.Locals("usID", fmt.Sprintf("%d", act.ID))

		return ctx.Next()
	}
}

func RemoteAccount(token string) (act *AuthTokenData, err error) {
	act = &AuthTokenData{}
	err = cache.GetObj(token, act)

	if err != nil || act.ID == 0 {
		err = nil
		resp := &AuthValidateResponse{}
		err := http.Get(system.Env("AUTH_API")+"/auth/validate", resp, map[string]string{
			"Authorization": "Bearer " + token,
		})
		if err != nil {
			return nil, err
		}
		if !resp.Success {
			return nil, errors.New(resp.Error.Message)
		}
		act = resp.Data
		duration, _ := utils.ParseDuration(system.Env("AUTH_CACHE_DURATION", "1h"))
		cache.SetObj(token, act, duration)
	}

	return act, err
}
