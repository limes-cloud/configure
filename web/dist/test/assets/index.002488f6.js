import{d as k,E as y,u as D,a as E,b as P,a0 as q,au as z,av as U,aw as R,ax as L,ay as M,B as N,S as j,az as Q,e as T,f as s,t as g,j as d,g as e,w as o,M as $,o as G,h as x,aA as H,aB as J,_ as K}from"./index.b4c32530.js";/* empty css               *//* empty css               *//* empty css               */import{d as O,F as W}from"./index.2084b81f.js";import{u as X}from"./loading.1cfe06b8.js";const Y={class:"container"},Z={class:"logo"},ee=["src"],se={class:"content"},oe={class:"content-inner"},ae={class:"login-form-wrapper"},te={class:"login-form-title"},ne={class:"login-form-sub-title"},re={class:"login-form-error-msg"},ue={class:"login-form-password-actions"},le={class:"footer"},ce=k({__name:"index",setup(ie){const v=y(null),c=D(),f=y(""),{loading:h,setLoading:b}=X(),B=E(),w=P(),a=O("login-config",{rememberPassword:!0,username:"",password:""});q(()=>{clearInterval(v.value),v.value=null});const r=z({username:a.value.username,password:a.value.password,captcha:"",captcha_id:""}),F=async({errors:u,values:t})=>{if(!h.value&&!u){b(!0);try{await B.login(t);const{redirect:n,...i}=c.currentRoute.value.query;n&&c.hasRoute(n)?c.push({name:n,query:{...i}}):c.push({path:"/",query:{...i}}),$.success("\u767B\u9646\u6210\u529F");const{rememberPassword:l}=a.value,{username:_,password:m}=t;a.value.username=l?_:"",a.value.password=l?m:""}catch(n){f.value=n.message}finally{b(!1)}}},S=u=>{a.value.rememberPassword=u};return(u,t)=>{const n=U,i=R,l=H,_=L,m=J,A=M,C=N,V=j,I=Q;return G(),T("div",Y,[s("div",Z,[s("img",{alt:"logo",style:{width:"100px"},src:u.$logo},null,8,ee)]),s("div",se,[s("div",oe,[s("div",ae,[s("div",te,g(d(w).title),1),s("div",ne,g(d(w).desc),1),s("div",re,g(f.value),1),e(I,{model:r,class:"login-form",layout:"vertical",onSubmit:F},{default:o(()=>[e(l,{field:"username",rules:[{required:!0,message:"\u8D26\u6237\u4E0D\u80FD\u4E3A\u7A7A"}],"validate-trigger":["change","blur"],"hide-label":""},{default:o(()=>[e(i,{modelValue:r.username,"onUpdate:modelValue":t[0]||(t[0]=p=>r.username=p),size:"large",placeholder:"\u8BF7\u8F93\u5165\u8D26\u6237"},{prefix:o(()=>[e(n)]),_:1},8,["modelValue"])]),_:1}),e(l,{field:"password",rules:[{required:!0,message:"\u5BC6\u7801\u4E0D\u80FD\u4E3A\u7A7A"}],"validate-trigger":["change","blur"],"hide-label":""},{default:o(()=>[e(m,{modelValue:r.password,"onUpdate:modelValue":t[1]||(t[1]=p=>r.password=p),size:"large",placeholder:"\u8BF7\u8F93\u5165\u5BC6\u7801","allow-clear":"",autocomplete:""},{prefix:o(()=>[e(_)]),_:1},8,["modelValue"])]),_:1}),e(V,{size:16,direction:"vertical"},{default:o(()=>[s("div",ue,[e(A,{checked:"rememberPassword","model-value":d(a).rememberPassword,onChange:S},{default:o(()=>[x(" \u8BB0\u4F4F\u5BC6\u7801 ")]),_:1},8,["model-value","onChange"])]),e(C,{size:"large",type:"primary","html-type":"submit",long:"",loading:d(h)},{default:o(()=>[x(" \u786E\u8BA4\u767B\u9646 ")]),_:1},8,["loading"])]),_:1})]),_:1},8,["model"])])]),s("div",le,[e(W)])])])}}});const fe=K(ce,[["__scopeId","data-v-888bcb1a"]]);export{fe as default};