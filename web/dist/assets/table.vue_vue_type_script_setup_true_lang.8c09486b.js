import"./index.f90c2f0c.js";/* empty css               *//* empty css              *//* empty css              *//* empty css               *//* empty css               *//* empty css                *//* empty css              *//* empty css               *//* empty css               */import{d as y,e as w,ae as T,bu as D,ad as S,bv as E,aU as $,bw as A,bx as B,aC as c,aI as e,A as u,bc as d,aL as n,aM as g,u as P,aJ as i,aD as I,bt as N}from"./arco.9b5dfa89.js";import{v as V}from"./type.a4a3f801.js";const R=y({__name:"table",props:{columns:null,loading:{type:Boolean},data:null,size:null,pagination:null,total:null},emits:["delete","update","add","set","resetToken","queryToken","pageChange"],setup(s,{emit:r}){const t=w({current:1,pageSize:10}),v=o=>{t.value.current=o,r("pageChange",t.value)},b=o=>{t.value.pageSize=o,r("pageChange",t.value)};return(o,q)=>{const m=T,l=D,C=S,h=E,_=$,k=A,z=B,p=N("permission");return u(),c(_,{direction:"vertical",fill:""},{default:e(()=>[d((u(),c(k,{"row-key":"id",loading:s.loading,columns:s.columns,pagination:!1,data:s.data,bordered:!1,size:s.size},{type:e(({record:a})=>[n(g(P(V)[a.type]),1)]),createdAt:e(({record:a})=>[n(g(o.$formatTime(a.created_at)),1)]),updatedAt:e(({record:a})=>[n(g(o.$formatTime(a.updated_at)),1)]),operations:e(({record:a})=>[i(_,{class:"cursor-pointer"},{default:e(()=>[d((u(),c(l,{color:"arcoblue",onClick:f=>r("set",a)},{icon:e(()=>[i(m)]),default:e(()=>[n(" \u914D\u7F6E ")]),_:2},1032,["onClick"])),[[p,"configure:business:value"]]),d((u(),c(l,{color:"orangered",onClick:f=>r("update",a)},{icon:e(()=>[i(m)]),default:e(()=>[n(" \u4FEE\u6539 ")]),_:2},1032,["onClick"])),[[p,"configure:business:update"]]),o.$hasPermission("configure:business:delete")?(u(),c(h,{key:0,content:"\u60A8\u786E\u8BA4\u5220\u9664\u6B64\u8D44\u6E90",type:"warning",onOk:f=>r("delete",a.id)},{default:e(()=>[i(l,{color:"red"},{icon:e(()=>[i(C)]),default:e(()=>[n(" \u5220\u9664 ")]),_:1})]),_:2},1032,["onOk"])):I("",!0)]),_:2},1024)]),_:1},8,["loading","columns","data","size"])),[[p,"configure:business:query"]]),i(z,{total:s.total,current:t.value.current,"page-size":t.value.pageSize,"show-total":"","show-jumper":"","show-page-size":"",onChange:v,onPageSizeChange:b},null,8,["total","current","page-size"])]),_:1})}}});export{R as _};