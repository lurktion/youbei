(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-ec98825e"],{"3a57":function(t,e,r){"use strict";r.r(e);var o=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("el-container",[r("el-main",[r("el-row",[r("el-col",{attrs:{span:22,offset:1}},[r("adddialog",{ref:"adddialog",attrs:{flush:t.getremoteStorage}})],1)],1),r("el-row",[r("el-col",{attrs:{span:22,offset:1}},[r("el-table",{attrs:{data:t.ftps,border:"",size:"mini"}},[r("el-table-column",{attrs:{prop:"types",label:"类型"}}),r("el-table-column",{attrs:{prop:"name",label:"名称"}}),r("el-table-column",{attrs:{prop:"path",label:"文件位置"}}),r("el-table-column",{attrs:{prop:"host",label:"主机地址"}}),r("el-table-column",{attrs:{prop:"port",label:"主机端口"}}),r("el-table-column",{attrs:{prop:"created",label:"创建时间"},scopedSlots:t._u([{key:"default",fn:function(e){return[r("span",[t._v(t._s(t.$public.comTime(e.row.created)))])]}}])}),r("el-table-column",{attrs:{prop:"status",label:"操作"},scopedSlots:t._u([{key:"default",fn:function(e){return[r("el-link",{attrs:{type:"primary",underline:!1},on:{click:function(r){return t.editTask(e.row.id)}}},[t._v("修改")]),r("el-link",{attrs:{type:"danger",underline:!1},on:{click:function(r){return t.deleteTask(e.row.id)}}},[t._v("删除")])]}}])})],1),r("div",{staticClass:"block"},[r("el-pagination",{attrs:{small:!0,background:"","page-size":t.pageSize,"page-count":t.currentPage,layout:"prev, pager, next",total:t.total},on:{"current-change":t.current_change}})],1)],1)],1),r("edit",{ref:"edit",attrs:{flush:t.getremoteStorage}})],1)],1)},a=[],l=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",[r("el-button",{attrs:{size:"mini"},on:{click:t.open}},[t._v("添加异地存储")]),r("el-dialog",{attrs:{title:"新建异地存储",visible:t.dialog,modal:!1,width:"500px"},on:{"update:visible":function(e){t.dialog=e}}},[r("el-form",{ref:"ftpform",attrs:{model:t.ftpform,size:"mini","label-position":"right","label-width":"auto","validate-on-rule-change":!1}},[r("el-form-item",{attrs:{label:"类型",prop:"types",rules:[{required:!0,message:"请选择类型名称",trigger:"blur"}]}},[r("el-select",{attrs:{placeholder:"请选择"},model:{value:t.ftpform.types,callback:function(e){t.$set(t.ftpform,"types",e)},expression:"ftpform.types"}},t._l(t.ftptypes,(function(t){return r("el-option",{key:t.value,attrs:{label:t.value,value:t.value}})})),1)],1),r("el-form-item",{attrs:{label:"名称",prop:"name",rules:[{required:!0,message:"请输入名称",trigger:"blur"}]}},[r("el-input",{model:{value:t.ftpform.name,callback:function(e){t.$set(t.ftpform,"name",e)},expression:"ftpform.name"}})],1),r("el-form-item",{attrs:{label:"主机地址",prop:"host",rules:[{required:!0,message:"请输入主机地址",trigger:"blur"}]}},[r("el-input",{model:{value:t.ftpform.host,callback:function(e){t.$set(t.ftpform,"host",e)},expression:"ftpform.host"}})],1),r("el-form-item",{attrs:{label:"主机端口",prop:"port",rules:[{required:!0,message:"请输入主机端口",trigger:"blur"}]}},[r("el-input",{model:{value:t.ftpform.port,callback:function(e){t.$set(t.ftpform,"port",e)},expression:"ftpform.port"}})],1),r("el-form-item",{attrs:{label:"用户名",prop:"username"}},[r("el-input",{model:{value:t.ftpform.username,callback:function(e){t.$set(t.ftpform,"username",e)},expression:"ftpform.username"}})],1),r("el-form-item",{attrs:{label:"密码",prop:"password"}},[r("el-input",{attrs:{"show-password":""},model:{value:t.ftpform.password,callback:function(e){t.$set(t.ftpform,"password",e)},expression:"ftpform.password"}})],1),r("el-form-item",{attrs:{label:"路径",prop:"path",rules:[{required:!0,message:"请输入 路径",trigger:"blur"}]}},[r("el-input",{model:{value:t.ftpform.path,callback:function(e){t.$set(t.ftpform,"path",e)},expression:"ftpform.path"}})],1),r("el-form-item",{attrs:{label:"断线重连"}},[r("el-input-number",{attrs:{min:0,max:10},model:{value:t.ftpform.relink,callback:function(e){t.$set(t.ftpform,"relink",e)},expression:"ftpform.relink"}})],1),r("el-form-item",[r("el-button",{attrs:{type:"primary"},on:{click:function(e){return t.addFtp("ftpform")}}},[t._v("确 定")])],1)],1)],1)],1)},n=[],s=(r("a9e3"),{props:{flush:{type:Function,default:function(){}}},data:function(){return{ftptypes:[{value:"ftp"},{value:"sftp"},{value:"Yserver"}],ftpform:{relink:0},dialog:!1}},methods:{open:function(){this.ftpform={relink:0},this.dialog=!0},addFtp:function(t){var e=this;this.$refs[t].validate((function(t){if(!t)return!1;e.ftpform["port"]=Number(e.ftpform["port"]),e.$http.Post(e.$api.ftp,e.ftpform,(function(){e.dialog=!1,e.flush()}),"添加异地存储")}))}}}),i=s,p=r("2877"),f=Object(p["a"])(i,l,n,!1,null,null,null),u=f.exports,m=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",[r("el-drawer",{attrs:{title:"修改异地存储配置",visible:t.dialog,direction:"rtl",size:550,modal:!1},on:{"update:visible":function(e){t.dialog=e}}},[r("el-row",[r("el-col",{attrs:{span:22,offset:1}},[r("el-form",{ref:"ftpform",attrs:{model:t.ftpform,size:"mini","label-width":"auto","label-position":"right"}},[r("el-form-item",{attrs:{label:"类型",prop:"types",rules:[{required:!0,message:"请选择类型名称",trigger:"blur"}]}},[r("el-select",{attrs:{placeholder:"请选择"},model:{value:t.ftpform.types,callback:function(e){t.$set(t.ftpform,"types",e)},expression:"ftpform.types"}},t._l(t.ftptypes,(function(t){return r("el-option",{key:t.value,attrs:{label:t.value,value:t.value}})})),1)],1),r("el-form-item",{attrs:{label:"名称",prop:"name",rules:[{required:!0,message:"请输入名称",trigger:"blur"}]}},[r("el-input",{model:{value:t.ftpform.name,callback:function(e){t.$set(t.ftpform,"name",e)},expression:"ftpform.name"}})],1),r("el-form-item",{attrs:{label:"主机地址",prop:"host",rules:[{required:!0,message:"请输入主机地址",trigger:"blur"}]}},[r("el-input",{model:{value:t.ftpform.host,callback:function(e){t.$set(t.ftpform,"host",e)},expression:"ftpform.host"}})],1),r("el-form-item",{attrs:{label:"主机端口",prop:"port",rules:[{required:!0,message:"请输入主机端口",trigger:"blur"}]}},[r("el-input",{model:{value:t.ftpform.port,callback:function(e){t.$set(t.ftpform,"port",e)},expression:"ftpform.port"}})],1),r("el-form-item",{attrs:{label:"用户名",prop:"username"}},[r("el-input",{model:{value:t.ftpform.username,callback:function(e){t.$set(t.ftpform,"username",e)},expression:"ftpform.username"}})],1),r("el-form-item",{attrs:{label:"密码",prop:"password"}},[r("el-input",{attrs:{"show-password":""},model:{value:t.ftpform.password,callback:function(e){t.$set(t.ftpform,"password",e)},expression:"ftpform.password"}})],1),r("el-form-item",{attrs:{label:"路径",prop:"path",rules:[{required:!0,message:"请输入 路径",trigger:"blur"}]}},[r("el-input",{model:{value:t.ftpform.path,callback:function(e){t.$set(t.ftpform,"path",e)},expression:"ftpform.path"}})],1),r("el-form-item",{attrs:{label:"断线重连"}},[r("el-input-number",{attrs:{min:0,max:10},model:{value:t.ftpform.relink,callback:function(e){t.$set(t.ftpform,"relink",e)},expression:"ftpform.relink"}})],1),r("el-form-item",[r("el-button",{attrs:{type:"primary"},on:{click:function(e){return t.editFtp("ftpform")}}},[t._v("修 改")])],1)],1)],1)],1)],1)],1)},c=[],d={props:{flush:{type:Function,default:function(){}}},data:function(){return{remoteStorageId:"",ftptypes:[{value:"ftp"},{value:"sftp"},{value:"Yserver"}],ftpform:{relink:0},dialog:!1}},methods:{open:function(t){var e=this;this.remoteStorageId=t,this.$http.Get(this.$api.ftp+"/"+t,{},(function(t){e.ftpform=t,e.dialog=!0}))},editFtp:function(t){var e=this;this.$refs[t].validate((function(t){if(!t)return!1;e.ftpform["port"]=Number(e.ftpform["port"]),e.$http.Put(e.$api.ftp+"/"+e.remoteStorageId,e.ftpform,(function(){e.dialog=!1,e.flush()}),"修改异地存储")}))}}},b=d,g=Object(p["a"])(b,m,c,!1,null,null,null),h=g.exports,v={components:{adddialog:u,edit:h},data:function(){return{title:!1,total:0,currentPage:1,pageSize:10,ftps:[]}},created:function(){this.getremoteStorage()},methods:{editTask:function(t){this.$refs.edit.open(t)},getremoteStorage:function(){var t=this;this.$http.Get(this.$api.ftps,{page:this.currentPage,count:this.pageSize},(function(e){t.ftps=e.data,t.total=e.count}))},deleteTask:function(t){var e=this;this.$confirm("此操作将永久删除该配置, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((function(){e.$http.Delete(e.$api.ftp+"/"+t,{},(function(){e.getremoteStorage()}))})).catch((function(){e.$message({type:"info",message:"已取消删除"})}))},current_change:function(t){this.currentPage=t,this.getremoteStorage()}}},k=v,y=Object(p["a"])(k,o,a,!1,null,null,null);e["default"]=y.exports},7156:function(t,e,r){var o=r("861d"),a=r("d2bb");t.exports=function(t,e,r){var l,n;return a&&"function"==typeof(l=e.constructor)&&l!==r&&o(n=l.prototype)&&n!==r.prototype&&a(t,n),t}},a9e3:function(t,e,r){"use strict";var o=r("83ab"),a=r("da84"),l=r("94ca"),n=r("6eeb"),s=r("5135"),i=r("c6b6"),p=r("7156"),f=r("c04e"),u=r("d039"),m=r("7c73"),c=r("241c").f,d=r("06cf").f,b=r("9bf2").f,g=r("58a8").trim,h="Number",v=a[h],k=v.prototype,y=i(m(k))==h,$=function(t){var e,r,o,a,l,n,s,i,p=f(t,!1);if("string"==typeof p&&p.length>2)if(p=g(p),e=p.charCodeAt(0),43===e||45===e){if(r=p.charCodeAt(2),88===r||120===r)return NaN}else if(48===e){switch(p.charCodeAt(1)){case 66:case 98:o=2,a=49;break;case 79:case 111:o=8,a=55;break;default:return+p}for(l=p.slice(2),n=l.length,s=0;s<n;s++)if(i=l.charCodeAt(s),i<48||i>a)return NaN;return parseInt(l,o)}return+p};if(l(h,!v(" 0o1")||!v("0b1")||v("+0x1"))){for(var w,x=function(t){var e=arguments.length<1?0:t,r=this;return r instanceof x&&(y?u((function(){k.valueOf.call(r)})):i(r)!=h)?p(new v($(e)),r,x):$(e)},_=o?c(v):"MAX_VALUE,MIN_VALUE,NaN,NEGATIVE_INFINITY,POSITIVE_INFINITY,EPSILON,isFinite,isInteger,isNaN,isSafeInteger,MAX_SAFE_INTEGER,MIN_SAFE_INTEGER,parseFloat,parseInt,isInteger".split(","),I=0;_.length>I;I++)s(v,w=_[I])&&!s(x,w)&&b(x,w,d(v,w));x.prototype=k,k.constructor=x,n(a,h,x)}}}]);