package presentation

import (
	"errors"
	"net/http"
	Carts "project/e-comerce/features/carts"
	RequestCart "project/e-comerce/features/carts/presentation/request"
	ResponseCart "project/e-comerce/features/carts/presentation/response"
	Helper "project/e-comerce/helper"
	Middlewares "project/e-comerce/middlewares"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartBusiness Carts.Business
}

func NewCartHandler(business Carts.Business) *CartHandler {
	return &CartHandler{
		cartBusiness: business,
	}
}

func (h *CartHandler) GetAll(c echo.Context) error {
	userID_token, errToken := Middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError,
			Helper.ResponseFailed("failed to get user id"))
	}
	UserId := userID_token
	result, err := h.cartBusiness.GetAllCart(UserId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Helper.ResponseFailed(" failed to get all data"))
	}
	return c.JSON(http.StatusOK, Helper.ResponseSuccessWithData("success get data", ResponseCart.FromCoreList(result)))
}

func (h *CartHandler) AddCart(c echo.Context) error {
	userID_token, errToken := Middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError,
			Helper.ResponseFailed("failed to get user id"))
	}

	cartData := RequestCart.Cart{}
	errBind := c.Bind(&cartData)

	if errBind != nil {
		return errors.New("failed to bind data")
	}

	cartCore := RequestCart.ToCore(cartData)
	cartCore.UserID = userID_token
	cartCore.Product.ID = cartData.ProductID

	result, err := h.cartBusiness.AddCart(cartCore)

	if err != nil {
		return errors.New("failed to insert data")
	}

	if result == 0 {
		return c.JSON(http.StatusBadRequest, Helper.ResponseFailed("failed to add data"))
	}

	return c.JSON(http.StatusOK, Helper.ResponseSuccessNoData("success to insert data"))
}

func (h *CartHandler) Update(c echo.Context) error {
	idCart, _ := strconv.Atoi(c.Param("id"))
	userID_token, errToken := Middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError,
			Helper.ResponseFailed("failed to get user id"))
	}

	cartData := RequestCart.Cart{}
	errBind := c.Bind(&cartData)

	if errBind != nil {
		return errors.New("failed to bind data")
	}

	cartCore := RequestCart.ToUpdateCore(cartData)
	cartCore.UserID = userID_token

	result, err := h.cartBusiness.UpdateCart(cartCore, idCart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Helper.ResponseFailed("failed to update cart"))
	}
	if result == 0 {
		return c.JSON(http.StatusBadRequest, Helper.ResponseFailed("failed to update cart"))
	}
	return c.JSON(http.StatusOK, Helper.ResponseSuccessNoData("success to update data"))
}

func (h *CartHandler) Destroy(c echo.Context) error {
	idCart, _ := strconv.Atoi(c.Param("id"))
	userID_token, errToken := Middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError,
			Helper.ResponseFailed("failed to get user id"))
	}
	UserId := userID_token

	result, err := h.cartBusiness.DestroyCart(UserId, idCart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Helper.ResponseFailed("error response to delete cart"))
	}
	if result == 0 {
		return c.JSON(http.StatusBadRequest, Helper.ResponseFailed("failed to delete cart"))
	}
	return c.JSON(http.StatusOK, Helper.ResponseSuccessNoData("success to delete data"))
}
