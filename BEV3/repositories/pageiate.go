package repositories

import (
	"math"
	"ocg-be/models"
)

type Entity interface {
	Count(search string) int64
	Take(limit int, offset int, orderBy string, sort string, search string) interface{}
}

func Pageniate(entity Entity, page int, limit int, orderBy string, sort string, search string) map[string]interface{} {

	offset := (page - 1) * limit

	data := entity.Take(limit, offset, orderBy, sort, search)
	total := entity.Count(search)

	return models.Map{"data": data,
		"meta": models.Map{
			"total":     total,
			"limit":     limit,
			"page":      page,
			"last_page": int(math.Ceil(float64(total) / float64(limit)))}}
}
