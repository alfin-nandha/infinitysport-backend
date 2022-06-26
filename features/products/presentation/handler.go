package presentation

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"project/e-comerce/features/products"
	_request_product "project/e-comerce/features/products/presentation/request"
	_response_product "project/e-comerce/features/products/presentation/response"
	"project/e-comerce/middlewares"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct{
	productBusiness products.Business
}

func NewProductHandler(business products.Business) *ProductHandler{
	return &ProductHandler{
		productBusiness: business,
	}
}

func (h *ProductHandler)GetAll(c echo.Context) error{
	result, err := h.productBusiness.GetAllProduct()
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to get all data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "success",
		"data" : _response_product.FromCoreList(result),
	})
}

func (h *ProductHandler)GetDataById(c echo.Context) error{
	id,_ := strconv.Atoi(c.Param("id"))
	result, err:= h.productBusiness.GetProductByID(id)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to get data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "success",
		"data" : _response_product.FromCore(result),
	})
}

func (h *ProductHandler)InsertData(c echo.Context)error{
	userID_token,errToken := middlewares.ExtractToken(c)
	if errToken != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to get user id",
		})
	}
	
	product := _request_product.Product{}
	err_bind := c.Bind(&product)
	if err_bind != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to bind insert data",
		})
	}
	
	file,fileErr := c.FormFile("file")
	if fileErr == http.ErrMissingFile {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"error missing file",
		})
	}else if fileErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed get file",
		})
	}

	filename := path.Base(file.Filename)

	src, errOpen := file.Open()
	fmt.Println(errOpen)
	defer src.Close()

	s3Config := &aws.Config{
		Region : aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("KeyID", "SecretKey", ""),

	}
	s3Session := session.New(s3Config)

    uploader := s3manager.NewUploader(s3Session)

	dest, errCreate := os.Create(filename)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"error upload file",
			
		})
	}
	defer dest.Close()

	if _, errio := io.Copy(dest, src); errio != nil {
		return errio
	}
	fmt.Printf("File %s uploaded successfully", filename,)

	productCore := _request_product.ToCore(product)
	productCore.UserID = userID_token
	err := h.productBusiness.InsertProduct(productCore)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "success insert data",
	})
	
}

func formFile(){

}


func (h *ProductHandler)DeleteData(c echo.Context) error{
	id,_ := strconv.Atoi(c.Param("id"))
	err:= h.productBusiness.DeleteProductByID(id)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to delete data"+err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "success delete data",
	})
}

func (h *ProductHandler)UpdateData(c echo.Context)error{
	id,_ := strconv.Atoi(c.Param("id"))

	productReq := _request_product.Product{}
	err_bind := c.Bind(&productReq)
	if err_bind != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to bind update data",
		})
	}

	productCore := _request_product.ToCore(productReq)
	err:= h.productBusiness.UpdateProductByID(productCore, id)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "success insert data",
	})
}

func (h *ProductHandler)GetProductByUser(c echo.Context) error{
	id_user,errToken := middlewares.ExtractToken(c)
	
	if errToken != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to get id user",
		})
	}
	result, err := h.productBusiness.GetProductByUserID(id_user)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":"failed to get all data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "success",
		"data" : _response_product.FromCoreList(result),
	})
}