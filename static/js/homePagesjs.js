
angular.module("myApp",[]).controller("myCtrl", function ($scope, $http){

    $scope.messageOfSaveUniversity = "";
    $scope.messageOfSaveStudent = "";
    $scope.messageOfSaveSC ="";
    $scope.messageOfSaveTC ="";
    $scope.messageOfSaveTeacher="";
    $scope.messageOfSaveClass="";

    $scope.university = {"name":"", "capacity": ""};
    $scope.student = {"name":"", "age": "" };
    $scope.className ="";
    $scope.teacherName = "";

    $scope.allStudents ="";
    $scope.allTeachers ="";
    $scope.allClasses ="";
    $scope.allUniversities ="";

    $scope.selectedUniversityOfStudent ={"id":"","name":"","capacity":"","number":""};

    $scope.selectedStudentOfSC={"id":"","name":"","age":"","uni":""};
    $scope.selectedClassOfSC = {"id":"","name":""};

    $scope.selectedTeacherOfTC={"id":"","name":""};
    $scope.selectedClassOfTC = {"id":"","name":""};

    $scope.sendUniversityData = function () {
        var url = "/saveUniversity";

        $http.post(url, $scope.university).then(function mySuccess(response) {

            $scope.messageOfSaveStudent = "";
            $scope.messageOfSaveSC ="";
            $scope.messageOfSaveTC ="";
            $scope.messageOfSaveTeacher="";
            $scope.messageOfSaveClass="";

            $scope.university="";

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

            $scope.student="";

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

        $http.post(url,  {"name":$scope.className}).then(function mySuccess(response) {

            $scope.messageOfSaveUniversity = "";
            $scope.messageOfSaveStudent = "";
            $scope.messageOfSaveSC ="";
            $scope.messageOfSaveTC ="";
            $scope.messageOfSaveTeacher="";

            $scope.className ="";

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

        }, function myError(response) {
            $scope.messageOfSaveTeacher = response.statusText;
        });
    };

    $scope.sendSCData = function () {
        var url = "/saveStudentAndClass", data = {"s_id": $scope.selectedStudentOfSC.id,
                                                  "c_id" : $scope.selectedClassOfSC.id};

        $http.post(url,  data).then(function mySuccess(response) {

            $scope.messageOfSaveUniversity = "";
            $scope.messageOfSaveStudent = "";
            $scope.messageOfSaveTC ="";
            $scope.messageOfSaveTeacher="";
            $scope.messageOfSaveClass="";

            $scope.selectedStudentOfSC="";
            $scope.selectedClassOfSC="";

            $scope.messageOfSaveSC = response.data;

        }, function myError(response) {
            $scope.messageOfSaveSC = response.statusText;
        });
    };

    $scope.sendTCData = function () {
        var url = "/saveTeacherAndClass", data = {"t_id": $scope.selectedTeacherOfTC.id,
            "c_id" : $scope.selectedClassOfTC.id};

        $http.post(url,  data).then(function mySuccess(response) {

            $scope.messageOfSaveUniversity = "";
            $scope.messageOfSaveStudent = "";
            $scope.messageOfSaveSC ="";
            $scope.messageOfSaveTeacher="";
            $scope.messageOfSaveClass="";

            $scope.selectedTeacherOfTC="";
            $scope.selectedClassOfTC="";

            $scope.messageOfSaveTC = response.data;

        }, function myError(response) {
            $scope.messageOfSaveTC = response.statusText;
        });
    };

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

    $http.get("/getAllClasses").then(function mySuccess(response) {
        $scope.allClasses = response.data;
    }, function myError(response) {
        $scope.messageOfSaveClass = response.statusText;
    });

    $http.get("/getAllTeachers").then(function mySuccess(response) {
        $scope.allTeachers = response.data;
    }, function myError(response) {
        $scope.messageOfSaveTeacher = response.statusText;
    });
});