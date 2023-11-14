import axios from 'axios';
import { LoginReq, LoginRes } from './types/auth';

export function captcha() {
  return axios.post('/configure/v1/login/captcha');
}

export function login(data: LoginReq) {
  return axios.post<LoginRes>('/configure/v1/login', data);
}

export function logout() {
  return {};
}

export function refreshToken() {
  return axios.post<LoginRes>('/configure/v1/token/refresh');
}
