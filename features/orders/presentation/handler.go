package presentation

import (
	"net/http"
	"project/e-comerce/features/orders"
	"project/e-comerce/features/orders/presentation/request"
	"project/e-comerce/helper"
	"project/e-comerce/middlewares"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct{
	orderBusiness orders.Bussiness
}

func NewOrderHandler(business orders.Bussiness) *OrderHandler{
	return &OrderHandler{
		orderBusiness: business,
	}
}

func (h OrderHandler)AddOrder(c echo.Context)(error){
	userID_token,errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to get user id",
		})
	}

	reqData := request.Order{}
	errBind := c.Bind(&reqData)
	if errBind != nil{
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to bind data"))
	}

	dataCore := request.ToCore(reqData)
	dataCore.UserID = userID_token
	err := h.orderBusiness.AddOrder(dataCore, reqData.CartID)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to add order"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success"))
}

func (h OrderHandler)GetOrder(c echo.Context)(error){
	userID_token,errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to get user id",
		})
	}
	dataOrder,err := h.orderBusiness.GetOrder(userID_token)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to get order",
		})
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("sueccess", dataOrder))
}

func (h OrderHandler)GeOrderDetailByOrderID(c echo.Context)(error){
	idOrder,errParam := strconv.Atoi(c.Param("orderid"))
	if errParam != nil{
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get param order id"))
	}

	dataOrderDetail,err := h.orderBusiness.GetOrderDetail(idOrder)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to get order",
		})
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success", dataOrderDetail))
}