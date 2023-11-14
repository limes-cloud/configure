import { defineStore } from 'pinia';
import { setToken, clearToken } from '@/utils/auth';
import { removeRouteListener } from '@/utils/route-listener';
import rsa from '@/utils/rsa';
import Message from '@arco-design/web-vue/es/message';
import { currentUser } from '@/api/configure/user';
import { login as userLogin, logout as userLogout } from '@/api/configure/auth';
import { LoginReq } from '@/api/configure/types/auth';
import { User } from '@/api/configure/types/user';
import useAppStore from '../app';

const useUserStore = defineStore('user', {
  state: (): User => ({
    id: 0,
    department_id: 0,
    role_id: 0,
    role_ids: [],
    name: '',
    nickname: '',
    gender: '',
    phone: '',
    avatar: '',
    email: '',
    status: false,
    disabled: '',
    last_login: 0,
    created_at: 0,
    updated_at: 0,
    role: undefined,
  }),

  getters: {
    userInfo(state: User): User {
      return { ...state };
    },
  },

  actions: {
    // Set user's information
    setInfo(partial: Partial<User>) {
      this.$patch(partial);
    },

    // Reset user's information
    resetInfo() {
      this.$reset();
    },

    // Get user's information
    async info() {
      const data = await currentUser();
      this.setInfo(data);
    },

    // Login
    async login(req: LoginReq) {
      const info = {
        ...req,
        password: rsa.encrypt({
          password: req.password,
          time: new Date().getTime(),
        }),
      };
      try {
        const { data } = await userLogin(info as LoginReq);
        setToken(data.token);
      } catch (err) {
        clearToken();
        throw err;
      }
    },
    clear() {
      const appStore = useAppStore();
      this.resetInfo();
      clearToken();
      removeRouteListener();
      appStore.clearApp();
    },
    // Logout
    async logout() {
      try {
        await userLogout();
      } finally {
        this.clear();
      }
    },
  },
});

export default useUserStore;
