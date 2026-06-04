<template>
  <div class="p-8">
    <!-- Info bar -->
    <div class="flex items-center gap-6 mb-6 bg-white rounded-xl px-6 py-3 shadow-sm border border-gray-100">
      <div class="flex items-center gap-3">
        <div :class="['w-8 h-8 rounded-full flex items-center justify-center text-white text-xs font-bold', getColorClass(form.patientId)]">
          {{ getInitials(selectedPasienName) }}
        </div>
        <div>
          <p class="text-xs text-gray-400 uppercase tracking-wide">Pasien</p>
          <p class="text-sm font-semibold text-slate-900">{{ selectedPasienName }}</p>
        </div>
        <span class="text-xs bg-teal-50 text-teal-700 px-2 py-0.5 rounded-full font-medium">{{ selectedPasienCode }}</span>
      </div>
      <div class="flex items-center gap-3 pl-6 border-l border-gray-200">
        <div class="w-8 h-8 rounded-full bg-purple-100 flex items-center justify-center">
          <svg class="w-4 h-4 text-purple-600" fill="currentColor" viewBox="0 0 20 20">
            <path d="M8 16.5a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0zM15 16.5a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0z"/>
            <path d="M3 4a1 1 0 00-1 1v10a1 1 0 001 1h1.05a2.5 2.5 0 014.9 0H10a1 1 0 001-1V5a1 1 0 00-1-1H3zM14 7a1 1 0 00-1 1v6.05A2.5 2.5 0 0115.95 16H17a1 1 0 001-1v-5a1 1 0 00-.293-.707l-2-2A1 1 0 0015 7h-1z"/>
          </svg>
        </div>
        <div>
          <p class="text-xs text-gray-400 uppercase tracking-wide">Obat</p>
          <p class="text-sm font-semibold text-slate-900">{{ form.nama_obat }}</p>
        </div>
      </div>
      <div class="flex items-center gap-3 pl-6 border-l border-gray-200">
        <div class="w-8 h-8 rounded-full bg-blue-100 flex items-center justify-center">
          <svg class="w-4 h-4 text-blue-600" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z" clip-rule="evenodd"/>
          </svg>
        </div>
        <div>
          <p class="text-xs text-gray-400 uppercase tracking-wide">Dosis</p>
          <p class="text-sm font-semibold text-slate-900">{{ form.jumlah_dosis }} {{ form.satuan }} · {{ form.kategori_obat }}</p>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-4 md:gap-6">
      <!-- Left: Frekuensi & Waktu Minum -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-100 p-4 md:p-6">
        <div class="flex items-center gap-3 mb-5">
          <div class="w-8 h-8 rounded-full bg-yellow-50 flex items-center justify-center">
            <svg class="w-4 h-4 text-yellow-600" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z" clip-rule="evenodd"/>
            </svg>
          </div>
          <div>
            <h3 class="text-sm font-semibold text-slate-900">Frekuensi & Waktu Minum</h3>
            <p class="text-xs text-gray-500">Tentukan seberapa sering dan kapan obat diminum</p>
          </div>
        </div>

        <div class="space-y-3 md:space-y-5">
          <div>
            <label class="block text-xs md:text-sm font-medium text-gray-700 mb-1">
              Frekuensi per Hari <span class="text-red-500">*</span>
            </label>
            <input :value="form.frekuensi_per_hari" @input="update('frekuensi_per_hari', $event.target.value)"
              type="text" placeholder="Contoh: 2x sehari"
              class="w-full px-3 py-2 text-sm border border-gray-200 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-transparent outline-none"/>
          </div>

          <div>
            <label class="block text-xs md:text-sm font-medium text-gray-700 mb-2">
              Waktu Minum <span class="text-red-500">*</span>
            </label>
            <div class="flex flex-wrap gap-2 md:gap-3">
              <button
                v-for="waktu in waktuMinum" :key="waktu.value"
                type="button"
                @click="$emit('toggle-waktu', waktu.value)"
                :class="[
                  'flex items-center gap-2 px-3 md:px-4 py-2 rounded-lg text-xs md:text-sm font-medium border transition flex-1',
                  selectedWaktuMinum.includes(waktu.value)
                    ? 'bg-teal-600 text-white border-teal-600'
                    : 'bg-white text-gray-600 border-gray-200 hover:border-teal-300'
                ]"
              >
                <span>{{ waktu.icon }}</span>
                {{ waktu.label }}
                <svg v-if="selectedWaktuMinum.includes(waktu.value)" class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"/>
                </svg>
              </button>
            </div>
          </div>

          <div>
            <label class="block text-xs md:text-sm font-medium text-gray-700 mb-2">
              Aturan Konsumsi <span class="text-red-500">*</span>
            </label>
            <div class="flex flex-wrap gap-2 md:gap-3">
              <button
                v-for="aturan in aturanKonsumsi" :key="aturan"
                type="button"
                @click="update('aturan_konsumsi', aturan)"
                :class="[
                  'flex items-center gap-2 px-3 md:px-4 py-2 rounded-lg text-xs md:text-sm font-medium border transition','flex-1',
                  form.aturan_konsumsi === aturan
                    ? 'bg-teal-600 text-white border-teal-600'
                    : 'bg-white text-gray-600 border-gray-200 hover:border-teal-300'
                ]"
              >
                <div :class="['w-4 h-4 rounded-full border-2 flex items-center justify-center flex-shrink-0',
                  form.aturan_konsumsi === aturan ? 'border-white' : 'border-gray-400']">
                  <div v-if="form.aturan_konsumsi === aturan" class="w-2 h-2 bg-white rounded-full"></div>
                </div>
                {{ aturan }}
              </button>
            </div>
          </div>

          <div>
            <label class="block text-xs md:text-sm font-medium text-gray-700 mb-1">Catatan untuk Pasien</label>
            <textarea :value="form.catatan" @input="update('catatan', $event.target.value)"
              rows="4"
              placeholder="Minum dengan segelas air penuh. Hindari konsumsi bersama jus jeruk."
              class="w-full px-3 py-2 text-sm border border-gray-200 rounded-lg focus:ring-2 focus:ring-teal-500 focus:border-transparent outline-none resize-none">
            </textarea>
          </div>
        </div>
      </div>

      <!-- Right: Durasi & Jadwal -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-100 p-4 md:p-6">
        <div class="flex items-center gap-3 mb-5">
          <div class="w-8 h-8 rounded-full bg-red-50 flex items-center justify-center">
            <svg class="w-4 h-4 text-red-500" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v2a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z" clip-rule="evenodd"/>
            </svg>
          </div>
          <div>
            <h3 class="text-sm font-semibold text-slate-900">Durasi & Jadwal</h3>
            <p class="text-xs text-gray-500">Tentukan lama pengobatan dan tanggal mulai</p>
          </div>
        </div>

        <div class="space-y-5">
          <!-- Tipe Durasi -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Tipe Durasi <span class="text-red-500">*</span></label>
            <div class="flex rounded-lg border border-gray-200 overflow-hidden">
              <button type="button" @click="update('tipe_durasi', 'hari')"
                :class="['flex-1 py-2.5 text-xs md:text-sm font-medium transition flex items-center justify-center gap-2','flex','flex-col','md:flex-row',
                  form.tipe_durasi === 'hari' ? 'bg-teal-600 text-white' : 'bg-white text-gray-600 hover:bg-gray-50']">
                📅 Jumlah Hari
              </button>
              <button type="button" @click="update('tipe_durasi', 'rutin')"
                :class="['flex-1 py-2.5 text-xs md:text-sm font-medium transition flex items-center justify-center gap-2','flex','flex-col','md:flex-row',
                  form.tipe_durasi === 'rutin' ? 'bg-teal-600 text-white' : 'bg-white text-gray-600 hover:bg-gray-50']">
                🔄 Rutin
              </button>
            </div>
          </div>

          <!-- Jumlah Hari -->
          <div v-if="form.tipe_durasi === 'hari'" class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">Jumlah Hari <span class="text-red-500">*</span></label>
              <div class="flex items-center gap-3 mb-3">
                <button type="button" @click="update('jumlah_hari', Math.max(1, form.jumlah_hari - 1))"
                  class="w-9 h-9 rounded-lg bg-gray-100 hover:bg-gray-200 flex items-center justify-center text-gray-700 font-bold transition">−</button>
                <span class="text-3xl font-bold text-slate-900 w-16 text-center">{{ form.jumlah_hari }}</span>
                <button type="button" @click="update('jumlah_hari', form.jumlah_hari + 1)"
                  class="w-9 h-9 rounded-lg bg-gray-100 hover:bg-gray-200 flex items-center justify-center text-gray-700 font-bold transition">+</button>
                <span class="text-sm text-gray-500">hari</span>
              </div>
              <div class="flex gap-2 flex-wrap">
                <button v-for="preset in [3, 7, 14, 30]" :key="preset" type="button"
                  @click="update('jumlah_hari', preset)"
                  :class="['px-2 md:px-3 py-1 rounded-full text-xs font-medium border transition','flex-1',
                    form.jumlah_hari === preset
                      ? 'bg-teal-600 text-white border-teal-600'
                      : 'bg-white text-gray-600 border-gray-200 hover:border-teal-300']">
                  {{ preset === 7 ? '1 minggu' : preset === 30 ? '1 bulan' : preset + ' hari' }}
                </button>
              </div>
            </div>

            <div class="grid grid-cols-2 gap-2 md:gap-3">
              <div>
                <label class="block text-xs md:text-sm font-medium text-gray-700 mb-1">Tanggal Mulai <span class="text-red-500">*</span></label>
                <input :value="form.tanggal_mulai" @input="update('tanggal_mulai', $event.target.value)"
                  type="date"
                  class="w-full px-3 py-2 text-sm border border-gray-200 rounded-lg focus:ring-2 focus:ring-teal-500 focus:border-transparent outline-none"/>
              </div>
              <div>
                <label class="block text-xs md:text-sm font-medium text-gray-700 mb-1">Tanggal Selesai</label>
                <input :value="tanggalSelesai" type="date" disabled
                  class="w-full px-3 py-2 text-sm border border-gray-200 rounded-lg bg-gray-50 text-gray-400 outline-none"/>
              </div>
            </div>
          </div>

          <!-- Rutin -->
          <div v-if="form.tipe_durasi === 'rutin'" class="space-y-4">
            <div class="p-3 bg-blue-50 rounded-lg border border-blue-200">
              <p class="text-xs text-blue-700">
                <strong>ℹ️</strong> Obat <strong>Rutin</strong> akan terus aktif dan tidak memiliki tanggal berakhir. Pasien perlu meminum obat ini setiap hari secara berkelanjutan.
              </p>
            </div>
            <div>
              <label class="block text-xs md:text-sm font-medium text-gray-700 mb-1">Tanggal Mulai <span class="text-red-500">*</span></label>
              <input :value="form.tanggal_mulai" @input="update('tanggal_mulai', $event.target.value)"
                type="date"
                class="w-full px-3 py-2 text-sm border border-gray-200 rounded-lg focus:ring-2 focus:ring-teal-500 focus:border-transparent outline-none"/>
            </div>
          </div>

          <!-- Pengingat Otomatis -->
          <div>
            <div class="flex items-center justify-between mb-1">
              <div>
                <p class="text-xs md:text-sm font-medium text-gray-700">Pengingat Otomatis</p>
                <p class="text-xs text-gray-500">Kirim notifikasi ke pasien</p>
              </div>
              <button type="button" @click="update('kirim_notifikasi', !form.kirim_notifikasi)"
                :class="['relative inline-flex h-6 w-11 items-center rounded-full transition',
                  form.kirim_notifikasi ? 'bg-teal-600' : 'bg-gray-200']">
                <span :class="['inline-block h-4 w-4 transform rounded-full bg-white transition shadow',
                  form.kirim_notifikasi ? 'translate-x-6' : 'translate-x-1']"></span>
              </button>
            </div>
            <p class="text-xs text-gray-500 mb-3">Pasien akan mendapat pengingat di aplikasi sesuai waktu minum</p>
          </div>

          <!-- Waktu Pengingat -->
          <div>
            <label class="block text-xs md:text-sm font-medium text-gray-700 mb-2">Waktu Pengingat Spesifik</label>
            <div class="grid grid-cols-2 gap-2 md:gap-3">
              <div>
                <label class="block text-xs text-gray-500 mb-1">☀️ Pagi</label>
                <input :value="form.waktu_reminder_pagi" @input="update('waktu_reminder_pagi', $event.target.value)"
                  type="time"
                  class="w-full px-3 py-2 text-sm border border-gray-200 rounded-lg focus:ring-2 focus:ring-teal-500 focus:border-transparent outline-none"/>
              </div>
              <div>
                <label class="block text-xs text-gray-500 mb-1">🌙 Malam</label>
                <input :value="form.waktu_reminder_malam" @input="update('waktu_reminder_malam', $event.target.value)"
                  type="time"
                  class="w-full px-3 py-2 text-sm border border-gray-200 rounded-lg focus:ring-2 focus:ring-teal-500 focus:border-transparent outline-none"/>
              </div>
            </div>
            <div class="mt-3 p-3 bg-teal-50 rounded-lg border border-teal-200">
              <p class="text-xs text-teal-700">
                🔔 Pengingat akan dikirim <strong>5 menit sebelum</strong> waktu minum yang ditentukan.
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Footer -->
    <div class="flex flex-col-reverse md:flex-row md:justify-end gap-3 mt-4 md:mt-6">
      <button @click="$emit('back')"
        class="w-full md:w-auto px-4 md:px-6 py-2.5 border border-gray-300 rounded-lg text-xs md:text-sm font-medium text-gray-700 hover:bg-gray-50 transition flex items-center justify-center gap-2">
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
        </svg>
        Kembali
      </button>
      <button @click="$emit('next')"
        class="w-full md:w-auto px-4 md:px-6 py-2.5 bg-teal-600 text-white rounded-lg text-xs md:text-sm font-medium hover:bg-teal-700 transition flex items-center justify-center gap-2">
        {{ form.tipe_durasi === 'rutin' ? 'Lihat Konfirmasi' : 'Konfirmasi' }}
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
        </svg>
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Step2AturanMinum',
  props: {
    form: { type: Object, required: true },
    selectedPasienName: { type: String, default: '' },
    selectedPasienCode: { type: String, default: '' },
    selectedWaktuMinum: { type: Array, default: () => [] },
    waktuMinum: { type: Array, default: () => [] },
    aturanKonsumsi: { type: Array, default: () => [] },
  },
  emits: ['update:form', 'toggle-waktu', 'back', 'next'],
  computed: {
    tanggalSelesai() {
      if (!this.form.tanggal_mulai || !this.form.jumlah_hari) return ''
      const d = new Date(this.form.tanggal_mulai)
      d.setDate(d.getDate() + this.form.jumlah_hari)
      return d.toISOString().split('T')[0]
    },
  },
  watch: {
    tanggalSelesai(newVal) {
      if (this.form.tipe_durasi === 'hari' && newVal) {
        this.update('tanggal_selesai', newVal)
      }
    },
  },
  methods: {
    update(field, value) {
      this.$emit('update:form', { ...this.form, [field]: value })
    },
    getInitials(nama) {
      if (!nama) return '?'
      return nama.split(' ').map(n => n[0]).join('').toUpperCase().slice(0, 2)
    },
    getColorClass(id) {
      const colors = ['bg-teal-500', 'bg-purple-500', 'bg-yellow-500', 'bg-red-500', 'bg-blue-500', 'bg-orange-500']
      return colors[(id || 0) % colors.length]
    },
  },
}
</script>