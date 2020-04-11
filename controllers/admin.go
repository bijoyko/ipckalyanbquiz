package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/bijoyko/ipckalyanbquiz/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//admin begins here
func LoadAdmin(c *gin.Context) {
	t, _ := template.ParseFiles("view/admin.html")
	t.Execute(c.Writer, nil)
}

func VerifyAdmin(c *gin.Context) {
	if c.PostForm("Username") == "Bijoy" && c.PostForm("Password") == "Lolol_123" {
		t, err := template.ParseFiles("view/editquestions.html")
		if err != nil {
			log.Println(err)
		}
		t.Execute(c.Writer, nil)
	} else {
		text1 := "You are not authorised to view this page"
		t, err := template.ParseFiles("view/admin.html")
		if err != nil {
			log.Println(err)
		}
		t.Execute(c.Writer, text1)
	}
}

func UpdateQuestions(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var quizD models.QuestionsForm
	if err := db.Table("updatequiz").First(&quizD).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Table("updatequiz").Delete(&quizD)

	var answersD models.Insertanswers
	if err := db.Table("answersquiz").First(&answersD).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Table("answersquiz").Delete(&answersD)

	update := models.UpdateForm{
		Q1e:  c.PostForm("Q1E"),
		Q1m:  c.PostForm("Q1M"),
		Ea1:  c.PostForm("EA1"),
		Eb1:  c.PostForm("EB1"),
		Ec1:  c.PostForm("EC1"),
		Ed1:  c.PostForm("ED1"),
		Ma1:  c.PostForm("MA1"),
		Mb1:  c.PostForm("MB1"),
		Mc1:  c.PostForm("MC1"),
		Md1:  c.PostForm("MD1"),
		Q2e:  c.PostForm("Q2E"),
		Q2m:  c.PostForm("Q2M"),
		Ea2:  c.PostForm("EA2"),
		Eb2:  c.PostForm("EB2"),
		Ec2:  c.PostForm("EC2"),
		Ed2:  c.PostForm("ED2"),
		Ma2:  c.PostForm("MA2"),
		Mb2:  c.PostForm("MB2"),
		Mc2:  c.PostForm("MC2"),
		Md2:  c.PostForm("MD2"),
		Q3e:  c.PostForm("Q3E"),
		Q3m:  c.PostForm("Q3M"),
		Ea3:  c.PostForm("EA3"),
		Eb3:  c.PostForm("EB3"),
		Ec3:  c.PostForm("EC3"),
		Ed3:  c.PostForm("ED3"),
		Ma3:  c.PostForm("MA3"),
		Mb3:  c.PostForm("MB3"),
		Mc3:  c.PostForm("MC3"),
		Md3:  c.PostForm("MD3"),
		Q4e:  c.PostForm("Q4E"),
		Q4m:  c.PostForm("Q4M"),
		Ea4:  c.PostForm("EA4"),
		Eb4:  c.PostForm("EB4"),
		Ec4:  c.PostForm("EC4"),
		Ed4:  c.PostForm("ED4"),
		Ma4:  c.PostForm("MA4"),
		Mb4:  c.PostForm("MB4"),
		Mc4:  c.PostForm("MC4"),
		Md4:  c.PostForm("MD4"),
		Q5e:  c.PostForm("Q5E"),
		Q5m:  c.PostForm("Q5M"),
		Ea5:  c.PostForm("EA5"),
		Eb5:  c.PostForm("EB5"),
		Ec5:  c.PostForm("EC5"),
		Ed5:  c.PostForm("ED5"),
		Ma5:  c.PostForm("MA5"),
		Mb5:  c.PostForm("MB5"),
		Mc5:  c.PostForm("MC5"),
		Md5:  c.PostForm("MD5"),
		Q6e:  c.PostForm("Q6E"),
		Q6m:  c.PostForm("Q6M"),
		Ea6:  c.PostForm("EA6"),
		Eb6:  c.PostForm("EB6"),
		Ec6:  c.PostForm("EC6"),
		Ed6:  c.PostForm("ED6"),
		Ma6:  c.PostForm("MA6"),
		Mb6:  c.PostForm("MB6"),
		Mc6:  c.PostForm("MC6"),
		Md6:  c.PostForm("MD6"),
		Q7e:  c.PostForm("Q7E"),
		Q7m:  c.PostForm("Q7M"),
		Ea7:  c.PostForm("EA7"),
		Eb7:  c.PostForm("EB7"),
		Ec7:  c.PostForm("EC7"),
		Ed7:  c.PostForm("ED7"),
		Ma7:  c.PostForm("MA7"),
		Mb7:  c.PostForm("MB7"),
		Mc7:  c.PostForm("MC7"),
		Md7:  c.PostForm("MD7"),
		Q8e:  c.PostForm("Q8E"),
		Q8m:  c.PostForm("Q8M"),
		Ea8:  c.PostForm("EA8"),
		Eb8:  c.PostForm("EB8"),
		Ec8:  c.PostForm("EC8"),
		Ed8:  c.PostForm("ED8"),
		Ma8:  c.PostForm("MA8"),
		Mb8:  c.PostForm("MB8"),
		Mc8:  c.PostForm("MC8"),
		Md8:  c.PostForm("MD8"),
		Q9e:  c.PostForm("Q9E"),
		Q9m:  c.PostForm("Q9M"),
		Ea9:  c.PostForm("EA9"),
		Eb9:  c.PostForm("EB9"),
		Ec9:  c.PostForm("EC9"),
		Ed9:  c.PostForm("ED9"),
		Ma9:  c.PostForm("MA9"),
		Mb9:  c.PostForm("MB9"),
		Mc9:  c.PostForm("MC9"),
		Md9:  c.PostForm("MD9"),
		Q10e: c.PostForm("Q10E"),
		Q10m: c.PostForm("Q10M"),
		Ea10: c.PostForm("EA10"),
		Eb10: c.PostForm("EB10"),
		Ec10: c.PostForm("EC10"),
		Ed10: c.PostForm("ED10"),
		Ma10: c.PostForm("MA10"),
		Mb10: c.PostForm("MB10"),
		Mc10: c.PostForm("MC10"),
		Md10: c.PostForm("MD10"),
	}

	insertanswers := models.Answers{
		Q1:  c.PostForm("QA1"),
		Q2:  c.PostForm("QA2"),
		Q3:  c.PostForm("QA3"),
		Q4:  c.PostForm("QA4"),
		Q5:  c.PostForm("QA5"),
		Q6:  c.PostForm("QA6"),
		Q7:  c.PostForm("QA7"),
		Q8:  c.PostForm("QA8"),
		Q9:  c.PostForm("QA9"),
		Q10: c.PostForm("QA10"),
	}

	updateQ := models.UpdateForm{
		Q1e:  update.Q1e,
		Q1m:  update.Q1m,
		Ea1:  update.Ea1,
		Eb1:  update.Eb1,
		Ec1:  update.Ec1,
		Ed1:  update.Ed1,
		Ma1:  update.Ma1,
		Mb1:  update.Mb1,
		Mc1:  update.Mc1,
		Md1:  update.Md1,
		Q2e:  update.Q2e,
		Q2m:  update.Q2m,
		Ea2:  update.Ea2,
		Eb2:  update.Eb2,
		Ec2:  update.Ec2,
		Ed2:  update.Ed2,
		Ma2:  update.Ma2,
		Mb2:  update.Mb2,
		Mc2:  update.Mc2,
		Md2:  update.Md2,
		Q3e:  update.Q3e,
		Q3m:  update.Q3m,
		Ea3:  update.Ea3,
		Eb3:  update.Eb3,
		Ec3:  update.Ec3,
		Ed3:  update.Ed3,
		Ma3:  update.Ma3,
		Mb3:  update.Mb3,
		Mc3:  update.Mc3,
		Md3:  update.Md3,
		Q4e:  update.Q4e,
		Q4m:  update.Q4m,
		Ea4:  update.Ea4,
		Eb4:  update.Eb4,
		Ec4:  update.Ec4,
		Ed4:  update.Ed4,
		Ma4:  update.Ma4,
		Mb4:  update.Mb4,
		Mc4:  update.Mc4,
		Md4:  update.Md4,
		Q5e:  update.Q5e,
		Q5m:  update.Q5m,
		Ea5:  update.Ea5,
		Eb5:  update.Eb5,
		Ec5:  update.Ec5,
		Ed5:  update.Ed5,
		Ma5:  update.Ma5,
		Mb5:  update.Mb5,
		Mc5:  update.Mc5,
		Md5:  update.Md5,
		Q6e:  update.Q6e,
		Q6m:  update.Q6m,
		Ea6:  update.Ea6,
		Eb6:  update.Eb6,
		Ec6:  update.Ec6,
		Ed6:  update.Ed6,
		Ma6:  update.Ma6,
		Mb6:  update.Mb6,
		Mc6:  update.Mc6,
		Md6:  update.Md6,
		Q7e:  update.Q7e,
		Q7m:  update.Q7m,
		Ea7:  update.Ea7,
		Eb7:  update.Eb7,
		Ec7:  update.Ec7,
		Ed7:  update.Ed7,
		Ma7:  update.Ma7,
		Mb7:  update.Mb7,
		Mc7:  update.Mc7,
		Md7:  update.Md7,
		Q8e:  update.Q8e,
		Q8m:  update.Q8m,
		Ea8:  update.Ea8,
		Eb8:  update.Eb8,
		Ec8:  update.Ec8,
		Ed8:  update.Ed8,
		Ma8:  update.Ma8,
		Mb8:  update.Mb8,
		Mc8:  update.Mc8,
		Md8:  update.Md8,
		Q9e:  update.Q9e,
		Q9m:  update.Q9m,
		Ea9:  update.Ea9,
		Eb9:  update.Eb9,
		Ec9:  update.Ec9,
		Ed9:  update.Ed9,
		Ma9:  update.Ma9,
		Mb9:  update.Mb9,
		Mc9:  update.Mc9,
		Md9:  update.Md9,
		Q10e: update.Q10e,
		Q10m: update.Q10m,
		Ea10: update.Ea10,
		Eb10: update.Eb10,
		Ec10: update.Ec10,
		Ed10: update.Ed10,
		Ma10: update.Ma10,
		Mb10: update.Mb10,
		Mc10: update.Mc10,
		Md10: update.Md10,
	}

	updateanswers := models.Answers{
		Q1:  insertanswers.Q1,
		Q2:  insertanswers.Q2,
		Q3:  insertanswers.Q3,
		Q4:  insertanswers.Q4,
		Q5:  insertanswers.Q5,
		Q6:  insertanswers.Q6,
		Q7:  insertanswers.Q7,
		Q8:  insertanswers.Q8,
		Q9:  insertanswers.Q9,
		Q10: insertanswers.Q10,
	}

	db.Table("updatequiz").Create(&updateQ)

	db.Table("answersquiz").Create(&updateanswers)
}
