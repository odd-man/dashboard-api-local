/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package common

// BaseSelectSection
// SELECT <field_key>[,<field_key>,<tag_key>] FROM <measurement_name>[,<measurement_name>]
const (
	// RequestSelectFields columns without index in influxdb, if multi(array) separated by comma
	RequestFieldKeys = "fields"
	// RequestSelectTags columns with index in influxdb, also as columns, if multi(array) separated by comma
	RequestTagKeys = "tags"
	// RequestMeasurementNames table names in influxdb, if multi(array) separated by comma
	RequestMeasurementNames = "measurements"
)

// WhereSection
// SELECT_clause FROM_clause WHERE <conditional_expression> [(AND|OR) <conditional_expression> [...]]
// The WHERE clause supports conditional_expressions on fields, tags, and timestamps.
// field: field_key <operator> ['string' | boolean | float | integer]
// tags: tag_key <operator> ['tag_value']
// timestamps:
// 	For most SELECT statements, the default time range is between 1677-09-21 00:12:43.145224194 and 2262-04-11T23:47:16.854775806Z UTC.
// 	For SELECT statements with a GROUP BY time() clause, the default time range is between 1677-09-21 00:12:43.145224194 UTC and now()
const (
	// RequestWhereExpressions where expression content in request
	// WHERE clause: =   equal to <> not equal to != not equal to =~ matches against !~ doesn’t match against
	// include all where conditional_expression, if multi(array) separated by comma
	RequestWhereExpressions = "whereExps"
)

/*
* GroupSection
* GROUP BY time intervals:	Basic Syntax	Advanced Syntax	GROUP BY time intervals and fill()
*
* 1. Syntax
* SELECT_clause FROM_clause [WHERE_clause] GROUP BY [* | <tag_key>[,<tag_key]]
*
* 2. GROUP BY time intervals
* SELECT <function>(<field_key>) FROM_clause WHERE <time_range> GROUP BY time(<time_interval>),[tag_key] [fill(<fill_option>)]
*
* 3. Advanced GROUP BY time() Syntax
* SELECT <function>(<field_key>) FROM_clause WHERE <time_range> GROUP BY time(<time_interval>,<offset_interval>),[tag_key] [fill(<fill_option>)]
*
* 4. GROUP BY time intervals and fill()
* SELECT <function>(<field_key>) FROM_clause WHERE <time_range> GROUP BY time(time_interval,[<offset_interval])[,tag_key] [fill(<fill_option>)]
 */
const (
	// RequestGroupBys group by, if multi(array) separated by comma
	RequestGroupBys = "groups"
)

const (
	// RequestFields columns for influxdb without index, separated by comma
	RequestFields = "fields"
	// RequestMeasurement table name
	RequestMeasurement = "table"

	// RequestDataBase database name
	RequestDataBase = "db"

	// RequestStartTime startTime field in request
	RequestStartTime = "startTime"

	// RequestEndTime endTime field in request
	RequestEndTime = "endTime"

	// RequestTimeSince timeSince field in request
	RequestTimeSince = "timeSince"

	// RequestWithExpression expression
	// WITH MEASUREMENT <regular_expression>
	// Regular expressions are surrounded by / characters
	RequestWithExpression = "withExp"

	// RequestOffset OFFSET <N> paginates N points in the query results
	RequestOffset = "offset"

	// RequestIntervals intervals field in request
	RequestIntervals = "intervals"

	// RequestIntervalsOffset intervalsOffset field in request
	RequestIntervalsOffset = "intervalsOffset"

	// RequestTag groupTag field in request
	// warn: only support single tag or not
	RequestTag = "tag"

	// RequestFillOption fillOption field in request
	RequestFillOption = "fillOption"

	// RequestOrder order field in request
	RequestOrder = "order"

	// RequestLimit limit field in request
	RequestLimit = "limit"

	// RequestTimeZone timeZone field in request
	RequestTimeZone = "timeZone"

	// RequestSQL the sql field in request
	RequestSQL = "sql"

	// RequestSQLs the sql array field in request
	RequestSQLs = "sqls"
)
