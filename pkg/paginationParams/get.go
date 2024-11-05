package paginationParams

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type pagination struct {
	Page   int
	Limit  int
	Offset int
}

func Get(ctx *gin.Context) *pagination {
	page, errPage := strconv.Atoi(ctx.Query("page"))
	if errPage != nil {
		page = 0
	}

	limit, errLimit := strconv.Atoi(ctx.Query("limit"))
	if errLimit != nil {
		limit = 0
	}

	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}

	return &pagination{
		Page:   page,
		Limit:  limit,
		Offset: offset,
	}
}
