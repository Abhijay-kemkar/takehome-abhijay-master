package common

type RequestType string

const (
	Echo RequestType = "ECHO"
	Sync RequestType = "Sync"
)

type BaseRequest struct {
	RequestId   string `json:"request_id"`
	RequestType string `json:"request_type"`
}

type BaseResponse struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
    RequestId   string `json:"requestId"`
    RequestType string `json:"requestType"`
}

type EchoRequest struct {
	BaseRequest
	Value string
}

type EchoResponse struct {
	BaseResponse
	Value string
}

type SyncRequest struct {
	BaseResponse
    Directory string `json:"directory"`
    Filename  string `json:"filename"`
    Contents  string `json:"contents"`
}



