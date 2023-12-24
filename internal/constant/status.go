package constant

const (
	// 100-global
	StatusSuccess = 100000
	StatusFailure = 100001

	// 20-authorization and security
	// 200-request
	RequestOriginInvalid    = 200000
	RequestTooFrequent      = 200001
	RequestMethodInvalid    = 200002
	RequestIpException      = 200003
	RequestParameterMissing = 200004
	RequestFormatError      = 200005

	// 201-request params
	TokenMissing     = 201001
	TokenExpired     = 201002
	TokenInvalid     = 201003
	SignatureMissing = 201004
	SignatureExpired = 201005
	SignatureError   = 201006

	// 300-account
	AccountFormatError         = 300000
	AccountAlreadyExist        = 300001
	AccountPasswordFormatError = 300002
	AccountNotExist            = 300003
	AccountPasswordError       = 300004
	AccountIsLocked            = 300005
	AccountIsInArrears         = 300006
)

func StatusText(code int) string {
	switch code {
	case StatusSuccess:
		return "success"
	case StatusFailure:
		return "failure"
	case RequestOriginInvalid:
		return "origin invalid"
	case RequestTooFrequent:
		return "request too frequent"
	case RequestMethodInvalid:
		return "request method invalid"
	case RequestIpException:
		return "request ip exception"
	case RequestParameterMissing:
		return "request parameter missing"
	case RequestFormatError:
		return "request format error"
	case TokenMissing:
		return "token missing"
	case TokenExpired:
		return "token expired"
	case TokenInvalid:
		return "token invalid"
	case SignatureMissing:
		return "signature missing"
	case SignatureExpired:
		return "signature expired"
	case SignatureError:
		return "signature error"
	case AccountFormatError:
		return "account format error"
	case AccountAlreadyExist:
		return "account already exist"
	case AccountPasswordFormatError:
		return "account password format error"
	case AccountNotExist:
		return "account not exist"
	case AccountPasswordError:
		return "account password error"
	case AccountIsLocked:
		return "account is locked"
	case AccountIsInArrears:
		return "account is arrears"
	default:
		return ""
	}
}
