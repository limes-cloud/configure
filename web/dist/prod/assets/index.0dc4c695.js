import"./index.69f3790f.js";/* empty css              *//* empty css               */import{d as x,e as a,br as B,B as C,aJ as n,aI as I,A as N,aR as c,aS as U}from"./arco.79b3a20c.js";import{p as b,a as k,u as R,d as T}from"./server.ac991692.js";import{u as E}from"./loading.4485fcab.js";import $ from"./tool.94122fc6.js";import{_ as L}from"./table.vue_vue_type_script_setup_true_lang.6ea29488.js";import{_ as M}from"./form.vue_vue_type_script_setup_true_lang.989d9886.js";import{_ as P}from"./search.vue_vue_type_script_setup_true_lang.42192447.js";import"./chart.c056e818.js";import"./vue.73de78ef.js";/* empty css               *//* empty css              *//* empty css               */import"./lodash.8c41d9ef.js";import"./sortable.esm.41bdaeb5.js";/* empty css               *//* empty css               *//* empty css               *//* empty css               *//* empty css              *//* empty css               *//* empty css                */const G={class:"container"},J={name:"ManagerServer"},ve=x({...J,setup(V){const r=a(),u=a({}),{setLoading:m}=E(!0),_=a(!1),p=a(),l=a("medium"),v=a(0),t=a({page:1,page_size:10}),d=a([{title:"\u670D\u52A1\u6807\u5FD7",dataIndex:"keyword",slotName:"keyword"},{title:"\u670D\u52A1\u540D\u79F0",dataIndex:"name",slotName:"name"},{title:"\u670D\u52A1\u63CF\u8FF0",dataIndex:"description",slotName:"description"},{title:"\u521B\u5EFA\u65F6\u95F4",dataIndex:"created_at",slotName:"createdAt"},{title:"\u66F4\u65B0\u65F6\u95F4",dataIndex:"updated_at",slotName:"updatedAt"},{title:"\u64CD\u4F5C",dataIndex:"operations",slotName:"operations"}]),o=async()=>{m(!0);try{const{data:e}=await b(t.value);p.value=e.list,v.value=e.total}finally{m(!1)}};o();const f=async e=>{await k(e),o(),c.success("\u521B\u5EFA\u6210\u529F")},g=async e=>{await R(e),o(),c.success("\u66F4\u65B0\u6210\u529F")},F=async e=>{await T(e),o(),c.success("\u5220\u9664\u6210\u529F")},h=async e=>{const s=t.value.page_size;t.value={...e,page:1,page_size:s},o()},A=()=>{u.value={},r.value.showAddDrawer()},z=e=>{u.value={...e},r.value.showUpdateDrawer()},w=e=>{u.value={},r.value.showAddDrawer()},D=async e=>{t.value.page=e.current,t.value.page_size=e.pageSize,o()};return(e,s)=>{const y=U("Breadcrumb"),S=B;return N(),C("div",G,[n(y),n(S,{class:"general-card"},{default:I(()=>[n(P,{onSearch:h}),n($,{size:l.value,"onUpdate:size":s[0]||(s[0]=i=>l.value=i),columns:d.value,"onUpdate:columns":s[1]||(s[1]=i=>d.value=i),onRefresh:o,onAdd:A},null,8,["size","columns"]),n(L,{columns:d.value,loading:_.value,data:p.value,size:l.value,total:v.value,pagination:{current:t.value.page,pageSize:t.value.page_size},onPageChange:D,onAdd:w,onUpdate:z,onDelete:F},null,8,["columns","loading","data","size","total","pagination"]),n(M,{ref_key:"formRef",ref:r,data:u.value,onAdd:f,onUpdate:g},null,8,["data"])]),_:1})])}}});export{ve as default};