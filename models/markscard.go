package models

//Markscard consists of marks  report
type Markscard struct {
	ExamID    int
	StudentID int
	SubID     int
	Marks     int
	Grade     byte
}

//Subject details
type Subject struct {
	SubjectID int
	SubName   string
	FacultyID int
}

//Assignment details
type Assignment struct {
	AssignmentID int
	StudentID    int
	Status       string
	FacultyID    int
}
