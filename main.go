package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Student struct {
	Id           int     `json:"id"`
	FirstName    string  `json:"firstName"`
	LastName     string  `json:"lastName"`
	Age          int     `json:"age"`
	Gpa          float32 `json:"gpa"`
	DepartmentId int     `json:"departmentId"`
}
type Department struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type Course struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Credit       int    `json:"credit"`
	DepartmentId int    `json:"departmentId"`
}

type Instructor struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}
type Section struct {
	Id           int `json:"id"`
	CourseId     int `json:"courseId"`
	Number       int `json:"number"`
	InstructorId int `json:"instructorId"`
}
type Enrollment struct {
	Id        int `json:"id"`
	StudentId int `json:"studentId"`
	SectionId int `json:"sectionId"`
}

type KomplexStudent struct {
	Id         int              `json:"id"`
	FirstName  string           `json:"firstName"`
	LastName   string           `json:"lastName"`
	Age        int              `json:"age"`
	Gpa        float32          `json:"gpa"`
	Department Department       `json:"department"`
	Section    []KomplexSection `json:"sections"`
}
type KomplexCourse struct {
	ID         int        `json:"id"`
	Title      string     `json:"title"`
	Credit     int        `json:"credit"`
	Department Department `json:"department"`
}
type KomplexSection struct {
	Id         int           `json:"id"`
	Course     KomplexCourse `json:"course"`
	Number     int           `json:"number"`
	Instructor Instructor    `json:"instructor"`
	Student    []Student     `json:"students"`
}
type KomplexEnrollment struct {
	Id      int            `json:"id"`
	Student Student        `json:"student"`
	Section KomplexSection `json:"section"`
}

