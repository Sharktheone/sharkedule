package user

type User struct {
	UUID     string
	Username string
	Password string
	//OAuth    OAuth?
	//MFA	  	MFA?
	//Email        Email
	Boards       []string
	CustomColors []string
	//Settings     Settings

}
