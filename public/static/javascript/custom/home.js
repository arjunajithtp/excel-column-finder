myApp.controller('AppCtrl', ['$scope', function($scope) {

    $scope.gridData = [];

    $scope.getWordCount = function() {

        $.ajax({
            url: '/excel-column-finder/home',
            type: 'POST',
            dataType: 'json',
            data : "&startingColumn=" + $scope.startingColumn + "&rows=" + $scope.rows + "&columns=" + $scope.columns,
            success : function(data) {
                $scope.$apply(function(){
                    $.extend(true,$scope.gridData,data);
                    console.log($scope.gridData);
                });
            }
        });
    };


}]);