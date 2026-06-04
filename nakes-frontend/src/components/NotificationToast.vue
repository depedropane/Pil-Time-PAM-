<template>
  <div class="fixed top-5 right-5 z-[9999] flex flex-col gap-3 max-w-sm md:max-w-md w-full pointer-events-none">
    <transition-group 
      name="toast" 
      tag="div" 
      class="flex flex-col gap-3 w-full"
    >
      <div 
        v-for="notif in notificationStore.notifications" 
        :key="notif.id"
        class="pointer-events-auto w-full rounded-xl border p-4 shadow-lg transition-all duration-300 flex items-start gap-3"
        :class="[
          notif.type === 'success' ? 'bg-[#F0FDF4] border-[#BBF7D0] text-slate-800' : '',
          notif.type === 'error' ? 'bg-[#FEF2F2] border-[#FCA5A5] text-slate-800' : '',
          notif.type === 'warning' || notif.type === 'info' ? 'bg-[#FFFDF0] border-[#FEF08A] text-slate-800' : ''
        ]"
      >
        <!-- Icons -->
        <!-- Success Check Icon -->
        <span v-if="notif.type === 'success'" class="text-[#166534] mt-0.5 flex-shrink-0">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
        </span>

        <!-- Error Warning Icon -->
        <span v-else-if="notif.type === 'error'" class="text-[#991B1B] mt-0.5 flex-shrink-0">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
          </svg>
        </span>

        <!-- Info/Warning Icon -->
        <span v-else class="text-[#854D0E] mt-0.5 flex-shrink-0">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
        </span>

        <!-- Content -->
        <div class="flex-1">
          <!-- Large Layout (Title exists) -->
          <div v-if="notif.title">
            <h4 class="font-bold text-sm"
              :class="[
                notif.type === 'success' ? 'text-[#166534]' : '',
                notif.type === 'error' ? 'text-[#991B1B]' : '',
                notif.type === 'warning' || notif.type === 'info' ? 'text-[#854D0E]' : ''
              ]"
            >
              {{ notif.title }}
            </h4>
            <p class="text-xs text-slate-600 mt-1 leading-relaxed">{{ notif.message }}</p>
          </div>

          <!-- Small Layout (Only Message) -->
          <div v-else class="text-sm font-semibold"
            :class="[
              notif.type === 'success' ? 'text-[#166534]' : '',
              notif.type === 'error' ? 'text-[#991B1B]' : '',
              notif.type === 'warning' || notif.type === 'info' ? 'text-[#854D0E]' : ''
            ]"
          >
            {{ notif.message }}
          </div>
        </div>

        <!-- Close Button -->
        <button 
          @click="notificationStore.removeNotification(notif.id)"
          class="text-slate-400 hover:text-slate-600 transition-colors duration-200 mt-0.5"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>
    </transition-group>
  </div>
</template>

<script>
import { useNotificationStore } from '../stores/notification'

export default {
  name: 'NotificationToast',
  setup() {
    const notificationStore = useNotificationStore()
    return {
      notificationStore
    }
  }
}
</script>

<style scoped>
.toast-enter-active,
.toast-leave-active {
  transition: all 0.35s cubic-bezier(0.25, 1, 0.5, 1);
}

.toast-enter-from {
  opacity: 0;
  transform: translateY(-20px) scale(0.95);
}

.toast-leave-to {
  opacity: 0;
  transform: translateX(40px) scale(0.95);
}
</style>
