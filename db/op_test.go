/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package db

import (
	"fmt"
	"testing"
)

func Test_Query(t *testing.T) {
	showDB := "show databases"
	res, err := Query(showDB)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v:\n%#v\n", showDB, res)
}

// all data is exist in the time point, if not exist use 0 or other data fill
func mergeMap(dataSet [][]map[string]interface{}) []map[string]interface{} {
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
