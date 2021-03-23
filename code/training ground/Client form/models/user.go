package models

type Client struct {
	ClientName     string
	ClientAddress  string
	ClientID       int
	ShortcodeType  int
	Requestor      string
	Position       string
	mobile         int
	email          string
	NameOfService  string
	GoLiveDate     string
	Description    string
	MO             string
	MT             string
	Pricepoint     float32
	VolumesPerWeek int
}

var (
	client []*Client
	nextID = 1
)

func AddClient(c Client) (Client, error) {
	c.ClientName = "SMS Global"
	client = append(client, &c)

	return c, nil
	/*ClientAddress: "128 Golding Street, Adelaide",
	ClientID:       1,
	ShortcodeType:  360,
	Requestor:      "John Morrison",
	Position:       "SDM",
	mobile:         0272761743
	email:          "testingRD@modica.co.nz",
	NameOfService:  "Alerts",
	GoLiveDate:   "12/02/2021" ,
	Description:  "This is a test",
	MO: "User to text keyword into 360"
	MT: "SMS Global: Thanks for your text"
	Pricepoint:     0.20
	VolumesPerWeek: 1000*/

}
