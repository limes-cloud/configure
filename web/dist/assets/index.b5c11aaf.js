import{n as f,a as g,_ as h}from"./index.8d26ebff.js";import{d as y,e as t,B as x,aJ as a,E as S,D as C,aP as u,A as B,aQ as b}from"./arco.9b5dfa89.js";import{C as E,a as w,c as A}from"./center.648f0f82.js";import{a as D}from"./env.02bee142.js";import F from"./left.e5f2f1c1.js";import T from"./right.d9d76ee6.js";import"./chart.a7f56037.js";import"./vue.b2201ea8.js";/* empty css               *//* empty css              *//* empty css              *//* empty css               *//* empty css               *//* empty css               */import"./index.02893194.js";import"./server.e5f610be.js";/* empty css               *//* empty css                *//* empty css               */import"./business.73c55f9b.js";import"./resource-server.903e2fab.js";function k(o){return f.put(`/configure/v1/configure/${o.env_keyword}`,o)}const G={class:"container"},I={name:"ConfigureResource"},N=y({...I,setup(o){const c=g(),s=t(56+58+20),n=t(0),r=t(),i=t();c.footer&&(s.value+=40),c.tabBar&&(s.value+=32),(async()=>{const{data:e}=await D();i.value=e})();const m=e=>{n.value=e,l(e)},d=async e=>{e.server_id=n.value,await w(e),l(e.server_id),u.success("\u6A21\u677F\u63D0\u4EA4\u6210\u529F")},l=async e=>{const{data:p}=await A(e);r.value=p},v=async e=>{await k(e),u.success("\u914D\u7F6E\u540C\u6B65\u6210\u529F")};return(e,p)=>{const _=b("Breadcrumb");return B(),x("div",G,[a(_),S("div",{class:"general-card",style:C({display:"flex",flexDirection:"row",height:"calc(100vh - "+s.value+"px)"})},[a(F,{class:"left",template:r.value,onSelect:m},null,8,["template"]),a(E,{template:r.value,envs:i.value,class:"center",onSubmit:d,onSync:v},null,8,["template","envs"]),a(T,{"server-id":n.value,class:"right"},null,8,["server-id"])],4)])}}});const oe=h(N,[["__scopeId","data-v-d1b05f57"]]);export{oe as default};