package api

type AppSettings struct {
	Host      string `json:"host"`
	Port      string `json:"port"`
	Endpoints string `json:"endpoints"`
}

type Status struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
