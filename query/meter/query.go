package meter

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/influxdata/influxdb/client/v2"

	"github.com/seeleteam/dashboard-api/common"
	"github.com/seeleteam/dashboard-api/db"
	"github.com/seeleteam/dashboard-api/query"
)

var (
	// ErrorQueryFields query fields error
	ErrorQueryFields = errors.New("query fields|columns(required) error")

	// ErrorQueryFieldsFormat query fields string format error
	ErrorQueryFieldsFormat = errors.New("fields|columns format error")

	// ErrorMeasurement measurement error
	ErrorMeasurement = errors.New("measurement|tablename(required) error")

	// ErrorWhereCondition where condition format error
	ErrorWhereCondition = errors.New("where condition error")
)

// Query meter query
type Query struct {
	stmt string // query string
	tag  string // flag can distinguish the data, if multi must be append with comma
}

// New get Query for Meter
func New(cond *query.Condition) *Query {
	stmt, err := generateQueryStmt(cond)
	if err != nil {
		return nil
	}
	tag := cond.Tag

	return &Query{
		stmt: stmt,
		tag:  tag,
	}
}

// Query query data from db for meter
func (m *Query) Query() (res []client.Result, err error) {
	if m.stmt == "" {
		return nil, errors.New("error query stmt")
	}
	return db.Query(m.stmt)
}

// GetChartData get chart data
func (m *Query) GetChartData(res interface{}) (chartData *common.ChartLineData, err error) {
	//[]map[string]interface{}
	result, ok := res.([]client.Result)
	if !ok {
		return nil, errors.New("input data type error")
	}
	return m.generateChartLineData(result)
}

func (m *Query) generateChartLineData(inputRes []client.Result) (chartData *common.ChartLineData, err error) {
	// one stmt one result, here index should 0
	if len(inputRes) == 0 {
		return nil, errors.New("error result, empty data")
	}
	if len(inputRes) != 1 {
		return nil, errors.New("error result, not support multi stmt query")
	}
	res := inputRes[0]
	lineCount := len(res.Series)
	lineDatas := make([][]map[string]interface{}, lineCount)
	// legend for multi lines, maybe nil when tag is blank
	legend := make([]string, lineCount)

	tagName := m.tag
	// only one line
	if tagName == "" {
		for i, seriesItem := range res.Series {
			// one series one chart line or bar
			name := seriesItem.Name
			columns := seriesItem.Columns
			values := seriesItem.Values
			fmt.Printf("index:%v, name:%v, columns:%v\n", i, name, columns)

			lineData := make([]map[string]interface{}, 1)
			for _, val := range values {
				lineDataItem := map[string]interface{}{
					"time": val[0],
					"val":  val[1],
				}

				lineData = append(lineData, lineDataItem)
			}
			lineDatas[i] = lineData
		}
		chartData = &common.ChartLineData{
			Legend: nil,
			Multi:  false,
			Data:   lineDatas,
		}
	} else {
		// may multi lines
		for i, seriesItem := range res.Series {
			// one series one chart line or bar
			name := seriesItem.Name
			columns := seriesItem.Columns
			tags := seriesItem.Tags
			values := seriesItem.Values

			fmt.Printf("index:%v, name:%v, columns:%v,tags:%v\n", i, name, columns, tags)
			legend[i] = tags[m.tag]

			lineData := make([]map[string]interface{}, 1)
			for _, val := range values {
				lineDataItem := map[string]interface{}{
					"time": val[0],
					"tag":  tags[tagName],
					"val":  val[1],
				}

				lineData = append(lineData, lineDataItem)
			}
			lineDatas[i] = lineData
		}
		// TODO mergeMultiLineData
		multiLineDatas := mergeMultiLineData(lineDatas)
		chartData = &common.ChartLineData{
			Legend: legend,
			Multi:  true,
			Data:   multiLineDatas,
		}
	}
	return chartData, nil
}

