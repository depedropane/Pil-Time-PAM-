<template>
  <div class="p-4 md:p-8">
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      
      <div class="bg-white rounded-2xl shadow-sm border border-gray-100 p-5 md:p-6 transition-all">
        <div class="flex items-center gap-3 mb-5">
          <div class="w-10 h-10 rounded-xl bg-teal-50 flex items-center justify-center text-teal-600">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/></svg>
          </div>
          <h3 class="text-sm font-bold text-slate-900">Data Pasien</h3>
        </div>

        <div class="relative mb-4">
          <input :value="searchPasien" @input="$emit('update:searchPasien', $event.target.value)"
            type="text" placeholder="Cari nama atau NIK..."
            class="w-full pl-10 pr-4 py-2.5 text-sm border border-gray-200 rounded-xl focus:ring-2 focus:ring-teal-500 outline-none transition-all"/>
          <svg class="absolute left-3.5 top-3 w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/></svg>
        </div>

        <div v-if="!form.patientId" class="mb-4 space-y-2 max-h-40 overflow-y-auto border-b pb-4">
          <div v-if="pasienList.length === 0" class="p-4 text-center text-xs text-gray-400 italic">Data pasien tidak tersedia...</div>
          <div v-for="pasien in pasienList" :key="pasien.pasien_id" @click="$emit('select-pasien', pasien)"
            class="p-3 rounded-xl cursor-pointer border border-gray-100 hover:bg-teal-50 hover:border-teal-200 flex items-center gap-3 transition-all">
            <div :class="['w-8 h-8 rounded-full flex items-center justify-center text-white text-[10px] font-bold', getColorClass(pasien.pasien_id)]">
              {{ getInitials(pasien.nama) }}
            </div>
            <div class="flex-1">
              <p class="text-xs font-bold text-slate-800">{{ pasien.nama }}</p>
              <p class="text-[9px] text-gray-400">{{ pasien.nik }}</p>
            </div>
          </div>
        </div>

        <transition name="fade">
          <div v-if="form.patientId" class="mt-4 p-5 bg-slate-50 rounded-2xl border border-slate-200 space-y-4 relative">
            <button @click="resetPasien" class="absolute top-4 right-4 text-[10px] font-bold text-teal-600 hover:text-red-500 uppercase tracking-widest">Ganti Pasien</button>
            <div class="flex items-center gap-4 border-b border-slate-200 pb-4">
              <div :class="['w-14 h-14 rounded-full flex items-center justify-center text-white font-bold text-xl shadow-inner', getColorClass(form.patientId)]">
                {{ getInitials(selectedPasienName) }}
              </div>
              <div>
                <h4 class="text-base font-bold text-slate-900 leading-tight">{{ selectedPasienName }}</h4>
                <p class="text-xs text-teal-600 font-bold tracking-wider mt-1">NIK: {{ selectedPasienNIK }}</p>
              </div>
            </div>
            <div class="grid grid-cols-2 gap-y-4 gap-x-2">
              <div>
                <p class="text-[10px] text-gray-400 uppercase font-bold tracking-tighter">Jenis Kelamin</p>
                <p class="text-xs text-slate-700 font-semibold">{{ selectedPasienJK === 'L' ? 'Laki-laki' : 'Perempuan' }}</p>
              </div>
              <div>
                <p class="text-[10px] text-gray-400 uppercase font-bold tracking-tighter">Telepon</p>
                <p class="text-xs text-slate-700 font-semibold">{{ selectedPasienTelepon || '-' }}</p>
              </div>
              <div class="col-span-2">
                <p class="text-[10px] text-gray-400 uppercase font-bold tracking-tighter">Alamat</p>
                <p class="text-xs text-slate-700 font-semibold leading-relaxed italic">"{{ selectedPasienAlamat || 'Alamat tidak tersedia' }}"</p>
              </div>
            </div>
          </div>
        </transition>
      </div>

      <div class="bg-white rounded-2xl shadow-sm border border-gray-100 p-5 md:p-6 transition-all">
        <div class="flex items-center gap-3 mb-5">
          <div class="w-10 h-10 rounded-xl bg-purple-50 flex items-center justify-center text-purple-600">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z"/></svg>
          </div>
          <h3 class="text-sm font-bold text-slate-900">Master Obat</h3>
        </div>

        <div class="relative mb-4">
          <input :value="searchObat" @input="$emit('update:searchObat', $event.target.value)"
            type="text" placeholder="Ketik nama obat untuk mencari..."
            class="w-full pl-10 pr-4 py-2.5 text-sm border border-gray-200 rounded-xl focus:ring-2 focus:ring-purple-500 outline-none transition-all"/>
          <svg class="absolute left-3.5 top-3 w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/></svg>
        </div>

        <div v-if="!form.obatId" class="mb-4 space-y-2 max-h-48 overflow-y-auto border-b pb-4">
          <div v-if="obatList.length === 0" class="p-4 text-center text-xs text-gray-400 italic">Obat tidak ditemukan...</div>
          <div v-for="obat in obatList" :key="obat.obat_id" @click="$emit('select-obat', obat)"
              class="p-3 rounded-xl cursor-pointer border border-gray-50 hover:bg-purple-50 hover:border-purple-200 flex items-center gap-3 transition-all">
            <div class="flex-1">
              <p class="text-xs font-bold text-slate-800">{{ obat.nama_obat }}</p>
              <p class="text-[9px] text-gray-400">Tersedia di Master Data</p>
            </div>
          </div>
        </div> 

        <transition name="fade">
          <div v-if="form.obatId" class="mt-4 p-5 bg-slate-50 rounded-2xl border border-slate-200 space-y-4 relative">
            <button @click="resetObat" class="absolute top-4 right-4 text-[10px] font-bold text-purple-600 hover:text-red-500 uppercase tracking-widest">Ganti Obat</button>
            
            <div class="flex items-start gap-4 border-b border-slate-200 pb-4">
              <div class="w-16 h-16 rounded-xl bg-white overflow-hidden border border-slate-200 flex-shrink-0 shadow-sm flex items-center justify-center">
                <img v-if="selectedObatGambar" :src="'https://pil-time-pam-production.up.railway.app' + selectedObatGambar" class="w-full h-full object-cover" />
                <span v-else class="text-3xl">💊</span>
              </div>
              <div class="flex-1 min-w-0">
                <h4 class="text-base font-bold text-slate-900 truncate">{{ form.nama_obat }}</h4>
                <p class="text-[10px] text-purple-600 font-bold uppercase mt-1 tracking-wider">Aturan Master</p>
                <p class="text-[11px] text-gray-500 leading-snug">{{ selectedObatAturan || 'Gunakan sesuai anjuran' }}</p>
              </div>
            </div>

            <div class="grid grid-cols-1 gap-3">
               <div>
                  <p class="text-[10px] text-gray-400 uppercase font-bold tracking-tighter">Fungsi Obat</p>
                  <p class="text-xs text-slate-700 leading-relaxed">{{ selectedObatFungsi || '-' }}</p>
               </div>
               <div v-if="selectedObatPantangan">
                  <p class="text-[10px] text-red-400 uppercase font-bold tracking-tighter">Pantangan</p>
                  <p class="text-xs text-red-600 font-medium">{{ selectedObatPantangan }}</p>
               </div>
            </div>

            <div class="mt-2 p-3 bg-blue-50 border border-blue-100 rounded-xl flex items-center gap-2">
              <svg class="w-4 h-4 text-blue-500" fill="currentColor" viewBox="0 0 20 20"><path d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"/></svg>
              <p class="text-[10px] text-blue-700 font-bold italic">Pengaturan dosis & jadwal ada di Langkah 2.</p>
            </div>
          </div>
        </transition>
      </div>
    </div>

    <div class="flex justify-end gap-3 mt-8">
      <button @click="$emit('cancel')" class="px-6 py-2.5 border border-gray-300 rounded-xl text-sm font-bold text-gray-600 hover:bg-gray-50">Batal</button>
      <button @click="$emit('next')" :disabled="!form.patientId || !form.obatId"
        class="px-8 py-2.5 bg-teal-600 text-white rounded-xl text-sm font-bold hover:bg-teal-700 disabled:opacity-40 flex items-center gap-2 transition-all">
        Lanjut ke Aturan Minum
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7l5 5m0 0l-5 5m5-5H6"/></svg>
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Step1PasienObat',
  props: {
    form: { type: Object, required: true },
    pasienList: { type: Array, default: () => [] },
    obatList: { type: Array, default: () => [] },
    searchPasien: { type: String, default: '' },
    searchObat: { type: String, default: '' },
    selectedPasienName: { type: String, default: '' },
    selectedPasienNIK: { type: String, default: '' },
    selectedPasienJK: { type: String, default: '' },
    selectedPasienTelepon: { type: String, default: '' },
    selectedPasienAlamat: { type: String, default: '' },
    selectedObatAturan: { type: String, default: '' },
    selectedObatFungsi: { type: String, default: '' },
    selectedObatPantangan: { type: String, default: '' },
    selectedObatGambar: { type: String, default: '' },
  },
  emits: ['update:form', 'update:searchPasien', 'update:searchObat', 'select-pasien', 'select-obat', 'next', 'cancel'],
  methods: {
    update(field, value) {
      this.$emit('update:form', { ...this.form, [field]: value })
    },
    resetPasien() {
      this.update('patientId', null);
      this.$emit('update:searchPasien', '');
    },
    resetObat() {
      this.update('obatId', null);
      this.$emit('update:searchObat', '');
    },
    getInitials(nama) {
      if (!nama) return '?'
      return nama.split(' ').map(n => n[0]).join('').toUpperCase().slice(0, 2)
    },
    getColorClass(id) {
      const colors = ['bg-teal-500', 'bg-purple-500', 'bg-orange-500', 'bg-rose-500', 'bg-blue-500']
      return colors[(id || 0) % colors.length]
    },
  },
}
</script>

<style scoped>
.fade-enter-active, .fade-leave-active { transition: opacity 0.3s ease; }
.fade-enter-from, .fade-leave-to { opacity: 0; }
</style>