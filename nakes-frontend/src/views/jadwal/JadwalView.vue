<template>
  <LayoutWrapper>
    <div class="min-h-screen bg-[#f8fafc] p-4 md:p-8">
      
      <!-- HEADER SECTION (Hanya muncul di daftar utama) -->
      <div v-if="jadwal.currentStep === 0" class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-8">
        <div>
          <h1 class="text-2xl md:text-[28px] font-bold text-slate-800">Jadwal Obat</h1>
        </div>
        <div class="flex flex-col sm:flex-row gap-3 items-stretch sm:items-center w-full md:w-auto">
          <div class="relative w-full sm:w-72">
            <span class="absolute inset-y-0 left-0 pl-3 flex items-center text-slate-400">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </span>
            <input v-model="jadwal.searchQuery" 
              class="search-input !w-full" 
              placeholder="Cari pasien atau nama obat..." />
          </div>
          <button @click="jadwal.openAddSchedule" class="btn-primary flex items-center justify-center gap-2 py-2.5">
            <span class="text-xl leading-none">+</span> Tambah Jadwal Obat
          </button>
        </div>
      </div>

      <!-- STEP PROGRESS INDICATOR (Muncul saat tambah data) -->
      <div v-if="jadwal.currentStep > 0 && jadwal.currentStep < 3" class="mb-8 bg-white p-4 rounded-2xl border border-slate-100 shadow-sm">
        <div class="flex items-center justify-center gap-2 sm:gap-6">
          <div v-for="(s, i) in jadwal.steps" :key="i" class="flex items-center gap-2">
            <div :class="['step-circle', jadwal.currentStep >= i+1 ? 'active' : 'inactive']">
              {{ i+1 }}
            </div>
            <span :class="['text-xs sm:text-sm font-semibold', jadwal.currentStep >= i+1 ? 'text-slate-800' : 'text-slate-400', 'hidden sm:inline-block']">{{ s }}</span>
            <div v-if="i < jadwal.steps.length - 1" class="w-6 sm:w-12 h-[2px] bg-slate-200 ml-1"></div>
          </div>
        </div>
      </div>

      <!-- MAIN CONTENT AREA -->
      <main>
        <!-- TABLE VIEW (Step 0) -->
        <div v-if="jadwal.currentStep === 0" class="pb-32">
          <!-- Desktop Table View -->
          <div class="hidden md:block bg-white rounded-[20px] border border-slate-100 shadow-[0_8px_30px_rgb(0,0,0,0.04)] overflow-hidden">
            <table class="w-full text-left border-collapse">
              <thead>
                <tr class="bg-slate-50/50 border-b border-slate-100">
                  <th class="table-th">Pasien</th>
                  <th class="table-th">Nama Obat</th>
                  <th class="table-th">Dosis</th>
                  <th class="table-th">Frekuensi</th>
                  <th class="table-th">Waktu</th>
                  <th class="table-th">Durasi</th>
                  <th class="table-th text-center">Aksi</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-slate-50">
                <tr v-for="j in jadwal.filteredJadwalList" :key="j.id" class="hover:bg-slate-50/50 transition-colors">
                  <td class="table-td font-semibold text-slate-700">{{ j.pasien_nama }}</td>
                  <td class="table-td text-slate-600">{{ j.nama_obat || 'Paracetamol' }}</td>
                  <td class="table-td text-slate-600">{{ j.jumlah_dosis || '1' }} {{ j.satuan || 'Tablet' }}</td>
                  <td class="table-td text-slate-600">{{ j.frekuensi_per_hari ? j.frekuensi_per_hari + 'x sehari' : '1x sehari' }}</td>
                  <td class="table-td text-slate-600 text-xs">{{ j.waktu_minum || '-' }}</td>
                  <td class="table-td">
                    <span :class="['px-3 py-1 rounded-full text-[11px] font-bold uppercase tracking-wider', 
                      j.tipe_durasi === 'hari' ? 'bg-blue-50 text-blue-500' : 'bg-teal-50 text-teal-600']">
                      {{ j.tipe_durasi === 'hari' ? j.jumlah_hari + ' Hari' : 'Rutin' }}
                    </span>
                  </td>
                  <td class="table-td text-center relative">
                    <div class="relative inline-block text-left" :ref="el => { if(el) menuRefs[j.id] = el }">
                      <button @click.stop="toggleDropdown(j.id)" class="action-btn">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-slate-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
                        </svg>
                      </button>
                      
                      <Teleport to="body">
                        <div v-if="openDropdownId === j.id" @click.stop :style="getDropdownStyle(j.id)" class="fixed w-40 bg-white border border-slate-100 rounded-xl shadow-lg z-50 py-1 overflow-hidden">
                          <button @click="viewDetail(j.id); closeDropdown()" class="w-full text-left px-4 py-2.5 text-sm text-slate-600 hover:bg-slate-50 hover:text-teal-600 transition-colors flex items-center gap-2">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" /><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" /></svg>
                            Lihat Detail
                          </button>
                          <button @click="editJadwal(j.id); closeDropdown()" class="w-full text-left px-4 py-2.5 text-sm text-slate-600 hover:bg-slate-50 hover:text-teal-600 transition-colors flex items-center gap-2">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" /></svg>
                            Edit
                          </button>
                          <button @click="deleteJadwal(j.id); closeDropdown()" class="w-full text-left px-4 py-2.5 text-sm text-red-500 hover:bg-red-50 transition-colors flex items-center gap-2">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" /></svg>
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
          <div class="md:hidden space-y-4">
            <div v-for="j in jadwal.filteredJadwalList" :key="j.id" 
              class="bg-white rounded-2xl border border-slate-100 p-5 shadow-[0_8px_30px_rgb(0,0,0,0.02)] space-y-4 relative text-left">
              
              <div class="flex justify-between items-start gap-4">
                <div class="space-y-1">
                  <span class="inline-flex items-center gap-1 bg-teal-50 text-teal-700 font-bold px-2 py-0.5 rounded text-[10px] uppercase">
                    👤 {{ j.pasien_nama }}
                  </span>
                  <h3 class="font-bold text-slate-800 text-base leading-snug">{{ j.nama_obat || 'Paracetamol' }}</h3>
                </div>

                <!-- Actions Dropdown -->
                <div class="relative inline-block text-left" :ref="el => { if(el) menuRefs['m_' + j.id] = el }">
                  <button @click.stop="toggleDropdown('m_' + j.id)" class="p-2 text-slate-400 hover:text-teal-600 border border-slate-100 rounded-lg hover:bg-white transition-all shadow-sm">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
                    </svg>
                  </button>
                  
                  <Teleport to="body">
                    <div v-if="openDropdownId === 'm_' + j.id" @click.stop :style="getDropdownStyle('m_' + j.id)" class="fixed w-40 bg-white rounded-xl shadow-lg border border-slate-100 z-50 py-1 overflow-hidden">
                      <button @click="viewDetail(j.id); closeDropdown()" class="w-full text-left px-3 py-2 text-xs text-slate-600 hover:bg-slate-50 hover:text-teal-600 transition-colors flex items-center gap-2">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" /><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" /></svg>
                        Lihat Detail
                      </button>
                      <button @click="editJadwal(j.id); closeDropdown()" class="w-full text-left px-3 py-2 text-xs text-slate-600 hover:bg-slate-50 hover:text-teal-600 transition-colors flex items-center gap-2">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" /></svg>
                        Edit
                      </button>
                      <button @click="deleteJadwal(j.id); closeDropdown()" class="w-full text-left px-3 py-2 text-xs text-red-500 hover:bg-red-50 transition-colors flex items-center gap-2">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" /></svg>
                        Hapus
                      </button>
                    </div>
                  </Teleport>
                </div>
              </div>

              <!-- Metadata Grid -->
              <div class="grid grid-cols-2 gap-y-3 gap-x-4 border-t border-slate-50 pt-3 text-xs">
                <div>
                  <span class="text-slate-400 font-medium block">Dosis</span>
                  <span class="text-slate-700 font-semibold">{{ j.jumlah_dosis || '1' }} {{ j.satuan || 'Tablet' }}</span>
                </div>
                <div>
                  <span class="text-slate-400 font-medium block">Frekuensi</span>
                  <span class="text-slate-700 font-semibold">{{ j.frekuensi_per_hari ? j.frekuensi_per_hari + 'x sehari' : '1x sehari' }}</span>
                </div>
                <div>
                  <span class="text-slate-400 font-medium block">Waktu Minum</span>
                  <span class="text-teal-600 font-bold bg-teal-50 px-2 py-0.5 rounded inline-block">{{ j.waktu_minum || '-' }}</span>
                </div>
                <div>
                  <span class="text-slate-400 font-medium block">Durasi</span>
                  <span :class="['inline-block px-2 py-0.5 rounded font-bold uppercase text-[10px] tracking-wider', 
                    j.tipe_durasi === 'hari' ? 'bg-blue-50 text-blue-500' : 'bg-emerald-50 text-emerald-600']">
                    {{ j.tipe_durasi === 'hari' ? j.jumlah_hari + ' Hari' : 'Rutin' }}
                  </span>
                </div>
              </div>
            </div>
          </div>

          <div v-if="jadwal.filteredJadwalList.length === 0" class="p-12 text-center text-slate-400 bg-white rounded-2xl border border-slate-100 shadow-[0_8px_30px_rgb(0,0,0,0.02)]">
             Belum ada data jadwal obat.
          </div>
        </div>

        <!-- STEP 1: PILIH PASIEN & OBAT -->
        <div v-if="jadwal.currentStep === 1" class="grid md:grid-cols-2 gap-8 animate-in fade-in duration-500">
          <div class="card">
            <h3 class="title">Pilih Pasien</h3>
            <p class="text-sm text-slate-400 mb-3">Cari dan pilih pasien yang akan dijadwalkan</p>

            <input v-model="jadwal.searchPasien" class="input mb-4" placeholder="Cari nama pasien..." />

            <div class="list-container">
              <div v-for="p in jadwal.filteredPasien" :key="p.pasien_id"
                @click="jadwal.selectPasien(p)"
                :class="['list-item flex items-center justify-between', jadwal.form.patientId === p.pasien_id ? 'active' : '']">

                <div class="flex items-center gap-3">
                  <div class="w-8 h-8 rounded-full bg-teal-100 flex items-center justify-center text-sm font-bold text-teal-600">
                    {{ p.nama.charAt(0) }}
                  </div>
                  <span>{{ p.nama }}</span>
                </div>

                <span v-if="jadwal.form.patientId === p.pasien_id" class="text-teal-500">✔</span>
              </div>
            </div>

            <!-- PASIEN TERPILIH -->
            <div v-if="jadwal.getSelectedPasienName()" class="mt-4 bg-teal-50 p-3 rounded-xl text-sm text-teal-700">
              ✔ Pasien Terpilih: <b>{{ jadwal.getSelectedPasienName() }}</b>
            </div>
          </div>

          <div class="card">
            <h3 class="title">Master Obat</h3>
            <p class="text-sm text-slate-400 mb-3">Pilih obat dari data master obat</p>

            <input v-model="jadwal.searchObat" class="input mb-4" placeholder="Cari nama obat..." />

            <div class="list-container">
              <div v-for="o in jadwal.filteredObat" :key="o.obat_id"
                @click="jadwal.selectObat(o)"
                :class="['list-item flex items-center justify-between', jadwal.form.obatId === o.obat_id ? 'active' : '']">

                <div class="flex items-center gap-3">
                  <div class="w-8 h-8 rounded-full bg-blue-100 flex items-center justify-center text-sm font-bold text-blue-600">
                    {{ o.nama_obat.charAt(0) }}
                  </div>
                  <span>{{ o.nama_obat }}</span>
                </div>

                <span v-if="jadwal.form.obatId === o.obat_id" class="text-teal-500">✔</span>
              </div>
            </div>

            <!-- OBAT TERPILIH -->
            <div v-if="jadwal.form.obatId" class="mt-4 bg-teal-50 p-3 rounded-xl text-sm text-teal-700">
              ✔ Obat Terpilih: <b>{{ jadwal.form.nama_obat }}</b>
            </div>
          </div>

        </div>

        <!-- STEP 2: DETAIL JADWAL -->
