

$(function(){
    //显示区域设置高度
    var winH = $(window).height();
    // console.log(winH)
    var showH = winH - 220 ;
    // console.log(showH)
    $(".tab-content").css('height',showH)

    //当浏览器大小变化时
    /*$(window).resize(function () {          
        var winH = $(window).height();
        // console.log(winH)
        var showH = winH - 230 ;
        // console.log(showH)
        $(".tab-content").css('height',showH)
    });*/

    //判断是否登录
    if(!token) {
        alert("找不到token，请重新登录！");
        location.href="login.html";
    }else{
        $.ajax({
            type:'post',
            url:PHPLINK + "admin/v1.check/checkToken",
            data:{},
            // headers:{'Content-Type':'application/x-www-form-urlencoded','token':token},
            success:function(data){
                console.log(data);

                if(data.code == 200){
                    console.log("token验证成功！！")
                    //网站头部管理员ID
                    $("#headname").text(username);
                    setToken();
                    bellRemind();

                    // 获取导航
                    getNav();
                    console.log(navArr);
                    if(navArr == ''){
                        alert("该账号暂未分配角色，请联系管理员.......")
                    }
                    getTopNav(localStorage.getItem('navname1'),true);
                }else{
                    alert("token已过期，请重新登录");
                    location.href="login.html";
                }
  
            },
            error:function(err){
                console.log(err)
                alert("token已过期，请重新登录");
                location.href="login.html";
            }
        });
        
    }

    

    //退出登录
    $("#signout").click(function(){
        alert('退出登录！');
        localStorage.clear();
        location.href = "login.html";
    })

    //选项卡初始化
    $('#main').bTabs({
        resize : function(){ 
            var winH = $(window).height();
            // console.log(winH)
            var showH = winH - 220 ;
            // console.log(showH)
            $(".tab-content").css('height',showH)
        }
    });


	// console.log(navSub);
	// console.log(navSub.id);
	// console.log(navSub.name[0]);

    // console.log(localStorage.getItem('navname1'))
	// getTopNav(localStorage.getItem('navname1'),true);
	

    //菜单点击
    /*$('a',$('#menuSideBar')).on('click', function(e) {
        e.stopPropagation();
        var li = $(this).closest('li');
        var menuId = $(li).attr('mid');
        var url = $(li).attr('funurl');
        var title = $(this).text();
        $('#mainFrameTabs').bTabsAdd(menuId,title,url);
        
    });*/




    //消息列表窗口
    $("#message").click(function(){
        dataShowLayer("消息列表","/message.html")
        // $("#message").find("badge").css("display","none")
    })
    


})


//定时查询token
function setToken(){
    setTimeout(function(){
        isToken();
    },600000);
}

function isToken(){
    //判断是否登录
    if(!token) {
        alert("找不到token，请重新登录！");
        location.href="login.html";
    }else{
        $.ajax({
            type:'post',
            url:PHPLINK + "admin/v1.check/checkToken",
            data:{},
            headers:{'Content-Type':'application/x-www-form-urlencoded','token':token},
            success:function(data){
                console.log(data);

                if(data.code == 200){
                    console.log("token验证成功！！")
                    //网站头部管理员ID
                    $("#headname").text(username);
                    setToken();
                }else{
                    alert("token已过期，请重新登录");
                    location.href="login.html";
                }
  
            },
            error:function(err){
                console.log(err)
                alert("token已过期，请重新登录");
                location.href="login.html";
            }
        });
        
    }
}

