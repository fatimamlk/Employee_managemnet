app.controller('AddEmployeeController', ['$http', '$location', function($http, $location) {
    var vm = this;
    vm.employee = {};

    // Fonction pour soumettre le formulaire et ajouter un employé
    vm.addEmployee = function() {
        $http.post('http://localhost:8080/employees', vm.employee)
            .then(function(response) {
                // Rediriger vers la liste des employés après l'ajout
                $location.path('/');
            }, function(error) {
                console.error("Erreur lors de l'ajout de l'employé", error);
            });
    };
}]);
