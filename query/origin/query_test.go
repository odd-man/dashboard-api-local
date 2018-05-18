package origin

import (
	"fmt"
	"testing"
)

func Test_New(t *testing.T) {
	stmt := "select stddev(count) as count  from \"chain.block.insert.meter\" " +
		" where time >= now() - 1d group by time(1m,-1m) ,addr fill(none) limit 20 tz('Asia/Shanghai')"

	query := New(stmt)
	fmt.Printf("Query is\n%#v\n", query)
}

func Test_GetChartData(t *testing.T) {
	stmt := "select stddev(count) as count  from \"chain.block.insert.meter\" " +
		" where time >= now() - 1d group by time(1m,-1m) ,addr fill(none) limit 20 tz('Asia/Shanghai')"

	query := New(stmt)
	fmt.Printf("Query is\n%#v\n", query)

	res, err := query.Query()
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("res is:\n%v\n", res)

}
