package enums

//QuesDifficulty hide
type QuesDifficulty = int

type level struct {
	EASY   QuesDifficulty
	MEDIUM QuesDifficulty
	HARD   QuesDifficulty
}

// DifficultyLvl for public use
var DifficultyLvl = &level{
	EASY:   0,
	MEDIUM: 1,
	HARD:   2,
}
