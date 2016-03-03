'use strict';

/**
 * @ngdoc overview
 * @name triviageekguiApp
 * @description
 * # triviageekguiApp
 *
 * Main module of the application.
 */
angular
  .module('triviageekguiApp', [
    'ngResource',
    'ngRoute',
    'ngSanitize',
    'ngTouch',
    'ngWebSocket'
  ])
  .config(function ($routeProvider) {
    $routeProvider
      .when('/main', {
        templateUrl: 'views/main.html',
        controller: 'MainCtrl'
      })
      .when('/game', {
        templateUrl: 'views/game.html',
        controller: 'GameCtrl',
        controllerAs: 'game'
      })
       .when('/results', {
        templateUrl: 'views/results.html',
        controller: 'ResultsCtrl',
        controllerAs: 'results'
      })
      .otherwise({
        redirectTo: '/main'
      });
  });