func saveStudent(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	student := Student{}
	json.Unmarshal(body, &student)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/University?charset=utf8")
	checkErr(err)
	defer db.Close()
	messageToClient := ""

	student.FirstName = strings.TrimSpace(student.FirstName)
	if len(student.FirstName) == 0 {
		messageToClient = messageToClient + "First name is empty. "
	}
	student.LastName = strings.TrimSpace(student.LastName)
	if len(student.LastName) == 0 {
		messageToClient = messageToClient + "Last name is empty. "
	}
	if student.Age < 1 {
		messageToClient = messageToClient + "Student age is wrong. "
	}
	if student.Gpa < 0 || student.Gpa > 4 {
		messageToClient = messageToClient + "Gpa is wrong. "
	}
	if student.DepartmentId == 0 {
		messageToClient = messageToClient + "Department is not selected. "
	}
	if messageToClient != "" {
		rw.Write([]byte(messageToClient))
		return
	}
	stmt, err := db.Prepare("INSERT Student SET firstName=?,lastName=?,age=?,gpa=?,department_Id=?")
	checkErr(err)
	_, err = stmt.Exec(student.FirstName, student.LastName, student.Age, student.Gpa, student.DepartmentId)
	checkErr(err)
}
func saveDepartment(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	dep := Department{}
	json.Unmarshal(body, &dep)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/University?charset=utf8")
	checkErr(err)
	defer db.Close()

	messageToClient := ""

	dep.Code = strings.TrimSpace(dep.Code)
	if len(dep.Code) == 0 {
		messageToClient = messageToClient + "Code is empty. "
	}
	dep.Name = strings.TrimSpace(dep.Name)
	if len(dep.Name) == 0 {
		messageToClient = messageToClient + "Name is empty. "
	}
	if messageToClient != "" {
		rw.Write([]byte(messageToClient))
		return
	}
	stmt, err := db.Prepare("INSERT Department SET code=?,name=?")
	checkErr(err)
	_, err = stmt.Exec(dep.Code, dep.Name)
	checkErr(err)
}
func saveCourse(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	course := Course{}
	json.Unmarshal(body, &course)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/University?charset=utf8")
	checkErr(err)
	defer db.Close()
	messageToClient := ""

	course.Title = strings.TrimSpace(course.Title)
	if len(course.Title) == 0 {
		messageToClient = messageToClient + "Title is empty. "
	}
	if course.Credit < 0 || course.Credit > 4 {
		messageToClient = messageToClient + "Credit is wrong. "
	}
	if course.DepartmentId == 0 {
		messageToClient = messageToClient + "Department is not selected. "
	}
	if messageToClient != "" {
		rw.Write([]byte(messageToClient))
		return
	}
	stmt, err := db.Prepare("INSERT Course SET title=?,credit=?,dep_id=?")
	checkErr(err)
	_, err = stmt.Exec(course.Title, course.Credit, course.DepartmentId)
	checkErr(err)
}
func saveInstructor(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	inst := Instructor{}
	json.Unmarshal(body, &inst)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/University?charset=utf8")
	checkErr(err)
	defer db.Close()
	messageToClient := ""

	inst.FirstName = strings.TrimSpace(inst.FirstName)
	if len(inst.FirstName) == 0 {
		messageToClient = messageToClient + "First name is empty. "
	}
	inst.LastName = strings.TrimSpace(inst.LastName)
	if len(inst.LastName) == 0 {
		messageToClient = messageToClient + "Last name is empty. "
	}
	if inst.Age < 0 || inst.Age > 130 {
		messageToClient = messageToClient + "Age is wrong. "
	}
	if messageToClient != "" {
		rw.Write([]byte(messageToClient))
		return
	}
	stmt, err := db.Prepare("INSERT Instructor SET firstName=?,lastName=?,age=?")
	checkErr(err)
	_, err = stmt.Exec(inst.FirstName, inst.LastName, inst.Age)
	checkErr(err)
}
func saveSection(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	section := Section{}
	json.Unmarshal(body, &section)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/University?charset=utf8")
	checkErr(err)
	defer db.Close()
	messageToClient := ""

	if section.CourseId == 0 {
		messageToClient = messageToClient + "Course is not selected. "
	}
	if section.Number < 1 {
		messageToClient = messageToClient + "Section number is wrong. "
	}
	if section.InstructorId == 0 {
		messageToClient = messageToClient + "Instructor is not selected. "
	}
	if messageToClient != "" {
		rw.Write([]byte(messageToClient))
		return
	}
	stmt, err := db.Prepare("INSERT Section SET course_id=?,number=?,instructor_id=?")
	checkErr(err)
	_, err = stmt.Exec(section.CourseId, section.Number, section.InstructorId)
	checkErr(err)
}
func saveEnrollment(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	enroll := Enrollment{}
	json.Unmarshal(body, &enroll)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/University?charset=utf8")
	checkErr(err)
	defer db.Close()
	messageToClient := ""

	if enroll.SectionId == 0 {
		messageToClient = messageToClient + "Section is not selected. "
	}
	if enroll.StudentId == 0 {
		messageToClient = messageToClient + "Student is not selected. "
	}
	if messageToClient != "" {
		rw.Write([]byte(messageToClient))
		return
	}
	stmt, err := db.Prepare("INSERT Enrollment SET student_id=?,section_id=?")
	checkErr(err)
	_, err = stmt.Exec(enroll.StudentId, enroll.SectionId)
	checkErr(err)
}

