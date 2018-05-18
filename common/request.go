/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package common

const (
	// RequestFields columns for influxdb without index, separated by comma
	RequestFields = "fields"
	// RequestMeasurement table name
	RequestMeasurement = "table"

	// RequestStartTime startTime field in request
	RequestStartTime = "startTime"

	// RequestEndTime endTime field in request
	RequestEndTime = "endTime"

	// RequestTimeSince timeSince field in request
	RequestTimeSince = "timeSince"

	// RequestWhereCondition where condition in request maybe in body ??
	RequestWhereCondition = "whereCondition"

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
)
