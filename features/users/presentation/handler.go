package presentation

import (
	"net/http"
	"project/e-comerce/features/users"
	"project/e-comerce/features/users/presentation/request"
	_request_user "project/e-comerce/features/users/presentation/request"
	_response_user "project/e-comerce/features/users/presentation/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{
	userBusiness users.Business
}

func NewUserHandler(business users.Business) *UserHandler{
	return &UserHandler{
		userBusiness: business,
	}
}

func (h *UserHandler)GetAll(c echo.Context) error{
	result, err := h.userBusiness.GetAllData("")
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to get all data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "success",
		"data" : _response_user.FromCoreList(result),
	})
}

func (h *UserHandler)GetDataById(c echo.Context) error{
	id,_ := strconv.Atoi(c.Param("id"))
	user_result, err:= h.userBusiness.GetDataById(id)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to get data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "success",
		"data" : _response_user.FromCore(user_result),
	})
}

func (h *UserHandler)InsertData(c echo.Context)error{
	
	user := _request_user.User{}
	err_bind := c.Bind(&user)
	if err_bind != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to bind insert data",
		})
	}
	userCore := request.ToCore(user)
	err := h.userBusiness.InsertData(userCore)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "success insert data",
	})
	
}

func (h *UserHandler)DeleteData(c echo.Context) error{
	id,_ := strconv.Atoi(c.Param("id"))
	err:= h.userBusiness.DeleteData(id)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to delete data"+err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "success delete data",
	})
}

func (h *UserHandler)UpdateData(c echo.Context)error{
	id,_ := strconv.Atoi(c.Param("id"))

	userReq := request.User{}
	err_bind := c.Bind(&userReq)
	if err_bind != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to bind update data",
		})
	}

	userCore := request.ToCore(userReq)
	err:= h.userBusiness.UpdateData(userCore, id)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "success insert data",
	})
}
