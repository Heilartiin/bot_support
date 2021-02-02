package models

type Server struct {
	ConnectionType     string `csv:"ConnectionType" json:"connection_type"`
	Name               string `csv:"Name" json:"name"`
	Host               string `csv:"Host" json:"host"`
	CredentialPassword string `csv:"CredentialPassword" json:"credential_password"`
	CredentialUserName string `csv:"CredentialUserName" json:"credential_user_name"`
	SubMode            string `csv:"SubMode" json:"sub_mode"`
}
