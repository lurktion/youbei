(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2d225f38"],{e71a:function(t,a,e){"use strict";e.r(a);var l=function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("el-container",[e("el-main",[e("el-row",[e("el-col",{attrs:{span:22,offset:1}},[e("el-table",{attrs:{data:t.rlogs,border:"",size:"mini"}},[e("el-table-column",{attrs:{prop:"name",label:"记录编号",width:"150"}}),e("el-table-column",{attrs:{prop:"rsinfo.name",label:"异地存储名称",width:"180"}}),e("el-table-column",{attrs:{prop:"loginfo.localfilepath",label:"上传文件"}}),e("el-table-column",{attrs:{prop:"rsinfo.types",label:"上传方式",width:"180"}}),e("el-table-column",{attrs:{prop:"status",label:"状态",width:"180"},scopedSlots:t._u([{key:"default",fn:function(a){return[e("el-tag",{attrs:{size:"mini",type:t.$public.StatusTagFormat(a.row.status)}},[t._v(t._s(t.$public.StatusFormat(a.row.status)))])]}}])}),e("el-table-column",{attrs:{prop:"created",label:"创建时间",width:"180"},scopedSlots:t._u([{key:"default",fn:function(a){return[e("span",[t._v(t._s(t.$public.comTime(a.row.created)))])]}}])}),e("el-table-column",{attrs:{prop:"",label:"操作",width:"180"},scopedSlots:t._u([{key:"default",fn:function(a){return[e("el-link",{attrs:{type:"primary",underline:!1},on:{click:function(e){return t.showlog(a.row.id)}}},[t._v("查看")])]}}])})],1),e("div",{staticClass:"block"},[e("el-pagination",{attrs:{small:!0,background:"","page-size":t.pageSize,"current-page":t.currentPage,layout:"prev, pager, next",total:t.total},on:{"current-change":t.current_change}})],1)],1)],1),t.showrlogdialog?e("el-dialog",{attrs:{title:"异地上传日志详情",width:"40%",visible:t.showrlogdialog},on:{"update:visible":function(a){t.showrlogdialog=a}}},[e("showrlog",{attrs:{"rlog-id":t.showrlogId,"show-data":t.formData},on:{close:t.dialogClose}})],1):t._e()],1)],1)},o=[],s=function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("div",[e("el-row",[e("el-col",{attrs:{span:22,offset:1}},[e("el-form",{attrs:{model:t.showData,size:"mini","label-width":"120px"}},[e("el-form-item",{attrs:{label:"上传文件"}},[e("span",[t._v(t._s(t.showData.loginfo.localfilepath))])]),e("el-form-item",{attrs:{label:"异地存储名称"}},[e("span",[t._v(t._s(t.showData.rsinfo.name))])]),e("el-form-item",{attrs:{label:"异地存储类型"}},[e("span",[t._v(t._s(t.showData.rsinfo.types))])]),e("el-form-item",{attrs:{label:"时间"}},[e("span",[t._v(t._s(t.$public.comTime(t.showData.created)))])]),e("el-form-item",{attrs:{label:"状态"}},[e("span",[t._v(t._s(t.$public.StatusFormat(t.showData.status)))])]),e("el-form-item",{attrs:{label:"文件尺寸"}},[e("span",[t._v(t._s(t.showData.ysuploadfile.size))])]),e("el-form-item",{attrs:{label:"文件切片数量"}},[e("span",[t._v(t._s(t.showData.ysuploadfile.packetnum))])])],1)],1)],1)],1)},r=[],n={props:{showData:{type:Object,default:null}},data:function(){return{}},methods:{}},i=n,u=e("2877"),c=Object(u["a"])(i,s,r,!1,null,null,null),p=c.exports,h={data:function(){return{showrlogdialog:!1,showrlogId:"",rlogs:[],total:0,currentPage:1,pageSize:10,formData:{}}},components:{showrlog:p},methods:{dialogClose:function(){this.showrlogdialog=!1},showlog:function(t){var a=this;this.$http.Get(this.$api.rlog+"/"+t,{},(function(t){a.formData=t,a.showrlogdialog=!0}))},taskGet:function(){var t=this;this.$http.Get(this.$api.rlogs,{page:this.currentPage,count:this.pageSize},(function(a){a&&(t.rlogs=a.data,t.total=a.count)}))},current_change:function(t){this.currentPage=t,this.taskGet()}},created:function(){this.taskGet()}},f=h,d=Object(u["a"])(f,l,o,!1,null,null,null);a["default"]=d.exports}}]);