package models

import "time"

// User is the users model
type User struct {
	ID 			int
	FirstName 	string
	LastName 	string
	Email 		string
	Password 	string
	AccessLevel int
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

// Room is the room model
type Room struct {
	ID 			int
	RoomName 	string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

// Restriction is the restriction model
type Restriction struct {
	ID 				int
	RestrictionName string
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}

// Reservation is the reservation model
type Reservation struct {
	ID 			int
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	FirstName 	string
	LastName 	string
	Email 		string
	Phone 		string
	StartDate 	time.Time
	EndDate 	time.Time
	RoomID 		int
	Room		Room
	Processed 	int
}

// RoomRestrictions is the room restriction model
type RoomRestriction struct {
	ID 				int
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	StartDate 		time.Time
	EndDate 		time.Time
	RoomID			int
	Room			Room
	ReservationID 	int
	Reservation		Reservation
	RestrictionID 	int
	Restriction		Restriction
}