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
