package presentation

import (
	"fmt"
	"net/http"
	"project/e-comerce/features/products"
	"time"

	_request_product "project/e-comerce/features/products/presentation/request"
	_response_product "project/e-comerce/features/products/presentation/response"
	"project/e-comerce/helper"
	"project/e-comerce/middlewares"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productBusiness products.Business
}

func NewProductHandler(business products.Business) *ProductHandler {
	return &ProductHandler{
		productBusiness: business,
	}
}

func (h *ProductHandler) GetAll(c echo.Context) error {
	result, err := h.productBusiness.GetAllProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _response_product.FromCoreList(result),
	})
}

func (h *ProductHandler) GetDataById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := h.productBusiness.GetProductByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _response_product.FromCore(result),
	})
}

func (h *ProductHandler) InsertData(c echo.Context) error {
	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get user id",
		})
	}

	product := _request_product.Product{}
	err_bind := c.Bind(&product)
	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to bind insert data",
		})
	}

	fileData, fileInfo, fileErr := c.Request().FormFile("file")
	if fileErr == http.ErrMissingFile || fileErr != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get file"))
	}

	extension, err_check_extension := helper.CheckFileExtension(fileInfo.Filename)
	if err_check_extension != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("file extension error"))
	}

	// check file size
	err_check_size := helper.CheckFileSize(fileInfo.Size)
	if err_check_size != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("file size error"))
	}

	// memberikan nama file
	fileName := strconv.Itoa(userID_token) + "_" + product.Name + time.Now().Format("2006-01-02 15:04:05") + "." + extension

	url, errUploadImg := helper.UploadImageToS3(fileName, fileData)

	if errUploadImg != nil {
		fmt.Println(errUploadImg)
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to upload file"))
	}

	productCore := _request_product.ToCore(product)
	productCore.UserID = userID_token
	productCore.Photo = fileInfo.Filename
	productCore.PhotoUrl = url

	err := h.productBusiness.InsertProduct(productCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success insert data",
	})

}

func formFile() {

}

func (h *ProductHandler) DeleteData(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get user id",
		})
	}

	err := h.productBusiness.DeleteProductByID(id, userID_token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to delete data" + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete data",
	})
}

func (h *ProductHandler) UpdateData(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	productReq := _request_product.Product{}
	err_bind := c.Bind(&productReq)
	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to bind update data",
		})
	}

	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get user id",
		})
	}

	productCore := _request_product.ToCore(productReq)

	fileData, fileInfo, fileErr := c.Request().FormFile("file")
	if fileErr != http.ErrMissingFile {
		if fileErr != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get file"))
		}

		extension, err_check_extension := helper.CheckFileExtension(fileInfo.Filename)
		if err_check_extension != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("file extension error"))
		}

		// check file size
		err_check_size := helper.CheckFileSize(fileInfo.Size)
		if err_check_size != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("file size error"))
		}

		// memberikan nama file
		fileName := strconv.Itoa(userID_token) + "_" + productReq.Name + time.Now().Format("2006-01-02 15:04:05") + "." + extension

		url, errUploadImg := helper.UploadImageToS3(fileName, fileData)

		if errUploadImg != nil {
			fmt.Println(errUploadImg)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to upload file"))
		}

		productCore.PhotoUrl = url
	}

	err := h.productBusiness.UpdateProductByID(productCore, id, userID_token)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed update data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("sucsess update product"))
}

func (h *ProductHandler) GetProductByUser(c echo.Context) error {
	id_user, errToken := middlewares.ExtractToken(c)

	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get id user",
		})
	}
	result, err := h.productBusiness.GetProductByUserID(id_user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _response_product.FromCoreList(result),
	})
}
