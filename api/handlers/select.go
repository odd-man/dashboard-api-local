/**
*
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
	"github.com/seeleteam/dashboard-api/db"
	"github.com/seeleteam/dashboard-api/db/query/origin"
	"github.com/seeleteam/dashboard-api/db/query/param"
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
		fields := c.QueryArray(common.RequestFields)
		tableName := c.Query(common.RequestMeasurement)
		if tableName == "" {
			errInfo := fmt.Sprintf("param %s error", common.RequestMeasurement)
			log.Errorln(errInfo)
			responseData := common.NewResponseData(404, errInfo, nil, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}

		whereExpressions := c.QueryArray(common.RequestWhereExpressions)
		startTime := c.Query(common.RequestStartTime)
		endTime := c.Query(common.RequestEndTime)
		timeSince := c.Query(common.RequestTimeSince)

		intervals := c.Query(common.RequestIntervals)
		intervalsOffset := c.Query(common.RequestIntervalsOffset)
		tags := c.QueryArray(common.RequestTags)

		fillOption := c.Query(common.RequestFillOption)
		orderBy := c.Query(common.RequestOrderBy)

		limit := 0
		limitVal := c.Query(common.RequestLimit)
		if limitVal != "" {
			limit1, err := strconv.ParseInt(limitVal, 10, 10)
			if err != nil {
				log.Error(err)
				responseData := common.NewResponseData(404, err, nil, c.Request.RequestURI)
				ResponseJSON(c, responseData)
				return
			}
			limit = int(limit1)
		}

		offset := 0
		offsetVal := c.Query(common.RequestOffset)
		if offsetVal != "" {
			offset1, err := strconv.ParseInt(limitVal, 10, 10)
			if err != nil {
				log.Error(err)
				responseData := common.NewResponseData(404, err, nil, c.Request.RequestURI)
				ResponseJSON(c, responseData)
				return
			}
			offset = int(offset1)
		}

		slimit := 0
		slimitVal := c.Query(common.RequestSLimit)
		if slimitVal != "" {
			slimit1, err := strconv.ParseInt(slimitVal, 10, 10)
			if err != nil {
				log.Error(err)
				responseData := common.NewResponseData(404, err, nil, c.Request.RequestURI)
				ResponseJSON(c, responseData)
				return
			}
			slimit = int(slimit1)
		}

		soffset := 0
		soffsetVal := c.Query(common.RequestSOffset)
		if soffsetVal != "" {
			soffset1, err := strconv.ParseInt(slimitVal, 10, 10)
			if err != nil {
				log.Error(err)
				responseData := common.NewResponseData(404, err, nil, c.Request.RequestURI)
				ResponseJSON(c, responseData)
				return
			}
			soffset = int(soffset1)
		}

		timeZone := c.Query(common.RequestTimeZone)

		condition := &db.Condition{
			Fields:           fields,
			Measurement:      tableName,
			WhereExpressions: whereExpressions,
			StartTime:        startTime,
			EndTime:          endTime,
			TimeSince:        timeSince,
			Intervals:        intervals,
			IntervalsOffset:  intervalsOffset,
			Tags:             tags,
			FillOption:       fillOption,
			OrderBy:          orderBy,
			Limit:            limit,
			Offset:           offset,
			SLimit:           slimit,
			SOffset:          soffset,
			TimeZone:         timeZone,
		}

		paramQuery, err := param.New(condition)
		if err != nil {
			log.Error("%v", err)
			responseData := common.NewResponseData(500, err, nil, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}

		log.Debug("stmt: %v\n", paramQuery.Stmt)

		res, err := paramQuery.Query()
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
