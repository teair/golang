
var pageNum = 1;

$(function(){

    // 游戏列表下拉事件
    $('#guipia').on('change',function(){
        $.post('/Index/listen',{'isfirst':'1'});
    })
	
	$("#add").on('click',function(){
    	showLayerLit("创建分发")
    })

     $("#all").on('click',function(){
       // alert("查询全部游戏")
       window.location.href='/Index/index/mark_online/2';
       /*$.ajax({
        url:'/Index/index',
        type:'post',
        data:{'mark_online':'2'},
        datatype:'json',
        success:function(a){
          if (a.code == '200'){
            var tbody = '';
            // 组装数据
            if (a.data.length != 0){
              $.each(a.data,function(k,v){
                if (v['mark_online'] == '1'){
                  tbody+= '<tr>';
                  tbody+= '<td>'+v['gid']+'</td>';
                  tbody+= '<td>'+v['app_name']+'</td>';
                  tbody+= '<td>'+v['gid']+'</td>';
                  tbody+= '<td>'+v['app_rename']+'</td>';
                  tbody+= '<td>'+v['create_time']+'</td>';
                  tbody+= '<td><button type="button" id=' + v['gid'] + ' class="btn btn-outline-success upload" >资料上传</button><button type="button" class="btn btn-outline-warning">修改</button><button type="button" id=' + v['gid'] + ",1" + ' class="btn btn-outline-secondary stateoff">上线</button><button type="button" id=' + v['gid'] + ' class="btn btn-outline-danger del">删除</button></td>';
                  tbody+= '</tr>';
                }else{
                  tbody+= '<tr>';
                  tbody+= '<td>{$vo.id}</td>';
                  tbody+= '<td>{$vo.app_name}</td>';
                  tbody+= '<td>{$vo.gid}</td>';
                  tbody+= '<td>{$vo.app_rename}</td>';
                  tbody+= '<td>{$vo.create_time}</td>';
                  tbody+= '<td><button type="button" id=' + v['gid'] + ' class="btn btn-outline-success upload" >资料上传</button><button type="button" class="btn btn-outline-warning">修改</button><button type="button" id=' + v['gid'] + ",0" + ' class="btn btn-outline-secondary stateoff">下线</button><button type="button" id=' + v['gid'] + ' class="btn btn-outline-danger del">删除</button></td>';
                  tbody+= '</tr>';
                }
              })
            }else{
                tbody+= '<tr>';
                tbody+= '<td colspan="6" align="center">';
                tbody+= '暂无分发游戏!';
                tbody+= '</td>';
                tbody+= '</tr>';
            }
            $('#list').html(tbody);
          }
        },error:function(){

        }
       })*/
       // $(".query-tabs button").removeClass("btn-info").removeClass("btn-outline-secondary")
       // $(".query-tabs button").addClass("btn-outline-secondary")
       // $(this).removeClass("btn-outline-secondary").addClass("btn-info")
     })

     $("#typeon").on('click',function(){
      // 上线
       window.location.href='/Index/index/mark_online/1';
       // $(".query-tabs button").removeClass("btn-info").removeClass("btn-outline-secondary")
       // $(".query-tabs button").addClass("btn-outline-secondary")
       // $(this).removeClass("btn-outline-secondary").addClass("btn-info")
     })

     $("#typeoff").on('click',function(){
       // 下线
       window.location.href='/Index/index/mark_online/0';
       // $(".query-tabs button").removeClass("btn-info").removeClass("btn-outline-secondary")
       // $(".query-tabs button").addClass("btn-outline-secondary")
       // $(this).removeClass("btn-outline-secondary").addClass("btn-info")
     })

     $("#query").on('click',function(){
       // alert("游戏名称查询")
       var app_name = $('#query-name').val();
       // window.location.href='/Index/index/app_name/' + encodeURIComponent(app_name);
       window.location.href='/Index/index/app_name/' + app_name;
     })

     $(".upload").on('click',function(){
        var gid = $(this).attr('id');
        window.location.href="/Upload/index/gid/" + gid;
     })

     
})





