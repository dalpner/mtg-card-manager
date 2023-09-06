package models

type Cards struct {
	Name          string
	Layout        string
	Cmc           int
	Colors        string
	ColorIdentity []string
	Type          string
	Supertypes    string
	Types         []string
	Subtypes      []string
	rarity        string
	Set           string
	SetName       string
	Text          string
	Flavor        string
	Artist        string
	Number        string
	Power         string
	Toughness     string
	Loyalty       string
	Language      string
	GameFormat    []string
	Legality      []string
	Page          int
	PageSize      int
	OrderBy       string
	Random        int
	Contains      []string
	Id            string
	MultiverseId  []string
}
