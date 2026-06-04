import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

const parseUser = () => {
  try {
    const userData = localStorage.getItem('user')
    return userData ? JSON.parse(userData) : null
  } catch (e) {
    return null
  }
}

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem('token') || null)
  const user = ref(parseUser())

  const isAuthenticated = computed(() => !!token.value)

  const setAuth = (newToken, userData) => {
    token.value = newToken
    user.value = userData
    localStorage.setItem('token', newToken)
    localStorage.setItem('user', JSON.stringify(userData))
  }

  const logout = () => {
    token.value = null
    user.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('user')
  }

  return {
    token,
    user,
    isAuthenticated,
    setAuth,
    logout
  }
})