<div v-if="jadwal.currentStep === 2" class="grid md:grid-cols-2 gap-8">

  <!-- KIRI -->
  <div class="card">
    <h3 class="title">Frekuensi & Waktu Minum</h3>
    <p class="text-sm text-slate-400 mb-4">Tentukan seberapa sering dan kapan obat diminum</p>

    <!-- FREKUENSI -->
    <div class="mb-4">
      <label class="label">Frekuensi per Hari</label>
      <input 
        type="text" 
        :value="jadwal.selectedWaktuMinum.length + 'x sehari'" 
        class="input bg-slate-100" 
        readonly
      />
    </div>

    <!-- WAKTU MINUM -->
    <div class="mb-4">
      <label class="label">Waktu Minum</label>
      <div class="flex flex-wrap gap-2 mt-2">
        <button 
          v-for="w in jadwal.waktuMinumOptions" 
          :key="w.value"
          @click="jadwal.toggleWaktuMinum(w.value)"
          :class="[
            'px-3 py-1.5 rounded-lg text-sm border transition',
            jadwal.selectedWaktuMinum.includes(w.value)
              ? 'bg-teal-100 text-teal-700 border-teal-300'
              : 'bg-white text-slate-500 border-slate-200'
          ]"
        >
          {{ w.icon }} {{ w.label }}
        </button>
      </div>
    </div>

    <!-- ATURAN KONSUMSI -->
    <div class="mb-4">
      <label class="label">Aturan Konsumsi</label>
      <div class="flex gap-2 mt-2 flex-wrap">
        <button
          v-for="a in jadwal.aturanKonsumsi"
          :key="a"
          @click="jadwal.form.aturan_konsumsi = a"
          :class="[
            'px-3 py-1.5 rounded-lg text-sm border',
            jadwal.form.aturan_konsumsi === a
              ? 'bg-teal-100 text-teal-700 border-teal-300'
              : 'bg-white text-slate-500 border-slate-200'
          ]"
        >
          {{ a }}
        </button>
      </div>
    </div>

    <!-- CATATAN -->
    <div>
      <label class="label">Catatan untuk Pasien</label>
      <textarea 
        v-model="jadwal.form.catatan" 
        class="input h-24"
        placeholder="Contoh: Minum dengan segelas air..."
      ></textarea>
    </div>
  </div>

  <!-- KANAN -->
  <div class="card">
    <h3 class="title">Durasi & Jadwal</h3>
    <p class="text-sm text-slate-400 mb-4">Tentukan lama pengobatan dan tanggal mulai</p>

    <!-- TIPE DURASI -->
    <div class="mb-4">
      <label class="label">Tipe Durasi</label>
      <div class="flex gap-2 mt-2">

        <!-- JUMLAH HARI -->
        <button 
          @click="jadwal.form.tipe_durasi = 'hari'"
          :class="[
            'flex-1 py-2 rounded-lg text-sm',
            jadwal.form.tipe_durasi === 'hari'
              ? 'bg-teal-500 text-white'
              : 'bg-white border border-slate-200'
          ]"
        >
          Jumlah Hari
        </button>

        <!-- RUTIN -->
        <button 
          @click="jadwal.form.tipe_durasi = 'rutin'"
          :class="[
            'flex-1 py-2 rounded-lg text-sm',
            jadwal.form.tipe_durasi === 'rutin'
              ? 'bg-teal-500 text-white'
              : 'bg-white border border-slate-200'
          ]"
        >
          Rutin
        </button>

      </div>
    </div>

    <div v-if="jadwal.form.tipe_durasi === 'hari'" class="mb-4">
      <label class="label">Jumlah Hari</label>

      <!-- INPUT + BUTTON -->
      <div class="flex items-center gap-2">
        
        <!-- MINUS -->
        <button 
          @click="jadwal.form.jumlah_hari = Math.max(1, (jadwal.form.jumlah_hari || 1) - 1)"
          class="px-3 py-2 rounded-lg bg-slate-100 hover:bg-slate-200"
        >
          -
        </button>

        <!-- INPUT -->
        <input 
          type="number"
          v-model="jadwal.form.jumlah_hari"
          class="input text-center"
          min="1"
        />

        <!-- PLUS -->
        <button 
          @click="jadwal.form.jumlah_hari = (jadwal.form.jumlah_hari || 1) + 1"
          class="px-3 py-2 rounded-lg bg-slate-100 hover:bg-slate-200"
        >
          +
        </button>

        <span class="text-sm text-slate-500">hari</span>
      </div>

      <!-- QUICK SELECT -->
      <div class="flex flex-wrap gap-2 mt-3">
        <button 
          @click="jadwal.form.jumlah_hari = 3"
          class="chip"
        >3 hari</button>

        <button 
          @click="jadwal.form.jumlah_hari = 7"
          class="chip"
        >7 hari</button>

        <button 
          @click="jadwal.form.jumlah_hari = 14"
          class="chip"
        >2 minggu</button>

        <button 
          @click="jadwal.form.jumlah_hari = 30"
          class="chip"
        >1 bulan</button>
      </div>
    </div>

    <!-- TANGGAL -->
    <div class="mb-4">
      <label class="label">Tanggal Mulai</label>
      <input 
        type="date" 
        v-model="jadwal.form.tanggal_mulai" 
        class="input"
      />
    </div>

    <!-- REMINDER -->
    <div class="bg-slate-50 p-4 rounded-xl">
      <label class="label mb-2 block">Waktu Pengingat</label>

      <div class="grid grid-cols-1 sm:grid-cols-3 gap-3">
        <div v-if="jadwal.selectedWaktuMinum.includes('Pagi')" class="space-y-1">
          <span class="text-[10px] font-bold text-slate-400 uppercase tracking-wider block text-center sm:text-left">🌅 Pagi</span>
          <input type="time" v-model="jadwal.form.waktu_reminder_pagi" class="input text-center"/>
        </div>
        <div v-if="jadwal.selectedWaktuMinum.includes('Siang')" class="space-y-1">
          <span class="text-[10px] font-bold text-slate-400 uppercase tracking-wider block text-center sm:text-left">☀️ Siang</span>
          <input type="time" v-model="jadwal.form.waktu_reminder_siang" class="input text-center"/>
        </div>
        <div v-if="jadwal.selectedWaktuMinum.includes('Malam')" class="space-y-1">
          <span class="text-[10px] font-bold text-slate-400 uppercase tracking-wider block text-center sm:text-left">🌙 Malam</span>
          <input type="time" v-model="jadwal.form.waktu_reminder_malam" class="input text-center"/>
        </div>
      </div>
    </div>

    <!-- INFO BAWAH -->
    <div class="bg-teal-50 text-teal-600 text-xs p-3 rounded-lg mt-4">
      Pengingat akan dikirim 5 menit sebelum waktu minum
    </div>
  </div>

