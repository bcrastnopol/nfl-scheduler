package league

//name interface
type IName interface {
	GetName() string
}

//get division interface
type IGetDivision interface {
	GetDivision() *Division
}

//get conference interface
type IGetConference interface {
	getConference() *Conference
}

//league interface
type ILeague interface {
	GetConferences() map[string]*Conference
}

//conference interface
type IConference interface {
	GetDivisions() map[string]*Division
}

//division interface
type IDivision interface {
	GetTeams() map[string]*Team
}

//team interface
type ITeam interface {
	PlayGame(team *Team)
	onSchedule(team *Team) bool
	GetWins() []*Team
	GetLosses() []*Team
	GetSchedule() []*Team
	appendWin()
	appendLoss()
	appendToSchedule()
}
