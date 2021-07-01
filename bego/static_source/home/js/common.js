

$(function(){
	// 获取推荐游戏
	$.ajax({
		url:'/Index/getlinks',
		datatype:'json',
		type:'post',
		success:function(a,b){
			if (a.code=='200'){
				// console.log(a);
				var gameobj = a.data.recgame;
				var linkobj = a.data.links;
				var recgame = '';
				var links = '';
				if (gameobj.length > 0){
					Object.keys(gameobj).forEach(function(key){
						// console.log(gameobj[key]);
						recgame += '<a href="http://web.ofwan.com/WebGames/TestOpenServer/g/'+gameobj[key]['gid']+'" target="_blank">'+gameobj[key]['app_name']+'</a>';
					});
				}else{
					recgame += '<a href="#" target="_blank">神戒</a><a href="#" target="_blank">灭神</a><a href="#" target="_blank">攻铩</a><a href="#" target="_blank">赤月传说Ⅱ</a><a href="#" target="_blank">武动苍穹</a><a href="#" target="_blank">血饮传说</a><a href="#" target="_blank">猛将天下</a><a href="#" target="_blank">创世热血战歌</a><a href="#" target="_blank">屠龙战记</a><a href="#" target="_blank">剑道仙语</a><a href="#" target="_blank">雷霆之怒</a><a href="#" target="_blank">皇图</a><a href="#" target="_blank">烈焰</a><a href="#" target="_blank">蓝月传奇</a><a href="#" target="_blank">战神风云</a><a href="#" target="_blank">魅影传说</a><a href="#" target="_blank">乱斗乾坤</a><a href="#" target="_blank">传奇世界</a><a href="#" target="_blank">赤月传说</a><a href="#" target="_blank">传奇荣耀</a><a href="#" target="_blank">屠龙传说</a><a href="#" target="_blank">九阴绝学</a>';
		            recgame += '<a href="#" target="_blank">传奇盛世</a><a href="#" target="_blank">斗破苍穹2</a><a href="#" target="_blank">红月传说</a><a href="#" target="_blank">天书世界</a>';
				}
				if (linkobj.length > 0){
					Object.keys(linkobj).forEach(function(key1){
						// console.log(gameobj[key]);
						links += '<a href="http://'+linkobj[key1]['weburl']+'" target="_blank">'+linkobj[key1]['webname']+'</a>';
					});
				}else{
					links += '<a href="#" target="_blank">网页游戏</a><a href="#" target="_blank">9377手游</a><a href="#" target="_blank">360游戏</a><a href="#" target="_blank">游侠页游</a><a href="#" target="_blank">太平洋页游</a><a href="#" target="_blank">腾讯游戏</a><a href="#" target="_blank">A9VG</a><a href="#" target="_blank">9K9K开服表</a><a href="#" target="_blank">2243手游网</a><a href="#" target="_blank">软件侠</a><a href="#" target="_blank">快吧游戏</a><a href="#" target="_blank">游戏交易平台</a><a href="#" target="_blank">手游交易</a><a href="#" target="_blank">ns游戏</a><a href="#" target="_blank">加速器</a>';
		            links += '<a href="#" target="_blank">网页游戏</a><a href="#" target="_blank">9377手游</a><a href="#" target="_blank">360游戏</a><a href="#" target="_blank">游侠页游</a><a href="#" target="_blank">太平洋页游</a><a href="#" target="_blank">腾讯游戏</a><a href="#" target="_blank">A9VG</a><a href="#" target="_blank">9K9K开服表</a><a href="#" target="_blank">2243手游网</a><a href="#" target="_blank">软件侠</a><a href="#" target="_blank">快吧游戏</a><a href="#" target="_blank">游戏交易平台</a><a href="#" target="_blank">手游交易</a><a href="#" target="_blank">ns游戏</a><a href="#" target="_blank">加速器</a>';
				}				
				$('.remgame-box .links-list .lbox').empty().append(recgame);
				$('.friends-box .links-list .lbox').empty().append(links);
			}
		}
	})
	$(".more").click(function(){
		$(this).toggleClass("morev")
		$(this).parent().find(".links-list").toggleClass("links-auto")
	})
})
