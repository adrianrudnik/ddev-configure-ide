package ddev

type DescribeResult struct {
	Raw struct {
		AppRoot string `json:"approot"`
		Status  string `json:"status"`

		Database struct {
			Type         string `json:"database_type"`
			Version      string `json:"database_version"`
			InternalPort string `json:"dbPort"`
			MappedPort   int64  `json:"published_port"`
			Name         string `json:"dbname"`
			Hostname     string `json:"host"`
			Username     string `json:"username"`
			Password     string `json:"password"`
		} `json:"dbinfo,omitempty"`
	} `json:"raw"`
}
