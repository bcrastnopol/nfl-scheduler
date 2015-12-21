package league

type Conference struct {
	name string
	divisions []*Division
}

type Division struct {
	name string
	teams []*Team
	confrence *Conference
}

type Team struct {
	name string
	division *Division
	conference *Conference
	wins []*Team
	losses []*Team
}

func NewTeam (name string, division *Division, conference *Conference) *Team {
	return &Team {name, division, conference, []*Team{}, []*Team{}}
}


func SetUpLeague () map[string]map[string][]string {
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
	// team_data := map[string]map[string]string{
	// 	"a": map[string]string{"a":"b"},
	// 	 "b": map[string]string{"b":"c"},
	// }
	return team_data
}