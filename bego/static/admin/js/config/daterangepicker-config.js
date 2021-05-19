$(function(){
    
    //时间段选择器
    $(".dateclear").on('click',function(){
        $(this).parent().find('input').val("");
    })
    $(".aud-date i").click(function(){
        $(this).parent().find('input').trigger("click");
    })
    $('.daterange').daterangepicker({
        "singleDatePicker": false,
        "autoApply": true,
        "locale" : {  
            format: "YYYY-MM-DD",   
            separator: ' ~ ', 
            applyLabel : '确定',  
            cancelLabel : '取消',  
            fromLabel : '起始时间',  
            toLabel : '结束时间',  
            minDate : '2019/11/30',
            customRangeLabel : '自定义',  
            daysOfWeek : [ '日', '一', '二', '三', '四', '五', '六' ],  
            monthNames : [ '一月', '二月', '三月', '四月', '五月', '六月', '七月', '八月', '九月', '十月', '十一月', '十二月' ],  
            firstDay : 1  
            }  
    })  

})