/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/seeleteam/dashboard-api/api/handlers"
	"github.com/seeleteam/dashboard-api/common"
)

// InitRouters init routers
func InitRouters(e *gin.Engine) {
	// set api handlers logger
	handlers.SetAPIHandlerLog("api-handlers", common.PrintLog)

	// routerGroup API
	routerGroupAPI := e.Group("/api")
	// base
	routerGroupAPI.GET("/ping", handlers.Ping())

	apiShowGroup := routerGroupAPI.Group("/show")
	apiShowGroup.GET("/databases", handlers.ShowDatabases())
	apiShowGroup.GET("/retentionPolices", handlers.ShowRetentionPolices())
	apiShowGroup.GET("/series", handlers.ShowSeries())
	apiShowGroup.GET("/measurements", handlers.ShowMeasurements())
	apiShowGroup.GET("/tagKeys", handlers.ShowTagKeys())
	apiShowGroup.GET("/tagValues", handlers.ShowTagValues())
	apiShowGroup.GET("/fieldKeys", handlers.ShowFieldKeys())

	// base sql in group api
	apiSelectGroup := routerGroupAPI.Group("/query")
	apiSelectGroup.GET("/sql", handlers.SelectBySQL())
	apiSelectGroup.GET("/multiSql", handlers.SelectByMultiSQL())
	apiSelectGroup.GET("/params", handlers.SelectWithParams())

	// api demo
	apiDemoGroup := routerGroupAPI.Group("/demo")
	apiDemoGroup.GET("/meterData/line", handlers.GetMeterLineData())
}
