// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmTemplatesCoreCommAction = `<script type="text/javascript">
	function {%.param%}Action(obj, id, callback, tabsId) {
		var o=$(obj);
		var ste=o.data("state");
		var url=o.data("url");
		var template=o.data("template");
		var contentId=o.attr("href");
		
		if(ste!="disable"){
			if(!{%.param%}TabsSwitch(contentId,ste,tabsId)){
				return;
			}
		}
		if(window.state){
			window.state=ste;
		}
		if(ste=="add"||ste=="edit"||ste=="show"||ste=="disable"){
			$(contentId).html("");
			if($.isNotBlank(url)){
				$.ajax({
					type : 'post',
					url : url,
					dataType : 'json',
					data : {sIdTR : id},
					success : function(result) {
						if(result.success){
							if(ste!="disable"){
								var data=result.data;
								data.state=ste;
								//将url参数放入data
								var winUrl = $.url();
								for(var key in winUrl.param()){
									data[key]=winUrl.param()[key];
								}
								$(contentId).handlebars(template,data);
							}
							try{
								if(typeof(callback) == "function"){
									callback();
								}
							}catch(e){}
						}else{
							alert(result.message);
						}
					}
				});
			}else{
				var data={state:ste};
				var winUrl = $.url();
				for(var key in winUrl.param()){
					data[key]=winUrl.param()[key];
				}
				$(contentId).handlebars(template,data);
				try{
					if(typeof(callback) == "function"){
						callback();
					}
				}catch(e){}
			}
		}else{
			try{
				if(typeof(callback) == "function"){
					callback();
				}
			}catch(e){}
		}
	}
	//返回将要切换的tab是否已经选中
	function {%.param%}TabsSwitch(id,state,tabsId){
		if($.isBlank(tabsId)){
			tabsId="#{%.param%}NavTabs";
		}
		var navTabs=$(tabsId+" li");
		var a=navTabs.find("a[href=\""+id+"\"]");
		// 防止重复点击：修改时点击tab后成添加
		var cls=a.parent().attr("class");
		if(cls=="active"){
			return false;
		}
		a.tab("show");
		if(a.parent().css("display")=="none"){
			a.parent().show();
		}
		if(state=="list"){
			$(tabsId).permission();
		}
		if($.isBlank(state)||state=="list"){
			state="add";
		}
		var tabTitleMsg={add:"<%message "tmpl.core.comm.action.add"%>",edit:"<%message "tmpl.core.comm.action.edit"%>",show:"<%message "tmpl.core.comm.action.show"%>"};
		$("#{%.param%}TabTitle").html(tabTitleMsg[state]);
		return true;
	}
</script>
`
