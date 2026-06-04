<template>
  <div class="p-4 md:p-8">
    <div class="max-w-3xl mx-auto">
      <!-- Success Card -->
      <div class="bg-white rounded-2xl shadow-sm border border-gray-100 p-6 md:p-10 text-center mb-6">
        <div class="inline-flex items-center justify-center w-20 h-20 bg-teal-600 rounded-full mb-6">
          <svg class="w-10 h-10 text-white" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"/>
          </svg>
        </div>
        <h2 class="text-xl md:text-2xl font-bold text-slate-900 mb-2">Jadwal Berhasil Disimpan!</h2>
        <p class="text-gray-600 mb-1">
          Jadwal minum obat untuk <span class="font-semibold text-slate-900">{{ selectedPasienName }}</span> telah aktif.
        </p>
        <p class="text-sm text-gray-500">Notifikasi pengingat akan dikirim ke aplikasi pasien sesuai jadwal.</p>

        <!-- Summary pills -->
        <div class="flex justify-center gap-2 md:gap-4 mt-6 md:mt-8 flex-wrap">
          <div class="flex flex-col items-center gap-2 px-6 py-4 bg-gray-50 rounded-xl min-w-28">
            <div class="w-8 h-8 md:w-10 md:h-10 bg-teal-100 rounded-full flex items-center justify-center">
              <svg class="w-4 h-4 md:w-5 md:h-5 text-teal-600" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z" clip-rule="evenodd"/>
              </svg>
            </div>
            <p class="text-xs text-gray-500 uppercase tracking-wide">Pasien</p>
            <p class="text-xs md:text-sm font-bold text-slate-900 text-center line-clamp-2">{{ selectedPasienName }}</p>
          </div>

          <div class="flex flex-col items-center gap-2 px-6 py-4 bg-gray-50 rounded-xl min-w-28">
            <div class="w-10 h-10 bg-purple-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-purple-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M8 16.5a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0zM15 16.5a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0z"/>
                <path d="M3 4a1 1 0 00-1 1v10a1 1 0 001 1h1.05a2.5 2.5 0 014.9 0H10a1 1 0 001-1V5a1 1 0 00-1-1H3zM14 7a1 1 0 00-1 1v6.05A2.5 2.5 0 0115.95 16H17a1 1 0 001-1v-5a1 1 0 00-.293-.707l-2-2A1 1 0 0015 7h-1z"/>
              </svg>
            </div>
            <p class="text-xs text-gray-500 uppercase tracking-wide">Obat</p>
            <p class="text-sm font-bold text-slate-900 text-center">{{ form.nama_obat }}</p>
          </div>

          <div class="flex flex-col items-center gap-2 px-6 py-4 bg-gray-50 rounded-xl min-w-28">
            <div class="w-10 h-10 bg-yellow-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-yellow-600" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z" clip-rule="evenodd"/>
              </svg>
            </div>
            <p class="text-xs text-gray-500 uppercase tracking-wide">Jadwal</p>
            <p class="text-sm font-bold text-slate-900 text-center">
              {{ form.frekuensi_per_hari }} · {{ form.tipe_durasi === 'rutin' ? 'Rutin' : form.jumlah_hari + ' hari' }}
            </p>
          </div>

          <div class="flex flex-col items-center gap-2 px-6 py-4 bg-gray-50 rounded-xl min-w-28">
            <div class="w-10 h-10 bg-orange-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-orange-500" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v2a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z" clip-rule="evenodd"/>
              </svg>
            </div>
            <p class="text-xs text-gray-500 uppercase tracking-wide">Mulai</p>
            <p class="text-sm font-bold text-slate-900 text-center">{{ formatDate(form.tanggal_mulai) }}</p>
          </div>
        </div>

        <div class="flex flex-col-reverse md:flex-row gap-3 mt-6 md:mt-8 justify-center">
          <button @click="$emit('back-to-list')"
            class="w-full md:w-auto px-4 md:px-6 py-2.5 border border-gray-300 rounded-lg text-xs md:text-sm font-medium text-gray-700 hover:bg-gray-50 transition flex items-center justify-center gap-2">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
            </svg>
            Kembali ke Daftar
          </button>
          <button @click="$emit('add-another')"
            class="w-full md:w-auto px-4 md:px-6 py-2.5 bg-teal-600 text-white rounded-lg text-xs md:text-sm font-medium hover:bg-teal-700 transition flex items-center justify-center gap-2">
            + Tambah Jadwal Lain
          </button>
        </div>
      </div>

      <!-- Bottom info panels -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4 md:gap-6">
        <!-- Perhatian -->
        <div class="bg-white rounded-xl shadow-sm border border-gray-100 p-4 md:p-5">
          <h4 class="text-xs md:text-sm font-semibold text-slate-900 mb-3 md:mb-4 flex items-center gap-2">
            <svg class="w-3 h-3 md:w-4 md:h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
            </svg>
            Perhatian
          </h4>
          <ol class="space-y-2 md:space-y-3">
            <li class="flex gap-3">
              <span class="flex-shrink-0 w-5 h-5 md:w-6 md:h-6 rounded-full bg-gray-100 flex items-center justify-center text-xs font-bold text-gray-600">1</span>
              <div>
                <p class="text-xs md:text-sm font-medium text-slate-900">Notifikasi dikirim ke pasien</p>
                <p class="text-xs text-gray-500 mt-0.5">
                  {{ selectedPasienName }} akan menerima pengingat di aplikasi pada pukul {{ form.waktu_reminder_pagi }} dan {{ form.waktu_reminder_malam }}
                </p>
              </div>
            </li>
            <li class="flex gap-3">
              <span class="flex-shrink-0 w-5 h-5 md:w-6 md:h-6 rounded-full bg-gray-100 flex items-center justify-center text-xs font-bold text-gray-600">2</span>
              <div>
                <p class="text-xs md:text-sm font-medium text-slate-900">Jadwal mulai aktif hari ini</p>
                <p class="text-xs text-gray-500 mt-0.5">
                  Terhitung sejak {{ formatDate(form.tanggal_mulai) }} dan akan berjalan secara {{ form.tipe_durasi }}
                </p>
              </div>
            </li>
            <li class="flex gap-3">
              <span class="flex-shrink-0 w-5 h-5 md:w-6 md:h-6 rounded-full bg-gray-100 flex items-center justify-center text-xs font-bold text-gray-600">3</span>
              <div>
                <p class="text-xs md:text-sm font-medium text-slate-900">Pantau kepatuhan pasien</p>
                <p class="text-xs text-gray-500 mt-0.5">Cek riwayat kepatuhan minum obat di menu Riwayat Kepatuhan</p>
              </div>
            </li>
          </ol>
        </div>

        <!-- Jadwal Aktif pasien -->
        <div class="bg-white rounded-xl shadow-sm border border-gray-100 p-4 md:p-5">
          <h4 class="text-xs md:text-sm font-semibold text-slate-900 mb-3 md:mb-4 flex items-center gap-2">
            <svg class="w-3 h-3 md:w-4 md:h-4 text-gray-400" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v2a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z" clip-rule="evenodd"/>
            </svg>
            Jadwal Aktif — {{ selectedPasienName }}
          </h4>
          <div class="space-y-2">
            <!-- Obat baru (highlighted) -->
            <div class="flex items-center justify-between p-2 md:p-3 rounded-lg border border-teal-200 bg-teal-50">
              <div class="flex-1 min-w-0">
                <div class="flex items-center gap-2 flex-wrap">
                  <p class="text-xs md:text-sm font-semibold text-slate-900 truncate">{{ form.nama_obat }}</p>
                  <span class="text-xs bg-teal-600 text-white px-2 py-0.5 rounded-full font-medium flex-shrink-0">Baru</span>
                </div>
                <p class="text-xs text-gray-500 mt-0.5">
                  {{ form.frekuensi_per_hari }} · {{ selectedWaktuMinum.join(' & ') }} · {{ form.tipe_durasi === 'rutin' ? 'Rutin' : form.jumlah_hari + ' hari' }}
                </p>
              </div>
              <span class="text-xs bg-green-100 text-green-700 px-2 py-1 rounded-full font-medium flex-shrink-0">Aktif</span>
            </div>
            <!-- Jadwal pasien lainnya -->
            <div v-for="jadwal in pasienJadwalList" :key="jadwal.id"
              class="flex items-center justify-between p-2 md:p-3 rounded-lg border border-gray-100 hover:bg-gray-50">
              <div class="flex-1 min-w-0">
                <p class="text-xs md:text-sm font-medium text-slate-900 truncate">{{ jadwal.nama_obat }}</p>
                <p class="text-xs text-gray-500 mt-0.5">
                  {{ jadwal.frekuensi_per_hari }}x sehari · {{ jadwal.waktu_minum }} · {{ jadwal.tipe_durasi === 'rutin' ? 'Rutin' : jadwal.jumlah_hari + ' hari' }}
                </p>
              </div>
              <span class="text-xs bg-green-100 text-green-700 px-2 py-1 rounded-full font-medium">Aktif</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Step3Konfirmasi',
  props: {
    form: { type: Object, required: true },
    selectedPasienName: { type: String, default: '' },
    selectedWaktuMinum: { type: Array, default: () => [] },
    pasienJadwalList: { type: Array, default: () => [] },
  },
  emits: ['back-to-list', 'add-another'],
  methods: {
    formatDate(dateStr) {
      if (!dateStr) return '-'
      return new Date(dateStr).toLocaleDateString('id-ID', {
        year: 'numeric', month: 'long', day: 'numeric'
      })
    },
  },
}
</script>