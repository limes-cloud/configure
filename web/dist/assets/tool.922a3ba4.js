import{f as q,o as R,p as U,q as H,_ as J}from"./index.3bd70e06.js";/* empty css               *//* empty css              *//* empty css               *//* empty css               *//* empty css               */import{d as K,e as x,a4 as M,aV as W,aW as Y,aZ as Z,bl as j,bq as G,aC as f,aI as o,br as O,A as u,aJ as n,bc as Q,aL as X,E as r,B as h,aK as k,C as ee,aM as B,aN as w,n as oe,bs as ne,aY as te,bt as se}from"./arco.9b5dfa89.js";import{l as g}from"./lodash.38496856.js";import{S as ae}from"./sortable.esm.41bdaeb5.js";import"./chart.a7f56037.js";import"./vue.b2201ea8.js";const ce={class:"action-icon"},le={class:"action-icon"},ie={id:"tableSetting"},re={style:{"margin-right":"4px",cursor:"move"}},ue={class:"title"},pe=K({__name:"tool",props:{columns:null,size:null},emits:["update:size","update:columns","add","refresh"],setup(C,{emit:p}){const D=C,l=x([]),d=x([]),y=t=>{p("update:size",t)};(t=>{l.value=g.exports.cloneDeep(t),l.value.forEach((e,c)=>{e.checked=!0}),d.value=g.exports.cloneDeep(l.value)})(D.columns);const b=(t,e,c,i=!1)=>{const s=i?g.exports.cloneDeep(t):t;return e>-1&&c>-1&&s.splice(e,1,s.splice(c,1,s[e]).pop()),s},V=t=>{t&&oe(()=>{const e=document.getElementById("tableSetting");new ae(e,{onEnd(c){const{oldIndex:i,newIndex:s}=c;b(l.value,i,s),b(d.value,i,s),p("update:columns",l.value)}})})},z=(t,e,c)=>{t?l.value.splice(c,0,e):l.value=d.value.filter(i=>i.checked),p("update:columns",l.value)};return(t,e)=>{const c=M,i=W,s=ne,E=q,m=Y,I=R,S=te,A=Z,F=U,L=H,N=j,T=G,P=O,$=se("permission");return u(),f(P,{style:{"margin-bottom":"16px","align-items":"center"}},{default:o(()=>[n(s,{span:12},{default:o(()=>[Q((u(),f(i,{type:"primary",onClick:e[0]||(e[0]=a=>p("add"))},{icon:o(()=>[n(c)]),default:o(()=>[X(" \u65B0\u5EFA\u670D\u52A1 ")]),_:1})),[[$,"configure:server:add"]])]),_:1}),n(s,{span:12,class:"tool"},{default:o(()=>[n(m,{content:"\u5237\u65B0"},{default:o(()=>[r("div",{class:"action-icon",onClick:e[1]||(e[1]=a=>p("refresh"))},[n(E,{size:"18"})])]),_:1}),n(A,{onSelect:y},{content:o(()=>[(u(!0),h(w,null,k(t.$densityList,(a,_)=>(u(),f(S,{key:_,value:_,class:ee({active:_===C.size})},{default:o(()=>[r("span",null,B(a),1)]),_:2},1032,["value","class"]))),128))]),default:o(()=>[n(m,{content:"\u5B57\u4F53\u5927\u5C0F"},{default:o(()=>[r("div",ce,[n(I,{size:"18"})])]),_:1})]),_:1}),n(m,{content:"\u884C\u8BBE\u7F6E"},{default:o(()=>[n(T,{trigger:"click",position:"bl",onPopupVisibleChange:V},{content:o(()=>[r("div",ie,[(u(!0),h(w,null,k(d.value,(a,_)=>(u(),h("div",{key:a.dataIndex,class:"setting"},[r("div",re,[n(L)]),r("div",null,[n(N,{modelValue:a.checked,"onUpdate:modelValue":v=>a.checked=v,onChange:v=>z(v,a,_)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),r("div",ue,B(a.title==="#"?"index":a.title),1)]))),128))])]),default:o(()=>[r("div",le,[n(F,{size:"18"})])]),_:1})]),_:1})]),_:1})]),_:1})}}});const we=J(pe,[["__scopeId","data-v-444cab39"]]);export{we as default};