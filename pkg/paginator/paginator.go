package paginator

import (
	"errors"
	"math"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/j23063519/clean_architecture/pkg/log"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Paging
type Paging struct {
	CurrentPage int    `json:"current_page"`
	PerPage     int    `json:"per_page"`
	TotalPage   int    `json:"total_page"`
	TotalCount  int64  `json:"total_count"`
	Sort        string `json:"sort"`  // id, created_at, updated_at
	Order       string `json:"order"` // asc(low to high), desc(high to low)
}

// Paginate
//
// c  - gin.context: c.Query
//
// db - GORM
//
// data - binding data from sql query
//
// PerPage - if nil => then use default value
func Paginate(c *gin.Context, db *gorm.DB, data interface{}) (Paging, error) {
	// 	init Paginator
	p := &Paginator{
		query: db,
		c:     c,
	}
	p.initProperties(0)

	// sql query
	err := p.query.Preload(clause.Associations). // associations
							Order(p.Sort + " " + p.Order). // ex: id asc
							Limit(p.PerPage).
							Offset(p.Offset).
							Find(data).Error

	// err
	if err != nil {
		log.ErrorJSON("分頁", "Paginate", err)
		return Paging{}, err
	}

	return Paging{
		CurrentPage: p.CurrentPage,
		PerPage:     p.PerPage,
		TotalPage:   p.TotalPage,
		TotalCount:  p.TotalCount,
		Sort:        p.Sort,
		Order:       p.Order,
	}, nil
}

// PaginateData
//
// c  - gin.context: c.Query
//
// data - binding data from datas
//
// PerPage - if nil => then use default value
func PaginateData(c *gin.Context, db *gorm.DB, data interface{}) (Paging, error) {
	// reflect data
	vf := reflect.ValueOf(data)

	// if vf is not pointer or not slice then error
	if vf.Kind() != reflect.Ptr || vf.Elem().Kind() != reflect.Slice {
		err := errors.New("data type is invalid")
		log.ErrorJSON("分頁", "PaginateData", err)
		return Paging{}, err
	}

	// if vf is nil then return Paging{} and nil error
	if vf.Elem().Len() < 1 {
		return Paging{}, nil
	}

	// init Paginator
	p := &Paginator{
		query: db,
		c:     c,
	}
	p.initProperties(int64(vf.Elem().Len()))

	// calculate pagination information
	totalRecords := p.TotalCount
	startIndex := (p.CurrentPage - 1) * p.PerPage
	endIndex := startIndex + p.PerPage

	// Prevent index from going out of range
	if startIndex >= int(totalRecords) {
		startIndex = 0
	}
	if endIndex > int(totalRecords) {
		endIndex = int(totalRecords)
	}

	// Use slicing to obtain a specified range of data
	sliceType := reflect.SliceOf(vf.Elem().Type().Elem())
	slicedData := reflect.MakeSlice(sliceType, endIndex-startIndex, endIndex-startIndex)
	reflect.Copy(slicedData, vf.Elem().Slice(startIndex, endIndex))

	// Reflect sliced ​​data back to original data indicators
	vf.Elem().Set(reflect.AppendSlice(reflect.MakeSlice(vf.Elem().Type(), 0, 0), slicedData))

	return Paging{
		CurrentPage: p.CurrentPage,
		PerPage:     p.PerPage,
		TotalPage:   p.TotalPage,
		TotalCount:  p.TotalCount,
		Sort:        p.Sort,
		Order:       p.Order,
	}, nil
}

// implements Pagination
type Paginator struct {
	PerPage     int
	CurrentPage int
	Offset      int
	TotalCount  int64
	TotalPage   int
	Sort        string       // id, created_at, updated_at
	Order       string       // asc(low to hight), desc(high to low)
	query       *gorm.DB     // db query，EX：a.ORM.Model(domain.Admin{})
	c           *gin.Context // gin context
}

func (p *Paginator) initProperties(dataLen int64) {
	p.PerPage = p.getPerPage(cast.ToInt(p.c.Query("per_page")))
	p.Order = p.getOrder(p.c.Query("order"))
	p.Sort = p.getSort(p.c.Query("sort"))
	p.TotalCount = p.getTotalCount(dataLen)
	p.TotalPage = p.getTotalPage()
	p.CurrentPage = p.getCurrentPage(cast.ToInt(p.c.Query("current_page")))
	p.Offset = (p.CurrentPage - 1) * p.PerPage
}

func (p *Paginator) getPerPage(perPage int) int {
	if perPage <= 0 {
		return 1
	}
	return perPage
}

func (p *Paginator) getCurrentPage(currentPage int) int {
	if currentPage <= 0 {
		return 1
	}
	if p.TotalPage == 0 {
		return 0
	}
	if currentPage > p.TotalPage {
		return p.TotalPage
	}
	return currentPage
}

func (p *Paginator) getTotalCount(dataLen int64) (count int64) {
	if dataLen > 0 {
		return dataLen
	}
	if err := p.query.Count(&count).Error; err != nil {
		return 0
	}
	return count
}

func (p *Paginator) getTotalPage() int {
	if p.TotalCount == 0 {
		return 0
	}
	nums := int64(math.Ceil(float64(p.TotalCount) / float64(p.PerPage)))
	if nums == 0 {
		nums = 1
	}
	return int(nums)
}

func (p *Paginator) getOrder(order string) string {
	if order == "asc" {
		return order
	}

	return "desc"
}

func (p *Paginator) getSort(sort string) string {
	switch sort {
	case "id", "created_at", "updated_at":
		return sort
	default:
		return "id"
	}
}
