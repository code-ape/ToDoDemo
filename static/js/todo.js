
var todoAppInterface = angular.module('todo.app.interface', []);

todoAppInterface.controller('InterfaceCtrl', [
	'$scope', '$rootScope', '$http', '$location', '$log', 
	function ($scope, $rootScope, $http, $location, $log) {

		$scope.to_dos = [];
		$scope.users  = [];

		$scope.logout = function () {
			$http.post("/api/logout", {"user": $rootScope.active_user, "token": $rootScope.auth_token}).
				success(function(data, status, headers, config) {
					if (data["status"] === "success") {
						$rootScope.auth_token = "";
						$rootScope.active_user = "";
						$log.log("Logout successful.");
						$location.path('/login');
					} else {
						$log.log("Logout failed.");
					}
				}).
				error(function(data, status, headers, config) {
					$log.log("status: " + status)
					$log.log("data: " + data)
					$log.log("Logout failed.");
				});
		};

		$scope.refresh = function () {
			$http.post("/api/get-to-dos", {"user": $rootScope.active_user, "token": $rootScope.auth_token}).
				success(function(data, status, headers, config) {
					if (data["status"] === "success") {
						$rootScope.to_dos = data["to_dos"];
						$scope.to_dos = data["to_dos"];
						$scope.users  = data["users"];
						$log.log("Get to-dos successful.");
					} else {
						$log.log("Get to-dos failed.");
					}
				}).
				error(function(data, status, headers, config) {
					$log.log("status: " + status)
					$log.log("data: " + data)
					$log.log("Logout failed.");
				});
		};

		$scope.new_todo_action = function () {
			$log.log("new_todo_field toggled");
			$rootScope.new_todo_field = !$rootScope.new_todo_field;
			$scope.new_todo_field = $rootScope.new_todo_field
		};

		$scope.send_new_todo = function (text) {
			$http.post("/api/add-to-dos", {"user": $rootScope.active_user, 
						"token": $rootScope.auth_token, 
						"to_dos":[{"user":$rootScope.active_user, "text": text}]}).
				success(function(data, status, headers, config) {
					$log.log("data[status] = " + data["status"])
					if (data["status"] === "success") {
						$scope.refresh();
						$log.log("Send to-dos successful.");
					} else {
						$log.log("Send to-dos failed.");
					}
				}).
				error(function(data, status, headers, config) {
					$log.log("status: " + status)
					$log.log("data: " + data)
					$log.log("Send to-dos failed.");
				});
		};

	}]);