func getAllStudents(rw http.ResponseWriter, req *http.Request) {
	allStudents := make([]KomplexStudent, 0)
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/University?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Student")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		komplexStudent := KomplexStudent{}
		student := Student{}
		err := rows.Scan(&student.Id, &student.FirstName, &student.LastName, &student.Age, &student.Gpa, &student.DepartmentId)
		checkErr(err)

		rows2, err2 := db.Query("SELECT * FROM Department WHERE id=?", student.DepartmentId)
		checkErr(err2)
		defer rows2.Close()
		department := Department{}
		for rows2.Next() {
			err := rows2.Scan(&department.ID, &department.Code, &department.Name)
			checkErr(err)
		}
		komplexStudent.Id = student.Id
		komplexStudent.FirstName = student.FirstName
		komplexStudent.LastName = student.LastName
		komplexStudent.Age = student.Age
		komplexStudent.Gpa = student.Gpa
		komplexStudent.Department = department

		rows3, err3 := db.Query("SELECT * FROM Enrollment WHERE student_id=?", student.Id)
		checkErr(err3)
		allEnroll := make([]Enrollment, 0)
		for rows3.Next() {
			enroll := Enrollment{}
			err := rows3.Scan(&enroll.Id, &enroll.StudentId, &enroll.SectionId)
			checkErr(err)
			allEnroll = append(allEnroll, enroll)
		}

		allSections := make([]KomplexSection, 0)
		for i := 0; i < len(allEnroll); i++ {
			rows4, err4 := db.Query("SELECT * FROM Section WHERE id=?", allEnroll[i].SectionId)
			checkErr(err4)

			for rows4.Next() {
				komplexSection := KomplexSection{}
				section := Section{}
				err := rows4.Scan(&section.Id, &section.CourseId, &section.Number, &section.InstructorId)
				checkErr(err)

				rows2, err2 := db.Query("SELECT * FROM Course WHERE id=?", section.CourseId)
				checkErr(err2)
				komplexCourse := KomplexCourse{}
				course := Course{}
				for rows2.Next() {
					err := rows2.Scan(&course.ID, &course.Title, &course.Credit, &course.DepartmentId)
					checkErr(err)

					rows3, err3 := db.Query("SELECT * FROM Department WHERE id=?", course.DepartmentId)
					checkErr(err3)
					department := Department{}
					for rows3.Next() {
						err := rows3.Scan(&department.ID, &department.Code, &department.Name)
						checkErr(err)
					}
					komplexCourse.ID = course.ID
					komplexCourse.Title = course.Title
					komplexCourse.Credit = course.Credit
					komplexCourse.Department = department
				}

				rows4, err4 := db.Query("SELECT * FROM Instructor WHERE id=?", section.InstructorId)
				checkErr(err4)
				instructor := Instructor{}
				for rows4.Next() {
					err := rows4.Scan(&instructor.Id, &instructor.FirstName, &instructor.LastName, &instructor.Age)
					checkErr(err)
				}

				rows5, err5 := db.Query("SELECT * FROM Enrollment WHERE section_id=?", section.Id)
				checkErr(err5)
				allEnroll := make([]Enrollment, 0)
				for rows5.Next() {
					enroll := Enrollment{}
					err := rows5.Scan(&enroll.Id, &enroll.StudentId, &enroll.SectionId)
					checkErr(err)
					allEnroll = append(allEnroll, enroll)
				}

				allStudents := make([]Student, 0)
				for i := 0; i < len(allEnroll); i++ {
					rows6, err6 := db.Query("SELECT * FROM Student WHERE id=?", allEnroll[i].StudentId)
					checkErr(err6)
					for rows6.Next() {
						student := Student{}
						err := rows6.Scan(&student.Id, &student.FirstName, &student.LastName, &student.Age, &student.Gpa, &student.DepartmentId)
						checkErr(err)
						allStudents = append(allStudents, student)
					}
				}

				komplexSection.Course = komplexCourse
				komplexSection.Id = section.Id
				komplexSection.Number = section.Number
				komplexSection.Instructor = instructor
				komplexSection.Student = allStudents

				allSections = append(allSections, komplexSection)
			}

		}
		komplexStudent.Section = allSections

		allStudents = append(allStudents, komplexStudent)
	}
	json.NewEncoder(rw).Encode(allStudents)
}
func getAllDepartments(rw http.ResponseWriter, req *http.Request) {

	allDepartments := make([]Department, 0)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/University?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Department")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		depart := Department{}
		err := rows.Scan(&depart.ID, &depart.Code, &depart.Name)
		checkErr(err)

		allDepartments = append(allDepartments, depart)
	}
	json.NewEncoder(rw).Encode(allDepartments)
}
func getAllCourses(rw http.ResponseWriter, req *http.Request) {

	allCourses := make([]KomplexCourse, 0)
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/University?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Course")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		komplexCourse := KomplexCourse{}
		course := Course{}
		err := rows.Scan(&course.ID, &course.Title, &course.Credit, &course.DepartmentId)
		checkErr(err)

		rows2, err2 := db.Query("SELECT * FROM Department WHERE id=?", course.DepartmentId)
		checkErr(err2)
		defer rows2.Close()

		department := Department{}
		for rows2.Next() {
			err := rows2.Scan(&department.ID, &department.Code, &department.Name)
			checkErr(err)
		}

		komplexCourse.Department = department
		komplexCourse.Credit = course.Credit
		komplexCourse.Title = course.Title
		komplexCourse.ID = course.ID

		allCourses = append(allCourses, komplexCourse)
	}
	json.NewEncoder(rw).Encode(allCourses)
}

