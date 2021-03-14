package defs

//Err Error Response Body
type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

//ErrorResponse Struct for sendingErrorResponse
type ErrorResponse struct {
	HTTPSc int
	Error  Err
}

var (
	//ErrorRequestBodyParseFailed Response for Request Body Parse Failed
	ErrorRequestBodyParseFailed = ErrorResponse{
		HTTPSc: 400,
		Error: Err{
			Error:     "Request body is not correct",
			ErrorCode: "001",
		},
	}
	//ErrorNotAuthUser Response for failed user authentication
	ErrorNotAuthUser = ErrorResponse{
		HTTPSc: 401,
		Error: Err{
			Error:     "User authentication failed",
			ErrorCode: "002",
		},
	}
	//ErrorDBError Response for DB op error
	ErrorDBError = ErrorResponse{
		HTTPSc: 500,
		Error: Err{
			Error:     "DB ops failed",
			ErrorCode: "003",
		},
	}
	//ErrorInternalFaults Response for InternalFaults
	ErrorInternalFaults = ErrorResponse{
		HTTPSc: 500,
		Error: Err{
			Error:     "Internal service error",
			ErrorCode: "004",
		},
	}
)
