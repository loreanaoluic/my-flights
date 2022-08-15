package model

type FlightStatus string

const (
	ACTIVE   FlightStatus = "ACTIVE"
	FULL     FlightStatus = "FULL"
	CANCELED FlightStatus = "CANCELED"
)

func (os FlightStatus) String() string {
	switch os {
	case ACTIVE:
		return "ACTIVE"
	case FULL:
		return "FULL"
	case CANCELED:
		return "CANCELED"
	}
	return "unknown"
}
