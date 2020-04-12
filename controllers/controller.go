package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/bijoyko/ipckalyanbquiz/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var Name models.Names
var Score int
var pulledquestions []string
var p1 models.QuestionsForm

type participation struct {
	Name   models.Names
	Score1 int
}

func MainPage(c *gin.Context) {
	t, _ := template.ParseFiles("view/main.html")
	t.Execute(c.Writer, nil)
}

func ValidateNames(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var questions models.QuestionsForm

	if err := db.Table("updatequiz").First(&questions).Error; err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	//questions pulled up from database which has right answers
	pulledquestions = []string{questions.Q1e, questions.Q1m, questions.Ea1, questions.Eb1, questions.Ec1,
		questions.Ed1, questions.Ma1, questions.Mb1, questions.Mc1, questions.Md1, questions.Q2e, questions.Q2m,
		questions.Ea2, questions.Eb2, questions.Ec2, questions.Ed2, questions.Ma2, questions.Mb2, questions.Mc2,
		questions.Md2, questions.Q3e, questions.Q3m, questions.Ea3, questions.Eb3, questions.Ec3, questions.Ed3,
		questions.Ma3, questions.Mb3, questions.Mc3, questions.Md3, questions.Q4e, questions.Q4m, questions.Ea4,
		questions.Eb4, questions.Ec4, questions.Ed4, questions.Ma4, questions.Mb4, questions.Mc4, questions.Md4,
		questions.Q5e, questions.Q5m, questions.Ea5, questions.Eb5, questions.Ec5, questions.Ed5, questions.Ma5,
		questions.Mb5, questions.Mc5, questions.Md5, questions.Q6e, questions.Q6m, questions.Ea6, questions.Eb6,
		questions.Ec6, questions.Ed6, questions.Ma6, questions.Mb6, questions.Mc6, questions.Md6, questions.Q7e,
		questions.Q7m, questions.Ea7, questions.Eb7, questions.Ec7, questions.Ed7, questions.Ma7, questions.Mb7,
		questions.Mc7, questions.Md7, questions.Q8e, questions.Q8m, questions.Ea8, questions.Eb8, questions.Ec8,
		questions.Ed8, questions.Ma8, questions.Mb8, questions.Mc8, questions.Md8, questions.Q9e, questions.Q9m,
		questions.Ea9, questions.Eb9, questions.Ec9, questions.Ed9, questions.Ma9, questions.Mb9, questions.Mc9,
		questions.Md9, questions.Q10e, questions.Q10m, questions.Ea10, questions.Eb10, questions.Ec10, questions.Ed10,
		questions.Ma10, questions.Mb10, questions.Mc10, questions.Md10}

	p1 = models.QuestionsForm{
		Q1e:  questions.Q1e,
		Q1m:  questions.Q1m,
		Ea1:  questions.Ea1,
		Eb1:  questions.Eb1,
		Ec1:  questions.Ec1,
		Ed1:  questions.Ed1,
		Ma1:  questions.Ma1,
		Mb1:  questions.Mb1,
		Mc1:  questions.Mc1,
		Md1:  questions.Md1,
		Q2e:  questions.Q2e,
		Q2m:  questions.Q2m,
		Ea2:  questions.Ea2,
		Eb2:  questions.Eb2,
		Ec2:  questions.Ec2,
		Ed2:  questions.Ed2,
		Ma2:  questions.Ma2,
		Mb2:  questions.Mb2,
		Mc2:  questions.Mc2,
		Md2:  questions.Md2,
		Q3e:  questions.Q3e,
		Q3m:  questions.Q3m,
		Ea3:  questions.Ea3,
		Eb3:  questions.Eb3,
		Ec3:  questions.Ec3,
		Ed3:  questions.Ed3,
		Ma3:  questions.Ma3,
		Mb3:  questions.Mb3,
		Mc3:  questions.Mc3,
		Md3:  questions.Md3,
		Q4e:  questions.Q4e,
		Q4m:  questions.Q4m,
		Ea4:  questions.Ea4,
		Eb4:  questions.Eb4,
		Ec4:  questions.Ec4,
		Ed4:  questions.Ed4,
		Ma4:  questions.Ma4,
		Mb4:  questions.Mb4,
		Mc4:  questions.Mc4,
		Md4:  questions.Md4,
		Q5e:  questions.Q5e,
		Q5m:  questions.Q5m,
		Ea5:  questions.Ea5,
		Eb5:  questions.Eb5,
		Ec5:  questions.Ec5,
		Ed5:  questions.Ed5,
		Ma5:  questions.Ma5,
		Mb5:  questions.Mb5,
		Mc5:  questions.Mc5,
		Md5:  questions.Md5,
		Q6e:  questions.Q6e,
		Q6m:  questions.Q6m,
		Ea6:  questions.Ea6,
		Eb6:  questions.Eb6,
		Ec6:  questions.Ec6,
		Ed6:  questions.Ed6,
		Ma6:  questions.Ma6,
		Mb6:  questions.Mb6,
		Mc6:  questions.Mc6,
		Md6:  questions.Md6,
		Q7e:  questions.Q7e,
		Q7m:  questions.Q7m,
		Ea7:  questions.Ea7,
		Eb7:  questions.Eb7,
		Ec7:  questions.Ec7,
		Ed7:  questions.Ed7,
		Ma7:  questions.Ma7,
		Mb7:  questions.Mb7,
		Mc7:  questions.Mc7,
		Md7:  questions.Md7,
		Q8e:  questions.Q8e,
		Q8m:  questions.Q8m,
		Ea8:  questions.Ea8,
		Eb8:  questions.Eb8,
		Ec8:  questions.Ec8,
		Ed8:  questions.Ed8,
		Ma8:  questions.Ma8,
		Mb8:  questions.Mb8,
		Mc8:  questions.Mc8,
		Md8:  questions.Md8,
		Q9e:  questions.Q9e,
		Q9m:  questions.Q9m,
		Ea9:  questions.Ea9,
		Eb9:  questions.Eb9,
		Ec9:  questions.Ec9,
		Ed9:  questions.Ed9,
		Ma9:  questions.Ma9,
		Mb9:  questions.Mb9,
		Mc9:  questions.Mc9,
		Md9:  questions.Md9,
		Q10e: questions.Q10e,
		Q10m: questions.Q10m,
		Ea10: questions.Ea10,
		Eb10: questions.Eb10,
		Ec10: questions.Ec10,
		Ed10: questions.Ed10,
		Ma10: questions.Ma10,
		Mb10: questions.Mb10,
		Mc10: questions.Mc10,
		Md10: questions.Md10,
	}

	if c.PostForm("Firstname") != "" && c.PostForm("Lastname") != "" && c.PostForm("Language") != "" {
		Name = models.Names{
			Firstname: c.PostForm("Firstname"),
			Lastname:  c.PostForm("Lastname"),
			Language:  c.PostForm("Language"),
		}
		if c.PostForm("Language") == "English" {
			t, _ := template.ParseFiles("view/FormEnglish.html")
			t.Execute(c.Writer, p1)
		} else {
			t, _ := template.ParseFiles("view/FormMalayalam.html")
			t.Execute(c.Writer, p1)
		}

	} else {
		text1 := "Please enter a valid first Name and last name. Also select a language of your choice to proceed\t Blank values are not allowed"
		t, _ := template.ParseFiles("view/main.html")
		t.Execute(c.Writer, text1)
	}
}

