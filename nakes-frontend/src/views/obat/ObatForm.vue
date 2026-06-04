<template>
  <div class="bg-white rounded-xl shadow-lg border border-slate-200 overflow-hidden">
    <div class="p-6 border-b border-slate-100 flex justify-between items-center bg-slate-50">
      <div>
        <h2 class="font-bold text-slate-800 text-lg">
          {{ isEditing ? 'Edit Data Obat' : 'Tambah Obat Baru' }}
        </h2>
        <p class="text-xs text-slate-500">Pastikan data obat untuk lansia Desa Nauli sudah akurat.</p>
      </div>
      <button @click="$emit('cancel')" class="text-slate-400 hover:text-slate-600 transition-colors">✕</button>
    </div>

    <form @submit.prevent="handleSubmit" class="p-6 md:p-8 space-y-6">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        
        <div class="space-y-5">
          <div>
            <label class="block text-xs font-bold text-teal-600 uppercase mb-2 tracking-wider">Nama Obat *</label>
            <input v-model="form.nama_obat" type="text" required 
              class="w-full px-4 py-2 border border-slate-300 rounded-lg focus:ring-2 focus:ring-teal-500 outline-none transition-all" 
              placeholder="Contoh: Metformin" />
          </div>

          <div>
            <label class="block text-xs font-bold text-teal-600 uppercase mb-2 tracking-wider">Fungsi Utama *</label>
            <textarea v-model="form.fungsi" rows="4" required 
              class="w-full px-4 py-2 border border-slate-300 rounded-lg focus:ring-2 focus:ring-teal-500 outline-none transition-all" 
              placeholder="Kegunaan obat ini untuk apa..."></textarea>
          </div>

          <div>
            <label class="block text-xs font-bold text-teal-600 uppercase mb-2 tracking-wider">Aturan Pemakaian *</label>
            <textarea v-model="form.aturan_pemakaian" rows="2" required 
              class="w-full px-4 py-2 border border-slate-300 rounded-lg focus:ring-2 focus:ring-teal-500 outline-none transition-all" 
              placeholder="Contoh: 3x1 hari, setelah makan pagi/siang/malam"></textarea>
          </div>
        </div>

        <div class="space-y-5">
          <div>
            <label class="block text-xs font-bold text-teal-600 uppercase mb-2 tracking-wider">Pantangan *</label>
            <input v-model="form.pantangan" type="text" required 
              class="w-full px-4 py-2 border border-slate-300 rounded-lg focus:ring-2 focus:ring-teal-500 outline-none transition-all" 
              placeholder="Hindari konsumsi bersama susu..." />
          </div>

          <div>
            <label class="block text-xs font-bold text-teal-600 uppercase mb-2 tracking-wider">Foto Obat</label>
            
            <div class="space-y-4">
              <div class="relative border-2 border-dashed rounded-2xl overflow-hidden bg-slate-50 flex flex-col items-center justify-center min-h-[220px] border-slate-300 shadow-inner">
                
                <video v-if="isCameraActive" ref="videoRef" autoplay playsinline class="w-full h-52 object-cover"></video>
                
                <img v-else-if="previewUrl" :src="previewUrl" class="h-52 w-full object-contain p-2" />
                
                <div v-else class="py-8 text-center text-slate-400">
                  <span class="text-5xl block mb-2">📸</span>
                  <p class="text-[10px] font-medium">Belum ada foto yang diambil</p>
                </div>

                <canvas ref="canvasRef" class="hidden"></canvas>
              </div>

              <div class="grid grid-cols-2 gap-3">
                <button v-if="!isCameraActive" type="button" @click="startCamera" 
                  class="flex items-center justify-center gap-2 py-2.5 bg-white border border-slate-300 text-slate-700 rounded-xl text-xs font-bold hover:bg-slate-50 transition-all shadow-sm">
                  <span>📷</span> Buka Kamera
                </button>
                
                <button v-else type="button" @click="capturePhoto" 
                  class="flex items-center justify-center gap-2 py-2.5 bg-teal-600 text-white rounded-xl text-xs font-bold hover:bg-teal-700 animate-pulse">
                  <span>⚪</span> Ambil Foto
                </button>

                <button type="button" @click="$refs.fileInput.click()" 
                  class="flex items-center justify-center gap-2 py-2.5 bg-white border border-slate-300 text-slate-700 rounded-xl text-xs font-bold hover:bg-slate-50 transition-all shadow-sm">
                  <span>📂</span> Pilih File
                </button>
                
                <input type="file" ref="fileInput" class="hidden" accept="image/*" @change="handleFileSelect" />
              </div>

              <p v-if="isCameraActive" class="text-center text-[10px] text-amber-600 font-bold italic animate-bounce">
                Kamera Aktif! Posisikan obat tepat di depan kamera.
              </p>
            </div>
          </div>
        </div>
      </div>

      <div class="pt-8 border-t border-slate-100 flex flex-col sm:flex-row justify-end gap-3">
        <button type="button" @click="$emit('cancel')" 
          class="px-8 py-2.5 text-slate-500 font-bold hover:text-slate-700 transition-all text-sm">
          Batal
        </button>
        <button type="submit" :disabled="submitting" 
          class="px-10 py-2.5 bg-teal-600 text-white rounded-xl font-bold hover:bg-teal-700 disabled:opacity-50 shadow-lg shadow-teal-100 transition-all text-sm">
          {{ submitting ? 'Menyimpan...' : (isEditing ? 'Update Data Obat' : 'Simpan Master Data') }}
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, watch } from 'vue'

