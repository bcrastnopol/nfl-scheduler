package league

import "log"

type Team struct {
	name         string
	division     *Division
	conference   *Conference
	league       *League
	wins         []*Team
	losses       []*Team
	scheduleInfo *scheduleInfo
}

type scheduleInfo struct {
	schedule []*Team
	// A little bit of a hack to reuse the getDivision interface
	// Same division name in both conferences
	opponentDivisions []*Division
	place             int
}

func (scheduleInfo *scheduleInfo) GetSchedule() []*Team {
	return scheduleInfo.schedule
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

func (team *Team) GetScheduleInfo() *scheduleInfo {
	return team.scheduleInfo
}

func (team *Team) GetSchedule() []*Team {
	return team.GetScheduleInfo().schedule
}

func (team *Team) GetOpponentDivisions() []*Division {
	return team.GetScheduleInfo().opponentDivisions
}

func (team *Team) GetDivision() *Division {
	return team.division
}

func (team *Team) GetConference() *Conference {
	return team.conference
}

func (team *Team) GetLeague() *League {
	return team.league
}

func (team *Team) appendWin(win *Team) {
	team.wins = append(team.wins, win)
}

func (team *Team) appendLoss(loss *Team) {
	team.losses = append(team.losses, loss)
}

func (team *Team) appendToSchedule(scheduled *Team) {
	sched := team.GetScheduleInfo().GetSchedule()
	sched = append(sched, scheduled)
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
	name string, d *Division, c *Conference, l *League, place int) *Team {
	return &Team{name, d, c, l, []*Team{}, []*Team{},
		&scheduleInfo{[]*Team{}, []*Division{}, place}}
}
