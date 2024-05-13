package common

import (
	"fmt"
	"regexp"
	"wh-hard01.kol.wb.ru/wh-tech/wh-tech-back/wh_tech_gitlab_api/internal/constant"
)

var (
	Rrus         = regexp.MustCompile("^[а-яА-Я]+$")
	RgitUserName = regexp.MustCompile(`^[a-z]+\.[a-z]+[0-9]*$`)
	RgitEmail    = regexp.MustCompile(fmt.Sprintf("^[a-z]+\\.[a-z]+[0-9]*(%s|%s)$", constant.EmailSuffixWbRu, constant.EmailSuffixWildberriesWork))
)
