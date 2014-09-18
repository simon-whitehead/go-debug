package godebug

// Not sure if this is "best practice" .. couldn't think of any other way though:

var debugger_markup = `
  <!DOCTYPE html>
  <html lang="en" ng-app="godebug">
  <head>
    <meta charset="utf-8" />
    <title>GoDebug - Browser-based debugger for Golang</title>
    <meta name="author" content="Simon Whitehead">
    <meta name="description" content="go, golang, de, debug, debugger, ide" />
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1" /> 
    <link rel="stylesheet" href="/debugger.css" type="text/css" />
    <link rel="stylesheet" href="//maxcdn.bootstrapcdn.com/bootstrap/3.2.0/css/bootstrap.min.css" type="text/css"/>
  </head>
  <body ng-app="godebug" ng-controller="debugController">

      <div class="navbar navbar-default navbar-fixed-top" role="navigation">
        <div class="container">
          <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target=".navbar-collapse">
              <span class="sr-only">Toggle navigation</span>
              <span class="icon-bar"></span>
              <span class="icon-bar"></span>
              <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="/">GoDebug</a>
          </div>
          <div class="collapse navbar-collapse">
            <ul class="nav navbar-nav pull-right">
              <li><a href="#" ng-click="continue()"><button class="btn btn-info btn-sm"><span class="glyphicon glyphicon-play" style="margin-right: 5px;"></span>Continue</button></a></li>
              <li><a href="#" ng-click="toggle()"><button class="btn btn-warning btn-sm"><span class="glyphicon glyphicon-fast-forward" style="margin-right: 5px;"></span>Skip</button></a></li>
            </ul>
          </div>
        </div>
      </div>

      <div class="container">
        <div class="col-sm-8">
        	<h4 ng-bind="fileName"></h4>
        	<pre style="float:left; text-align:right;" ng-bind-html="lines">
        	</pre>
          <div ng-bind-html="code">
          </div>
        </div>
        <div class="col-sm-4">
          <h4>Watches</h4>
              <table class="table table-bordered table-striped">
                <thead>
                  <tr>
                      <th>#</th>
                      <th>Variable</th>
                      <th>Value</th>
                  </tr>
                </thead>
                <tbody>
                  <tr ng-repeat="watch in watches">
                    <td>
                      <span ng-bind="watch.Index"></span>
                    </td>
                    <td>
                      <span ng-bind="watch.Name"></span>
                    </td>
                    <td>
                      <span ng-bind="watch.Value"></span>
                    </td>
                  </tr>
                </tbody>
            </table>
        </div>
      </div>

        <div class="container">
          <div class="col-sm-12">
          <h4>Runtime stats</h4>
          <table class="table table-bordered table-striped">
              <thead>
                <tr>
                    <th>Stat</th>
                    <th>Value</th>
                </tr>
              </thead>
              <tbody>
                <tr ng-repeat="stat in stats">
                  <td>
                    <span ng-bind="stat.Name"></span>
                  </td>
                  <td>
                    <span ng-bind="stat.Value"></span>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
    
    <script src="//cdnjs.cloudflare.com/ajax/libs/angular.js/1.2.20/angular.min.js"></script>
    <script src="//cdnjs.cloudflare.com/ajax/libs/angular.js/1.2.20/angular-sanitize.min.js"></script>
    <script src="/debugger.js"></script>
  </body>
  </html>

`
