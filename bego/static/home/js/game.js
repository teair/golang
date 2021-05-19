
// 搜索下拉效果执行代码
window.onload = ListenerSearch();

//实时监控搜索框文本输入事件
function ListenerSearch(){
    
    var name = [];
    // 游戏代号
    $.ajax({
        url:'/Game/gamesel',
        datatype:'json',
        type:'post',
        success:function(a,b){
            if (a.code=='200'){
                $.each(a.data,function(index, el) {
                    name.push(el.app_name);
                    name.push(el.app_rename);
                });
            }
        }
    });

    // console.log(name);return false;
    //实时监控文本输入
    $("#search_val").bind('input propertychange',function () {
        // alert($(this).val());
        // return false;
        QueryKeyword($(this).val(),name);
    });
}

//检索数组里是否有对应关键词
function QueryKeyword(SearchName,ArrayList) {
    //初始化数组
    var Keyword = [];
    //遍历数组内容
    for(var i = 0; i < ArrayList.length; i++){
        //查询判断数组内是否包含关键词
        if(SearchName.length != 0){
            //搜索框输入数据全等于数组内字符串参数
            if (ArrayList[i].indexOf(SearchName) >= 0){
                //添加数据
                Keyword.push(ArrayList[i]);
            }
        }
    }
    if(Keyword.length != 0){
    	// console.log('1');return false;
        //创建table表单
        CreateTable(Keyword);
    } else {
        //删除table表单
        RemoveTable();
    }
}

//监控table表单点击事件,修改input内容
function TableOnclick() {
   
    $(".choose").on('click',function(){
        var id = $(this).data("id");
        $("#search_val").val($("#choose"+id).html());
        $("#searchTips").remove();
    })
}

//创建table表单
function CreateTable(Keyword) {
    var TableContent = "";
    for(var i = 0; i < Keyword.length; i++){
        TableContent += "" +
            "<tr>" +
            "<td id='choose"+i+"' data-id='"+i+"' class='choose' onclick=selgames('"+Keyword[i]+"')>"+Keyword[i]+"</td>" +
            "</tr>";
    }

    // console.log(TableContent);return false;
    //table表单不存在时进行创建
    if(!document.getElementById("#searchTips")){
        // console.log(TableContent);return false;
        var Table = document.createElement('table');
        Table.id = "searchTips";
        $(".search_item").append(Table);
    }

    $("#searchTips").html(TableContent);

    // 控制键盘事件
    var num = 0;
    var selength = Keyword.length;
    $('#choose'+num).attr('bgcolor','#888888');
	document.onkeydown=function(event){
		var e = event || window.event || arguments.callee.caller.arguments[0];
		// 下
		if(e && e.keyCode==40){
			$('#choose'+num).attr('bgcolor','#DDDDDD');
			num++;
			if (num > (selength-1)){
				num = 0;
			}
			$('#choose'+num).attr('bgcolor','#888888');
		}
		
		// 上
		if(e && e.keyCode==38){
			$('#choose'+num).attr('bgcolor','#DDDDDD');
			num--
			if (num < 0 ){
				num = (selength-1);
			}
			$('#choose'+num).attr('bgcolor','#888888');
		}

		// 兼容FF和IE和Opera
	　　var theEvent = event || window.event;
	　　var code = theEvent.keyCode || theEvent.which || theEvent.charCode;
	　　if (code == 13) {
			// 回车执行查询
	        $("#search_val").val($("#choose"+num).html());
	        $("#searchTips").remove();
	        selgame();
	　　}
		
	};

    TableOnclick()
}

//删除table表单
function RemoveTable() {
    $("#searchTips").remove();
}


// 文本框搜索游戏
function selgame()
{
	var selname = $('#search_val').val();
	$.ajax({
		url:'/Game/gamelist',
		data:{'gamename':selname},
		datatype:'json',
		type:'get',
		success:function(a,b){
			datares(a);
		}
	})
}
// 文本框搜索游戏
function selgames(gamename)
{
	$.ajax({
		url:'/Game/gamelist',
		data:{'gamename':gamename},
		datatype:'json',
		type:'get',
		success:function(a,b){
			datares(a);
		}
	})
}




