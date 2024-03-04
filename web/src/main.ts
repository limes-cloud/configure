import { createApp } from 'vue';
import ArcoVue from '@arco-design/web-vue';
import ArcoVueIcon from '@arco-design/web-vue/es/icon';
import globalComponents from '@/components';
import logo from '@/assets/logo.png';
import router from './router';
import store from './store';
import directive from './directive';
// import './mock';
import App from './App.vue';

import '@/assets/style/global.less';
import '@/utils/interceptor';

import { formatTime, parseTime } from './utils/time';
import { densityList, genderList } from './utils/consts';
import { hasPermission } from './utils/permission';

const app = createApp(App);

app.config.globalProperties.$formatUrl = (url: string) => {
  return `/${url}`;
};
app.config.globalProperties.$logo = logo;
app.config.globalProperties.$formatTime = formatTime;
app.config.globalProperties.$parseTime = parseTime;
app.config.globalProperties.$densityList = densityList;
app.config.globalProperties.$genderList = genderList;
app.config.globalProperties.$hasPermission = hasPermission;

app.use(ArcoVue, {});
app.use(ArcoVueIcon);

app.use(router);
app.use(store);
app.use(globalComponents);
app.use(directive);


app.mount('#app');
