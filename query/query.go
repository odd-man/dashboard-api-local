/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package query

// Query service for query
type Query interface {
	// query data from db
	Query() (res interface{}, err error)

	// convert db data to chart data
	GetChartData(interface{}) (data interface{}, err error)
}
