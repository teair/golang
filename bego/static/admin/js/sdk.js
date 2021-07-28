
$(function(){

    //显示区域设置高度
    var winH = $(window).height();
    // console.log(winH)
    var showH = winH - 100 ;
    // console.log(showH)
    $(".content").css({'height':showH,'overflow-y':'scroll'})

    //当浏览器大小变化时
    $(window).resize(function () {          
        var winH = $(window).height();
        // console.log(winH)
        var showH = winH - 100 ;
        // console.log(showH)
        $(".content").css({'height':showH,'overflow-y':'scroll'})
    });
  
    del();
    state();

})

//上线-下线
function state(){
    $(".stateoff").on('click',function(e){
        var str = $(this).attr('data');
        var arr = new Array()   // 定义一组数组
        arr = str.split(",")
        $.ajax({
            type:'post',
            url:'/admin/markonline',
            datatype:'json',
            data:{gid:arr[0],mark_online:arr[1],id:arr[2]},
            // contentType:"application/x-www-form-urlencoded",
            // headers:{'Content-Type':'application/x-www-form-urlencoded','token':token},
            success:function(a){
                if (a.code == 200){
                    window.location.href='/admin/index?mark_online=' + a.data;
                    return false;
                }
                if (a.code == 1000){
                    layer.open({
                      content: '请上传游戏资料!'
                      ,title:'提示'
                      ,btn: ['确认']
                      ,yes: function(index, layero){
                        //按钮【按钮一】的回调
                        window.location.href="/admin/upindex/" + a.data;
                      }
                      ,cancel: function(){ 
                        //右上角关闭回调
                        
                        // return false //开启该代码可禁止点击该按钮关闭
                      }
                    });
                }else{
                    layer.msg('系统繁忙...',{icon:5});
                }
            },
            error:function(){
                layer.msg('意外错误，请联系管理员...',{icon:5});
            }
        });
    })

    $(".stateon").on('click',function(){
          alert("解禁成功");
    })
}


//删除菜单
function del(){
    // 已分发游戏的删除
    $(".del").on('click',function(){
        var gid = $(this).attr('id');
        layer.confirm('是否确认删除？', {
            btn: ['是', '否'] //可以无限个按钮
        },function(index,layero){
            $.ajax({
                type:'post',
                url:'/admin/gamedel',
                datatype:'json',
                data:{gid:gid},
                success:function(a){
                    if (a.Code == '200'){
                        // layer.alert('删除成功!',{icon:1});
                        window.location.href='../../../admin/index';
                    }else{
                        layer.alert('系统繁忙...',{icon:5});
                    }
                },
                error:function(){
                    layer.alert('意外错误，请联系管理员...',{icon:5});
                }
            });
        });
    })
}


// 游戏区服删除
$(".serverdel").on('click',function(){
    var id = $(this).attr('id');
    // layer.alert(id);
    layer.confirm('是否确认删除？', {
            btn: ['是', '否'] //可以无限个按钮
        },function(index,layero){
        $.ajax({
            type:'post',
            url:'/Server/serverdel',
            datatype:'json',
            data:{id:id},
            // contentType:"application/x-www-form-urlencoded",
            success:function(data){
                var data = eval('(' + data + ')');
                if (data.msg == 'success'){
                    // layer.alert('删除成功!',{icon:1});
                    window.location.href='/Server/index';
                }else{
                    layer.alert('系统繁忙...',{icon:5});
                }
            },
            error:function(){
                layer.alert('意外错误，请联系管理员...',{icon:5});
            }
        });
    });
})


// 友情链接删除
$(".linkdel").on('click',function(){
    if (confirm('是否确认删除?')){
        var id = $(this).attr('id');
        $.ajax({
            type:'post',
            url:'/Links/linkdel',
            datatype:'json',
            data:{id:id},
            // contentType:"application/x-www-form-urlencoded",
            success:function(data){
                var data = eval('(' + data + ')');
                if (data.msg == 'success'){
                    window.location.href='/Links/index';
                }else{
                    layer.alert('系统繁忙...',{icon:5});
                }
            },
            error:function(a,b,c){
                layer.alert('意外错误，请联系管理员...',{icon:5});
            }
        });
    }
})

/**********************************************************************公用方法---  */
//导航高亮
function navText(text1,text2){
    // console.log(text1)
    // console.log(text2)
    $(".sidebar-nav .nav-link-one").parent().removeClass("open");
    $(".sidebar-nav .nav-link-two").removeClass("active");

    $(".sidebar-nav .nav-link-one").each(function(key,value){
        var thisText1=$(this).find('span').text();
        // console.log(thisText1)
        if(text1==thisText1){
            // console.log(thisText1)
            $(this).parent().addClass("open");
        }
    })
    $(".sidebar-nav .nav-link-two").each(function(key,value){
        var thisText2=$(this).find('span').text();  
        // console.log(thisText2)   
        if(text2==thisText2){
            // console.log(thisText2)
            $(this).addClass("active");
        }
    })
    
}


//字符串去空格
function noSpaceText(text) {  
    // console.log(text)
    
    return text.replace(/(^\s*)|(\s*$)/g, "");

};


//时间格式化方法  yyyy-mm-dd hh:ii
function formatDate(date) {  
    var y = date.getFullYear();  
    var m = date.getMonth() + 1;  
    m = m < 10 ? ('0' + m) : m;  
    var d = date.getDate();  
    d = d < 10 ? ('0' + d) : d;  
    var h = date.getHours();  
    h=h < 10 ? ('0' + h) : h;  
    var minute = date.getMinutes();  
    minute = minute < 10 ? ('0' + minute) : minute;  
    var second=date.getSeconds();  
    second=second < 10 ? ('0' + second) : second;  
    //return y + '-' + m + '-' + d+' '+h+':'+minute+':'+second; 
    return y + '-' + m + '-' + d;  
}; 

