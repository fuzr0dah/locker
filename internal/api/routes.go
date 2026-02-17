package api

type Endpoint struct {
	Method string
	Path   string
}

var (
	SecretList   = Endpoint{Method: "GET", Path: "/secrets"}
	CreateSecret = Endpoint{Method: "POST", Path: "/secrets"}
	GetSecret    = Endpoint{Method: "GET", Path: "/secrets/{id}"}
	UpdateSecret = Endpoint{Method: "PUT", Path: "/secrets/{id}"}
	DeleteSecret = Endpoint{Method: "DELETE", Path: "/secrets/{id}"}
)
