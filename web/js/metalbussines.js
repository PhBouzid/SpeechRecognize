/**
 * Created by РђР»РµРєСЃР°РЅРґСЂ on 20.03.2018.
 */

$(window).resize(function(){
    if($(window).width() >= 1024) {
        $('.news').css('margin-bottom', $('.newsAbsBlock').height() - 130);
    } else {
        $('.news').css('margin-bottom', '20px');
    }
});
$(document).ready(function(){
    if($(window).width() >= 1024) {
        $('.news').css('margin-bottom', $('.newsAbsBlock').height() - 130);
    } else {
        $('.news').css('margin-bottom', '20px');
    }
});

$(document).ready(function(){
    var stream = {
            title: "Demo Stream",
            mp3: "http://127.0.0.1:8030/audio"
        },
        ready = false;

    $("#jquery_jplayer_1").jPlayer({
        ready: function (event) {
            ready = true;
            //$(this).jPlayer("setMedia", stream).jPlayer("play");
            $(this).jPlayer("setMedia", stream);
            $('.jp-play').click();
        },
        pause: function() {
            $(this).jPlayer("clearMedia");
        },
        error: function(event) {
            if(ready && event.jPlayer.error.type === $.jPlayer.error.URL_NOT_SET) {
                // Setup the media stream again and play it.
                $(this).jPlayer("setMedia", stream).jPlayer("play");
            }
        },
        swfPath: "js",
        supplied: "mp3",
        preload: "none",
        wmode: "window",
        autoPlay: true,
        keyEnabled: true
    });


    var img_array = [1, 2, 3],
        newIndex = 0,
        index = 0,
        interval = 30000;
    function changeBg() {

        //  --------------------------
        //  For random image rotation:
        //  --------------------------

        //  newIndex = Math.floor(Math.random() * 10) % img_array.length;
        //  index = (newIndex === index) ? newIndex -1 : newIndex;

        //  ------------------------------
        //  For sequential image rotation:
        //  ------------------------------

        index = (index + 1) % img_array.length;
        $("#leftSide").fadeOut(100);
        $("#leftSide").css('background-image', 'url("../images/slide' + img_array[index] + '.jpg');
        $("#leftSide").fadeIn(100);
        setTimeout(changeBg, interval);


    };
    changeBg();


});