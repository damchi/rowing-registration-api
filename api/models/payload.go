package models

type ClubRegistrationParam struct {
	RoleId         uint64 `json:"role"`
	FirstName      string `gorm:"size:255;not null;" json:"first_name"`
	LastName       string `gorm:"size:255;not null;" json:"last_name"`
	Email          string `gorm:"size:100;not null;unique" json:"email"`
	Password       string `gorm:"size:100;not null;" json:"password"`
	ClubName       string `gorm:"size:100;not null;" json:"club_name"`
	Address1       string `gorm:"size:100;not null;" json:"address1"`
	Address2       string `gorm:"size:100;" json:"address2,omitempty"`
	Address3       string `gorm:"size:100;" json:"address3,omitempty"`
	AddressCity    string `gorm:"size:100;not null;" json:"address_city"`
	AddressCountry string `gorm:"size:100;not null;" json:"address_country"`
	AddressState   string `gorm:"size:100;not null;" json:"address_state"`
	PostCode       string `gorm:"size:100;not null;" json:"post_code"`
	Phone          string `gorm:"size:100;not null;" json:"phone"`
}

type UserLoginParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
