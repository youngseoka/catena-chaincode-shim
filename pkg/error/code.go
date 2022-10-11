package error

const (
	// Common code allocated to 0 ~ 999
	Ok                 uint32 = 0
	BadRequest         uint32 = 10
	UnknownMethod      uint32 = 11
	Unauthorized       uint32 = 12
	NotFound           uint32 = 13
	JsonMarshalError   uint32 = 20
	JsonUnmarshalError uint32 = 21
	UnknownError       uint32 = 999
)
