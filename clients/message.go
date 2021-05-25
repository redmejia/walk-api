package clients

type Message struct {
	Signin bool `json:"signin"`
	UserId int  `json:"user_id"`
}
