app.factory('employeeService', ['$http', function($http) {
    var apiBaseUrl = 'http://localhost:8080/employees';
  
    return {
      getAll: function() {
        return $http.get(apiBaseUrl);
      },
      getById: function(id) {
        return $http.get(`${apiBaseUrl}/${id}`);
      },
      create: function(employee) {
        return $http.post(apiBaseUrl, employee);
      },
      update: function(id, employee) {
        return $http.put(`${apiBaseUrl}/${id}`, employee);
      },
      delete: function(id) {
        return $http.delete(`${apiBaseUrl}/${id}`);
      }
    };
  }]);
  