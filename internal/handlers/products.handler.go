package handlers

import (
	"coffee-shop-golang/internal/models"
	"coffee-shop-golang/internal/repositories"
	"fmt"
	"net/http"

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

	if returnName || returnCategory || returnMinrange || returnMaxrange || returnPage || returnLimit { 

		result, err := h.RepsitoryGetFilterProducts(name, category, minrange, maxrange, page, limit)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			fmt.Println(err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": result,
			"message": "get all product success",
		})
		return
	}

	result, err := h.RepsitoryGetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "get all product success",
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
		"data": result,
		"message": "get product by id success",
	})
}

func (h *HandlerProducts) CreateProducts(ctx *gin.Context) {
	var body models.ProductsModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepsitoryCreateProducts(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "create product success",
	})
}

func (h *HandlerProducts) UpdateProducts(ctx *gin.Context) {
	id := ctx.Param("id")

	var body models.ProductsModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepsitoryUpdateProducts(&body, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "update product success",
	})
}

func (h *HandlerProducts) DeleteProducts(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.RepositoryDeleteProducts(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "delete product success",
	})
}