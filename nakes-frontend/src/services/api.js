import axios from 'axios'
import { useAuthStore } from '../stores/auth'

const apiClient = axios.create({
  baseURL: 'https://pil-time-pam-production.up.railway.app/api',
  headers: {
    'Content-Type': 'application/json'
  }
})

// jadwalApiClient sekarang juga mengarah ke main backend (port 8080)
// karena jadwal-service sudah di-merge ke dalam backend
const jadwalApiClient = axios.create({
  baseURL: 'https://pil-time-pam-production.up.railway.app/api/admin',
  headers: {
    'Content-Type': 'application/json'
  }
})

// Interceptor untuk apiClient (Main Backend)
apiClient.interceptors.request.use(
  config => {
    const authStore = useAuthStore()
    if (authStore.token) {
      config.headers.Authorization = `Bearer ${authStore.token}`
    }
    return config
  },
  error => Promise.reject(error)
)

apiClient.interceptors.response.use(
  response => response,
  error => {
    if (error.response?.status === 401) {
      const currentPath = window.location.pathname
      if (currentPath !== '/login') {
        const authStore = useAuthStore()
        authStore.logout()
        window.location.href = '/login'
      }
    }
    return Promise.reject(error)
  }
)

// Interceptor untuk jadwalApiClient (Jadwal Service)
jadwalApiClient.interceptors.request.use(
  config => {
    const authStore = useAuthStore()
    if (authStore.token) {
      config.headers.Authorization = `Bearer ${authStore.token}`
    }
    return config
  },
  error => Promise.reject(error)
)

jadwalApiClient.interceptors.response.use(
  response => response,
  error => {
    if (error.response?.status === 401) {
      const currentPath = window.location.pathname
      if (currentPath !== '/login') {
        const authStore = useAuthStore()
        authStore.logout()
        window.location.href = '/login'
      }
    }
    return Promise.reject(error)
  }
)

export { jadwalApiClient }
export default apiClient
