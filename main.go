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
	ID        int    `json:"id"`
	Name      string `json:"name"`
	TeacherID int    `json:"teacherId"`
}

// Teacher for get response(saveTeacherAndClass)
type Teacher struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	ClassID int    `json:"classId"`
	UniID   int    `json:"uniId"`
}

// KomplexTeacher for getTeacherData
type KomplexTeacher struct {
	ID      int        `json:"id"`
	Name    string     `json:"name"`
	ClassID int        `json:"classId"`
	Uni     University `json:"uni"`
}

// StudentAndClass for get response(saveTeacherAndClass)
type StudentAndClass struct {
	SID int `json:"s_id"`
	CID int `json:"c_id"`
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

	stmt, err := db.Prepare("INSERT University SET University_Name=?,University_Capacity=?,University_RecordedStudent=?")
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

	stmt, err := db.Prepare("INSERT Class SET Name=?")
	checkErr(err)

	messageToClient := ""

	class.Name = strings.TrimSpace(class.Name)
	if len(class.Name) == 0 {
		messageToClient = messageToClient + "Name is empty. "
		rw.Write([]byte(messageToClient))
		return
	}

	_, err = stmt.Exec(class.Name)
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
	}
	if teacher.ClassID < 1 {
		messageToClient = messageToClient + "Class is not selected. "
	}
	if teacher.UniID < 1 {
		messageToClient = messageToClient + "University is not selected."
	}
	if messageToClient != "" {
		rw.Write([]byte(messageToClient))
		return
	}

	stmt, err := db.Query("INSERT Teacher SET Name=?,ClassId=?,UniId=?", teacher.Name, teacher.ClassID, teacher.UniID)
	checkErr(err)
	defer stmt.Close()
}

//insert teacher-class to the database
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
	teachers := make([]KomplexTeacher, 0)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Teacher WHERE Name LIKE ?", "%"+requestedName+"%")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		komplexTeacher := KomplexTeacher{}
		teacher := Teacher{}
		err := rows.Scan(&teacher.ID, &teacher.Name, &teacher.ClassID, &teacher.UniID)
		checkErr(err)

		university := University{}
		stm, err := db.Query("SELECT * FROM University WHERE University_ID= ?", teacher.UniID)
		checkErr(err)
		defer stm.Close()

		for stm.Next() {
			err := stm.Scan(&university.ID, &university.Name, &university.Capacity, &university.Number)
			checkErr(err)
		}

		komplexTeacher.ID = teacher.ID
		komplexTeacher.Name = teacher.Name
		komplexTeacher.ClassID = teacher.ClassID
		komplexTeacher.Uni = university

		teachers = append(teachers, komplexTeacher)
	}
	json.NewEncoder(rw).Encode(teachers)
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

	allStudents := make([]Student, 0)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Student")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		stu := Student{}
		err := rows.Scan(&stu.ID, &stu.Name, &stu.Age, &stu.UniversityID)
		checkErr(err)

		allStudents = append(allStudents, stu)
	}
	json.NewEncoder(rw).Encode(allStudents)
}

//return all students in the database
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
		err := rows.Scan(&class.ID, &class.Name, &class.TeacherID)
		checkErr(err)

		allClasses = append(allClasses, class)
	}
	json.NewEncoder(rw).Encode(allClasses)
}

func getNotSelectedClasses(rw http.ResponseWriter, req *http.Request) {

	allClasses := make([]Class, 0)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/StudentSystem?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Class")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		class := Class{}
		err := rows.Scan(&class.ID, &class.Name, &class.TeacherID)
		checkErr(err)

		if class.TeacherID == 0 {
			allClasses = append(allClasses, class)
		}

	}
	json.NewEncoder(rw).Encode(allClasses)
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

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/homePage.html")
	})
	http.HandleFunc("/findPage", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/findPage.html")
	})
	http.HandleFunc("/js/homePagesjs.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/js/homePagesjs.js")
	})

	http.HandleFunc("/js/findPagejs.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/js/findPagejs.js")
	})
	http.HandleFunc("/css/bootstrap.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/css/bootstrap.css")
	})
	http.HandleFunc("/css/design.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/css/design.css")
	})

	http.HandleFunc("/saveStudent", saveStudent)

	http.HandleFunc("/saveUniversity", saveUniversity)

	http.HandleFunc("/saveClass", saveClass)

	http.HandleFunc("/saveTeacher", saveTeacher)

	http.HandleFunc("/saveStudentAndClass", saveStudentAndClass)

	http.HandleFunc("/getStudentData", getStudentData)

	http.HandleFunc("/getUniversityData", getUniversityData)

	http.HandleFunc("/getTeacherData", getTeacherData)

	http.HandleFunc("/getAllUniversities", getAllUniversities)

	http.HandleFunc("/getAllStudents", getAllStudents)

	http.HandleFunc("/getAllClasses", getAllClasses)

	http.HandleFunc("/getNotSelectedClasses", getNotSelectedClasses)

	http.HandleFunc("/deleteStudent", deleteStudent)

	http.HandleFunc("/deleteUniversity", deleteUniversity)

	log.Fatal(http.ListenAndServe(":1112", nil))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}