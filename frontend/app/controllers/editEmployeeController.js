app.controller('EditEmployeeController', ['$http', '$routeParams', '$location', function($http, $routeParams, $location) {
    var vm = this;
    vm.employee = {};

    // Récupérer les informations de l'employé à éditer avec l'ID passé dans l'URL
    var employeeId = $routeParams.id;
    $http.get('http://localhost:8080/employees/' + employeeId)
        .then(function(response) {
            vm.employee = response.data;
        }, function(error) {
            console.error("Erreur lors de la récupération de l'employé", error);
        });

    // Fonction pour mettre à jour l'employé
    vm.updateEmployee = function() {
        $http.put('http://localhost:8080/employees/' + employeeId, vm.employee)
            .then(function(response) {
                // Rediriger vers la liste des employés après la mise à jour
                $location.path('/');
            }, function(error) {
                console.error("Erreur lors de la mise à jour de l'employé", error);
            });
    };
}]);