func mergeMultiLineData(dataSet [][]map[string]interface{}) []map[string]interface{} {
	lineCount := len(dataSet)
	if lineCount == 0 {
		return nil
	}
	fmt.Printf("dataSet is %v\n", dataSet)

	dataSize := len(dataSet[0])
	dataSetsNew := make([]map[string]interface{}, 0)
continueL:
	for i := 0; i < dataSize; i++ {
		innerMap := make(map[string]interface{})
		for j := 0; j < lineCount; j++ {
			data := dataSet[j][i]
			if data == nil {
				continue continueL
			}
			// fmt.Printf("data is %#v\n", data)
			tag := data["tag"].(string)
			innerMap[tag] = data["val"]
		}
		innerMap["time"] = dataSet[0][i]["time"].(string)
		dataSetsNew = append(dataSetsNew, innerMap)
	}
	// fmt.Printf("valid dataSetsNew:\n%v\n", dataSetsNew)
	return dataSetsNew
}

// generateQueryStmt generate meter query string
func generateQueryStmt(condition *query.Condition) (stmt string, err error) {
	var buffer bytes.Buffer
	buffer.WriteString("select ")
	if condition == nil {
		return "", errors.New("error query condition")
	}

	fields := condition.Fields
	if fields == "" {
		fields = "stddev(count) as count "
		// return "", ErrorQueryFields
	}
	if strings.HasSuffix(fields, ",") {
		return "", ErrorQueryFieldsFormat
	}

	// fields or columns
	buffer.WriteString(fields + " ")
	buffer.WriteString("from ")

	measurement := condition.Measurement
	if measurement == "" {
		return "", ErrorMeasurement
	}

	measurement = strings.Trim(measurement, "\"")
	if measurement == "" {
		return "", ErrorMeasurement
	}
	// measurement or tablename
	buffer.WriteString("\"" + measurement + "\" ")

	//where condition related time
	buffer.WriteString("where ")
	// must exist time condition append the where
	var timeCondition bytes.Buffer
	timeStart := condition.StartTime
	timeEnd := condition.EndTime
	timeSince := condition.TimeSince
	if timeStart != "" && timeEnd != "" {
		timeCondition.WriteString(fmt.Sprintf("time >= %s ", timeStart))
		timeCondition.WriteString(fmt.Sprintf("and time <= %s ", timeEnd))
	} else if timeStart != "" {
		timeCondition.WriteString(fmt.Sprintf("time >= %s ", timeStart))
	} else if timeEnd != "" {
		timeCondition.WriteString(fmt.Sprintf("time <= %s ", timeEnd))
	} else {
		if timeSince == "" {
			timeSince = "30s"
		}
		timeCondition.WriteString(fmt.Sprintf("time >= now() - %s ", timeSince))
	}
	buffer.WriteString(timeCondition.String())

	// where condition append
	var whereConditionBuffer bytes.Buffer
	whereCondition := condition.WhereCondition
	if whereCondition != nil && len(whereCondition) != 0 {
		// each must one condition, not support multi nest
		for k, v := range whereCondition {
			if v != "" {
				if strings.Contains(v, "and") {
					return "", ErrorWhereCondition
				}
				whereConditionBuffer.WriteString("and " + k + v)
			}
		}
		// append where condition
		buffer.WriteString(whereConditionBuffer.String() + " ")
	}

	// must required, group by time(30s,offsetInterval)
	var groupTimeCondition string
	intervals := condition.Intervals
	if intervals == "" {
		intervals = "30s"
		groupTimeCondition = intervals
	}

	intervalsOffset := condition.IntervalsOffset
	if intervalsOffset != "" {
		groupTimeCondition = intervals + "," + intervalsOffset
	}
	buffer.WriteString(fmt.Sprintf("group by time(%s) ", groupTimeCondition))

	groupTag := condition.Tag
	if groupTag != "" {
		buffer.WriteString("," + groupTag + " ")
	}

	if condition.FillOption != "" {
		buffer.WriteString(fmt.Sprintf("fill(%s) ", condition.FillOption))
	}

	// order by default asc
	if strings.ToLower(condition.Order) == "desc" {
		buffer.WriteString(fmt.Sprintf("fill(%s) ", "desc"))
	}

	// limit default ?
	limit := condition.Limit
	if limit <= 0 {
		limit = 100
	}

	buffer.WriteString(fmt.Sprintf("limit %d ", limit))

	// format like Asia/SHanghai
	zoneVal := condition.TimeZone
	if zoneVal != "" {
		buffer.WriteString(fmt.Sprintf("tz('%s')", zoneVal))
	}
	return buffer.String(), nil
}
