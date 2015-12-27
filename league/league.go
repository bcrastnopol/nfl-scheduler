package league

// type Ninterface interface {
// 	GetName() string
// }

//league interface
type Linterface interface {
	GetConferences() map[string]*Conference
}

type League struct {
	name        string
	conferences map[string]*Conference
}

func (league *League) GetConferences() map[string]*Conference {
	return league.conferences
}

//conference interface
type Cinterface interface {
	GetDivisions() map[string]*Division
}

type Conference struct {
	name      string
	divisions map[string]*Division
}

func (conference *Conference) GetDivisions() map[string]*Division {
	return conference.divisions
}

//division interface
type Dinterface interface {
	GetTeams() map[string]*Team
}

type Division struct {
	name      string
	teams     map[string]*Team
	confrence *Conference
}

func (division *Division) GetTeams() map[string]*Team {
	return division.teams
}

//team interface
type Tinterface interface {
	AddWin(team *Team) *Team
	// AddLoss() *Team
	onSchedule(team *Team) bool
	GetWins() []*Team
	GetLosser() []*Team
	GetSchedule() []*Team
}

type Team struct {
	name       string
	division   *Division
	conference *Conference
	wins       []*Team
	losses     []*Team
	schedule   []*Team
}

func NewTeam(
	name string, division *Division, conference *Conference) *Team {
	return &Team{name, division, conference, []*Team{}, []*Team{}, []*Team{}}
}

func (team *Team) GetWins() []*Team {
	return team.wins
}

func (team *Team) GetLosses() []*Team {
	return team.losses
}

func (team *Team) GetSchedule() []*Team {
	return team.schedule
}

// func (team *Team) onSchedule(scheduled *Team) bool {
// 	_ = "breakpoint"
// 	for _, t := range team.GetSchedule() {
// 		if scheduled == t {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (team *Team) AddWin(defeated *Team) *Team{
// 	if team.onSchedule(defeated){
// 		log.Fatal("Team isn't on schedule")
// 	}
// 	played := 0
// 	for _, w :range team.GetWins() {
// 		if team == w {
// 			played++
// 			if
// 		}
// 	}
// }

// func (team *Team) GetName() string {
// 	return team.name
// }

func SetUpLeague(name string) *League {
	//XXX very NFL specific, but I imagine this could be adapted for other league
	//configurations
	team_data := map[string]map[string][]string{
		"NFC": map[string][]string{
			"North": []string{
				"Packers", "Vikings", "Lions", "Bears",
			},
			"South": []string{
				"Falcons", "Saints", "Buccaneers", "Panthers",
			},
			"East": []string{
				"Eagles", "Cowboys", "Giants", "Redskinds",
			},
			"West": []string{
				"Cardinals", "Seahawks", "49ers", "Rams",
			},
		},
		"AFC": map[string][]string{
			"North": []string{
				"Steelers", "Bengals", "Browns", "Ravens",
			},
			"South": []string{
				"Titans", "Colts", "Texans", "Jaguars",
			},
			"East": []string{
				"Jets", "Dolphins", "Patriots", "Bills",
			},
			"West": []string{
				"Chargers", "Raiders", "Broncos", "Chiefs",
			},
		},
	}
	league := &League{name, map[string]*Conference{}}
	for con_name, con_vals := range team_data {
		con := &Conference{con_name, map[string]*Division{}}
		league.GetConferences()[con_name] = con
		for div_name, div_vals := range con_vals {
			div := &Division{div_name, map[string]*Team{}, con}
			con.GetDivisions()[div_name] = div
			for _, team_name := range div_vals {
				div.GetTeams()[team_name] = NewTeam(team_name, div, con)
			}
		}
	}

	return league
}
