import{d as y,E as S,aO as T,aP as w,aQ as E,aR as D,S as $,aS as B,aT as P,k as l,w as e,o as r,ac as d,h as t,t as g,j as A,g as i,l as N,aN as O}from"./index.b4c32530.js";/* empty css               *//* empty css               *//* empty css              *//* empty css               *//* empty css               *//* empty css               */import{v as V}from"./type.a4a3f801.js";const J=y({__name:"table",props:{columns:null,loading:{type:Boolean},data:null,size:null,pagination:null,total:null},emits:["delete","update","add","set","resetToken","queryToken","pageChange"],setup(s,{emit:c}){const o=S({current:1,pageSize:10}),v=n=>{o.value.current=n,c("pageChange",o.value)},h=n=>{o.value.pageSize=n,c("pageChange",o.value)};return(n,j)=>{const m=T,u=w,C=E,k=D,_=$,z=B,b=P,p=O("permission");return r(),l(_,{direction:"vertical",fill:""},{default:e(()=>[d((r(),l(z,{"row-key":"id",loading:s.loading,columns:s.columns,pagination:!1,data:s.data,bordered:!1,size:s.size},{type:e(({record:a})=>[t(g(A(V)[a.type]),1)]),createdAt:e(({record:a})=>[t(g(n.$formatTime(a.created_at)),1)]),updatedAt:e(({record:a})=>[t(g(n.$formatTime(a.updated_at)),1)]),operations:e(({record:a})=>[i(_,{class:"cursor-pointer"},{default:e(()=>[d((r(),l(u,{color:"arcoblue",onClick:f=>c("set",a)},{icon:e(()=>[i(m)]),default:e(()=>[t(" \u914D\u7F6E ")]),_:2},1032,["onClick"])),[[p,"configure:business:value"]]),d((r(),l(u,{color:"orangered",onClick:f=>c("update",a)},{icon:e(()=>[i(m)]),default:e(()=>[t(" \u4FEE\u6539 ")]),_:2},1032,["onClick"])),[[p,"configure:business:update"]]),n.$hasPermission("configure:business:delete")?(r(),l(k,{key:0,content:"\u60A8\u786E\u8BA4\u5220\u9664\u6B64\u8D44\u6E90",type:"warning",onOk:f=>c("delete",a.id)},{default:e(()=>[i(u,{color:"red"},{icon:e(()=>[i(C)]),default:e(()=>[t(" \u5220\u9664 ")]),_:1})]),_:2},1032,["onOk"])):N("",!0)]),_:2},1024)]),_:1},8,["loading","columns","data","size"])),[[p,"configure:business:query"]]),i(b,{total:s.total,current:o.value.current,"page-size":o.value.pageSize,"show-total":"","show-jumper":"","show-page-size":"",onChange:v,onPageSizeChange:h},null,8,["total","current","page-size"])]),_:1})}}});export{J as _};