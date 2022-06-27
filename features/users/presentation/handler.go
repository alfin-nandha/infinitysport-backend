package presentation

import (
	"net/http"
	"project/e-comerce/features/users"
	_requestUser "project/e-comerce/features/users/presentation/request"
	_responseUser "project/e-comerce/features/users/presentation/response"
	"project/e-comerce/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userBusiness users.Business
}

func NewUserHandler(business users.Business) *UserHandler {
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
			helper.ResponseFailed("failed to get all data"))
	}

	return c.JSON(http.StatusOK,
		helper.ResponseSuccessWithData("success", _responseUser.FromCoreList(result)))
}

func (h *UserHandler) GetDataById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := h.userBusiness.GetDataById(id)
	if id != result.ID {
		return c.JSON(http.StatusBadGateway,
			helper.ResponseFailed("no get data user"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailed("failed to get data"))
	}
	return c.JSON(http.StatusOK,
		helper.ResponseSuccessWithData("success", _responseUser.FromCore(result)))
}

func (h *UserHandler) Insert(c echo.Context) error {
	user := _requestUser.User{}
	err_bind := c.Bind(&user)
	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailed("failed to bind insert data"))
	}
	userCore := _requestUser.ToCore(user)
	_, err := h.userBusiness.InsertData(userCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailed("failed to insert data"))
	}
	return c.JSON(http.StatusOK,
		helper.ResponseSuccessNoData("success insert data"))
}

func (h *UserHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := h.userBusiness.DeleteData(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailed("failed to delete data"))
	}
	return c.JSON(http.StatusOK,
		helper.ResponseFailed("success delete data"))
}

func (h *UserHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	userReq := _requestUser.User{}
	err_bind := c.Bind(&userReq)
	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailed("failed to bind update data"))
	}

	userCore := _requestUser.ToCore(userReq)
	_, err := h.userBusiness.UpdateData(userCore, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailed("failed to insert data"))
	}
	return c.JSON(http.StatusOK,
		helper.ResponseSuccessNoData("success update data"))
}
