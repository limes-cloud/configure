import { Menu } from './types/menu';

import menuData from './menu.json';

export function getMenuTree() {
  return menuData as Menu[];
}

export default null;
