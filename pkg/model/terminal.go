package model

type Terminal struct {
	Name           string `json:"name"`
	Comment        string `json:"comment"`
	ServiceAccount struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		AccessKey struct {
			Id     string `json:"id"`
			Secret string `json:"secret"`
		}
	} `json:"service_account"`
}

type TerminalTask struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Args       string `json:"args"`
	IsFinished bool
}
