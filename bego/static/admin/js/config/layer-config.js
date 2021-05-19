
//默认弹窗iframe
function ifrShowLayerLar(title,url){
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
        content: [PORTLINK + url, 'yes'], //iframe的url，no代表不显示滚动条
     
    });
}

function ifrShowLayerLit(title,url){
    layer.open({
        type: 2,
        title: [title,'font-size:16px; color:#333;'],  //标题文本 , 样式 ， 不显示标题 title:false
        area:['500px','500px'],   //定义宽高
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
        content: [PORTLINK + url, 'yes'], //iframe的url，no代表不显示滚动条
     
    });
}

function showLayerLit(title){
    layer.open({
        type: 1,
        //skin: 'pop-class',
        title: [title,'font-size:16px; color:#333;'],  //标题文本 , 样式 ， 不显示标题 title:false
        area:['500px','300px'],   //定义宽高
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
        content:$('#popup'), //注意，如果str是object，那么需要字符拼接。
        success:function(){
            
        },
        end:function(){
            
        }
    });

}


//关闭layer
function hideLayer(){
    layer.close(layer.index)
}
