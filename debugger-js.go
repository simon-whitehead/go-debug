package godebug

var debugger_js = `
	var godebug = angular.module('godebug', ['ngSanitize'])

	godebug.controller('debugController', function($scope, $http) {

		var refreshUI = function() {
			$http.get('/code').then(function(result) {
				$scope.code = result.data.Code;
				$scope.fileName = "File: " + result.data.FileName;
				$scope.lines = result.data.Lines;
        		$scope.breakpoint = result.data.Breakpoint;
        		$scope.stats = result.data.Stats;
			});

			$http.get('/watches').then(function(result) {
				$scope.watches = result.data;
			});
		};

		refreshUI();

		$scope.continue = function() {
			$http.get('/continue').then(function(result) {
				refreshUI();
			},
			function(result) {
				if (result.status == 0) {
					alert('Debugging session has ended');
				}
			});
		};

	    $scope.toggle = function() {
	      $http.get('/toggle?id=' + $scope.breakpoint).then(function() {
	        $scope.continue();
	      });
	    };
	});
`
