package model

type TravelClass string

const (
	ECONOMY  TravelClass = "ECONOMY"
	BUSINESS TravelClass = "BUSINESS"
	FIRST    TravelClass = "FIRST"
)

func (os TravelClass) String() string {
	switch os {
	case ECONOMY:
		return "ECONOMY"
	case BUSINESS:
		return "BUSINESS"
	case FIRST:
		return "FIRST"
	}
	return "unknown"
}
