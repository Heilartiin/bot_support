package vdsin


type Balance struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      struct {
		Real    string `json:"real"`
		Bonus   string `json:"bonus"`
		Partner string `json:"partner"`
	} `json:"data"`
}

type ServerResponse struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      []*Server `json:"data"`
}

type Server struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	FullName   string      `json:"full_name"`
	Created    string      `json:"created"`
	Updated    string      `json:"updated"`
	End        string      `json:"end"`
	Status     string      `json:"status"`
	StatusText interface{} `json:"status_text"`
	IP         struct {
		ID      int    `json:"id"`
		IP      string `json:"ip"`
		Type    string `json:"type"`
		Host    string `json:"host"`
		Gateway string `json:"gateway"`
		Netmask string `json:"netmask"`
		Mac     string `json:"mac"`
		System  bool   `json:"system"`
	} `json:"ip"`
	ServerPlan struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"server-plan"`
	Template struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"template"`
	Datacenter struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"datacenter"`
	Can struct {
		Reboot  bool `json:"reboot"`
		Update  bool `json:"update"`
		Delete  bool `json:"delete"`
		Prolong bool `json:"prolong"`
	} `json:"can"`
}

type ServerCredential struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Password string `json:"password"`
	} `json:"data"`
}