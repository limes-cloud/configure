export interface User {
  id: number;
  department_id: number;
  role_id: number;
  role_ids: number[];
  name: string;
  nickname: string;
  gender: string;
  phone: string;
  avatar: string;
  email: string;
  status: boolean;
  disabled: string;
  last_login: number;
  token?: string;
  password?: string;
  created_at: number;
  updated_at: number;
}
