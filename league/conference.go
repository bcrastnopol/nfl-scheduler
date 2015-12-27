package league

type Conference struct {
	name      string
	divisions map[string]*Division
}

func (conference *Conference) GetDivisions() map[string]*Division {
	return conference.divisions
}