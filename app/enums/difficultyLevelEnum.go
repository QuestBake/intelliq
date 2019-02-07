package enums

//Difficulty hide
type Difficulty = int

type level struct {
	EASY   Difficulty
	MEDIUM Difficulty
	HARD   Difficulty
}

// DifficultyLvl for public use
var DifficultyLvl = &level{
	EASY:   0,
	MEDIUM: 1,
	HARD:   2,
}
