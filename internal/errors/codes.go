package errors

var (
	NetworkError = &Err{code: 100001, name: "network error"}
)

type Err struct {
	code int32
	name string
}
