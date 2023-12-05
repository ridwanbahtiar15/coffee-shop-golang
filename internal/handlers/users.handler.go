package handlers

import (
	"coffee-shop-golang/internal/helpers"
	"coffee-shop-golang/internal/models"
	"coffee-shop-golang/internal/repositories"
	"coffee-shop-golang/pkg"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
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
		return
	}

	if _, err := govalidator.ValidateStruct(body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	i := pkg.InitHashConfig().UseDefaultConfig()
	hashedPassword, err := i.GenHashedPassword(body.Users_password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	id, errs := h.RepositoryCreateUsers(&body, hashedPassword)
	if errs != nil {
		pgErr, _ := errs.(*pq.Error)
		if pgErr.Code == "23505" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "email or phone alredy registered",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errs)
		return
	}

	cld, err := helpers.InitCloudinary()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	fieldName := "users_image"
	formFile, err := ctx.FormFile(fieldName)

	urlImage := "Profile.jpg"
	if formFile == nil {
		errUpdate := h.RepositoryUpdateImgUsers(urlImage, strconv.Itoa(id))
		if errUpdate != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": errUpdate.Error(),
		})
		return
	}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "create user success",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	file, err := formFile.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer file.Close()
	
	publicId := fmt.Sprintf("%s_%s-%s", "user-profile", fieldName, strconv.Itoa(id))
	folder := ""
	res, errs := cld.Uploader(ctx, file, publicId, folder)
	urlImage = res.SecureURL

	if errs != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": errs.Error(),
		})
		return
	}

	errUpdate := h.RepositoryUpdateImgUsers(urlImage, strconv.Itoa(id))
	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": errUpdate.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "create user success",
		"data": gin.H{
			"url": urlImage,
		},
	})
}

func (h *HandlerUsers) UpdateUsers(ctx *gin.Context) {
	id := ctx.Param("id")

	var body models.UpdateUserModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if _, err := govalidator.ValidateStruct(body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
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

	i := pkg.InitHashConfig().UseDefaultConfig()
	hashedPassword, err := i.GenHashedPassword(body.Users_password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	errs := h.RepositoryUpdateUsers(&body, hashedPassword, id)
	fmt.Println(errs)
	if errs != nil {
		pgErr, _ := errs.(*pq.Error)
		if pgErr.Code == "23505" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "phone alredy registered",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	cld, err := helpers.InitCloudinary()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	fieldName := "users_image"
	formFile, err := ctx.FormFile(fieldName)

	if formFile == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "update user success",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	file, err := formFile.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer file.Close()
	
	publicId := fmt.Sprintf("%s_%s-%s", "user-profile", fieldName, id)
	folder := ""
	res, errs := cld.Uploader(ctx, file, publicId, folder)

	if errs != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": errs.Error(),
		})
		return
	}

	errUpdate := h.RepositoryUpdateImgUsers(res.SecureURL, id)
	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": errUpdate.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "update user success",
		"data": gin.H{
			"url": res.SecureURL,
		},
	})
}

func (h *HandlerUsers) DeleteUsers(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := h.RepositoryDeleteUsers(id)

	if rows, _ := res.RowsAffected(); rows == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "id user not found",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete user success",
	})
}

func (h *HandlerUsers) UserProfile(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	token := strings.Replace(bearerToken, "Bearer ", "", -1)
	payload, _ := pkg.VerifyToken(token)
	id := payload.Users_id

	var body models.UpdateUserModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if _, err := govalidator.ValidateStruct(body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	result, err := h.RepositoryUsersById(id)
	if err != nil {
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

	i := pkg.InitHashConfig().UseDefaultConfig()
	hashedPassword, err := i.GenHashedPassword(body.Users_password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	errs := h.RepositoryUpdateUsers(&body, hashedPassword, id)
	if errs != nil {
		pgErr, _ := errs.(*pq.Error)
		if pgErr.Code == "23505" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "phone alredy registered",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	cld, err := helpers.InitCloudinary()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	fieldName := "users_image"
	formFile, err := ctx.FormFile(fieldName)

	if formFile == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "update user success",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	file, err := formFile.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer file.Close()
	
	publicId := fmt.Sprintf("%s_%s-%s", "user-profile", fieldName, id)
	folder := ""
	res, errs := cld.Uploader(ctx, file, publicId, folder)

	if errs != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": errs.Error(),
		})
		return
	}

	errUpdate := h.RepositoryUpdateImgUsers(res.SecureURL, id)
	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": errUpdate.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "update user success",
		"data": gin.H{
			"url": res.SecureURL,
		},
	})
}