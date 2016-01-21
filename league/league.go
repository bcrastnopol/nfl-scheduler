package league

// import (
// 	"nfl-scheduler/linter"
// )

type League struct {
	name        string
	conferences map[string]*Conference
}

func (league *League) GetConferences() map[string]*Conference {
	return league.conferences
}

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
		con := &Conference{con_name, map[string]*Division{}, league}
		league.GetConferences()[con_name] = con
		for div_name, div_vals := range con_vals {
			div := &Division{div_name, map[string]*Team{}, con, league}
			con.GetDivisions()[div_name] = div
			place := 0
			for _, team_name := range div_vals {
				place++
				div.GetTeams()[team_name] = NewTeam(team_name, div, con, league, place)
			}
		}
	}
	return league
}

func setSchedules(league *League) {
	for c_name, con := range league.GetConferences() {
		var first_div *Division
		for d_name, div := range con.GetDivisions() {
			for t_name, team := range div.GetTeams() {

			}
		}
	}
}
