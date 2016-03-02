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
    'ngTouch'
  ])
  .config(function ($routeProvider) {
    $routeProvider
      .when('/main', {
        templateUrl: 'views/main.html',
        controller: 'MainCtrl',
        controllerAs: 'main'
      })
      .when('/game', {
        templateUrl: 'views/game.html',
        controller: 'GameCtrl',
        controllerAs: 'game'
      })
      .otherwise({
        redirectTo: '/main'
      });
  });
