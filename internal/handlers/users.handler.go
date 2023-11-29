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

type HandlerUsers struct {
	*repositories.UsersRepository
}

func InitializeHandlerUsers(r *repositories.UsersRepository) *HandlerUsers {
	return &HandlerUsers{r}
}

func (h *HandlerUsers) GetAllUsers(ctx *gin.Context) {
	name, returnName := ctx.GetQuery("name")
	page, returnPage := ctx.GetQuery("page")
	limit, returnLimit := ctx.GetQuery("limit")
	sort, returnSort := ctx.GetQuery("sort")
	
	if returnName || returnPage || returnLimit || returnSort {
		result, err := h.RepositoryGetFilterUsers(name, page, limit, sort)

		if len(result) == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
			})
			return
		}

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		count, err := h.RepositoryCountUsers(name)
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

		data := models.MetaUsers{}
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
	

	result, err := h.RepositoryGetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get all user success",
		"data": result,
	})
}

func (h *HandlerUsers) GetUsersById(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.RepositoryUsersById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get user by id success",
		"data": result,
	})
}

func (h *HandlerUsers) CreateUsers(ctx *gin.Context) {
	var body models.UsersModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	err := h.RepositoryCreateUsers(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "create user success",
	})
}

func (h *HandlerUsers) UpdateUsers(ctx *gin.Context) {
	id := ctx.Param("id")

	var body models.UsersModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepositoryUsersById(id)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	// cek partial
	if body.Users_fullname == "" {
		body.Users_fullname = result[0].Users_fullname
	}
	if body.Users_password == "" {
		body.Users_password = result[0].Users_password
	}
	if body.Users_phone == "" {
		body.Users_phone = result[0].Users_phone
	}
	if body.Users_address == "" {
		body.Users_address = result[0].Users_address
	}
	if body.Users_image == "" {
		body.Users_image = result[0].Users_image
	}

	errs := h.RepositoryUpdateUsers(&body, id)
	if errs != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "update user success",
	})
}

func (h *HandlerUsers) DeleteUsers(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.RepositoryDeleteUsers(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete user success",
	})
}