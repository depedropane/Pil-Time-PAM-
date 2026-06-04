<template>
  <div class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50 overflow-y-auto p-4">
    <div class="bg-white rounded-xl shadow-xl w-full max-w-2xl my-8">
      <!-- Header -->
      <div class="flex justify-between items-center border-b border-gray-200 p-6 md:p-8">
        <div>
          <h2 class="text-2xl md:text-3xl font-bold text-slate-900">
            {{ editingObat?.obat_id ? 'Edit Info Obat' : 'Tambah Info Obat' }}
          </h2>
          <p class="text-xs md:text-sm text-gray-500 mt-1">
            Lengkapi semua informasi obat untuk pasien
          </p>
        </div>
        <button 
          @click="$emit('close')"
          class="text-gray-400 hover:text-gray-600 text-2xl font-bold"
        >
          ×
        </button>
      </div>

      <!-- Content -->
      <div class="px-6 md:px-8 py-6 space-y-6 max-h-[calc(100vh-250px)] overflow-y-auto">
        <!-- Gambar Obat (Image Upload) -->
        <ImageUpload 
          :initial-image="form.gambar"
          @update:image="form.gambar = $event"
        />

        <!-- Identitas Obat -->
        <div class="bg-blue-50 rounded-lg border border-blue-200 p-4">
          <h3 class="text-sm font-semibold text-gray-700 mb-4">Identitas Obat</h3>
          <div class="space-y-3">
            <label class="block">
              <span class="text-sm font-medium text-gray-700 mb-1 block">
                Nama Obat <span class="text-red-500">*</span>
              </span>
              <input 
                v-model="form.nama_obat" 
                type="text" 
                placeholder="cth. Paracetamol, Amoxicillin..."
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none text-sm"
              />
            </label>

            <label class="block">
              <span class="text-sm font-medium text-gray-700 mb-1 block">
                Kategori / Indikasi <span class="text-red-500">*</span>
              </span>
              <input 
                v-model="form.kategori_indikasi" 
                type="text" 
                placeholder="cth. Pereda Nyeri & Demam"
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none text-sm"
              />
            </label>
          </div>
        </div>

        <!-- Aturan Penggunaan -->
        <div class="bg-blue-50 rounded-lg border border-blue-200 p-4">
          <h3 class="text-sm font-semibold text-gray-700 mb-4">Aturan Penggunaan</h3>
          <div class="space-y-3">
            <div class="grid grid-cols-2 gap-3">
              <label class="block">
                <span class="text-sm font-medium text-gray-700 mb-1 block">
                  Frekuensi Min <span class="text-red-500">*</span>
                </span>
                <input 
                  v-model.number="form.frekuensi_min" 
                  type="number" 
                  min="1"
                  placeholder="1"
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none text-sm"
                />
              </label>
              <label class="block">
                <span class="text-sm font-medium text-gray-700 mb-1 block">
                  Frekuensi Max <span class="text-red-500">*</span>
                </span>
                <input 
                  v-model.number="form.frekuensi_max" 
                  type="number" 
                  min="1"
                  placeholder="1"
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none text-sm"
                />
              </label>
            </div>

            <div class="grid grid-cols-2 gap-3">
              <label class="block">
                <span class="text-sm font-medium text-gray-700 mb-1 block">
                  Durasi Min (hari) <span class="text-red-500">*</span>
                </span>
                <input 
                  v-model.number="form.durasi_min" 
                  type="number" 
                  min="1"
                  placeholder="1"
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none text-sm"
                />
              </label>
              <label class="block">
                <span class="text-sm font-medium text-gray-700 mb-1 block">
                  Durasi Max (hari) <span class="text-red-500">*</span>
                </span>
                <input 
                  v-model.number="form.durasi_max" 
                  type="number" 
                  min="1"
                  placeholder="1"
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none text-sm"
                />
              </label>
            </div>

            <label class="block">
              <span class="text-sm font-medium text-gray-700 mb-2 block">
                Waktu Konsumsi <span class="text-red-500">*</span>
              </span>
              <div class="flex flex-wrap gap-2">
                <button
                  v-for="option in waktuOptions"
                  :key="option"
                  @click="toggleWaktu(option)"
                  :class="[
                    'px-3 py-1.5 rounded-full border text-xs transition',
                    form.waktu_konsumsi.includes(option)
                      ? 'border-blue-600 text-blue-700 bg-blue-50 font-medium'
                      : 'border-gray-300 text-gray-500 bg-white hover:bg-gray-50'
                  ]"
                >
                  {{ option }}
                </button>
              </div>
            </label>
          </div>
        </div>

        <!-- Informasi Klinis -->
        <div class="bg-purple-50 rounded-lg border border-purple-200 p-4">
          <h3 class="text-sm font-semibold text-gray-700 mb-4">Informasi Klinis</h3>
          <div class="space-y-3">
            <label class="block">
              <span class="text-sm font-medium text-gray-700 mb-1 block">
                Fungsi Obat <span class="text-red-500">*</span>
              </span>
              <textarea 
                v-model="form.fungsi"
                rows="3"
                placeholder="Jelaskan fungsi utama obat ini secara singkat dan jelas..."
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent outline-none text-sm resize-none"
              ></textarea>
            </label>

            <label class="block">
              <span class="text-sm font-medium text-gray-700 mb-1 block">
                Aturan Pakai <span class="text-red-500">*</span>
              </span>
              <textarea 
                v-model="form.aturan_penggunaan"
                rows="3"
                placeholder="Petunjuk lengkap cara mengonsumsi, dosis, dan interval minimum..."
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent outline-none text-sm resize-none"
              ></textarea>
            </label>

            <label class="block">
              <span class="text-sm font-medium text-gray-700 mb-1 block">
                Perhatian / Peringatan <span class="text-red-500">*</span>
              </span>
              <textarea 
                v-model="form.perhatian"
                rows="3"
                placeholder="Peringatan penting, efek samping, kondisi khusus, dan kontra indikasi..."
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent outline-none text-sm resize-none"
              ></textarea>
            </label>
          </div>
        </div>
      </div>

      <!-- Footer -->
      <div class="border-t border-gray-200 px-6 md:px-8 py-4 flex flex-col-reverse md:flex-row justify-end gap-3">
        <button 
          @click="$emit('close')"
          class="w-full md:w-auto px-6 py-2.5 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50 transition font-medium text-sm md:text-base flex items-center justify-center gap-2"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
          </svg>
          Batal
        </button>
        <button 
          @click="handleSubmit"
          :disabled="!isFormValid"
          :class="['w-full md:w-auto px-6 py-2.5 bg-teal-600 text-white rounded-lg transition font-medium text-sm md:text-base flex items-center justify-center gap-2',
            isFormValid ? 'hover:bg-teal-700 cursor-pointer' : 'opacity-50 cursor-not-allowed']"
        >
          Simpan Info Obat
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed } from 'vue'
import ImageUpload from './ImageUpload.vue'

