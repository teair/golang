
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
       window.location.href='/admin/index/2';
     })

     $("#typeon").on('click',function(){
        // 上线
        window.location.href='/admin/index/1';
     })

     $("#typeoff").on('click',function(){
        // 下线
        window.location.href='/admin/index/0';
     })

     $("#query").on('click',function(){
       // alert("游戏名称查询")
       var app_name = $('#query-name').val();
       // window.location.href='/Index/index/app_name/' + encodeURIComponent(app_name);
       window.location.href='/admin/index//' + app_name;
     })

     $(".upload").on('click',function(){
        var gid = $(this).attr('id');
        window.location.href="/Upload/index/gid/" + gid;
     })

     
})