</div>

        <!-- STEP 3: SUCCESS -->
        <div v-if="jadwal.currentStep === 3" class="max-w-md mx-auto card text-center py-12 animate-in zoom-in duration-300">
          <div class="w-20 h-20 bg-green-100 text-green-600 rounded-full flex items-center justify-center mx-auto mb-6">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-10 w-10" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7" />
            </svg>
          </div>
          <h2 class="text-2xl font-bold text-slate-800 mb-2">Berhasil Disimpan!</h2>
          <p class="text-slate-500 mb-8">Jadwal konsumsi obat telah berhasil didaftarkan ke sistem.</p>
          <button @click="jadwal.cancelAdd" class="btn-primary w-full">Selesai</button>
        </div>
      </main>

      <!-- GLOBAL ACTION BUTTON -->
      <div v-if="jadwal.currentStep === 1 || jadwal.currentStep === 2"
          class="flex justify-between items-center mt-8">

        <!-- KIRI -->
        <button 
          @click="jadwal.currentStep === 1 ? jadwal.cancelAdd() : jadwal.currentStep--"
          class="px-5 py-2 rounded-lg bg-slate-100 hover:bg-slate-200 text-slate-600 text-sm font-medium"
        >
          ← {{ jadwal.currentStep === 1 ? 'Batal' : 'Kembali' }}
        </button>

        <!-- KANAN STEP 1 -->
        <button 
          v-if="jadwal.currentStep === 1"
          @click="jadwal.goToStep2()"
          class="px-6 py-2 rounded-lg bg-teal-500 hover:bg-teal-600 text-white text-sm font-semibold"
        >
          Lanjut Aturan Minum →
        </button>

        <!-- KANAN STEP 2 -->
        <button 
          v-if="jadwal.currentStep === 2"
          @click="handleSubmit"
          class="px-6 py-2 rounded-lg bg-teal-500 hover:bg-teal-600 text-white text-sm font-semibold"
        >
          Lihat Konfirmasi →
        </button>

      </div>

      <!-- LOADING OVERLAY -->
      <div v-if="jadwal.loading" class="overlay">
        <div class="bg-white p-6 rounded-2xl shadow-xl flex flex-col items-center">
          <div class="animate-spin rounded-full h-10 w-10 border-4 border-teal-500 border-t-transparent mb-4"></div>
          <p class="text-slate-600 font-medium">Memproses data...</p>
        </div>
      </div>

    </div>
  </LayoutWrapper>
