package sources

var (
	walkscoreBaseUrl string = "https://www.walkscore.com"
)

type Score struct {
	Walk    int8
	Transit int8
	Bike    int8
}

type Neighborhood struct {
	Name       string
	Population int32
	Score      Score
}

type Walkscore struct {
	AverageScore  Score
	Neighborhoods []Neighborhood
}

func getPath(location string) {

}

// func Find(location string) (Walkscore, error) {

// }
