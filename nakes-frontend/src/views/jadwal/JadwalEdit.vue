<template>
  <LayoutWrapper>
    <div class="min-h-screen bg-gray-50">
      <!-- Header -->
      <div class="bg-white border-b border-gray-200 px-4 md:px-8 py-4">
        <div class="flex items-center gap-2 md:gap-4">
          <div>
            <h2 class="text-base md:text-lg font-bold text-slate-900">Edit Jadwal Obat</h2>
          </div>
        </div>
      </div>

      <!-- Main Content -->
      <div class="p-4 md:p-8">
        <div v-if="!loading && editForm" class="space-y-6">
          <!-- Info bar -->
          <div class="flex flex-col md:flex-row md:items-center md:gap-6 mb-6 bg-white rounded-xl px-4 md:px-6 py-4 shadow-sm border border-gray-100">
            <div class="flex items-center gap-3">
              <div>
                <p class="text-xs text-gray-400 uppercase tracking-wide">Pasien</p>
                <p class="text-sm font-semibold text-slate-900">{{ jadwal?.pasien_nama }}</p>
              </div>
            </div>
          </div>

          <!-- Edit Form -->
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4 md:gap-6">
            <!-- Left: Adjustable Fields -->
            <div class="bg-white rounded-xl shadow-sm border border-gray-100 p-4 md:p-6">
              <div class="flex items-center gap-3 mb-5">
                <div>
                  <h1 class="title font-bold text-slate-900">Informasi Obat</h1>
                </div>
              </div>

              <div class="space-y-3 md:space-y-5">
                <div>
                  <label class="block text-xs md:text-sm font-medium text-gray-700 mb-1">
                    Nama Obat <span class="text-red-500">*</span>
                  </label>
                  <input v-model="editForm.nama_obat" type="text" placeholder="Nama obat"
                    class="w-full px-3 py-2 text-sm border border-gray-200 rounded-lg focus:ring-2 focus:ring-teal-500 focus:border-transparent outline-none"/>
                </div>

                <div class="grid grid-cols-2 gap-3">
                  <div>
                    <label class="block text-xs md:text-sm font-medium text-gray-700 mb-1">
                      Dosis <span class="text-red-500">*</span>
                    </label>
                    <input v-model="editForm.jumlah_dosis" type="text" placeholder="1"
                      class="w-full px-3 py-2 text-sm border border-gray-200 rounded-lg focus:ring-2 focus:ring-teal-500 focus:border-transparent outline-none"/>
                  </div>
                  <div>
                    <label class="block text-xs md:text-sm font-medium text-gray-700 mb-1">
                      Satuan <span class="text-red-500">*</span>
                    </label>
                    <input v-model="editForm.satuan" type="text" placeholder="tablet"
                      class="w-full px-3 py-2 text-sm border border-gray-200 rounded-lg focus:ring-2 focus:ring-teal-500 focus:border-transparent outline-none"/>
                  </div>
                </div>

                <div>
                  <label class="block text-xs md:text-sm font-medium text-gray-700 mb-1">
                    Frekuensi per Hari <span class="text-red-500">*</span>
                  </label>
                  <input 
                    type="text" 
                    :value="selectedWaktuMinum.length + 'x sehari'" 
                    class="w-full px-3 py-2 text-sm border border-gray-200 rounded-lg bg-gray-50 text-gray-500 outline-none" 
                    readonly
                  />
                </div>

                <div>
                  <label class="block text-xs md:text-sm font-medium text-gray-700 mb-2">
                    Waktu Minum <span class="text-red-500">*</span>
                  </label>
                  <div class="flex flex-wrap gap-2 md:gap-3">
                    <button
                      v-for="waktu in wakteMinum" :key="waktu.value"
                      type="button"
                      @click="toggleWaktuMinum(waktu.value)"
                      :class="[
                        'flex items-center gap-2 px-3 md:px-4 py-2 rounded-lg text-xs md:text-sm font-medium border transition flex-1',
                        selectedWaktuMinum.includes(waktu.value)
                          ? 'bg-teal-600 text-white border-teal-600'
                          : 'bg-white text-gray-600 border-gray-200 hover:border-teal-300'
                      ]">
                      <span>{{ waktu.icon }}</span>
                      {{ waktu.label }}
                      <svg v-if="selectedWaktuMinum.includes(waktu.value)" class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"/>
                      </svg>
                    </button>
                  </div>
                </div>

                <!-- Waktu Pengingat Section -->
                <div v-if="selectedWaktuMinum.length > 0" class="p-4 bg-slate-50 rounded-xl border border-slate-100">
                  <label class="block text-xs md:text-sm font-semibold text-slate-700 mb-2">Waktu Pengingat</label>
                  <div class="grid grid-cols-1 sm:grid-cols-3 gap-3">
                    <div v-if="selectedWaktuMinum.includes('Pagi')" class="space-y-1 bg-white p-3 rounded-lg border border-gray-150 shadow-sm text-center">
                      <span class="text-[10px] font-bold text-slate-400 uppercase tracking-wider block">🌅 Pagi</span>
                      <input type="time" v-model="waktuReminder.Pagi" class="w-full text-center border-none focus:ring-0 outline-none text-sm font-semibold text-gray-700" />
                    </div>
                    <div v-if="selectedWaktuMinum.includes('Siang')" class="space-y-1 bg-white p-3 rounded-lg border border-gray-150 shadow-sm text-center">
                      <span class="text-[10px] font-bold text-slate-400 uppercase tracking-wider block">☀️ Siang</span>
                      <input type="time" v-model="waktuReminder.Siang" class="w-full text-center border-none focus:ring-0 outline-none text-sm font-semibold text-gray-700" />
                    </div>
                    <div v-if="selectedWaktuMinum.includes('Malam')" class="space-y-1 bg-white p-3 rounded-lg border border-gray-150 shadow-sm text-center">
                      <span class="text-[10px] font-bold text-slate-400 uppercase tracking-wider block">🌙 Malam</span>
                      <input type="time" v-model="waktuReminder.Malam" class="w-full text-center border-none focus:ring-0 outline-none text-sm font-semibold text-gray-700" />
                    </div>
                  </div>
                </div>

                <div>
                  <label class="block text-xs md:text-sm font-medium text-gray-700 mb-2">
                    Aturan Konsumsi <span class="text-red-500">*</span>
                  </label>
                  <div class="flex flex-wrap gap-2">
                    <button v-for="aturan in aturanKonsumsi" :key="aturan" type="button"
                      @click="editForm.aturan_konsumsi = aturan"
                      :class="[
                        'px-3 py-2 rounded-lg text-xs md:text-sm font-medium border transition',
                        editForm.aturan_konsumsi === aturan
                          ? 'bg-teal-600 text-white border-teal-600'
                          : 'bg-white text-gray-600 border-gray-200 hover:border-teal-300'
                      ]">
                      {{ aturan }}
                    </button>
                  </div>
                </div>

                <div>
                  <label class="block text-xs md:text-sm font-medium text-gray-700 mb-1">Catatan</label>
                  <textarea v-model="editForm.catatan" placeholder="Catatan untuk pasien"
                    rows="4"
                    class="w-full px-3 py-2 text-sm border border-gray-200 rounded-lg focus:ring-2 focus:ring-teal-500 focus:border-transparent outline-none resize-none">
                  </textarea>
                </div>
              </div>
            </div>

            <!-- Right: Durasi Info -->
            <div class="bg-white rounded-xl shadow-sm border border-gray-100 p-4 md:p-6">
              <div class="flex items-center gap-3 mb-5">
                <div>
                  <h2 class="totle font-bold text-slate-900">Durasi & Jadwal</h2>
                </div>
              </div>

              <div class="space-y-5">
                <!-- Edit Duration Type -->
                <div>
                  <label class="block text-xs md:text-sm font-medium text-gray-700 mb-2">
                    Tipe Durasi <span class="text-red-500">*</span>
                  </label>
                  <div class="flex gap-2">
                    <button type="button" @click="editForm.tipe_durasi = 'hari'"
                      :class="[
                        'flex-1 px-3 py-2 rounded-lg text-xs md:text-sm font-medium border transition',
                        editForm.tipe_durasi === 'hari'
                          ? 'bg-teal-500 text-white border-teal-300'
                          : 'bg-white text-gray-600 border-gray-200 hover:bg-teal-600'
                      ]">
                      Hari
                    </button>
                    <button type="button" @click="editForm.tipe_durasi = 'rutin'"
                      :class="[
                        'flex-1 px-3 py-2 rounded-lg text-xs md:text-sm font-medium border transition',
                        editForm.tipe_durasi === 'rutin'
                          ? 'bg-teal-500 text-white border-teal-300'
                          : 'bg-white text-gray-600 border-gray-200 hover:bg-teal-600'
                      ]">
                      Rutin
                    </button>
                  </div>
                </div>

                <!-- Edit Duration Days (if tipe_durasi is hari) -->
                <div v-if="editForm.tipe_durasi === 'hari'">
                  <label class="block text-xs md:text-sm font-medium text-gray-700 mb-2">
                    Jumlah Hari <span class="text-red-500">*</span>
                  </label>
                  <input v-model="editForm.jumlah_hari" type="number" placeholder="Contoh: 7"
                    class="w-full px-3 py-2 text-sm border border-gray-200 rounded-lg focus:ring-2 focus:border-transparent outline-none"/>
                </div>
                <!-- Jika Rutin -->
                <div v-if="editForm.tipe_durasi === 'rutin'" class="space-y-3">
                  <div class="p-3 bg-gray-50 rounded-lg border border-gray-200">
                    <p class="text-xs text-gray-500">
                      Jadwal akan berjalan setiap hari sesuai waktu yang dipilih.
                    </p>
                  </div>

                  <div class="p-3 bg-teal-50 rounded-lg border border-teal-100">
                    <p class="text-xs text-teal-600">
                      Pengingat akan dikirim 5 menit sebelum waktu minum
                    </p>
                  </div>
                </div>

                <!-- Display Dates -->
                <div class="space-y-3">
                  <div class="p-3 bg-gray-50 rounded-lg border border-gray-200">
                    <p class="text-xs text-gray-500 uppercase font-medium mb-1">Tanggal Mulai</p>
                    <p class="text-sm font-semibold text-slate-900">{{ formatDate(editForm.tanggal_mulai) }}</p>
                  </div>
                  <div v-if="editForm.tanggal_selesai" class="p-3 bg-gray-50 rounded-lg border border-gray-200">
                    <p class="text-xs text-gray-500 uppercase font-medium mb-1">Tanggal Selesai</p>
                    <p class="text-sm font-semibold text-slate-900">{{ formatDate(editForm.tanggal_selesai) }}</p>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Action Buttons -->
          <div class="flex gap-3 flex-col-reverse md:flex-row md:justify-end">
            <button 
              @click="goBack"
              class="flex-1 md:flex-none px-6 py-3 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50 transition font-medium text-sm flex items-center justify-center gap-2">
              ← Batal
            </button>
            <button 
              @click="handleSave"
              :disabled="isSaving"
              class="flex-1 md:flex-none px-6 py-3 bg-teal-600 text-white rounded-lg hover:bg-teal-700 transition font-medium text-sm flex items-center justify-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed">
              <svg v-if="!isSaving" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M5 13l4 4L19 7"/>
              </svg>
              <span v-if="isSaving" class="inline-block animate-spin rounded-full h-4 w-4 border-b-2 border-current"></span>
              {{ isSaving ? 'Menyimpan...' : 'Simpan Perubahan' }}
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
import { useNotificationStore } from '../../stores/notification'

