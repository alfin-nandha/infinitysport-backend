package presentation

import (
	"net/http"
	Auth "project/e-comerce/features/auth"
	Request "project/e-comerce/features/auth/presentation/request"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	userBusiness Auth.Business
}

func NewAuthHandler(business Auth.Business) *AuthHandler {
	return &AuthHandler{
		userBusiness: business,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	reqBody := Request.User{}
	errBind := c.Bind(&reqBody)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get bind data",
		})
	}

	authCore := Request.ToCore(reqBody)
	result, name, err := h.userBusiness.Login(authCore)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get token data" + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"name":    name,
		"token":   result,
	})
}
