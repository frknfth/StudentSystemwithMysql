angular.module("myApp",[]).controller("myCtrl", function($scope, $http) {

    $scope.allStudents ="";
    $scope.allUniversities ="";
    $scope.allTeachers ="";
    $scope.allClasses ="";

    $scope.university={"name":"", "capacity": ""};
    $scope.student={"name":"", "age": ""};

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
            $http.post("/deleteClass", {"id": a, "name": b}).then(function mySuccess(response) {
                $scope.getAllData();
            }, function myError(response) {
            });
        } else{
            return false;
        }
    };

    $scope.getUniData = function () {
        $http.get("/getAllUniversities").then(function mySuccess(response) {
            $scope.allUniversities = response.data;
        }, function myError(response) {
            $scope.messageOfSaveStudent = response.statusText;
        });
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
        $http.get("/getAllStudents").then(function mySuccess(response) {
            $scope.allStudents = response.data;
        }, function myError(response) {
            $scope.messageOfSaveStudent = response.statusText;
        });
    };

    $http.get("/getAllUniversities").then(function mySuccess(response) {
        $scope.allUniversities = response.data;
    }, function myError(response) {
        $scope.messageOfSaveUniversity = response.statusText;
    });
});