$(function(){
	$(".tabs_label a").click(function(){
		$(".tabs_label a").removeClass('active');
		$(this).addClass('active');
	})

	$(".tools .tab .item").click(function(){
		$(".tools .tab .item").removeClass('active');
		$(this).addClass('active');

		if($(this).data("tab")==2){
			$(".game_list").addClass("picture")
		}else if($(this).data("tab")==1){
			$(".game_list").removeClass("picture")
		}
	})


	$.ajax({
		url:'/Index/serverinfo',
		type:'get',
		data:{'time':'today'},
		datatype:'json',
		success:function(a,b){
			// 区服数据处理
			serveres(a);
		}
	})
	$(".game-open-tabs dd").click(function(){
		$(".game-open-tabs dd").removeClass('active');
		$(this).addClass('active');

		// 游戏开服
		var data = {};
		var txt = $(this).text();
		if (txt == '今日开服'){
			data.time = 'today'
		}
		if (txt == '明日开服'){
			data.time = 'tomorrow';
		}
		if (txt == '昨日开服'){
			data.time = 'yesterday';
		}
		$.ajax({
			url:'/Index/serverinfo',
			type:'get',
			data:data,
			datatype:'json',
			success:function(a,b){
				// 区服数据处理
				serveres(a);
			}
		})
	})
})

// 开服分页
function getserver(data){
	var page = data.split('=')[1];
	$.ajax({
		url:'/Index/serverinfo',
		type:'get',
		data:{'page':page},
		datatype:'json',
		success:function(a,b){
			// 区服数据处理
			serveres(a);
			// console.log(a);
		}
	})
}

// 区服数据处理
function serveres(a){
	$(function(){
		var server = '';
		if (a.code=='200'){
			var obj = a.data.data.data;
			var zdarr = a.data.zhiding;
			var tjarr = a.data.recgame;

			// 置顶区服
			Object.keys(obj).forEach(function(key){
				var zd = false;
				Object.keys(zdarr).forEach(function(key1){
					if (zdarr[key1]['gid'] == obj[key]['gid']){
						zd = true;
						server += '<li class="zhiding now clearfix">';
						server += '<div class="bd clearfix">';
						server += '<span class="date fl">' + obj[key]['server_date'].substr(5) + ' ' + obj[key]['server_time'].substr(0,5) + '</span>';
						server += '<em class="ico fl"></em>';
						server += '<a href="http://web.ofwan.com/WebGames/TestOpenServer/g/'+obj[key]['gid']+'" target="_blank" title="" class="gname fl">' + obj[key]['app_name'] + '</a>';
						server += '<a class="gline fr" href="http://web.ofwan.com/WebGames/TestOpenServer/g/'+obj[key]['gid']+'" target="_blank">' + obj[key]['server_name']+ '服' + '</a>';
						server += '</div>';
						server += '<div class="bili-fuli" style="display: none;">';
						server += '<p class="bili">充值比例 1:1000</p>';
						server += '<p class="fuli">上线赠送VIP3，永久黄金王权卡，绝版称号免费领</p>';
						server += '</div>';
						server += '</li>';
						return false;
					}
				})
			})

			// console.log(zdarr);return false;
			Object.keys(obj).forEach(function(key){
				var tj = false;
				// console.log(obj[key]['gid']);
				var zd = false;
				Object.keys(zdarr).forEach(function(key1){
					if (zdarr[key1]['gid'] == obj[key]['gid']){
						zd = true;
						return false;
					}
				})
				Object.keys(tjarr).forEach(function(key2){
					if (tjarr[key2]['gid'] == obj[key]['gid']){
						tj = true;
						server += '<li class="tuijian clearfix">';
						server += '<div class="bd clearfix">';
						server += '<span class="date fl">' + obj[key]['server_date'].substr(5) + ' ' + obj[key]['server_time'].substr(0,5) + '</span>';
						server += '<em class="ico fl"></em>';
						server += '<a href="http://web.ofwan.com/WebGames/TestOpenServer/g/'+obj[key]['gid']+'" target="_blank" title="" class="gname fl">' + obj[key]['app_name'] + '</a>';
						server += '<a class="gline fr" href="http://web.ofwan.com/WebGames/TestOpenServer/g/'+obj[key]['gid']+'" target="_blank">' + obj[key]['server_name']+ '服' + '</a>';
						server += '</div>';
						server += '<div class="bili-fuli" style="display: none;">';
						server += '<p class="bili">充值比例 1:1000</p>';
						server += '<p class="fuli">上线赠送VIP3，永久黄金王权卡，绝版称号免费领</p>';
						server += '</div>';
						server += '</li>';
						return false;
					}
				})
				if (!tj && !zd){
					server += '<li class=" clearfix">';
					server += '<div class="bd clearfix">';
					server += '<span class="date fl">' + obj[key]['server_date'].substr(5) + ' ' + obj[key]['server_time'].substr(0,5) + '</span>';
					server += '<em class="ico fl"></em>';
					server += '<a href="http://web.ofwan.com/WebGames/TestOpenServer/g/'+obj[key]['gid']+'" target="_blank" title="" class="gname fl">' + obj[key]['app_name'] + '</a>';
					server += '<a class="gline fr" href="http://web.ofwan.com/WebGames/TestOpenServer/g/'+obj[key]['gid']+'" target="_blank">' + obj[key]['server_name']+ '服' + '</a>';
					server += '</div>';
					server += '<div class="bili-fuli" style="display: none;">';
					server += '<p class="bili">充值比例 1:1000</p>';
					server += '<p class="fuli">上线赠送VIP3，永久黄金王权卡，绝版称号免费领</p>';
					server += '</div>';
					server += '</li>';
				}
				// server += '<li class="zhiding now clearfix">';
			});
			
		}else{
				server += '<li class="zhiding now clearfix"><div class="bd clearfix">	<span class="date fl">12-30 10:00</span>	<em class="ico fl"></em>	<a href="#" target="_blank" title="" class="gname fl">倾国之怒</a>	<a class="gline fr" href="#" target="_blank">5服</a></div><div class="bili-fuli" style="display: none;">	<p class="bili">充值比例 1:1000</p>	<p class="fuli">上线赠送VIP3，永久黄金王权卡，绝版称号免费领</p></div></li>';
		}
		$('.game-open ul').empty().append(server);
		// console.log(a.data.page);return false;
		if (a.data.page !== null){
			$('.open_pages').empty().append(a.data.page);
			var lastpage = a.data.data.last_page;
			// console.log(a.data.data.current_page);
			// 修改返回分页属性
			for (var i = 0; i <= (lastpage+1); i++) {
				if (a.data.data.current_page == i){
					continue;
				}
				var pagedata = $('.open_pages .pagination li:eq(' + i + ') a').attr('href');
				// console.log(pagedata);
				$('.open_pages .pagination li:eq(' + i + ') a').attr('href','javascript:void(getserver("'+pagedata+'"))');
			}
		}else{
			$('.open_pages').empty();
		}
	})
}






