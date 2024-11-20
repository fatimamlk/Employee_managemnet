var app = angular.module('employeeApp', ['ngRoute']);

app.config(['$routeProvider', function($routeProvider) {
  $routeProvider
    .when('/', {
      templateUrl:'app/views/employeeList.html',
      controller: 'employeeListController'
    })
    .when('/add', {
      templateUrl: 'app/views/addEmployee.html',
      controller: 'addEmployeeController'
    })
    .when('/edit/:id', {
      templateUrl: 'app/views/editEmployee.html',
      controller: 'editEmployeeController'
    })
    .otherwise({
      redirectTo: '/'
}]);
