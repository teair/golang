//////定义上传方法函数

function PreviewImage(imgFile,name) { 

    //定义图片路径 
    var path;
    
    $(".picbox").prepend("<div class='showcol'><div class='imgcol'><img src='"+ imgFile +"' data-name='"+name+"' /></div><a class='hide delete-btn'>X</a></div>");
    
    // $(".picbox").append("<div class='filebox "+ classfb +"'><input type='file' class='upmul' name='upmul" + fb + "' onchange='PreviewImage(this)'/></div>");

} 


//控制"按钮"显示与隐藏
$(".picbox").off("mouseenter",".showcol").on("mouseenter",".showcol",function(){
    var that=this;
    var dom=$(that).children(".delete-btn");
    dom.removeClass("hide");
    //为点击事件解绑，防止重复执行
    dom.off("click");
    dom.on("click",function(){
    	//删除当前图片
      var thissrc = dom.parents('.showcol').find("img").data("name");
      for(var i=0;i<picArr.length;i++){
        if(thissrc==picArr[i]){
          console.log(picArr[i])
          // console.log($.inArray(picArr[i],arr));
          picArr.splice(i,1);
          console.log(picArr)
        }
      }
      dom.parents('.showcol').remove();

  });
}).off("mouseleave",".showcol").on("mouseleave",".showcol",function(){
    var that=this;
    $(that).children(".delete-btn").addClass("hide");
})


