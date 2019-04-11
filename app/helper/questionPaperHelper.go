package helper

import (
	utility "intelliq/app/common"
	"intelliq/app/dto"
	"intelliq/app/enums"
	"intelliq/app/model"
	"sort"
)

//GetSectionCountMap generate K,V for requestted questions per section i.e OBJECTIVE-10,SHORT-5 ...
func GetSectionCountMap(sections []dto.QuesLength) map[enums.QuesLength]int {
	sectionCountMap := make(map[enums.QuesLength]int)
	for _, section := range sections {
		sectionCountMap[section.Type] = section.Count
	}
	return sectionCountMap
}

//PrioritiseDifficultyList sort difficulty based on higher percentage i.e MEIDUM-40,EASY-30,HARD-30
func PrioritiseDifficultyList(difficulty []dto.QuesDifficulty) {
	sort.Slice(difficulty, func(i, j int) bool {
		return difficulty[i].Percent > difficulty[j].Percent
	})
}

//PopulateSectionalQuestionMap group available db questions per section i.e OBJECTIVE-ques[] , LONG-ques[] ...
func PopulateSectionalQuestionMap(ques *model.Question,
	sectionQuesMap map[enums.QuesLength]map[enums.QuesDifficulty]model.Questions) {
	lvlMap := sectionQuesMap[ques.Length]
	if lvlMap == nil {
		lvlMap = make(map[enums.QuesDifficulty]model.Questions)
	}
	quesList := lvlMap[ques.Difficulty]
	quesList = append(quesList, *ques)
	lvlMap[ques.Difficulty] = quesList
	sectionQuesMap[ques.Length] = lvlMap
}

//GetResultSectionQuesList get questions for each section passed, split into mulitple sets requested per level
//e.g OBJECTIVE -> {EASY: [],[]},{MEDIUM: [],[]} ,{HARD : [],[]} where set =2
func GetResultSectionQuesList(levelBasedQuesMap map[enums.QuesDifficulty]model.Questions, difficultyList []dto.QuesDifficulty,
	section enums.QuesLength, requestedQuesCount int, requestedSets int, sectionChannel chan<- *dto.Section) { // go routine
	// generate ques count per each difficuly level i.e EASY-10,MEIDUM-20,HARD-6
	levelBasedQuesCountMap := make(map[enums.QuesDifficulty]int)
	for difficulty, quesList := range levelBasedQuesMap {
		levelBasedQuesCountMap[difficulty] = len(quesList)
	}
	//deduce ques count per difficulty level based on percentage requested
	levelQuesProportionMap := getLevelProportionMap(requestedQuesCount, levelBasedQuesCountMap, difficultyList)
	// get questions for each set per level where level = EASY,MEIUM,HARD & set count = requested Sets
	levelQuesSetMap := make(map[enums.QuesDifficulty][]model.Questions)
	for level, deducedCount := range levelQuesProportionMap {
		levelSets := getLevelBasedQuestionSets(levelBasedQuesMap[level], deducedCount, requestedSets)
		levelQuesSetMap[level] = levelSets
	}
	// return section model with section name and per level ques sets map
	sectionChannel <- &dto.Section{
		Type:     section,
		LevelMap: levelQuesSetMap,
	}
}

