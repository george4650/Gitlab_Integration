package gitlab_api_client

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"strconv"
	"wh-hard01.kol.wb.ru/wh-tech/wh-tech-back/wh_tech_gitlab_api/internal/models"
	client "wh-hard01.kol.wb.ru/wh_core/gocore_http"
	"wh-hard01.kol.wb.ru/wh_core/gocore_service_core/configs"
)

type GitLabApiClient struct {
	fastHttpClient *client.HttpClient
}

const (
	apiVersionPath = "api/v4"
)

func NewGitLabApiClient(ctx context.Context, config configs.Config) *GitLabApiClient {
	fastHttpClient := client.NewHttpClient()
	fastHttpClient.InitHttpClient("gitlab_api_client")(ctx, config)

	return &GitLabApiClient{
		fastHttpClient: fastHttpClient,
	}
}

func (r *GitLabApiClient) Registration(fullName, userName, email, password string) (*models.AllUserInfo, error) {
	api := fmt.Sprintf("%s/users", apiVersionPath)

	userOpt := models.CreateUserOptions{
		Email:            email,
		Name:             fullName,
		Username:         userName,
		Password:         password,
		CanCreateGroup:   false,
		SkipConfirmation: false, // пропустить подтверждение админом учётной записи
	}

	resp, err := r.fastHttpClient.HTTPRequest(nil, http.MethodPost, &userOpt, api, nil)
	if err != nil {
		return nil, r.fastHttpClient.GenerateError(resp, err, api)
	}

	if resp.StatusCode() == http.StatusCreated {
		user := models.User{}
		err = jsoniter.Unmarshal(resp.Body(), &user)
		if err != nil {
			return nil, fmt.Errorf("[Registration] Unmarshal body error - %w", err)
		}
		return &models.AllUserInfo{MainInfo: userOpt, OtherInfo: user}, nil
	}

	return nil, r.fastHttpClient.GenerateError(resp, err, api)
}

func (r *GitLabApiClient) ListUsers(request models.GetUsersRequest) ([]models.User, error) {
	api := fmt.Sprintf("%s/users", apiVersionPath)

	params := map[string]string{
		"page": strconv.Itoa(request.Page),
	}
	if request.Username != nil {
		params["username"] = *request.Username
	}
	if request.Search != nil {
		params["search"] = *request.Search
	}
	if request.OrderBy != nil {
		params["order_by"] = *request.OrderBy
	}
	if request.Sort != nil {
		params["sort"] = *request.Sort
	}
	if request.Blocked != nil {
		params["blocked"] = strconv.FormatBool(*request.Blocked)
	}
	if request.Admins != nil {
		params["admins"] = strconv.FormatBool(*request.Admins)
	}

	resp, err := r.fastHttpClient.HTTPRequest(nil, http.MethodGet, nil, api, params)
	if err != nil {
		return nil, r.fastHttpClient.GenerateError(resp, err, api)
	}

	if resp.StatusCode() == http.StatusOK {
		users := []models.User{}
		err = jsoniter.Unmarshal(resp.Body(), &users)
		if err != nil {
			return nil, fmt.Errorf("[ListUsers] Unmarshal body error - %w", err)
		}
		return users, nil
	}
	return nil, r.fastHttpClient.GenerateError(resp, err, api)
}

func (r *GitLabApiClient) GetUser(userID int64) (*models.User, error) {
	api := fmt.Sprintf("%s/users/%d", apiVersionPath, userID)

	resp, err := r.fastHttpClient.HTTPRequest(nil, http.MethodGet, nil, api, nil)
	if err != nil {
		return nil, r.fastHttpClient.GenerateError(resp, err, api)
	}

	if resp.StatusCode() == http.StatusOK {
		user := &models.User{}
		err = jsoniter.Unmarshal(resp.Body(), user)
		if err != nil {
			return nil, fmt.Errorf("[GetUser] Unmarshal body error - %w", err)
		}
		return user, nil
	}
	return nil, r.fastHttpClient.GenerateError(resp, err, api)
}

func (r *GitLabApiClient) BlockUser(userID int64) error {
	api := fmt.Sprintf("%s/users/%d/block", apiVersionPath, userID)

	resp, err := r.fastHttpClient.HTTPRequest(nil, http.MethodPost, nil, api, nil)
	if err != nil {
		return r.fastHttpClient.GenerateError(resp, err, api)
	}

	if resp.StatusCode() == http.StatusCreated {
		return nil
	}
	return r.fastHttpClient.GenerateError(resp, err, api)
}

// Проверка валиден ли токен
func (r *GitLabApiClient) CheckValidToken() error {
	api := "check_token"

	resp, err := r.fastHttpClient.HTTPRequest(nil, http.MethodGet, nil, api, nil)
	if err != nil {
		return r.fastHttpClient.GenerateError(resp, err, api)
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	}

	return r.fastHttpClient.GenerateError(resp, err, api)
}
