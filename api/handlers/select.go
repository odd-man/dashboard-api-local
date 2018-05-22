/**
*  sql ref: https://docs.influxdata.com/influxdb/v1.5/query_language/data_exploration
*
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/seeleteam/dashboard-api/common"
	"github.com/seeleteam/dashboard-api/query"
	"github.com/seeleteam/dashboard-api/query/meter"
	"github.com/seeleteam/dashboard-api/query/origin"
)

// SelectBySQL get data from influxdb by influxdb sql
func SelectBySQL() gin.HandlerFunc {
	return func(c *gin.Context) {
		sql := c.Query(common.RequestSQL)
		if sql == "" {
			errInfo := fmt.Sprintf("param field sql required!")
			log.Errorln(errInfo)
			responseData := common.NewResponseData(500, errors.New(errInfo), nil, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}
		query := origin.New(sql)
		log.Debug("stmt: %v", query.Stmt)

		res, err := query.Query()
		if err != nil {
			log.Error("SelectBySQL, err:\n%v\n")
			responseData := common.NewResponseData(500, err, res, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}
		responseData := common.NewResponseData(200, nil, res, c.Request.RequestURI)
		ResponseJSON(c, responseData)
	}
}

// SelectByMultiSQL get data from influxdb by multi influxdb sql
func SelectByMultiSQL() gin.HandlerFunc {
	return func(c *gin.Context) {
		sqls := c.QueryArray(common.RequestSQLs)

		if sqls == nil || len(sqls) == 0 {
			errInfo := fmt.Sprintf("param field sqls required!")
			log.Errorln(errInfo)
			responseData := common.NewResponseData(400, errors.New(errInfo), nil, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}

		// be separated by semicolon
		var sqlStr bytes.Buffer
		for _, sql := range sqls {
			if sql != "" {
				if strings.HasSuffix(sql, ";") {
					sqlStr.WriteString(sql)
				} else {
					sqlStr.WriteString(sql + ";")
				}
			}
		}
		if sqlStr.String() == "" {
			errInfo := fmt.Sprintf("param field sqls content error")
			log.Errorln(errInfo)
			responseData := common.NewResponseData(500, errors.New(errInfo), nil, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}

		query := origin.New(sqlStr.String())
		log.Debug("stmt: %v", query.Stmt)

		res, err := query.Query()
		if err != nil {
			log.Error("SelectByMultiSQL, err:\n%v\n")
			responseData := common.NewResponseData(500, err, res, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}
		responseData := common.NewResponseData(200, nil, res, c.Request.RequestURI)
		ResponseJSON(c, responseData)
	}
}

// SelectWithParams select with params(generate sql)
func SelectWithParams() gin.HandlerFunc {
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
		responseData := common.NewResponseData(200, err, res, c.Request.RequestURI)
		ResponseJSON(c, responseData)
	}
}
