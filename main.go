package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Student for get response(saveStudent)
type Student struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	UniversityID int    `json:"uni"`
}

// KomplexStudent for getStudentData
type KomplexStudent struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Age        int        `json:"age"`
	University University `json:"uni"`
}

// University for get response(saveStudent)
type University struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Number   int    `json:"number"`
}

// Class for get response(saveTeacherAndClass)
type Class struct {
	ID         int    `json:"id"`
	Department string `json:"dep"`
	Code       int    `json:"code"`
	Name       string `json:"name"`
}

// Teacher for get response(saveTeacherAndClass)
type Teacher struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// StudentAndClass for get response(saveTeacherAndClass)
type StudentAndClass struct {
	SID int `json:"s_id"`
	CID int `json:"c_id"`
}

type TeacherAndClass struct {
	TID int `json:"t_id"`
	CID int `json:"c_id"`
}

type Department struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//insert student to the database
func saveStudent(rw http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	student := Student{}
	json.Unmarshal(body, &student)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("SELECT * FROM University WHERE University_ID= ?", student.UniversityID)
	checkErr(err)
	defer rows.Close()

	university := University{}
	for rows.Next() {
		err := rows.Scan(&university.ID, &university.Name, &university.Capacity, &university.Number)
		checkErr(err)
	}

	student.Name = strings.TrimSpace(student.Name)
	messageToClient := ""
	if len(student.Name) == 0 {
		messageToClient = messageToClient + "Student name is empty. "
	}
	if student.Age < 1 {
		messageToClient = messageToClient + "Student age is wrong. "
	}
	if student.UniversityID == 0 {
		messageToClient = messageToClient + "University is not selected. "

	} else if university.Number >= university.Capacity {
		messageToClient = messageToClient + "Insufficient universite capacity. "
	}
	if messageToClient != "" {
		rw.Write([]byte(messageToClient))
		return
	}

	stmt, err := db.Prepare("INSERT Student SET Student_Name=?,Student_Age=?,Student_UniversityID=?")
	checkErr(err)
	_, err = stmt.Exec(student.Name, student.Age, student.UniversityID)
	checkErr(err)
}

//insert university to the database
func saveUniversity(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	uni := University{}
	json.Unmarshal(body, &uni)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	stmt, err :=
		db.Prepare("INSERT University SET University_Name=?,University_Capacity=?,University_RecordedStudent=?")
	checkErr(err)

	messageToClient := ""

	uni.Name = strings.TrimSpace(uni.Name)
	if len(uni.Name) == 0 {
		messageToClient = messageToClient + "Name is empty. "
	}
	if uni.Capacity < 1 {
		messageToClient = messageToClient + "Capacity is wrong."
	}
	if messageToClient != "" {
		rw.Write([]byte(messageToClient))
		return
	}
	_, err = stmt.Exec(uni.Name, uni.Capacity, 0)
	checkErr(err)
}

func saveClass(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	class := Class{}
	json.Unmarshal(body, &class)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("INSERT Class SET Department=?, Code=?,Name=?")
	checkErr(err)

	messageToClient := ""

	class.Department = strings.TrimSpace(class.Department)
	if len(class.Department) == 0 {
		messageToClient = messageToClient + " Department is empty. "
	}

	if class.Code < 1 {
		messageToClient = messageToClient + " Code is empty. "
	}

	class.Name = strings.TrimSpace(class.Name)
	if len(class.Name) == 0 {
		messageToClient = messageToClient + "Name is empty. "
		rw.Write([]byte(messageToClient))
		return
	}

	_, err = stmt.Exec(class.Department, class.Code, class.Name)
	checkErr(err)
}

//insert teacher-class to the database
func saveTeacher(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	teacher := Teacher{}
	json.Unmarshal(body, &teacher)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	messageToClient := ""

	teacher.Name = strings.TrimSpace(teacher.Name)
	if len(teacher.Name) == 0 {
		messageToClient = messageToClient + "Name is empty. "
		rw.Write([]byte(messageToClient))
		return
	}

	stmt, err := db.Query("INSERT Teacher SET Name=?", teacher.Name)
	checkErr(err)
	defer stmt.Close()
}

//insert student-class to the database
func saveStudentAndClass(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	sc := StudentAndClass{}
	json.Unmarshal(body, &sc)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	messageToClient := ""

	if sc.SID < 1 {
		messageToClient = messageToClient + "Student is not selected."
	}
	if sc.CID < 1 {
		messageToClient = messageToClient + "Class is not selected."
	}
	if messageToClient != "" {
		rw.Write([]byte(messageToClient))
		return
	}

	rows, err := db.Query("INSERT StudentClass SET s_id=?, c_id=?", sc.SID, sc.CID)
	checkErr(err)
	defer rows.Close()
}

