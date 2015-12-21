package team

struct Team {
	name string
	division string //probably make ptr to division object
	conference string //probably make ptr to conference object
	wins []*Team
	losses []*Team
}

func team() {
	return 
}