</template>

<script setup>
import { onMounted, ref, computed, watch, onBeforeUnmount } from 'vue'
import { useRouter } from 'vue-router'
import LayoutWrapper from '../../components/LayoutWrapper.vue'
import { useJadwalStore } from '../../stores/jadwal'
import { usePasienStore } from '../../stores/pasien'
import { useObatStore } from '../../stores/obat'
import { useNotificationStore } from '../../stores/notification'
import { useConfirmStore } from '../../stores/confirm'

const router = useRouter()
const jadwal = useJadwalStore()
const pasien = usePasienStore()
const obat = useObatStore()
const notificationStore = useNotificationStore()
const confirmStore = useConfirmStore()
const reminderTimes = ref([])

const frekuensiCount = computed(() => {
  const f = jadwal.form.frekuensi || ''
  const match = f.match(/\d+/)
  return match ? parseInt(match[0]) : 0
})

const waktuLabels = ['Pagi', 'Siang', 'Malam']

watch(frekuensiCount, (val) => {
  reminderTimes.value = Array.from(
    { length: val },
    (_, i) => reminderTimes.value[i] || ''
  )
})

const openDropdownId = ref(null)
const menuRefs = ref({})

const toggleDropdown = (id) => {
  if (openDropdownId.value === id) {
    openDropdownId.value = null
  } else {
    openDropdownId.value = id
  }
}

