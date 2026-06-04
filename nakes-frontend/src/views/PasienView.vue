<template>
  <LayoutWrapper>
    <div class="p-4 md:p-8">
      <!-- Header -->
      <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-8">
        <h1 class="text-2xl md:text-3xl font-bold text-slate-900">Data Pasien</h1>
      </div>

      <!-- Main Content Grid -->
      <div class="grid grid-cols-1 gap-6" :class="selectedPasien ? 'lg:grid-cols-3' : 'lg:grid-cols-1'">
        <!-- Left: Table -->
        <div :class="selectedPasien ? 'lg:col-span-2' : 'lg:col-span-3'">
          <!-- Search Bar -->
          <div class="mb-6">
            <div class="relative">
              <svg class="absolute left-3 top-3 w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
              <input
                v-model="searchQuery"
                type="text"
                placeholder="Cari pasien atau nama obat..."
                class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-teal-500"
              />
            </div>
          </div>

          <!-- Desktop Table View -->
          <div class="hidden md:block bg-white rounded-lg shadow overflow-hidden border border-gray-200">
            <table class="w-full">
              <thead class="bg-gray-50 border-b border-gray-200">
                <tr>
                  <th class="px-6 py-3 text-left text-xs font-semibold text-gray-900">NAMA PASIEN</th>
                  <th class="px-6 py-3 text-left text-xs font-semibold text-gray-900">JENIS KELAMIN</th>
                  <th class="px-6 py-3 text-left text-xs font-semibold text-gray-900">KOTA</th>
                  <th class="px-6 py-3 text-left text-xs font-semibold text-gray-900">JUMLAH OBAT</th>
                  <th class="px-6 py-3 text-left text-xs font-semibold text-gray-900">STATUS</th>

                </tr>
              </thead>
              <tbody class="divide-y divide-gray-200">
                <tr 
                  v-for="pasien in filteredPasien" 
                  :key="pasien.pasien_id"
                  @click="selectPasien(pasien)"
                  class="hover:bg-gray-50 transition cursor-pointer"
                  :class="{ 'bg-teal-50': selectedPasien?.pasien_id === pasien.pasien_id }"
                >
                  <td class="px-6 py-4 text-sm font-medium text-slate-900">
                    <div class="flex items-center gap-3">
                      <div class="w-8 h-8 rounded-full bg-teal-100 flex items-center justify-center">
                        <span class="text-xs font-semibold text-teal-700">{{ getInitials(pasien.nama) }}</span>
                      </div>
                      {{ pasien.nama }}
                    </div>
                  </td>
                  <td class="px-6 py-4 text-sm text-gray-600">{{ pasien.jenis_kelamin || '-' }}</td>
                  <td class="px-6 py-4 text-sm text-gray-600">{{ pasien.tempat_lahir || '-' }}</td>
                  <td class="px-6 py-4 text-sm font-medium text-slate-900">
                    <span class="inline-flex items-center justify-center w-6 h-6 bg-blue-100 text-blue-700 rounded-full text-xs font-bold">
                      {{ pasien.jumlah_obat || 0 }}
                    </span>
                  </td>
                  <td class="px-6 py-4 text-sm">
                    <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-teal-100 text-teal-800">
                      Aktif
                    </span>
                  </td>

                </tr>
              </tbody>
            </table>

            <!-- Empty State -->
            <div v-if="filteredPasien.length === 0" class="px-6 py-12 text-center">
              <svg class="w-12 h-12 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.856-1.487M15 10a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
              <p class="text-gray-500 text-sm">Belum ada data pasien</p>
            </div>
          </div>

          <!-- Mobile Card View -->
          <div class="md:hidden space-y-3">
            <div 
              v-for="pasien in filteredPasien" 
              :key="pasien.pasien_id"
              @click="selectPasien(pasien)"
              class="bg-white rounded-lg shadow p-4 border border-gray-200 cursor-pointer hover:shadow-md transition"
              :class="{ 'border-teal-500 bg-teal-50': selectedPasien?.pasien_id === pasien.pasien_id }"
            >
              <div class="flex items-center gap-2 mb-3">
                <div class="w-8 h-8 rounded-full bg-teal-100 flex items-center justify-center flex-shrink-0">
                  <span class="text-xs font-semibold text-teal-700">{{ getInitials(pasien.nama) }}</span>
                </div>
                <div>
                  <h3 class="font-semibold text-slate-900 text-sm">{{ pasien.nama }}</h3>
                  <p class="text-xs text-gray-500">{{ pasien.tempat_lahir }}</p>
                </div>
              </div>
              <div class="grid grid-cols-2 gap-2 text-xs">
                <div>
                  <span class="text-gray-600">Jenis Kelamin</span>
                  <p class="font-medium text-slate-900">{{ pasien.jenis_kelamin || '-' }}</p>
                </div>
                <div>
                  <span class="text-gray-600">Jumlah Obat</span>
                  <p class="font-medium text-slate-900">{{ pasien.jumlah_obat || 0 }}</p>
                </div>
              </div>
            </div>

            <!-- Empty State Mobile -->
            <div v-if="filteredPasien.length === 0" class="text-center py-8 bg-white rounded-lg">
              <svg class="w-12 h-12 text-gray-400 mx-auto mb-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.856-1.487M15 10a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
              <p class="text-gray-500 text-sm">Belum ada data pasien</p>
            </div>
          </div>

          <!-- Pagination info -->
          <div class="mt-4 text-xs text-gray-500">
            Menampilkan {{ filteredPasien.length }} dari {{ pasienStore.pasienList.length }} pasien
          </div>
        </div>

        <!-- Right: Detail Panel -->
        <div v-if="selectedPasien" class="lg:col-span-1">
          <div v-if="selectedPasien" class="bg-white rounded-lg shadow border border-gray-200 overflow-hidden sticky top-4">
            <!-- Header -->
            <div class="bg-gradient-to-r from-teal-500 to-teal-600 px-6 py-4">
              <div class="flex items-center gap-3">
                <div class="w-12 h-12 rounded-full bg-white bg-opacity-20 flex items-center justify-center">
                  <span class="text-lg font-bold text-white">{{ getInitials(selectedPasien.nama) }}</span>
                </div>
                <div>
                  <h3 class="font-bold text-white">{{ selectedPasien.nama }}</h3>
                  <p class="text-teal-100 text-xs">{{ selectedPasien.nik }}</p>
                </div>
              </div>
            </div>

            <!-- Content -->
            <div class="p-6 space-y-4">
              <!-- Basic Info Section -->
              <div>
                <h4 class="text-xs font-semibold text-gray-500 uppercase mb-3">Data Pribadi</h4>
                <div class="space-y-3">
                  <div>
                    <p class="text-xs text-gray-500">EMAIL</p>
                    <p class="text-sm font-medium text-slate-900 break-words">{{ selectedPasien.email }}</p>
                  </div>
                  <div>
                    <p class="text-xs text-gray-500">NO TELEPON</p>
                    <p class="text-sm font-medium text-slate-900">{{ selectedPasien.no_telepon || '-' }}</p>
                  </div>
                  <div>
                    <p class="text-xs text-gray-500">NO TELEPON PENDAMPING (WA)</p>
                    <p class="text-sm font-medium text-slate-900 mt-1">{{ selectedPasien.no_telepon_pendamping || '-' }}</p>
                  </div>
                  <div>
                    <p class="text-xs text-gray-500">JENIS KELAMIN</p>
                    <p class="text-sm font-medium text-slate-900">{{ selectedPasien.jenis_kelamin || '-' }}</p>
                  </div>
                </div>
              </div>

              <!-- Divider -->
              <div class="border-t border-gray-200"></div>

              <!-- Address Info Section -->
              <div>
                <h4 class="text-xs font-semibold text-gray-500 uppercase mb-3">Informasi Alamat</h4>
                <div class="space-y-3">
                  <div>
                    <p class="text-xs text-gray-500">TEMPAT LAHIR</p>
                    <p class="text-sm font-medium text-slate-900">{{ selectedPasien.tempat_lahir || '-' }}</p>
                  </div>
                  <div>
                    <p class="text-xs text-gray-500">TANGGAL LAHIR</p>
                    <p class="text-sm font-medium text-slate-900">{{ formatDate(selectedPasien.tanggal_lahir) }}</p>
                  </div>
                  <div>
                    <p class="text-xs text-gray-500">ALAMAT</p>
                    <p class="text-sm font-medium text-slate-900 break-words">{{ selectedPasien.alamat || '-' }}</p>
                  </div>
                </div>
              </div>

              <!-- Divider -->
              <div class="border-t border-gray-200"></div>

              <!-- Medicine Info Section -->
              <div>
                <h4 class="text-xs font-semibold text-gray-500 uppercase mb-3">Informasi Obat</h4>
                <div class="bg-blue-50 rounded-lg p-4 text-center">
                  <p class="text-xs text-blue-600 mb-1">Total Obat Aktif</p>
                  <p class="text-2xl font-bold text-blue-700">{{ selectedPasien.jumlah_obat || 0 }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </LayoutWrapper>
</template>

<script>
import { onMounted, ref, computed } from 'vue'
import LayoutWrapper from '../components/LayoutWrapper.vue'
import { usePasienStore } from '../stores/pasien'

export default {
  name: 'PasienView',
  components: {
    LayoutWrapper
  },
  setup() {
    const pasienStore = usePasienStore()
    const searchQuery = ref('')
    const selectedPasien = ref(null)

    const filteredPasien = computed(() => {
      const list = pasienStore.pasienList || []
      if (!searchQuery.value) return list
      const q = searchQuery.value.toLowerCase()
      return list.filter(p =>
        p.nama.toLowerCase().includes(q) ||
        p.email.toLowerCase().includes(q) ||
        (p.nik && p.nik.includes(q))
      )
    })

    const getInitials = (nama) => {
      if (!nama) return '?'
      return nama.split(' ').map(n => n[0]).join('').toUpperCase().slice(0, 2)
    }

    const formatDate = (dateStr) => {
      if (!dateStr) return '-'
      return new Date(dateStr).toLocaleDateString('id-ID', {
        year: 'numeric', month: 'long', day: 'numeric'
      })
    }

    const selectPasien = (pasien) => {
      selectedPasien.value = pasien
    }

    const formatWA = (number) => {
      if (!number) return ''
      let clean = number.replace(/\D/g, '')
      if (clean.startsWith('0')) {
        clean = '62' + clean.slice(1)
      }
      return clean
    }

    onMounted(() => {
      pasienStore.fetchPasiens()
    })

    return {
      pasienStore,
      searchQuery,
      selectedPasien,
      filteredPasien,
      selectPasien,
      getInitials,
      formatDate,
      formatWA
    }
  }
}
</script>