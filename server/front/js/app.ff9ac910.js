(function(t){function e(e){for(var s,n,a=e[0],u=e[1],c=e[2],d=0,p=[];d<a.length;d++)n=a[d],Object.prototype.hasOwnProperty.call(o,n)&&o[n]&&p.push(o[n][0]),o[n]=0;for(s in u)Object.prototype.hasOwnProperty.call(u,s)&&(t[s]=u[s]);l&&l(e);while(p.length)p.shift()();return i.push.apply(i,c||[]),r()}function r(){for(var t,e=0;e<i.length;e++){for(var r=i[e],s=!0,a=1;a<r.length;a++){var u=r[a];0!==o[u]&&(s=!1)}s&&(i.splice(e--,1),t=n(n.s=r[0]))}return t}var s={},o={app:0},i=[];function n(e){if(s[e])return s[e].exports;var r=s[e]={i:e,l:!1,exports:{}};return t[e].call(r.exports,r,r.exports,n),r.l=!0,r.exports}n.m=t,n.c=s,n.d=function(t,e,r){n.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:r})},n.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},n.t=function(t,e){if(1&e&&(t=n(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var r=Object.create(null);if(n.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var s in t)n.d(r,s,function(e){return t[e]}.bind(null,s));return r},n.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="/";var a=window["webpackJsonp"]=window["webpackJsonp"]||[],u=a.push.bind(a);a.push=e,a=a.slice();for(var c=0;c<a.length;c++)e(a[c]);var l=u;i.push([0,"chunk-vendors"]),r()})({0:function(t,e,r){t.exports=r("56d7")},"034f":function(t,e,r){"use strict";var s=r("85ec"),o=r.n(s);o.a},"56d7":function(t,e,r){"use strict";r.r(e);r("e260"),r("e6cf"),r("cca6"),r("a79d");var s=r("2b0e"),o=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",[r("div",{staticClass:"layout"},[r("Layout",[r("Content",{style:{padding:"0 50px"}},[r("div",[r("Breadcrumb",{style:{margin:"16px 0"}},[r("BreadcrumbItem",[t._v("Status")])],1),r("div",{staticClass:"div-status"},[r("Row",{attrs:{gutter:16,type:"flex",justify:"start"}},t._l(this.serverStatus,(function(e){return r("i-col",{key:e.name,staticStyle:{"margin-bottom":"15px"},attrs:{span:"24"}},[r("Card",[r("Row",{style:{margin:"10px 0"},attrs:{type:"flex",justify:"center",align:"middle"}},[r("i-col",{attrs:{span:"4"}},[r("h3",[r("Icon",{attrs:{type:"logo-buffer",size:32,color:t.nowTime-1e3*e.lastReportTime>1e3*t.timeLimit?"#ed4014":"#19be6b"}}),t._v(t._s(e.name)+" ")],1)]),r("i-col",{attrs:{span:"20"}},[r("p",[t._v(" "+t._s(e.description)+" ")])])],1),r("Row",{style:{margin:"15px 0"},attrs:{type:"flex",justify:"start",align:"middle"}},[r("i-col",{attrs:{span:"6"}},[r("center",[r("p",[t._v("IP Address")]),r("Tag",{attrs:{color:"default",size:"medium"}},[t._v(t._s(e.ip))])],1)],1),r("i-col",{attrs:{span:"6"}},[r("center",[r("p",[t._v("CPU Load")]),r("Tag",{attrs:{color:"default",size:"medium"}},[t._v(t._s(e.cpuUsage)+"%")])],1)],1),r("i-col",{attrs:{span:"6"}},[r("center",[r("p",[t._v("Memory Usage")]),r("Tag",{attrs:{color:"default",size:"medium"}},[t._v(t._s(e.memUsage>=1024?(e.memUsage/1024).toFixed(2):e.memUsage)+" / "+t._s(e.memUsage>=1024?(e.maxMemLimit/1024).toFixed(2):e.maxMemLimit)+" "+t._s(e.memUsage>=1024?"GB":"MB"))])],1)],1),r("i-col",{attrs:{span:"6"}},[r("center",[r("p",[t._v("Storage Usage")]),r("Tag",{attrs:{color:"default",size:"medium"}},[t._v(t._s(e.diskUsage>=1024?(e.diskUsage/1024).toFixed(2):e.diskUsage)+" / "+t._s(e.diskUsage>=1024?(e.maxDiskLimit/1024).toFixed(2):e.maxDiskLimit)+" "+t._s(e.diskUsage>=1024?"GB":"MB"))])],1)],1)],1),r("Row",{style:{margin:"15px 0"},attrs:{type:"flex",justify:"start",align:"middle"}},[r("i-col",{attrs:{span:"6"}},[r("center",[r("p",[t._v("Network I/O")]),r("Tag",{attrs:{color:"default",size:"medium"}},[t._v(t._s(e.inBound.toFixed(2))+" / "+t._s(e.outBound.toFixed(2))+" Mbps")])],1)],1),r("i-col",{attrs:{span:"6"}},[r("center",[r("p",[t._v("Bandwidth Usage")]),r("Tag",{attrs:{color:"default",size:"medium"}},[t._v(t._s(e.inBoundTotalUsage>=1024?(e.inBoundTotalUsage/1024).toFixed(2):e.inBoundTotalUsage.toFixed(2))+" / "+t._s(e.inBoundTotalUsage>=1024?(e.outBoundTotalUsage/1024).toFixed(2):e.outBoundTotalUsage.toFixed(2))+" "+t._s(e.inBoundTotalUsage>=1024?"GB":"MB"))])],1)],1),r("i-col",{attrs:{span:"12"}},[r("center",[r("p",[t._v("Service Status")]),t._l(e.serviceStatus,(function(s,o){return r("Tag",{key:o,attrs:{type:"dot",color:1==s?t.nowTime-1e3*e.lastReportTime>3e4?"error":"success":"error"}},[t._v(t._s(o))])}))],2)],1)],1),r("Divider"),r("Row",{attrs:{type:"flex",justify:"center",align:"middle"}},[r("i-col",{attrs:{span:"24"}},[t._v(" Update Time "),r("Time",{attrs:{time:1e3*e.lastReportTime,type:"datetime"}}),t._v(" ("),r("Time",{attrs:{time:e.lastReportTime,interval:1}}),t._v(") ")],1)],1)],1)],1)})),1),t.spinShow?r("Spin",{attrs:{size:"large",fix:""}}):t._e()],1)],1)]),r("Footer",{staticClass:"layout-footer-center"},[t._v("Server Status")])],1)],1)])},i=[],n=(r("d3b7"),r("25f0"),{data:function(){return{serverStatus:null,spinShow:!0,msgNetworkError:!1,msgNetworkSuccess:!1,lastReportError:!1,timer:null,nowTime:null,timeLimit:window.g.timeLimit}},mounted:function(){this.getServerStatus(),this.timer=setInterval(this.getServerStatus,1e3*window.g.timeDuration)},beforeDestroy:function(){clearInterval(this.timer)},methods:{getServerStatus:function(){var t=this;this.$axios.post(window.g.apiUrl+"/status",{}).then((function(e){t.serverStatus=e.data["status"],t.spinShow=!1,1==t.lastReportError&&(t.lastReportError=!1,t.showMsgNetworkSuccess())})).catch((function(e){t.$Message.error(e.toString()),t.showMsgNetworkError(),t.lastReportError=!0})),this.nowTime=(new Date).getTime()},showMsgNetworkError:function(){var t=this;0==this.msgNetworkError&&(this.msgNetworkError=!0,this.$Notice.error({title:"Network error",desc:"Can't connect to server, retry in next duration.",top:50,duration:0,onClose:function(){t.msgNetworkError=!1}}))},showMsgNetworkSuccess:function(){var t=this;0==this.msgNetworkSuccess&&(this.msgNetworkSuccess=!0,this.$Notice.success({title:"Retry success",desc:"Connect to server success.",duration:0,onClose:function(){t.msgNetworkSuccess=!1}}))},showUnixTime:function(t){var e=1e3*t,r=new Date(e);return r.toLocaleString()}}}),a=n,u=(r("034f"),r("2877")),c=Object(u["a"])(a,o,i,!1,null,null,null),l=c.exports,d=r("82ae"),p=r.n(d),m=r("f825"),f=r.n(m);r("f8ce");s["default"].use(f.a),s["default"].config.productionTip=!1,s["default"].prototype.$axios=p.a,new s["default"]({render:function(t){return t(l)}}).$mount("#app")},"85ec":function(t,e,r){}});
//# sourceMappingURL=app.ff9ac910.js.map