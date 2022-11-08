package gorm_generics

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type GormModel[E any] interface {
	ToEntity() E
	FromEntity(entity E) interface{}
}

func NewRepository[M GormModel[E], E any](db *gorm.DB) *GormRepository[M, E] {
	return &GormRepository[M, E]{
		db: db,
	}
}

type GormRepository[M GormModel[E], E any] struct {
	db *gorm.DB
}

func (r *GormRepository[M, E]) Insert(ctx context.Context, entity *E) error {
	var start M
	model := start.FromEntity(*entity).(M)

	err := r.db.WithContext(ctx).Create(&model).Error
	if err != nil {
		return err
	}

	*entity = model.ToEntity()
	return nil
}

func (r *GormRepository[M, E]) FindByID(ctx context.Context, id uint) (E, error) {
	var model M
	err := r.db.WithContext(ctx).First(&model, id).Error
	if err != nil {
		return *new(E), err
	}

	return model.ToEntity(), nil
}

func (r *GormRepository[M, E]) Find(ctx context.Context, specification Specification) ([]E, error) {
	var models []M
	err := r.db.WithContext(ctx).Where(specification.GetQuery(), specification.GetValues()...).Find(&models).Error
	if err != nil {
		return nil, err
	}

	result := make([]E, 0, len(models))
	for _, row := range models {
		result = append(result, row.ToEntity())
	}

	return result, nil
}

func (r *GormRepository[M, E]) Update(ctx context.Context, entity *E) error {
	var start M
	model := start.FromEntity(*entity).(M)

	err := r.db.WithContext(ctx).Updates(&model).Error
	if err != nil {
		return err
	}

	*entity = model.ToEntity()
	return nil
}

type PaginationQuery struct {
	Page     int                       `form:"page,default=1"`
	PageSize int                       `form:"pageSize,default=10"`
	OrderBy  string                    `form:"orderBy,default=id"`
	Order    string                    `form:"order,default=DESC" binding:"oneof='DESC' 'ASC' 'asc' 'desc'"`
	Filter   map[string]interface{}    `form:"filter"`
	SearchIn map[string]interface{}    `form:"searchIn"`
	Search   map[string]interface{}    `form:"search"`
	SearchOr map[string]interface{}    `form:"searchOr"`
	Between  map[string][2]interface{} `form:"between"`
}


//PaginationResponse schema for paginated responses
type PaginationResponse struct {
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}


//Paginate add pagination functions to existing db
func Paginate(query *PaginationQuery, tx *gorm.DB, result interface{}) (response PaginationResponse, err error) {
	countQuery := tx.WithContext(context.Background())
	offset := (query.Page - 1) * query.PageSize

	pageQuery := tx.Order(fmt.Sprintf(`"%s" %s`, query.OrderBy, query.Order)).
		Limit(query.PageSize).
		Offset(offset)

	err = pageQuery.Find(result).Error

	if err != nil {
		return response, err
	}

	countQuery.Statement.Preloads = map[string][]interface{}{}
	var count int64
	err = GormRepository.db.Table("(?) as results", countQuery).Count(&count).Error

	response = PaginationResponse{
		Data:  result,
		Total: count,
	}

	return response, err
}
