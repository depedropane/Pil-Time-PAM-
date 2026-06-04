<template>
  <div class="p-4 md:p-8">
    <!-- Header -->
    <div class="flex flex-col md:flex-row md:justify-between md:items-center gap-4 mb-6 md:mb-8">
      <div>
        <h1 class="text-xl md:text-2xl font-bold text-slate-900">Jadwal Obat</h1>
        <p class="text-xs md:text-sm text-gray-500 mt-1">Atur frekuensi dan durasi konsumsi obat pasien di sini</p>
      </div>
      <button
        @click="$emit('open-add')"
        class="w-full md:w-auto px-4 md:px-5 py-2.5 bg-teal-600 text-white rounded-lg hover:bg-teal-700 transition font-medium text-xs md:text-sm flex items-center justify-center gap-2 shadow-sm"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
        </svg>
        Tambah Jadwal Obat
      </button>
    </div>

    <!-- Search -->
    <div class="mb-4 md:mb-5">
      <div class="relative w-full md:max-w-sm">
        <svg class="absolute left-3 top-2.5 w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
        </svg>
        <input
          :value="searchQuery"
          @input="$emit('update:searchQuery', $event.target.value)"
          type="text"
          placeholder="Cari pasien atau nama obat..."
          class="w-full pl-9 pr-4 py-2 text-sm border border-gray-200 rounded-lg focus:ring-2 focus:ring-teal-500 focus:border-transparent outline-none bg-white"
        />
      </div>
    </div>

    <!-- Table: hilangkan overflow-hidden supaya dropdown tidak terpotong -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-100">
      <div class="hidden md:block overflow-x-auto">
        <table class="w-full">
          <thead>
            <tr class="border-b border-gray-100">
              <th class="px-6 py-3.5 text-left text-xs font-semibold text-gray-500 uppercase tracking-wide">Pasien</th>
              <th class="px-6 py-3.5 text-left text-xs font-semibold text-gray-500 uppercase tracking-wide">Nama Obat</th>
              <th class="px-6 py-3.5 text-left text-xs font-semibold text-gray-500 uppercase tracking-wide">Dosis</th>
              <th class="px-6 py-3.5 text-left text-xs font-semibold text-gray-500 uppercase tracking-wide">Frekuensi</th>
              <th class="px-6 py-3.5 text-left text-xs font-semibold text-gray-500 uppercase tracking-wide">Waktu</th>
              <th class="px-6 py-3.5 text-left text-xs font-semibold text-gray-500 uppercase tracking-wide">Durasi</th>
              <th class="px-6 py-3.5 text-left text-xs font-semibold text-gray-500 uppercase tracking-wide">Status</th>
              <th class="px-6 py-3.5 text-left text-xs font-semibold text-gray-500 uppercase tracking-wide">Aksi</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-50">
            <tr v-for="jadwal in jadwalList" :key="jadwal.id" class="hover:bg-gray-50 transition">
              <td class="px-6 py-4 text-sm font-medium text-slate-900">{{ jadwal.pasien_nama }}</td>
              <td class="px-6 py-4 text-sm text-gray-600">{{ jadwal.nama_obat }}</td>
              <td class="px-6 py-4 text-sm text-gray-600">{{ jadwal.jumlah_dosis }} {{ jadwal.satuan }}</td>
              <td class="px-6 py-4 text-sm text-gray-600">{{ jadwal.frekuensi_per_hari }}x sehari</td>
              <td class="px-6 py-4 text-sm text-gray-600">{{ jadwal.waktu_minum }}</td>
              <td class="px-6 py-4 text-sm">
                <span v-if="jadwal.tipe_durasi === 'hari'"
                  class="text-xs bg-teal-50 text-teal-700 border border-teal-200 px-2.5 py-1 rounded-full font-medium">
                  {{ jadwal.jumlah_hari }} hari
                </span>
                <span v-else class="text-xs bg-gray-100 text-gray-600 border border-gray-200 px-2.5 py-1 rounded-full font-medium">
                  Rutin
                </span>
              </td>
              <td class="px-6 py-4">
                <span v-if="jadwal.status === 'Aktif'" class="text-xs bg-green-50 text-green-700 px-2.5 py-1 rounded-full font-medium flex items-center gap-1 w-fit">
                  <span class="w-1.5 h-1.5 bg-green-500 rounded-full"></span>
                  Aktif
                </span>
                <span v-else class="text-xs bg-red-50 text-red-700 px-2.5 py-1 rounded-full font-medium flex items-center gap-1 w-fit">
                  <span class="w-1.5 h-1.5 bg-red-500 rounded-full"></span>
                  Tidak Aktif
                </span>
              </td>
              <td class="px-6 py-4">
                <div class="relative inline-block" :ref="el => { if(el) menuRefs[jadwal.id] = el }">
                  <button
                    @click.stop="toggleMenu(jadwal.id)"
                    class="p-1.5 text-gray-400 hover:text-gray-600 hover:bg-gray-100 rounded transition"
                  >
                    <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                      <path d="M6 10a2 2 0 11-4 0 2 2 0 014 0zM12 10a2 2 0 11-4 0 2 2 0 014 0zM16 12a2 2 0 100-4 2 2 0 000 4z"/>
                    </svg>
                  </button>

                  <!-- Teleport ke body agar tidak terpotong overflow -->
                  <Teleport to="body">
                    <div
                      v-if="openMenuId === jadwal.id"
                      :style="getDropdownStyle(jadwal.id)"
                      class="fixed w-44 bg-white rounded-lg shadow-xl border border-gray-200 z-50"
                      @click.stop
                    >
                      <button @click="handleAction('view', jadwal)"
                        class="w-full text-left px-4 py-2.5 flex items-center gap-3 text-sm text-gray-700 hover:bg-teal-50 hover:text-teal-700 transition rounded-t-lg">
                        <svg class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/>
                        </svg>
                        Lihat Detail
                      </button>
                      <div class="border-t border-gray-100"></div>
                      <button @click="handleAction('edit', jadwal)"
                        class="w-full text-left px-4 py-2.5 flex items-center gap-3 text-sm text-gray-700 hover:bg-blue-50 hover:text-blue-700 transition">
                        <svg class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/>
                        </svg>
                        Edit
                      </button>
                      <div class="border-t border-gray-100"></div>
                      <button @click="handleAction('delete', jadwal)"
                        class="w-full text-left px-4 py-2.5 flex items-center gap-3 text-sm text-gray-700 hover:bg-red-50 hover:text-red-600 transition rounded-b-lg">
                        <svg class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/>
                        </svg>
                        Hapus
                      </button>
                    </div>
                  </Teleport>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Mobile Card View -->
      <div class="md:hidden divide-y divide-gray-100">
        <div v-for="jadwal in jadwalList" :key="'m_' + jadwal.id" class="p-4">
          <div class="flex justify-between items-start mb-3">
            <div class="flex-1 min-w-0">
              <p class="font-semibold text-slate-900 text-sm">{{ jadwal.pasien_nama }}</p>
              <p class="text-xs text-gray-500 mt-0.5">{{ jadwal.nama_obat }}</p>
            </div>
            <div class="ml-2 flex-shrink-0 relative inline-block" :ref="el => { if(el) menuRefs['m_' + jadwal.id] = el }">
              <button
                @click.stop="toggleMenu('m_' + jadwal.id)"
                class="p-1.5 text-gray-400 hover:text-gray-600 hover:bg-gray-100 rounded transition"
              >
                <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M6 10a2 2 0 11-4 0 2 2 0 014 0zM12 10a2 2 0 11-4 0 2 2 0 014 0zM16 12a2 2 0 100-4 2 2 0 000 4z"/>
                </svg>
              </button>
              <Teleport to="body">
                <div
                  v-if="openMenuId === 'm_' + jadwal.id"
                  :style="getDropdownStyle('m_' + jadwal.id)"
                  class="fixed w-44 bg-white rounded-lg shadow-xl border border-gray-200 z-50"
                  @click.stop
                >
                  <button @click="handleAction('view', jadwal)"
                    class="w-full text-left px-4 py-2.5 flex items-center gap-3 text-sm text-gray-700 hover:bg-teal-50 hover:text-teal-700 transition rounded-t-lg">
                    <svg class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/>
                    </svg>
                    Lihat Detail
                  </button>
                  <div class="border-t border-gray-100"></div>
                  <button @click="handleAction('edit', jadwal)"
                    class="w-full text-left px-4 py-2.5 flex items-center gap-3 text-sm text-gray-700 hover:bg-blue-50 hover:text-blue-700 transition">
                    <svg class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/>
                    </svg>
                    Edit
                  </button>
                  <div class="border-t border-gray-100"></div>
                  <button @click="handleAction('delete', jadwal)"
                    class="w-full text-left px-4 py-2.5 flex items-center gap-3 text-sm text-gray-700 hover:bg-red-50 hover:text-red-600 transition rounded-b-lg">
                    <svg class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/>
                    </svg>
                    Hapus
                  </button>
                </div>
              </Teleport>
            </div>
          </div>
          <div class="grid grid-cols-2 gap-2 text-xs">
            <div class="bg-gray-50 rounded-lg p-2">
              <p class="text-gray-500 font-medium">Dosis</p>
              <p class="font-semibold text-slate-900 mt-1">{{ jadwal.jumlah_dosis }} {{ jadwal.satuan }}</p>
            </div>
            <div class="bg-gray-50 rounded-lg p-2">
              <p class="text-gray-500 font-medium">Frekuensi</p>
              <p class="font-semibold text-slate-900 mt-1">{{ jadwal.frekuensi_per_hari }}x/hari</p>
            </div>
            <div class="bg-gray-50 rounded-lg p-2">
              <p class="text-gray-500 font-medium">Waktu</p>
              <p class="font-semibold text-slate-900 mt-1">{{ jadwal.waktu_minum }}</p>
            </div>
            <div class="bg-gray-50 rounded-lg p-2">
              <p class="text-gray-500 font-medium">Durasi</p>
              <p class="font-semibold text-slate-900 mt-1">
                <span v-if="jadwal.tipe_durasi === 'hari'">{{ jadwal.jumlah_hari }} hari</span>
                <span v-else>Rutin</span>
              </p>
            </div>
          </div>
          <div class="mt-3 flex justify-end">
            <span v-if="jadwal.status === 'Aktif'" class="text-xs bg-green-50 text-green-700 px-2.5 py-1 rounded-full font-medium flex items-center gap-1">
              <span class="w-1.5 h-1.5 bg-green-500 rounded-full"></span>
              Aktif
            </span>
            <span v-else class="text-xs bg-red-50 text-red-700 px-2.5 py-1 rounded-full font-medium flex items-center gap-1">
              <span class="w-1.5 h-1.5 bg-red-500 rounded-full"></span>
              Tidak Aktif
            </span>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-if="jadwalList.length === 0" class="py-16 text-center">
        <svg class="w-12 h-12 text-gray-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"
            d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"/>
        </svg>
        <p class="text-gray-400 text-sm">Belum ada data jadwal obat</p>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted, onBeforeUnmount } from 'vue'

