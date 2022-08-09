package helper

import (
	"go-course-basic/restfulapi/model/domain"
	"go-course-basic/restfulapi/model/web"
)

func ToCategoryReponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryReponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryReponse(category))
	}
	return categoryResponses
}
