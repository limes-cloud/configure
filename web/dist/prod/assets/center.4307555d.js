import{o as R,v as oe,w as ne,_ as ae}from"./index.69f3790f.js";/* empty css               *//* empty css               */import{d as ue,e as l,r as le,w as se,aX as re,T as ie,aP as de,bo as ce,bm as ve,B as N,E as _,be as h,aC as b,aI as a,aJ as u,aR as U,aS as me,aG as pe,A as m,aL as C,aN as fe,aK as _e,aM as ye,aQ as ge,bp as Fe,bv as be}from"./arco.79b3a20c.js";import{b as Ce,p as we}from"./template.cd4580b9.js";import Ee from"./compare.06c878d1.js";function Me(r){return R.put(`/configure/v1/configure/${r.env_keyword}`,r)}function ke(r){return R.post(`/configure/v1/configure/compare/${r.env_keyword}`,r)}const Be={class:"configure"},he={class:"header"},De={class:"header-item"},Ve={class:"header-item"},Se={class:"header-item"},Te={class:"edit"},Oe=ue({__name:"center",props:{template:null,envs:null},emits:["submit","sync"],setup(r,{emit:D}){const i=r,V="submit",w="sync",S="preview",T=l(),c=l(!1),d=l(""),p=l(!1),O=l(),s=l({}),y=l(!1),$=l(),v=l({}),n=le({description:"",content:"",format:"json"}),f=l(!1),g=l(""),E=l([]),M=async()=>(n.description=v.value.description,n.format==="json"?n.content=JSON.stringify(JSON.parse(n.content)):g.value=n.content,D("submit",n),!0),P=()=>{var o;return D("sync",{server_id:(o=i.template)==null?void 0:o.server_id,env_keyword:s.value.env_keyword,description:v.value.description}),!0},j=async()=>{var e;const{data:o}=await Ce({id:(e=i.template)==null?void 0:e.id,content:n.content,format:n.format});if(!o||!o.length){U.error("\u6A21\u677F\u4E0D\u5B58\u5728\u53D8\u66F4");return}d.value=V,c.value=!0,E.value=o},J=async()=>{var e;const{data:o}=await ke({server_id:(e=i.template)==null?void 0:e.server_id,env_keyword:s.value.env_keyword});if(!o||!o.length){U.error("\u914D\u7F6E\u4E0D\u5B58\u5728\u53D8\u66F4");return}c.value=!0,E.value=o},L=()=>(c.value=!1,v.value.description="",y.value=!0,!0),G=async()=>{d.value===V&&M(),d.value===w&&P()},z=()=>{p.value=!0,s.value.env_keyword=void 0,s.value.title="\u9884\u89C8",d.value=S},K=()=>{p.value=!0,s.value.env_keyword=void 0,s.value.title="\u540C\u6B65",d.value=w},Q=async()=>{var e;const{data:o}=await we({env_keyword:s.value.env_keyword,content:n.content,format:n.format,server_id:(e=i.template)==null?void 0:e.server_id});n.format==="json"?g.value=JSON.stringify(JSON.parse(o.content),null,3):g.value=o.content,f.value=!0},W=async()=>await O.value.validate()?!1:(d.value===S&&Q(),d.value===w&&J(),!0);return se(()=>i.template,o=>{!o||(n.content=o.content,n.format=o.format,T.value.setEditLang(o.format))},{deep:!0}),(o,e)=>{const X=oe,k=re,Y=ie,q=ne,x=me("CodeEditor"),H=ge,Z=de,A=Fe,I=ce,F=pe,ee=ve,B=be("permission");return m(),N("div",Be,[_("div",he,[_("div",De,[h((m(),b(k,{type:"primary",status:"success",disabled:!i.template,onClick:z},{icon:a(()=>[u(X)]),default:a(()=>[C(" \u751F\u6210\u9884\u89C8 ")]),_:1},8,["disabled"])),[[B,"configure:configure:preview"]])]),_("div",Ve,[h((m(),b(k,{type:"primary",disabled:!i.template,onClick:j},{icon:a(()=>[u(Y)]),default:a(()=>[C(" \u63D0\u4EA4\u6A21\u677F ")]),_:1},8,["disabled"])),[[B,"configure:template:add"]])]),_("div",Se,[h((m(),b(k,{type:"primary",disabled:!i.template,onClick:K},{icon:a(()=>[u(q)]),default:a(()=>[C(" \u540C\u6B65\u914D\u7F6E ")]),_:1},8,["disabled"])),[[B,"configure:configure:sync"]])])]),_("div",Te,[u(x,{ref_key:"coder",ref:T,modelValue:n.content,"onUpdate:modelValue":e[0]||(e[0]=t=>n.content=t),lang:n.format,"show-line":!1,"switch-lang":!0,style:{width:"100%",height:"100%"},onChangeLang:e[1]||(e[1]=t=>{n.format=t})},null,8,["modelValue","lang"])]),u(F,{visible:p.value,"onUpdate:visible":e[3]||(e[3]=t=>p.value=t),simple:"",title:"\u8BF7\u9009\u62E9\u73AF\u5883",onCancel:e[4]||(e[4]=t=>p.value=!1),onBeforeOk:W},{default:a(()=>[u(I,{ref_key:"envFormRef",ref:O,model:s.value,layout:"vertical"},{default:a(()=>[u(A,{field:"env_keyword",label:"\u8BF7\u9009\u62E9\u9700\u8981"+s.value.title+"\u7684\u73AF\u5883",rules:[{required:!0,message:"\u8BF7\u9009\u62E9\u9700\u8981\u6E32\u67D3\u7684\u73AF\u5883"}],"validate-trigger":["change","input"]},{default:a(()=>[u(Z,{modelValue:s.value.env_keyword,"onUpdate:modelValue":e[2]||(e[2]=t=>s.value.env_keyword=t),"allow-search":"",placeholder:"\u8BF7\u9009\u62E9"},{default:a(()=>[(m(!0),N(fe,null,_e(r.envs,(t,te)=>(m(),b(H,{key:te,value:t.keyword},{default:a(()=>[C(ye(t.name),1)]),_:2},1032,["value"]))),128))]),_:1},8,["modelValue"])]),_:1},8,["label"])]),_:1},8,["model"])]),_:1},8,["visible"]),u(F,{visible:f.value,"onUpdate:visible":e[5]||(e[5]=t=>f.value=t),title:"\u914D\u7F6E\u8BE6\u60C5",width:"900px","unmount-on-close":"","body-style":{height:"500px",padding:"0"},onCancel:e[6]||(e[6]=t=>f.value=!1),onBeforeOk:e[7]||(e[7]=t=>f.value=!1)},{default:a(()=>[u(x,{value:g.value,lang:n.format,"show-line":!1,"read-only":!0,style:{width:"100%",height:"100%"}},null,8,["value","lang"])]),_:1},8,["visible"]),u(F,{visible:c.value,"onUpdate:visible":e[8]||(e[8]=t=>c.value=t),title:"\u53D8\u66F4\u8BE6\u60C5",width:"900px","ok-text":"\u786E\u8BA4\u53D8\u66F4","unmount-on-close":"","body-style":{height:"500px",padding:"0"},onCancel:e[9]||(e[9]=t=>c.value=!1),onBeforeOk:L},{default:a(()=>[u(Ee,{data:E.value},null,8,["data"])]),_:1},8,["visible"]),u(F,{visible:y.value,"onUpdate:visible":e[11]||(e[11]=t=>y.value=t),simple:"",title:"\u53D8\u66F4\u63CF\u8FF0",onCancel:e[12]||(e[12]=t=>y.value=!1),onBeforeOk:G},{default:a(()=>[u(I,{ref_key:"descFormRef",ref:$,model:v.value,layout:"vertical"},{default:a(()=>[u(A,{field:"description",label:"\u8BF7\u7B80\u8981\u6982\u8FF0\u4E00\u4E0B\u672C\u6B21\u63D0\u4EA4\u7684\u4FEE\u6539",rules:[{required:!0,message:"\u6A21\u677F\u63CF\u8FF0\u662F\u5FC5\u586B\u9879"}],"validate-trigger":["change","input"]},{default:a(()=>[u(ee,{modelValue:v.value.description,"onUpdate:modelValue":e[10]||(e[10]=t=>v.value.description=t),placeholder:"\u8BF7\u8F93\u5165"},null,8,["modelValue"])]),_:1})]),_:1},8,["model"])]),_:1},8,["visible"])])}}});const xe=ae(Oe,[["__scopeId","data-v-08034d41"]]),Pe=Object.freeze(Object.defineProperty({__proto__:null,default:xe},Symbol.toStringTag,{value:"Module"}));export{xe as C,Pe as c,Me as s};