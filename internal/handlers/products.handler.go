package handlers

import (
	"coffee-shop-golang/internal/models"
	"coffee-shop-golang/internal/repositories"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "The range your input is not correct",
				})
				return
			}
		}

		result, err := h.RepsitoryGetFilterProducts(name, category, minrange, maxrange, page, limit, sort)
		fmt.Println(err)

		if len(result) == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "product not found",
			})
			return
		}

		count, err := h.RepositryCountProducts(name, category, minrange, maxrange)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			fmt.Println(err)
			return
		}

		totalData, _ := strconv.Atoi(count[0])
		resultLimit, _ := strconv.Atoi(limit)
		resultPage, _ := strconv.Atoi(page)
		isLastPage := math.Ceil(float64(totalData) / float64(resultLimit))
		resultIsLastPage := int(isLastPage) <= resultPage
		fmt.Println(resultIsLastPage)

		// fmt.Println(ctx.Request.URL.Path)
		
		linkNext := fmt.Sprintf("%s?page=%d&limit=%d", ctx.Request.URL.Path, resultPage + 1, resultLimit) 
		linkPrev := fmt.Sprintf("%s?page=%d&limit=%d", ctx.Request.URL.Path, resultPage - 1, resultLimit) 

		var isNext string
		var isPrev string

		if resultIsLastPage {
			isNext = "null"
		} else {
			isNext = linkNext
		}

		if resultPage == 1 {
			isPrev = "null"
		} else {
			isPrev = linkPrev
		}

		data := models.Meta{}
		data.Page = resultPage
		data.TotalData = totalData
		data.Next = isNext
		data.Prev = isPrev


		ctx.JSON(http.StatusOK, gin.H{
			"message": "get product success",
			"result": result,
			"meta": data,
		})
		return
	}

	result, err := h.RepsitoryGetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get all product success",
		"result": result,
	})
}

func (h *HandlerProducts) GetProductsById(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.RepsitoryProductsById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get product by id success",
		"result": result,
	})
}

func (h *HandlerProducts) CreateProducts(ctx *gin.Context) {
	var body models.ProductsModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	err := h.RepsitoryCreateProducts(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "create product success",
	})
}

func (h *HandlerProducts) UpdateProducts(ctx *gin.Context) {
	id := ctx.Param("id")

	var body models.ProductsModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	err := h.RepsitoryUpdateProducts(&body, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "update product success",
	})
}

func (h *HandlerProducts) DeleteProducts(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.RepositoryDeleteProducts(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete product success",
	})
}