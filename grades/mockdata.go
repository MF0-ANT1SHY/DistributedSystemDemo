package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "Nick",
			LastName:  "Carter",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 89,
				},
				{
					Title: "Quiz 3",
					Type:  GradeQuiz,
					Score: 92,
				},
			},
		},
		{
			ID:        2,
			FirstName: "komeiji",
			LastName:  "cenrili",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 75,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 91,
				},
				{
					Title: "Quiz 3",
					Type:  GradeQuiz,
					Score: 62,
				},
			},
		},
	}
}