// 全部游戏
$(function(){
	$.ajax({
		url:'/Game/gamelist',
		datatype:'json',
		type:'get',
		success:function(a,b){
			// console.log(a);return false;
			datares(a);
		},error:function(c,d){
			console.log(d);
		}
	})
})

// 根据字母关键字搜索游戏
$(function(){
	$('.f_zimu').click(function(){
		var txt = $(this).text().split('-');
		// console.log(txt);return false;
		if (txt == '全部'){
			$.ajax({
				url:'/Game/gamelist',
				datatype:'json',
				type:'get',
				success:function(a,b){
					datares(a);
				},error:function(c,d){
					console.log(d);
				}
			})
			window.event.returnValue=false;
		}
		$.ajax({
			url:'/Game/selname',
			data:{'arr':txt},
			datatype:'json',
			type:'get',
			success:function(c,d){
				// console.log(c);
				datares(c);
			}
		})
	})
})

// 根据游戏类型搜索游戏
$(function(){
	$('.f_type').click(function(){
		var txt = $(this).text();

		$.ajax({
			url:'/Game/seltheme',
			data:{'txt':txt},
			datatype:'json',
			type:'get',
			success:function(f,g){
				// console.log(f);
				datares(f);
			}
		})
	})
})

// 游戏结果分页
function gamelist(data)
{
	// alert(data);
	var url = data.split('?')[0];
	var page= data.split('?')[1].substr(5);
	// alert(page);return false;
	$.ajax({
		url:url,
		datatype:'json',
		data:{page:page},
		type:'get',
		success:function(a,b){
			datares(a);
		}
	})
}


