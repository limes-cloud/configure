import{d as V,E as d,$ as B,aw as C,b0 as k,aV as E,az as h,af as x,k as U,w as t,o as y,g as a,h as F,aA as q,b1 as R}from"./index.b4c32530.js";/* empty css              *//* empty css               *//* empty css               *//* empty css                *//* empty css               */const H=V({__name:"form",props:{data:null},emits:["add","update"],setup(_,{expose:c,emit:i}){const m=_,p=d(),o=d(!1),r=d(!1),e=d({...m.data});B(()=>m.data,n=>{e.value=n}),c({showAddDrawer:()=>{o.value=!0,r.value=!0},showUpdateDrawer:()=>{o.value=!0,r.value=!1},closeDrawer:()=>{o.value=!1}});const w=async()=>await p.value.validate()?!1:(r.value?i("add",{...e.value}):i("update",{...e.value}),!0);return(n,u)=>{const f=C,s=q,v=k,D=R,A=E,g=h,b=x;return y(),U(b,{visible:o.value,"onUpdate:visible":u[4]||(u[4]=l=>o.value=l),title:r.value?"\u65B0\u5EFA":"\u4FEE\u6539",width:"380px",onCancel:u[5]||(u[5]=l=>o.value=!1),onBeforeOk:w},{default:t(()=>[a(g,{ref_key:"formRef",ref:p,model:e.value,"label-align":"left",layout:"horizontal","auto-label-width":""},{default:t(()=>[a(s,{field:"name",label:"\u670D\u52A1\u540D\u79F0",rules:[{required:!0,message:"\u670D\u52A1\u540D\u79F0\u662F\u5FC5\u586B\u9879"}],"validate-trigger":["change","input"]},{default:t(()=>[a(f,{modelValue:e.value.name,"onUpdate:modelValue":u[0]||(u[0]=l=>e.value.name=l),placeholder:"\u8BF7\u8F93\u5165\u670D\u52A1\u540D\u79F0","allow-clear":""},null,8,["modelValue"])]),_:1}),a(s,{field:"keyword",label:"\u670D\u52A1\u6807\u5FD7",rules:[{required:!0,message:"\u670D\u52A1\u6807\u5FD7\u662F\u5FC5\u586B\u9879"}],"validate-trigger":["change","input"]},{default:t(()=>[a(f,{modelValue:e.value.keyword,"onUpdate:modelValue":u[1]||(u[1]=l=>e.value.keyword=l),placeholder:"\u8BF7\u8F93\u5165\u670D\u52A1\u6807\u5FD7","allow-clear":""},null,8,["modelValue"])]),_:1}),a(s,{field:"is_business",label:"\u4E1A\u52A1\u670D\u52A1",rules:[{required:!0,message:"\u4E1A\u52A1\u670D\u52A1\u662F\u5FC5\u586B\u9879"}],"validate-trigger":["change","input"]},{default:t(()=>[a(D,{modelValue:e.value.is_business,"onUpdate:modelValue":u[2]||(u[2]=l=>e.value.is_business=l)},{default:t(()=>[a(v,{value:!0},{default:t(()=>[F("\u662F")]),_:1}),a(v,{value:!1},{default:t(()=>[F("\u5426")]),_:1})]),_:1},8,["modelValue"])]),_:1}),a(s,{field:"description",label:"\u670D\u52A1\u63CF\u8FF0",rules:[{required:!0,message:"\u670D\u52A1\u63CF\u8FF0\u662F\u5FC5\u586B\u9879"}],"validate-trigger":["change","input"]},{default:t(()=>[a(A,{modelValue:e.value.description,"onUpdate:modelValue":u[3]||(u[3]=l=>e.value.description=l),placeholder:"\u8BF7\u8F93\u5165\u670D\u52A1\u63CF\u8FF0","allow-clear":""},null,8,["modelValue"])]),_:1})]),_:1},8,["model"])]),_:1},8,["visible","title"])}}});export{H as _};