export default {
  name: 'JadwalEdit',
  components: {
    LayoutWrapper
  },
  setup() {
    const router = useRouter()
    const route = useRoute()
    const jadwalStore = useJadwalStore()
    const notificationStore = useNotificationStore()
    const loading = ref(false)
    const isSaving = ref(false)
    const editForm = ref(null)
    const selectedWaktuMinum = ref([])

    const waktuReminder = ref({
      Pagi: '',
      Siang: '',
      Malam: ''
    })

    const aturanKonsumsi = ['Sebelum makan', 'Sesudah makan', 'Bersama makan']
    const wakteMinum = [
      { value: 'Pagi', label: 'Pagi', icon: '☀️' },
      { value: 'Siang', label: 'Siang', icon: '🌤️' },
      { value: 'Malam', label: 'Malam', icon: '🌙' },
    ]

    const jadwal = computed(() => {
      const id = parseInt(route.params.id)
      return jadwalStore.jadwalList.find(j => j.id === id)
    })

    const toggleWaktuMinum = (val) => {
      const idx = selectedWaktuMinum.value.indexOf(val)
      if (idx > -1) {
        selectedWaktuMinum.value.splice(idx, 1)
        waktuReminder.value[val] = ''
      } else {
        selectedWaktuMinum.value.push(val)
        // Set defaults if empty
        if (val === 'Pagi' && !waktuReminder.value.Pagi) waktuReminder.value.Pagi = '07:00'
        if (val === 'Siang' && !waktuReminder.value.Siang) waktuReminder.value.Siang = '13:00'
        if (val === 'Malam' && !waktuReminder.value.Malam) waktuReminder.value.Malam = '19:00'
      }
      
      // Sort waktu in canonical order: Pagi, Siang, Malam
      const order = { 'Pagi': 0, 'Siang': 1, 'Malam': 2 }
      selectedWaktuMinum.value.sort((a, b) => order[a] - order[b])
      
      // Build waktu_minum string
      const jamMinumList = []
      if (selectedWaktuMinum.value.includes('Pagi')) jamMinumList.push(waktuReminder.value.Pagi)
      if (selectedWaktuMinum.value.includes('Siang')) jamMinumList.push(waktuReminder.value.Siang)
      if (selectedWaktuMinum.value.includes('Malam')) jamMinumList.push(waktuReminder.value.Malam)
      editForm.value.waktu_minum = jamMinumList.join(', ')
    }

    const getInitials = (nama) => {
      if (!nama) return '?'
      return nama.split(' ').map(n => n[0]).join('').toUpperCase().slice(0, 2)
    }

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

    const handleSave = async () => {
      if (!editForm.value.nama_obat) {
        notificationStore.error('Nama obat harus diisi', 'Gagal Menyimpan')
        return
      }
      if (!editForm.value.jumlah_dosis) {
        notificationStore.error('Dosis harus diisi', 'Gagal Menyimpan')
        return
      }
      if (!editForm.value.satuan) {
        notificationStore.error('Satuan harus diisi', 'Gagal Menyimpan')
        return
      }
      if (selectedWaktuMinum.value.length === 0) {
        notificationStore.warning('Pilih minimal satu waktu minum', 'Peringatan')
        return
      }
      for (const waktu of selectedWaktuMinum.value) {
        if (!waktuReminder.value[waktu]) {
          notificationStore.error(`Jam untuk ${waktu} harus diisi`, 'Gagal Menyimpan')
          return
        }
      }
      if (!editForm.value.tipe_durasi) {
        notificationStore.error('Tipe durasi harus dipilih', 'Gagal Menyimpan')
        return
      }
      if (editForm.value.tipe_durasi === 'hari' && !editForm.value.jumlah_hari) {
        notificationStore.error('Jumlah hari harus diisi', 'Gagal Menyimpan')
        return
      }

      // Build waktu_minum list
      const jamMinumList = []
      if (selectedWaktuMinum.value.includes('Pagi')) jamMinumList.push(waktuReminder.value.Pagi)
      if (selectedWaktuMinum.value.includes('Siang')) jamMinumList.push(waktuReminder.value.Siang)
      if (selectedWaktuMinum.value.includes('Malam')) jamMinumList.push(waktuReminder.value.Malam)
      editForm.value.waktu_minum = jamMinumList.join(', ')

      const frekuensi_per_hari = String(selectedWaktuMinum.value.length)
      const waktu_reminder_pagi = selectedWaktuMinum.value.includes('Pagi') ? waktuReminder.value.Pagi : ''
      const waktu_reminder_malam = selectedWaktuMinum.value.includes('Malam') ? waktuReminder.value.Malam : ''

      // Check if there are any changes
      const hasChanges =
        editForm.value.nama_obat !== jadwal.value.nama_obat ||
        parseInt(editForm.value.jumlah_dosis) !== parseInt(jadwal.value.jumlah_dosis) ||
        editForm.value.satuan !== jadwal.value.satuan ||
        frekuensi_per_hari !== String(jadwal.value.frekuensi_per_hari) ||
        editForm.value.waktu_minum !== jadwal.value.waktu_minum ||
        editForm.value.aturan_konsumsi !== jadwal.value.aturan_konsumsi ||
        editForm.value.catatan !== jadwal.value.catatan ||
        editForm.value.tipe_durasi !== jadwal.value.tipe_durasi ||
        (editForm.value.tipe_durasi === 'hari' ? parseInt(editForm.value.jumlah_hari) : 0) !== parseInt(jadwal.value.jumlah_hari) ||
        editForm.value.status !== jadwal.value.status ||
        waktu_reminder_pagi !== (jadwal.value.waktu_reminder_pagi || '') ||
        waktu_reminder_malam !== (jadwal.value.waktu_reminder_malam || '')

      if (!hasChanges) {
        notificationStore.info('Tidak ada perubahan untuk disimpan', 'Informasi')
        return
      }

      isSaving.value = true
      try {
        await jadwalStore.updateJadwal(editForm.value.id, {
          nama_obat: editForm.value.nama_obat,
          jumlah_dosis: parseInt(editForm.value.jumlah_dosis),
          satuan: editForm.value.satuan,
          frekuensi_per_hari: frekuensi_per_hari,
          waktu_minum: editForm.value.waktu_minum,
          aturan_konsumsi: editForm.value.aturan_konsumsi,
          catatan: editForm.value.catatan,
          tipe_durasi: editForm.value.tipe_durasi,
          jumlah_hari: editForm.value.tipe_durasi === 'hari' ? parseInt(editForm.value.jumlah_hari) : 0,
          status: editForm.value.status,
          waktu_reminder_pagi: waktu_reminder_pagi,
          waktu_reminder_malam: waktu_reminder_malam,
        })
        notificationStore.success('Jadwal berhasil diperbarui', 'Sukses')
        router.push({ name: 'jadwal' })
      } catch (error) {
        notificationStore.error('Gagal memperbarui jadwal: ' + error.message, 'Gagal')
      } finally {
        isSaving.value = false
      }
    }

    const initializeForm = (data) => {
      editForm.value = { ...data }
      
      // Initialize waktuReminder
      waktuReminder.value = {
        Pagi: data.waktu_reminder_pagi || '',
        Siang: '',
        Malam: data.waktu_reminder_malam || ''
      }
      
      selectedWaktuMinum.value = []
      
      if (data.waktu_minum) {
        // waktu_minum contains actual hours, e.g. "08:00, 13:00, 20:00"
        const times = data.waktu_minum.split(',').map(t => t.trim()).filter(Boolean)
        
        times.forEach(time => {
          const parts = time.split(':')
          if (parts.length === 2) {
            const hour = parseInt(parts[0])
            if (hour >= 4 && hour < 11) {
              selectedWaktuMinum.value.push('Pagi')
              waktuReminder.value.Pagi = time
            } else if (hour >= 11 && hour < 15) {
              selectedWaktuMinum.value.push('Siang')
              waktuReminder.value.Siang = time
            } else {
              selectedWaktuMinum.value.push('Malam')
              waktuReminder.value.Malam = time
            }
          } else {
            // Fallback for non-time strings
            if (['Pagi', 'Siang', 'Malam'].includes(time)) {
              selectedWaktuMinum.value.push(time)
            }
          }
        })
        
        // Sort in canonical order: Pagi, Siang, Malam
        const order = { 'Pagi': 0, 'Siang': 1, 'Malam': 2 }
        selectedWaktuMinum.value.sort((a, b) => order[a] - order[b])
      }
    }

    onMounted(() => {
      if (jadwalStore.jadwalList.length === 0) {
        loading.value = true
        jadwalStore.fetchJadwals().finally(() => {
          loading.value = false
          if (jadwal.value) {
            initializeForm(jadwal.value)
          }
        })
      } else if (jadwal.value) {
        initializeForm(jadwal.value)
      }
    })

    return {
      waktuReminder,
      jadwal,
      editForm,
      loading,
      isSaving,
      aturanKonsumsi,
      wakteMinum,
      selectedWaktuMinum,
      formatDate,
      getInitials,
      goBack,
      handleSave,
      toggleWaktuMinum
    }
  }
}
</script>
