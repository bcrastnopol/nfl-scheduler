package league

type Conference struct {
	name      string
	divisions map[string]*Division
	league    *League
}

func (conference *Conference) GetDivisions() map[string]*Division {
	return conference.divisions
}

func (conference *Conference) GetName() string {
	return conference.name
}

func (conference *Conference) getLeague() *League {
	return conference.league
}
