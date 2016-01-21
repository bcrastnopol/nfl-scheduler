package league

//Struct Interfaces

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
	GetWins() []*Team
	GetLosses() []*Team
	GetScheduleInfo() *scheduleInfo
	GetOpponentDivisions() []*Division
	setOpponentDivisions()
	onSchedule(team *Team) bool
	appendWin()
	appendLoss()
}

type IScheduleInfo interface {
	ScheduleOpponent()
	appendToSchedule()
	// makes sure a team never encounters a non division opponent more than once
	// and a division opponent more than twice
	limitedEncounters() bool
}

//Getters For Struct Properties

//name interface
type IGetName interface {
	GetName() string
}

//get league interface
type IGetLeague interface {
	GetLeague() *League
}

//get conference interface
type IGetConference interface {
	getConference() *Conference
}

//get division interface
type IGetDivision interface {
	GetDivision() *Division
}

//get team schedule interface
type IGetTeamSchedule interface {
	GetSchedule() map[string]*Team
}
