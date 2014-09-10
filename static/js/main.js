
jQuery(function ()
{
    jQuery("#f_elem_city1").autocomplete({
        source: function (request, response) {
            jQuery.getJSON(
                "http://gd.geobytes.com/AutoCompleteCity?callback=?&q="+request.term,
                function (data) {
                    response(data);
                }
            );
        },
        minLength: 3,
        select: function (event, ui) {
            var selectedObj = ui.item;
            jQuery("#f_elem_city1").val(selectedObj.value);
            return false;
        },
        open: function () {
            jQuery(this).removeClass("ui-corner-all").addClass("ui-corner-top");
        },
        close: function () {
            jQuery(this).removeClass("ui-corner-top").addClass("ui-corner-all");
        }
    });
    jQuery("#f_elem_city1").autocomplete("option", "delay", 100);


    jQuery("#f_elem_city2").autocomplete({
        source: function (request, response) {
            jQuery.getJSON(
                "http://gd.geobytes.com/AutoCompleteCity?callback=?&q="+request.term,
                function (data) {
                    response(data);
                }
            );
        },
        minLength: 3,
        select: function (event, ui) {
            var selectedObj = ui.item;
            jQuery("#f_elem_city2").val(selectedObj.value);
            return false;
        },
        open: function () {
            jQuery(this).removeClass("ui-corner-all").addClass("ui-corner-top");
        },
        close: function () {
            jQuery(this).removeClass("ui-corner-top").addClass("ui-corner-all");
        }
    });
    jQuery("#f_elem_city2").autocomplete("option", "delay", 100);


    jQuery("#f_elem_city3").autocomplete({
        source: function (request, response) {
            jQuery.getJSON(
                "http://gd.geobytes.com/AutoCompleteCity?callback=?&q="+request.term,
                function (data) {
                    response(data);
                }
            );
        },
        minLength: 3,
        select: function (event, ui) {
            var selectedObj = ui.item;
            jQuery("#f_elem_city3").val(selectedObj.value);
            return false;
        },
        open: function () {
            jQuery(this).removeClass("ui-corner-all").addClass("ui-corner-top");
        },
        close: function () {
            jQuery(this).removeClass("ui-corner-top").addClass("ui-corner-all");
        }
    });
    jQuery("#f_elem_city3").autocomplete("option", "delay", 100);


    jQuery("#f_elem_city4").autocomplete({
        source: function (request, response) {
            jQuery.getJSON(
                "http://gd.geobytes.com/AutoCompleteCity?callback=?&q="+request.term,
                function (data) {
                    response(data);
                }
            );
        },
        minLength: 3,
        select: function (event, ui) {
            var selectedObj = ui.item;
            jQuery("#f_elem_city4").val(selectedObj.value);
            return false;
        },
        open: function () {
            jQuery(this).removeClass("ui-corner-all").addClass("ui-corner-top");
        },
        close: function () {
            jQuery(this).removeClass("ui-corner-top").addClass("ui-corner-all");
        }
    });
    jQuery("#f_elem_city4").autocomplete("option", "delay", 100);


    jQuery("#f_elem_city5").autocomplete({
        source: function (request, response) {
            jQuery.getJSON(
                "http://gd.geobytes.com/AutoCompleteCity?callback=?&q="+request.term,
                function (data) {
                    response(data);
                }
            );
        },
        minLength: 3,
        select: function (event, ui) {
            var selectedObj = ui.item;
            jQuery("#f_elem_city5").val(selectedObj.value);
            return false;
        },
        open: function () {
            jQuery(this).removeClass("ui-corner-all").addClass("ui-corner-top");
        },
        close: function () {
            jQuery(this).removeClass("ui-corner-top").addClass("ui-corner-all");
        }
    });
    jQuery("#f_elem_city5").autocomplete("option", "delay", 100);
});

$(document).ready(
function()
{
    $('#flyin').effect("slide", {}, 300);
    $('.flyin').effect("shake", {}, 300);
});