/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package query

// Condition meter condition for query
type Condition struct {
	Fields string // columns for influxdb without index, separated by comma
	// tags        string // columns for influxdb with index, separated by comma

	Measurement string // table name for influxdb

	// where time >= {startTime} and time <= {endTime}, format like 2015-08-18T00:00:00Z,
	// 2015-08-18 00:12:00, 1439856000000000000, 1439856000s, 24043524m,
	// now() - 30s, now() - 1m, now() - 1d, ...
	// UTC time
	StartTime string
	// format like startTime
	EndTime string

	// if timeSince set startTime and endTime will be disabled
	TimeSince string

	// append time condition with and
	// format like: {"host": "='a'"},
	// {"flag": "=1"}
	WhereCondition map[string]string

	Intervals       string // required, default 30s ex: 5s, 5m, 5h..., group by time(5s)
	IntervalsOffset string // should be same unit for intervals, if exist will append intervals with comma

	// warn: only support single tag or not
	// format like "host"  "email"
	Tag string

	// fill options, in linear, none, null, previous
	// linear - Reports the results of linear interpolation for time intervals with no data.
	// none - Reports no timestamp and no value for time intervals with no data.
	// null -  Reports null for time intervals with no data but returns a timestamp. This is the same as the default behavior.
	// previous -   Reports the value from the previous time interval for time intervals with no data.
	// number
	FillOption string

	// default is asc, you can use desc replace it
	Order string

	Limit    int    // LIMIT <N> returns the first N points from the specified measurement.
	TimeZone string // default use UTC
}
