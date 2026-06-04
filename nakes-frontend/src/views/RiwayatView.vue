<template>
  <LayoutWrapper>
    <div class="p-4 md:p-8">
      <!-- Header -->
      <div class="mb-6 md:mb-8">
        <h1 class="text-2xl md:text-3xl font-bold text-slate-900 mb-1 md:mb-2">Riwayat Kepatuhan</h1>
      </div>

      <!-- Statistics Cards -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-3 md:gap-4 lg:gap-6 mb-6 md:mb-8">
        <!-- Diminum -->
        <div class="bg-white rounded-lg md:rounded-2xl border-2  p-4 md:p-6 shadow-sm hover:shadow-md transition">
          <div class="flex items-start justify-between mb-3 md:mb-4">
            <div class="flex items-center justify-center">
              <svg class="w-5 h-5" fill="none" stroke="#0F9B76" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.8"
                  d="M5 13l4 4L19 7"/>
              </svg>
            </div>
          </div>
          <p class="text-2xl font-bold text-slate-800 font-bold text-slate-900 mb-1">{{ riwayatStore.statistics?.diminum ?? 0 }}</p>
          <p class="text-xs md:text-sm text-gray-600">Diminum</p>
        </div>

        <!-- Terlambat -->
        <div class="bg-white rounded-lg md:rounded-2xl border-2  p-4 md:p-6 shadow-sm hover:shadow-md transition">
          <div class="flex items-start justify-between mb-3 md:mb-4">
            <div class="flex items-center justify-center">
              <svg class="w-5 h-5" fill="none" stroke="#D97706" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.8"
                  d="M12 8v4l3 3"/>
                <circle cx="12" cy="12" r="9" stroke-width="1.8"/>
              </svg>
            </div>
          </div>
          <p class="text-2xl font-bold text-slate-800 font-bold text-slate-900 mb-1">{{ riwayatStore.statistics?.terlambat ?? 0 }}</p>
          <p class="text-xs md:text-sm text-gray-600">Terlambat</p>
        </div>

        <!-- Terlewat -->
        <div class="bg-white rounded-lg md:rounded-2xl border-2  p-4 md:p-6 shadow-sm hover:shadow-md transition">
          <div class="flex items-start justify-between mb-3 md:mb-4">
            <div class="flex items-center justify-center">
              <svg class="w-5 h-5" fill="none" stroke="#DC2626" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.8"
                  d="M6 6l12 12M6 18L18 6"/>
              </svg>
            </div>
          </div>
          <p class="text-2xl font-bold text-slate-800 font-bold text-slate-900 mb-1">{{ riwayatStore.statistics?.terlewat ?? 0 }}</p>
          <p class="text-xs md:text-sm text-gray-600">Terlewat</p>
        </div>
      </div>

      <!-- Search and Filter -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-4 md:p-6 mb-6 md:mb-8">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-3 md:gap-4">
          <!-- Search Input -->
          <div>
            <label class="block text-xs md:text-sm font-semibold text-gray-700 mb-1 md:mb-2">Cari Pasien atau Obat</label>
            <div class="relative">
              <svg class="absolute left-3 top-2.5 w-4 h-4 md:w-5 md:h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
              <input
                v-model="riwayatStore.searchQuery"
                @input="riwayatStore.setSearchQuery(riwayatStore.searchQuery)"
                type="text"
                placeholder="Ketik nama pasien atau obat..."
                class="w-full pl-9 md:pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-teal-500 focus:border-transparent text-sm"
              />
            </div>
          </div>

          <!-- Status Filter -->
          <div>
            <label class="block text-xs md:text-sm font-semibold text-gray-700 mb-1 md:mb-2">Filter Status</label>
            <select
              v-model="riwayatStore.filterStatus"
              @change="riwayatStore.setFilterStatus(riwayatStore.filterStatus)"
              class="w-full px-3 md:px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-teal-500 focus:border-transparent text-sm"
            >
              <option value="all">Semua Status</option>
              <option value="diminum">Diminum</option>
              <option value="terlambat">Terlambat</option>
              <option value="terlewat">Terlewat</option>
            </select>
          </div>

          <!-- Reset Button -->
          <div class="flex items-end">
            <button
              @click="riwayatStore.resetFilters"
              class="w-full px-3 md:px-4 py-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50 transition font-medium text-sm"
            >
              Reset Filter
            </button>
          </div>
        </div>
      </div>

      <!-- Desktop Table View -->
      <div class="hidden md:block bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
        <div v-if="riwayatStore.loading" class="text-center py-12">
          <svg class="animate-spin h-12 w-12 text-teal-500 mx-auto mb-4" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <p class="text-gray-500 text-sm">Memuat data...</p>
        </div>

        <div v-else-if="riwayatStore.filteredRiwayat.length === 0" class="text-center py-12">
          <svg class="w-12 h-12 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
          </svg>
          <p class="text-gray-500 mb-2 text-sm">Tidak ada data riwayat kepatuhan</p>
          <p class="text-gray-400 text-xs">Cobalah ubah filter atau cari dengan kata kunci lain</p>
        </div>

        <table v-else class="w-full text-sm">
          <thead class="bg-gray-50 border-b border-gray-200 sticky top-0">
            <tr>
              <th class="px-4 py-3 text-left text-xs font-semibold text-gray-700">PASIEN</th>
              <th class="px-4 py-3 text-left text-xs font-semibold text-gray-700">NAMA OBAT</th>
              <th class="px-4 py-3 text-left text-xs font-semibold text-gray-700">Tanggal</th>
              <th class="px-4 py-3 text-left text-xs font-semibold text-gray-700">Jadwal</th>
              <th class="px-4 py-3 text-left text-xs font-semibold text-gray-700">Waktu Minum</th>
              <th class="px-4 py-3 text-left text-xs font-semibold text-gray-700">Status</th>
              <th class="px-4 py-3 text-left text-xs font-semibold text-gray-700">Catatan</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="item in riwayatStore.filteredRiwayat?.filter(i => i)" :key="item?.id || Math.random()" class="hover:bg-gray-50 transition">
              <td class="px-4 py-3 text-xs text-gray-900 font-medium">{{ item?.namaPasien || '-' }}</td>
              <td class="px-4 py-3 text-xs text-gray-600">{{ item?.namaObat || '-' }}</td>
              <td class="px-4 py-3 text-xs text-gray-600">{{ item?.tanggal || '-' }}</td>
              <td class="px-4 py-3 text-xs text-gray-600">{{ item?.jadwal || '-' }}</td>
              <td class="px-4 py-3 text-xs text-gray-600">{{ item?.waktuMinum || '-' }}</td>
              <td class="px-4 py-3">
                <span
                  :class="{
                    'px-2 py-1 rounded text-xs font-semibold inline-block': true,
                    'bg-green-100 text-green-800': item?.status === 'Diminum',
                    'bg-orange-100 text-orange-800': item?.status === 'Terlambat',
                    'bg-red-100 text-red-800': item?.status === 'Terlewat'
                  }"
                >
                  {{ item?.status || '-' }}
                </span>
              </td>
              <td class="px-4 py-3 text-xs text-gray-600">
                <span v-if="item?.catatan" class="inline-flex items-center gap-1 px-2 py-1 bg-blue-50 text-blue-700 rounded text-xs">
                  {{ item.catatan }}
                </span>
                <span v-else class="text-gray-400">-</span>
              </td>
            </tr>
          </tbody>
        </table>

        <!-- Pagination Info -->
        <div class="px-4 md:px-6 py-3 md:py-4 border-t border-gray-200 bg-gray-50 flex items-center justify-between">
          <p class="text-xs md:text-sm text-gray-600">
            Menampilkan <span class="font-semibold">{{ riwayatStore.filteredRiwayat.length }}</span> dari <span class="font-semibold">{{ riwayatStore.riwayatList.length }}</span> data
          </p>
        </div>
      </div>

      <!-- Mobile Card View -->
      <div class="md:hidden space-y-3">
        <div v-if="riwayatStore.loading" class="text-center py-12 bg-white rounded-lg">
          <svg class="animate-spin h-12 w-12 text-teal-500 mx-auto mb-4" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <p class="text-gray-500 text-sm">Memuat data...</p>
        </div>

        <div v-else-if="riwayatStore.filteredRiwayat.length === 0" class="text-center py-8 bg-white rounded-lg">
          <svg class="w-10 h-10 text-gray-400 mx-auto mb-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
          </svg>
          <p class="text-gray-500 text-sm mb-1">Tidak ada data riwayat kepatuhan</p>
          <p class="text-gray-400 text-xs">Cobalah ubah filter atau cari dengan kata kunci lain</p>
        </div>

        <div v-else>
          <div v-for="item in riwayatStore.filteredRiwayat?.filter(i => i)" :key="item?.id || Math.random()" class="bg-white rounded-lg shadow-sm p-4 border border-gray-200">
            <div class="flex justify-between items-start mb-3">
              <div class="flex-1 min-w-0">
                <h3 class="font-semibold text-slate-900 text-sm truncate">{{ item?.namaPasien || '-' }}</h3>
                <p class="text-xs text-gray-500 mt-0.5">{{ item?.namaObat || '-' }}</p>
              </div>
              <span
                :class="{
                  'px-2 py-1 rounded text-xs font-semibold flex-shrink-0 ml-2': true,
                  'bg-green-100 text-green-800': item?.status === 'Diminum',
                  'bg-orange-100 text-orange-800': item?.status === 'Terlambat',
                  'bg-red-100 text-red-800': item?.status === 'Terlewat'
                }"
              >
                {{ item?.status || '-' }}
              </span>
            </div>

            <div class="space-y-2 text-xs mb-3">
              <div class="flex justify-between">
                <span class="text-gray-600">Tanggal:</span>
                <span class="text-slate-900 font-medium">{{ item?.tanggal || '-' }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">Jadwal:</span>
                <span class="text-slate-900 font-medium">{{ item?.jadwal || '-' }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">Waktu Minum:</span>
                <span class="text-slate-900 font-medium">{{ item?.waktuMinum || '-' }}</span>
              </div>
              <div v-if="item?.catatan" class="flex justify-between pt-2 border-t border-gray-200">
                <span class="text-gray-600">Catatan:</span>
                <span class="text-blue-600 font-medium">{{ item.catatan }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </LayoutWrapper>
</template>

<script setup>
import { onMounted } from 'vue'
import LayoutWrapper from '../components/LayoutWrapper.vue'
import { useRiwayatStore } from '../stores/riwayat'

const riwayatStore = useRiwayatStore()

onMounted(() => {
  riwayatStore.fetchRiwayat()
})
</script>

<style scoped>
</style>
