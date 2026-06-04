<template>
  <LayoutWrapper>
    <div class="min-h-screen bg-[#f8fafc] p-8">
      <h1 class="text-[28px] font-bold text-slate-800">Dashboard</h1>
      <p class="text-slate-500 mt-1 mb-8">Selamat datang di Portal Tenaga Kesehatan </p>

      <!-- Stats Grid -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4 md:gap-6 mb-6 md:mb-8">
        <!-- Total Pasien Card -->
        <div class="bg-white rounded-[20px] shadow-[0_8px_30px_rgb(0,0,0,0.04)] border border-slate-100 p-6">
          <div class="flex justify-between items-center">
            <div>
              <p class="text-slate-500 text-sm">Total Pasien</p>
              <p class="text-3xl font-bold text-slate-800 mt-1">{{ stats.totalPasien }}</p>
            </div>

            <div class="w-12 h-12 bg-teal-50 rounded-xl flex items-center justify-center">
              <svg class="w-5 h-5 md:w-6 md:h-6 text-blue-600" fill="none" stroke="#0F9B76" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M12 14c-4 0-7 2-7 4v2h14v-2c0-2-3-4-7-4z"/>
                <circle cx="12" cy="7" r="4" stroke-width="2"/>
              </svg>
            </div>
          </div>
        </div>

        <!-- Total Jadwal Card -->
        <div class="bg-white rounded-[20px] shadow-[0_8px_30px_rgb(0,0,0,0.04)] border border-slate-100 p-6">
          <div class="flex justify-between items-center">
            <div>
              <p class="text-slate-500 text-sm">Total Jadwal</p>
              <p class="text-3xl font-bold text-slate-800 mt-1">{{ stats.totalJadwal }}</p>
            </div>

            <div class="w-12 h-12 bg-teal-50 rounded-xl flex items-center justify-center">
              <svg class="w-5 h-5 md:w-6 md:h-6 text-green-600" fill="none" stroke="#0F9B76" viewBox="0 0 24 24">
                <rect x="3" y="5" width="18" height="16" rx="2" stroke-width="2"/>
                <line x1="3" y1="10" x2="21" y2="10" stroke-width="2"/>
                <line x1="8" y1="3" x2="8" y2="7" stroke-width="2"/>
                <line x1="16" y1="3" x2="16" y2="7" stroke-width="2"/>
              </svg>
            </div>
          </div>
        </div>
      </div>

      <!-- Quick Actions -->
      <div class="bg-white rounded-[20px] shadow-[0_8px_30px_rgb(0,0,0,0.04)] border border-slate-100 p-6">
        <h2 class="text-base md:text-lg font-semibold text-slate-900 mb-4">Akses Cepat</h2>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-3 md:gap-4">
          <router-link 
            to="/pasien"
            class="block p-4 bg-white rounded-[16px] border border-slate-100 hover:shadow-md transition"
          >
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 bg-teal-50 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 md:w-6 md:h-6 text-white" fill="none" stroke="#0F9B76" viewBox="0 0 24 24">
                  <circle cx="12" cy="8" r="3" stroke-width="2"/>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M6 20c0-3 3-5 6-5s6 2 6 5"/>
                </svg>
              </div>

              <div>
                <p class="font-semibold text-slate-800">Kelola Pasien</p>
                <p class="text-xs text-slate-400">Manage patients</p>
              </div>
            </div>
          </router-link>

          <router-link 
            to="/jadwal"
            class="block p-4 bg-white rounded-[16px] border border-slate-100 hover:shadow-md transition"
          >
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 bg-teal-50 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 md:w-6 md:h-6 text-white" fill="none" stroke="#0F9B76" viewBox="0 0 24 24">
                  <rect x="3" y="5" width="18" height="16" rx="2" stroke-width="2"/>
                  <line x1="3" y1="10" x2="21" y2="10" stroke-width="2"/>
                </svg>
              </div>

              <div>
                <p class="font-semibold text-slate-800">Kelola Jadwal</p>
                <p class="text-xs text-slate-400">Manage patients</p>
              </div>
            </div>
          </router-link>

          <router-link 
            to="/obat"
            class="block p-4 bg-white rounded-[16px] border border-slate-100 hover:shadow-md transition"
          >
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 bg-teal-50 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 md:w-6 md:h-6 text-white" fill="none" stroke="#0F9B76" viewBox="0 0 24 24">
                  <rect x="7" y="3" width="10" height="18" rx="5" stroke-width="2"/>
                  <line x1="7" y1="12" x2="17" y2="12" stroke-width="2"/>
                </svg>
              </div>

              <div>
                <p class="font-semibold text-slate-800">Info Obat</p>
                <p class="text-xs text-slate-400">Manage patients</p>
              </div>
            </div>
          </router-link>
        </div>
      </div>
    </div>
  </LayoutWrapper>
</template>

<script>
import { ref, onMounted } from 'vue'
import LayoutWrapper from '../components/LayoutWrapper.vue'
import { dashboardService } from '../services'

export default {
  name: 'DashboardView',
  components: {
    LayoutWrapper
  },
  setup() {
    const stats = ref({
      totalPasien: 0,
      totalJadwal: 0
    })
    const loading = ref(false)

    const loadStats = async () => {
      loading.value = true
      try {
        const response = await dashboardService.getDashboard()
        if (response.data) {
          stats.value = {
            totalPasien: response.data.total_pasien || 0,
            totalJadwal: response.data.total_jadwal || 0
          }
        }
      } catch (error) {
        console.error('Failed to load dashboard stats:', error)
      } finally {
        loading.value = false
      }
    }

    onMounted(() => {
      loadStats()
    })

    return {
      stats,
      loading
    }
  }
}
</script>