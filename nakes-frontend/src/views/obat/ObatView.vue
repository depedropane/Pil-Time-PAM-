<template>
  <LayoutWrapper>
    <div>
      <!-- LIST MODE -->
      <section v-if="pageMode === 'list'">
        <ObatListView
          :obat-list="obatStore.obatList"
          :search-term="searchTerm"
          :current-page="currentPage"
          @update:searchTerm="searchTerm = $event"
          @add-obat="openCreatePage"
          @edit-obat="openEditPage"
          @page-change="goToPage"
        />
      </section>

      <!-- FORM MODE -->
      <section v-else>
        <!-- Header with back button + title + step indicator -->
        <div class="flex items-center justify-between gap-3 mb-5">
          <div class="flex items-center gap-3">
            <button
              @click="backFromForm"
              class="h-9 px-3 rounded-lg border border-slate-200 text-slate-500 text-sm hover:bg-slate-100 flex items-center gap-1"
            >
              ← Kembali
            </button>

            <div>
              <h1 class="text-xl font-bold text-slate-800">Tambah Info Obat</h1>
              <p class="text-xs text-slate-500">Lengkapi semua informasi obat yang akan ditampilkan kepada pasien.</p>
            </div>
          </div>

          <!-- Step indicator top-right -->
          <div class="flex items-center gap-2 text-[11px]">
            <!-- Step 1 -->
            <div class="flex flex-col items-center">
              <div
                :class="[
                  'w-7 h-7 rounded-full flex items-center justify-center font-bold text-sm',
                  formStep >= 1 ? 'bg-teal-600 text-white' : 'bg-slate-200 text-slate-500'
                ]"
              >
                <span v-if="formStep > 1">✓</span>
                <span v-else>1</span>
              </div>
              <p :class="['mt-1 text-center', formStep === 1 ? 'text-teal-700 font-semibold' : 'text-slate-400']">
                Data Info Obat
              </p>
            </div>

            <!-- Connector line -->
            <div class="w-12 h-px bg-slate-300 mb-3"></div>

            <!-- Step 2 -->
            <div class="flex flex-col items-center">
              <div
                :class="[
                  'w-7 h-7 rounded-full flex items-center justify-center font-bold text-sm',
                  formStep >= 2 ? 'bg-teal-600 text-white' : 'bg-slate-200 text-slate-500'
                ]"
              >
                2
              </div>
              <p :class="['mt-1 text-center', formStep === 2 ? 'text-teal-700 font-semibold' : 'text-slate-400']">
                Konfirmasi
              </p>
            </div>
          </div>
        </div>

        <!-- STEP 1: Form -->
        <ObatFormStep1
          v-if="formStep === 1"
          v-model="form"
          @next-step="goToConfirmation"
        />

        <!-- STEP 2: Confirmation -->
        <ObatFormStep2
          v-else-if="formStep === 2 && !showSuccess"
          :form="form"
          :submitting="submitting"
          @back="formStep = 1"
          @submit="submitForm"
        />

        <!-- SUCCESS -->
        <ObatSuccess
          v-else
          :form="form"
          @back-to-list="finishToList"
          @add-more="resetAndAddMore"
        />
      </section>
    </div>
  </LayoutWrapper>
</template>

<script>
import { computed, onMounted, ref } from 'vue'
import LayoutWrapper from '../../components/LayoutWrapper.vue'
import ObatListView from './ObatListView.vue'
import ObatFormStep1 from './ObatFormStep1.vue'
import ObatFormStep2 from './ObatFormStep2.vue'
import ObatSuccess from './ObatSuccess.vue'
import { useObatStore } from '../../stores/obat'
import { useNotificationStore } from '../../stores/notification'
import { useConfirmStore } from '../../stores/confirm'

