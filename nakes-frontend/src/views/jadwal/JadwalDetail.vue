<template>
  <LayoutWrapper>
    <div class="min-h-screen bg-gray-50">
      <!-- Header -->
      <div class="bg-white border-b border-gray-200 px-4 md:px-8 py-4">
        <div class="flex items-center gap-2 md:gap-4">
          <div>
            <h2 class="text-base md:text-lg font-bold text-slate-900">Detail Jadwal Obat</h2>
            <p class="text-xs text-gray-500">Informasi lengkap jadwal minum obat</p>
          </div>
        </div>
      </div>

      <!-- Main Content -->
      <div class="p-4 md:p-8 max-w-4xl mx-auto">
        <div v-if="!loading && jadwal" class="space-y-6">
          <!-- Patient Info Card -->
          <div class="bg-white rounded-xl shadow-sm border border-gray-100 p-6">
            <h3 class="text-sm font-semibold text-gray-700 uppercase tracking-wide mb-4">Informasi Pasien</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <p class="text-xs text-gray-500 mb-1">Nama Pasien</p>
                <p class="text-sm font-semibold text-slate-900">{{ jadwal.pasien_nama }}</p>
              </div>
              <div>
                <p class="text-xs text-gray-500 mb-1">Kode Pasien</p>
                <p class="text-sm font-semibold text-slate-900">{{ jadwal.pasien_kode || '-' }}</p>
              </div>
            </div>
          </div>

          <!-- Medicine Info Card -->
          <div class="bg-white rounded-xl shadow-sm border border-gray-100 p-6">
            <h3 class="text-sm font-semibold text-gray-700 uppercase tracking-wide mb-4">Informasi Obat</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <p class="text-xs text-gray-500 mb-1">Nama Obat</p>
                <p class="text-sm font-semibold text-slate-900">{{ jadwal.nama_obat }}</p>
              </div>
              <div>
                <p class="text-xs text-gray-500 mb-1">Dosis</p>
                <p class="text-sm font-semibold text-slate-900">{{ jadwal.jumlah_dosis }} {{ jadwal.satuan }}</p>
              </div>
              <div>
                <p class="text-xs text-gray-500 mb-1">Kategori Obat</p>
                <p class="text-sm font-semibold text-slate-900">{{ jadwal.kategori_obat || '-' }}</p>
              </div>
              <div>
                <p class="text-xs text-gray-500 mb-1">Takaran Obat</p>
                <p class="text-sm font-semibold text-slate-900">{{ jadwal.takaran_obat || '-' }}</p>
              </div>
            </div>
          </div>

          <!-- Consumption Schedule Card -->
          <div class="bg-white rounded-xl shadow-sm border border-gray-100 p-6">
            <h3 class="text-sm font-semibold text-gray-700 uppercase tracking-wide mb-4">Jadwal Konsumsi</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <p class="text-xs text-gray-500 mb-1">Frekuensi per Hari</p>
                <p class="text-sm font-semibold text-slate-900">{{ jadwal.frekuensi_per_hari }}</p>
              </div>
              <div>
                <p class="text-xs text-gray-500 mb-1">Waktu Minum</p>
                <p class="text-sm font-semibold text-slate-900">{{ jadwal.waktu_minum }}</p>
              </div>
              <div>
                <p class="text-xs text-gray-500 mb-1">Aturan Konsumsi</p>
                <p class="text-sm font-semibold text-slate-900">{{ jadwal.aturan_konsumsi || '-' }}</p>
              </div>
              <div>
                <p class="text-xs text-gray-500 mb-1">Tanggal Mulai</p>
                <p class="text-sm font-semibold text-slate-900">{{ formatDate(jadwal.tanggal_mulai) }}</p>
              </div>
            </div>
          </div>

          <!-- Duration Card -->
          <div class="bg-white rounded-xl shadow-sm border border-gray-100 p-6">
            <h3 class="text-sm font-semibold text-gray-700 uppercase tracking-wide mb-4">Durasi & Jadwal</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <p class="text-xs text-gray-500 mb-1">Tipe Durasi</p>
                <p class="text-sm font-semibold text-slate-900">
                  <span v-if="jadwal.tipe_durasi === 'hari'" class="inline-block bg-blue-50 text-blue-700 px-3 py-1 rounded-full text-xs font-medium">
                    {{ jadwal.jumlah_hari }} hari
                  </span>
                  <span v-else class="inline-block bg-gray-100 text-gray-600 px-3 py-1 rounded-full text-xs font-medium">
                    Rutin
                  </span>
                </p>
              </div>
              <div v-if="jadwal.tanggal_selesai">
                <p class="text-xs text-gray-500 mb-1">Tanggal Selesai</p>
                <p class="text-sm font-semibold text-slate-900">{{ formatDate(jadwal.tanggal_selesai) }}</p>
              </div>
            </div>
          </div>

          <!-- Notes Card -->
          <div v-if="jadwal.catatan" class="bg-white rounded-xl shadow-sm border border-gray-100 p-6">
            <h3 class="text-sm font-semibold text-gray-700 uppercase tracking-wide mb-4">Catatan</h3>
            <p class="text-sm text-slate-900 leading-relaxed">{{ jadwal.catatan }}</p>
          </div>

          <!-- Action Buttons -->
          <div class="flex justify-end">
            <button 
              @click="goBack"
              class="px-6 py-3 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50 transition font-medium text-sm">
              ← Kembali
            </button>
          </div>
        </div>

        <!-- Loading State -->
        <div v-else-if="loading" class="flex items-center justify-center py-12">
          <div class="text-center">
            <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-teal-600 mb-4"></div>
            <p class="text-gray-500 text-sm">Memuat data...</p>
          </div>
        </div>

        <!-- Error State -->
        <div v-else class="text-center py-12">
          <svg class="w-12 h-12 text-gray-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"
              d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"/>
          </svg>
          <p class="text-gray-400 text-sm">Data jadwal tidak ditemukan</p>
        </div>
      </div>
    </div>
  </LayoutWrapper>
</template>

<script>
import { onMounted, ref, computed } from 'vue'
import LayoutWrapper from '../../components/LayoutWrapper.vue'
import { useJadwalStore } from '../../stores/jadwal'
import { useRouter, useRoute } from 'vue-router'

export default {
  name: 'JadwalDetail',
  components: {
    LayoutWrapper
  },
  setup() {
    const router = useRouter()
    const route = useRoute()
    const jadwalStore = useJadwalStore()
    const loading = ref(false)

    const jadwal = computed(() => {
      const id = parseInt(route.params.id)
      return jadwalStore.jadwalList.find(j => j.id === id)
    })

    const formatDate = (dateStr) => {
      if (!dateStr) return '-'
      return new Date(dateStr).toLocaleDateString('id-ID', {
        year: 'numeric',
        month: 'long',
        day: 'numeric'
      })
    }

    const goBack = () => {
      router.back()
    }

    const goToEdit = () => {
      router.push({ name: 'jadwal-edit', params: { id: route.params.id } })
    }

    onMounted(() => {
      if (jadwalStore.jadwalList.length === 0) {
        loading.value = true
        jadwalStore.fetchJadwals().finally(() => {
          loading.value = false
        })
      }
    })

    return {
      jadwal,
      loading,
      formatDate,
      goBack,
      goToEdit
    }
  }
}
</script>
