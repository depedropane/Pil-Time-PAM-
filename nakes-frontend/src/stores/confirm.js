import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useConfirmStore = defineStore('confirm', () => {
  const isOpen = ref(false)
  const title = ref('Konfirmasi')
  const message = ref('')
  const confirmText = ref('Ya')
  const cancelText = ref('Batal')
  const type = ref('danger') // 'danger' | 'warning' | 'info'
  
  let resolvePromise = null

  const show = (options = {}) => {
    title.value = options.title || 'Konfirmasi'
    message.value = options.message || ''
    confirmText.value = options.confirmText || 'Ya'
    cancelText.value = options.cancelText || 'Batal'
    type.value = options.type || 'danger'
    isOpen.value = true

    return new Promise((resolve) => {
      resolvePromise = resolve
    })
  }

  const handleConfirm = () => {
    isOpen.value = false
    if (resolvePromise) resolvePromise(true)
  }

  const handleCancel = () => {
    isOpen.value = false
    if (resolvePromise) resolvePromise(false)
  }

  return {
    isOpen,
    title,
    message,
    confirmText,
    cancelText,
    type,
    show,
    handleConfirm,
    handleCancel
  }
})
