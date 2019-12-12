package entity

//
type DBEntity struct {
	UserName  string `json:"userName"`
	PassWord  string `json:"passWord"`
	IpAddr    string `json:"ipAddr"`
	Port      string `json:"port"`
	DBName    string `json:"DBName"`
	TableName string `json:"tableName"`
}
