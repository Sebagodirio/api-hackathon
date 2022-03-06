package helpers

import (
	"math/rand"
)

func GetRandomState() string {
	states := []string{
		"Alabama",
		"Alaska",
		"Arizona",
		"California",
		"Connecticut",
		"Florida",
		"Colorado",
		"Hawaii",
		"Idaho",
		"Illinois",
		"Kansas",
		"Mississippi",
		"Missouri",
		"Nevada",
		"Montana",
		"Nebraska",
		"New York",
		"Rhode Island",
		"Oregon",
		"Vermont",
	}
	return states[rand.Intn(len(states))]
}

func GetRandomDevelopmentsName() string {
	states := []string{
		"Kent Cockroaches",
		"Charming Cockroaches",
		"Beautiful Frogs",
		"Beautiful Butterflies",
		"Modest Wanderers",
		"Charming Wanderers",
		"Kent Butterflies",
		"Cant think of a name",
		"The best",
		"Are we programming?",
		"We are programming",
		"We Rock",
		"Yeah",
		"Nope",
		"Hello team",
		"OMG team",
		"Yuhu team",
		"Goku team",
		"Gohan ha",
		"Idk team",
	}
	return states[rand.Intn(len(states))]
}

func GetRandomDevelopmentsDescriptions() string {
	states := []string{
		"This is a great project",
		"Really a great project",
		"A project that will save the world",
		"I dont hace more ideas for a good description",
		"Neither of these descriptions are good really",
		"ok I think this is enough",
	}
	return states[rand.Intn(len(states))]
}

func GenerateScore(topScore float64) float64 {
	return rand.Float64() * topScore
}
