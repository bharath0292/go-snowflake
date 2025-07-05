package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	service "github.com/go-snowflake/internal/services"
)

var (
	ErrTableNotFound  = errors.New("table not found")
	ErrInvalidTable   = errors.New("invalid table name")
	ErrColumnNotFound = errors.New("column not found")
	ErrDatabaseError  = errors.New("database error")
)

type ColumnListHandler struct {
	columnListService service.IColumnListService
}

func NewColumnListHandler(columnListService service.IColumnListService) *ColumnListHandler {
	return &ColumnListHandler{
		columnListService: columnListService,
	}
}

func (h *ColumnListHandler) GetAllColumnsInfo(c *gin.Context) {
	tableName := strings.TrimSpace(c.Query("tableName"))

	columns, err := h.columnListService.GetAllColumnsInfo(c.Request.Context(), tableName)
	if err != nil {
		if errors.Is(err, ErrTableNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("Table '%s' not found", tableName),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  columns,
		"count": len(columns),
	})
}
