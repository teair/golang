
var pageNum = 1;

$(function(){

    // 游戏列表下拉事件
    $('#guipia').on('change',function(){
        $.post('/admin/listen',{'isfirst':'1'});
    })
	
	$("#add").on('click',function(){
    	showLayerLit("创建分发")
    })

     $("#all").on('click',function(){
       // alert("查询全部游戏")
       window.location.href='../../../../admin/index?mark_online=2';
     })

     $("#typeon").on('click',function(){
        var data = $(this).data()
        window.location.href='/admin/index?mark_online=1';
     })

     $("#typeoff").on('click',function(){
        // 下线
        var data = $(this).data()
        window.location.href='/admin/index?mark_online=0';
     })

     $("#query").on('click',function(){
       var app_name = $('#query-name').val();
       window.location.href='/admin/index?app_name=' + app_name;
     })

     $(".upload").on('click',function(){
        var gid = $(this).attr('id');
        window.location.href="/admin/upindex/" + gid;
     })

})





