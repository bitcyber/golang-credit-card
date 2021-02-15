package banksdb

var banksDB = &banksDBImpl{
	prefixToBank: make(map[int]*Bank),
}

// Bank represent bank info.
type Bank struct {
	Name     string `json:"name"`
	Country  string `json:"country"`
	EngTitle string `json:"engTitle"`
	URL      string `json:"url"`
	Prefixes []int  `json:"prefixes"`
}

// FindBank search bank info in all countries.
func FindBank(creditCard string) *Bank {
	return banksDB.FindBank(creditCard)
}

func init() {
	for _, banks := range banksByCountry {
		banksDB.addBanksToDB(banks)
	}
}
