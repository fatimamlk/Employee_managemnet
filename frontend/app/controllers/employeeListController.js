app.controller('employeeListController', function($scope, employeeService) {
    $scope.employees = [];

    employeeService.getEmployees().then(function(response) {
        $scope.employees = response.data;
    }).catch(function(error) {
        console.error('Erreur lors de la récupération des employés:', error);
    });

    $scope.deleteEmployee = function(id) {
        employeeService.deleteEmployee(id).then(function() {
            $scope.employees = $scope.employees.filter(emp => emp.id !== id);
        }).catch(function(error) {
            console.error('Erreur lors de la suppression de l\'employé:', error);
        });
    };
});
