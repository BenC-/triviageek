'use strict';

/**
 * @ngdoc function
 * @name triviageekguiApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the triviageekguiApp
 */
 angular.module('triviageekguiApp')
 .controller('MainCtrl', function ($scope, $rootScope, $websocket, $interval) {

 	$scope.screenMode = 'start';
 	$scope.player = {pseudo:""};

 	var dataStream;

 	$scope.startGame = function(){
 		dataStream = $websocket('ws://localhost:9001');
 		dataStream.send(JSON.stringify($scope.player));

 		dataStream.onMessage(function(m) {
    // Log event
    var object = JSON.parse(m.data);
    if(object.hasOwnProperty('name')){ // Game is starting
    	$scope.game = object;
    	var startTime = Date.parse(object.startTime);
    	var now = new Date();
    	$scope.countDown = Math.ceil((startTime-now.getTime())/1000);
    	$interval(function(){$scope.countDown--;},1000);
    } else if (object.hasOwnProperty('smell')) { // Question/Response
    	$scope.screenMode = 'game';
    	$scope.question = object;
if(object.smell.name===""){  // Question, solution is hidden
   $scope.countDown = 20;
   $scope.played = false;
} else {
  $scope.played = true;
}
    } else { // Result
    	$scope.screenMode = 'results';
    	$scope.result = object;
    }
    });
 	};

 	$scope.submitAnswer = function(proposal){
 		if($scope.played){
 			return;
 		}
 		var response = {step : $scope.question.step, value : proposal}; 		
 		dataStream.send(JSON.stringify(response));
 	};

 	

 });
