'use strict';

/**
 * @ngdoc function
 * @name triviageekguiApp.controller:AboutCtrl
 * @description
 * # AboutCtrl
 * Controller of the triviageekguiApp
 */
angular.module('triviageekguiApp', ['ngWebSocket'])
  .controller('GameCtrl', function ($websocket) {

  	var dataStream = $websocket('ws://localhost:9000');


    dataStream.onMessage(function(message) {
        // If game

        // If question

        // If result
    });




    $scope.sendQuestion = function youpi(){

    dataStream.send(JSON.stringify({ action: 'get' }));
}

    this.awesomeThings = [];
  });
