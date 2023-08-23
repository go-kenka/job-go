package risk

type Request struct {
	AppID  string      `json:"app_id"`
	Time   int64       `json:"time"`
	Data   interface{} `json:"data"`
	Random string      `json:"random"`
	Sign   string      `json:"sign"`
}

type LogData struct {
	Organize  string     `json:"organize"`
	User      string     `json:"user"`
	Ip        string     `json:"ip"`
	UserAgent string     `json:"user_agent"`
	PageType  string     `json:"page_type"`
	Value     int64      `json:"value"`
	Request   string     `json:"request"`
	DbCount   []*DBCount `json:"db_count"`
}
type DBCount struct {
	DB        string `json:"db"`
	UniqueKey string `json:"unique_key"`
}

type SyncData struct {
	DbKey   string `json:"db_key"`
	DbName  string `json:"db_name"`
	DbCount int64  `json:"db_count"`
}
