import{d as k,E as n,$ as E,az as F,af as b,k as v,w as u,M as C,o as i,g as m,e as D,a3 as B,a_ as S,a2 as x,aA as N}from"./index.527ddea7.js";/* empty css              *//* empty css               *//* empty css               */const $=k({__name:"value",props:{values:null,envs:null},emits:["update"],setup(d,{expose:p,emit:w}){const f=d,c=n(),l=n(!1),r=n({});E(()=>f.values,t=>{!t||(r.value={},t.forEach(e=>{r.value[e.environment_id]=e.values}))}),p({showUpdateDrawer:()=>{l.value=!0},closeDrawer:()=>{l.value=!1}});const _=async()=>{if(await c.value.validate())return!1;const e=[];try{f.envs.forEach(o=>{const s=JSON.stringify(JSON.parse(r.value[o.id]));e.push({env_keyword:o.keyword,environment_id:o.id,values:JSON.stringify(JSON.parse(s))})})}catch{return C.error("\u6570\u636E\u683C\u5F0F\u5B58\u5728\u9519\u8BEF"),!1}return w("update",[...e]),!0};return(t,e)=>{const o=N,s=F,h=b;return i(),v(h,{visible:l.value,"onUpdate:visible":e[0]||(e[0]=a=>l.value=a),title:"\u53D8\u91CF\u503C\u914D\u7F6E",width:"380px","unmount-on-close":"",onCancel:e[1]||(e[1]=a=>l.value=!1),onBeforeOk:_},{default:u(()=>[m(s,{ref_key:"formRef",ref:c,model:r.value,"label-align":"left",layout:"vertical","auto-label-width":""},{default:u(()=>[(i(!0),D(x,null,B(d.envs,(a,g)=>(i(),v(o,{key:g,field:String(a.id),label:a.name+"("+a.keyword+")",rules:[{required:!0,message:a.name+"("+a.keyword+")\u914D\u7F6E\u662F\u5FC5\u586B\u9879"}],"validate-trigger":["change","input"]},{default:u(()=>[m(S,{modelValue:r.value[a.id],"onUpdate:modelValue":y=>r.value[a.id]=y,style:{width:"100%",height:"150px"},"show-folding":!1,"show-line":!1},null,8,["modelValue","onUpdate:modelValue"])]),_:2},1032,["field","label","rules"]))),128))]),_:1},8,["model"])]),_:1},8,["visible"])}}});export{$ as _};