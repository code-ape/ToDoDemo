
var todoAppInterface = angular.module('todo.app.interface', []);

todoAppInterface.controller('InterfaceCtrl', ['$scope', '$log', function ($scope, $log) {

		$scope.login = function (user) {
			if (typeof user == "undefined") {
				$scope.auth_failed();
				return;
			}
			if (!(user["name"] in user_dict)) {
				$scope.auth_failed();
				return;
			}

			var password_hash = sha256_digest(user["password"]);
			if (AttemptAuth(user["name"], password_hash)) {
				$log.log("Auth success.")
			} else {
				$scope.auth_failed();
				return;
			}
		};

		$scope.auth_failed = function () {
			$log.log("Auth failed.");
		};
	}]);