//deduce ques count per difficulty level based on percentage requested
//e.g requested : EASY-30%,MEDIUM-40%,HARD-30% ; MaxQuesInSection = 10 ;; output : EASY-3,MEIDIUM-4,HARD-3
func getLevelProportionMap(requestedQuesCountPerSection int, levelQuesCountMap map[enums.QuesDifficulty]int,
	difficultyList []dto.QuesDifficulty) map[enums.QuesDifficulty]int {
	quesLeft, levelExhaustCount, availableLevelsCount := requestedQuesCountPerSection, 0, len(levelQuesCountMap)
	levelProportionMap := make(map[enums.QuesDifficulty]int)
	levelPercentMap := make(map[enums.QuesDifficulty]int)
	for _, difficultyLvl := range difficultyList {
		levelPercentMap[difficultyLvl.Level] = difficultyLvl.Percent
	}
	//calculate ques count(lower int value of percent result) as per individual level percentage
	for difficultyLevel, availabelDBQuesCount := range levelQuesCountMap {
		percent := levelPercentMap[difficultyLevel]
		projectedQuesCount := getPercentProportion(percent, requestedQuesCountPerSection)
		if projectedQuesCount < availabelDBQuesCount {
			levelProportionMap[difficultyLevel] = projectedQuesCount
			levelQuesCountMap[difficultyLevel] = availabelDBQuesCount - projectedQuesCount
			quesLeft -= projectedQuesCount
		} else {
			levelProportionMap[difficultyLevel] = availabelDBQuesCount
			levelQuesCountMap[difficultyLevel] = 0
			quesLeft -= availabelDBQuesCount
			levelExhaustCount++
		}
	}
	//if for some reason the quesleft count > 0 , iterate through each lvl again and add 1 quesCount per level until
	//quesLeft == 0 or all levels are maxed out as per question availablity count
	for quesLeft > 0 && levelExhaustCount < availableLevelsCount {
		for _, difficultyLvl := range difficultyList {
			availableQuesCount := levelQuesCountMap[difficultyLvl.Level]
			if availableQuesCount == 0 {
				continue
			}
			reducedCount := availableQuesCount - 1
			levelProportionMap[difficultyLvl.Level] = levelProportionMap[difficultyLvl.Level] + 1
			levelQuesCountMap[difficultyLvl.Level] = reducedCount
			quesLeft--
			if quesLeft == 0 {
				break
			}
			if reducedCount == 0 {
				levelExhaustCount++
			}
		}
	}
	return levelProportionMap
}

func getPercentProportion(percent int, sum int) int {
	return percent * sum / 100
}

//extract actual DB ques for each level and create multiple copies as per set count
func getLevelBasedQuestionSets(availableQuesList model.Questions, deducedCount int,
	requestedSets int) []model.Questions {
	if deducedCount == 0 {
		return nil
	}
	shuffledIndexArr := getShuffleIndexArray(len(availableQuesList), deducedCount*requestedSets)
	var questionSets []model.Questions
	var quesListPerSet model.Questions
	startIndex, shuffleLen := 0, len(shuffledIndexArr)
	for requestedSets > 0 {
		for {
			quesListPerSet = append(quesListPerSet, availableQuesList[shuffledIndexArr[startIndex%shuffleLen]])
			startIndex++
			if startIndex%deducedCount == 0 {
				break
			}
		}
		questionSets = append(questionSets, quesListPerSet)
		quesListPerSet = nil
		requestedSets--
	}
	return questionSets
}

func getShuffleIndexArray(totalQuestions, requiredQuestions int) []int {
	shuffledLenth := utility.GetMin(totalQuestions, requiredQuestions)
	var shuffledArr []int
	indexMap := make(map[int]struct{})
	size := 0
	for size < shuffledLenth {
		index := utility.GenerateRandom(0, totalQuestions)
		if _, ok := indexMap[index]; ok {
			continue
		}
		shuffledArr = append(shuffledArr, index)
		indexMap[index] = struct{}{}
		size++
	}
	return shuffledArr
}

//GenerateQuestionPaper generated ques paper from given section
func GenerateQuestionPaper(sectionList []dto.Section, currentSet int,
	difficultyList []dto.QuesDifficulty, paperChannel chan<- *dto.QuestionPaperDto) {
	var sections []dto.Section
	var sectionalQuesList model.Questions
	for _, section := range sectionList {
		for _, difficulty := range difficultyList {
			if quesSets, ok := section.LevelMap[difficulty.Level]; ok {
				if len(quesSets) >= (currentSet + 1) {
					sectionalQuesList = append(sectionalQuesList, section.LevelMap[difficulty.Level][currentSet]...)
				}
			}
		}
		sections = append(sections, dto.Section{Type: section.Type, Questions: sectionalQuesList})
		sectionalQuesList = nil
	}
	paperChannel <- &dto.QuestionPaperDto{Set: currentSet + 1, Sections: sections}
}
