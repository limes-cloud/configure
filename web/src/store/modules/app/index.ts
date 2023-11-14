import { defineStore } from 'pinia';
import type { RouteRecordNormalized } from 'vue-router';
import defaultSettings from '@/config/settings.json';
import { App, Home } from '@/router/types';
import { AppState } from './types';

const useAppStore = defineStore('app', {
  state: (): AppState => {
    const setting: AppState = {
      ...defaultSettings,
      menus: new Map(),
      permissions: new Map(),
      apps: [],
      app: '',
      homes: new Map(),
    };
    return setting;
  },

  getters: {
    appCurrentSetting(state: AppState): AppState {
      return { ...state };
    },
    appDevice(state: AppState) {
      return state.device;
    },
    appMenu(state: AppState) {
      return state.menus.get(state.app) || [];
    },
    appList(state: AppState) {
      return state.apps || [];
    },
    appPermissions(state: AppState) {
      return state.permissions;
    },
    appHome(state: AppState) {
      return state.homes.get(state.app);
    },
    appHomePath(state: AppState) {
      const home = state.homes.get(state.app);
      if (home) {
        return home.path;
      }

      const menus = state.menus.get(state.app);
      if (menus) {
        return menus[0].path;
      }

      return '';
    },
    currentAppKey(state: AppState) {
      return state.app;
    },
    currentAppinfo(state: AppState) {
      let app = <App>{};
      state.apps.forEach((item) => {
        if (item.keyword === state.app) {
          app = item;
        }
      });
      return app;
    },
  },

  actions: {
    // Update app settings
    updateSettings(partial: Partial<AppState>) {
      // @ts-ignore-next-line
      this.$patch(partial);
    },

    // Change theme color
    toggleTheme(dark: boolean) {
      if (dark) {
        this.theme = 'dark';
        document.body.setAttribute('arco-theme', 'dark');
      } else {
        this.theme = 'light';
        document.body.removeAttribute('arco-theme');
      }
    },
    toggleDevice(device: string) {
      this.device = device;
    },
    toggleMenu(value: boolean) {
      this.hideMenu = value;
    },
    clearApp() {
      this.menus = new Map();
      this.permissions = new Map();
      this.apps = [];
      this.app = '';
      this.homes = new Map();
    },
    setApps(apps: App[]) {
      this.apps = apps;
    },
    setMenus(menus: Map<string, RouteRecordNormalized[]>) {
      this.menus = menus;
    },
    setPermissions(ps: Map<string, string>) {
      this.permissions = ps;
    },
    setHomes(homes: Map<string, Home>) {
      this.homes = homes;
    },
    setAppHome(app: string, home: Home) {
      this.homes.set(app, home);
    },
    setCurrentApp(app?: string) {
      this.app = app || this.apps[0].keyword;
    },
  },
});

export default useAppStore;
