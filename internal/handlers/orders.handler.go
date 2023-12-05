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
)

type HandlerOrders struct {
	*repositories.OrdersRepository
}

func InitializeHandlerOrders(r *repositories.OrdersRepository) *HandlerOrders {
	return &HandlerOrders{r}
}

func (h *HandlerOrders) GetAllOrders(ctx *gin.Context) {
	orderNumber, returnOrderNumber := ctx.GetQuery("orderNumber")
	page, returnPage := ctx.GetQuery("page")
	limit, returnLimit := ctx.GetQuery("limit")
	sort, returnSort := ctx.GetQuery("sort")

	if returnOrderNumber || returnPage || returnLimit || returnSort {
		result, err := h.RepositoryGetFilterOrders(orderNumber, page, limit, sort)
		fmt.Println(err)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		if len(result) == 0 {
			ctx.JSON(http.StatusNotFound, helpers.GetResponse("user not found", nil))
			return
		}

		count, err := h.RepositoryCountOrders(orderNumber)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		totalData, _ := strconv.Atoi(count[0])
		resultLimit, _ := strconv.Atoi(limit)
		resultPage, _ := strconv.Atoi(page)
		isLastPage := math.Ceil(float64(totalData) / float64(resultLimit))
		resultIsLastPage := int(isLastPage) <= resultPage

		fmt.Println(resultPage)
		
		linkNext := fmt.Sprintf("%s?page=%d&limit=%d", ctx.Request.URL.Path, resultPage + 1, resultLimit)
		if returnSort {
			linkNext = fmt.Sprintf("%s&sort=%s", linkNext, sort)
		}

		linkPrev := fmt.Sprintf("%s?page=%d&limit=%d", ctx.Request.URL.Path, resultPage - 1, resultLimit)
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
			"message": "get order success",
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


	result, err := h.RepositoryGetAllOrders()
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, helpers.GetResponse("order not found", nil))
		return
	}

	ctx.JSON(http.StatusOK, helpers.GetResponse("get all order success", result))
}

func (h *HandlerOrders) CreateOrders(ctx *gin.Context) {
	var body models.OrdersModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	err := h.RepositoryCreateOrders(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, helpers.GetResponse("create order success", nil))
}

func (h *HandlerOrders) UpdateOrders(ctx *gin.Context) {
	var body models.OrdersModel
	id := ctx.Param("id")
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err := h.RepositoryUpdateOrders(&body, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, helpers.GetResponse("update order success", nil))
}