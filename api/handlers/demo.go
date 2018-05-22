/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/seeleteam/dashboard-api/common"
	"github.com/seeleteam/dashboard-api/query"
	"github.com/seeleteam/dashboard-api/query/meter"
)

// GetMeterLineData get meter metrics data for line chart
func GetMeterLineData() gin.HandlerFunc {
	return func(c *gin.Context) {
		tableName := c.Query(common.RequestMeasurement)
		if tableName == "" {
			responseData := common.NewResponseData(404, errors.New("param table error"), nil, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}

		limit := 10
		limitVal := c.Query(common.RequestLimit)
		if limitVal != "" {
			limit1, _ := strconv.ParseInt(limitVal, 10, 10)
			limit = int(limit1)
		}
		if limit <= 0 {
			limit = 20
		}

		timeSince := c.Query(common.RequestTimeSince)
		startTime := c.Query(common.RequestStartTime)
		endTime := c.Query(common.RequestEndTime)

		fillOption := c.Query(common.RequestFillOption)
		if fillOption == "" {
			fillOption = "null"
		}
		intervals := c.Query(common.RequestIntervals)
		intervalsOffset := c.Query(common.RequestIntervalsOffset)
		if intervalsOffset == "" {
			intervalsOffset = "30s"
		}
		order := c.Query(common.RequestOrder)

		tag := c.Query(common.RequestTag)
		timeZone := c.Query(common.RequestTimeZone)

		condition := &query.Condition{
			// Fields:      "stddev(count) as cc",
			Measurement:     tableName,
			Limit:           limit,
			TimeSince:       timeSince,
			StartTime:       startTime,
			EndTime:         endTime,
			FillOption:      fillOption,
			Tag:             tag,
			TimeZone:        timeZone,
			Intervals:       intervals,
			IntervalsOffset: intervalsOffset,
			Order:           order,
		}

		meterQuery := meter.New(condition)
		log.Debug("stmt: %v\n", meterQuery.Stmt)

		res, err := meterQuery.Query()
		if err != nil {
			log.Error("%v", err)
			responseData := common.NewResponseData(500, err, nil, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}

		chartData, err := meterQuery.GetChartData(res)
		if err != nil {
			responseData := common.NewResponseData(500, err, nil, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}
		responseData := common.NewResponseData(200, nil, chartData, c.Request.RequestURI)
		ResponseJSON(c, responseData)
	}
}
