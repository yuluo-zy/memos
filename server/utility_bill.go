package server

import (
	"github.com/labstack/echo/v4"
	"github.com/usememos/memos/store"
	"net/http"
)

func (s *Server) registerUtilityBillRoutes(g *echo.Group, mysqlUrl string) {

	g.GET("/utility_bill", func(c echo.Context) error {
		res, err := store.FindUtilityBill(mysqlUrl)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to list memo relations").SetInternal(err)
		}
		return c.JSON(http.StatusOK, composeResponse(res))
	})
}
