import{d as q,E as r,$ as U,aw as X,aX as j,B as H,S as J,b0 as K,bb as O,aP as Q,bc as W,aT as Y,e as c,f as e,g as t,w as l,a2 as g,a3 as y,b1 as Z,r as ee,s as se,M as ae,o,h as E,ac as S,k as V,t as u,aN as te,_ as oe}from"./index.b4c32530.js";/* empty css               *//* empty css               *//* empty css               */import{p as ne}from"./business.6eb1e581.js";import{p as ce}from"./resource-server.7b0d7086.js";const _=i=>(ee("data-v-fc284d97"),i=i(),se(),i),le={class:"fields"},ie=_(()=>e("div",{class:"hr"},[e("span",{class:"icon"}),e("span",null,"\u53EF\u7528\u53D8\u91CF")],-1)),re={class:"header"},ue={class:"field-list"},_e={class:"item-row"},de=_(()=>e("span",{class:"label"},"\u5B57\u6BB5",-1)),pe={style:{"margin-right":"5px"}},ve={class:"item-row"},he=_(()=>e("span",{class:"label"},"\u8BF4\u660E",-1)),me={class:"value"},ge={class:"item-row"},ye=_(()=>e("span",{class:"label"},"\u5B57\u6BB5",-1)),fe={class:"keyword"},we={class:"item-row"},be=_(()=>e("span",{class:"label"},"\u8BF4\u660E",-1)),Be={class:"value"},ke={style:{"margin-right":"5px"}},De=q({__name:"right",props:{serverId:{type:Number,default:0}},setup(i){const d=i,p=r("business"),v=r(0),f=r([]),w=r([]),n=r({page:1,page_size:10,keyword:""}),b=s=>`\${${s}}`,N=async()=>{const{data:s}=await ce({...n.value,server_id:d.serverId});f.value=s.list,v.value=s.total},B=async()=>{const{data:s}=await ne({...n.value,server_id:d.serverId});w.value=s.list,v.value=s.total},k=()=>{if(!d.serverId){ae.error("\u8BF7\u5148\u9009\u62E9\u670D\u52A1");return}p.value==="business"?B():N()},R=s=>{n.value.page=s,k()},D=()=>{n.value.page=1,n.value.page_size=10,v.value=0,k()};return U(()=>d.serverId,s=>{s&&B()}),(s,F)=>{const T=X,A=j,G=H,L=J,I=K,$=Z,x=O,C=Q,M=W,P=Y,z=te("copy");return o(),c("div",le,[ie,e("div",re,[t(L,null,{default:l(()=>[t(T,{style:{width:"180px"},placeholder:"\u8BF7\u8F93\u5165\u53D8\u91CF\u540D\u79F0","allow-clear":""}),t(G,{type:"primary",onClick:D},{icon:l(()=>[t(A)]),_:1})]),_:1})]),t($,{modelValue:p.value,"onUpdate:modelValue":F[0]||(F[0]=a=>p.value=a),type:"button",style:{width:"100%",didplay:"flex"},onChange:D},{default:l(()=>[t(I,{class:"tag-item",value:"business"},{default:l(()=>[E("\u4E1A\u52A1\u53D8\u91CF")]),_:1}),t(I,{class:"tag-item",value:"resource"},{default:l(()=>[E("\u8D44\u6E90\u53D8\u91CF")]),_:1})]),_:1},8,["modelValue"]),e("div",ue,[p.value=="business"?(o(!0),c(g,{key:0},y(w.value,(a,h)=>(o(),c("div",{key:h,class:"item"},[e("div",_e,[de,S((o(),V(C,{size:"small",color:"arcoblue"},{default:l(()=>[e("span",pe,u(a.keyword),1),t(x)]),_:2},1024)),[[z,b(a.keyword)]])]),e("div",ve,[he,e("span",me,u(a.description),1)])]))),128)):(o(!0),c(g,{key:1},y(f.value,(a,h)=>(o(),c("div",{key:h,class:"item"},[e("div",ge,[ye,e("span",fe,u(a.keyword),1)]),e("div",we,[be,e("span",Be,u(a.description),1)]),t(M),(o(!0),c(g,null,y(a.fields.split(","),m=>(o(),c("div",{key:m,class:"item-left"},[S((o(),V(C,{size:"small",color:"arcoblue"},{default:l(()=>[e("span",ke,u(a.keyword+"."+m),1),t(x)]),_:2},1024)),[[z,b(a.keyword+"."+m)]])]))),128))]))),128))]),t(P,{size:"small",total:v.value,current:n.value.page,"page-size":n.value.page_size,"show-total":"",onChange:R},null,8,["total","current","page-size"])])}}});const Se=oe(De,[["__scopeId","data-v-fc284d97"]]);export{Se as default};