export default {
  name: 'InfoObatForm',
  components: {
    ImageUpload
  },
  props: {
    editingObat: { type: Object, default: null }
  },
  emits: ['close', 'submit'],
  setup(props, { emit }) {
    const form = ref({
      nama_obat: props.editingObat?.nama_obat || '',
      kategori_indikasi: props.editingObat?.kategori_indikasi || '',
      frekuensi_min: Number(props.editingObat?.frekuensi_min || 1),
      frekuensi_max: Number(props.editingObat?.frekuensi_max || 1),
      durasi_min: Number(props.editingObat?.durasi_min || 1),
      durasi_max: Number(props.editingObat?.durasi_max || 1),
      waktu_konsumsi: Array.isArray(props.editingObat?.waktu_konsumsi) ? [...props.editingObat.waktu_konsumsi] : [],
      fungsi: props.editingObat?.fungsi || '',
      aturan_penggunaan: props.editingObat?.aturan_penggunaan || '',
      perhatian: props.editingObat?.perhatian || '',
      gambar: props.editingObat?.gambar || null
    })

    const waktuOptions = ['Sebelum makan', 'Sesudah makan', 'Saat makan', 'Kapan saja']

    const isFormValid = computed(() => {
      return form.value.nama_obat?.trim() && 
             form.value.kategori_indikasi?.trim() &&
             Number(form.value.frekuensi_min) > 0 &&
             Number(form.value.frekuensi_max) > 0 &&
             Number(form.value.durasi_min) > 0 &&
             Number(form.value.durasi_max) > 0 &&
             form.value.waktu_konsumsi.length > 0 &&
             form.value.fungsi?.trim() &&
             form.value.aturan_penggunaan?.trim() &&
             form.value.perhatian?.trim()
    })

    const toggleWaktu = (value) => {
      if (form.value.waktu_konsumsi.includes(value)) {
        form.value.waktu_konsumsi = form.value.waktu_konsumsi.filter((item) => item !== value)
      } else {
        form.value.waktu_konsumsi = [...form.value.waktu_konsumsi, value]
      }
    }

    const handleSubmit = async () => {
      if (!isFormValid.value) return
      
      try {
        emit('submit', form.value)
      } catch (error) {
        console.error('Error submitting form:', error)
      }
    }

    return {
      form,
      waktuOptions,
      isFormValid,
      toggleWaktu,
      handleSubmit
    }
  }
}
</script>