function bellRemind(){
    /**
     * 与GatewayWorker建立websocket连接，域名和端口改为你实际的域名端口，
     * 其中端口为Gateway端口，即start_gateway.php指定的端口。
     * start_gateway.php 中需要指定websocket协议，像这样
     * $gateway = new Gateway(websocket://0.0.0.0:7272);
     */
    ws = new WebSocket("ws://192.168.8.115:8282");
    // ws = new WebSocket("ws://zixishi-ws.fangyoucn.com");
    // 服务端主动推送消息时会触发这里的onmessage
    ws.onmessage = function(e){
        console.log(e)
        // json数据转换成js对象
        var data = eval("("+e.data+")");
        var type = data.type || '';
        switch(type){
            // Events.php中返回的init类型的消息，将client_id发给后台进行uid绑定
            case 'init':
                // 利用jquery发起ajax请求，将client_id发给后端进行uid绑定
                // $.post('/admin/v1.worker/index', {client_id: data.client_id}, function(data){
                //     console.log(data);
                // }, 'json');  //8f3131ab70ee27f58b2ed2b93df798e4  93895ac3d26381e9fb6b4d87ab2146e6
                $.ajax({
                    type:'post',
                    url: PHPLINK + 'admin/v1.worker/index',
                    data:{'client_id': data.client_id},
                    headers:{'Content-Type':'application/x-www-form-urlencoded','token':token},
                    success:function(data){
                        console.log(data);
                    },
                    error:function(err){
                        console.log(err)
                    }
                });
                break;
            case 'message':
                console.log(data.msg);
                dataShowLayer2('新消息提醒',"/new_message.html?content="+data.msg);
                // $("#message").find("badge").css("display","block")
                break;    
            // 当mvc框架调用GatewayClient发消息时直接alert出来
            default :
                console.log(data.msg);
        }
    };

    ws.onopen = function(){
        console.info("与服务端连接成功");
        ws.send('test msg\n');//相当于发送一个初始化信息
        console.info("向服务端发送心跳包字符串");
        setInterval(show,8000);
    }

    function show(){
        ws.send('heart beat\n');
    }
}
    


//iframe监听
function ifrchange(){
    // console.log($("#iniframe").attr("src"))
    
}

// function backout(){
//     document.frames['iframe'].history.back();
//     return false;
// }


//导航跳转
function navTopSkip(){

   $(".topnav-item").on('click',function(e){
		var thisname = $(this).find('a').text();
		// console.log(thistext)
		getTopNav(thisname,false);

        localStorage.setItem('navname1',thisname);
        // console.log(localStorage.getItem('navname1'))
        // console.log(localStorage)

	})     
}

function navSideSkip(){
    $(".nav3-link-active").parents(".nav-dropdown").addClass("open");
    $(".nav3-link").on('click',function(){
		var thisurl = $(this).data('url');
		var thisupname = $(this).data('up');;
		var thisname = $(this).find('a').text();
		
        localStorage.setItem('navname2',thisupname);
        localStorage.setItem('navname3',thisname);

		// navSub.name[1] = thisupname;
		// navSub.name[2] = thisname;

		// console.log(localStorage)
        // console.log(localStorage.getItem('navname3'))
		//跳转页面
		// console.log(thisurl)
        
		getSideNav(localStorage.getItem('navid'),localStorage.getItem('navname3'));

		
	})     
}

function changeIfr(e){
    // console.log(e)
    var url = $("#iniframe").attr("src");
    // history.replaceState(null,null,'./'+url+'.html');
}