func getAllInstructors(rw http.ResponseWriter, req *http.Request) {

	allInstructors := make([]Instructor, 0)
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/University?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Instructor")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		instructor := Instructor{}
		err := rows.Scan(&instructor.Id, &instructor.FirstName, &instructor.LastName, &instructor.Age)
		checkErr(err)

		allInstructors = append(allInstructors, instructor)
	}
	json.NewEncoder(rw).Encode(allInstructors)
}
func getAllSections(rw http.ResponseWriter, req *http.Request) {
	allSections := make([]KomplexSection, 0)
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/University?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Section")
	checkErr(err)

	for rows.Next() {
		komplexSection := KomplexSection{}
		section := Section{}
		err := rows.Scan(&section.Id, &section.CourseId, &section.Number, &section.InstructorId)
		checkErr(err)

		rows2, err2 := db.Query("SELECT * FROM Course WHERE id=?", section.CourseId)
		checkErr(err2)
		komplexCourse := KomplexCourse{}
		course := Course{}
		for rows2.Next() {
			err := rows2.Scan(&course.ID, &course.Title, &course.Credit, &course.DepartmentId)
			checkErr(err)

			rows3, err3 := db.Query("SELECT * FROM Department WHERE id=?", course.DepartmentId)
			checkErr(err3)
			department := Department{}
			for rows3.Next() {
				err := rows3.Scan(&department.ID, &department.Code, &department.Name)
				checkErr(err)
			}
			komplexCourse.ID = course.ID
			komplexCourse.Title = course.Title
			komplexCourse.Credit = course.Credit
			komplexCourse.Department = department
		}

		rows4, err4 := db.Query("SELECT * FROM Instructor WHERE id=?", section.InstructorId)
		checkErr(err4)
		instructor := Instructor{}
		for rows4.Next() {
			err := rows4.Scan(&instructor.Id, &instructor.FirstName, &instructor.LastName, &instructor.Age)
			checkErr(err)
		}

		rows5, err5 := db.Query("SELECT * FROM Enrollment WHERE section_id=?", section.Id)
		checkErr(err5)
		allEnroll := make([]Enrollment, 0)
		for rows5.Next() {
			enroll := Enrollment{}
			err := rows5.Scan(&enroll.Id, &enroll.StudentId, &enroll.SectionId)
			checkErr(err)
			allEnroll = append(allEnroll, enroll)
		}

		allStudents := make([]Student, 0)
		for i := 0; i < len(allEnroll); i++ {
			rows6, err6 := db.Query("SELECT * FROM Student WHERE id=?", allEnroll[i].StudentId)
			checkErr(err6)
			for rows6.Next() {
				student := Student{}
				err := rows6.Scan(&student.Id, &student.FirstName, &student.LastName, &student.Age, &student.Gpa, &student.DepartmentId)
				checkErr(err)
				allStudents = append(allStudents, student)
			}
		}

		komplexSection.Course = komplexCourse
		komplexSection.Id = section.Id
		komplexSection.Number = section.Number
		komplexSection.Instructor = instructor
		komplexSection.Student = allStudents

		allSections = append(allSections, komplexSection)
	}
	json.NewEncoder(rw).Encode(allSections)
}
func getAllEnrollments(rw http.ResponseWriter, req *http.Request) {

	allEnroll := make([]KomplexEnrollment, 0)
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/University?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Enrollment")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		komplexEnroll := KomplexEnrollment{}

		enroll := Enrollment{}
		err := rows.Scan(&enroll.Id, &enroll.StudentId, &enroll.SectionId)
		checkErr(err)

		rows2, err2 := db.Query("SELECT * FROM Student WHERE id=?", enroll.StudentId)
		checkErr(err2)
		student := Student{}
		for rows2.Next() {
			err := rows2.Scan(&student.Id, &student.FirstName, &student.LastName, &student.Age, &student.Gpa, &student.DepartmentId)
			checkErr(err)
		}
		komplexEnroll.Student = student

		rows3, err3 := db.Query("SELECT * FROM Section WHERE id=?", enroll.SectionId)
		checkErr(err3)
		komplexSection := KomplexSection{}

		for rows3.Next() {
			section := Section{}
			err := rows3.Scan(&section.Id, &section.CourseId, &section.Number, &section.InstructorId)
			checkErr(err)

			rows4, err4 := db.Query("SELECT * FROM Course WHERE id=?", section.CourseId)
			checkErr(err4)
			komplexCourse := KomplexCourse{}
			course := Course{}
			for rows4.Next() {
				err := rows4.Scan(&course.ID, &course.Title, &course.Credit, &course.DepartmentId)
				checkErr(err)

				rows5, err5 := db.Query("SELECT * FROM Department WHERE id=?", course.DepartmentId)
				checkErr(err5)
				department := Department{}
				for rows5.Next() {
					err := rows5.Scan(&department.ID, &department.Code, &department.Name)
					checkErr(err)
				}
				komplexCourse.ID = course.ID
				komplexCourse.Title = course.Title
				komplexCourse.Credit = course.Credit
				komplexCourse.Department = department
			}

			rows6, err6 := db.Query("SELECT * FROM Instructor WHERE id=?", section.InstructorId)
			checkErr(err6)
			instructor := Instructor{}
			for rows6.Next() {
				err := rows6.Scan(&instructor.Id, &instructor.FirstName, &instructor.LastName, &instructor.Age)
				checkErr(err)
			}

			rows7, err7 := db.Query("SELECT * FROM Enrollment WHERE section_id=?", section.Id)
			checkErr(err7)
			allEnroll := make([]Enrollment, 0)
			for rows7.Next() {
				enroll := Enrollment{}
				err := rows7.Scan(&enroll.Id, &enroll.StudentId, &enroll.SectionId)
				checkErr(err)
				allEnroll = append(allEnroll, enroll)
			}

			allStudents := make([]Student, 0)
			for i := 0; i < len(allEnroll); i++ {
				rows8, err8 := db.Query("SELECT * FROM Student WHERE id=?", allEnroll[i].StudentId)
				checkErr(err8)
				for rows8.Next() {
					student := Student{}
					err := rows8.Scan(&student.Id, &student.FirstName, &student.LastName, &student.Age, &student.Gpa, &student.DepartmentId)
					checkErr(err)
					allStudents = append(allStudents, student)
				}
			}

			komplexSection.Course = komplexCourse
			komplexSection.Id = section.Id
			komplexSection.Number = section.Number
			komplexSection.Instructor = instructor
			komplexSection.Student = allStudents
		}

		komplexEnroll.Section = komplexSection

		allEnroll = append(allEnroll, komplexEnroll)
	}
	json.NewEncoder(rw).Encode(allEnroll)
}

