package restapi

// RestAPI is object for rest APi
type RestAPI struct {
	baseURL  string
	endPoint string
}

// New it Open connection to server
func New() (RepoRest *RestAPI, err error) {
	return &RestAPI{
		baseURL:  "172.31.4.92:8977/",
		endPoint: "getgeragestatus",
	}, nil
}
