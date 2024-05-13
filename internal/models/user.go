package models

import (
	"net"
	"time"
)

type User struct {
	ID                             int             `json:"id"`
	Username                       string          `json:"username"`
	Email                          string          `json:"email"`
	Name                           string          `json:"name"`
	State                          string          `json:"state"`
	Locked                         bool            `json:"locked"`
	AvatarURL                      string          `json:"avatar_url"`
	WebURL                         string          `json:"web_url"`
	CreatedAt                      *time.Time      `json:"created_at"`
	IsAdmin                        bool            `json:"is_admin"`
	Bio                            string          `json:"bio"`
	Location                       string          `json:"location"`
	Skype                          string          `json:"skype"`
	Linkedin                       string          `json:"linkedin"`
	Twitter                        string          `json:"twitter"`
	Discord                        string          `json:"discord"`
	WebsiteURL                     string          `json:"website_url"`
	Organization                   string          `json:"organization"`
	JobTitle                       string          `json:"job_title"`
	LastSignInAt                   *time.Time      `json:"last_sign_in_at"`
	ConfirmedAt                    *time.Time      `json:"confirmed_at"`
	ThemeID                        int             `json:"theme_id"`
	LastActivityOn                 string          `json:"last_activity_on"`
	ColorSchemeID                  int             `json:"color_scheme_id"`
	ProjectsLimit                  int             `json:"projects_limit"`
	CurrentSignInAt                *time.Time      `json:"current_sign_in_at"`
	Note                           string          `json:"note"`
	Identities                     []*UserIdentity `json:"identities"`
	CanCreateGroup                 bool            `json:"can_create_group"`
	CanCreateProject               bool            `json:"can_create_project"`
	TwoFactorEnabled               bool            `json:"two_factor_enabled"`
	External                       bool            `json:"external"`
	PrivateProfile                 bool            `json:"private_profile"`
	CurrentSignInIP                *net.IP         `json:"current_sign_in_ip"`
	LastSignInIP                   *net.IP         `json:"last_sign_in_ip"`
	NamespaceID                    int             `json:"namespace_id"`
	Bot                            bool            `json:"bot"`
	PublicEmail                    string          `json:"public_email"`
	SharedRunnersMinutesLimit      int             `json:"shared_runners_minutes_limit"`
	ExtraSharedRunnersMinutesLimit int             `json:"extra_shared_runners_minutes_limit"`
	UsingLicenseSeat               bool            `json:"using_license_seat"`
}

type UserIdentity struct {
	Provider  string `json:"provider"`
	ExternUID string `json:"extern_uid"`
}

type CreateUserOptions struct {
	Name             string `json:"name"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	CanCreateGroup   bool   `json:"can_create_group"`
	SkipConfirmation bool   `json:"skip_confirmation"`
}

type AllUserInfo struct {
	MainInfo  CreateUserOptions `json:"main_info"`
	OtherInfo User              `json:"other_info"`
}

type GetUsersRequest struct {
	Page     int     `form:"page" binding:"required,gt=0,lte=500"`
	Search   *string `form:"search" binding:"omitempty,max=100"`
	Username *string `form:"username" binding:"omitempty,max=100"`
	OrderBy  *string `form:"order_by" binding:"omitempty,oneof=id name username created_at"`
	Sort     *string `form:"sort" binding:"omitempty,oneof=asc desc"`
	Blocked  *bool   `form:"blocked" binding:"omitempty"`
	Admins   *bool   `form:"admins" binding:"omitempty"`
}
