package user

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/labstack/echo"
	"github.com/nguyentrungtin/go-echo-boilerplate/db"
	"github.com/nguyentrungtin/go-echo-boilerplate/lib"
)

// GetAll is used to get all users
// using query params to offset/limit/sort (pagination)
// sortKeys := []string{"email", "name", "role", "created_at", "updated_at"}
// sortValues := []string{"asc", "desc"}
func GetAll(c echo.Context) error {
	users := []User{}
	db := db.Session()

	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit == 0 {
		limit = 10
	}
	offset := limit * (page - 1)

	sortKeys := []string{"email", "first_name", "last_name", "username", "role", "created_at", "updated_at"}
	sortValues := []string{"asc", "desc"}
	searchKeys := []string{"email", "first_name", "last_name", "username", "role", "created_at", "updated_at"}

	sort := c.QueryParam("sort")
	if _, found := lib.Find(sortKeys, sort); !found {
		sort = "updated_at"
	}

	sortBy := c.QueryParam("sortBy")
	if _, found := lib.Find(sortValues, sortBy); !found {
		sortBy = "desc"
	}

	sortStr := fmt.Sprintf("%s %s", sort, sortBy)

	searchKey := c.QueryParam("searchKey")
	if _, found := lib.Find(searchKeys, searchKey); !found {
		searchKey = "username"
	}

	search := c.QueryParam("search")
	search = fmt.Sprintf("%s%s%s", "%%", search, "%%")

	var count int
	if err := db.Model(&User{}).Where(gorm.ToColumnName(searchKey)+" LIKE ?", search).Order(sortStr).Offset(offset).Limit(limit).Find(&users).Offset(0).Limit(-1).Count(&count).Error; err != nil {
		return echo.ErrInternalServerError
	}

	pages := int(math.Ceil(float64(count) / float64(limit)))

	resp := struct {
		Data   []User `json:"data"`
		Count  int    `json:"count"`
		Pages  int    `json:"pages"`
		Page   int    `json:"page"`
		Limit  int    `json:"limit"`
		Sort   string `json:"sort"`
		SortBy string `json:"sortBy"`
	}{
		Data:   users,
		Count:  count,
		Pages:  pages,
		Page:   page,
		Limit:  limit,
		Sort:   sort,
		SortBy: sortBy,
	}

	return c.JSON(http.StatusOK, &resp)
}