//顶部菜单导航获取
function getTopNav(text,initial) {  
    // console.log(text) 
    // console.log(navArr)
    // console.log(initial)

    var topNavHtml = '';
            
    for(var i=0;i<navArr.length;i++){

        if(navArr[i].title == text){
        	// navSub.id = i; 
            localStorage.setItem('navid',i);

        	//当前选中导航
            if(navArr[i].url == "#" ){
            	// console.log("当前有二级导航");
                topNavHtml += '<li class="topnav-item active"> <a>'+ navArr[i].title +'</a> </li>';
                //添加数据
                // navSub.name[0] = navArr[i].title;
                // navSub.name[1] = navArr[i].sub[0].title;
                // navSub.name[2] = navArr[i].sub[0].sub[0].title;

                // localStorage.setItem('navname1',navArr[i].title);
                // localStorage.setItem('navname2',navArr[i].sub[0].title);
                // localStorage.setItem('navname3',navArr[i].sub[0].sub[0].title);
				// console.log(localStorage.getItem('navid'))
                // console.log(localStorage.getItem('navname3'))
                // 获取二级导航
                if(!initial){
                    if(navArr[i].title){
                        localStorage.setItem('navname1',navArr[i].title);
                    }
                    if(navArr[i].sub[0].title){
                        localStorage.setItem('navname2',navArr[i].sub[0].title);
                    }
                    if(navArr[i].sub[0].sub[0].title){
                        localStorage.setItem('navname3',navArr[i].sub[0].sub[0].title);
                    }
  
                }

                // console.log(tabArr)
                // $('#main').bTabsClose('bTabs_tab1');
                for(var t=0;t<tabArr.length;t++){
                    // console.log(tabArr[t])
                    $('#main').bTabsClose('bTabs_'+tabArr[t]);
                }
                tabArr = []
                getSideNav(localStorage.getItem('navid'), localStorage.getItem('navname3'));
                // 显示sidebar
		        $("#sidebar-btn").show();
		        $("body").removeClass("sidebar-hidden");
            }else{
            	// console.log("当前没有二级导航");
                topNavHtml += '<li class="topnav-item active"> <a>'+ navArr[i].title +'</a> </li>';
                //添加数据
                // navSub.name[0] = navArr[i].title;
                // navSub.name[1] = '';
                // navSub.name[2] = '';

                localStorage.setItem('navname1',navArr[i].title);
                localStorage.setItem('navname2','');
                localStorage.setItem('navname3','');
				// console.log(navSub)
                // 隐藏sidebar
		        $("#sidebar-btn").hide();
		        $("body").addClass("sidebar-hidden");
		        //跳转页面
                // console.log('tab'+navArr[i].id);
		        // console.log(navArr[i].title);
                // console.log(navArr[i].myalias);
                // console.log(tabArr)
                // $('#main').bTabsClose('bTabs_tab1');
                for(var t=0;t<tabArr.length;t++){
                    // console.log(tabArr[t])
                    $('#main').bTabsClose('bTabs_'+tabArr[t]);
                }
                tabArr = []
                // console.log(navArr[i])
                tabArr.push('tab'+navArr[i].id)
                // console.log(tabArr)
                $('#main').bTabsAdd('tab'+navArr[i].id,navArr[i].title,navArr[i].myalias);
                // $("#iniframe").attr("src","./" + navArr[i].myalias + ".html");
                // localStorage.setItem('url',navArr[i].myalias);
                // history.replaceState(null,null,'./'+navArr[i].myalias+'.html');

            }

        }else{
        	//其他导航
            topNavHtml += '<li class="topnav-item"> <a>'+ navArr[i].title +'</a> </li>'; 
        }  
    }

    // 顶部导航DOM加载
	$("#top-navbar").html(topNavHtml);

	//顶部导航跳转
	navTopSkip();
	
        
};