const props = defineProps(['initialData', 'isEditing', 'submitting'])
const emit = defineEmits(['submit', 'cancel'])

const form = ref({
  nama_obat: '',
  fungsi: '',
  aturan_pemakaian: '',
  pantangan: '',
  imageFile: null
})

const previewUrl = ref(null)
const isCameraActive = ref(false)
const videoRef = ref(null)
const canvasRef = ref(null)
const stream = ref(null)
const fileInput = ref(null)

// Mengisi data saat edit
const fillForm = () => {
  if (props.initialData) {
    form.value.nama_obat = props.initialData.nama_obat || ''
    form.value.fungsi = props.initialData.fungsi || ''
    form.value.aturan_pemakaian = props.initialData.aturan_pemakaian || ''
    form.value.pantangan = props.initialData.pantangan || ''
    form.value.imageFile = null
    
    if (props.initialData.gambar) {
      previewUrl.value = `https://pil-time-pam-production.up.railway.app${props.initialData.gambar}`
    } else {
      previewUrl.value = null
    }
  } else {
    form.value = { nama_obat: '', fungsi: '', aturan_pemakaian: '', pantangan: '', imageFile: null }
    previewUrl.value = null
  }
}

onMounted(() => fillForm())
watch(() => props.initialData, () => fillForm(), { deep: true })

// Logika Kamera
const startCamera = async () => {
  try {
    isCameraActive.value = true
    stream.value = await navigator.mediaDevices.getUserMedia({ 
      video: { facingMode: 'environment' } 
    })
    if (videoRef.value) {
      videoRef.value.srcObject = stream.value
    }
  } catch (err) {
    alert("Gagal mengakses kamera. Pastikan izin kamera diberikan.")
    isCameraActive.value = false
  }
}

const capturePhoto = () => {
  const video = videoRef.value
  const canvas = canvasRef.value
  if (video && canvas) {
    canvas.width = video.videoWidth
    canvas.height = video.videoHeight
    canvas.getContext('2d').drawImage(video, 0, 0)
    
    canvas.toBlob((blob) => {
      const file = new File([blob], `cam_capture_${Date.now()}.jpg`, { type: "image/jpeg" })
      form.value.imageFile = file
      previewUrl.value = URL.createObjectURL(file)
      stopCamera()
    }, 'image/jpeg')
  }
}

const stopCamera = () => {
  if (stream.value) {
    stream.value.getTracks().forEach(track => track.stop())
  }
  isCameraActive.value = false
}

const handleFileSelect = (e) => {
  stopCamera()
  const file = e.target.files[0]
  if (file && file.type.startsWith('image/')) {
    form.value.imageFile = file
    previewUrl.value = URL.createObjectURL(file)
  }
}

onUnmounted(() => stopCamera())

const handleSubmit = () => {
  emit('submit', { ...form.value })
}
</script>