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

func (scheduleInfo *scheduleInfo) GetOpponentDivisions() []*Division {
	return scheduleInfo.opponentDivisions
}

func(team *Team) SetOpponentDivision(div *Division) {
	ods := team.GetScheduleInfo().GetOpponentDivisions()
	for _, division := range ods {
		if div == division{
			log.Print("Opponent division already set")
			return
		}  
	}
	td := team.GetDivision() 
	if div != td && len(ods) < 2 {
		if len(ods) == 1 && ods[0].GetConference() == div.GetConference() {
			log.Fatal("Can't play two divisions from the same conference")
		}
		for t_name, t := range td {
			
		}
	}
	scheduleInfo = append(scheduleInfo(div))
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

func (team *Team) limitedEncounters(encountered *Team, encounters []*Team) bool {
	div := team.getDivision()
	for _, opponent := range encounters {
		seen := 0
		if encountered == opponent {
			seen++
			if div == opponent.getDivision() && seen < 1 {
				continue
			}
			return false 
		}
	}
	return true
}

func (team *Team) appendToSchedule(scheduled *Team) {
	sched := team.GetSchedule()
	if team.limitedEncounters(scheduled, team.GetSchedule()) {
		sched = append(sched, scheduled)
	}
	else {
			log.Print("Team has already been scheduled enough times")
	}

}

func (team *Team) onSchedule(scheduled *Team check) bool {
	_ = "breakpoint"
	for _, t := range team.GetSchedule() {
		if scheduled == t {
			return true
		}
	}
	return false
}

func (team *Team) ScheduleOpponent(scheduled *Team) {
		team.appendToSchedule(scheduled)
		scheduled.appendToSchedule(team)
}


func (team *Team) PlayGame(opponent *Team, win bool) {
	if team.onSchedule(opponent) {
		log.Fatal("Team isn't on schedule")
	}
	if team.limitedEncounters(opponent,
	 		append(team.GetWins(), team.GetLosses()...)){
		if win {
			team.appendWin(opponent)
		} else {
			team.appendLoss(opponent)
		}
	}
	else {
			log.Fatal("Team has already played enough times")
	}
}

func NewTeam(
	name string, d *Division, c *Conference, l *League, place int) *Team {
	return &Team{name, d, c, l, []*Team{}, []*Team{},
		&scheduleInfo{[]*Team{}, []*Division{}, place}}
}
