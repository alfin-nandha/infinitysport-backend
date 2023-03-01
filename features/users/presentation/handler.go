package presentation

import (
	"net/http"
	Users "project/e-comerce/features/users"
	Request "project/e-comerce/features/users/presentation/request"
	Response "project/e-comerce/features/users/presentation/response"
	Helper "project/e-comerce/helper"
	Middlewares "project/e-comerce/middlewares"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userBusiness Users.Business
}

func NewUserHandler(business Users.Business) *UserHandler {
	return &UserHandler{
		userBusiness: business,
	}
}

func (h *UserHandler) GetAll(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	result, err := h.userBusiness.GetAllData(limitint, offsetint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			Helper.ResponseFailed("failed to get all data"))
	}

	return c.JSON(http.StatusOK,
		Helper.ResponseSuccessWithData("success", Response.FromCoreList(result)))
}

func (h *UserHandler) GetDataById(c echo.Context) error {
	userID_token, errToken := Middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get user id",
		})
	}

	result, err := h.userBusiness.GetDataById(userID_token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			Helper.ResponseFailed("failed to get data"))
	}
	return c.JSON(http.StatusOK,
		Helper.ResponseSuccessWithData("success", Response.FromCore(result)))
}

func (h *UserHandler) Insert(c echo.Context) error {
	user := Request.User{}
	err_bind := c.Bind(&user)
	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError,
			Helper.ResponseFailed("failed to bind insert data"))
	}
	userCore := Request.ToCore(user)
	_, err := h.userBusiness.InsertData(userCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			Helper.ResponseFailed("failed to insert data"))
	}
	return c.JSON(http.StatusOK,
		Helper.ResponseSuccessNoData("success insert data"))
}

func (h *UserHandler) Delete(c echo.Context) error {
	userID_token, errToken := Middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get user id",
		})
	}
	_, err := h.userBusiness.DeleteData(userID_token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			Helper.ResponseFailed("failed to delete data"))
	}
	return c.JSON(http.StatusOK,
		Helper.ResponseSuccessNoData("success delete data"))
}

func (h *UserHandler) Update(c echo.Context) error {
	userID_token, errToken := Middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get user id",
		})
	}

	userReq := Request.User{}
	err_bind := c.Bind(&userReq)
	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError,
			Helper.ResponseFailed("failed to bind update data"))
	}

	userCore := Request.ToCore(userReq)
	_, err := h.userBusiness.UpdateData(userCore, userID_token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			Helper.ResponseFailed("failed to insert data"))
	}
	return c.JSON(http.StatusOK,
		Helper.ResponseSuccessNoData("success update data"))
}
