

var todoApp = angular.module('todo.app', [
	'ngRoute',
	'ngAnimate',
	'todo.app.login',
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