var app = angular.module("myApp",['angularUtils.directives.dirPagination']);

app.controller("myCtrl", function($scope, $http) {

    $scope.messageOfSaveDepartment = "";
    $scope.messageOfSaveStudent = "";
    $scope.messageOfSaveCourse="";
    $scope.messageOfSaveInstructor="";
    $scope.messageOfSaveSection="";
    $scope.messageOfSaveEnrollment="";

    $scope.student = {"id":"","firstName":"", "lastName": "","age":"","gpa":"" };
    $scope.department ={"id":"","code":"","name":""};
    $scope.course = {"id":"","title":"","credit":""};
    $scope.instructor = {"id":"","firstName":"", "lastName": "" ,"age":""};
    $scope.section = {"id":"","courseId":"", "number": "" ,"instructorId":""};
    $scope.enrollment ={"studentId":"","sectionId":""};

    $scope.allStudents ="";
    $scope.allDepartments="";
    $scope.allCourses="";
    $scope.allInstructors ="";
    $scope.allSections ="";
    $scope.allEnrollments ="";

    $scope.currentPageDpt = 1;
    $scope.currentPageStd = 1;
    $scope.currentPageCourse=1;
    $scope.currentPageInstructor=1;
    $scope.currentPageSection=1;
    $scope.currentPageEnrollment=1;

    $scope.pageSize = 10;

    $scope.selectedDepartmentOfStudent ={"id":"","code":"","name":""};
    $scope.selectedDepartmentOfCourse = {"id":"","code":"","name":""};

    $scope.selectedCourseOfSection= {"id":"","title":"","credit":""};
    $scope.selectedInstructorOfSection = {"id":"","firstName":"", "lastName": "" ,"age":""};

    $scope.selectedStudentOfEnrollment = {"id":"","firstName":"", "lastName": "","age":"","gpa":"" };
    $scope.selectedSectionOfEnrollment = {"id":"","courseId":"", "number": "" ,"instructorId":""};



    $scope.sendStudentData = function () {
        var url = "/saveStudent", data = {"firstName": $scope.student.firstName,"lastName": $scope.student.lastName,
            "age": $scope.student.age,"gpa": $scope.student.gpa,"departmentId": $scope.selectedDepartmentOfStudent.id};

        $http.post(url, data).then(function mySuccess(response) {

            $scope.student = {"firstName":"", "lastName": "","age":"","gpa":"" };
            $scope.selectedDepartmentOfStudent ={"id":"","code":"","name":""};
            $scope.messageOfSaveStudent = response.data;
            $scope.getStudentsData();

        }, function myError(response) {
            $scope.messageOfSaveStudent = response.statusText;
        });
    };
    $scope.sendDepartmentData = function () {
        var url = "/saveDepartment", data = {"code": $scope.department.code,"name": $scope.department.name};

        $http.post(url, data).then(function mySuccess(response) {

            $scope.department ={"code":"","name":""};
            $scope.messageOfSaveDepartment = response.data;
            $scope.getDepartmentData();

        }, function myError(response) {
            $scope.messageOfSaveDepartment = response.statusText;
        });
    };
    $scope.sendCourseData = function () {
        var url = "/saveCourse", data = {"title": $scope.course.title,"credit": $scope.course.credit,
            "departmentId":$scope.selectedDepartmentOfCourse.id};

        $http.post(url, data).then(function mySuccess(response) {

            $scope.course = {"id":"","title":"","credit":""};
            $scope.messageOfSaveCourse = response.data;
            $scope.getCourseData();

        }, function myError(response) {
            $scope.messageOfSaveCourse = response.statusText;
        });
    };
    $scope.sendInstructorData = function () {
        var url = "/saveInstructor", data = {"firstName": $scope.instructor.firstName,"lastName":$scope.instructor.lastName,
        "age":$scope.instructor.age};

        $http.post(url, data).then(function mySuccess(response) {

            $scope.instructor = {"id":"","firstName":"", "lastName": "" };
            $scope.messageOfSaveInstructor = response.data;
            $scope.getInstructorData();

        }, function myError(response) {
            $scope.messageOfSaveInstructor = response.statusText;
        });
    };
    $scope.sendSectionData = function () {
        var url = "/saveSection", data = {"courseId": $scope.selectedCourseOfSection.id,"number":$scope.section.number,
            "instructorId":$scope.selectedInstructorOfSection.id};

        $http.post(url, data).then(function mySuccess(response) {

            $scope.section = {"id":"","courseId":"", "number": "" ,"instructorId":""};
            $scope.messageOfSaveSection = response.data;
            $scope.getSectionData();

        }, function myError(response) {
            $scope.messageOfSaveSection = response.statusText;
        });
    };
    $scope.sendEnrollmentData = function () {
        var url = "/saveEnrollment", data = {"studentId": $scope.selectedStudentOfEnrollment.id,
            "sectionId":$scope.selectedSectionOfEnrollment.id};

        $http.post(url, data).then(function mySuccess(response) {

            $scope.messageOfSaveEnrollment = response.data;
            $scope.getEnrollmentData();

        }, function myError(response) {
            $scope.messageOfSaveEnrollment = response.statusText;
        });
    };



    $scope.deleteStudent=function (a, b, c) {
        if(confirm("Are you sure delete this STUDENT : " + b + " "+ "c " +" ?")) {
            $http.post("/deleteStudent", {"id": a, "name": b, "age": c}).then(function mySuccess(response) {

                $scope.getStudentsData();
            }, function myError(response) {
            });
        } else{
            return false;
        }
    };
    $scope.deleteDepartment=function (a, b, c) {
        if(confirm("Are you sure delete this DEPARTMENT : " +b +" - "+ c + " ?")) {
            $http.post("/deleteDepartment", {"id": a, "code": b, "name": c}).then(function mySuccess(response) {

                $scope.getDepartmentData();
            }, function myError(response) {
            });
        } else{
            return false;
        }
    };
    $scope.deleteCourse=function (a, b, c, d) {
        if(confirm("Are you sure delete this COURSE : " + b + " ?")) {
            $http.post("/deleteCourse", {"id": a, "title": b, "credit": c,"departmentId":d}).then(function mySuccess(response) {

                $scope.getCourseData();
            }, function myError(response) {
            });
        } else{
            return false;
        }
    };
    $scope.deleteInstructor=function (a, b, c, d) {
        if(confirm("Are you sure delete this INSTRUCTOR : " + b + " " + c + " " + d +" ?")) {
            $http.post("/deleteInstructor", {"id": a, "firstName": b, "lastName": c,"age":d}).then(function mySuccess(response) {

                $scope.getInstructorData();
            }, function myError(response) {
            });
        } else{
            return false;
        }
    };
    $scope.deleteSection=function (a, b, c, d) {
        if(confirm("Are you sure delete this SECTİON : " + b + "-" + c + " " + d +" ?")) {
            $http.post("/deleteSection", {"id": a, "courseId": b, "number": c,"instructorId":d}).then(function mySuccess(response) {

                $scope.getSectionData();
            }, function myError(response) {
            });
        } else{
            return false;
        }
    };
    $scope.deleteEnrollment=function (a, b) {
        if(confirm("Are you sure delete this Enrollment : " + a + " " + b  +" ?")) {
            $http.post("/deleteEnrollment", {"studentId": a, "sectionId": b}).then(function mySuccess(response) {

                $scope.getEnrollmentData();
            }, function myError(response) {
            });
        } else{
            return false;
        }
    };



    $scope.getStudentsData = function () {
        $http.get("/getAllStudents").then(function mySuccess(response) {
            $scope.allStudents = response.data;
        }, function myError(response) {
            $scope.messageOfSaveStudent = response.statusText;
        });
    };
    $scope.getDepartmentData = function () {
        $http.get("/getAllDepartments").then(function mySuccess(response) {
            $scope.allDepartments = response.data;
        }, function myError(response) {
            $scope.messageOfSaveDepartment = response.statusText;
        });
    };
    $scope.getCourseData = function () {
        $http.get("/getAllCourses").then(function mySuccess(response) {
            $scope.allCourses = response.data;
        }, function myError(response) {
            $scope.messageOfSaveCourse = response.statusText;
        });
    };
    $scope.getInstructorData = function () {
        $http.get("/getAllInstructors").then(function mySuccess(response) {
            $scope.allInstructors = response.data;
        }, function myError(response) {
            $scope.messageOfSaveInstructor = response.statusText;
        });
    };
    $scope.getSectionData = function () {
        $http.get("/getAllSections").then(function mySuccess(response) {
            $scope.allSections = response.data;
        }, function myError(response) {
            $scope.messageOfSaveSection = response.statusText;
        });
    };
    $scope.getEnrollmentData = function () {
        $http.get("/getAllEnrollments").then(function mySuccess(response) {
            $scope.allEnrollments = response.data;
        }, function myError(response) {
            $scope.messageOfSaveEnrollment = response.statusText;
        });
    };

    $scope.getStudentsData();
    $scope.getDepartmentData();
    $scope.getCourseData();
    $scope.getInstructorData();
    $scope.getSectionData();
    $scope.getEnrollmentData();

    $scope.setSection= function(number,  courseTitle,  courseCredit,  departCode,  departName,  instFirst,  instLast,  instAge,students){

        $scope.komplexSection.course = courseTitle +"-"+ number + "    Credit:" + courseCredit ;
        $scope.komplexSection.instructor = instFirst +" "+ instLast + "  " + instAge;
        $scope.komplexSection.departmant =  departCode +" - "+ departName;
        $scope.komplexSection.students =  students;
    };

    $scope.komplexSection ={"course":" ","instructor":"","departmant":"","students":""};


    $scope.komplexStudent = {"student":" ","sections":""};

    $scope.setStudent= function(firstName, lastName, sections){

        $scope.komplexStudent.student = firstName +" "+ lastName ;
        $scope.komplexStudent.sections = sections;
    };


});
