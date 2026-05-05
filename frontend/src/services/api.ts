import axios from 'axios';

const api = axios.create({
  baseURL: '/api/v1',
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json',
  },
});

api.interceptors.request.use((config) => {
  const token = localStorage.getItem('langops_token');
  if (token) {
    config.headers.Authorization = Bearer ;
  }
  return config;
});

api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('langops_token');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

export default api;

// Version management APIs
export const versionApi = {
  list: () => api.get('/versions'),
  create: (data: any) => api.post('/versions', data),
  get: (id: string) => api.get(/versions/),
  rollback: (id: string) => api.post(/versions//rollback),
  branch: (id: string, data: any) => api.post(/versions//branch, data),
};

// Evaluation APIs
export const evalApi = {
  list: () => api.get('/evaluations'),
  create: (data: any) => api.post('/evaluations', data),
  run: (id: string) => api.post(/evaluations//run),
  getReport: (id: string) => api.get(/evaluations//report),
};

// Release APIs
export const releaseApi = {
  list: () => api.get('/releases'),
  create: (data: any) => api.post('/releases', data),
  approve: (id: string) => api.post(/releases//approve),
  startCanary: (id: string, data: any) => api.post(/releases//canary, data),
  rollback: (id: string) => api.post(/releases//rollback),
};

// A/B Test APIs
export const abtestApi = {
  list: () => api.get('/abtests'),
  create: (data: any) => api.post('/abtests', data),
  start: (id: string) => api.post(/abtests//start),
  stop: (id: string) => api.post(/abtests//stop),
  getResult: (id: string) => api.get(/abtests//result),
};