package models

type Quiz struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Language  string `json:"language"`
	Q1        string `json:"q1"`
	Q2        string `json:"q2"`
	Q3        string `json:"q3"`
	Q4        string `json:"q4"`
	Q5        string `json:"q5"`
	Q6        string `json:"q6"`
	Q7        string `json:"q7"`
	Q8        string `json:"q8"`
	Q9        string `json:"q9"`
	Q10       string `json:"q10"`
	Points    int    `json:"points"`
}

type EnterQuizData struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Language  string `json:"language" binding:"required"`
	Q1        string `json:"q1" binding:"required"`
	Q2        string `json:"q2" binding:"required"`
	Q3        string `json:"q3" binding:"required"`
	Q4        string `json:"q4" binding:"required"`
	Q5        string `json:"q5" binding:"required"`
	Q6        string `json:"q6" binding:"required"`
	Q7        string `json:"q7" binding:"required"`
	Q8        string `json:"q8" binding:"required"`
	Q9        string `json:"q9" binding:"required"`
	Q10       string `json:"q10" binding:"required"`
}

type Names struct {
	Firstname string
	Lastname  string
	Language  string
}

type Insertanswers struct {
	ID  int    `json:"id" gorm:"primary_key"`
	Q1  string `json:"q1"`
	Q2  string `json:"q2"`
	Q3  string `json:"q3"`
	Q4  string `json:"q4"`
	Q5  string `json:"q5"`
	Q6  string `json:"q6"`
	Q7  string `json:"q7"`
	Q8  string `json:"q8"`
	Q9  string `json:"q9"`
	Q10 string `json:"`
}

type Answers struct {
	Q1  string `json:"q1" binding:"required"`
	Q2  string `json:"q2" binding:"required"`
	Q3  string `json:"q3" binding:"required"`
	Q4  string `json:"q4" binding:"required"`
	Q5  string `json:"q5" binding:"required"`
	Q6  string `json:"q6" binding:"required"`
	Q7  string `json:"q7" binding:"required"`
	Q8  string `json:"q8" binding:"required"`
	Q9  string `json:"q9" binding:"required"`
	Q10 string `json:"q10" binding:"required"`
}

