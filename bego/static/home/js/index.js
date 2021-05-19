
// 点击轮播图页面跳转
function gamedetail(gid)
{
	window.location.href='http://web.ofwan.com/WebGames/TestOpenServer/g/'+gid;
}

// 新闻资讯
$(function(){
	$.ajax({
		url:'/Index/gamenews',
		datatype:'json',
		type:'get',
		success:function(a,b){
			datares(a);
		}
	})
})

// 新闻分页
function newspages(data)
{
	// layer.alert(data);
	var page = data.split('=')[1];
	$.ajax({
		url:'/Index/gamenews',
		data:{'page':page},
		type:'get',
		success:function(a,b){
			datares(a);
		},error:function(c,d){
			console.log(d);
		}
	})
}

// 新闻数据处理
function datares(a)
{
	if (a.code=='200'){
		var obj = a.data.data.data;
		// console.log(a.data);return false;
		var news = '<ul class="clearfix">';
		Object.keys(obj).forEach(function(key){
			// console.log(obj[key]['title']);
			news += '<li class="rel"><a href="#" title="' + obj[key]['title'] + '" target="_blank">' + obj[key]['title'] + '</a></li>';
		});
		news += '</ul>';
		if (a.data.page !== null){
			news += a.data.page;
		}
		$('.news-list').empty().append(news);
		// console.log(a.data.data.current_page);
		var lastpage = a.data.data.last_page;
		// $('.news-list').append(a.data.page);
		if (a.data.page !== null){
			// console.log(a.data.page);return false;
			// $('.news-list').empty().append(a.data.page);
			// 修改返回分页属性
			for (var i = 0; i <= (lastpage+1); i++) {
				if (a.data.data.current_page == i){
					continue;
				}
				var pagedata = $('.news-list .pagination li:eq(' + i + ') a ').attr('href');
				// console.log(pagedata);
				$('.news-list .pagination li:eq(' + i + ') a').attr('href','javascript:void(newspages("'+pagedata+'"))');
				// $('.pagination li:eq(' + i + ') a').attr('οnclick','newspages()');
			}
		}

	}else{
		var txt = '';
		txt += '<li><a href="#" title="【重要公告】实名认证注册和防沉迷系统开启公告" target="_blank">【重要公告】实名认证注册和防沉迷系统开启公告</a></li><li><a href="#" title="【问卷有礼】诚邀您填写问卷，有礼相赠~" target="_blank">【问卷有礼】诚邀您填写问卷，有礼相赠~</a></li>';
		txt += '<li><a href="#" title="《王城英雄》跨服攻城战季后赛来临！" target="_blank">《王城英雄》跨服攻城战季后赛来临！</a></li><li><a href="#" title="《将星》打boss送首充！vip经验免费拿！" target="_blank">《将星》打boss送首充！vip经验免费拿！</a></li>';
		txt += '<li><a href="#" title="《传奇霸业》五周年！9重豪礼送送送" target="_blank">《传奇霸业》五周年！9重豪礼送送送</a></li><li><a href="#" title="【重要公告】实名认证注册和防沉迷系统开启公告" target="_blank">【重要公告】实名认证注册和防沉迷系统开启公告</a></li>';
		txt += '<li><a href="#" title="【问卷有礼】诚邀您填写问卷，有礼相赠~" target="_blank">【问卷有礼】诚邀您填写问卷，有礼相赠~</a></li><li><a href="#" title="《王城英雄》跨服攻城战季后赛来临！" target="_blank">《王城英雄》跨服攻城战季后赛来临！</a></li>';
		txt += '<li><a href="#" title="《将星》打boss送首充！vip经验免费拿！" target="_blank">《将星》打boss送首充！vip经验免费拿！</a></li><li><a href="#" title="《传奇霸业》五周年！9重豪礼送送送" target="_blank">《传奇霸业》五周年！9重豪礼送送送</a></li>';
		txt += '<li><a href="#" title="【重要公告】实名认证注册和防沉迷系统开启公告" target="_blank">【重要公告】实名认证注册和防沉迷系统开启公告</a></li><li><a href="#" title="【问卷有礼】诚邀您填写问卷，有礼相赠~" target="_blank">【问卷有礼】诚邀您填写问卷，有礼相赠~</a></li>';
		txt += '<li><a href="#" title="《王城英雄》跨服攻城战季后赛来临！" target="_blank">《王城英雄》跨服攻城战季后赛来临！</a></li><li><a href="#" title="《将星》打boss送首充！vip经验免费拿！" target="_blank">《将星》打boss送首充！vip经验免费拿！</a></li>';
		txt += '<li><a href="#" title="《传奇霸业》五周年！9重豪礼送送送" target="_blank">《传奇霸业》五周年！9重豪礼送送送</a></li>';
		$('.news-list').append(txt);
	}
}



// 区服信息
$(function(){
	$.ajax({
		url:'/Index/serverinfo',
		type:'get',
		data:{'time':'today'},
		datatype:'json',
		success:function(a,b){
			// console.log(a);
			// 区服数据处理
			serveres(a);
		}
	})
	$(".mod-open-tabs dd").click(function(){
		$(".mod-open-tabs dd").removeClass('active');
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

			// var flg = false;
			// var obj5 = [{'name':'1'},{'name':2},{'name':3}];
			// Object.keys(obj5).forEach(function(key5){
			// 	// console.log(obj5[key5]);
			// 	if (obj5[key5]['name'] == 1){
			// 		flg = true;
			// 	}
			// })
			// console.log(flg);return false;

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
		// console.log(server);return false;
		$('.serverdata').empty().append(server);
		// return false;
		if (a.data.page !== null){
			$('.open_page').empty().append(a.data.page);
			var lastpage = a.data.data.last_page;
			// console.log(a.data.data.current_page);
			// 修改返回分页属性
			for (var i = 0; i <= (lastpage+1); i++) {
				if (a.data.data.current_page == i){
					continue;
				}
				var pagedata = $('.open_page .pagination li:eq(' + i + ') a').attr('href');
				// console.log(pagedata);
				$('.open_page .pagination li:eq(' + i + ') a').attr('href','javascript:void(getserver("'+pagedata+'"))');
			}
		}else{
			$('.open_page').empty();
		}
	})
}



