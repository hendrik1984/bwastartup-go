package user

// struct ini yang mewakili apa yang di input user di web
// yang mana akan prosesnya nanti akan di handle oleh "handler" untuk memappingkan
//register user menjadi struct ini
type RegisterUserInput struct {
	Name       string
	Occupation string
	Email      string
	Password   string
}