const closeDropdown = () => {
  openDropdownId.value = null
}

const getDropdownStyle = (id) => {
  const el = menuRefs.value[id]
  if (!el) return {}
  const rect = el.getBoundingClientRect()
  return {
    top: rect.bottom + 4 + 'px',
    left: rect.right - 160 + 'px', // 160px is w-40 dropdown width
  }
}

const viewDetail = (id) => {
  router.push(`/jadwal/${id}`)
}

const editJadwal = (id) => {
  router.push(`/jadwal/${id}/edit`)
}

const deleteJadwal = async (id) => {
  const ok = await confirmStore.show({
    title: 'Hapus Jadwal Obat?',
    message: 'Apakah Anda yakin ingin menghapus jadwal obat ini? Tindakan ini tidak dapat dibatalkan.',
    confirmText: 'Ya, Hapus',
    cancelText: 'Batal',
    type: 'danger'
  })
  if (ok) {
    try {
      await jadwal.deleteJadwal(id)
      notificationStore.success('Jadwal berhasil dihapus.', 'Sukses')
    } catch (error) {
      notificationStore.error('Gagal menghapus jadwal.', 'Error')
    }
  }
}

const handleClickOutside = (e) => {
  const isInsideAnyMenu = Object.values(menuRefs.value).some(
    el => el && el.contains(e.target)
  )
  if (!isInsideAnyMenu) {
    openDropdownId.value = null
  }
}

