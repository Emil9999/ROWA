(function(t){function e(e){for(var n,r,i=e[0],l=e[1],c=e[2],u=0,p=[];u<i.length;u++)r=i[u],Object.prototype.hasOwnProperty.call(o,r)&&o[r]&&p.push(o[r][0]),o[r]=0;for(n in l)Object.prototype.hasOwnProperty.call(l,n)&&(t[n]=l[n]);d&&d(e);while(p.length)p.shift()();return s.push.apply(s,c||[]),a()}function a(){for(var t,e=0;e<s.length;e++){for(var a=s[e],n=!0,r=1;r<a.length;r++){var i=a[r];0!==o[i]&&(n=!1)}n&&(s.splice(e--,1),t=l(l.s=a[0]))}return t}var n={},r={app:0},o={app:0},s=[];function i(t){return l.p+"js/"+({}[t]||t)+"."+{"chunk-2d0b2c3a":"d4b8d628","chunk-33360dae":"059d1c5b","chunk-73393eee":"da37603b"}[t]+".js"}function l(e){if(n[e])return n[e].exports;var a=n[e]={i:e,l:!1,exports:{}};return t[e].call(a.exports,a,a.exports,l),a.l=!0,a.exports}l.e=function(t){var e=[],a={"chunk-33360dae":1,"chunk-73393eee":1};r[t]?e.push(r[t]):0!==r[t]&&a[t]&&e.push(r[t]=new Promise((function(e,a){for(var n="css/"+({}[t]||t)+"."+{"chunk-2d0b2c3a":"31d6cfe0","chunk-33360dae":"4d458ea3","chunk-73393eee":"d1742363"}[t]+".css",o=l.p+n,s=document.getElementsByTagName("link"),i=0;i<s.length;i++){var c=s[i],u=c.getAttribute("data-href")||c.getAttribute("href");if("stylesheet"===c.rel&&(u===n||u===o))return e()}var p=document.getElementsByTagName("style");for(i=0;i<p.length;i++){c=p[i],u=c.getAttribute("data-href");if(u===n||u===o)return e()}var d=document.createElement("link");d.rel="stylesheet",d.type="text/css",d.onload=e,d.onerror=function(e){var n=e&&e.target&&e.target.src||o,s=new Error("Loading CSS chunk "+t+" failed.\n("+n+")");s.code="CSS_CHUNK_LOAD_FAILED",s.request=n,delete r[t],d.parentNode.removeChild(d),a(s)},d.href=o;var v=document.getElementsByTagName("head")[0];v.appendChild(d)})).then((function(){r[t]=0})));var n=o[t];if(0!==n)if(n)e.push(n[2]);else{var s=new Promise((function(e,a){n=o[t]=[e,a]}));e.push(n[2]=s);var c,u=document.createElement("script");u.charset="utf-8",u.timeout=120,l.nc&&u.setAttribute("nonce",l.nc),u.src=i(t);var p=new Error;c=function(e){u.onerror=u.onload=null,clearTimeout(d);var a=o[t];if(0!==a){if(a){var n=e&&("load"===e.type?"missing":e.type),r=e&&e.target&&e.target.src;p.message="Loading chunk "+t+" failed.\n("+n+": "+r+")",p.name="ChunkLoadError",p.type=n,p.request=r,a[1](p)}o[t]=void 0}};var d=setTimeout((function(){c({type:"timeout",target:u})}),12e4);u.onerror=u.onload=c,document.head.appendChild(u)}return Promise.all(e)},l.m=t,l.c=n,l.d=function(t,e,a){l.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:a})},l.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},l.t=function(t,e){if(1&e&&(t=l(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var a=Object.create(null);if(l.r(a),Object.defineProperty(a,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var n in t)l.d(a,n,function(e){return t[e]}.bind(null,n));return a},l.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return l.d(e,"a",e),e},l.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},l.p="/",l.oe=function(t){throw console.error(t),t};var c=window["webpackJsonp"]=window["webpackJsonp"]||[],u=c.push.bind(c);c.push=e,c=c.slice();for(var p=0;p<c.length;p++)e(c[p]);var d=u;s.push([0,"chunk-vendors"]),a()})({0:function(t,e,a){t.exports=a("56d7")},"034f":function(t,e,a){"use strict";var n=a("2ac1"),r=a.n(n);r.a},"092d":function(t,e,a){"use strict";var n=a("7001"),r=a.n(n);r.a},"0dac":function(t,e,a){"use strict";var n=a("4c10"),r=a.n(n);r.a},"0f94":function(t,e,a){},"1d7a":function(t,e,a){},"1e82":function(t,e,a){"use strict";var n=a("1d7a"),r=a.n(n);r.a},"2ac1":function(t,e,a){},"35a2":function(t,e,a){},"3ecd":function(t,e,a){},"4a76":function(t,e,a){},"4c10":function(t,e,a){},"4d98":function(t,e,a){},"56d7":function(t,e,a){"use strict";a.r(e);a("e260"),a("e6cf"),a("cca6"),a("a79d");var n=a("a026"),r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("v-app",{attrs:{id:"app"}},[a("router-view")],1)},o=[],s={name:"App",components:{},data:function(){return{}}},i=s,l=(a("034f"),a("0c7c")),c=a("6544"),u=a.n(c),p=a("7496"),d=Object(l["a"])(i,r,o,!1,null,null,null),v=d.exports;u()(d,{VApp:p["a"]});a("d3b7");var f=a("8c4f"),h=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"Home"}},[a("v-container",[a("HomeTopRow"),t.farm_active?a("div",[a("v-row",{staticStyle:{margin:"30px 0 0 50px"}},[a("v-col",[a("v-chip",{attrs:{color:"transparent"}},[a("v-avatar",{staticStyle:{"margin-right":"7px"},attrs:{color:"primary",size:"16px"}}),a("span",[t._v("Ready to"),a("br"),t._v("harvest")])],1),a("v-chip",{attrs:{color:"transparent"}},[a("v-avatar",{staticStyle:{"margin-right":"7px"},attrs:{color:"accent",size:"16px"}}),a("span",[t._v("Ready to"),a("br"),t._v("plant")])],1)],1)],1),a("v-row",{staticStyle:{"padding-top":"40px"},attrs:{justify:"center"}},[a("CatTree")],1)],1):a("div",[a("StatGraphic")],1)],1),a("FarmInfo")],1)},m=[],g=(a("a4d3"),a("4de4"),a("4160"),a("accc"),a("0d03"),a("e439"),a("dbb4"),a("b64b"),a("159b"),a("2fa7")),b=a("bc3a"),y=a.n(b),_=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("v-row",[a("v-col",{staticStyle:{"padding-left":"30px"},attrs:{cols:"4"}},[a("v-btn",{attrs:{icon:"","x-large":""}},[a("v-icon",[t._v("$vuetify.icons.adminSettings")])],1)],1),a("v-col",{attrs:{"align-content-center":"",cols:"4"}},[a("ButtonPill")],1),a("v-col",{staticStyle:{"text-align":"right","padding-right":"30px"},attrs:{cols:"4",id:"logo-text"}},[t._v(" ROWA ")])],1)},w=[],x=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("v-btn-toggle",{staticClass:"shadow",attrs:{rounded:"",borderless:""}},[a("v-btn",{style:[t.farm_active?{"background-color":"var(--v-primary-base)",color:"white"}:{"background-color":"white",color:"var(--v-primary-base)"}],on:{click:t.change_dash_state}},[t._v(" Farm ")]),a("v-btn",{style:[t.farm_active?{"background-color":"white",color:"var(--v-primary-base)"}:{"background-color":"var(--v-primary-base)",color:"white"}],on:{click:t.change_dash_state}},[t._v(" Stats ")])],1)},S=[],O=a("2f62");function j(t,e){var a=Object.keys(t);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(t);e&&(n=n.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),a.push.apply(a,n)}return a}function P(t){for(var e=1;e<arguments.length;e++){var a=null!=arguments[e]?arguments[e]:{};e%2?j(a,!0).forEach((function(e){Object(g["a"])(t,e,a[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(a)):j(a).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(a,e))}))}return t}var C={name:"ButtonPill",computed:P({},Object(O["d"])(["farm_active"])),methods:P({},Object(O["b"])(["change_dash_state"]))},k=C,V=(a("e4fd"),a("8336")),E=a("a609"),D=Object(l["a"])(k,x,S,!1,null,"d1d1babc",null),I=D.exports;u()(D,{VBtn:V["a"],VBtnToggle:E["a"]});var H={name:"HomeTopRow",components:{ButtonPill:I}},L=H,A=(a("ae85"),a("62ad")),B=a("132d"),T=a("0fd9"),M=Object(l["a"])(L,_,w,!1,null,"09941a6d",null),$=M.exports;u()(M,{VBtn:V["a"],VCol:A["a"],VIcon:B["a"],VRow:T["a"]});var F=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("FarmTransition",{attrs:{"y-positions":t.yPositions}},[a("div",[a("v-row",{attrs:{justify:"center"}},[a("h2",[t._v("Farming")])]),a("v-row",{attrs:{justify:"center"}},[a("p",[t._v("Plant and Harvest your very own plant from this farm for your lunch, or to take home.")])]),a("v-row",{staticStyle:{"margin-top":"40px"},attrs:{justify:"center"}},[a("v-btn",{staticClass:"text-capitalize",attrs:{id:"button",rounded:"",color:"accent",height:"75",width:"360",to:{name:"Farming"}},on:{click:function(e){return e.stopPropagation(),t.hello(e)}}},[t._v(" Start Farming Now "),a("v-icon",{attrs:{right:"",dark:""}},[t._v("mdi-arrow-right")])],1)],1),a("v-row",{staticStyle:{padding:"0 80px"}},[a("v-col",{staticClass:"info-box"},[a("InfoBoxPlants",{attrs:{heading:"Harvestable",plants:t.harvestable}})],1),a("v-col",{staticClass:"info-box"},[a("InfoBoxPlants",{attrs:{heading:"Plantable",plants:t.plantable}})],1)],1)],1)])},N=[],R=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[n("div",{ref:"interactElement",staticClass:"moving-card",class:{isAnimating:t.isInteractAnimating},style:{transform:t.transformString}},[n("v-row",{attrs:{justify:"center"}},[n("img",{staticStyle:{"padding-top":"12px"},attrs:{src:a("9e20"),alt:"line"}})]),t._t("default")],2)])},G=[],z=(a("99af"),a("c975"),a("d81d"),a("284c")),U=a("fb3a"),q=a.n(U),J={name:"FarmTransition",props:{yPositions:Array},data:function(){return{interactPosition:{x:0,y:0},isInteractAnimating:!0}},methods:{interactSetPosition:function(t){var e=t.x,a=void 0===e?0:e,n=t.y,r=void 0===n?0:n;this.interactPosition={x:a,y:r}},stopCardAtPosition:function(){var t=this,e=Object(z["a"])(this.yPositions),a=e.map((function(e){return Math.abs(e-t.interactPosition.y)})),n=a.indexOf(Math.min.apply(Math,Object(z["a"])(a))),r=this.yPositions[n];this.interactSetPosition({x:0,y:r})}},computed:{transformString:function(){var t=this.interactPosition,e=t.x,a=t.y;return"translate3D(".concat(e,"px, ").concat(a,"px, 0)")}},mounted:function(){var t=this,e=this.$refs.interactElement;q()(e).draggable({lockAxis:"y",onstart:function(){t.isInteractAnimating=!1},onmove:function(e){var a=t.interactPosition.x+e.dx,n=t.interactPosition.y+e.dy;t.interactSetPosition({x:a,y:n})},onend:function(){t.isInteractAnimating=!0,t.stopCardAtPosition()}})},beforeDestroy:function(){q()(this.$refs.interactElement).unset()}},Z=J,K=(a("70af"),Object(l["a"])(Z,R,G,!1,null,"a6488d82",null)),W=K.exports;u()(K,{VRow:T["a"]});var Q=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("v-row",{attrs:{justify:"center"}},[a("h2",[t._v(t._s(t.heading))])]),a("v-row",{attrs:{justify:"center"}},[a("table",t._l(t.plants,(function(e){return a("tr",{key:e.plant_type},[a("td",[t._v(t._s(e.plant_type))]),a("td",[t._v(t._s(e.available_plants))])])})),0)])],1)},X=[],Y={name:"InfoBoxPlants",props:{heading:String,plants:Array},data:function(){return{}}},tt=Y,et=(a("0dac"),Object(l["a"])(tt,Q,X,!1,null,"657e6720",null)),at=et.exports;u()(et,{VRow:T["a"]});var nt={name:"FarmInfo",components:{FarmTransition:W,InfoBoxPlants:at},data:function(){return{yPositions:[0,-160],harvestable:null,plantable:null}},methods:{getHarvestablePlants:function(){var t=this;y.a.get("http://127.0.0.1:3000/dashboard/harvestable-plants").then((function(e){t.harvestable=e.data})).catch((function(t){console.log(t)}))},getPlantablePlants:function(){var t=this;y.a.get("http://127.0.0.1:3000/dashboard/plantable-plants").then((function(e){t.plantable=e.data})).catch((function(t){console.log(t)}))}},created:function(){this.getHarvestablePlants(),this.getPlantablePlants()}},rt=nt,ot=(a("092d"),Object(l["a"])(rt,F,N,!1,null,"3b093d82",null)),st=ot.exports;u()(ot,{VBtn:V["a"],VCol:A["a"],VIcon:B["a"],VRow:T["a"]});var it=a("5c24"),lt=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticStyle:{padding:"0 30px"}},[a("v-row",[a("v-col",{staticStyle:{padding:"15px"}},[a("StatBox",{attrs:{heading:"Temperature",type:"temperature",value:t.temperature}})],1),a("v-col",{staticStyle:{padding:"15px"}},[a("StatBox",{attrs:{heading:"Light Intensity",type:"light",value:t.light_intensity}})],1)],1),a("v-row",{attrs:{justify:"center"}},[a("p")])],1)},ct=[],ut=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"info-box"},[a("v-row",{attrs:{justify:"center"}},[a("h2",[t._v(t._s(t.heading))])]),a("v-row",{staticStyle:{"padding-top":"10px"},attrs:{justify:"center"}},["temperature"===t.type?a("VueSpeedometer",{attrs:{value:t.value,height:130,width:210,segments:5,segmentColors:["tomato","gold","limegreen","gold","tomato"],minValue:0,maxValue:40,textColor:"white",currentValueText:"${value} °C"}}):t._e(),"light"===t.type?a("VueSpeedometer",{attrs:{value:t.value,height:130,width:210,segments:1e3,minValue:0,maxValue:1e3,maxSegmentLabels:5,currentValueText:"${value} Lux",textColor:"white"}}):t._e()],1)],1)},pt=[],dt=(a("a9e3"),a("7840")),vt=a.n(dt),ft={name:"StatBox",components:{VueSpeedometer:vt.a},props:{heading:String,value:Number,type:String}},ht=ft,mt=(a("b94e"),Object(l["a"])(ht,ut,pt,!1,null,"bf3fd420",null)),gt=mt.exports;u()(mt,{VRow:T["a"]});var bt={name:"StatGraphic",components:{StatBox:gt},data:function(){return{temperature:null,light_intensity:null,interval:null}},methods:{getSensorData:function(){var t=this;y.a.get("http://localhost:3000/dashboard/sensor-data").then((function(e){t.temperature=e.data.temperature,t.light_intensity=e.data.light_intensity,console.log(e.data)}))}},mounted:function(){this.getSensorData(),this.interval=setInterval(this.getSensorData,3e4)},destroyed:function(){clearInterval(this.interval)}},yt=bt,_t=Object(l["a"])(yt,lt,ct,!1,null,"265e3dc7",null),wt=_t.exports;function xt(t,e){var a=Object.keys(t);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(t);e&&(n=n.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),a.push.apply(a,n)}return a}function St(t){for(var e=1;e<arguments.length;e++){var a=null!=arguments[e]?arguments[e]:{};e%2?xt(a,!0).forEach((function(e){Object(g["a"])(t,e,a[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(a)):xt(a).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(a,e))}))}return t}u()(_t,{VCol:A["a"],VRow:T["a"]});var Ot={name:"Home",components:{HomeTopRow:$,FarmInfo:st,CatTree:it["a"],StatGraphic:wt},data:function(){return{msg:"Farm Overview",harvestable_plants:null,all_plants:null,sensor_data:{datetime:null,temperature:null,light_intensity:null},sensor_data_updated:null,interval:null}},methods:{getSensorData:function(){var t=this;this.sensor_data_updated=(new Date).toISOString(),y.a.get("http://127.0.0.1:3000/dashboard/sensor-data").then((function(e){t.sensor_data=e.data[0],console.log(t.sensor_data)})).catch((function(t){console.log(t)}))},getIntervalData:function(){this.interval=setInterval(this.getSensorData,5e3)}},created:function(){this.getSensorData()},beforeDestroy:function(){clearInterval(this.interval)},computed:St({},Object(O["d"])(["farm_active"]))},jt=Ot,Pt=(a("eb42"),a("8212")),Ct=a("cc20"),kt=a("a523"),Vt=Object(l["a"])(jt,h,m,!1,null,"592d6d66",null),Et=Vt.exports;u()(Vt,{VAvatar:Pt["a"],VChip:Ct["a"],VCol:A["a"],VContainer:kt["a"],VRow:T["a"]});var Dt=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("div",{staticClass:"box"},[a("div",{staticClass:"moving-card",class:{up:t.farmInfo},on:{click:function(e){t.farmInfo=!t.farmInfo}}},[a("h1",[t._v("CSS3 slide up")]),a("p",[t._v("This is a quick demo of slide-up effect using CSS animation when hover the box. No JS required!")])])])])},It=[],Ht={name:"demo",data:function(){return{farmInfo:!1}}},Lt=Ht,At=(a("b32d"),Object(l["a"])(Lt,Dt,It,!1,null,"346448ae",null)),Bt=At.exports,Tt=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"Farming"}},[a("v-container",{attrs:{align:"center"}},[a("v-stepper",{staticClass:"step-number",model:{value:t.e1,callback:function(e){t.e1=e},expression:"e1"}},[a("v-stepper-items",[a("v-stepper-content",{attrs:{step:"1"}},[a("v-row",{attrs:{"align-center":""}},[a("v-col",{staticStyle:{"padding-left":"30px"},attrs:{cols:"2"}}),a("v-col",{attrs:{align:"center","align-self":"center",cols:"8"}},[a("h3",{staticStyle:{color:"#828282"}},[t._v("Choose what you would like to do")])]),a("v-col",{staticStyle:{"text-align":"right","padding-right":"30px"},attrs:{cols:"2"}},[a("v-btn",{attrs:{dark:"",fab:"",color:"white",to:{name:"Home"}}},[a("v-icon",{attrs:{color:"primary"}},[t._v("mdi-close")])],1)],1)],1)],1)],1),a("v-stepper-header",{staticClass:"step-number",attrs:{id:"header-steps"}},[a("v-stepper-step",{attrs:{complete:!0,step:""}},[t._v("Name of step 1")]),a("v-divider"),a("v-stepper-step",{attrs:{complete:t.e1>1,step:""}},[t._v("Name of step 2")]),a("v-divider"),a("v-stepper-step",{attrs:{complete:t.e1>1,step:""}},[t._v("Name of step 3")]),a("v-divider"),a("v-stepper-step",{attrs:{complete:t.e1>2,step:""}},[t._v("Name of step 4")]),a("v-divider"),a("v-stepper-step",{attrs:{complete:t.e1>3,step:""}},[t._v("Name of step 4")]),a("v-divider"),a("v-stepper-step",{attrs:{complete:t.e1>4,step:""}},[t._v("Name of step 4")])],1)],1),a("v-row",{attrs:{justify:"center"}},[a("h1",[t._v("Harvest a plant")])]),a("v-row",{attrs:{justify:"center"}},[a("p",[t._v("Go on a journey to become a famer.Harvest a plant now, for your daily dose of fresh greens.")])]),a("v-row",{staticClass:"info-box",attrs:{justify:"center",margin:"auto",padding:"auto"}},[a("InfoBoxPlants",{attrs:{heading:"Harvestable",plants:t.harvestable_plants}})],1),a("v-row",{staticStyle:{"margin-top":"40px"},attrs:{justify:"center"}},[a("v-btn",{attrs:{id:"button",disabled:t.HarvestDisable,to:{name:"Harvest"},rounded:"",color:"accent",height:"75",width:"360"}},[t._v(" Start Harvesting "),a("v-icon",{attrs:{right:"",dark:""}},[t._v("mdi-arrow-right")])],1)],1),a("v-row",{attrs:{justify:"center"}},[a("h1",[t._v("Plant a new Plant")])]),a("v-row",{attrs:{justify:"center"}},[a("p",[t._v("Go on a journey to become a famer.Plant a new seed now and enjoy fresh greens in the future.")])]),a("v-row",{staticClass:"info-box",attrs:{justify:"center",margin:"auto",padding:"auto"}},[a("InfoBoxPlants",{attrs:{heading:"Plantable",plants:t.plantable_plants}})],1),a("v-row",{staticStyle:{"margin-top":"40px"},attrs:{justify:"center"}},[a("v-btn",{attrs:{id:"button",disabled:t.PlantDisable,to:{name:"Plant"},rounded:"",color:"accent",height:"75",width:"360"}},[t._v(" Start Planting "),a("v-icon",{attrs:{right:"",dark:""}},[t._v("mdi-arrow-right")])],1)],1)],1)],1)},Mt=[],$t={name:"Farming",components:{InfoBoxPlants:at},data:function(){return{harvestingIsDisabled:!1,step:0,e1:0,harvestable_plants:null,plantable_plants:null}},methods:{getHarvestablePlants:function(){var t=this;y.a.get("http://127.0.0.1:3000/dashboard/harvestable-plants").then((function(e){t.harvestable_plants=e.data,console.log(t.harvestable_plants)})).catch((function(t){console.log(t)}))},getPlantablePlants:function(){var t=this;y.a.get("http://127.0.0.1:3000/dashboard/plantable-plants").then((function(e){t.plantable_plants=e.data,console.log(t.plantable_plants)})).catch((function(t){console.log(t)}))}},created:function(){this.getHarvestablePlants(),this.getPlantablePlants()},computed:{PlantDisable:function(){return null==this.plantable_plants},HarvestDisable:function(){return null==this.harvestable_plants}}},Ft=$t,Nt=(a("6e0f"),a("ce7e")),Rt=a("7e85"),Gt=a("e516"),zt=a("9c54"),Ut=a("56a4"),qt=Object(l["a"])(Ft,Tt,Mt,!1,null,"9c84120a",null),Jt=qt.exports;u()(qt,{VBtn:V["a"],VCol:A["a"],VContainer:kt["a"],VDivider:Nt["a"],VIcon:B["a"],VRow:T["a"],VStepper:Rt["a"],VStepperContent:Gt["a"],VStepperHeader:zt["a"],VStepperItems:zt["b"],VStepperStep:Ut["a"]}),n["a"].use(f["a"]);var Zt=[{path:"/",name:"Home",component:Et},{path:"/test",name:"test",component:Bt},{path:"/plant",name:"Plant",component:function(){return a.e("chunk-33360dae").then(a.bind(null,"e8cb"))}},{path:"/harvest",name:"Harvest",component:function(){return a.e("chunk-73393eee").then(a.bind(null,"ccb6"))}},{path:"/plant/:pos",name:"PlantingInstructions",component:function(){return a.e("chunk-2d0b2c3a").then(a.bind(null,"2600"))}},{path:"/farming",name:"Farming",component:Jt}],Kt=new f["a"]({mode:"history",base:"/",routes:Zt}),Wt=Kt;n["a"].use(O["a"]);var Qt=new O["a"].Store({state:{farm_active:!0,farm_info:{1:{type:"lettuce",pos:{1:{age:3,harvestable:!1},2:{age:10,harvestable:!1},3:{age:17,harvestable:!1},4:{age:24,harvestable:!1},5:{age:31,harvestable:!1},6:{age:38,harvestable:!1}}},2:{type:"lettuce",pos:{1:{age:3,harvestable:!1},2:{age:10,harvestable:!1},3:{age:17,harvestable:!1},4:{age:24,harvestable:!1},5:{age:31,harvestable:!1},6:{age:42,harvestable:!0}}},3:{type:"basil",pos:{1:{age:3,harvestable:!1},2:{age:10,harvestable:!1},3:{age:17,harvestable:!1},4:{age:24,harvestable:!1},5:{age:31,harvestable:!1},6:{age:42,harvestable:!0}}},4:{type:"basil",pos:{1:!1,2:{age:10,harvestable:!1},3:{age:17,harvestable:!1},4:{age:24,harvestable:!1},5:{age:31,harvestable:!1},6:{age:42,harvestable:!1}}},5:{type:"lettuce",pos:{1:{age:3,harvestable:!1},2:{age:10,harvestable:!1},3:{age:17,harvestable:!1},4:{age:24,harvestable:!1},5:{age:31,harvestable:!1},6:{age:42,harvestable:!0}}},6:{type:"lettuce",pos:{1:!1,2:{age:10,harvestable:!1},3:{age:17,harvestable:!1},4:{age:24,harvestable:!1},5:{age:31,harvestable:!1},6:{age:42,harvestable:!1}}}}},mutations:{CHANGE_DASH_STATE:function(t){t.farm_active=!t.farm_active}},actions:{change_dash_state:function(t){var e=t.commit;e("CHANGE_DASH_STATE")}},modules:{}}),Xt=a("f309"),Yt=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("svg",{attrs:{width:"47",height:"46",viewBox:"0 0 47 46",fill:"none",xmlns:"http://www.w3.org/2000/svg"}},[a("path",{attrs:{d:"M39.2206 15.4064L42.2617 9.39323L37.0181 4.20792L30.9373 7.21523C30.3411 6.92955 29.7317 6.67932 29.1156 6.46629L26.9664 0H19.5508L17.4012 6.46629C16.7851 6.67932 16.1758 6.92955 15.5796 7.21523L9.49878 4.20792L4.2552 9.39323L7.2963 15.4064C7.00741 15.996 6.75437 16.5986 6.53895 17.2079L0 19.3333V26.6664L6.53895 28.7921C6.75437 29.4014 7.00741 30.004 7.2963 30.5936L4.2552 36.6068L9.49878 41.7921L15.5796 38.7848C16.1758 39.0704 16.7851 39.3207 17.4012 39.5337L19.5505 46H26.966L29.1156 39.5337C29.7317 39.3207 30.3411 39.0704 30.9373 38.7848L37.0181 41.7921L42.2617 36.6068L39.2206 30.5936C39.5094 30.004 39.7625 29.4014 39.9779 28.7921L46.5169 26.6667V19.3336L39.9779 17.2079C39.7625 16.5986 39.5094 15.9957 39.2206 15.4064ZM31.4352 32.4336H15.0816V29.7383C15.0816 25.2794 18.7495 21.6523 23.2584 21.6523C20.2521 21.6523 17.8072 19.2346 17.8072 16.2617C17.8072 13.2888 20.2521 10.8711 23.2584 10.8711C26.2647 10.8711 28.7096 13.2888 28.7096 16.2617C28.7096 19.2346 26.2647 21.6523 23.2584 21.6523C27.7674 21.6523 31.4352 25.2794 31.4352 29.7383V32.4336Z",fill:"#789659"}})])},te=[],ee={name:"AdminSettings"},ae=ee,ne=Object(l["a"])(ae,Yt,te,!1,null,"0af0d50a",null),re=ne.exports;n["a"].use(Xt["a"]);var oe=new Xt["a"]({theme:{themes:{light:{primary:"#789659",secondary:"#F4F0E6",accent:"#E3927B"}},options:{customProperties:!0}},icons:{values:{adminSettings:{component:re}}}}),se=a("0086"),ie=a.n(se);n["a"].use(ie.a),n["a"].config.productionTip=!1,new n["a"]({router:Wt,store:Qt,vuetify:oe,render:function(t){return t(v)}}).$mount("#app")},"5c24":function(t,e,a){"use strict";var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticStyle:{position:"relative",width:"410px",height:"522px"}},[a("svg",{attrs:{width:"410",height:"522",viewBox:"0 0 410 522",fill:"none",xmlns:"http://www.w3.org/2000/svg"}},[a("rect",{attrs:{x:"190",width:"30",height:"494",fill:"#828282"}}),a("path",{attrs:{d:"M0 522L0 494H410V522H0Z",fill:"#828282"}}),a("rect",{attrs:{width:"190",height:"7",transform:"matrix(-1 0 0 1 190 0)",fill:"#828282"}}),a("rect",{attrs:{width:"190",height:"7",transform:"matrix(-1 0 0 1 410 0)",fill:"#828282"}})]),a("div",{staticClass:"module",staticStyle:{bottom:"371px"}},[a("Module",{attrs:{id:1,reverse:!1}})],1),a("div",{staticClass:"module",staticStyle:{bottom:"325px",right:"0px"}},[a("Module",{attrs:{id:2,reverse:!0}})],1),a("div",{staticClass:"module",staticStyle:{bottom:"247px"}},[a("Module",{attrs:{id:3,reverse:!1}})],1),a("div",{staticClass:"module",staticStyle:{bottom:"191px",right:"0"}},[a("Module",{attrs:{id:4,reverse:!0}})],1),a("div",{staticClass:"module",staticStyle:{bottom:"123px"}},[a("Module",{attrs:{id:5,reverse:!1}})],1),a("div",{staticClass:"module",staticStyle:{bottom:"77px",right:"0"}},[a("Module",{attrs:{id:6,reverse:!0}})],1)])},r=[],o=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticStyle:{position:"relative",width:"190px",height:"160px"}},[a("div",{staticStyle:{position:"absolute",bottom:"20px",width:"190px",height:"80px"}},t._l(t.positions,(function(e){return a("div",{key:e,staticStyle:{width:"15%",display:"inline-block",height:"50px"}},[t.module_plants.pos[e]?a("img",{attrs:{height:t.calculate_img_size(t.module_plants.pos[e].age),src:t.getImgUrl(t.module_plants.type)}}):a("div",{staticStyle:{display:"inline-block",width:"26px"}})])})),0),a("div",{staticStyle:{position:"absolute",bottom:"0"}},[a("svg",{attrs:{width:"190",height:"48",viewBox:"0 0 190 48",fill:"none",xmlns:"http://www.w3.org/2000/svg"}},[a("rect",{attrs:{width:"189.999",height:"41",transform:"matrix(-1 0 0 1 190 0)",fill:"#C4C4C4"}}),a("rect",{attrs:{width:"189.999",height:"7",transform:"matrix(-1 0 0 1 190 41)",fill:"#828282"}})])]),a("div",{staticStyle:{position:"absolute",bottom:"30px",width:"210px",height:"80px"}},t._l(t.positions,(function(e){return a("div",{key:e,staticStyle:{width:"15%",display:"inline-flex",height:"50px","z-index":"2","align-items":"center","justify-content":"center"}},[t.module_plants.pos[e].harvestable?a("span",{staticClass:"dot",style:[t.module_plants.pos[e].harvestable?{"background-color":"#789659"}:{}]}):t._e(),t.module_plants.pos[e]?t._e():a("span",{staticClass:"dot",style:[{"background-color":"#E3927B"}]})])})),0)])},s=[],i=(a("a4d3"),a("4de4"),a("4160"),a("26e9"),a("e439"),a("dbb4"),a("b64b"),a("159b"),a("2fa7")),l=a("2f62");function c(t,e){var a=Object.keys(t);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(t);e&&(n=n.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),a.push.apply(a,n)}return a}function u(t){for(var e=1;e<arguments.length;e++){var a=null!=arguments[e]?arguments[e]:{};e%2?c(a,!0).forEach((function(e){Object(i["a"])(t,e,a[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(a)):c(a).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(a,e))}))}return t}var p={name:"Module",props:["id","reverse"],computed:u({},Object(l["d"])(["farm_info"]),{module_plants:function(){return this.farm_info[this.$props.id]},positions:function(){return this.reverse?Object.keys(this.module_plants.pos):Object.keys(this.module_plants.pos).reverse()}}),methods:u({},Object(l["c"])(["get_module"]),{getImgUrl:function(t){return a("9f45")("./"+t+".svg")},calculate_img_size:function(t){var e;return e=t<7?10:t>36?50:t/42*50,e}})},d=p,v=(a("73ee"),a("0c7c")),f=Object(v["a"])(d,o,s,!1,null,"6aa34a62",null),h=f.exports,m={name:"CatTree",components:{Module:h}},g=m,b=(a("1e82"),Object(v["a"])(g,n,r,!1,null,"0973860d",null));e["a"]=b.exports},"6e0f":function(t,e,a){"use strict";var n=a("0f94"),r=a.n(n);r.a},7001:function(t,e,a){},"70af":function(t,e,a){"use strict";var n=a("fa84"),r=a.n(n);r.a},"73ee":function(t,e,a){"use strict";var n=a("4a76"),r=a.n(n);r.a},"9e20":function(t,e,a){t.exports=a.p+"img/rectangle_28.e6fa5003.svg"},"9f45":function(t,e,a){var n={"./Module.svg":"ac05","./basil.svg":"fdca","./lettuce.svg":"d483"};function r(t){var e=o(t);return a(e)}function o(t){if(!a.o(n,t)){var e=new Error("Cannot find module '"+t+"'");throw e.code="MODULE_NOT_FOUND",e}return n[t]}r.keys=function(){return Object.keys(n)},r.resolve=o,t.exports=r,r.id="9f45"},a04c:function(t,e,a){},ac05:function(t,e,a){t.exports=a.p+"img/Module.89b9f5f8.svg"},ae85:function(t,e,a){"use strict";var n=a("4d98"),r=a.n(n);r.a},b32d:function(t,e,a){"use strict";var n=a("a04c"),r=a.n(n);r.a},b94e:function(t,e,a){"use strict";var n=a("35a2"),r=a.n(n);r.a},d483:function(t,e,a){t.exports=a.p+"img/lettuce.9cf41183.svg"},e4fd:function(t,e,a){"use strict";var n=a("3ecd"),r=a.n(n);r.a},eb42:function(t,e,a){"use strict";var n=a("f99d"),r=a.n(n);r.a},f99d:function(t,e,a){},fa84:function(t,e,a){},fdca:function(t,e,a){t.exports=a.p+"img/basil.89b81ae8.svg"}});
//# sourceMappingURL=app.903e218e.js.map