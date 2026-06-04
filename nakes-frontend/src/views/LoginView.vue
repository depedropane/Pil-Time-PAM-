<template>
  <div class="min-h-screen bg-gray-50 flex">
    <!-- Left Side - Dark Sidebar -->
    <div class="hidden lg:flex lg:w-2/5 bg-gradient-to-b from-slate-900 via-slate-800 to-slate-900 flex-col justify-center px-10 py-12">
      <!-- Logo & Title -->
      <div class="flex items-center gap-3 mb-10">
        <div class="w-14 h-14 rounded-full overflow-hidden bg-gradient-to-br from-teal-400 to-cyan-500 flex items-center justify-center">
          <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4"></path>
          </svg>
        </div>
        <div>
          <h1 class="text-xl font-bold text-white leading-tight">Pil Time</h1>
          <p class="text-slate-400 text-xs">Portal Tenaga Kesehatan</p>
        </div>
      </div>

      <!-- Headline -->
      <div class="mb-8">
        <h2 class="text-3xl font-bold text-white leading-snug mb-4">
          Kelola Pasien<br>
          dengan <span class="text-teal-400">Lebih<br>Mudah</span>
        </h2>
        <p class="text-slate-400 text-sm leading-relaxed">
          Platform terintegrasi untuk monitoring kepatuhan minum obat pasien Desa Nauli secara real-time.
        </p>
      </div>

      <!-- Feature List -->
      <div class="space-y-5">
        <div class="flex items-start gap-3">
          <div class="w-8 h-8 flex items-center justify-center flex-shrink-0 mt-0.5">
            <svg class="w-5 h-5 text-teal-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
            </svg>
          </div>
          <div>
            <p class="text-white text-sm font-semibold">Input Jadwal Obat</p>
            <p class="text-slate-400 text-xs mt-0.5">Atur jadwal dan dosis obat pasien langsung dari sistem</p>
          </div>
        </div>

        <div class="flex items-start gap-3">
          <div class="w-8 h-8 flex items-center justify-center flex-shrink-0 mt-0.5">
            <svg class="w-5 h-5 text-teal-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
          <div>
            <p class="text-white text-sm font-semibold">Monitoring Kepatuhan</p>
            <p class="text-slate-400 text-xs mt-0.5">Pantau riwayat konsumsi obat seluruh pasien secara akurat</p>
          </div>
        </div>

        <div class="flex items-start gap-3">
          <div class="w-8 h-8 flex items-center justify-center flex-shrink-0 mt-0.5">
            <svg class="w-5 h-5 text-teal-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path>
            </svg>
          </div>
          <div>
            <p class="text-white text-sm font-semibold">Notifikasi Otomatis</p>
            <p class="text-slate-400 text-xs mt-0.5">Pasien mendapat pengingat tepat waktu sesuai jadwal</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Right Side - Form -->
    <div class="w-full lg:w-3/5 flex flex-col items-center justify-center px-6 py-12 sm:px-12">
      <div class="w-full max-w-md">
        <!-- Header -->
        <div class="text-center mb-8">
          <h2 class="text-3xl font-bold text-slate-900 mb-2">Selamat Datang</h2>
          <p class="text-gray-500 text-sm">Masuk ke portal tenaga kesehatan Pil Time</p>
        </div>

        <!-- Error Message -->
        <div v-if="errorMessage" class="mb-6 p-4 bg-red-50 border border-red-200 rounded-lg flex items-start justify-between">
          <div class="flex items-start space-x-3 flex-1">
            <svg class="w-5 h-5 text-red-500 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd"></path>
            </svg>
            <span class="text-sm text-red-700">{{ errorMessage }}</span>
          </div>
          <button 
            type="button"
            @click="errorMessage = ''"
            class="text-red-500 hover:text-red-700 ml-3 flex-shrink-0"
          >
            <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"></path>
            </svg>
          </button>
        </div>

        <!-- Form -->
        <form @submit.prevent="handleLogin" class="space-y-5">
          <!-- Email -->
          <div>
            <label for="email" class="block text-sm font-medium text-slate-700 mb-1.5">Email</label>
            <input
              id="email"
              v-model="form.email"
              type="email"
              required
              @input="errorMessage = ''"
              class="w-full px-4 py-2.5 border border-gray-300 rounded-lg focus:ring-2 focus:ring-teal-500 focus:border-transparent outline-none transition text-sm"
              placeholder="Enter your email"
            />
          </div>

          <!-- Password -->
          <div>
            <label for="password" class="block text-sm font-medium text-slate-700 mb-1.5">Password</label>
            <div class="relative">
              <input
                id="password"
                v-model="form.password"
                :type="showPassword ? 'text' : 'password'"
                required
                @input="errorMessage = ''"
                class="w-full px-4 py-2.5 border border-gray-300 rounded-lg focus:ring-2 focus:ring-teal-500 focus:border-transparent outline-none transition pr-10 text-sm"
                placeholder="Enter your password"
              />
              <button
                type="button"
                @click="showPassword = !showPassword"
                class="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600 transition"
              >
                <svg v-if="showPassword" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M10 12a2 2 0 100-4 2 2 0 000 4z"></path>
                  <path fill-rule="evenodd" d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z" clip-rule="evenodd"></path>
                </svg>
                <svg v-else class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M3.707 2.293a1 1 0 00-1.414 1.414l14 14a1 1 0 001.414-1.414l-1.473-1.473A10.014 10.014 0 0019.542 10C18.268 5.943 14.478 3 10 3a9.958 9.958 0 00-4.512 1.074l-1.78-1.781zm4.261 4.26l1.514 1.515a2.003 2.003 0 012.45 2.45l1.514 1.514a4 4 0 00-5.478-5.478z" clip-rule="evenodd"></path>
                  <path d="M12.454 16.697L9.75 13.992a4 4 0 01-3.742-3.741L2.335 6.578A9.98 9.98 0 00.458 10c1.274 4.057 5.064 7 9.542 7 .847 0 1.669-.105 2.454-.303z"></path>
                </svg>
              </button>
            </div>
          </div>

          <!-- Login Button -->
          <button
            type="submit"
            :disabled="loading"
            class="w-full px-4 py-2.5 mt-2 font-semibold text-white bg-gradient-to-r from-teal-500 to-cyan-500 rounded-lg hover:from-teal-600 hover:to-cyan-600 disabled:opacity-50 disabled:cursor-not-allowed transition text-sm"
          >
            {{ loading ? 'Signing in...' : 'Login' }}
          </button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import { authService } from '../services'

