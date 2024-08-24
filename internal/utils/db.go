package utils

import "gorm.io/gorm"

type Paginate struct {
	limit int
	page  int
}

func (p *Paginate) PaginatedResult(db *gorm.DB) *gorm.DB {
	offset := (p.page - 1) * p.limit

	return db.Offset(offset).
		Limit(p.limit)
}

func NewPaginate(limit int, page int) *Paginate {
	return &Paginate{limit: limit, page: page}
}
