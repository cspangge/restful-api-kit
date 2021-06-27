package apiHelpers

const (
	//Success as constant
	Success = "Success."
	//Error as constant
	Error = "Error."

	SUCCESS                  = 200
	NOT_ENOUGH_PARAMTERS     = 1000
	FAIL_TO_PARSE_PARAMETERS = 1002
	FAIL_TO_LOGIN            = 1003
	FAIL_TO_GET_TOKEN        = 1004
	FAIl_TO_GET_CACHE_RES    = 1005
)

var (
	CODE2MSG = map[int]string{
		SUCCESS:                  "Success",
		NOT_ENOUGH_PARAMTERS:     "Not enough parameters",
		FAIL_TO_PARSE_PARAMETERS: "Fail to parse parameters",
		FAIL_TO_LOGIN:            "Check your email and password, please",
		FAIL_TO_GET_TOKEN:        "Fail to get token",
		FAIl_TO_GET_CACHE_RES:    "Fail to get cache result",
	}
)
