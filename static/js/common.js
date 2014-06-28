$(function() {
	$('.nav-container').slideDown(500);
	$(".nav-list").on("mouseenter", "li.nav-item", function(evt) {
		getFocus($(this));
	});

	$(".nav-list").on("mouseout", function(evt) {
		var $dom = $(".nav-list").find(".cur");
		getFocus($dom);
	});
});

function getFocus($dom,flag){
	var _top = ($dom.index()) * $dom.outerHeight()+130;
	$(".hover-layer").css('top', _top);
	$(".nav-list>li.active").removeClass("active");
	$dom.addClass('active');
	if(flag){
		$dom.addClass("cur");
	}
}