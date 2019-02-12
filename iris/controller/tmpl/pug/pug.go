package pug

import "github.com/kataras/iris"

func PugTest(ctx iris.Context) {
	type Job struct {
		Employer string
		Role     string
	}
	type Person struct {
		Name   string
		Age    int
		Emails []string
		Jobs   []*Job
	}

	job1 := Job{Employer: "Monash B", Role: "Honorary"}
	job2 := Job{Employer: "Box Hill", Role: "Head of HE"}

	person := Person{
		Name:   "jan",
		Age:    50,
		Emails: []string{"jan@newmarch.name", "jan.newmarch@gmail.com"},
		Jobs:   []*Job{&job1, &job2},
	}

	ctx.View("index.pug", person)
}
