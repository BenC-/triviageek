'use strict';

/**
 * @ngdoc function
 * @name triviageekguiApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the triviageekguiApp
 */
angular.module('triviageekguiApp')
  .controller('MainCtrl', function ($scope, $rootScope, $websocket, $location) {

  	$scope.player={pseudo : "yuyu"};
	
  	 $scope.startGame = function(){
  	 	$rootScope.dataStream = $websocket('ws://localhost:9001');


  	 	console.log($scope.player.pseudo);
  	 	$location.path( "/game" );

  	 };
    
  });
