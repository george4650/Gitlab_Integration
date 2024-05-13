package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math/rand"
	"strings"
	gitlab_api_client "wh-hard01.kol.wb.ru/wh-tech/wh-tech-back/wh_tech_gitlab_api/client"
	"wh-hard01.kol.wb.ru/wh-tech/wh-tech-back/wh_tech_gitlab_api/internal/common"
	"wh-hard01.kol.wb.ru/wh-tech/wh-tech-back/wh_tech_gitlab_api/internal/constant"
	"wh-hard01.kol.wb.ru/wh-tech/wh-tech-back/wh_tech_gitlab_api/internal/models"
	"wh-hard01.kol.wb.ru/wh_core/gocore_rest_auto_api/validators"
	"wh-hard01.kol.wb.ru/wh_core/gocore_service_core/configs"
	utils "wh-hard01.kol.wb.ru/wh_core/gocore_utils"
)

type gitLabApiService struct {
	gitLabApiClient *gitlab_api_client.GitLabApiClient
}

func NewGitLabApiService() GitLabApiService {
	return &gitLabApiService{}
}

type GitLabApiService interface {
	Registration(c *gin.Context, params map[string]interface{})
	BlockUser(c *gin.Context, params map[string]interface{})
	GetUsers(c *gin.Context, params map[string]interface{})
	GetUser(c *gin.Context, params map[string]interface{})
	CheckValidToken(c *gin.Context, params map[string]interface{})
	Init(ctx context.Context, config configs.Config)
}

func (g *gitLabApiService) Init(ctx context.Context, config configs.Config) {
	g.gitLabApiClient = gitlab_api_client.NewGitLabApiClient(ctx, config)
}

func (g *gitLabApiService) Registration(c *gin.Context, params map[string]interface{}) {
	body, ok := params[validators.BodyValidatorBODY].(*struct {
		FirstName string `json:"first_name" binding:"required,max=50"`
		LastName  string `json:"last_name" binding:"required,max=50"`
		UserName  string `json:"user_name" binding:"required,max=110"`
		Email     string `json:"email" binding:"required,email"`
	})
	if !ok {
		utils.BindValidationErrorWithAbort(c, "BODY validation error")
		return
	}

	if !common.Rrus.MatchString(body.FirstName) {
		utils.BindValidationErrorWithAbort(c, fmt.Sprintf("Ошибка валидации: имя должно быть написано кириллицей - '%s'", body.FirstName))
		return
	}
	if !common.Rrus.MatchString(body.LastName) {
		utils.BindValidationErrorWithAbort(c, fmt.Sprintf("Ошибка валидации: фамилия должна быть написана кириллицей - '%s'", body.LastName))
		return
	}
	if !common.RgitUserName.MatchString(body.UserName) {
		utils.BindValidationErrorWithAbort(c, fmt.Sprintf("Ошибка валидации: неверный формат для username - '%s'", body.UserName))
		return
	}
	if !common.RgitEmail.MatchString(body.Email) {
		utils.BindValidationErrorWithAbort(c, fmt.Sprintf("Ошибка валидации: неверный формат для email - '%s'", body.Email))
		return
	}

	fullName := cases.Title(language.Russian).String(fmt.Sprintf(constant.FullName, body.FirstName, body.LastName))

	addUser, err := g.gitLabApiClient.Registration(fullName, body.UserName, body.Email, generateRandomPassword())
	if err != nil {
		utils.BindServiceErrorWithAbort(c, "create user error", fmt.Errorf("Не удалось создать пользователя: %w", err))
		return
	}

	utils.BindObjectToRestData(c, &addUser)
}

func (g *gitLabApiService) GetUsers(c *gin.Context, _ map[string]interface{}) {
	var request models.GetUsersRequest
	if err := c.ShouldBind(&request); err != nil {
		utils.BindValidationErrorWithAbort(c, err.Error())
		return
	}

	users, err := g.gitLabApiClient.ListUsers(request)
	if err != nil {
		utils.BindServiceErrorWithAbort(c, "list users error", fmt.Errorf("Не удалось получить список пользователей: %w", err))
		return
	}

	utils.BindObjectToRestData(c, &users)
}

func (g *gitLabApiService) GetUser(c *gin.Context, params map[string]interface{}) {
	userID, ok := params["id"].(int64)
	if !ok {
		utils.BindValidationErrorWithAbort(c, "param 'id' not set or wrong")
		return
	}

	user, err := g.gitLabApiClient.GetUser(userID)
	if err != nil {
		utils.BindServiceErrorWithAbort(c, "get user error", fmt.Errorf("Не удалось получить пользователя по ID: %w", err))
		return
	}

	utils.BindObjectToRestData(c, &user)
}

func (g *gitLabApiService) BlockUser(c *gin.Context, params map[string]interface{}) {
	body, ok := params[validators.BodyValidatorBODY].(*struct {
		ID int64 `json:"id" binding:"required,gt=0,lte=10000"`
	})
	if !ok {
		utils.BindValidationErrorWithAbort(c, "BODY validation error")
		return
	}

	err := g.gitLabApiClient.BlockUser(body.ID)
	if err != nil {
		utils.BindServiceErrorWithAbort(c, "block user error", fmt.Errorf("Не удалось заблокировать пользователя: %w", err))
		return
	}

	utils.BindNoContent(c)
}

func (g *gitLabApiService) CheckValidToken(c *gin.Context, _ map[string]interface{}) {
	err := g.gitLabApiClient.CheckValidToken()
	if err != nil {
		utils.BindServiceErrorWithAbort(c, "token is not valid", fmt.Errorf("Access токен не валиден: %w", err))
		return
	}

	utils.BindNoContent(c)
}

func generateRandomPassword() string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz")
	length := 16
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
