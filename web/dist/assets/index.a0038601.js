import{n as l}from"./index.34172a3f.js";/* empty css              *//* empty css               */import{d as T,e as a,bp as V,B as G,aJ as s,aI as $,A as j,aP as c,aQ as J}from"./arco.9b5dfa89.js";import{u as M}from"./loading.eb6d651b.js";import{a as P}from"./resource-server.516ebeb2.js";import{a as Q}from"./env.88bc2fe1.js";import q from"./tool.51bf6a5d.js";import{_ as H}from"./table.vue_vue_type_script_setup_true_lang.e106b222.js";import{_ as K}from"./form.vue_vue_type_script_setup_true_lang.44acddc1.js";import{_ as O}from"./search.vue_vue_type_script_setup_true_lang.8f34cda1.js";import{_ as W}from"./value.vue_vue_type_script_setup_true_lang.5a82f933.js";import"./chart.a7f56037.js";import"./vue.b2201ea8.js";/* empty css               *//* empty css              *//* empty css               *//* empty css               *//* empty css               */import"./lodash.d7f1e1af.js";import"./sortable.esm.41bdaeb5.js";/* empty css               *//* empty css              *//* empty css               *//* empty css                *//* empty css               *//* empty css              *//* empty css               *//* empty css                */import"./server.abab0313.js";import"./index.8f3462d0.js";function X(t){return l.get("/configure/v1/resources",{params:{...t}})}function Y(t){return l.post("/configure/v1/resource",t)}function Z(t){return l.put("/configure/v1/resource",t)}function ee(t){return l.delete("/configure/v1/resource",{params:{id:t}})}function ae(t){return l.get("/configure/v1/resource/values",{params:{resource_id:t}})}function te(t,u){return l.put("/configure/v1/resource/value",{resource_id:t,list:u})}const oe={class:"container"},se={name:"ConfigureResource"},Le=T({...se,setup(t){const u=a(),v=a(),n=a({}),{setLoading:f}=M(!0),y=a(!1),_=a(),g=a([]),d=a("medium"),F=a(0),h=a([]),w=a([]),R=a(0),i=a({page:1,page_size:10}),p=a([{title:"\u53D8\u91CF\u6807\u5FD7",dataIndex:"keyword",slotName:"keyword"},{title:"\u53D8\u91CF\u5B57\u6BB5",dataIndex:"fields",slotName:"fields"},{title:"\u53D8\u91CF\u6807\u7B7E",dataIndex:"tag",slotName:"tag"},{title:"\u53D8\u91CF\u63CF\u8FF0",dataIndex:"description",slotName:"description"},{title:"\u521B\u5EFA\u65F6\u95F4",dataIndex:"created_at",slotName:"createdAt"},{title:"\u66F4\u65B0\u65F6\u95F4",dataIndex:"updated_at",slotName:"updatedAt"},{title:"\u64CD\u4F5C",dataIndex:"operations",slotName:"operations"}]),D=async()=>{const{data:e}=await Q();h.value=e},A=async e=>{const{data:o}=await ae(e);w.value=o},B=async e=>{const{data:o}=await P(e);g.value=o},r=async()=>{f(!0);try{const{data:e}=await X(i.value);_.value=e.list,F.value=e.total}finally{f(!1)}};r(),D();const z=async e=>{e.fields=e.fieldList.join(","),await Y(e),r(),c.success("\u521B\u5EFA\u6210\u529F")},C=async e=>{e.fields=e.fieldList.join(","),await Z(e),r(),c.success("\u66F4\u65B0\u6210\u529F")},x=async e=>{await ee(e),r(),c.success("\u5220\u9664\u6210\u529F")},U=async e=>{const o=i.value.page_size;i.value={...e,page:1,page_size:o},r()},I=async e=>{await te(R.value,e),c.success("\u914D\u7F6E\u6210\u529F")},N=()=>{n.value={},u.value.showAddDrawer()},S=e=>{B(e.id),n.value={...e},n.value.fieldList=n.value.fields.split(","),u.value.showUpdateDrawer()},b=e=>{n.value={},u.value.showAddDrawer()},E=e=>{R.value=e,A(e),v.value.showUpdateDrawer()};return(e,o)=>{const k=J("Breadcrumb"),L=V;return j(),G("div",oe,[s(k),s(L,{class:"general-card"},{default:$(()=>[s(O,{onSearch:U}),s(q,{size:d.value,"onUpdate:size":o[0]||(o[0]=m=>d.value=m),columns:p.value,"onUpdate:columns":o[1]||(o[1]=m=>p.value=m),onRefresh:r,onAdd:N},null,8,["size","columns"]),s(H,{columns:p.value,loading:y.value,data:_.value,size:d.value,total:F.value,pagination:{current:i.value.page,pageSize:i.value.page_size},onAdd:b,onUpdate:S,onDelete:x,onSet:E},null,8,["columns","loading","data","size","total","pagination"]),s(K,{ref_key:"formRef",ref:u,data:n.value,servers:g.value,onAdd:z,onUpdate:C},null,8,["data","servers"]),s(W,{ref_key:"valueRef",ref:v,values:w.value,envs:h.value,onUpdate:I},null,8,["values","envs"])]),_:1})])}}});export{Le as default};