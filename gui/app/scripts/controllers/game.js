'use strict';

/**
 * @ngdoc function
 * @name triviageekguiApp.controller:AboutCtrl
 * @description
 * # AboutCtrl
 * Controller of the triviageekguiApp
 */
angular.module('triviageekguiApp')
  .controller('GameCtrl', function ($rootscope) {

  	
    $rootscope.dataStream.onMessage(function(message) {
        // If game
        console.log(message);

        // If question

        // If result
    });



/*
    $scope.sendQuestion = function youpi(){

    dataStream.send(JSON.stringify({ action: 'get' }));*/


    this.awesomeThings = [];
  });
