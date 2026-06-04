<template>
  <LayoutWrapper>
    <div class="p-4 md:p-8">
      
      <transition name="fade">
        <div v-if="isDetailOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm" @click.self="closeDetail">
          <div class="bg-white rounded-2xl shadow-2xl w-full max-w-lg overflow-hidden animate-modal">
            <div class="bg-teal-600 p-4 flex justify-between items-center text-white">
              <div class="flex items-center gap-2">
                <span class="text-xl">💊</span>
                <h3 class="font-bold uppercase tracking-wider text-sm">Detail Informasi Obat</h3>
              </div>
              <button @click="closeDetail" class="hover:bg-teal-700 p-1 rounded-full transition-colors text-xl w-8 h-8 flex items-center justify-center">✕</button>
            </div>
            
            <div class="p-6 space-y-6">
              <div class="flex justify-center bg-slate-50 rounded-xl p-4 border border-dashed border-slate-200">
                <img 
                  v-if="selectedObat?.gambar"
                  :src="getImageUrl(selectedObat.gambar)" 
                  class="h-48 w-full object-contain rounded-lg shadow-sm" 
                  @error="(e) => e.target.src = 'https://via.placeholder.com/400x200?text=Gambar+Tidak+Ditemukan'"
                  alt="Foto Obat"
                />
                <div v-else class="h-48 w-full flex flex-col items-center justify-center text-slate-400">
                  <span class="text-5xl">🖼️</span>
                  <p class="text-xs mt-2 font-medium">Foto obat belum diunggah</p>
                </div>
              </div>

              <div class="space-y-4 text-left">
                <div>
                  <label class="text-[10px] font-bold text-teal-600 uppercase tracking-widest">Nama Obat</label>
                  <p class="text-slate-800 font-bold text-xl">{{ selectedObat?.nama_obat }}</p>
                </div>

                <div>
                  <label class="text-[10px] font-bold text-teal-600 uppercase tracking-widest block mb-1">Fungsi Utama</label>
                  <p class="text-slate-700 text-sm leading-relaxed">{{ selectedObat?.fungsi || '-' }}</p>
                </div>

                <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                  <div class="bg-teal-50 p-3 rounded-xl border border-teal-100 text-left">
                    <label class="text-[10px] font-bold text-teal-600 uppercase block mb-1">Aturan Pemakaian</label>
                    <p class="text-slate-800 font-semibold text-sm">{{ selectedObat?.aturan_pemakaian || '-' }}</p>
                  </div>
                  <div class="bg-red-50 p-3 rounded-xl border border-red-100 text-left">
                    <label class="text-[10px] font-bold text-red-500 uppercase block mb-1">Pantangan</label>
                    <p class="text-red-700 font-semibold text-sm">{{ selectedObat?.pantangan || '-' }}</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </transition>

      <transition name="fade" mode="out-in">
        <div v-if="isFormOpen">
          <ObatForm 
            :initial-data="selectedObat" 
            :is-editing="!!selectedObat"
            :submitting="loading"
            @submit="handleFormSubmit"
            @cancel="closeForm"
          />
        </div>

        <div v-else class="space-y-6">
          <div class="w-full">
            <ObatListView 
              :obat-list="obatStore.obatList"
              @edit-obat="openEditForm"
              @select-obat="openDetailModal" 
              @add-obat="openAddForm"
              :selected-id="selectedObat?.obat_id"
            />
          </div>
        </div>
      </transition>
    </div>
  </LayoutWrapper>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import LayoutWrapper from '../components/LayoutWrapper.vue'
import { useObatStore } from '../stores/obat.js'
import ObatListView from './obat/ObatListView.vue'
import ObatForm from './obat/ObatForm.vue'

const obatStore = useObatStore()
const isFormOpen = ref(false)
const isDetailOpen = ref(false)
const selectedObat = ref(null)
const loading = ref(false)

onMounted(() => {
  obatStore.fetchObats()
})

const getImageUrl = (path) => {
  if (!path) return ''
  if (path.startsWith('http')) return path
  const baseUrl = 'https://pil-time-pam-production.up.railway.app'
  return `${baseUrl}${path}`
}

const openAddForm = () => {
  selectedObat.value = null
  isFormOpen.value = true
}

const openEditForm = (obat) => {
  selectedObat.value = { ...obat }
  isFormOpen.value = true
}

const openDetailModal = (obat) => {
  selectedObat.value = obat
  isDetailOpen.value = true
}

const closeDetail = () => {
  isDetailOpen.value = false
}

const closeForm = () => {
  isFormOpen.value = false
  selectedObat.value = null
}

const handleFormSubmit = async (formData) => {
  loading.value = true
  try {
    const data = new FormData()
    data.append('nama_obat', formData.nama_obat)
    data.append('fungsi', formData.fungsi)
    data.append('aturan_pemakaian', formData.aturan_pemakaian)
    data.append('pantangan', formData.pantangan)
    
    if (formData.imageFile instanceof File) {
      data.append('gambar_file', formData.imageFile)
      console.log("📁 File ditemukan:", formData.imageFile.name)
    }

    // Perbaikan: Gunakan selectedObat.value secara langsung, bukan isEditing yang tidak didefinisikan
    if (selectedObat.value && selectedObat.value.obat_id) {
      await obatStore.updateObat(selectedObat.value.obat_id, data)
    } else {
      await obatStore.createObat(data)
    }
    
    closeForm()
    await obatStore.fetchObats()
  } catch (error) {
    console.error("Kesalahan saat simpan:", error)
    alert("Gagal menyimpan data obat")
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

@keyframes modal-appear {
  from { opacity: 0; transform: translateY(20px) scale(0.95); }
  to { opacity: 1; transform: translateY(0) scale(1); }
}

.animate-modal {
  animation: modal-appear 0.3s ease-out forwards;
}
</style>