export default {
  name: 'LoginView',
  setup() {
    const router = useRouter()
    const authStore = useAuthStore()
    const loading = ref(false)
    const errorMessage = ref('')
    const showPassword = ref(false)
    const form = ref({
      email: '',
      password: ''
    })

    const handleLogin = async () => {
      if (!form.value.email || !form.value.password) {
        errorMessage.value = 'Email dan password harus diisi'
        return
      }

      loading.value = true
      errorMessage.value = ''

      try {
        const axiosResponse = await authService.loginAdmin(form.value.email, form.value.password)
        const response = axiosResponse.data

        // Auth-service mengembalikan: { token, message, data: { nakes_id, nama, email } }
        if (response.token) {
          const userData = {
            nakes_id: response.data?.nakes_id,
            nama: response.data?.nama,
            email: response.data?.email
          }

          authStore.setAuth(response.token, userData)
          router.push('/dashboard')
        } else {
          errorMessage.value = response.message || 'Login gagal'
          loading.value = false
        }
      } catch (error) {
        const errMsg = error.response?.data?.message || error.message || 'Terjadi kesalahan saat login'
        errorMessage.value = errMsg
        loading.value = false
        console.error('Login error:', error)
      }
    }

    return {
      form,
      loading,
      errorMessage,
      showPassword,
      handleLogin
    }
  }
}
</script>