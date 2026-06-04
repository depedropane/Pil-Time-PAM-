import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { riwayatService } from '../services'

export const useRiwayatStore = defineStore('riwayat', () => {
  // ── API State ──────────────────────────────────────────────
  const riwayatList = ref([])
  const loading = ref(false)
  const error = ref(null)

  // ── UI State ───────────────────────────────────────────────
  const searchQuery = ref('')
  const filterStatus = ref('all') // 'all', 'diminum', 'terlambat', 'terlewat'

  // ── Statistics ─────────────────────────────────────────────
  const statistics = computed(() => {
    const stats = {
      diminum: 0,
      terlambat: 0,
      terlewat: 0
    }

    if (!Array.isArray(riwayatList.value)) {
      return stats
    }

    riwayatList.value.forEach(item => {
      if (item?.status === 'Diminum') {
        stats.diminum++
      } else if (item?.status === 'Terlambat') {
        stats.terlambat++
      } else if (item?.status === 'Terlewat') {
        stats.terlewat++
      }
    })

    return stats
  })

  // ── Filtered Data ──────────────────────────────────────────
  const filteredRiwayat = computed(() => {
    if (!Array.isArray(riwayatList.value)) {
      return []
    }

    let filtered = riwayatList.value

    // Filter by status
    if (filterStatus.value === 'diminum') {
      filtered = filtered.filter(item => item?.status === 'Diminum')
    } else if (filterStatus.value === 'terlambat') {
      filtered = filtered.filter(item => item?.status === 'Terlambat')
    } else if (filterStatus.value === 'terlewat') {
      filtered = filtered.filter(item => item?.status === 'Terlewat')
    }

    // Filter by search query
    if (searchQuery.value) {
      const query = searchQuery.value.toLowerCase()
      filtered = filtered.filter(item =>
        item?.namaPasien?.toLowerCase().includes(query) ||
        item?.namaObat?.toLowerCase().includes(query)
      )
    }

    return filtered
  })

  // ── Actions ────────────────────────────────────────────────
  const fetchRiwayat = async () => {
    loading.value = true
    error.value = null

    try {
      const response = await riwayatService.getAllRiwayat()
      const data = response?.data?.data || response?.data || []
      riwayatList.value = Array.isArray(data) ? data.map(item => ({
        ...item,
        namaPasien: item?.namaPasien || item?.nama_pasien || '-',
        namaObat: item?.namaObat || item?.nama_obat || '-',
        waktuMinum: item?.waktuMinum || item?.waktu_minum || '-',
        jadwal: item?.jadwal || '-',
        pasienId: item?.pasienId || item?.pasien_id,
        jadwalId: item?.jadwalId || item?.jadwal_id,
        buktiFoto: item?.buktiFoto || item?.bukti_foto
      })) : []
    } catch (err) {
      console.error('Error fetching riwayat:', err)
      error.value = err?.message || 'Gagal mengambil data riwayat'
      riwayatList.value = []
    } finally {
      loading.value = false
    }
  }

  const setSearchQuery = (query) => {
    searchQuery.value = query
  }

  const setFilterStatus = (status) => {
    filterStatus.value = status
  }

  const resetFilters = () => {
    searchQuery.value = ''
    filterStatus.value = 'all'
  }

  return {
    riwayatList,
    loading,
    error,
    searchQuery,
    filterStatus,
    statistics,
    filteredRiwayat,
    fetchRiwayat,
    setSearchQuery,
    setFilterStatus,
    resetFilters
  }
})
