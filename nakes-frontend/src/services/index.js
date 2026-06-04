import apiClient from './api'
import axios from 'axios'

// Auth Service client (port 8080) — login, register, reset password
const authApiClient = axios.create({
  baseURL: 'https://pil-time-pam-production.up.railway.app',
  headers: {
    'Content-Type': 'application/json'
  }
})

export const authService = {
  // Admin/Nakes login → auth-service
  loginAdmin(email, password) {
    return authApiClient.post('/auth/nakes/login', {
      email,
      password
    })
  }
}

export const pasienService = {
  getAllPasien() {
    return apiClient.get('/admin/pasien')
  },

  deletePasien(id) {
    return apiClient.delete(`/admin/pasien/${id}`)
  }
}

export const jadwalService = {
  getAllJadwal() {
    return apiClient.get('/admin/jadwal')
  },

  createJadwal(data) {
    return apiClient.post('/admin/jadwal', data)
  },

  updateJadwal(id, data) {
    return apiClient.put(`/admin/jadwal/${id}`, data)
  },

  deleteJadwal(id) {
    return apiClient.delete(`/admin/jadwal/${id}`)
  }
}

export const obatService = {
  getAllObat() {
    return apiClient.get('/admin/info-obat')
  },

  createObat(data) {
    return apiClient.post('/admin/info-obat', data)
  },

  updateObat(id, data) {
    return apiClient.put(`/admin/info-obat/${id}`, data)
  },

  deleteObat(id) {
    return apiClient.delete(`/admin/info-obat/${id}`)
  }
}

export const dashboardService = {
  getDashboard() {
    return apiClient.get('/admin/dashboard')
  }
}

export const riwayatService = {
  getAllRiwayat() {
    return apiClient.get('/admin/riwayat')
  }
}
