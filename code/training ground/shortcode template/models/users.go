package models

type ShortCodeRequest struct {
	// Omni User variables
	FirstName string
	LastName  string
	Phone     int
	ID        int

	// Client User variables
	CompanyName       string
	CompanyAddress    string
	RequestorFullName string
	Position          string
	Email             string
	SCNumber          int
	SCPricePoint      float32
	SCType            string
}

var (
	scRequest []*ShortCodeRequest
	nextID    = 1
)

func NewSC(sc ShortCodeRequest) (ShortCodeRequest, error) {
	// This method is used when a new code needs to be setup
	sc.ID = nextID
	sc.CompanyAddress = "39 Omapere Street"
	nextID++
	scRequest = append(scRequest, &sc)
	return sc, nil

}

func ChangeServiceSC() []*ShortCodeRequest {
	// This method is used when the code is change from one client to another
}

func MigrateSC() {
	// This method is invoked when a code is migrating from one provider to another or
	// one connection to another connection
}

func DonationSC() {
	// This method is invoked when a request a donation service to be setup
}

func PricePointChangeSC() {
	// This method is invoked when a request for a price change to a shortcode

}*/
