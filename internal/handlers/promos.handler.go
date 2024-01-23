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

type HandlerPromos struct {
	repositories.IPromosRepository
}

func InitializeHandlerPromos(r repositories.IPromosRepository) *HandlerPromos {
	return &HandlerPromos{r}
}

func (h *HandlerPromos) GetAllPromos(ctx *gin.Context) {
	// var body *models.QueryParamsPromo
	page, _ := ctx.GetQuery("page")
	limit, _ := ctx.GetQuery("limit")

	// if returnPage || returnLimit {
		result, err := h.RepositoryGetAllPromos(page, limit)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		if len(result) == 0 {
			ctx.JSON(http.StatusNotFound, helpers.GetResponse("promo not found", nil, nil))
			return
		}

		count, err := h.RepositoryCountPromos()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		totalData, _ := strconv.Atoi(count[0])
		resultLimit, _ := strconv.Atoi(limit)
		resultPage, _ := strconv.Atoi(page)
		isLastPage := math.Ceil(float64(totalData) / float64(resultLimit))
		resultIsLastPage := int(isLastPage) <= resultPage
		
		linkNext := fmt.Sprintf("%s?page=%d&limit=%d", ctx.Request.URL.Path, resultPage + 1, resultLimit)

		linkPrev := fmt.Sprintf("%s?page=%d&limit=%d", ctx.Request.URL.Path, resultPage - 1, resultLimit)

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

		meta := helpers.GetPagination(resultPage, totalData, isNext, isPrev)
		ctx.JSON(http.StatusOK, helpers.GetResponse("get promo success", result, &meta))
		return
	// }

	
	// result, err := h.RepsitoryGetAllPromos()
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, err)
	// 	return
	// }

	// if len(result) == 0 {
	// 	ctx.JSON(http.StatusNotFound, helpers.GetResponse("promo not found", nil, nil))
	// 	return
	// }

	// ctx.JSON(http.StatusOK, helpers.GetResponse("get all promo success", result, nil))
}

func (h *HandlerPromos) CreateProomos(ctx *gin.Context) {
	var body models.PromosModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err := h.RepositoryCreatePromos(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, helpers.GetResponse("create promo success", nil, nil))
}

func (h *HandlerPromos) UpdatePromos(ctx *gin.Context) {
	id := ctx.Param("id")

	var body models.UpdatePromosModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	result, errs := h.RepositoryGetPromosById(id)
	if errs != nil {
		ctx.JSON(http.StatusInternalServerError, errs)
		return
	}

	// cek partial
	if body.Promos_name == "" {
		body.Promos_name = result[0].Promos_name
	}
	if body.Promos_start == "" {
		body.Promos_start = result[0].Promos_start
	}
	if body.Promos_end == "" {
		body.Promos_end = result[0].Promos_end
	}
	
	err := h.RepositoryUpdatePromos(&body, id)
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, helpers.GetResponse("update promo success", nil, nil))
}

func (h *HandlerPromos) DeletePromos(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := h.RepositoryDeletePromos(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		fmt.Println(err)
		return
	}
	
	if rows := res; rows == 0 {
		ctx.JSON(http.StatusNotFound, helpers.GetResponse("id promo not found", nil, nil))
		return
	}
	ctx.JSON(http.StatusOK, helpers.GetResponse("delete promo success", nil, nil))
}

func (h *HandlerPromos) GetPromosById(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.RepositoryGetPromosById(id)
	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, helpers.GetResponse("promo not found", nil, nil))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, helpers.GetResponse("get promo by id success", result, nil))
}