export default {
  name: 'JadwalTable',
  props: {
    jadwalList: { type: Array, default: () => [] },
    searchQuery: { type: String, default: '' },
  },
  emits: ['open-add', 'view', 'edit', 'delete', 'update:searchQuery'],
  setup(props, { emit }) {
    const openMenuId = ref(null)
    const menuRefs = ref({})

    const toggleMenu = (id) => {
      openMenuId.value = openMenuId.value === id ? null : id
    }

    // Hitung posisi dropdown berdasarkan posisi tombol di layar (pakai getBoundingClientRect)
    const getDropdownStyle = (id) => {
      const el = menuRefs.value[id]
      if (!el) return {}
      const rect = el.getBoundingClientRect()
      return {
        top: rect.bottom + 4 + 'px',
        left: rect.right - 176 + 'px', // 176px = w-44
      }
    }

    const handleAction = (action, jadwal) => {
      openMenuId.value = null
      if (action === 'view') emit('view', jadwal)
      else if (action === 'edit') emit('edit', jadwal)
      else if (action === 'delete') emit('delete', jadwal.id)
    }

    const handleClickOutside = (e) => {
      const isInsideAnyMenu = Object.values(menuRefs.value).some(
        el => el && el.contains(e.target)
      )
      if (!isInsideAnyMenu) openMenuId.value = null
    }

    onMounted(() => document.addEventListener('click', handleClickOutside))
    onBeforeUnmount(() => document.removeEventListener('click', handleClickOutside))

    return { openMenuId, menuRefs, toggleMenu, handleAction, getDropdownStyle }
  }
}
</script>