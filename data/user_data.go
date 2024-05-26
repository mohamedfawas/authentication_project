package data

type User struct {
	Name     string
	Email    string
	Password string
}

type SessionData struct {
	SessionId string
	Email     string
}

var UserData = make(map[string]User)
var Sessions = make(map[string]SessionData)

func init() {
	//User Datas are initialized
	UserData["vinijr@gmail.com"] = User{"Vinicius Junior", "vinijr@gmail.com", "777"}
	UserData["rogrygo@gmail.com"] = User{"Rodrygo Goes", "rogrygo@gmail.com", "1111"}
	UserData["jbellingham@gmail.com"] = User{"Jude Victor Bellingham", "jbellingham@gmail.com", "555"}
}