func deleteStudent(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	student := Student{}
	json.Unmarshal(body, &student)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/University?charset=utf8")
	checkErr(err)
	defer db.Close()

	_, err = db.Query("DELETE FROM Student WHERE id=?", student.Id)
	checkErr(err)
}
func deleteDepartment(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	dep := Department{}
	json.Unmarshal(body, &dep)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/University?charset=utf8")
	checkErr(err)
	defer db.Close()

	_, err = db.Query("DELETE FROM Department WHERE id=?", dep.ID)
	checkErr(err)
}
func deleteCourse(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	course := Course{}
	json.Unmarshal(body, &course)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/University?charset=utf8")
	checkErr(err)
	defer db.Close()

	_, err = db.Query("DELETE FROM Course WHERE id=?", course.ID)
	checkErr(err)
}
func deleteInstructor(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	inst := Instructor{}
	json.Unmarshal(body, &inst)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/University?charset=utf8")
	checkErr(err)
	defer db.Close()

	_, err = db.Query("DELETE FROM Instructor WHERE id=?", inst.Id)
	checkErr(err)
}
func deleteSection(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	section := Section{}
	json.Unmarshal(body, &section)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/University?charset=utf8")
	checkErr(err)
	defer db.Close()

	_, err = db.Query("DELETE FROM Section WHERE id=?", section.Id)
	checkErr(err)
}
func deleteEnrollment(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	enroll := Enrollment{}
	json.Unmarshal(body, &enroll)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/University?charset=utf8")
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM Enrollment WHERE student_id =? AND section_id=?")
	checkErr(err)
	_, err = stmt.Exec(&enroll.Id, enroll.StudentId, enroll.SectionId)
	checkErr(err)
}

