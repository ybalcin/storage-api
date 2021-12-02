package port

type RecordHttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Records interface{} `json:"records"`
}