//insert teacher-class to the database
func saveTeacherAndClass(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	tc := TeacherAndClass{}
	json.Unmarshal(body, &tc)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	messageToClient := ""

	if tc.TID < 1 {
		messageToClient = messageToClient + "teacher is not selected."
	}
	if tc.CID < 1 {
		messageToClient = messageToClient + "Class is not selected."
	}
	if messageToClient != "" {
		rw.Write([]byte(messageToClient))
		return
	}

	rows, err := db.Query("INSERT TeacherClass SET t_id=?, c_id=?", tc.TID, tc.CID)
	checkErr(err)
	defer rows.Close()
}

//return specific students with contains response writer's Name
func getStudentData(rw http.ResponseWriter, req *http.Request) {
	requestedName := req.Header.Get("Name")
	students := make([]KomplexStudent, 0)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	requestedName = strings.TrimSpace(requestedName)
	if len(requestedName) == 0 {
		return
	}

	komplexStudent := KomplexStudent{}
	student := Student{}
	rows, err := db.Query("SELECT * FROM Student WHERE Student_Name LIKE ?", "%"+requestedName+"%")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.UniversityID)
		checkErr(err)

		university := University{}
		rows, err := db.Query("SELECT * FROM University WHERE University_ID= ?", student.UniversityID)
		checkErr(err)
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&university.ID, &university.Name, &university.Capacity, &university.Number)
			checkErr(err)
		}
		komplexStudent.Age = student.Age
		komplexStudent.Name = student.Name
		komplexStudent.ID = student.ID
		komplexStudent.University = university

		students = append(students, komplexStudent)
	}
	json.NewEncoder(rw).Encode(students)
}

//return specific universities with contains response writer's Name
func getUniversityData(rw http.ResponseWriter, req *http.Request) {

	requestedName := req.Header.Get("Name")
	requestedName = strings.TrimSpace(requestedName)
	if len(requestedName) == 0 {
		return
	}
	universities := make([]University, 0)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM University WHERE University_Name LIKE ?", "%"+requestedName+"%")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		university := University{}
		err := rows.Scan(&university.ID, &university.Name, &university.Capacity, &university.Number)
		checkErr(err)
		universities = append(universities, university)
	}
	json.NewEncoder(rw).Encode(universities)
}

func getTeacherData(rw http.ResponseWriter, req *http.Request) {

	requestedName := req.Header.Get("Name")
	requestedName = strings.TrimSpace(requestedName)

	if len(requestedName) == 0 {
		return
	}
	teachers := make([]Teacher, 0)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Teacher WHERE Name LIKE ?", "%"+requestedName+"%")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		teacher := Teacher{}
		err := rows.Scan(&teacher.ID, &teacher.Name)
		checkErr(err)

		teachers = append(teachers, teacher)
	}
	json.NewEncoder(rw).Encode(teachers)
}

func getClassData(rw http.ResponseWriter, req *http.Request) {

	requestedName := req.Header.Get("Name")
	requestedName = strings.TrimSpace(requestedName)

	if len(requestedName) == 0 {
		return
	}
	classes := make([]Class, 0)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Class WHERE Name LIKE ?", "%"+requestedName+"%")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		class := Class{}
		err := rows.Scan(&class.ID, &class.Name)
		checkErr(err)

		classes = append(classes, class)
	}
	json.NewEncoder(rw).Encode(classes)
}

//return all universities in the database
func getAllUniversities(rw http.ResponseWriter, req *http.Request) {

	allUniversities := make([]University, 0)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM University")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		uni := University{}
		err := rows.Scan(&uni.ID, &uni.Name, &uni.Capacity, &uni.Number)
		checkErr(err)

		allUniversities = append(allUniversities, uni)
	}
	json.NewEncoder(rw).Encode(allUniversities)
}

//return all students in the database
func getAllStudents(rw http.ResponseWriter, req *http.Request) {

	allStudents := make([]KomplexStudent, 0)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	komplexStudent := KomplexStudent{}
	student := Student{}
	rows, err := db.Query("SELECT * FROM Student")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.UniversityID)
		checkErr(err)

		university := University{}
		rows, err := db.Query("SELECT * FROM University WHERE University_ID= ?", student.UniversityID)
		checkErr(err)
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&university.ID, &university.Name, &university.Capacity, &university.Number)
			checkErr(err)
		}
		komplexStudent.Age = student.Age
		komplexStudent.Name = student.Name
		komplexStudent.ID = student.ID
		komplexStudent.University = university

		allStudents = append(allStudents, komplexStudent)
	}
	json.NewEncoder(rw).Encode(allStudents)
}