onMounted(async () => {
  document.addEventListener('click', handleClickOutside)
  
  await Promise.all([
    jadwal.fetchJadwals(),
    pasien.fetchPasiens(),
    obat.fetchObats()
  ])

  if (!jadwal.form.jumlah_hari) {
    jadwal.form.jumlah_hari = 7
  }

  if (!jadwal.form.tipe_durasi) {
    jadwal.form.tipe_durasi = 'rutin'
  }
})

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside)
})

const handleSubmit = () => {
  if (!jadwal.form.tanggal_mulai) {
    notificationStore.error('Tanggal mulai wajib diisi', 'Gagal')
    return
  }
  if (jadwal.selectedWaktuMinum.length === 0) {
    notificationStore.warning('Pilih minimal satu waktu minum (Pagi/Siang/Malam)', 'Peringatan')
    return
  }
  jadwal.form.waktu_reminder = reminderTimes.value
  jadwal.form.tipe_durasi = jadwal.form.tipe_durasi || 'rutin'
  jadwal.submitJadwal()
}
</script>

<style scoped lang="postcss">
/* Layout & Cards */
.card {
  @apply bg-white p-6 rounded-[20px] shadow-[0_8px_30px_rgb(0,0,0,0.04)] border border-slate-100;
}

/* Typography */
.title {
  @apply text-lg font-bold text-slate-800 mb-5 flex items-center gap-2;
}
.label {
  @apply text-xs font-bold text-slate-500 mb-1.5 block;
}

