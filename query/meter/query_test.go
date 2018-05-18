package meter

import (
	"fmt"
	"testing"

	"github.com/seeleteam/dashboard-api/query"
)

func Test_New(t *testing.T) {
	tagOrLegend := "addr"
	condition := &query.Condition{
		// Fields:      "stddev(count) as cc",
		Measurement: "chain.block.insert.meter",
		Limit:       10,
		TimeSince:   "30d",
		StartTime:   "",
		EndTime:     "",
		FillOption:  "0",
		WhereCondition: map[string]string{
			"addr": "='127.0.0.1:65027'",
		},
		Tag:      tagOrLegend,
		TimeZone: "Asia/Shanghai",
	}
	query := New(condition)
	fmt.Printf("Query is\n%#v\n", query)
}

func Test_GetChartData(t *testing.T) {
	tagOrLegend := "addr"
	condition := &query.Condition{
		// Fields:      "stddev(count) as cc",
		Measurement:    "chain.block.insert.meter",
		Limit:          20,
		TimeSince:      "30d",
		StartTime:      "",
		EndTime:        "",
		FillOption:     "0",
		WhereCondition: map[string]string{
			// "addr": "='127.0.0.1:65027'",
		},
		Tag:      tagOrLegend,
		TimeZone: "Asia/Shanghai",
	}
	query := New(condition)
	fmt.Printf("Query is\n%#v\n", query)

	res, err := query.Query()
	if err != nil {
		t.Error(err)
	}
	chartData, err := query.GetChartData(res)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("chartData is:\n%v\n", chartData)

}
