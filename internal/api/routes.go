package api

type Endpoint struct {
	Method string
	Path   string
}

var Secrets = struct {
	Create Endpoint
	Get    Endpoint
	Update Endpoint
	Delete Endpoint
	List   Endpoint

	Versions SecretVersions
}{
	Create: Endpoint{Method: "POST", Path: "/secrets"},
	Get:    Endpoint{Method: "GET", Path: "/secrets/{id}"},
	Update: Endpoint{Method: "PATCH", Path: "/secrets/{id}"},
	Delete: Endpoint{Method: "DELETE", Path: "/secrets/{id}"},
	List:   Endpoint{Method: "GET", Path: "/secrets"},

	Versions: SecretVersions{
		Get:  Endpoint{Method: "GET", Path: "/secrets/{id}/versions/{version}"},
		List: Endpoint{Method: "GET", Path: "/secrets/{id}/versions"},
	},
}

type SecretVersions struct {
	Get  Endpoint
	List Endpoint
}
