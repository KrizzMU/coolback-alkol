package core

type ModLes struct {
	Module
	Lessons []Lesson
}

type LesMd struct {
	Lesson
	Mdfile []string
}

type СourseСontent struct {
	Course
	Modules []ModLes
}