<template>
  <transition name="modal">
    <div 
      v-if="confirmStore.isOpen" 
      class="fixed inset-0 z-[9999] flex items-center justify-center p-4"
    >
      <!-- Backdrop Overlay -->
      <div 
        class="absolute inset-0 bg-slate-900/60 backdrop-blur-sm"
        @click="confirmStore.handleCancel"
      ></div>

      <!-- Modal Card -->
      <div 
        class="bg-white rounded-2xl w-full max-w-sm p-6 shadow-2xl relative z-10 border border-slate-100 transform transition-all duration-300"
      >
        <div class="text-center">
          <!-- Icon Header -->
          <div 
            v-if="confirmStore.type === 'danger'" 
            class="mx-auto flex items-center justify-center h-14 w-14 rounded-full bg-red-50 text-red-600 mb-4"
          >
            <svg class="h-7 w-7" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
            </svg>
          </div>

          <div 
            v-else-if="confirmStore.type === 'warning'" 
            class="mx-auto flex items-center justify-center h-14 w-14 rounded-full bg-amber-50 text-amber-600 mb-4"
          >
            <svg class="h-7 w-7" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
            </svg>
          </div>

          <div 
            v-else 
            class="mx-auto flex items-center justify-center h-14 w-14 rounded-full bg-teal-50 text-teal-600 mb-4"
          >
            <svg class="h-7 w-7" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>

          <!-- Title & Message -->
          <h3 class="text-lg font-bold text-slate-900 mb-2">
            {{ confirmStore.title }}
          </h3>
          <p class="text-sm text-slate-500 leading-relaxed px-2">
            {{ confirmStore.message }}
          </p>
        </div>

        <!-- Buttons -->
        <div class="flex gap-3 mt-6">
          <button 
            @click="confirmStore.handleCancel"
            class="flex-1 px-4 py-2.5 border border-gray-200 rounded-xl text-sm font-semibold text-slate-600 hover:bg-slate-50 active:bg-slate-100 transition-colors duration-200"
          >
            {{ confirmStore.cancelText }}
          </button>
          <button 
            @click="confirmStore.handleConfirm"
            class="flex-1 px-4 py-2.5 rounded-xl text-sm font-bold text-white transition-all duration-200 shadow-md active:scale-[0.98]"
            :class="[
              confirmStore.type === 'danger' ? 'bg-red-600 hover:bg-red-700 hover:shadow-red-200' : '',
              confirmStore.type === 'warning' ? 'bg-amber-600 hover:bg-amber-700 hover:shadow-amber-200' : '',
              confirmStore.type === 'info' ? 'bg-teal-600 hover:bg-teal-700 hover:shadow-teal-200' : '',
            ]"
          >
            {{ confirmStore.confirmText }}
          </button>
        </div>
      </div>
    </div>
  </transition>
</template>

<script>
import { useConfirmStore } from '../stores/confirm'

export default {
  name: 'ConfirmModal',
  setup() {
    const confirmStore = useConfirmStore()
    return {
      confirmStore
    }
  }
}
</script>

<style scoped>
.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.3s ease;
}

.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}

.modal-enter-active div:last-child {
  animation: scale-up 0.3s cubic-bezier(0.34, 1.56, 0.64, 1);
}

.modal-leave-active div:last-child {
  animation: scale-down 0.2s cubic-bezier(0.34, 1.56, 0.64, 1) reverse;
}

@keyframes scale-up {
  from {
    transform: scale(0.95);
    opacity: 0;
  }
  to {
    transform: scale(1);
    opacity: 1;
  }
}

@keyframes scale-down {
  from {
    transform: scale(1);
    opacity: 1;
  }
  to {
    transform: scale(0.95);
    opacity: 0;
  }
}
</style>
