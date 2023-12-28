package user_model

var (
	COST = "COST"
)

type userDomain struct {
	id       string
	email    string
	password string
	name     string
}

func (ud *userDomain) SetID(id string) {
	ud.id = id
}

func (ud *userDomain) GetId() string {
	return ud.id
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) SetName(name string) {
	ud.name = name
}