type QuestionsForm struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Q1e  string `json:"q1e"`
	Q1m  string `json:"q1m"`
	Ea1  string `json:"ea1"`
	Eb1  string `json:"eb1"`
	Ec1  string `json:"ec1"`
	Ed1  string `json:"ed1"`
	Ma1  string `json:"ma1"`
	Mb1  string `json:"mb1"`
	Mc1  string `json:"mc1"`
	Md1  string `json:"md1"`
	Q2e  string `json:"q2e"`
	Q2m  string `json:"q2m"`
	Ea2  string `json:"ea2"`
	Eb2  string `json:"eb2"`
	Ec2  string `json:"ec2"`
	Ed2  string `json:"ed2"`
	Ma2  string `json:"ma2"`
	Mb2  string `json:"mb2"`
	Mc2  string `json:"mc2"`
	Md2  string `json:"md2"`
	Q3e  string `json:"q3e"`
	Q3m  string `json:"q3m"`
	Ea3  string `json:"ea3"`
	Eb3  string `json:"eb3"`
	Ec3  string `json:"ec3"`
	Ed3  string `json:"ed3"`
	Ma3  string `json:"ma3"`
	Mb3  string `json:"mb3"`
	Mc3  string `json:"mc3"`
	Md3  string `json:"md3"`
	Q4e  string `json:"q4e"`
	Q4m  string `json:"q4m"`
	Ea4  string `json:"ea4"`
	Eb4  string `json:"eb4"`
	Ec4  string `json:"ec4"`
	Ed4  string `json:"ed4"`
	Ma4  string `json:"ma4"`
	Mb4  string `json:"mb4"`
	Mc4  string `json:"mc4"`
	Md4  string `json:"md4"`
	Q5e  string `json:"q5e"`
	Q5m  string `json:"q5m"`
	Ea5  string `json:"ea5"`
	Eb5  string `json:"eb5"`
	Ec5  string `json:"ec5"`
	Ed5  string `json:"ed5"`
	Ma5  string `json:"ma5"`
	Mb5  string `json:"mb5"`
	Mc5  string `json:"mc5"`
	Md5  string `json:"md5"`
	Q6e  string `json:"q6e"`
	Q6m  string `json:"q6m"`
	Ea6  string `json:"ea6"`
	Eb6  string `json:"eb6"`
	Ec6  string `json:"ec6"`
	Ed6  string `json:"ed6"`
	Ma6  string `json:"ma6"`
	Mb6  string `json:"mb6"`
	Mc6  string `json:"mc6"`
	Md6  string `json:"md6"`
	Q7e  string `json:"q7e"`
	Q7m  string `json:"q7m"`
	Ea7  string `json:"ea7"`
	Eb7  string `json:"eb7"`
	Ec7  string `json:"ec7"`
	Ed7  string `json:"ed7"`
	Ma7  string `json:"ma7"`
	Mb7  string `json:"mb7"`
	Mc7  string `json:"mc7"`
	Md7  string `json:"md7"`
	Q8e  string `json:"q8e"`
	Q8m  string `json:"q8m"`
	Ea8  string `json:"ea8"`
	Eb8  string `json:"eb8"`
	Ec8  string `json:"ec8"`
	Ed8  string `json:"ed8"`
	Ma8  string `json:"ma8"`
	Mb8  string `json:"mb8"`
	Mc8  string `json:"mc8"`
	Md8  string `json:"md8"`
	Q9e  string `json:"q9e"`
	Q9m  string `json:"q9m"`
	Ea9  string `json:"ea9"`
	Eb9  string `json:"eb9"`
	Ec9  string `json:"ec9"`
	Ed9  string `json:"ed9"`
	Ma9  string `json:"ma9"`
	Mb9  string `json:"mb9"`
	Mc9  string `json:"mc9"`
	Md9  string `json:"md9"`
	Q10e string `json:"q10e"`
	Q10m string `json:"q10m"`
	Ea10 string `json:"ea10"`
	Eb10 string `json:"eb10"`
	Ec10 string `json:"ec10"`
	Ed10 string `json:"ed10"`
	Ma10 string `json:"ma10"`
	Mb10 string `json:"mb10"`
	Mc10 string `json:"mc10"`
	Md10 string `json:"md10"`
}

