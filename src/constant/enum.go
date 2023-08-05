package constant

const (
	Male Gender = iota + 1
	Female
	NonBinary
)

type Gender int64

func (g Gender) String() string {
	return [...]string{"Male", "Female", "NonBinary"}[g-1]
}

const (
	Hindu Religion = iota + 1
	Muslim
	Katolik
	Protestan
	Budha
	Khonghucu
)

type Religion int64

func (r Religion) String() string {
	return [...]string{"Hindu", "Muslim", "Katolik", "Protestan", "Budha", "Khonghucu"}[r-1]
}

const (
	UserTypeNormal UserType = iota + 1
	UserTypePremium
)

type UserType int64

func (u UserType) String() string {
	return [...]string{"Normal", "Premium"}[u-1]
}

func (u *UserType) IsEligibleUpgrade() bool {
	return map[UserType]bool{
		UserTypeNormal:  false,
		UserTypePremium: true,
	}[*u]
}
