package presentation

import (
	"log"
	"net/http"
	Orders "project/e-comerce/features/orders"
	Request "project/e-comerce/features/orders/presentation/request"
	Helper "project/e-comerce/helper"
	Middlewares "project/e-comerce/middlewares"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderBusiness Orders.Bussiness
}

func NewOrderHandler(business Orders.Bussiness) *OrderHandler {
	return &OrderHandler{
		orderBusiness: business,
	}
}

func (h OrderHandler) AddOrder(c echo.Context) error {
	userID_token, errToken := Middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, Helper.ResponseFailed("failed to get id user"))
	}

	reqData := Request.Order{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, Helper.ResponseFailed("failed to bind data"))
	}

	log.Print("request = ", reqData.Address)

	dataCore := Request.ToCore(reqData)
	dataCore.UserID = userID_token

	log.Print("core address = ", dataCore.Address)
	err := h.orderBusiness.AddOrder(dataCore, reqData.CartID, userID_token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Helper.ResponseFailed(err.Error()))
	}

	return c.JSON(http.StatusOK, Helper.ResponseSuccessNoData("success"))
}

func (h OrderHandler) GetOrder(c echo.Context) error {
	userID_token, errToken := Middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, Helper.ResponseFailed("failed to get user id"))
	}
	dataOrder, err := h.orderBusiness.GetOrder(userID_token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Helper.ResponseFailed(err.Error()))
	}

	//responOrder := response.FromCoreList(dataOrder)

	return c.JSON(http.StatusOK, Helper.ResponseSuccessWithData("success", dataOrder))
}

func (h OrderHandler) GeOrderDetailByOrderID(c echo.Context) error {
	userID_token, errToken := Middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, Helper.ResponseFailed("failed to get user id"))
	}

	idOrder, errParam := strconv.Atoi(c.Param("orderid"))
	if errParam != nil {
		return c.JSON(http.StatusInternalServerError, Helper.ResponseFailed("failed to get param order id"))
	}

	dataOrderDetail, err := h.orderBusiness.GetOrderDetail(idOrder)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, Helper.ResponseSuccessWithData("success", dataOrderDetail))
}

func (h OrderHandler) ConfirmOrder(c echo.Context) error {
	userID_token, errToken := Middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, Helper.ResponseFailed("failed to get user id"))
	}

	orderId, errParam := strconv.Atoi(c.Param("orderid"))
	if errParam != nil {
		return c.JSON(http.StatusInternalServerError, Helper.ResponseFailed("failed to get param order id"))
	}
	payReq := Request.Payment{}
	errBind := c.Bind(&payReq)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, Helper.ResponseFailed("failed to get bind payment data"))
	}

	payCore := Request.ToPaymentCore(payReq)

	errRespon := h.orderBusiness.ConfirmOrder(payCore, orderId, userID_token)
	if errRespon != nil {
		return c.JSON(http.StatusInternalServerError, Helper.ResponseFailed(errRespon.Error()))
	}

	return c.JSON(http.StatusOK, Helper.ResponseSuccessNoData("success"))
}
func (h OrderHandler) CancelOrder(c echo.Context) error {
	userID_token, errToken := Middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, Helper.ResponseFailed("failed to get user id"))
	}

	orderId, errParam := strconv.Atoi(c.Param("orderid"))
	if errParam != nil {
		return c.JSON(http.StatusInternalServerError, Helper.ResponseFailed("failed to get param order id"))
	}

	errRespon := h.orderBusiness.CancelOrder(orderId, userID_token)
	if errRespon != nil {
		return c.JSON(http.StatusInternalServerError, Helper.ResponseFailed(errRespon.Error()))
	}
	return c.JSON(http.StatusOK, Helper.ResponseSuccessNoData("success"))
}
