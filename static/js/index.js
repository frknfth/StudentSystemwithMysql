var app = angular.module("myApp",['angularUtils.directives.dirPagination']);

app.controller("myCtrl", function($scope, $http) {

    $scope.messageOfSaveUniversity = "";
    $scope.messageOfSaveStudent = "";
    $scope.messageOfSaveSC ="";
    $scope.messageOfSaveTC ="";
    $scope.messageOfSaveTeacher="";
    $scope.messageOfSaveClass="";

    $scope.university = {"name":"", "capacity": ""};
    $scope.student = {"name":"", "age": "" };
    $scope.class ={"dep":"","code":"","name":""};
    $scope.teacherName = "";

    $scope.allUniversities ="";
    $scope.allStudents ="";
    $scope.allTeachers ="";
    $scope.allClasses ="";
    $scope.allDepartments="";

    $scope.selectedUniversityOfStudent ={"id":"","name":"","capacity":"","number":""};
    $scope.selectedDepartmentOfClass = {"id":"","name":""};

    $scope.selectedStudentOfSC={"id":"","name":"","age":"","uni":""};
    $scope.selectedClassOfSC = {"id":"","name":""};

    $scope.selectedTeacherOfTC={"id":"","name":""};
    $scope.selectedClassOfTC = {"id":"","name":""};


    $scope.currentPageUni = 1;
    $scope.currentPageStd = 1;
    $scope.currentPageClass=1;
    $scope.currentPageTeacher=1;
    $scope.pageSize = 10;


    $scope.getUniData = function () {
        $http.get("/getAllUniversities").then(function mySuccess(response) {
            $scope.allUniversities = response.data;
            $scope.currentPage = 1;
        }, function myError(response) {
            $scope.messageOfSaveStudent = response.statusText;
        });
    };

    $scope.sendUniversityData = function () {
        var url = "/saveUniversity";

        $http.post(url, $scope.university).then(function mySuccess(response) {

            $scope.messageOfSaveStudent = "";
            $scope.messageOfSaveSC ="";
            $scope.messageOfSaveTC ="";
            $scope.messageOfSaveTeacher="";
            $scope.messageOfSaveClass="";

            $scope.university={"name":"", "capacity": ""};

            $scope.messageOfSaveUniversity = response.data;

            $http.get("/getAllUniversities").then(function mySuccess(response) {
                $scope.allUniversities = response.data;
            }, function myError(response) {
                $scope.messageOfSaveUniversity = response.statusText;
            });

        }, function myError(response) {
            $scope.messageOfSaveUniversity = response.statusText;
        });
    };

    $scope.sendStudentData = function () {
        var url = "/saveStudent", data = {"name": $scope.student.name,"age": $scope.student.age,
            "uni": $scope.selectedUniversityOfStudent.id};

        $http.post(url, data).then(function mySuccess(response) {
            $scope.messageOfSaveUniversity = "";
            $scope.messageOfSaveSC ="";
            $scope.messageOfSaveTC ="";
            $scope.messageOfSaveTeacher="";
            $scope.messageOfSaveClass="";

            $scope.student={"name":"", "age": ""};
            $scope.selectedUniversityOfStudent ={"id":"","name":"","capacity":"","number":""};

            $scope.messageOfSaveStudent = response.data;

            $http.get("/getAllUniversities").then(function mySuccess(response) {
                $scope.allUniversities = response.data;
            }, function myError(response) {
                $scope.messageOfSaveUniversity = response.statusText;
            });

            $http.get("/getAllStudents").then(function mySuccess(response) {
                $scope.allStudents = response.data;
            }, function myError(response) {
                $scope.messageOfSaveStudent = response.statusText;
            });

        }, function myError(response) {
            $scope.messageOfSaveStudent = response.statusText;
        });
    };

    $scope.sendClassData = function () {
        var url = "/saveClass";

        $scope.class.dep = $scope.selectedDepartmentOfClass.name;

        $http.post(url,  $scope.class).then(function mySuccess(response) {

            $scope.messageOfSaveUniversity = "";
            $scope.messageOfSaveStudent = "";
            $scope.messageOfSaveSC ="";
            $scope.messageOfSaveTC ="";
            $scope.messageOfSaveTeacher="";

            $scope.class ={"dep":"","code":"","name":""};
            $scope.selectedDepartmentOfClass = {"id":"","name":""};

            $scope.getClassData();
            $scope.messageOfSaveClass = response.data;

            $http.get("/getAllClasses").then(function mySuccess(response) {
                $scope.allClasses = response.data;
            }, function myError(response) {
                $scope.messageOfSaveClass = response.statusText;
            });

        }, function myError(response) {
            $scope.messageOfSaveClass = response.statusText;
        });
    };

    $scope.sendTeacherData = function () {
        var url = "/saveTeacher", data = {"name": $scope.teacherName};

        $http.post(url,  data).then(function mySuccess(response) {

            $scope.messageOfSaveUniversity = "";
            $scope.messageOfSaveStudent = "";
            $scope.messageOfSaveSC ="";
            $scope.messageOfSaveTC ="";
            $scope.messageOfSaveTeacher="";
            $scope.messageOfSaveClass="";

            $scope.teacherName ="";

            $scope.messageOfSaveTeacher = response.data;

            $http.get("/getAllTeachers").then(function mySuccess(response) {
                $scope.allTeachers = response.data;
            }, function myError(response) {
                $scope.messageOfSaveTeacher = response.statusText;
            });

        }, function myError(response) {
            $scope.messageOfSaveTeacher = response.statusText;
        });
    };

    $scope.deleteUniversity=function (a, b, c, d) {
        if(confirm("are you sure delete this university and all students in it : " + b + " ?")) {
            $http.post("/deleteUniversity", {"id": a, "name": b, "capacity": c, "number": d}).then(function mySuccess
                (response) {
                $scope.getUniData();
            }, function myError(response) {
            });
        } else{
            return false;
        }
    };

    $scope.deleteStudent=function (a, b, c, d) {
        if(confirm("are you sure delete this student : " + b + " ?")) {
            $http.post("/deleteStudent", {"id": a, "name": b, "age": c, "uni": d}).then(function mySuccess(response) {

                $scope.getStudentsData();
                $scope.getUniData();
            }, function myError(response) {
            });
        } else{
            return false;
        }
    };

    $scope.deleteTeacher=function (a, b) {
        if(confirm("are you sure delete this university and all students in it? : " + b)) {
            $http.post("/deleteTeacher", {"id": a, "name": b}).then(function mySuccess(response) {
                $scope.getAllData();
            }, function myError(response) {
            });
        } else{
            return false;
        }
    };

    $scope.deleteClass=function (a, b) {
        if(confirm("are you sure delete this class? : " + b)) {
            $http.post("/deleteClass", {"id": a}).then(function mySuccess(response) {
                $scope.getClassData();
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

    $scope.getClassData = function () {
        $http.get("/getAllClasses").then(function mySuccess(response) {
            $scope.allClasses = response.data;
        }, function myError(response) {
            $scope.messageOfSaveClass = response.statusText;
        });
    };

    $scope.getTeacherData = function () {
        $http.get("/getAllTeachers").then(function mySuccess(response) {
            $scope.allTeachers = response.data;
        }, function myError(response) {
            $scope.messageOfSaveStudent = response.statusText;
        });
    };

    $http.get("/getAllUniversities").then(function mySuccess(response) {
        $scope.allUniversities = response.data;
    }, function myError(response) {
        $scope.messageOfSaveUniversity = response.statusText;
    });

    $http.get("/getAllDepartments").then(function mySuccess(response) {
        $scope.allDepartments = response.data;
    }, function myError(response) {
        $scope.messageOfSaveClass = response.statusText;
    });





    $scope.sort = function(keyname){
        $scope.sortKey = keyname;   //set the sortKey to the param passed
        $scope.reverse = !$scope.reverse; //if true make it false and vice versa
    }
});
