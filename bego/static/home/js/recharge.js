
$(function(){
	$(".pay-list li").click(function(){
		$(".pay-list li a").removeClass("active");
		$(this).find("a").addClass("active");
		var navitem = $(this).data("navitem");
		// console.log(navitem)
		var navitemtext = $(this).find("span").text();
		// console.log(navitemtext)
		$(".order-tit i").removeClass().addClass("o-tit-"+navitem);
		$(".order-tit span").text(navitemtext)
	})

	$("#tr-pay-for .o-div a").click(function(){
		$("#tr-pay-for .o-div a").removeClass('active');
		$(this).addClass('active');
		if($(this).data("payfornum")==2){
			$(".pay-for-tip").show();
		}else{
			$(".pay-for-tip").hide();
		}
	})

	$(".pay-money-list .money-item").click(function(){
		$(".pay-money-list .money-item").removeClass('active');
		$(this).addClass('active');
		var sum = $(this).data("sum");
		var base = $("#coin-radio").text();

		$("#game-coin").text(parseInt(sum)*parseInt(base))

	})

	$('.other-money-box').blur(function () {
        var sum = $(this).val();
        if(sum != '' &&sum != 0){
        	var base = $("#coin-radio").text();

			$("#game-coin").text(parseInt(sum)*parseInt(base))
        }
		
    });
})