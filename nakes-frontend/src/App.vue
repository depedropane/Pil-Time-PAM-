<template>
  <router-view />
  <NotificationToast />
  <ConfirmModal />
</template>

<script>
import NotificationToast from './components/NotificationToast.vue'
import ConfirmModal from './components/ConfirmModal.vue'

export default {
  name: 'App',
  components: {
    NotificationToast,
    ConfirmModal
  }
}
</script>
