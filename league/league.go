package league

type Conference struct {
	name      string
	divisions []*Division
}

type Division struct {
	name      string
	teams     []*Team
	confrence *Conference
}

type Team struct {
	name       string
	division   *Division
	conference *Conference
	wins       []*Team
	losses     []*Team
}

func NewTeam(name string, division *Division, conference *Conference) *Team {
	return &Team{name, division, conference, []*Team{}, []*Team{}}
}

func SetUpLeague() []*Conference {
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
	league := []*Conference{}

	for con_name, con_vals := range team_data {
		con := &Conference{con_name, []*Division{}}
		league = append(league, con)
		for div_name, div_vals := range con_vals {
			div := &Division{div_name, []*Team{}, con}
			con.divisions = append(con.divisions, div)
			for _, team_name := range div_vals {
				div.teams = append(div.teams, NewTeam(team_name, div, con))
			}
		}
	}
	return league
}
