package place

type Place struct {
	Address    string `json:"address"`
	Country    string `json:"country"`
	Prefecture string `json:"prefecture"`
	City       string `json:"city"`
	Postal     string `json:"postal"`
}
