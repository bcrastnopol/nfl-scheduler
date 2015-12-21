package team

type Team struct {
	name string
	division string //probably make ptr to division object
	conference string //probably make ptr to conference object
	wins []*Team
	losses []*Team
}
func NewTeam (name, division, conference string) *Team {
	return &Team {name, division, conference, []*Team{}, []*Team{}}
}