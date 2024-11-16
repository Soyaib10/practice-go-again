package encapsulation

import (
	"errors"
)

type Student struct {
	Name       string
	age        int
	rollNumber int
	grades     []int
}

// NewStudent is a constructor function of Student
func NewStudent(name string, age int, rollNumber int) *Student {
	return &Student{
		Name:       name,
		age:        age,
		rollNumber: rollNumber,
		grades:     make([]int, 0),
	}
}

// getter methods
func (s *Student) GetAge() int {
	return s.age
}

func (s *Student) GetRollNumber() int {
	return s.rollNumber
}

func (s *Student) GetGrades() []int {
	grades := make([]int, len(s.grades))
	copy(grades, s.grades)
	return grades
}

// setter methods
func (s *Student) SetAge(age int) error {
	if age < 0 || age > 150 {
		return errors.New("unvalid age")
	}
	s.age = age
	return nil
}

func (s *Student) AddGrade(grade int) error {
	if grade < 0 || grade > 100 {
		return errors.New("অবৈধ গ্রেড")
	}
	s.grades = append(s.grades, grade)
	return nil
}

func (s *Student) CalcualteAvageGrade() float64 {
	if len(s.grades) == 0 {
		return 0.0
	}
	sum := 0
	for _, x := range s.grades {
		sum += x
	}
	return float64(sum) / float64(len(s.grades))
}
