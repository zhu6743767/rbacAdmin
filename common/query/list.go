package query

import (
	"fmt"
	"rbacAdmin/global"
	"rbacAdmin/models"

	"gorm.io/gorm"
)

type Option struct {
	Page     models.Page
	Likes    []string
	Where    *gorm.DB
	Debug    bool
	Preloads []string
}

func List[T any](model T, option Option) (list []T, count int64, err error) {

	// 带入model
	baseDB := global.DB.Model(&model).Where(model)
	// 带入Debug
	if option.Where != nil {
		baseDB = baseDB.Where(option.Where)
	}
	// 带入Where
	if option.Where != nil {
		baseDB = baseDB.Where(option.Where)
	}
	// 带入Likes
	if option.Page.Key != "" && len(option.Likes) != 0 {
		query := global.DB.Where("")
		for _, column := range option.Likes {
			query.Or(fmt.Sprintf("%s LIKE ?", column), fmt.Sprintf("%%%s%%", option.Page.Key))
		}
		baseDB = baseDB.Where(query)
	}
	// 带入Preloads
	for _, preload := range option.Preloads {
		baseDB.Preload(preload)
	}

	if option.Page.Limit <= 0 {
		option.Page.Limit = 10
	}
	if option.Page.Page <= 0 {
		option.Page.Page = 1
	}

	offset := (option.Page.Page - 1) * option.Page.Limit

	if option.Page.Sort == "" {
		option.Page.Sort = "created_at desc"
	}

	baseDB.Limit(option.Page.Limit).Offset(offset).Order(option.Page.Sort).Find(&list)
	baseDB.Model(&model).Count(&count)
	return
}
