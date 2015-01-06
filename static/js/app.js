

var todoApp = angular.module('todo.app', [
	'ngRoute',
	'ngAnimate',
	'todo.app.login',
	'todo.app.interface',
	]);

todoApp.config(['$routeProvider', function ($routeProvider) {
		$routeProvider.
		  when('/login', {
		  	templateUrl: 'login.html',
		  	controller: 'AuthCtrl'
		  }).
		  when('/to-do', {
		  	templateUrl: 'todo.html',
		  	controller: 'InterfaceCtrl'
		  }).
		  otherwise({
		  	redirectTo: '/login'
		  });
	}]);

todoApp.run(["$rootScope", '$location', function ($rootScope, $location) {
	$rootScope.auth_token = "";
	$rootScope.active_user = "";
	$rootScope.to_dos = {};
	$rootScope.new_todo_field = false;
	$location.path("/login")
}]);