//return all classes in the database
func getAllClasses(rw http.ResponseWriter, req *http.Request) {

	allClasses := make([]Class, 0)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Class")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		class := Class{}
		err := rows.Scan(&class.ID, &class.Department, &class.Code, &class.Name)
		checkErr(err)

		allClasses = append(allClasses, class)
	}
	json.NewEncoder(rw).Encode(allClasses)
}

//return all teachers in the database
func getAllTeachers(rw http.ResponseWriter, req *http.Request) {

	allTeachers := make([]Teacher, 0)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Teacher")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		teacher := Teacher{}
		err := rows.Scan(&teacher.ID, &teacher.Name)
		checkErr(err)

		allTeachers = append(allTeachers, teacher)
	}
	json.NewEncoder(rw).Encode(allTeachers)
}

//return all departments in the database
func getAllDepartments(rw http.ResponseWriter, req *http.Request) {

	allDepartments := make([]Department, 0)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Department")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		depart := Department{}
		err := rows.Scan(&depart.ID, &depart.Name)
		checkErr(err)

		allDepartments = append(allDepartments, depart)
	}
	json.NewEncoder(rw).Encode(allDepartments)
}

func deleteStudent(rw http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	student := Student{}
	json.Unmarshal(body, &student)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	_, err = db.Query("DELETE FROM Student WHERE Student_ID=?", student.ID)
	checkErr(err)
}

func deleteUniversity(rw http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	university := University{}
	json.Unmarshal(body, &university)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	_, err = db.Query("DELETE FROM University WHERE University_ID=?", university.ID)
	checkErr(err)
}

func deleteTeacher(rw http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	teacher := Teacher{}
	json.Unmarshal(body, &teacher)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	_, err = db.Query("DELETE FROM Teacher WHERE Id=?", teacher.ID)
	checkErr(err)
}

func deleteClass(rw http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	class := Class{}
	json.Unmarshal(body, &class)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	_, err = db.Query("DELETE FROM Class WHERE Id=?", class.ID)
	checkErr(err)
}

func main() {

	http.HandleFunc("/dirPagination.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/dirPagination.js")
	})

	http.HandleFunc("/dirPagination.tpl.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/dirPagination.tpl.html")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/findPage", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/findPage.html")
	})

	http.HandleFunc("/homePage", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/homePage.html")
	})

	http.HandleFunc("/js/homePagesjs.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/js/homePagesjs.js")
	})

	http.HandleFunc("/js/findPagejs.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/js/findPagejs.js")
	})

	http.HandleFunc("/js/jquery.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/js/jquery.js")
	})

	http.HandleFunc("/js/bootstrap.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/js/bootstrap.js")
	})

	http.HandleFunc("/js/all.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/js/all.js")
	})

	http.HandleFunc("/js/index.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/js/index.js")
	})

	http.HandleFunc("/css/bootstrap.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/css/bootstrap.css")
	})

	http.HandleFunc("/css/design.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/css/design.css")
	})

	http.HandleFunc("/css/index.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/css/index.css")
	})

	http.HandleFunc("/saveUniversity", saveUniversity)

	http.HandleFunc("/saveStudent", saveStudent)

	http.HandleFunc("/saveClass", saveClass)

	http.HandleFunc("/saveTeacher", saveTeacher)

	http.HandleFunc("/saveStudentAndClass", saveStudentAndClass)

	http.HandleFunc("/saveTeacherAndClass", saveTeacherAndClass)

	http.HandleFunc("/getUniversityData", getUniversityData)

	http.HandleFunc("/getStudentData", getStudentData)

	http.HandleFunc("/getTeacherData", getTeacherData)

	http.HandleFunc("/getClassData", getClassData)

	http.HandleFunc("/getAllUniversities", getAllUniversities)

	http.HandleFunc("/getAllStudents", getAllStudents)

	http.HandleFunc("/getAllClasses", getAllClasses)

	http.HandleFunc("/getAllTeachers", getAllTeachers)

	http.HandleFunc("/getAllDepartments", getAllDepartments)

	http.HandleFunc("/deleteStudent", deleteStudent)

	http.HandleFunc("/deleteUniversity", deleteUniversity)

	http.HandleFunc("/deleteTeacher", deleteTeacher)

	http.HandleFunc("/deleteClass", deleteClass)

	log.Fatal(http.ListenAndServe(":1112", nil))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
