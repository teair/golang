$(function(){
	$.ajax({
		url:'/Server/serverinfo',
		datatype:'json',
		type:'get',
		success:function(a,b){
			serveres(a);
		}
	})
})

// 区服分页
function getserver(data)
{
	var page = data.split('=')[1];
	$.ajax({
		url:'/Server/serverinfo',
		datatype:'json',
		data:{page:page},
		type:'get',
		success:function(a,b){
			serveres(a);
		}
	})
}


function serveres(a)
{
	if (a.code=='200'){
		// console.log(a.data.labelres);return false;
		var txt = '';
		var serobj = a.data.serveres.data;
		var labobj = a.data.labelres;
		Object.keys(serobj).forEach(function(key){
			var labelarr = [];
			// console.log(serobj[key]);
			txt += '<tr>';
			txt += '<td><a href="javascript:;" >' + serobj[key]['app_name'] + '</a></td>';
			txt += '<td><a href="javascript:;" >' + serobj[key]['server_name'] + '</a></td>';
			txt += '<td><a href="javascript:;" >';
			Object.keys(labobj).forEach(function(key1){
				if (serobj[key]['gid'] == labobj[key1]['gid']){
					labelarr.push(labobj[key1]['label_name']);
				}
			})
			// console.log(labelarr);
			if (labelarr.length > 0){
				txt += labelarr.join('，');
			}else{
				txt += '暂无游戏题材!';
			}
			
			txt += '</a></td>';
			txt += '<td>'+serobj[key]['server_date'].substr(5)+'</td>';
			txt += '<td>'+serobj[key]['server_time']+'</td>';
			txt += '<td><a href="javascript:;" onclick=getgift("' + serobj[key]['gid'] + ',' + serobj[key]['sid'] + '")>领取礼包</a></td>';
			txt += '<td><a href="http://web.ofwan.com/WebGames/TestOpenServer/g/'+serobj[key]['gid']+'" target="_blank" class="go">开始游戏</a></td>';
			txt += '<td>  火爆开启中 </td>';
			txt += '</tr>';
		});
	}else{
		txt += '<tr><td><a target="_blank">神戒</a></td><td><a href="#" target="_blank">公测858区</a></td><td><a href="#" target="_blank">角色扮演</a></td><td>01-08</td><td>15:00</td><!-- <td>9377</td> --><td><a href="#" target="_blank">领取礼包</a></td><td><a href="#" target="_blank" class="go">开始游戏</a></td><td>  火爆开启中 </td></tr>';
		txt += '<tr><td><a target="_blank">神戒</a></td><td><a href="#" target="_blank">公测858区</a></td><td><a href="#" target="_blank">角色扮演</a></td><td>01-08</td><td>15:00</td><!-- <td>9377</td> --><td><a href="#" target="_blank">领取礼包</a></td><td><a href="#" target="_blank" class="go">开始游戏</a></td><td>  火爆开启中 </td></tr>';
		txt += '<tr><td><a target="_blank">神戒</a></td><td><a href="#" target="_blank">公测858区</a></td><td><a href="#" target="_blank">角色扮演</a></td><td>01-08</td><td>15:00</td><!-- <td>9377</td> --><td><a href="#" target="_blank">领取礼包</a></td><td><a href="#" target="_blank" class="go">开始游戏</a></td><td>  火爆开启中 </td></tr>';
	}
	// return false;
	// console.log(txt);
	$('.kaifu-table tbody').empty().append(txt);
	if (a.data.page !== null){
		$('.paging').empty().append(a.data.page);
		var lastpage = a.data.serveres.last_page;
		// console.log(a.data.data.current_page);
		// 修改返回分页属性
		for (var i = 0; i <= (lastpage+1); i++) {
			if (a.data.serveres.current_page == i){
				continue;
			}
			var pagedata = $('.paging .pagination li:eq(' + i + ') a').attr('href');
			// console.log(pagedata);
			$('.paging .pagination li:eq(' + i + ') a').attr('href','javascript:void(getserver("'+pagedata+'"))');
		}
	}else{
		$('.paging').empty();
	}
}


function getgift(gid,sid)
{
	// alert(gid,sid);return false;
	$.ajax({
		url:'/Server/giftrecord',
		datatype:'json',
		type:'post',
		success:function(data){
			if (data.code == '200'){
				window.location.href='http://web.h5.cc/WebGames/getgift/g/' + gid + '/username/' + data.data.uid + '/server_id/' + sid;
			}
			/*switch(data.code){
				case 1:
					layer.open({
						title:'礼包领取成功：',
						content:data.data.giftname + "：" + "<input type='text' name='giftcode' style='width:80px;text-align:center;' value='"+data.data.giftcode+"'><a href='javascript:;' style='text-decoration:none;' onclick=copys('"+data.data.giftcode+"');>点击复制</a>"
					})
					break;
				case 4:
					layer.open({
						title:'您已经领取过该礼包了！',
						content:data.data.giftname + "：" +  "<input type='text' name='giftcode' style='width:80px;text-align:center;' value='"+data.data.giftcode+"'><a href='javascript:;' style='text-decoration:none;' onclick=copys('"+data.data.giftcode+"');>点击复制</a>"
					})

					break;
				default :
					layer.msg(data.msg);
					break;
			}*/
		}
	})
}

function copys(msg){
	$("input[name='giftcode']").select();
    document.execCommand("Copy")
}


