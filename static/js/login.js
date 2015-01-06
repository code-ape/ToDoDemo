
var todoAppLogin = angular.module('todo.app.login', []);

todoAppLogin.controller('AuthCtrl', ['$scope', '$rootScope', '$location','$http', '$log', 
	function ($scope, $rootScope, $location, $http, $log) {

		$scope.invalid = false;
		$scope.attempt_number = 1;

		$scope.login = function (user) {
			$scope.invalid = false;
			if (typeof user == "undefined") {
				$scope.auth_failed();
				return;
			}

			var password_hash = sha256_digest(user["password"]);
			
			$http.post("/api/login", {"user": user["name"], "password": password_hash}).
				success(function(data, status, headers, config) {
					if (data["status"] === "success") {
						$rootScope.auth_token = data["token"];
						$rootScope.active_user = user["name"];
						$log.log("Token: " + data["token"]);
						$scope.auth_success();
					} else {
						$scope.auth_failed();
					}
				}).
				error(function(data, status, headers, config) {
					$log.log("status: " + status)
					$log.log("data: " + data)
					$scope.auth_failed();
				});
		};

	$scope.auth_success = function () {
			$log.log("Auth success.");
			$scope.attempt_number = 0;
			$location.path('/to-do');
		};

	$scope.auth_failed = function () {
			$scope.invalid = true;
			$scope.attempt_number++;
			$log.log("Auth failed.");
		};
	}]);

//var user_dict = {"TestUser": "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"};

//var AttemptAuth = function(username, password_hash) {
//	return password_hash == user_dict[username];
//};