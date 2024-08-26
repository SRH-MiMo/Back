package entities

type AuthDAO struct {
	UUID     string `gorm:"primary_key"`
	Nickname string
	Age      string
	Job      string
}

type AuthDTO struct {
	Nickname string `json:"nickname"`
	Age      string `json:"age"`
	Job      string `json:"job"`
}
