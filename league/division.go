package league

type Division struct {
	name      string
	teams     map[string]*Team
	confrence *Conference
}

func (division *Division) GetTeams() map[string]*Team {
	return division.teams
}
