
angular.module("myApp",[]).controller("myCtrl", function ($scope, $http){

    $scope.university = {"name":"", "capacity": ""};
    $scope.messageOfSaveUniversity = "";

    $scope.student = {"name":"", "age": "" };
    $scope.messageOfSaveStudent = "";

    $scope.allUniversities ="";
    $scope.selectedUniversityOfStudent ={"id":"","name":"","capacity":"","number":""};

    $scope.className ="";

    $scope.selectedUniversityOfTeacher = {"id":"","name":"","capacity":"","number":""};
    $scope.selectedClassOfTeacher = {"id":"","name":""};
    $scope.teacherName = "";
    $scope.notSelectedAllClasses ="";

    $scope.allStudents ="";
    $scope.allClasses ="";
    $scope.selectedStudentOfSC={"id":"","name":"","age":"","uni":""};
    $scope.selectedClassOfSC = {"id":"","name":""};
    $scope.messageOfSaveSC ="";

    $scope.sendUniversityData = function () {
        var url = "/saveUniversity";

        $http.post(url, $scope.university).then(function mySuccess(response) {

            $scope.messageOfSaveStudent = "";
            $scope.messageOfTeacherAndClass = "";
            $scope.messageOfSaveSC ="";

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
            $scope.messageOfTeacherAndClass = "";
            $scope.messageOfSaveSC ="";

            $scope.messageOfSaveStudent = response.data;

            $http.get("/getAllUniversities").then(function mySuccess(response) {
                $scope.allUniversities = response.data;
            }, function myError(response) {
                $scope.messageOfSaveUniversity = response.statusText;
            });

            $http.get("/getAllStudents").then(function mySuccess(response) {
                $scope.allStudents = response.data;
            }, function myError(response) {
                $scope.messageOfTeacherAndClass = response.statusText;
            });

            $http.get("/getAllClasses").then(function mySuccess(response) {
                $scope.allClasses = response.data;
            }, function myError(response) {
                $scope.messageOfTeacherAndClass = response.statusText;
            });

        }, function myError(response) {
            $scope.messageOfSaveStudent = response.statusText;
        });
    };

    $scope.sendClassData = function () {
        var url = "/saveClass";

        $http.post(url,  {"name":$scope.className}).then(function mySuccess(response) {

            $scope.messageOfSaveStudent = "";
            $scope.messageOfSaveUniversity = "";
            $scope.messageOfSaveSC ="";
            $scope.messageOfTeacherAndClass="";
            $scope.className ="";

            $scope.messageClass = response.data;


            $http.get("/getAllClasses").then(function mySuccess(response) {
                $scope.allClasses = response.data;
            }, function myError(response) {
                $scope.messageClass = response.statusText;
            });
            $http.get("/getNotSelectedClasses").then(function mySuccess(response) {
                $scope.notSelectedAllClasses = response.data;
            }, function myError(response) {
                $scope.messageOfTeacherAndClass = response.statusText;
            });

        }, function myError(response) {
            $scope.messageClass = response.statusText;
        });
    };

    $scope.sendTeacherData = function () {
        var url = "/saveTeacher", data = {"name": $scope.teacherName,
            "classId" : $scope.selectedClassOfTeacher.id,
            "uniId"     : $scope.selectedUniversityOfTeacher.id};

        $http.post(url,  data).then(function mySuccess(response) {

            $scope.messageOfSaveStudent = "";
            $scope.messageOfSaveUniversity = "";
            $scope.messageOfSaveSC ="";

            $scope.messageOfTeacherAndClass = response.data;

            $http.get("/getNotSelectedClasses").then(function mySuccess(response) {
                $scope.notSelectedAllClasses = response.data;
            }, function myError(response) {
                $scope.messageOfTeacherAndClass = response.statusText;
            });

        }, function myError(response) {
            $scope.messageOfSaveUniversity = response.statusText;
        });
    };

    $scope.sendSCData = function () {
        var url = "/saveStudentAndClass", data = {"s_id": $scope.selectedStudentOfSC.id,
                                                  "c_id" : $scope.selectedClassOfSC.id};

        $http.post(url,  data).then(function mySuccess(response) {

            $scope.messageOfSaveStudent = "";
            $scope.messageOfSaveUniversity = "";
            $scope.messageOfTeacherAndClass ="";

            $scope.messageOfSaveSC = response.data;

        }, function myError(response) {
            $scope.messageOfSaveSC = response.statusText;
        });
    };


    $http.get("/getAllUniversities").then(function mySuccess(response) {
        $scope.allUniversities = response.data;
    }, function myError(response) {
        $scope.messageOfTeacherAndClass = response.statusText;
    });

    $http.get("/getAllStudents").then(function mySuccess(response) {
        $scope.allStudents = response.data;
    }, function myError(response) {
        $scope.messageOfTeacherAndClass = response.statusText;
    });

    $http.get("/getAllClasses").then(function mySuccess(response) {
        $scope.allClasses = response.data;
    }, function myError(response) {
        $scope.messageOfTeacherAndClass = response.statusText;
    });

    $http.get("/getNotSelectedClasses").then(function mySuccess(response) {
        $scope.notSelectedAllClasses = response.data;
    }, function myError(response) {
        $scope.messageOfTeacherAndClass = response.statusText;
    });


});