(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2d21f81c"],{d9b3:function(t,e,a){"use strict";a.r(e);var l=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("el-container",[a("el-main",[a("el-row",[a("el-col",{attrs:{span:22,offset:1}},[a("el-card",[a("div",{staticClass:"indexsearch"},[a("el-form",{attrs:{model:t.query,size:"mini",inline:!0}},[a("el-form-item",{attrs:{label:""}},[a("el-input",{attrs:{placeholder:"日志编号",size:"mini",clearable:""},model:{value:t.query.name,callback:function(e){t.$set(t.query,"name",e)},expression:"query.name"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{attrs:{placeholder:"任务编号",size:"mini",clearable:""},model:{value:t.query.taskname,callback:function(e){t.$set(t.query,"taskname",e)},expression:"query.taskname"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{attrs:{placeholder:"数据库名",size:"mini",clearable:""},model:{value:t.query.dbname,callback:function(e){t.$set(t.query,"dbname",e)},expression:"query.dbname"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-button",{attrs:{type:"primary",icon:"el-icon-search",size:"mini"},on:{click:t.search}},[t._v("搜索")])],1)],1)],1)])],1)],1),a("el-row",[a("el-col",{attrs:{span:22,offset:1}},[a("el-table",{attrs:{data:t.tasks,border:"",size:"mini"}},[a("el-table-column",{attrs:{prop:"name",label:"日志编号",width:"150"}}),a("el-table-column",{attrs:{prop:"dbinfo.name",label:"任务编号",width:"150"}}),a("el-table-column",{attrs:{prop:"dbinfo.dbtype",label:"数据库类型",width:"140"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-tag",{attrs:{type:"primary",size:"mini"}},[t._v(t._s(e.row.dbinfo.dbtype))])]}}])}),a("el-table-column",{attrs:{prop:"status",label:"状态",width:"135"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-tag",{attrs:{size:"mini",type:t.$public.StatusTagFormat(e.row.status)}},[t._v(t._s(t.$public.StatusFormat(e.row.status))+t._s(3===e.row.status?"("+e.row.recoveryProgress+"%)":""))])]}}])}),a("el-table-column",{attrs:{prop:"localfilepath",label:"本地路径"},scopedSlots:t._u([{key:"default",fn:function(e){return[2===e.row.status&&1===e.row.status||!e.row.localfilepath||""===e.row.localfilepath?t._e():a("div",[t._v(" "+t._s(e.row.localfilepath)+" "),0===e.row.deleted?a("el-link",{attrs:{type:"primary",href:t.$api.httpUrl+t.$api.downloadfile+"/"+e.row.id,target:"_blank"},on:{"~click":function(e){return t.preventMultiClicks(e)}}},[t._v("下载")]):t._e()],1)]}}])}),a("el-table-column",{attrs:{prop:"created",label:"执行时间",width:"140"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("span",[t._v(t._s(t.$public.comTime(e.row.created)))])]}}])}),a("el-table-column",{attrs:{prop:"created",label:"是否本地保存",width:"100"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-tag",{attrs:{type:t.$public.tasklogifstatustagfmt(e.row.ifsavelocal),size:"mini"}},[t._v(t._s(t.$public.tasklogifstatusfmt(e.row.ifsavelocal)))])]}}])}),a("el-table-column",{attrs:{prop:"created",label:"异地存储记录",width:"140"},scopedSlots:t._u([{key:"default",fn:function(e){return[e.row.rlogs&&e.row.rlogs.length&&e.row.rlogs.length>0?a("div",t._l(e.row.rlogs,(function(e,l){return a("div",{key:l+"rlogs"},[a("el-tag",{attrs:{size:"mini"}},[t._v(t._s(e.name))])],1)})),0):t._e()]}}])}),a("el-table-column",{attrs:{prop:"created",label:"文件状态",width:"100"},scopedSlots:t._u([{key:"default",fn:function(e){return[2!==e.row.status||1!==e.row.status?a("div",[0===e.row.deleted?a("el-tag",{attrs:{type:"success",size:"mini"}},[t._v("存在")]):a("el-tag",{attrs:{type:"danger",size:"mini"}},[t._v("已删除")])],1):t._e()]}}])}),a("el-table-column",{attrs:{prop:"created",label:"压缩包加密",width:"100"},scopedSlots:t._u([{key:"default",fn:function(e){return[2!==e.row.status||1!==e.row.status?a("div",[1===e.row.passwordStatus?a("el-tag",{attrs:{type:"success",size:"mini"}},[t._v("有密码")]):a("el-tag",{attrs:{type:"primary",size:"mini"}},[t._v("无密码")])],1):t._e()]}}])}),a("el-table-column",{attrs:{prop:"recoveryStatus",label:"成功恢复次数",width:"100"}}),a("el-table-column",{attrs:{prop:"created",label:"操作",width:"100"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("div",[a("el-link",{attrs:{type:"primary",underline:!1},on:{click:function(a){return t.showlog(e.row.id)}}},[t._v("详情")])],1)]}}])})],1),a("div",{staticClass:"block"},[a("el-pagination",{attrs:{background:"",small:"","page-size":t.query.pageSize,"current-page":t.query.currentPage,layout:"total, prev, pager, next",total:t.total},on:{"current-change":t.current_change}})],1)],1)],1),t.showrlogdialog?a("el-dialog",{attrs:{title:"备份日志详情",width:"40%",visible:t.showrlogdialog},on:{"update:visible":function(e){t.showrlogdialog=e}}},[a("showlog",{attrs:{"show-data":t.formData},on:{close:t.dialogClose}})],1):t._e()],1)],1)},s=[],r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("el-row",[a("el-col",{attrs:{span:22,offset:1}},[a("el-form",{attrs:{model:t.showData,size:"mini","label-width":"120px"}},[a("el-form-item",{attrs:{label:"任务编号"}},[a("span",[t._v(t._s(t.showData.dbinfo.name))])]),a("el-form-item",{attrs:{label:"备份类型"}},[a("span",[t._v(t._s(t.showData.dbinfo.dbtype))])]),a("el-form-item",{attrs:{label:"备份文件路径"}},[a("span",[t._v(t._s(t.showData.localfilepath))])]),a("el-form-item",{attrs:{label:"时间"}},[a("span",[t._v(t._s(t.$public.comTime(t.showData.created)))])]),a("el-form-item",{attrs:{label:"状态"}},[a("span",[t._v(t._s(t.$public.StatusFormat(t.showData.status)))])]),a("el-form-item",{attrs:{label:"压缩包加密"}},[1===t.showData.passwordStatus?a("el-tag",{attrs:{type:"success",size:"mini"}},[t._v("是")]):a("el-tag",{attrs:{type:"primary",size:"mini"}},[t._v("否")])],1),1===t.showData.passwordStatus?a("el-form-item",{attrs:{label:"压缩包密码"}},[t._v(" "+t._s(t.showData.password)+" ")]):t._e(),0!=t.showData.status?a("el-form-item",{attrs:{label:"错误信息"}},[a("span",[t._v(t._s(t.showData.msg))])]):t._e(),t.showData.status<=0&&0===t.showData.deleted&&1===t.showData.recovery?a("el-form-item",{attrs:{label:"操作"}},[a("el-button",{attrs:{type:"warning"},on:{click:t.sqlrecovery}},[t._v("还原数据库")])],1):t._e(),3===t.showData.status?a("el-form-item",{attrs:{label:"还原进度"}},[t._v(t._s(t.showData.recoveryProgress)+"%")]):t._e(),t.showData.recoveryStatus>0?a("el-form-item",{attrs:{label:"还原数据库次数"}},[t._v(" "+t._s(t.showData.recoveryStatus)+" ")]):t._e(),a("el-form-item",{attrs:{label:"最后还原时间"}},[t._v(" "+t._s(t.$public.comTime(t.showData.recoveryTime))+" ")]),-1===t.showData.status?a("el-form-item",{attrs:{label:"恢复错误信息"}},[a("span",[t._v(t._s(t.showData.recoveryErrMsg))])]):t._e()],1)],1)],1)],1)},o=[],i={props:{showData:{type:Object,default:null}},data:function(){return{}},methods:{sqlrecovery:function(){var t=this;this.$http.Put(this.$api.sqlrecovery+"/"+this.showData.id,{},(function(e){console.log(e),t.$notify({title:"还原任务提示",message:e}),t.$emit("close")}),"调用数据还原任务")}}},n=i,u=a("2877"),c=Object(u["a"])(n,r,o,!1,null,null,null),p=c.exports,d={data:function(){return{showrlogdialog:!1,tasks:[],total:0,formData:{},query:{currentPage:1,pageSize:15,name:"",taskname:"",dbname:""}}},components:{showlog:p},methods:{preventMultiClicks:function(t){t.preventDefault(),this.downloadFile(t.target.href)},downloadFile:function(t){var e=document.createElement("a");e.href=t,e.style.display="none",document.body.appendChild(e),e.click(),document.body.removeChild(e)},dialogClose:function(){this.showrlogdialog=!1,this.taskGet()},showlog:function(t){var e=this;this.$http.Get(this.$api.log+"/"+t,{},(function(t){e.formData=t,e.showrlogdialog=!0}))},taskGet:function(){var t=this;this.$http.Get(this.$api.logs,this.query,(function(e){e&&(t.tasks=e.data,t.total=e.count)}))},search:function(){this.query.currentPage=1,this.taskGet()},current_change:function(t){this.query.currentPage=t,this.taskGet()}},created:function(){this.taskGet()}},m=d,f=Object(u["a"])(m,l,s,!1,null,null,null);e["default"]=f.exports}}]);