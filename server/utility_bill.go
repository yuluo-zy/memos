package server

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/usememos/memos/store"
	"net/http"
	"strconv"
)

func (s *Server) registerUtilityBillRoutes(g *echo.Group) {

	g.GET("/utility_bill", func(c echo.Context) error {
		ctx := c.Request().Context()
		db, err := sql.Open("mysql", "用户名:密码@tcp(localhost:3306)/数据库名")

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("db cannot open!")).SetInternal(err)
		}

		defer db.Close()
		memoID, err := strconv.Atoi(c.Param("memoId"))
		if err != nil {

		}

		memoRelationList, err := s.Store.ListMemoRelations(ctx, &store.FindMemoRelationMessage{
			MemoID: &memoID,
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to list memo relations").SetInternal(err)
		}
		return c.JSON(http.StatusOK, composeResponse(memoRelationList))
	})
}
