package errors

var values = map[string]ErrorDto{
	"MM0000000": ErrorDto{
		ID:     "MM0000000",
		Status: 500,
		Msg:    "Error Not found",
	},
	"MM0000001": ErrorDto{
		ID:     "MM0000001",
		Status: 400,
		Msg:    "Invalid inbound entity",
	},
	"MM0000002": ErrorDto{
		ID:     "MM0000002",
		Status: 404,
		Msg:    "Resource not found",
	},
	"MM0000003": ErrorDto{
		ID:     "MM0000003",
		Status: 500,
		Msg:    "Undefined Error",
	},
	"MM0000004": ErrorDto{
		ID:     "MM0000004",
		Status: 500,
		Msg:    "Connector is down",
	},
	"MM0000005": ErrorDto{
		ID:     "MM0000005",
		Status: 401,
		Msg:    "Authentication error",
	},
	"MM0000006": ErrorDto{
		ID:     "MM0000006",
		Status: 500,
		Msg:    "private keys errors",
	},
	"MM0000007": ErrorDto{
		ID:     "MM0000007",
		Status: 500,
		Msg:    "public keys errors",
	},
	"MM0000008": ErrorDto{
		ID:     "MM0000008",
		Status: 500,
		Msg:    "sign token error",
	},
	"MM0000009": ErrorDto{
		ID:     "MM0000009",
		Status: 415,
		Msg:    "content-type application/json it's required",
	},
	"MM0000010": ErrorDto{
		ID:     "MM0000010",
		Status: 500,
		Msg:    "mysql fatal error",
	},
	"MM0000011": ErrorDto{
		ID:     "MM0000011",
		Status: 400,
		Msg:    "unssuported auth algorithm",
	},
	"MM0000012": ErrorDto{
		ID:     "MM0000012",
		Status: 409,
		Msg:    "Resource conflict",
	},
	"MM0000013": ErrorDto{
		ID:     "MM0000013",
		Status: 500,
		Msg:    "Redis fatal error",
	},
	"MM0000014": ErrorDto{
		ID:     "MM0000014",
		Status: 500,
		Msg:    "Unmarshal redis error",
	},
	"MM0000015": ErrorDto{
		ID:     "MM0000015",
		Status: 404,
		Msg:    "Redis not found",
	},
	"MM0000016": ErrorDto{
		ID:     "MM0000016",
		Status: 401,
		Msg:    "authorization error",
	},
	"MM0000017": ErrorDto{
		ID:     "MM0000017",
		Status: 400,
		Msg:    "invalid authorization_token format",
	},
	"MM0000018": ErrorDto{
		ID:     "MM0000018",
		Status: 401,
		Msg:    "Token expires",
	},
	"MM0000019": ErrorDto{
		ID:     "MM0000019",
		Status: 409,
		Msg:    "Notification conflict",
	},
	"MM0000020": ErrorDto{
		ID:     "MM0000020",
		Status: 504,
		Msg:    "Gateway Timeout",
	},
	"MM0000021": ErrorDto{
		ID:     "MM0000021",
		Status: 504,
		Msg:    "Bad Gateway",
	},
}
