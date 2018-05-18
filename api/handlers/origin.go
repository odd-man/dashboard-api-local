/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/seeleteam/dashboard-api/common"
	"github.com/seeleteam/dashboard-api/query/origin"
)

// GetBySQL get data from influxdb by influxdb sql
func GetBySQL() gin.HandlerFunc {
	return func(c *gin.Context) {
		sql := c.Query(common.RequestSQL)
		if sql == "" {
			errInfo := fmt.Sprintf("param field sql error")
			log.Errorln(errInfo)
			responseData := common.NewResponseData(500, errInfo, nil, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}
		query := origin.New(sql)
		log.Debug("stmt: %v\n", query.Stmt)

		res, err := query.Query()
		if err != nil {
			log.Errorln(err)
			responseData := common.NewResponseData(500, err.Error(), res, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}
		responseData := common.NewResponseData(200, err.Error(), res, c.Request.RequestURI)
		ResponseJSON(c, responseData)
	}
}
