package league

import "log"

type Team struct {
	name       string
	division   *Division
	conference *Conference
	wins       []*Team
	losses     []*Team
	schedule   []*Team
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

func (team *Team) onSchedule(scheduled *Team) bool {
	_ = "breakpoint"
	for _, t := range team.GetSchedule() {
		if scheduled == t {
			return true
		}
	}
	return false
}

func NewTeam(
	name string, division *Division, conference *Conference) *Team {
	return &Team{name, division, conference, []*Team{}, []*Team{}, []*Team{}}
}

func (team *Team) AddWin(defeated *Team) *Team {
	if team.onSchedule(defeated) {
		log.Fatal("Team isn't on schedule")
	}
	played := 0
	for _, w := range team.GetWins() {
		if team == w {
			played++
		}
	}
	return team
}