function formatTime(date) {  
    var y = date.getFullYear();  
    var m = date.getMonth() + 1;  
    m = m < 10 ? ('0' + m) : m;  
    var d = date.getDate();  
    d = d < 10 ? ('0' + d) : d;  
    var h = date.getHours();  
    h=h < 10 ? ('0' + h) : h;  
    var minute = date.getMinutes();  
    minute = minute < 10 ? ('0' + minute) : minute;  
    var second=date.getSeconds();  
    second=second < 10 ? ('0' + second) : second;  
    //return y + '-' + m + '-' + d+' '+h+':'+minute+':'+second; 
    return h + ':' + minute + ':' + second;  
}; 

//获取URL
function getPar(par){
    //获取当前URL
    var local_url = document.location.href; 
    //获取要取得的get参数位置
    var get = local_url.indexOf(par +"=");

    if(get == -1){
        return false;   
    }   
    //截取字符串
    var get_par = local_url.slice(par.length + get + 1);  
  
    //判断截取后的字符串是否还有其他get参数
    var nextPar = get_par.indexOf("&");


    if(nextPar != -1){
        get_par = get_par.slice(0, nextPar);
    }
    return decodeURI(get_par);
}

//时间字符串转换时间戳
function getUnix(date){
    var obdate = new Date(Date.parse(date.replace(/-/g, "/")));
        obdate = obdate.getTime() / 1000;
        return obdate;
}   

//时间补0
function getNow(s) {
    return s < 10 ? '0' + s: s;
}
//获取当前时间
function getCurDateTime(t){
    var myDate = new Date();

    var year=myDate.getFullYear();   //获取当前年
    var month=myDate.getMonth()+1;   //获取当前月
    var date=myDate.getDate();       //获取当前日

    var h=myDate.getHours();         //获取当前小时数(0-23)
    var m=myDate.getMinutes();       //获取当前分钟数(0-59)
    var s=myDate.getSeconds(); 

    // var nowtime = year+'-'+getNow(month)+"-"+getNow(date)+" "+getNow(h)+':'+getNow(m)+":"+getNow(s);

    // console.log(nowtime)

    var nowtime ;

    switch(t){
        case 'year':
        nowtime = year
        break;
        case 'month':
        nowtime = getNow(month)
        break;
        case 'date':
        nowtime = getNow(date)
        break;
        case 'h':
        nowtime = getNow(h)
        break;
        case 'm':
        nowtime = getNow(m)
        break;
        case 's':
        nowtime = getNow(s)
        break;
        default:
        nowtime = year + '-' + getNow(month) + '-' + getNow(date) +' '+ getNow(h) + ':' + getNow(m) + ':' + getNow(s) ;
        break;
    }

    return nowtime ;
}
//时间戳转化为倒计时时间
function cTime(intDiff){
    var day=0,
        hour=0,
        minute=0,
        second=0;//时间默认值        
    if(intDiff > 0){
        day = Math.floor(intDiff / (60 * 60 * 24));
        hour = Math.floor(intDiff / (60 * 60)) - (day * 24);
        minute = Math.floor(intDiff / 60) - (day * 24 * 60) - (hour * 60);
        second = Math.floor(intDiff) - (day * 24 * 60 * 60) - (hour * 60 * 60) - (minute * 60);
    }
    if (minute <= 9) minute = '0' + minute;
    if (second <= 9) second = '0' + second;
    /*$('#day_show').html(day+"天");*/
    // $('#hour_show').html('<s id="h"></s>'+hour);
    // $('#minute_show').html('<s></s>'+minute);
    // $('#second_show').html('<s></s>'+second);
    if(day == 0){
        return hour + '时'+ minute + '分' + second + '秒'
    }else if(hour == 0){
        return minute + '分' + second + '秒'
    }else{
        return day + '天' + hour + '时'+ minute + '分' + second + '秒' 
    }
    

} 

//倒计时
function timer(intDiff){
    window.setInterval(function(){
    var day=0,
        hour=0,
        minute=0,
        second=0;//时间默认值        
    if(intDiff > 0){
        day = Math.floor(intDiff / (60 * 60 * 24));
        hour = Math.floor(intDiff / (60 * 60)) - (day * 24);
        minute = Math.floor(intDiff / 60) - (day * 24 * 60) - (hour * 60);
        second = Math.floor(intDiff) - (day * 24 * 60 * 60) - (hour * 60 * 60) - (minute * 60);
    }
    if (minute <= 9) minute = '0' + minute;
    if (second <= 9) second = '0' + second;
    /*$('#day_show').html(day+"天");*/
    // $('#hour_show').html('<s id="h"></s>'+hour);
    // $('#minute_show').html('<s></s>'+minute);
    // $('#second_show').html('<s></s>'+second);

    console.log(day + '天' + hour + '时'+ minute + '分' + second + '秒') 

    intDiff--;
    }, 1000);
} 


//普通字符转换成转意符
function escape2Html(str) {

 var arrEntities={'lt':'<','gt':'>','nbsp':' ','amp':'&','quot':'"'};

 return str.replace(/&(lt|gt|nbsp|amp|quot);/ig,function(all,t){return arrEntities[t];});

}

//转意符换成普通字符
function escape2Html(str) {

 var arrEntities={'lt':'<','gt':'>','nbsp':' ','amp':'&','quot':'"'};

 return str.replace(/&(lt|gt|nbsp|amp|quot);/ig,function(all,t){return arrEntities[t];});

}


//刷新
function refresh(){
    window.location.reload()
}




