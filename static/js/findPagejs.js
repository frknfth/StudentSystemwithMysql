
angular.module("myApp",[]).controller("myCtrl", function($scope, $http){

    $scope.studentName="";
    $scope.students = "";

    $scope.universityName="";
    $scope.universities = "";

    $scope.teacherName ="";
    $scope.teachers ="";

    $scope.className ="";
    $scope.classes ="";

    $scope.findStudent=function () {
        $http.get("/getStudentData",{headers:{'Name':$scope.studentName}}).then(function mySuccess(response) {
            $scope.students = response.data;
        }, function myError(response) {
        });
    };

    $scope.findUniversity=function () {
        $http.get("/getUniversityData",{headers:{'Name':$scope.universityName}}).then(function mySuccess(response) {
            $scope.universities = response.data;
        }, function myError(response) {
        });
    };

    $scope.findTeacher = function () {
        $http.get("/getTeacherData",{headers:{'Name':$scope.teacherName}}).then(function mySuccess(response) {
            $scope.teachers = response.data;
        }, function myError(response) {
        });
    };

    $scope.findClass = function () {
        $http.get("/getClassData",{headers:{'Name':$scope.className}}).then(function mySuccess(response) {
            $scope.classes = response.data;
        }, function myError(response) {
        });
    };


    $scope.deleteStudent=function (a, b, c, d) {
        if(confirm("are you sure delete this student? : " + b)) {
            $http.post("/deleteStudent", {"id": a, "name": b, "age": c, "uni": d}).then(function mySuccess(response) {
                $scope.findStudent();
                $scope.findUniversity();
            }, function myError(response) {
            });
        } else{
            return false;
        }
    };

    $scope.deleteUniversity=function (a, b, c, d) {
        if(confirm("are you sure delete this university and all students in it? : " + b)) {
            $http.post("/deleteUniversity", {"id": a, "name": b, "capacity": c, "number": d}).then(function mySuccess(response) {
                $scope.findStudent();
                $scope.findUniversity();
            }, function myError(response) {
            });
        } else{
            return false;
        }
    };

    $scope.deleteTeacher=function (a, b) {
        if(confirm("are you sure delete this university and all students in it? : " + b)) {
            $http.post("/deleteTeacher", {"id": a, "name": b}).then(function mySuccess(response) {
                $scope.findTeacher();
            }, function myError(response) {
            });
        } else{
            return false;
        }
    };

    $scope.deleteClass=function (a, b) {
        if(confirm("are you sure delete this class? : " + b)) {
            $http.post("/deleteClass", {"id": a, "name": b}).then(function mySuccess(response) {
                $scope.findClass();
            }, function myError(response) {
            });
        } else{
            return false;
        }
    };

});