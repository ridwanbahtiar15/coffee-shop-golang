package handlers

import (
	"coffee-shop-golang/internal/helpers"
	"coffee-shop-golang/internal/models"
	"coffee-shop-golang/internal/repositories"
	"coffee-shop-golang/pkg"
	"fmt"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type HandlerAuth struct {
	*repositories.RepositoryAuth
}

func InitializeHandlerAuth(r *repositories.RepositoryAuth) *HandlerAuth {
	return &HandlerAuth{r}
}

func (h *HandlerAuth) RegisterUsers(ctx *gin.Context) {
	var body models.GetUserInfoModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
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

	errs := h.RepositoryRegisterUsers(&body, hashedPassword)
	if errs != nil {
		pgErr, _ := errs.(*pq.Error)
		if pgErr.Code == "23505" {
			ctx.JSON(http.StatusBadRequest, helpers.GetResponse("email alredy registered", nil, nil))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errs)
		return
	}
	ctx.JSON(http.StatusOK, helpers.GetResponse("register success", nil, nil))
}

func (h *HandlerAuth) LoginUsers(ctx *gin.Context) {
	body := &models.GetUserInfoModel{}
	if err := ctx.ShouldBind(body); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if _, err := govalidator.ValidateStruct(body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	result, err := h.GetUsers(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, helpers.GetResponse("email not registered", nil, nil))
		return
	}

	hc := pkg.HashConfig{}
	isValid, err := hc.ComparePasswordAndHash(body.Users_password, result[0].Users_password)
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !isValid {
		ctx.JSON(http.StatusUnauthorized, helpers.GetResponse("email or password is wrong", nil, nil))
		return
	}

	// generate JWT
	payload := pkg.NewPayload(result[0].Users_id, result[0].Roles_id)
	token, err := payload.GenerateToken()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	message := fmt.Sprintf("Welcome %s", result[0].Users_fullname)
	ctx.JSON(http.StatusOK, helpers.GetResponse(message, gin.H{
		"token": token,
		"userInfo": gin.H{
			"users_id": result[0].Users_id,
			"users_email": result[0].Users_email,
			"users_fullname": result[0].Users_fullname,
			"roles_id": result[0].Roles_id,
		},
	}, nil))

	if errs := h.InsertJwt(result[0].Users_id, token); errs != nil {
		ctx.JSON(http.StatusInternalServerError, errs.Error())
		return
	}
}

func (h *HandlerAuth) LogoutUsers(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	token := strings.Replace(bearerToken, "Bearer ", "", -1)

	if err := h.DeleteJwt(token); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, helpers.GetResponse("logout success", nil, nil))

}