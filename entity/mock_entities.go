package entity

// EventMock mocks Food Menu Event
var EventMock = Event{
	ID:          1,
	Name:        "Mock Event 01",
	Description: "Mock Event 01 Description",
	Location:    "Mock Event 01 Location",
	Time:        "Mock Event 01 Time",
	Image:       "mock_cat.png",
}

// RoleMock mocks user role entity
var RoleMock = Role{
	ID:    1,
	Name:  "Mock Role 01",
	Users: []User{},
}

// UserMock mocks application user
var UserMock = User{
	ID:       1,
	FullName: "Mock User 01",
	Email:    "mockuser@example.com",
	Phone:    "0900000000",
	Password: "P@$$w0rd",
	RoleID:   1,
	// Orders:   []Order{},
}

// SessionMock mocks sessions of loged in user
var SessionMock = Session{
	ID:         1,
	UUID:       "_session_one",
	SigningKey: []byte("RestaurantApp"),
	Expires:    0,
}
