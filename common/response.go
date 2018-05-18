/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package common

// ResponseData response data
type ResponseData struct {
	Code int         `json:"code"` // http status code
	Msg  string      `json:"msg"`  // error msg
	Data interface{} `json:"data"` // data, can include data code
	URI  string      `json:"uri"`  // request uri
}

// ChartLineData line chart data
type ChartLineData struct {
	Legend []string    `json:"legend"` // chart legend
	Data   interface{} `json:"data"`   // chart line data, inner struct maybe changed with different multi value
	Multi  bool        `json:"multi"`  // if have multi lines, this should be true, default is single line
}

// NewResponseData get the ResponseData according the fields
func NewResponseData(code int, msg string, data interface{}, uri string) *ResponseData {
	return &ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
		URI:  uri,
	}
}
