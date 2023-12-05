package handlers

import (
	"coffee-shop-golang/internal/helpers"
	"coffee-shop-golang/internal/models"
	"coffee-shop-golang/internal/repositories"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type HandlerProducts struct {
	*repositories.ProductsRepository
}


func InitializeHandlerProducts(r *repositories.ProductsRepository) *HandlerProducts {
	return &HandlerProducts{r}
}


func (h *HandlerProducts) GetAllProducts(ctx *gin.Context) {
	name, returnName := ctx.GetQuery("name")
	category, returnCategory := ctx.GetQuery("category")
	minrange, returnMinrange := ctx.GetQuery("minrange")
	maxrange, returnMaxrange := ctx.GetQuery("maxrange")
	page, returnPage := ctx.GetQuery("page")
	limit, returnLimit := ctx.GetQuery("limit")
	sort, returnSort := ctx.GetQuery("sort")

	if returnName || returnCategory || returnMinrange || returnMaxrange || returnPage || returnLimit || returnSort {
		
		if returnMinrange && returnMaxrange {
			if minrange >= maxrange {
				ctx.JSON(http.StatusBadRequest, helpers.GetResponse("The range your input is not correct", nil))
				return
			}
		}

		result, err := h.RepsitoryGetFilterProducts(name, category, minrange, maxrange, page, limit, sort)

		if len(result) == 0 {
			ctx.JSON(http.StatusNotFound, helpers.GetResponse("product not found", nil))
			return
		}

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		count, _ := h.RepositoryCountProducts(name, category, minrange, maxrange)

		totalData, _ := strconv.Atoi(count[0])
		resultLimit, _ := strconv.Atoi(limit)
		resultPage, _ := strconv.Atoi(page)
		isLastPage := math.Ceil(float64(totalData) / float64(resultLimit))
		resultIsLastPage := int(isLastPage) <= resultPage
		
		linkNext := fmt.Sprintf("%s?page=%d&limit=%d", ctx.Request.URL.Path, resultPage + 1, resultLimit) 
		if returnName {
			linkNext = fmt.Sprintf("%s&name=%s", linkNext, name)
		}
		if returnCategory {
			linkNext = fmt.Sprintf("%s&category=%s", linkNext, category)
		}
		if returnMinrange && returnMaxrange {
			linkNext = fmt.Sprintf("%s&minrange=%s&maxrange=%s", linkNext, minrange, maxrange)
		}
		if returnSort {
			linkNext = fmt.Sprintf("%s&sort=%s", linkNext, sort)
		}

		linkPrev := fmt.Sprintf("%s?page=%d&limit=%d", ctx.Request.URL.Path, resultPage - 1, resultLimit)
		if returnName {
			linkPrev = fmt.Sprintf("%s&name=%s", linkPrev, name)
		}
		if returnCategory {
			linkPrev = fmt.Sprintf("%s&category=%s", linkPrev, category)
		}
		if returnMinrange && returnMaxrange {
			linkPrev = fmt.Sprintf("%s&minrange=%s&maxrange=%s", linkPrev, minrange, maxrange)
		}
		if returnSort {
			linkPrev = fmt.Sprintf("%s&sort=%s", linkPrev, sort)
		}

		var isNext string
		var isPrev string

		if resultIsLastPage {
			isNext = "null"
		} else {
			isNext = linkNext
		}

		if resultPage == 1 || resultPage == 0 {
			isPrev = "null"
		} else {
			isPrev = linkPrev
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "get product success",
			"result": result,
			"meta": gin.H{
				"page": resultPage,
				"totalData": totalData,
				"next": isNext,
				"prev": isPrev,
			},
		})
		return
	}

	result, err := h.RepositoryGetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, helpers.GetResponse("get all product success", result))
}

func (h *HandlerProducts) GetProductsById(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.RepositoryProductsById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, helpers.GetResponse("get product by id success", result))
}

func (h *HandlerProducts) CreateProducts(ctx *gin.Context) {
	var body models.ProductsModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if _, err := govalidator.ValidateStruct(body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	id, err := h.RepositoryCreateProducts(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	cld, err := helpers.InitCloudinary()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.GetResponse(err.Error(), nil))
		return
	}

	fieldName := "products_image"
	formFile, err := ctx.FormFile(fieldName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.GetResponse(err.Error(), nil))
		return
	}

	file, err := formFile.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.GetResponse(err.Error(), nil))
		return
	}
	defer file.Close()
	
	publicId := fmt.Sprintf("%s_%s-%s", "product", fieldName, strconv.Itoa(id))
	folder := ""
	res, errs := cld.Uploader(ctx, file, publicId, folder)

	if errs != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.GetResponse(errs.Error(), nil))
		return
	}

	errUpdate := h.RepositoryUpdateImgProducts(res.SecureURL, strconv.Itoa(id))
	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.GetResponse(errUpdate.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, helpers.GetResponse("create product success", gin.H{ "url": res.SecureURL }))
}

func (h *HandlerProducts) UpdateProducts(ctx *gin.Context) {
	id := ctx.Param("id")

	var body models.UpdateProductsModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	result, err := h.RepositoryProductsById(id)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	// cek partial
	if body.Products_name == "" {
		body.Products_name = result[0].Products_name
	}
	if body.Products_price == "" {
		body.Products_price = result[0].Products_price
	}
	if body.Products_desc == "" {
		body.Products_desc = result[0].Products_desc
	}
	if body.Products_stock == "" {
		body.Products_stock = result[0].Products_stock
	}
	if body.Products_image == "" {
		body.Products_image = result[0].Products_image
	}
	if body.Categories_id == "" {
		body.Categories_id = result[0].Categories_id
	}

	errs := h.RepositoryUpdateProducts(&body, id)
	if errs != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	
	cld, err := helpers.InitCloudinary()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.GetResponse(err.Error(), nil))
		return
	}

	fieldName := "products_image"
	formFile, err := ctx.FormFile(fieldName)

	if formFile == nil {
		ctx.JSON(http.StatusOK, helpers.GetResponse("update product success", nil))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.GetResponse(err.Error(), nil))
		return
	}

	file, err := formFile.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.GetResponse(err.Error(), nil))
		return
	}
	defer file.Close()
	
	publicId := fmt.Sprintf("%s_%s-%s", "product", fieldName, id)
	folder := ""
	res, errs := cld.Uploader(ctx, file, publicId, folder)

	if errs != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.GetResponse(err.Error(), nil))
		return
	}

	errUpdate := h.RepositoryUpdateImgProducts(res.SecureURL, id)
	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": errUpdate.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, helpers.GetResponse("update product success", gin.H{ "url": res.SecureURL }))
}

func (h *HandlerProducts) DeleteProducts(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := h.RepositoryDeleteProducts(id)

	if rows, _ := res.RowsAffected(); rows == 0 {
		ctx.JSON(http.StatusNotFound, helpers.GetResponse("id product not found", nil))
		return
	}

	if err != nil {
		pgErr, _ := err.(*pq.Error)
		if pgErr.Code == "23503" {
			ctx.JSON(http.StatusInternalServerError, helpers.GetResponse("error constraint", nil))
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, helpers.GetResponse("delete product success", nil))
}