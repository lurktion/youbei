(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-18075cb4"],{"25ef":function(e,t,a){"use strict";a.r(t);var l=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("el-container",[a("el-main",[a("el-row",[a("el-col",{attrs:{span:22,offset:1}},[a("el-button",{attrs:{size:"mini"},on:{click:e.opendialog}},[e._v("添加任务")])],1)],1),a("el-row",[a("el-col",{attrs:{span:22,offset:1}},[a("el-table",{attrs:{data:e.sshtasks,border:"",size:"mini"}},[a("el-table-column",{attrs:{prop:"host",label:"主机地址"}}),a("el-table-column",{attrs:{prop:"dbhost",label:"数据库地址"}}),a("el-table-column",{attrs:{prop:"dbname",label:"数据库名"}}),a("el-table-column",{attrs:{prop:"crontab",label:"计划任务"}}),a("el-table-column",{attrs:{prop:"savepath",label:"保存地址"}}),a("el-table-column",{attrs:{prop:"created",label:"创建时间"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(e.$public.comTime(t.row.created)))])]}}])}),a("el-table-column",{attrs:{prop:"",label:"操作"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-link",{attrs:{type:"primary",underline:!1},on:{click:function(a){return e.runJob(t.row.id)}}},[e._v("立即执行")]),a("el-link",{attrs:{type:"primary",underline:!1},on:{click:function(a){return e.openedit(t.row.id)}}},[e._v("修改")]),a("el-link",{attrs:{type:"primary",underline:!1},on:{click:function(a){return e.deleteTask(t.row.id)}}},[e._v("删除")])]}}])})],1)],1)],1),e.dialogFormVisible?a("el-dialog",{attrs:{title:e.AddTaskTitle?"修改任务":"添加任务",size:"mini",visible:e.dialogFormVisible},on:{"update:visible":function(t){e.dialogFormVisible=t}}},[a("addsshtask",{attrs:{"task-id":e.taskId,dialogtitle:e.AddTaskTitle},on:{close:e.addclose}})],1):e._e()],1)],1)},s=[],i=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("el-container",[a("el-main",[a("el-row",[a("el-col",{attrs:{span:24}},[a("el-form",{attrs:{model:e.formData,"label-width":"120px",size:"mini","label-position":"right"}},[a("el-alert",{attrs:{title:"提示:当前远程备份只支持ssh远程和mysql数据库",type:"warning",center:"","show-icon":""}}),a("el-divider",[e._v("SSH配置")]),a("el-form-item",{attrs:{label:"主机地址"}},[a("el-input",{attrs:{placeholder:"主机地址"},model:{value:e.formData.host,callback:function(t){e.$set(e.formData,"host",t)},expression:"formData.host"}})],1),a("el-form-item",{attrs:{label:"SSH端口"}},[a("el-input",{attrs:{placeholder:"端口"},model:{value:e.formData.sshport,callback:function(t){e.$set(e.formData,"sshport",t)},expression:"formData.sshport"}})],1),a("el-form-item",{attrs:{label:"SSH用户名"}},[a("el-input",{attrs:{placeholder:"用户名"},model:{value:e.formData.sshuser,callback:function(t){e.$set(e.formData,"sshuser",t)},expression:"formData.sshuser"}})],1),a("el-form-item",{attrs:{label:"SSH密码"}},[a("el-input",{attrs:{placeholder:"密码"},model:{value:e.formData.sshpwd,callback:function(t){e.$set(e.formData,"sshpwd",t)},expression:"formData.sshpwd"}})],1),a("el-divider",[e._v("MySQL配置")]),a("el-form-item",{attrs:{label:"SQL主机地址"}},[a("el-input",{attrs:{placeholder:"主机地址"},model:{value:e.formData.dbhost,callback:function(t){e.$set(e.formData,"dbhost",t)},expression:"formData.dbhost"}})],1),a("el-form-item",{attrs:{label:"SQL端口"}},[a("el-input",{attrs:{placeholder:"端口"},model:{value:e.formData.dbport,callback:function(t){e.$set(e.formData,"dbport",t)},expression:"formData.dbport"}})],1),a("el-form-item",{attrs:{label:"SQL用户名"}},[a("el-input",{attrs:{placeholder:"用户名"},model:{value:e.formData.dbuser,callback:function(t){e.$set(e.formData,"dbuser",t)},expression:"formData.dbuser"}})],1),a("el-form-item",{attrs:{label:"SQL密码"}},[a("el-input",{attrs:{placeholder:"密码"},model:{value:e.formData.dbpwd,callback:function(t){e.$set(e.formData,"dbpwd",t)},expression:"formData.dbpwd"}})],1),a("el-form-item",{attrs:{label:"数据库名"}},[a("el-input",{attrs:{placeholder:"数据库名"},model:{value:e.formData.dbname,callback:function(t){e.$set(e.formData,"dbname",t)},expression:"formData.dbname"}})],1),a("el-divider",[e._v("任务配置")]),a("el-form-item",{attrs:{label:"本地保存"}},[a("el-switch",{attrs:{"active-text":"保存","inactive-text":"不保存"},on:{change:e.savepathfunc},model:{value:e.localsave,callback:function(t){e.localsave=t},expression:"localsave"}})],1),e.localsave?a("div",[a("el-form-item",{attrs:{label:"存储路径"}},[a("el-input",{attrs:{disabled:!0},model:{value:e.formData.savepath,callback:function(t){e.$set(e.formData,"savepath",t)},expression:"formData.savepath"}},[a("el-button",{attrs:{slot:"append",icon:"el-icon-search"},on:{click:function(t){return e.enableDirSearch("savepath")}},slot:"append"})],1)],1),a("el-form-item",{attrs:{label:"是否加密"}},[a("el-switch",{model:{value:e.enablezippwd,callback:function(t){e.enablezippwd=t},expression:"enablezippwd"}})],1),e.enablezippwd?a("el-form-item",{attrs:{label:"压缩密码"}},[a("el-input",{attrs:{"show-password":""},model:{value:e.formData.zippwd,callback:function(t){e.$set(e.formData,"zippwd",t)},expression:"formData.zippwd"}})],1):e._e(),a("el-form-item",{attrs:{label:"异地存储"}},[a("el-select",{attrs:{multiple:""},model:{value:e.formData.rs,callback:function(t){e.$set(e.formData,"rs",t)},expression:"formData.rs"}},e._l(e.remotestorages,(function(e){return a("el-option",{key:e.id,attrs:{label:e.name,value:e.id}})})),1)],1)],1):e._e(),a("el-form-item",{attrs:{label:"执行频率"}},[a("el-select",{on:{change:e.changepinlv},model:{value:e.pinlv,callback:function(t){e.pinlv=t},expression:"pinlv"}},e._l(e.pinlvs,(function(e){return a("el-option",{key:e.value,attrs:{label:e.label,value:e.value}})})),1)],1),"month"==e.pinlv?a("el-form-item",{attrs:{label:"日期"}},[a("el-select",{attrs:{placeholder:"请选择"},on:{change:e.cjobs},model:{value:e.day,callback:function(t){e.day=t},expression:"day"}},e._l(e.days,(function(e){return a("el-option",{key:e.value,attrs:{label:e.label,value:e.value}})})),1)],1):e._e(),"week"==e.pinlv?a("el-form-item",{attrs:{label:"周"}},[a("el-select",{attrs:{placeholder:"请选择"},on:{change:e.cjobs},model:{value:e.week,callback:function(t){e.week=t},expression:"week"}},e._l(e.weeks,(function(e){return a("el-option",{key:e.value,attrs:{label:e.label,value:e.value}})})),1)],1):e._e(),a("el-form-item",{attrs:{label:"时间点"}},[a("el-time-picker",{attrs:{placeholder:"任意时间点","value-format":"H m s"},on:{change:e.createjobs},model:{value:e.times,callback:function(t){e.times=t},expression:"times"}})],1),a("el-form-item",{attrs:{label:"计划任务",prop:"crontab",rules:[{required:!0,message:"请输入 计划任务",trigger:"blur"}]}},[a("el-input",{model:{value:e.formData.crontab,callback:function(t){e.$set(e.formData,"crontab",t)},expression:"formData.crontab"}})],1),a("el-form-item",{attrs:{label:"过期删除"}},[a("el-switch",{on:{change:e.expirefunc},model:{value:e.expire,callback:function(t){e.expire=t},expression:"expire"}})],1),e.expire?a("el-form-item",{attrs:{label:"过期"}},[a("el-input-number",{attrs:{min:0,max:100},model:{value:e.formData.expire,callback:function(t){e.$set(e.formData,"expire",t)},expression:"formData.expire"}}),e._v("天 ")],1):e._e(),a("el-form-item",{attrs:{label:""}},[""==e.taskId?a("el-button",{attrs:{type:"primary"},on:{click:e.addsshtask}},[e._v("添加")]):a("el-button",{attrs:{type:"primary"},on:{click:e.updatesshtask}},[e._v("修改")])],1)],1)],1)],1),e.dialogDirVisible?a("el-dialog",{attrs:{title:"请选择",visible:e.dialogDirVisible,width:"80%","append-to-body":""},on:{"update:visible":function(t){e.dialogDirVisible=t}}},[a("dirlist",{attrs:{dialogDirVisible:e.dialogDirVisible,"key-name":e.keyName},on:{close:e.changeDir}})],1):e._e()],1)],1)},o=[],r=(a("b0c0"),a("9794")),n={components:{Dirlist:r["a"]},data:function(){return{formData:{crontab:""},localsave:!1,keyName:"",expire:!1,dialogDirVisible:!1,enablezippwd:!1,remotestorages:[],day:"1",week:"1",times:"",pinlv:"day",pinlvs:[{value:"day",label:"天"},{value:"week",label:"周"},{value:"month",label:"月"}],weeks:[{value:"1",label:"周一"},{value:"2",label:"周二"},{value:"3",label:"周三"},{value:"4",label:"周四"},{value:"5",label:"周五"},{value:"6",label:"周六"},{value:"0",label:"周日"}],days:[{value:"1",label:"1"},{value:"2",label:"2"},{value:"3",label:"3"},{value:"4",label:"4"},{value:"5",label:"5"},{value:"6",label:"6"},{value:"7",label:"7"},{value:"8",label:"8"},{value:"9",label:"9"},{value:"10",label:"10"},{value:"11",label:"11"},{value:"12",label:"12"},{value:"13",label:"13"},{value:"14",label:"14"},{value:"15",label:"15"},{value:"16",label:"16"},{value:"17",label:"17"},{value:"18",label:"18"},{value:"19",label:"19"},{value:"20",label:"20"},{value:"21",label:"21"},{value:"22",label:"22"},{value:"23",label:"23"},{value:"24",label:"24"},{value:"25",label:"25"},{value:"26",label:"26"},{value:"27",label:"27"},{value:"28",label:"28"},{value:"29",label:"29"},{value:"30",label:"30"},{value:"31",label:"31"},{value:"L",label:"最后一天"}]}},props:{dialogtitle:{type:Boolean,default:!1},taskId:{type:String,default:""}},created:function(){""!==this.taskId&&this.getsshtask(this.taskId),this.ftpGet()},methods:{getsshtask:function(e){var t=this;this.$http.Get(this.$api.sshtask+"/"+e,{},(function(e){""===e.savepath?t.localsave=!1:t.localsave=!0,""===t.zippwd?t.enablezippwd=!1:t.enablezippwd=!0,console.log(t.enablezippwd=!1),t.formData=e}))},savepathfunc:function(e){e||(this.formData.savepath=""),console.log(e)},expirefunc:function(){this.formData.expire=0},createjobs:function(e){this.times=e,this.cjobs()},changeDir:function(e){this.formData[e.name]=e.value,this.dialogDirVisible=!1},enableDirSearch:function(e){this.keyName=e,this.dialogDirVisible=!0},ftpGet:function(){var e=this;this.$http.Get(this.$api.ftps,{page:1,count:-1},(function(t){console.log(t),e.remotestorages=t.data}))},changepinlv:function(e){this.pinlv=e,this.cjobs()},cjobs:function(){var e=this.times.split(" "),t="";3===e.length&&(t=e[2]+" "+e[1]+" "+e[0],"month"===this.pinlv&&this.day?t=t+" "+this.day+" * *":"week"===this.pinlv&&this.week?t=t+" * * "+this.week:t+=" * * *"),this.formData.crontab=t},addsshtask:function(){var e=this;this.$http.Post(this.$api.sshtask,this.formData,(function(){e.$emit("close")}))},updatesshtask:function(){var e=this;!1===this.localsave&&(this.formData.rs=[],this.formData.savepath="",this.formData.zippwd=""),this.$http.Put(this.$api.sshtask+"/"+this.taskId,this.formData,(function(){e.$emit("close")}))}}},c=n,u=a("2877"),p=Object(u["a"])(c,i,o,!1,null,null,null),d=p.exports,m={components:{addsshtask:d},data:function(){return{sshtasks:[],taskId:"",AddTaskTitle:!1,dialogFormVisible:!1}},created:function(){this.getsshtasks()},methods:{opendialog:function(){this.$router.push({path:"addsshtask",meta:{title:"添加任务1"}})},getsshtasks:function(){var e=this;this.$http.Get(this.$api.sshtasks,{},(function(t){e.sshtasks=t.data}))},addclose:function(){this.dialogFormVisible=!1,this.getsshtasks()},runJob:function(e){this.$http.Put(this.$api.runsshjob+"/"+e,{},(function(){}))},openedit:function(e){this.$router.push({path:"addsshtask",query:{id:e},meta:{title:"修改任务"}})},deleteTask:function(e){var t=this;this.$http.Delete(this.$api.sshtask+"/"+e,{},(function(){t.getsshtasks()}))}}},b=m,h=Object(u["a"])(b,l,s,!1,null,null,null);t["default"]=h.exports},9794:function(e,t,a){"use strict";var l=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",[a("div",{staticClass:"mydiv"},[a("el-tree",{ref:"tree",attrs:{props:e.props,load:e.loadNode,"check-on-click-node":!0,lazy:"","node-key":"path","show-checkbox":"","check-strictly":!0},on:{check:e.nodeClick}})],1),a("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[a("el-button",{attrs:{type:"primary",size:"mini"},on:{click:e.clickDirPath}},[e._v("确 定")])],1)])},s=[],i={props:{keyName:{type:String,default:""}},data:function(){return{input:"",props:{label:"label",children:"children",isLeaf:"isLeaf",path:"path"}}},methods:{clickDirPath:function(){var e={name:this.keyName,value:this.input};this.$emit("close",e)},loadNode:function(e,t){0===e.level?this.$http.Get(this.$api.dirlist,{},(function(e){return t(e)})):this.$http.Get(this.$api.dirlist,{dir:e.data.path},(function(e){return t(e)}))},nodeClick:function(e){this.$refs.tree.setCheckedNodes([e]),this.input=e.path}}},o=i,r=a("2877"),n=Object(r["a"])(o,l,s,!1,null,null,null);t["a"]=n.exports}}]);