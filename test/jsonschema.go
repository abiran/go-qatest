package test

type Header struct {
	Version  string   `json:"version"`
	Customer Customer `json:"customer"`
	Soda     Soda     `json:"soda"`
	MetalLb  string   `json:"metalLb"`
}

type Customer struct {
	Name       string `json:"name"`
	Location   string `json:"location"`
	Deployment string `json:"deployment"`
}

type Soda struct {
	App         string      `json:"app"`
	Version     string      `json:"version"`
	FQDN        string      `json:"fqdn"`
	NameSpace   string      `json:"namespace"`
	Bandwidth   string      `json:"bandwidth"`
	Persistence Persistence `json:"persistence"`
	Backup      Backup      `json:"backup"`
}

type Persistence struct {
	Enabled bool   `json:"enabled"`
	Volume  string `json:"volume"`
}

type Backup struct {
	Enabled         bool   `json:"enabled"`
	AwsAccessKey    string `json:"awsAccessKey"`
	AwsAccessSecret string `json:"awsAccessSecret"`
	AwsBucketName   string `json:"awsBucketName"`
}