export default {
  name: 'ObatView',
  components: {
    LayoutWrapper,
    ObatListView,
    ObatFormStep1,
    ObatFormStep2,
    ObatSuccess
  },
  setup() {
    const obatStore = useObatStore()
    const notificationStore = useNotificationStore()
    const confirmStore = useConfirmStore()

    const pageMode = ref('list')
    const formStep = ref(1)
    const showSuccess = ref(false)
    const submitting = ref(false)

    const isEditing = ref(false)
    const editingId = ref(null)

    const searchTerm = ref('')
    const currentPage = ref(1)

    const form = ref(getInitialForm())

    const openCreatePage = () => {
      isEditing.value = false
      editingId.value = null
      pageMode.value = 'form'
      formStep.value = 1
      showSuccess.value = false
      form.value = getInitialForm()
    }

    const openEditPage = (obat) => {
      isEditing.value = true
      editingId.value = obat.obat_id
      pageMode.value = 'form'
      formStep.value = 1
      showSuccess.value = false
      form.value = {
        nama_obat: obat.nama_obat || '',
        kategori_indikasi: obat.kategori_indikasi || '',
        frekuensi_min: Number(obat.frekuensi_min || 1),
        frekuensi_max: Number(obat.frekuensi_max || 1),
        durasi_min: Number(obat.durasi_min || 1),
        durasi_max: Number(obat.durasi_max || 1),
        waktu_konsumsi: Array.isArray(obat.waktu_konsumsi) ? [...obat.waktu_konsumsi] : [],
        fungsi: obat.fungsi || '',
        aturan_penggunaan: obat.aturan_penggunaan || '',
        perhatian: obat.perhatian || ''
      }
    }

    const goToPage = (page) => {
      if (page < 1) return
      currentPage.value = page
    }

    const goToConfirmation = () => {
      if (form.value.frekuensi_min > form.value.frekuensi_max || form.value.durasi_min > form.value.durasi_max) {
        notificationStore.error('Rentang frekuensi atau durasi tidak valid.', 'Gagal')
        return
      }
      formStep.value = 2
    }

    const submitForm = async () => {
      if (submitting.value) return

      submitting.value = true
      try {
        const payload = {
          nama_obat: form.value.nama_obat.trim(),
          kategori_indikasi: form.value.kategori_indikasi.trim(),
          frekuensi_min: Number(form.value.frekuensi_min),
          frekuensi_max: Number(form.value.frekuensi_max),
          durasi_min: Number(form.value.durasi_min),
          durasi_max: Number(form.value.durasi_max),
          waktu_konsumsi: form.value.waktu_konsumsi,
          fungsi: form.value.fungsi.trim(),
          aturan_penggunaan: form.value.aturan_penggunaan.trim(),
          perhatian: form.value.perhatian.trim()
        }

        if (isEditing.value && editingId.value) {
          await obatStore.updateObat(editingId.value, payload)
        } else {
          await obatStore.createObat(payload)
        }

        showSuccess.value = true
      } catch (error) {
        notificationStore.error('Gagal menyimpan data: ' + error.message, 'Gagal')
      } finally {
        submitting.value = false
      }
    }

    const backFromForm = async () => {
      if (showSuccess.value) {
        finishToList()
        return
      }

      if (formStep.value === 2) {
        formStep.value = 1
        return
      }

      const ok = await confirmStore.show({
        title: 'Batalkan Pengisian?',
        message: 'Apakah Anda yakin ingin membatalkan pengisian data obat ini? Data yang belum disimpan akan hilang.',
        confirmText: 'Ya, Batal',
        cancelText: 'Kembali',
        type: 'warning'
      })
      if (ok) {
        pageMode.value = 'list'
      }
    }

    const finishToList = () => {
      pageMode.value = 'list'
      showSuccess.value = false
      formStep.value = 1
      isEditing.value = false
      editingId.value = null
      form.value = getInitialForm()
      currentPage.value = 1
      searchTerm.value = ''
    }

    const resetAndAddMore = () => {
      isEditing.value = false
      editingId.value = null
      showSuccess.value = false
      formStep.value = 1
      form.value = getInitialForm()
    }

    onMounted(() => {
      obatStore.fetchObats()
    })

    return {
      obatStore,
      pageMode,
      formStep,
      form,
      submitting,
      showSuccess,
      searchTerm,
      currentPage,
      openCreatePage,
      openEditPage,
      goToPage,
      goToConfirmation,
      submitForm,
      backFromForm,
      finishToList,
      resetAndAddMore
    }
  }
}

function getInitialForm() {
  return {
    nama_obat: '',
    kategori_indikasi: '',
    frekuensi_min: 1,
    frekuensi_max: 1,
    durasi_min: 1,
    durasi_max: 1,
    waktu_konsumsi: [],
    fungsi: '',
    aturan_penggunaan: '',
    perhatian: ''
  }
}
</script>