//侧菜单导航获取
function getSideNav(num,text) {  
	// console.log(num) 
	// console.log(text) 
 	// console.log(navArr) 
    var sideNav3Url = '';

    var sideNavHtml = '';
    var sideNavTitleHtml = '<li class="nav-title"></li>';
    var sideNav2Html = '';
    var sideNav2HtmlHead= '<li class="nav-item nav-dropdown">';
    var sideNav2HtmlFoot= '</li>';            
    var sideNav3Html = '';
    var sideNav3HtmlHead = '<ul class="nav-dropdown-items">';
    var sideNav3HtmlFoot = '</ul>';

    for(var j=0;j<navArr[num].sub.length;j++){
        // console.log(j)
        // console.log(navArr[navNum].sub[j].title)

        // sideNavHtml = '';
        sideNav2Html = '';
        sideNav3Html = '';
        // sideNav2HtmlHead='';
        // sideNav2HtmlFoot='';
        // sideNav3Html = '';
        // sideNav3HtmlHead='';
        // sideNav3HtmlFoot='';

        sideNav2Html += '<a href="#" class="nav-link nav-dropdown-toggle"><i class="icon"></i> '+ navArr[num].sub[j].title +'<i class="fa fa-caret-left"></i></a>'
        
        // console.log(navArr[num].sub[j].sub)
        if( !navArr[num].sub[j].sub ){
            sideNav3Html = '';
        }else{
            for(var k=0;k<navArr[num].sub[j].sub.length;k++){
               
                // console.log(navArr[num].sub[j].sub[k].icon);
                
                if(navArr[num].sub[j].sub[k].title == text){
                    sideNav3Html += '<li class="nav-item nav3-link" data-url="'+
                                    navArr[num].sub[j].sub[k].myalias +'" data-up="'+
                                    navArr[num].sub[j].title +'" ><a href="#" class="nav-link nav3-link-active active"><i class="'+ 
                                    navArr[num].sub[j].sub[k].icon + '"></i>'+ 
                                    navArr[num].sub[j].sub[k].title + '</a></li>' ;  
                    //跳转页面
                    // $("#iniframe").attr("src","./" + navArr[num].sub[j].sub[k].myalias + ".html") ;
                    tabArr.push('tab'+navArr[num].sub[j].sub[k].id)
                    $('#main').bTabsAdd('tab'+navArr[num].sub[j].sub[k].id,navArr[num].sub[j].sub[k].title,navArr[num].sub[j].sub[k].myalias);
                    // localStorage.setItem('url',navArr[num].sub[j].sub[k].myalias);
                    // history.replaceState(null,null,'./'+navArr[i].myalias+'.html');                 
                    
                }else{
                    sideNav3Html += '<li class="nav-item nav3-link" data-url="'+
                                    navArr[num].sub[j].sub[k].myalias +'" data-up="'+
                                    navArr[num].sub[j].title +'" ><a href="#" class="nav-link"><i class="'+
                                    navArr[num].sub[j].sub[k].icon + '"></i>' +
                                    navArr[num].sub[j].sub[k].title + '</a></li>' ;
                }

            }  
        }
            
        
         

        sideNavHtml += sideNav2HtmlHead + sideNav2Html + sideNav3HtmlHead + sideNav3Html + sideNav3HtmlFoot + sideNav2HtmlFoot;

        // console.log(sideNavHtml)

    }

    $("#sidebar-nav").html(sideNavHtml);


    navSideSkip();

    $('.nav-dropdown-toggle').on('click', function (e) {
        e.preventDefault();
        $(this).parent().toggleClass('open');
    });


};

//数据中心模块弹窗处理
//弹窗iframe
function dataShowLayer(title,url){
    layer.open({
        type: 2,
        title: [title,'font-size:16px; color:#333;'],  //标题文本 , 样式 ， 不显示标题 title:false
        area:['66%','500px'],   //定义宽高
        shade: [0.8, '#393D49'],  //遮罩层 【透明度 ， 颜色】
        shadeClose: true,     //点击遮罩层关闭弹窗
        //time: 5000,   //自动关闭
        closeBtn: 1,  //关闭按钮
        anim: 0,  //弹出动画
        isOutAnim: true,  //关闭动画
        maxmin: true,   //最大最小化
        fixed: false,  //固定 - 鼠标滚动时，层是否固定在可视区域
        resize:true,  //是否允许拉伸
        scrollbar: false,   //是否允许浏览器出现滚动条
        move: '.layui-layer-title',   //触发拖动的元素
        moveOut: false,    //是否允许拖拽到窗口外
        zIndex:999,   //层叠顺序
        content: [PORTLINK + url, 'no'], //iframe的url，no代表不显示滚动条
     
    });
}

function dataShowLayer2(title,url){
    layer.open({
        type: 2,
        title: [title,'font-size:16px; color:#333;'],  //标题文本 , 样式 ， 不显示标题 title:false
        area:['600px','400px'],   //定义宽高
        shade: [0.8, '#393D49'],  //遮罩层 【透明度 ， 颜色】
        shadeClose: true,     //点击遮罩层关闭弹窗
        //time: 5000,   //自动关闭
        closeBtn: 1,  //关闭按钮
        anim: 0,  //弹出动画
        isOutAnim: true,  //关闭动画
        maxmin: true,   //最大最小化
        fixed: false,  //固定 - 鼠标滚动时，层是否固定在可视区域
        resize:true,  //是否允许拉伸
        scrollbar: false,   //是否允许浏览器出现滚动条
        move: '.layui-layer-title',   //触发拖动的元素
        moveOut: false,    //是否允许拖拽到窗口外
        zIndex:999,   //层叠顺序
        content: [PORTLINK + url, 'no'], //iframe的url，no代表不显示滚动条
     
    });
}



function dataHideLayer(){
    layer.close(layer.index)
}


