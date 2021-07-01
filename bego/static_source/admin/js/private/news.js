    // var id ;
    $(function(){

        // id = getPar("id");

        $("#add").click(function(){
            window.location.href="/Gamenews/newsadd"
	    })

	    $(".edit").click(function(){
            var id = $(this).attr('id');
	        window.location.href="/Gamenews/newsedit/id/" + id;
	    })

        $('.newsdel').click(function(){
            var id = $(this).attr('id');
            layer.confirm('是否确认删除？', {
                btn: ['是', '否'] //可以无限个按钮
            },function(index,layero){
                $.ajax({
                    url:'/Gamenews/newsdel',
                    type:'post',
                    datatype:'json',
                    data:{'id':id},
                    success:function(data){
                        if (data.msg == 'success'){
                            window.location.href='/Gamenews/index';
                        }else{
                            alert('删除失败!');
                        }
                    },
                    error:function(a,b,c){
                        // alert(c);
                        alert('意外错误，请联系管理员!');
                    }
                })
            })
        })

    })