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

func (team *Team) GetName() string {
	return team.name
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

func (team *Team) GetDivision() *Division {
	return team.division
}

func (team *Team) GetConference() *Conference {
	return team.conference
}

func (team *Team) appendWin(win *Team) {
	team.wins = append(team.wins, win)
}

func (team *Team) appendLoss(loss *Team) {
	team.losses = append(team.losses, loss)
}

func (team *Team) appendToSchedule(scheduled *Team) {
	team.schedule = append(team.schedule, scheduled)
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

func (team *Team) PlayGame(opponent *Team, win bool) {
	if team.onSchedule(opponent) {
		log.Fatal("Team isn't on schedule")
	}
	played := 0
	for _, opponent := range append(team.GetWins(), team.GetLosses()...) {
		if team == opponent {
			played++
			if team.GetDivision() == opponent.GetDivision() && played < 1 {
				continue
			}
			log.Fatal("Team has already played 2x")
		}
	}
	if win {
		team.appendWin(opponent)
	} else {
		team.appendLoss(opponent)
	}
}

func NewTeam(
	name string, division *Division, conference *Conference) *Team {
	return &Team{name, division, conference, []*Team{}, []*Team{}, []*Team{}}
}
