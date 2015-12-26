package league

//league interface
type Linterface interface {
	GetConferences() map[string]Cinterface
}

type League struct {
	name        string
	conferences map[string]Cinterface
}

func (league *League) GetConferences() map[string]Cinterface {
	return league.conferences
}

//conference interface
type Cinterface interface {
	GetDivisions() map[string]Dinterface
}

type Conference struct {
	name      string
	divisions map[string]Dinterface
}

func (conference *Conference) GetDivisions() map[string]Dinterface {
	return conference.divisions
}

//division interface
type Dinterface interface {
	GetTeams() map[string]Tinterface
}

type Division struct {
	name      string
	teams     map[string]Tinterface
	confrence *Conference
}

func (division *Division) GetTeams() map[string]Tinterface {
	return division.teams
}

//team interface
type Tinterface interface {
	// AddWin() *Team
	// AddLoss() *Team
	// onSchedule() bool
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

func onSchedule(team, scheduled *Team) bool {
	_ = "breakpoint"
	for _, t := range team.schedule {
		if scheduled == t {
			return true
		}
	}
	return false
}

// func AddWin(team, defeated *Team) *Team{
// 	if team
// 		log.Fatal("Team isn't on schedule")
// 	if team.division == defeated.division
// 	for _, w :range.team {
// 		if team == w {

// 		}
// 	}
// }

func SetUpLeague(name string) Linterface {
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
	var league Linterface = &League{name, map[string]Cinterface{}}
	for con_name, con_vals := range team_data {
		con := &Conference{con_name, map[string]Dinterface{}}
		league.GetConferences()[con_name] = con
		for div_name, div_vals := range con_vals {
			div := &Division{div_name, map[string]Tinterface{}, con}
			con.GetDivisions()[div_name] = div
			for _, team_name := range div_vals {
				div.GetTeams()[team_name] = NewTeam(team_name, div, con)
			}
		}
	}

	return league
}