/* Buttons */
.btn-primary {
  @apply bg-[#00a79d] hover:bg-[#008f86] text-white px-6 py-2.5 rounded-xl font-semibold transition-all shadow-md active:scale-[0.98];
}
.btn-secondary {
  @apply bg-slate-100 hover:bg-slate-200 text-slate-600 px-6 py-2.5 rounded-xl font-semibold transition-all;
}
.action-btn {
  @apply p-2 text-slate-400 hover:text-[#00a79d] border border-slate-100 rounded-lg hover:bg-white transition-all shadow-sm;
}

/* Inputs */
.search-input {
  @apply pl-10 pr-4 py-2.5 border border-slate-200 rounded-xl w-72 focus:ring-4 focus:ring-teal-500/10 focus:border-[#00a79d] outline-none transition-all bg-slate-50/50;
}
.input {
  @apply w-full border border-slate-200 p-3 rounded-xl focus:ring-4 focus:ring-teal-500/10 focus:border-[#00a79d] outline-none transition-all bg-slate-50/30 text-slate-700;
}

/* Table Styles */
.table-th {
  @apply px-6 py-4 text-[11px] font-bold text-slate-400 uppercase tracking-widest;
}
.table-td {
  @apply px-6 py-5 text-sm align-middle;
}

/* Selection Lists */
.list-container {
  @apply border border-slate-100 rounded-xl overflow-hidden max-h-60 overflow-y-auto;
}
.list-item {
  @apply p-4 cursor-pointer hover:bg-slate-50 border-b border-slate-50 last:border-none transition-colors text-slate-600;
}
.list-item.active {
  @apply bg-teal-50 text-teal-700 font-bold border-l-4 border-l-teal-500;
}

/* Chips */
.chip {
  @apply px-4 py-2 border border-slate-200 rounded-full cursor-pointer transition-all text-sm font-medium text-slate-500 bg-white;
}
.chip.active {
  @apply bg-[#00a79d] border-[#00a79d] text-white shadow-md shadow-teal-200;
}

/* Stepper */
.step-circle {
  @apply w-8 h-8 flex items-center justify-center rounded-full text-xs font-bold transition-all;
}
.step-circle.active {
  @apply bg-[#00a79d] text-white ring-4 ring-teal-100;
}
.step-circle.inactive {
  @apply bg-slate-200 text-slate-500;
}

/* Overlays */
.overlay {
  @apply fixed inset-0 bg-slate-900/40 backdrop-blur-sm flex items-center justify-center z-50;
}
</style>