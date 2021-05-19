


$(function(){

	$("#add").on('click',function(){
        ifrShowLayerLit("创建标签","tagadd")
    })
    add()
    edit()
})




//添加阶梯
function add(){
    $(".add").on('click',function(){
        ifrShowLayerLit("添加标签","tagadd")
    })
}


//编辑菜单
function edit(){
    $(".edit").on('click',function(){
        var tag = $(this).attr('id');
        tag = tag.split(',');
        ifrShowLayerLit("修改标签","tagedit?id=" + tag[0] + '&pid=' + tag[1] + '&tag_name=' + tag[2])
    })
}



//删除菜单
function del(){
    $(".del").on('click',function(){
        var id = $(this).attr('id');
        layer.confirm('是否确认删除？', {
                btn: ['是', '否'] //可以无限个按钮
            },function(index,layero){
            $.ajax({
                url: '/Tag/tagdel',
                type: 'post',
                dataType: 'json',
                data: {'id': id},
                success: function(a){
                    var a = eval('(' + a + ')');
                    if (a.msg == 'success'){
                        window.location.href='/Tag/index';
                    }else{
                        alert('删除失败!');
                    }
                },
                error:function(c,d,e){
                    console.log(e);
                }
            });
        });
        // if (confirm('确定要删除吗？')){
        //     $.ajax({
        //         url: '/Tag/tagdel',
        //         type: 'post',
        //         dataType: 'json',
        //         data: {'id': id},
        //         success: function(a){
        //             var a = eval('(' + a + ')');
        //             if (a.msg == 'success'){
        //                 window.location.href='/Tag/index';
        //             }else{
        //                 alert('删除失败!');
        //             }
        //         },
        //         error:function(c,d,e){
        //             console.log(e);
        //         }
        //     })
        // }
        /*layer.confirm('是否确认删除？', {
            btn: ['是', '否'] //可以无限个按钮
        },function(index,layero){
            
        }*/
    })
}

