$(document).ready(
    function()
    {
        $('#flyin').effect("slide", {}, 300);
        $('.flyin').effect("shake", {}, 300);
    }
);

var init = function(){
    var input1 = document.getElementById('f_elem_city1');
    var input2 = document.getElementById('f_elem_city2');
    var input3 = document.getElementById('f_elem_city3');
    var input4 = document.getElementById('f_elem_city4');
    var input5 = document.getElementById('f_elem_city5');
    var options = {
        types: ['(cities)']
    };
    autocomplete1 = new google.maps.places.Autocomplete(input1, options);
    autocomplete2 = new google.maps.places.Autocomplete(input2, options);
    autocomplete3 = new google.maps.places.Autocomplete(input3, options);
    autocomplete4 = new google.maps.places.Autocomplete(input4, options);
    autocomplete5 = new google.maps.places.Autocomplete(input5, options);
}

google.maps.event.addDomListener(window, 'load', init);