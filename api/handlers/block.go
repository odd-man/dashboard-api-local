/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/seeleteam/dashboard-api/common"
	"github.com/seeleteam/dashboard-api/query"
	"github.com/seeleteam/dashboard-api/query/meter"
)

// GetMeterLineData get meter metrics data for line chart
func GetMeterLineData() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Request.
		tableName := c.Query(common.RequestMeasurement)
		if tableName == "" {
			c.JSON(404, gin.H{
				"message": "tableName error for" + c.Request.URL.Path,
			})
		}

		limit := c.GetInt(common.RequestLimit)
		if limit <= 0 {
			limit = 20

		}

		timeSince := c.GetString(common.RequestTimeSince)
		startTime := c.GetString(common.RequestStartTime)
		endTime := c.GetString(common.RequestEndTime)
		fillOption := c.GetString(common.RequestFillOption)
		if fillOption == "" {
			fillOption = "0"
		}

		tag := c.GetString(common.RequestTag)
		timeZone := c.GetString(common.RequestTimeZone)

		condition := &query.Condition{
			// Fields:      "stddev(count) as cc",
			Measurement: tableName,
			Limit:       limit,
			TimeSince:   timeSince,
			StartTime:   startTime,
			EndTime:     endTime,
			FillOption:  fillOption,
			Tag:         tag,
			TimeZone:    timeZone,
		}

		meterQuery := meter.New(condition)

		res, err := meterQuery.Query()
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error() + c.Request.URL.Path,
			})
		}

		chartData, err := meterQuery.GetChartData(res)

		responseData := &common.ResponseData{
			Code: 200,
			Msg:  "",
			Data: chartData,
			URI:  c.Request.RequestURI,
		}

		log.Info("GetMeterLineData() response:\n%v", responseData)
		c.JSON(200, responseData)
	}
}