func Form(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	input := models.EnterQuizData{
		Firstname: Name.Firstname,
		Lastname:  Name.Lastname,
		Language:  Name.Language,
		Q1:        c.PostForm("Q1"),
		Q2:        c.PostForm("Q2"),
		Q3:        c.PostForm("Q3"),
		Q4:        c.PostForm("Q4"),
		Q5:        c.PostForm("Q5"),
		Q6:        c.PostForm("Q6"),
		Q7:        c.PostForm("Q7"),
		Q8:        c.PostForm("Q8"),
		Q9:        c.PostForm("Q9"),
		Q10:       c.PostForm("Q10"),
	}

	//answers submitted by the client
	checker := []string{c.PostForm("Q1"), c.PostForm("Q2"), c.PostForm("Q3"), c.PostForm("Q4"), c.PostForm("Q5"), c.PostForm("Q6"), c.PostForm("Q7"),
		c.PostForm("Q8"), c.PostForm("Q9"), c.PostForm("Q10")}

	//method to fetch correct answers from the db from another table
	var quiz models.Quiz
	if err := db.Table("answersquiz").First(&quiz).Error; err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	//answers pulled up from database which has right answers
	givenAnswers := []string{quiz.Q1, quiz.Q2, quiz.Q3, quiz.Q4, quiz.Q5, quiz.Q6, quiz.Q7, quiz.Q8, quiz.Q9, quiz.Q10}
	Score = 0
	// compare answers given by client with the right answers and increase the score
	for i, _ := range checker {
		if givenAnswers[i] == checker[i] {
			Score++
		} else {
		}
	}

	//insert quiz data into db
	quizD := models.Quiz{Firstname: input.Firstname, Lastname: input.Lastname, Language: input.Language, Q1: input.Q1, Q2: input.Q2,
		Q3: input.Q3, Q4: input.Q4, Q5: input.Q5, Q6: input.Q6, Q7: input.Q7, Q8: input.Q8, Q9: input.Q9, Q10: input.Q10, Points: Score}
	db.Table("quiztable").Create(&quizD)

	participant := participation{
		Name: models.Names{
			Firstname: Name.Firstname,
			Lastname:  Name.Lastname,
		},
		Score1: Score,
	}

	t, _ := template.ParseFiles("view/ThankYou.html")
	t.Execute(c.Writer, participant)
}

func ScoreView(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var quizF []models.Scorers
	db.Table("quiztable").Select("firstname, lastname, points").Where("points = ?", "10").Find(&quizF)
	t, err := template.ParseFiles("view/Topscorers.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(c.Writer, quizF)
}

func Quiztable(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var quizT []models.Quiz
	db.Table("quiztable").Find(&quizT)
	t, err := template.ParseFiles("view/Quiztable.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(c.Writer, quizT)
}
