$(function(){


    //日期选择
    datecur = new Date();
    datecur = formatDate(datecur);
    timecur = new Date();
    timecur = formatTime(timecur);
    //console.log(timecur);

    $(".datetime2").val(datecur + ' ' + timecur);
    $(".datepicker").val(datecur);

    $(".datetime2").datetimepicker({
        format:'Y-m-d H:i',
        formatDate:'Y-m-d H:i',
    }) 

    $(".datepicker").datetimepicker({
        lang:'ch',
        timepicker:false,
        format:'Y-m-d',
        formatDate:'Y-m-d',
       
    }) 

    //时间选择
    $(".time").datetimepicker({
        datepicker:false,
        format:'H:i',
        // allowTimes:['12:00','13:00','15:00','17:00','17:05','17:20','19:00','20:00'],
        step:30
    }) 
    
})