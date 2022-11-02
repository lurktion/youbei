(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-6da46dca"],{1148:function(t,e,i){"use strict";var n=i("a691"),a=i("1d80");t.exports="".repeat||function(t){var e=String(a(this)),i="",s=n(t);if(s<0||s==1/0)throw RangeError("Wrong number of repetitions");for(;s>0;(s>>>=1)&&(e+=e))1&s&&(i+=e);return i}},"408a":function(t,e,i){var n=i("c6b6");t.exports=function(t){if("number"!=typeof t&&"Number"!=n(t))throw TypeError("Incorrect invocation");return+t}},"42b7":function(t,e,i){"use strict";i("5839")},4601:function(t,e,i){"use strict";i.r(e);var n=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("div",[i("el-row",[i("el-col",{attrs:{span:22,offset:1}},[i("el-row",{attrs:{gutter:20}},[i("el-col",{attrs:{span:8}},[i("el-card",{staticClass:"card-css"},[i("div",{staticClass:"div1"},[i("cpu",{ref:"cpu",attrs:{cpuinfo:t.cpuinfo}})],1)])],1),i("el-col",{attrs:{span:8}},[i("el-card",{staticClass:"card-css"},[i("div",{staticClass:"div1"},[i("mem",{ref:"mem",attrs:{meminfo:t.meminfo}})],1)])],1),i("el-col",{attrs:{span:8}},[i("el-card",{staticClass:"card-css"},[i("div",{staticClass:"div1"},[i("disk",{ref:"disk",attrs:{diskinfo:t.diskinfo}})],1)])],1)],1),i("el-row",{attrs:{gutter:20}},[i("el-col",{attrs:{span:12}},[i("el-card",[i("div",{attrs:{slot:"header"},slot:"header"},[i("span",[t._v("系统信息")])]),i("div",[t._v(" 主机名称:"+t._s(t.hostinfo.name)+" ")]),i("el-divider"),i("div",[t._v(" 系统类型:"+t._s(t.hostinfo.os)+" ")]),i("el-divider"),i("div",[t._v(" 架构:"+t._s(t.hostinfo.arch)+" ")]),i("el-divider"),i("div",[t._v(" 操作系统版本:"+t._s(t.hostinfo.plate)+" ")]),i("el-divider"),i("div",[t._v(" 主机id:"+t._s(t.hostinfo.id)+" ")])],1)],1),i("el-col",{attrs:{span:12}},[i("el-card",[i("div",{attrs:{slot:"header"},slot:"header"},[i("span",[t._v("任务统计")])]),i("div",[t._v(" 已完成:"+t._s(t.taskcounts.task_ok)+" ")]),i("el-divider"),i("div",[t._v(" 进行中:"+t._s(t.taskcounts.task_ing)+" ")]),i("el-divider"),i("div",[t._v(" 失败:"+t._s(t.taskcounts.task_err)+" ")])],1)],1)],1)],1)],1)],1)},a=[],s=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("div",[i("div",{style:{width:"270px",height:"270px"},attrs:{id:"MemChart"}}),i("span",{staticStyle:{align:"center"}},[t._v(t._s((t.meminfo.used/1024/1024/1024).toFixed(2))+"/"+t._s((t.meminfo.total/1024/1024/1024).toFixed(2))+"(GB)")])])},r=[],o=(i("b680"),{name:"mem",props:{meminfo:{type:Object,default:function(){return{percent:49,used:0,total:0}}}},data:function(){return{myChart:{},option:{title:{text:"内存使用率",left:"center"},tooltip:{formatter:"{a} <br/>{b} : {c}%"},series:[{name:"mem",type:"gauge",detail:{formatter:"{value}%"},data:[{value:49,name:"使用率"}]}]}}},mounted:function(){this.draw();var t=this;setInterval((function(){t.getnewdata()}),2e3)},methods:{draw:function(){var t=this.$echarts.init(document.getElementById("MemChart"));t.setOption(this.option),this.myChart=t},getnewdata:function(){"undefined"!==typeof this.meminfo.percent&&(this.option.series[0].data[0].value=this.meminfo.percent.toFixed(2),this.myChart.setOption(this.option,!0))}}}),c=o,d=i("2877"),u=Object(d["a"])(c,s,r,!1,null,null,null),l=u.exports,f=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("div",[i("div",{style:{width:"270px",height:"270px"},attrs:{id:"DiskChart"}}),i("span",[t._v(t._s((t.diskinfo.used/1024/1024/1024).toFixed(2))+"/"+t._s((t.diskinfo.total/1024/1024/1024).toFixed(2))+"(GB)")])])},h=[],p={name:"disk",props:{diskinfo:{type:Object,default:function(){return{percent:49,used:0,total:0}}}},data:function(){return{myChart:{},option:{title:{text:"磁盘空间",left:"center"},tooltip:{formatter:"{a} <br/>{b} : {c}%"},series:[{name:"mem",type:"gauge",detail:{formatter:"{value}%"},data:[{value:49,name:"使用率"}]}]}}},mounted:function(){this.draw()},methods:{draw:function(){var t=this.$echarts.init(document.getElementById("DiskChart"));t.setOption(this.option),this.myChart=t},getnewdata:function(){"undefined"!==typeof this.diskinfo.percent&&(this.option.series[0].data[0].value=this.diskinfo.percent.toFixed(2),this.myChart.setOption(this.option,!0))}}},m=p,v=Object(d["a"])(m,f,h,!1,null,null,null),w=v.exports,g=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("div",[i("div",{style:{width:"270px",height:"270px"},attrs:{id:"CpuChart"}}),i("span",[t._v("核心数:"+t._s(t.cpuinfo.cpucounts)+"核")])])},_=[],k={name:"cpu",props:{cpuinfo:{type:Object,default:function(){return{percent:49,used:0,total:0}}}},data:function(){return{msg:"",myChart:{},option:{title:{text:"CPU使用率",left:"center"},tooltip:{formatter:"{a} <br/>{b} : {c}%"},series:[{name:"cpu",type:"gauge",detail:{formatter:"{value}%"},data:[{value:49,name:"使用率"}]}]}}},mounted:function(){this.draw()},methods:{draw:function(){var t=this.$echarts.init(document.getElementById("CpuChart"));t.setOption(this.option),this.myChart=t},getnewdata:function(){"undefined"!==typeof this.cpuinfo.percent&&(this.option.series[0].data[0].value=this.cpuinfo.percent.toFixed(2),this.myChart.setOption(this.option,!0))}}},x=k,y=Object(d["a"])(x,g,_,!1,null,null,null),b=y.exports,C={components:{cpu:b,disk:w,mem:l},data:function(){return{taskcounts:{},meminfo:{},hostinfo:{},cpuinfo:{},diskinfo:{},resData:{task:100},Interval:Object}},mounted:function(){this.getnewdata(),this.taskcount(),this.setIntervalfunc()},methods:{setIntervalfunc:function(){var t=this;this.Interval=setInterval((function(){t.getnewdata()}),3e3)},taskcount:function(){var t=this;this.$http.Get(this.$api.dashboard,{},(function(e){t.taskcounts=e}))},getnewdata:function(){var t=this;this.$http.Get(this.$api.sysinfo,{},(function(e){t.meminfo=e.meminfo,t.hostinfo=e.hostinfo,t.diskinfo=e.diskinfo,t.cpuinfo=e.cpuinfo,t.$refs["cpu"].getnewdata(),t.$refs["mem"].getnewdata(),t.$refs["disk"].getnewdata()}))}},destroyed:function(){clearInterval(this.Interval)}},O=C,$=(i("42b7"),Object(d["a"])(O,n,a,!1,null,null,null));e["default"]=$.exports},5839:function(t,e,i){},b680:function(t,e,i){"use strict";var n=i("23e7"),a=i("a691"),s=i("408a"),r=i("1148"),o=i("d039"),c=1..toFixed,d=Math.floor,u=function(t,e,i){return 0===e?i:e%2===1?u(t,e-1,i*t):u(t*t,e/2,i)},l=function(t){var e=0,i=t;while(i>=4096)e+=12,i/=4096;while(i>=2)e+=1,i/=2;return e},f=c&&("0.000"!==8e-5.toFixed(3)||"1"!==.9.toFixed(0)||"1.25"!==1.255.toFixed(2)||"1000000000000000128"!==(0xde0b6b3a7640080).toFixed(0))||!o((function(){c.call({})}));n({target:"Number",proto:!0,forced:f},{toFixed:function(t){var e,i,n,o,c=s(this),f=a(t),h=[0,0,0,0,0,0],p="",m="0",v=function(t,e){var i=-1,n=e;while(++i<6)n+=t*h[i],h[i]=n%1e7,n=d(n/1e7)},w=function(t){var e=6,i=0;while(--e>=0)i+=h[e],h[e]=d(i/t),i=i%t*1e7},g=function(){var t=6,e="";while(--t>=0)if(""!==e||0===t||0!==h[t]){var i=String(h[t]);e=""===e?i:e+r.call("0",7-i.length)+i}return e};if(f<0||f>20)throw RangeError("Incorrect fraction digits");if(c!=c)return"NaN";if(c<=-1e21||c>=1e21)return String(c);if(c<0&&(p="-",c=-c),c>1e-21)if(e=l(c*u(2,69,1))-69,i=e<0?c*u(2,-e,1):c/u(2,e,1),i*=4503599627370496,e=52-e,e>0){v(0,i),n=f;while(n>=7)v(1e7,0),n-=7;v(u(10,n,1),0),n=e-1;while(n>=23)w(1<<23),n-=23;w(1<<n),v(1,1),w(2),m=g()}else v(0,i),v(1<<-e,0),m=g()+r.call("0",f);return f>0?(o=m.length,m=p+(o<=f?"0."+r.call("0",f-o)+m:m.slice(0,o-f)+"."+m.slice(o-f))):m=p+m,m}})}}]);