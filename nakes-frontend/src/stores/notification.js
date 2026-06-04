import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useNotificationStore = defineStore('notification', () => {
  const notifications = ref([])

  const addNotification = ({ type = 'info', title = '', message = '', duration = 4000 }) => {
    const id = Date.now() + Math.random().toString(36).substring(2, 9)
    const newNotif = { id, type, title, message, duration }
    notifications.value.push(newNotif)

    if (duration > 0) {
      setTimeout(() => {
        removeNotification(id)
      }, duration)
    }

    return id
  }

  const removeNotification = (id) => {
    notifications.value = notifications.value.filter(n => n.id !== id)
  }

  // Helper shortcuts
  const success = (message, title = '') => addNotification({ type: 'success', title, message })
  const error = (message, title = '') => addNotification({ type: 'error', title, message })
  const warning = (message, title = '') => addNotification({ type: 'warning', title, message })
  const info = (message, title = '') => addNotification({ type: 'info', title, message })

  return {
    notifications,
    addNotification,
    removeNotification,
    success,
    error,
    warning,
    info
  }
})
