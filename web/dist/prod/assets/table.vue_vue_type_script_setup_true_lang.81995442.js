import"./index.69f3790f.js";/* empty css               *//* empty css               *//* empty css              *//* empty css               *//* empty css               *//* empty css               */import{d as w,e as B,bw as D,aW as E,ae as S,ad as T,bx as $,by as A,bz as N,aC as r,aI as e,A as t,be as m,aJ as n,B as P,aK as x,aL as s,aM as _,u as F,aN as I,aD as V,bv as q}from"./arco.79b3a20c.js";import{l as L}from"./lodash.8c41d9ef.js";const U=w({__name:"table",props:{columns:null,loading:{type:Boolean},data:null,size:null,pagination:null,total:null},emits:["delete","update","add","set","resetToken","queryToken","pageChange"],setup(l,{emit:c}){const i=B({current:1,pageSize:10}),v=o=>{i.value.current=o,c("pageChange",i.value)},h=o=>{i.value.pageSize=o,c("pageChange",i.value)};return(o,O)=>{const u=D,p=E,f=S,C=T,k=$,z=A,b=N,g=q("permission");return t(),r(p,{direction:"vertical",fill:""},{default:e(()=>[m((t(),r(z,{"row-key":"id",loading:l.loading,columns:l.columns,pagination:!1,data:l.data,bordered:!1,size:l.size},{fields:e(({record:a})=>[n(p,null,{default:e(()=>[(t(!0),P(I,null,x(F(L.exports.split)(a.fields,","),(d,y)=>(t(),r(u,{key:y},{default:e(()=>[s(_(d),1)]),_:2},1024))),128))]),_:2},1024)]),createdAt:e(({record:a})=>[s(_(o.$formatTime(a.created_at)),1)]),updatedAt:e(({record:a})=>[s(_(o.$formatTime(a.updated_at)),1)]),operations:e(({record:a})=>[n(p,{class:"cursor-pointer"},{default:e(()=>[m((t(),r(u,{color:"arcoblue",onClick:d=>c("set",a.id)},{icon:e(()=>[n(f)]),default:e(()=>[s(" \u914D\u7F6E ")]),_:2},1032,["onClick"])),[[g,"configure:resource:value"]]),m((t(),r(u,{color:"orangered",onClick:d=>c("update",a)},{icon:e(()=>[n(f)]),default:e(()=>[s(" \u4FEE\u6539 ")]),_:2},1032,["onClick"])),[[g,"configure:resource:update"]]),o.$hasPermission("configure:resource:delete")?(t(),r(k,{key:0,content:"\u60A8\u786E\u8BA4\u5220\u9664\u6B64\u8D44\u6E90",type:"warning",onOk:d=>c("delete",a.id)},{default:e(()=>[n(u,{color:"red"},{icon:e(()=>[n(C)]),default:e(()=>[s(" \u5220\u9664 ")]),_:1})]),_:2},1032,["onOk"])):V("",!0)]),_:2},1024)]),_:1},8,["loading","columns","data","size"])),[[g,"configure:resource:query"]]),n(b,{total:l.total,current:i.value.current,"page-size":i.value.pageSize,"show-total":"","show-jumper":"","show-page-size":"",onChange:v,onPageSizeChange:h},null,8,["total","current","page-size"])]),_:1})}}});export{U as _};