func main() {

	http.HandleFunc("/js/dirPagination.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/js/dirPagination.js")
	})
	http.HandleFunc("/dirPagination.tpl.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/dirPagination.tpl.html")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/new.html")
	})

	http.HandleFunc("/js/new.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/js/new.js")
	})

	http.HandleFunc("/js/angular1.6.4.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/js/angular1.6.4.js")
	})
	http.HandleFunc("/js/angular-route1.6.4.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/js/angular-route1.6.4.js")
	})

	http.HandleFunc("/js/jquery3.3.1.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/js/jquery3.3.1.js")
	})

	http.HandleFunc("/js/bootstrap3.3.7.min.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/js/bootstrap3.3.7.min.js")
	})

	http.HandleFunc("/js/fontawesome5.0.6.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/js/fontawesome5.0.6.js")
	})

	http.HandleFunc("/css/bootstrap3.3.7.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/css/bootstrap3.3.7.css")
	})
	http.HandleFunc("/pages/page1.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/pages/page1.html")
	})
	http.HandleFunc("/pages/page2.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/pages/page2.html")
	})
	http.HandleFunc("/pages/page3.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/pages/page3.html")
	})
	http.HandleFunc("/pages/page4.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/pages/page4.html")
	})
	http.HandleFunc("/pages/page5.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/pages/page5.html")
	})
	http.HandleFunc("/pages/page6.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/pages/page6.html")
	})
	http.HandleFunc("/css/design.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/css/design.css")
	})
	http.HandleFunc("/css/index.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/css/index.css")
	})
	http.HandleFunc("/saveStudent", saveStudent)

	http.HandleFunc("/saveDepartment", saveDepartment)

	http.HandleFunc("/saveCourse", saveCourse)

	http.HandleFunc("/saveInstructor", saveInstructor)

	http.HandleFunc("/saveSection", saveSection)

	http.HandleFunc("/saveEnrollment", saveEnrollment)

	http.HandleFunc("/getAllStudents", getAllStudents)

	http.HandleFunc("/getAllDepartments", getAllDepartments)

	http.HandleFunc("/getAllCourses", getAllCourses)

	http.HandleFunc("/getAllInstructors", getAllInstructors)

	http.HandleFunc("/getAllSections", getAllSections)

	http.HandleFunc("/getAllEnrollments", getAllEnrollments)

	http.HandleFunc("/deleteStudent", deleteStudent)

	http.HandleFunc("/deleteDepartment", deleteDepartment)

	http.HandleFunc("/deleteCourse", deleteCourse)

	http.HandleFunc("/deleteInstructor", deleteInstructor)

	http.HandleFunc("/deleteSection", deleteSection)

	http.HandleFunc("/deleteEnrollment", deleteEnrollment)

	log.Fatal(http.ListenAndServe(":1112", nil))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