// 游戏数据处理
function datares(a)
{
	var txt = '';
	if (a.code=='200'){
		var obj = a.data.gameres.data;
		var type= a.data.labelres;
		var hotgame = a.data.hotgame;
		var hotarr = new Array();
		// console.log(hotgame);
		if (obj.length > 0){
			// 根据推荐游戏排序
			if (hotgame !== undefined){
				// 轮播与推荐游戏数组
				Object.keys(hotgame).forEach(function(key2){
					hotarr.push(hotgame[key2]);
				})
				// console.log(JSON.stringify(obj));
				// console.log(JSON.stringify(hotarr));
				// return false;
				// console.log(hotarr.length);
				// return false;
				for (var hotnum = 0;hotnum < hotarr.length; hotnum++){
					for (var objnum = 0;objnum < obj.length; objnum++){
						if (hotarr[hotnum] == obj[objnum]['gid']){
							txt += '<li><div class="item-v clearfix">';
							txt += '<a href="http://web.ofwan.com/Web/webGames/TestOpenServer/g/' + hotarr[hotnum] + '" target="_blank" class="gimg fl">';
							txt += '<img src="http://download.qwyouxi.com/' + obj[objnum]['osspath'] + '" class="lazy-img" width="280" height="140" style="display: inline;">';
							txt += '<div class="img_info">';
							txt += '<p class="gname">' + obj[objnum]['app_name'] +'</p>';
							// Object.keys(type).forEach(function(key1){
							// 	if (type[key1]['gid'] == hotgame[hotnum]['gid'] && type[key1]['type'] == 'gametype'){
									// txt += '<p class="txt">' + type[key1]['label_name'] + '</p>';
									txt += '<p class="txt">即时战斗</p>';
							// 	}
							// })
							txt += '<p class="txt">比例 1:100</p><p class="ico"><span class="hotico"><font>HOT</font></span></p><p class="play">开始游戏</p></div>';
							txt += '</a>';

							txt += '<div class="ginfo webinfo fl"><div class="clearfix"><a href="#" target="_blank" class="fl gname" title="' + obj[objnum]['app_name'] + '">' + obj[objnum]['app_name'] + '</a></div><div class="fuli" title="首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包">首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包</div>';
							txt += '<div class="gtype mtop15 clearfix">';
							Object.keys(type).forEach(function(key2){
								if (type[key2]['gid'] == obj[objnum]['gid']){
									txt += '<span class="type">' + type[key2]['label_name'] + '</span>';
								}
							})
							txt += '<a href="http://web.ofwan.com/Web/webGames/TestOpenServer/g/' + hotarr[hotnum] + '" target="_blank" class="play">开始游戏</a><a href="#" target="_blank" class="wd_down">微端下载</a></div></div>';
							txt += '</div></li>';
						}
					}
				}
				// console.log(txt);return false;
			}
			
			// console.log(txt);return false;

			Object.keys(obj).forEach(function(key){
				if (hotarr.indexOf(obj[key]['gid']) < 0){
					txt += '<li><div class="item-v clearfix">';
					txt += '<a href="#" target="_blank" class="gimg fl">';
					txt += '<img src="http://download.qwyouxi.com/' + obj[key]['osspath'] + '" class="lazy-img" width="280" height="140" style="display: inline;">';
					txt += '<div class="img_info">';
					txt += '<p class="gname">' + obj[key]['app_name'] +'</p>';
					// Object.keys(type).forEach(function(key1){
					// 	if (type[key1]['gid'] == obj[key]['gid'] && type[key1]['type'] == 'gametype'){
							// txt += '<p class="txt">' + type[key1]['label_name'] + '</p>';
							txt += '<p class="txt">即时战斗</p>';
					// 	}
					// })
					txt += '<p class="txt">比例 1:100</p><p class="ico"><span class="hotico"><font>HOT</font></span></p><p class="play">开始游戏</p></div>';
					txt += '</a>';

					txt += '<div class="ginfo webinfo fl"><div class="clearfix"><a href="#" target="_blank" class="fl gname" title="' + obj[key]['app_name'] + '">' + obj[key]['app_name'] + '</a></div><div class="fuli" title="首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包">首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包</div>';
					txt += '<div class="gtype mtop15 clearfix">';
					Object.keys(type).forEach(function(key2){
						if (type[key2]['gid'] == obj[key]['gid']){
							txt += '<span class="type">' + type[key2]['label_name'] + '</span>';
						}
					})
					txt += '<a href="http://web.ofwan.com/Web/webGames/TestOpenServer/g/' + obj[key]['gid'] + '" target="_blank" class="play">开始游戏</a><a href="#" target="_blank" class="wd_down">微端下载</a></div></div>';
					txt += '</div></li>';
				}
				// console.log(obj[key]['app_name']);
			});
		}else{
				txt += '<li><div class="item-v clearfix"><a href="#" target="_blank" class="gimg fl"><img src="/static/home/images/default/img1-2.jpg" class="lazy-img" width="280" height="140" style="display: inline;"><div class="img_info"><p class="gname">创世OL</p><p class="txt">即时战斗</p><p class="txt">比例 1:100</p><p class="ico"><span class="hotico"><font>HOT</font></span></p>	<p class="play">开始游戏</p>	</div></a><div class="ginfo webinfo fl"><div class="clearfix"><a href="#" target="_blank" class="fl gname" title="创世OL">创世OL</a></div><div class="fuli" title="首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包">首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包</div><div class="gtype mtop15 clearfix"><span class="type">绿色游戏</span><span class="type">即时战斗</span><span class="type">传奇</span><a href="#" target="_blank" class="play">开始游戏</a><a href="#" target="_blank" class="wd_down">微端下载</a></div></div></div></li>';
				txt += '<li><div class="item-v clearfix"><a href="#" target="_blank" class="gimg fl"><img src="/static/home/images/default/img1-2.jpg" class="lazy-img" width="280" height="140" style="display: inline;"><div class="img_info"><p class="gname">创世OL</p><p class="txt">即时战斗</p><p class="txt">比例 1:100</p><p class="ico"><span class="hotico"><font>HOT</font></span></p>	<p class="play">开始游戏</p>	</div></a><div class="ginfo webinfo fl"><div class="clearfix"><a href="#" target="_blank" class="fl gname" title="创世OL">创世OL</a></div><div class="fuli" title="首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包">首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包</div><div class="gtype mtop15 clearfix"><span class="type">绿色游戏</span><span class="type">即时战斗</span><span class="type">传奇</span><a href="#" target="_blank" class="play">开始游戏</a><a href="#" target="_blank" class="wd_down">微端下载</a></div></div></div></li>';
				txt += '<li><div class="item-v clearfix"><a href="#" target="_blank" class="gimg fl"><img src="/static/home/images/default/img1-2.jpg" class="lazy-img" width="280" height="140" style="display: inline;"><div class="img_info"><p class="gname">创世OL</p><p class="txt">即时战斗</p><p class="txt">比例 1:100</p><p class="ico"><span class="hotico"><font>HOT</font></span></p>	<p class="play">开始游戏</p>	</div></a><div class="ginfo webinfo fl"><div class="clearfix"><a href="#" target="_blank" class="fl gname" title="创世OL">创世OL</a></div><div class="fuli" title="首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包">首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包</div><div class="gtype mtop15 clearfix"><span class="type">绿色游戏</span><span class="type">即时战斗</span><span class="type">传奇</span><a href="#" target="_blank" class="play">开始游戏</a><a href="#" target="_blank" class="wd_down">微端下载</a></div></div></div></li>';
				txt += '<li><div class="item-v clearfix"><a href="#" target="_blank" class="gimg fl"><img src="/static/home/images/default/img1-2.jpg" class="lazy-img" width="280" height="140" style="display: inline;"><div class="img_info"><p class="gname">创世OL</p><p class="txt">即时战斗</p><p class="txt">比例 1:100</p><p class="ico"><span class="hotico"><font>HOT</font></span></p>	<p class="play">开始游戏</p>	</div></a><div class="ginfo webinfo fl"><div class="clearfix"><a href="#" target="_blank" class="fl gname" title="创世OL">创世OL</a></div><div class="fuli" title="首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包">首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包</div><div class="gtype mtop15 clearfix"><span class="type">绿色游戏</span><span class="type">即时战斗</span><span class="type">传奇</span><a href="#" target="_blank" class="play">开始游戏</a><a href="#" target="_blank" class="wd_down">微端下载</a></div></div></div></li>';
				txt += '<li><div class="item-v clearfix"><a href="#" target="_blank" class="gimg fl"><img src="/static/home/images/default/img1-2.jpg" class="lazy-img" width="280" height="140" style="display: inline;"><div class="img_info"><p class="gname">创世OL</p><p class="txt">即时战斗</p><p class="txt">比例 1:100</p><p class="ico"><span class="hotico"><font>HOT</font></span></p>	<p class="play">开始游戏</p>	</div></a><div class="ginfo webinfo fl"><div class="clearfix"><a href="#" target="_blank" class="fl gname" title="创世OL">创世OL</a></div><div class="fuli" title="首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包">首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包</div><div class="gtype mtop15 clearfix"><span class="type">绿色游戏</span><span class="type">即时战斗</span><span class="type">传奇</span><a href="#" target="_blank" class="play">开始游戏</a><a href="#" target="_blank" class="wd_down">微端下载</a></div></div></div></li>';
				txt += '<li><div class="item-v clearfix"><a href="#" target="_blank" class="gimg fl"><img src="/static/home/images/default/img1-2.jpg" class="lazy-img" width="280" height="140" style="display: inline;"><div class="img_info"><p class="gname">创世OL</p><p class="txt">即时战斗</p><p class="txt">比例 1:100</p><p class="ico"><span class="hotico"><font>HOT</font></span></p>	<p class="play">开始游戏</p>	</div></a><div class="ginfo webinfo fl"><div class="clearfix"><a href="#" target="_blank" class="fl gname" title="创世OL">创世OL</a></div><div class="fuli" title="首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包">首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包</div><div class="gtype mtop15 clearfix"><span class="type">绿色游戏</span><span class="type">即时战斗</span><span class="type">传奇</span><a href="#" target="_blank" class="play">开始游戏</a><a href="#" target="_blank" class="wd_down">微端下载</a></div></div></div></li>';
		}
		// console.log(a.data.gameres);return false;
		// 分页
		if (a.data.page === null){
			$('.paging').empty();
		}else{
			$('.paging').empty().append(a.data.page);
		}

		// 修改返回分页属性
		var lastpage = a.data.gameres.last_page;
		// $('.news-list').append(a.data.page);
		// console.log(lastpage);return false;
		if (a.data.page !== null){
			for (var i = 0; i <= (lastpage+1); i++) {
				if (a.data.gameres.current_page == i){
					continue;
				}
				var pagedata = $('.pagination li:eq(' + i + ') a ').attr('href');
				// console.log(pagedata);
				$('.pagination li:eq(' + i + ') a').attr('href','javascript:void(gamelist("'+pagedata+'"))');
				// $('.pagination li:eq(' + i + ') a').attr('οnclick','newspages()');
			}
		}

	}else{
		txt += '<li><div class="item-v clearfix"><a href="#" target="_blank" class="gimg fl"><img src="/static/home/images/default/img1-2.jpg" class="lazy-img" width="280" height="140" style="display: inline;"><div class="img_info"><p class="gname">创世OL</p><p class="txt">即时战斗</p><p class="txt">比例 1:100</p><p class="ico"><span class="hotico"><font>HOT</font></span></p>	<p class="play">开始游戏</p>	</div></a><div class="ginfo webinfo fl"><div class="clearfix"><a href="#" target="_blank" class="fl gname" title="创世OL">创世OL</a></div><div class="fuli" title="首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包">首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包</div><div class="gtype mtop15 clearfix"><span class="type">绿色游戏</span><span class="type">即时战斗</span><span class="type">传奇</span><a href="#" target="_blank" class="play">开始游戏</a><a href="#" target="_blank" class="wd_down">微端下载</a></div></div></div></li>';
		txt += '<li><div class="item-v clearfix"><a href="#" target="_blank" class="gimg fl"><img src="/static/home/images/default/img1-2.jpg" class="lazy-img" width="280" height="140" style="display: inline;"><div class="img_info"><p class="gname">创世OL</p><p class="txt">即时战斗</p><p class="txt">比例 1:100</p><p class="ico"><span class="hotico"><font>HOT</font></span></p>	<p class="play">开始游戏</p>	</div></a><div class="ginfo webinfo fl"><div class="clearfix"><a href="#" target="_blank" class="fl gname" title="创世OL">创世OL</a></div><div class="fuli" title="首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包">首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包</div><div class="gtype mtop15 clearfix"><span class="type">绿色游戏</span><span class="type">即时战斗</span><span class="type">传奇</span><a href="#" target="_blank" class="play">开始游戏</a><a href="#" target="_blank" class="wd_down">微端下载</a></div></div></div></li>';
		txt += '<li><div class="item-v clearfix"><a href="#" target="_blank" class="gimg fl"><img src="/static/home/images/default/img1-2.jpg" class="lazy-img" width="280" height="140" style="display: inline;"><div class="img_info"><p class="gname">创世OL</p><p class="txt">即时战斗</p><p class="txt">比例 1:100</p><p class="ico"><span class="hotico"><font>HOT</font></span></p>	<p class="play">开始游戏</p>	</div></a><div class="ginfo webinfo fl"><div class="clearfix"><a href="#" target="_blank" class="fl gname" title="创世OL">创世OL</a></div><div class="fuli" title="首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包">首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包</div><div class="gtype mtop15 clearfix"><span class="type">绿色游戏</span><span class="type">即时战斗</span><span class="type">传奇</span><a href="#" target="_blank" class="play">开始游戏</a><a href="#" target="_blank" class="wd_down">微端下载</a></div></div></div></li>';
		txt += '<li><div class="item-v clearfix"><a href="#" target="_blank" class="gimg fl"><img src="/static/home/images/default/img1-2.jpg" class="lazy-img" width="280" height="140" style="display: inline;"><div class="img_info"><p class="gname">创世OL</p><p class="txt">即时战斗</p><p class="txt">比例 1:100</p><p class="ico"><span class="hotico"><font>HOT</font></span></p>	<p class="play">开始游戏</p>	</div></a><div class="ginfo webinfo fl"><div class="clearfix"><a href="#" target="_blank" class="fl gname" title="创世OL">创世OL</a></div><div class="fuli" title="首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包">首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包</div><div class="gtype mtop15 clearfix"><span class="type">绿色游戏</span><span class="type">即时战斗</span><span class="type">传奇</span><a href="#" target="_blank" class="play">开始游戏</a><a href="#" target="_blank" class="wd_down">微端下载</a></div></div></div></li>';
		txt += '<li><div class="item-v clearfix"><a href="#" target="_blank" class="gimg fl"><img src="/static/home/images/default/img1-2.jpg" class="lazy-img" width="280" height="140" style="display: inline;"><div class="img_info"><p class="gname">创世OL</p><p class="txt">即时战斗</p><p class="txt">比例 1:100</p><p class="ico"><span class="hotico"><font>HOT</font></span></p>	<p class="play">开始游戏</p>	</div></a><div class="ginfo webinfo fl"><div class="clearfix"><a href="#" target="_blank" class="fl gname" title="创世OL">创世OL</a></div><div class="fuli" title="首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包">首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包</div><div class="gtype mtop15 clearfix"><span class="type">绿色游戏</span><span class="type">即时战斗</span><span class="type">传奇</span><a href="#" target="_blank" class="play">开始游戏</a><a href="#" target="_blank" class="wd_down">微端下载</a></div></div></div></li>';
		txt += '<li><div class="item-v clearfix"><a href="#" target="_blank" class="gimg fl"><img src="/static/home/images/default/img1-2.jpg" class="lazy-img" width="280" height="140" style="display: inline;"><div class="img_info"><p class="gname">创世OL</p><p class="txt">即时战斗</p><p class="txt">比例 1:100</p><p class="ico"><span class="hotico"><font>HOT</font></span></p>	<p class="play">开始游戏</p>	</div></a><div class="ginfo webinfo fl"><div class="clearfix"><a href="#" target="_blank" class="fl gname" title="创世OL">创世OL</a></div><div class="fuli" title="首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包">首充任意金额可获得绝版神兵(20阶武器)+万宝盼盼(自动拾取)+大量元宝+七日礼包</div><div class="gtype mtop15 clearfix"><span class="type">绿色游戏</span><span class="type">即时战斗</span><span class="type">传奇</span><a href="#" target="_blank" class="play">开始游戏</a><a href="#" target="_blank" class="wd_down">微端下载</a></div></div></div></li>';
	}
	$('#gamelist').empty().append(txt);
}



// 游戏区服信息





