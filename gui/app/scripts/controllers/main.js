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


 	var dataStream = $websocket('ws://localhost:9001');

 	$scope.startGame = function(){
 		dataStream.send(JSON.stringify($scope.player));
 	};

 	$scope.submitAnswer = function(proposal){
 		var response = {step : $scope.question.step, success : (proposal===$scope.question.smell.name)}
 		dataStream.send(JSON.stringify(response));
 	};

 	dataStream.onMessage(function(m) {
    // Log event
    console.log(m);
    var object = JSON.parse(m.data);
    if(object.hasOwnProperty('name')){ // Game is starting
    	$scope.game = object;
    	var startTime = Date.parse(object.startTime);
    	var now = new Date();
    	$scope.countDown = Math.ceil((startTime-now.getTime())/1000);
    	$interval(function(){$scope.countDown--;},1000);
    } else if (object.hasOwnProperty('smell')) { // Question
    	$scope.screenMode = 'game';
    	$scope.question = object;
    } else { // Result
    	$scope.screenMode = 'results';
    	$scope.result = object;
    }

        // If question

        // If result
    });

 });
