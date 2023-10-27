package service

import (
	"blog/app/dto"
)

type articleDTO interface {
	dto.ArticleR
	dto.ArticleW
	dto.TagsR
	dto.TagsW
	dto.CateR
	dto.CateW
}
type userDTO interface {
	dto.UserR
	dto.UserW
	dto.Authorizer
}

type DomainService struct {
	articleDTO
	userDTO
}