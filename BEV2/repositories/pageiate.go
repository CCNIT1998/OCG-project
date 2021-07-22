package repositories

import (
	"math"
	"ocg-be/models"
)

type Entity interface {
	Count(category_id int, search string) int64
	TakeByParams(limit int, offset int, categoryId int, orderBy string, sort string, search string) interface{}
}

func Pageniate(entity Entity, page int, limit int, categoryId int, orderBy string, sort string, search string) map[string]interface{} {

	offset := (page - 1) * limit

	data := entity.TakeByParams(limit, offset, categoryId, orderBy, sort, search)
	total := entity.Count(categoryId, search)

	return models.Map{"data": data,
		"meta": models.Map{
			"total":     total,
			"limit":     limit,
			"page":      page,
			"last_page": int(math.Ceil(float64(total) / float64(limit)))}}
}
