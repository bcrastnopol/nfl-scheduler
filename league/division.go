package league

type Division struct {
	name      string
	teams     map[string]*Team
	confrence *Conference
	league    *League
}

func (division *Division) GetTeams() map[string]*Team {
	return division.teams
}

func (division *Division) GetName() string {
	return division.name
}

func (division *Division) GetConference() *Conference {
	return division.confrence
}

func (division *Division) getLeague() *League {
	return division.league
}
