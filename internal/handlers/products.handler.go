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
	repositories.IProductsRepository
}


func InitializeHandlerProducts(r repositories.IProductsRepository) *HandlerProducts {
	return &HandlerProducts{r}
}


func (h *HandlerProducts) GetAllProducts(ctx *gin.Context) {
	name, _ := ctx.GetQuery("name")
	category, _ := ctx.GetQuery("category")
	minrange, returnMinrange := ctx.GetQuery("minrange")
	maxrange, returnMaxrange := ctx.GetQuery("maxrange")
	page, _ := ctx.GetQuery("page")
	limit, _ := ctx.GetQuery("limit")
	sort, _ := ctx.GetQuery("sort")

	// if returnName || returnCategory || returnMinrange || returnMaxrange || returnPage || returnLimit || returnSort {
		
		if returnMinrange && returnMaxrange {
			if minrange >= maxrange {
				ctx.JSON(http.StatusBadRequest, helpers.GetResponse("The range your input is not correct", nil, nil))
				return
			}
		}

		if page == "" {
			page = "1"
		}

		if limit == "" {
			limit = "6"
		}

		result, err := h.RepositoryGetAllProducts(name, category, minrange, maxrange, page, limit, sort)

		if len(result) == 0 {
			ctx.JSON(http.StatusNotFound, helpers.GetResponse("product not found", nil, nil))
			return
		}

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		count, _ := h.RepositoryCountProducts(name, category, minrange, maxrange)

		totalData, _ := strconv.Atoi(count[0])
		fmt.Println(totalData)
		resultLimit, _ := strconv.Atoi(limit)
		resultPage, _ := strconv.Atoi(page)
		isLastPage := math.Ceil(float64(totalData) / float64(resultLimit))
		resultIsLastPage := int(isLastPage) <= resultPage
		
		linkNext := fmt.Sprintf("%s?page=%d&limit=%d", ctx.Request.URL.Path, resultPage + 1, resultLimit) 
		if name != "" {
			linkNext = fmt.Sprintf("%s&name=%s", linkNext, name)
		}
		if category != "" {
			linkNext = fmt.Sprintf("%s&category=%s", linkNext, category)
		}
		if returnMinrange && returnMaxrange {
			linkNext = fmt.Sprintf("%s&minrange=%s&maxrange=%s", linkNext, minrange, maxrange)
		}
		if sort != "" {
			linkNext = fmt.Sprintf("%s&sort=%s", linkNext, sort)
		}

		linkPrev := fmt.Sprintf("%s?page=%d&limit=%d", ctx.Request.URL.Path, resultPage - 1, resultLimit)
		if name != "" {
			linkPrev = fmt.Sprintf("%s&name=%s", linkPrev, name)
		}
		if category != "" {
			linkPrev = fmt.Sprintf("%s&category=%s", linkPrev, category)
		}
		if returnMinrange && returnMaxrange {
			linkPrev = fmt.Sprintf("%s&minrange=%s&maxrange=%s", linkPrev, minrange, maxrange)
		}
		if sort != "" {
			linkPrev = fmt.Sprintf("%s&sort=%s", linkPrev, sort)
		}

		isNext := linkNext
		isPrev := linkPrev

		if resultIsLastPage {
			isNext = "null"
		}

		if resultPage == 1 || resultPage == 0 {
			isPrev = "null"
		} else {
			isPrev = linkPrev
		}

		

		meta := helpers.GetPagination(resultPage, totalData, isNext, isPrev)
		ctx.JSON(http.StatusOK, helpers.GetResponse("get product success", result, &meta))
	}

	// result, err := h.RepositoryGetAllProducts(name, category, minrange, maxrange, page, limit, sort)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, err)
	// 	return
	// }
	// ctx.JSON(http.StatusOK, helpers.GetResponse("get all product success", result, nil))
// }

func (h *HandlerProducts) GetProductsById(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.RepositoryProductsById(id)
	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, helpers.GetResponse("product not found", nil, nil))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, helpers.GetResponse("get product by id success", result, nil))
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
	fmt.Println(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	cld, err := helpers.InitCloudinary()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.GetResponse(err.Error(), nil, nil))
		return
	}

	fieldName := "products_image"
	formFile, err := ctx.FormFile(fieldName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.GetResponse(err.Error(), nil, nil))
		return
	}

	file, err := formFile.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.GetResponse(err.Error(), nil, nil))
		return
	}
	defer file.Close()
	
	publicId := fmt.Sprintf("%s_%s-%s", "product", fieldName, strconv.Itoa(id))
	folder := ""
	res, errs := cld.Uploader(ctx, file, publicId, folder)

	if errs != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.GetResponse(errs.Error(), nil, nil))
		return
	}

	errUpdate := h.RepositoryUpdateImgProducts(res.SecureURL, strconv.Itoa(id))
	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.GetResponse(errUpdate.Error(), nil, nil))
		return
	}

	ctx.JSON(http.StatusOK, helpers.GetResponse("create product success", gin.H{ "url": res.SecureURL }, nil))
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
		ctx.JSON(http.StatusInternalServerError, helpers.GetResponse(err.Error(), nil, nil))
		return
	}

	fieldName := "products_image"
	formFile, err := ctx.FormFile(fieldName)

	if formFile == nil {
		ctx.JSON(http.StatusOK, helpers.GetResponse("update product success", nil, nil))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.GetResponse(err.Error(), nil, nil))
		return
	}

	file, err := formFile.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.GetResponse(err.Error(), nil, nil))
		return
	}
	defer file.Close()
	
	publicId := fmt.Sprintf("%s_%s-%s", "product", fieldName, id)
	folder := ""
	res, errs := cld.Uploader(ctx, file, publicId, folder)

	if errs != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.GetResponse(err.Error(), nil, nil))
		return
	}

	errUpdate := h.RepositoryUpdateImgProducts(res.SecureURL, id)
	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": errUpdate.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, helpers.GetResponse("update product success", gin.H{ "url": res.SecureURL }, nil))
}

func (h *HandlerProducts) DeleteProducts(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := h.RepositoryDeleteProducts(id)

	if err != nil {
		pgErr, _ := err.(*pq.Error)
		if pgErr.Code == "23503" {
			ctx.JSON(http.StatusInternalServerError, helpers.GetResponse("error constraint", nil, nil))
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if rows := res; rows == 0 {
		ctx.JSON(http.StatusNotFound, helpers.GetResponse("id product not found", nil, nil))
		return
	}
	ctx.JSON(http.StatusOK, helpers.GetResponse("delete product success", nil, nil))
}