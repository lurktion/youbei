(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2d21ab75"],{bd63:function(t,e,a){"use strict";a.r(e);var l=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("el-container",[a("el-main",[a("el-row",[a("el-col",{attrs:{span:22,offset:1}},[a("el-table",{attrs:{data:t.tasks,border:"",size:"mini"}},[a("el-table-column",{attrs:{prop:"filename",label:"文件名"}}),a("el-table-column",{attrs:{prop:"created",label:"状态",width:"180"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("span",[t._v(t._s(t.$public.StatusFormat(e.row.status)))])]}}])}),a("el-table-column",{attrs:{prop:"filesize",label:"文件大小",width:"180"}}),a("el-table-column",{attrs:{prop:"packet",label:"分片数",width:"180"}}),a("el-table-column",{attrs:{prop:"created",label:"创建时间",width:"180"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("span",[t._v(t._s(t.$public.comTime(e.row.created)))])]}}])}),a("el-table-column",{attrs:{prop:"created",label:"操作",width:"180"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-link",{attrs:{type:"primary",underline:!1},on:{click:function(a){return t.showlog(e.row.id)}}},[t._v("查看")])]}}])})],1),a("div",{staticClass:"block"},[a("el-pagination",{attrs:{small:!0,background:"","page-size":t.pageSize,"page-count":t.currentPage,layout:"prev, pager, next",total:t.total},on:{"current-change":t.current_change}})],1)],1)],1),t.showrlogdialog?a("el-dialog",{attrs:{title:"Yserver日志详情",width:"40%",visible:t.showrlogdialog},on:{"update:visible":function(e){t.showrlogdialog=e}}},[a("showlocalstronge",{attrs:{"show-data":t.formData},on:{close:t.dialogClose}})],1):t._e()],1)],1)},o=[],s=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("el-row",[a("el-col",{attrs:{span:22,offset:1}},[a("el-form",{attrs:{model:t.showData,size:"mini","label-width":"120px"}},[a("el-form-item",{attrs:{label:"文件名"}},[a("span",[t._v(t._s(t.showData.filename))])]),a("el-form-item",{attrs:{label:"文件大小"}},[a("span",[t._v(t._s(t.showData.filesize))])]),a("el-form-item",{attrs:{label:"时间"}},[a("span",[t._v(t._s(t.$public.comTime(t.showData.created)))])]),a("el-form-item",{attrs:{label:"状态"}},[a("span",[t._v(t._s(t.$public.StatusFormat(t.showData.status)))])]),a("el-form-item",{attrs:{label:"文件切片数量"}},[a("span",[t._v(t._s(t.showData.packet))])])],1)],1)],1)],1)},n=[],r={props:{showData:{type:Object,default:null}},created:function(){},data:function(){return{}},methods:{}},i=r,c=a("2877"),u=Object(c["a"])(i,s,n,!1,null,null,null),p=u.exports,d={components:{showlocalstronge:p},data:function(){return{tasks:[],total:0,currentPage:1,pageSize:10,formData:{},showrlogdialog:!1}},methods:{dialogClose:function(){this.showrlogdialog=!0},showlog:function(t){var e=this;this.$http.Get(this.$api.uploadlog+"/"+t,{},(function(t){e.formData=t,e.showrlogdialog=!0}))},taskGet:function(){var t=this;this.$http.Get(this.$api.uploadlogs,{page:this.currentPage,count:this.pageSize},(function(e){e&&(t.tasks=e.data,t.total=e.count)}))},current_change:function(t){this.currentPage=t,this.taskGet()}},created:function(){this.taskGet()}},h=d,f=Object(c["a"])(h,l,o,!1,null,null,null);e["default"]=f.exports}}]);