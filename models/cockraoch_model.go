package models

type CockroachModel struct {
	DriverID    string `db:"driver_id" json:"driver_id"`
	DriverName  string `db:"driver_name" json:"driver_name"`
	DriverCode  string `db:"driver_code" json:"driver_code"`
	NumberSim   string `db:"number_sim" json:"number_sim"`
	PhoneNumber string `db:"phone_number" json:"phone_number"`
	MemberID    string `db:"member_id" json:"member_id"`
	Imei        string `db:"imei" json:"imei"`
}