type UpdateForm struct {
	Q1e  string `json:"q1e" binding:"required"`
	Q1m  string `json:"q1m" binding:"required"`
	Ea1  string `json:"ea1" binding:"required"`
	Eb1  string `json:"eb1" binding:"required"`
	Ec1  string `json:"ec1" binding:"required"`
	Ed1  string `json:"ed1" binding:"required"`
	Ma1  string `json:"ma1" binding:"required"`
	Mb1  string `json:"mb1" binding:"required"`
	Mc1  string `json:"mc1" binding:"required"`
	Md1  string `json:"md1" binding:"required"`
	Q2e  string `json:"q2e" binding:"required"`
	Q2m  string `json:"q2m" binding:"required"`
	Ea2  string `json:"ea2" binding:"required"`
	Eb2  string `json:"eb2" binding:"required"`
	Ec2  string `json:"ec2" binding:"required"`
	Ed2  string `json:"ed2" binding:"required"`
	Ma2  string `json:"ma2" binding:"required"`
	Mb2  string `json:"mb2" binding:"required"`
	Mc2  string `json:"mc2" binding:"required"`
	Md2  string `json:"md2" binding:"required"`
	Q3e  string `json:"q3e" binding:"required"`
	Q3m  string `json:"q3m" binding:"required"`
	Ea3  string `json:"ea3" binding:"required"`
	Eb3  string `json:"eb3" binding:"required"`
	Ec3  string `json:"ec3" binding:"required"`
	Ed3  string `json:"ed3" binding:"required"`
	Ma3  string `json:"ma3" binding:"required"`
	Mb3  string `json:"mb3" binding:"required"`
	Mc3  string `json:"mc3" binding:"required"`
	Md3  string `json:"md3" binding:"required"`
	Q4e  string `json:"q4e" binding:"required"`
	Q4m  string `json:"q4m" binding:"required"`
	Ea4  string `json:"ea4" binding:"required"`
	Eb4  string `json:"eb4" binding:"required"`
	Ec4  string `json:"ec4" binding:"required"`
	Ed4  string `json:"ed4" binding:"required"`
	Ma4  string `json:"ma4" binding:"required"`
	Mb4  string `json:"mb4" binding:"required"`
	Mc4  string `json:"mc4" binding:"required"`
	Md4  string `json:"md4" binding:"required"`
	Q5e  string `json:"q5e" binding:"required"`
	Q5m  string `json:"q5m" binding:"required"`
	Ea5  string `json:"ea5" binding:"required"`
	Eb5  string `json:"eb5" binding:"required"`
	Ec5  string `json:"ec5" binding:"required"`
	Ed5  string `json:"ed5" binding:"required"`
	Ma5  string `json:"ma5" binding:"required"`
	Mb5  string `json:"mb5" binding:"required"`
	Mc5  string `json:"mc5" binding:"required"`
	Md5  string `json:"md5" binding:"required"`
	Q6e  string `json:"q6e" binding:"required"`
	Q6m  string `json:"q6m" binding:"required"`
	Ea6  string `json:"ea6" binding:"required"`
	Eb6  string `json:"eb6" binding:"required"`
	Ec6  string `json:"ec6" binding:"required"`
	Ed6  string `json:"ed6" binding:"required"`
	Ma6  string `json:"ma6" binding:"required"`
	Mb6  string `json:"mb6" binding:"required"`
	Mc6  string `json:"mc6" binding:"required"`
	Md6  string `json:"md6" binding:"required"`
	Q7e  string `json:"q7e" binding:"required"`
	Q7m  string `json:"q7m" binding:"required"`
	Ea7  string `json:"ea7" binding:"required"`
	Eb7  string `json:"eb7" binding:"required"`
	Ec7  string `json:"ec7" binding:"required"`
	Ed7  string `json:"ed7" binding:"required"`
	Ma7  string `json:"ma7" binding:"required"`
	Mb7  string `json:"mb7" binding:"required"`
	Mc7  string `json:"mc7" binding:"required"`
	Md7  string `json:"md7" binding:"required"`
	Q8e  string `json:"q8e" binding:"required"`
	Q8m  string `json:"q8m" binding:"required"`
	Ea8  string `json:"ea8" binding:"required"`
	Eb8  string `json:"eb8" binding:"required"`
	Ec8  string `json:"ec8" binding:"required"`
	Ed8  string `json:"ed8" binding:"required"`
	Ma8  string `json:"ma8" binding:"required"`
	Mb8  string `json:"mb8" binding:"required"`
	Mc8  string `json:"mc8" binding:"required"`
	Md8  string `json:"md8" binding:"required"`
	Q9e  string `json:"q9e" binding:"required"`
	Q9m  string `json:"q9m" binding:"required"`
	Ea9  string `json:"ea9" binding:"required"`
	Eb9  string `json:"eb9" binding:"required"`
	Ec9  string `json:"ec9" binding:"required"`
	Ed9  string `json:"ed9" binding:"required"`
	Ma9  string `json:"ma9" binding:"required"`
	Mb9  string `json:"mb9" binding:"required"`
	Mc9  string `json:"mc9" binding:"required"`
	Md9  string `json:"md9" binding:"required"`
	Q10e string `json:"q10e" binding:"required"`
	Q10m string `json:"q10m" binding:"required"`
	Ea10 string `json:"ea10" binding:"required"`
	Eb10 string `json:"eb10" binding:"required"`
	Ec10 string `json:"ec10" binding:"required"`
	Ed10 string `json:"ed10" binding:"required"`
	Ma10 string `json:"ma10" binding:"required"`
	Mb10 string `json:"mb10" binding:"required"`
	Mc10 string `json:"mc10" binding:"required"`
	Md10 string `json:"md10" binding:"required"`
}

type Scorers struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Points    int    `json:"points"`
}
