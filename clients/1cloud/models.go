package _cloud



type Balance struct {
	ID                 int    `json:"ID"`
	FirstName          string `json:"FirstName"`
	LastName           string `json:"LastName"`
	Email              string `json:"Email"`
	Balance            float64    `json:"Balance"`
	BalanceTillDateUtc string `json:"BalanceTillDateUtc"`
	//ProjectsBalance    []struct {
	//	ProjectID int `json:"ProjectId"`
	//	Balance   struct {
	//		Real  float64 `json:"Real"`
	//		Bonus int `json:"Bonus"`
	//		Test  int `json:"Test"`
	//	} `json:"Balance"`
	//} `json:"ProjectsBalance"`
}

type Server struct {
	ID                int         `json:"ID"`
	Name              string      `json:"Name"`
	HostName          interface{} `json:"HostName"`
	State             string      `json:"State"`
	IsPowerOn         bool        `json:"IsPowerOn"`
	CPU               int         `json:"CPU"`
	RAM               int         `json:"RAM"`
	HDD               int         `json:"HDD"`
	IP                string      `json:"IP"`
	AdminUserName     string      `json:"AdminUserName"`
	AdminPassword     string      `json:"AdminPassword"`
	Image             string      `json:"Image"`
	IsHighPerformance bool        `json:"IsHighPerformance"`
	HDDType           string      `json:"HDDType"`
	PrimaryNetworkIP  string      `json:"PrimaryNetworkIp"`
	LinkedNetworks    []struct {
		LinkID      int    `json:"LinkID"`
		NetworkID   int    `json:"NetworkID"`
		LinkState   string `json:"LinkState"`
		NetworkType string `json:"NetworkType"`
		NetworkName string `json:"NetworkName"`
		IP          string `json:"IP"`
		MAC         string `json:"MAC"`
		Mask        string `json:"Mask"`
		Gateway     string `json:"Gateway"`
		Bandwidth   int    `json:"Bandwidth"`
	} `json:"LinkedNetworks"`
	DCLocation    string        `json:"DCLocation"`
	ImageFamily   string        `json:"ImageFamily"`
	LinkedSSHKeys []interface{} `json:"LinkedSshKeys"`
}

type OneCloudServer struct {
	ConnectionType     string `csv:"ConnectionType" json:"connection_type"`
	Name               string `csv:"Name" json:"name"`
	Host               string `csv:"Host" json:"host"`
	CredentialPassword string `csv:"CredentialPassword" json:"credential_password"`
	CredentialUserName string `csv:"CredentialUserName" json:"credential_user_name"`
	SubMode            string `csv:"SubMode" json:"sub_mode"`
}
