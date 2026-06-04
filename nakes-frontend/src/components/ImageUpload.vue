<template>
  <div class="bg-blue-50 rounded-lg border border-blue-200 p-4">
    <label class="block text-sm font-medium text-gray-700 mb-4">
      Gambar Obat <span class="text-red-500">*</span>
    </label>

    <!-- Upload Area / Image Preview -->
    <div
      v-if="!previewImage"
      @drop.prevent="handleDrop"
      @dragover.prevent="isDragging = true"
      @dragleave="isDragging = false"
      :class="['border-2 border-dashed rounded-lg p-8 text-center transition cursor-pointer',
        isDragging ? 'border-blue-500 bg-blue-100' : 'border-blue-300 bg-blue-50 hover:bg-blue-100']"
    >
      <svg class="w-12 h-12 mx-auto text-blue-400 mb-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
      </svg>
      <p class="text-sm text-gray-600 font-medium mb-1">
        Seret dan lepas gambar di sini
      </p>
      <p class="text-xs text-gray-500 mb-4">
        atau gunakan tombol di bawah
      </p>

      <div class="flex flex-col sm:flex-row gap-2 justify-center">
        <!-- File Picker -->
        <button
          type="button"
          @click="triggerFileInput"
          class="px-4 py-2 bg-blue-500 text-white text-sm font-medium rounded-lg hover:bg-blue-600 transition flex items-center justify-center gap-2"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
          Pilih File
        </button>

        <!-- Camera Button -->
        <button
          type="button"
          @click="openCamera"
          class="px-4 py-2 bg-indigo-500 text-white text-sm font-medium rounded-lg hover:bg-indigo-600 transition flex items-center justify-center gap-2"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9a2 2 0 012-2h.93a2 2 0 001.664-.89l.812-1.22A2 2 0 0110.07 4h3.86a2 2 0 011.664.89l.812 1.22A2 2 0 0018.07 7H19a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2V9z"/>
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 13a3 3 0 11-6 0 3 3 0 016 0z"/>
          </svg>
          Ambil Foto
        </button>
      </div>

      <input
        ref="fileInput"
        type="file"
        accept="image/*"
        capture="environment"
        class="hidden"
        @change="handleFileSelect"
      />
    </div>

    <!-- Image Preview -->
    <div v-else class="relative">
      <img
        :src="previewImage"
        alt="Preview"
        class="w-full h-64 object-cover rounded-lg"
      />
      <div class="absolute top-2 right-2 flex gap-2">
        <!-- Delete Image Button -->
        <button
          type="button"
          @click="removeImage"
          class="p-2 bg-red-500 text-white rounded-lg hover:bg-red-600 transition shadow-lg"
          title="Hapus gambar"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
          </svg>
        </button>
      </div>

      <!-- File info -->
      <div class="mt-3 text-xs text-gray-600">
        <p v-if="uploadProgress > 0 && uploadProgress < 100" class="font-medium">
          Upload {{ uploadProgress }}%
        </p>
        <p v-if="fileName" class="font-medium">
          {{ fileName }}
        </p>
      </div>
    </div>

    <!-- Camera Modal -->
    <div
      v-if="showCameraModal"
      class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50 p-4"
    >
      <div class="bg-white rounded-xl shadow-xl w-full max-w-2xl overflow-hidden">
        <div class="flex justify-between items-center border-b border-gray-200 p-4 md:p-6">
          <h3 class="text-lg md:text-xl font-bold text-slate-900">Ambil Foto</h3>
          <button
            @click="closeCamera"
            class="text-gray-400 hover:text-gray-600 text-2xl font-bold"
          >
            ×
          </button>
        </div>

        <div class="p-4 md:p-6">
          <!-- Video stream -->
          <video
            v-if="!photoTaken"
            ref="videoElement"
            class="w-full rounded-lg bg-black"
            playsinline
            autoplay
            muted
          ></video>

          <!-- Captured photo -->
          <canvas
            v-else
            ref="canvasElement"
            class="w-full rounded-lg"
          ></canvas>
        </div>

        <div class="border-t border-gray-200 px-4 md:px-6 py-4 flex gap-2 justify-end flex-col-reverse md:flex-row">
          <button
            @click="closeCamera"
            class="w-full md:w-auto px-4 py-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50 transition font-medium text-sm md:text-base"
          >
            Batal
          </button>

          <button
            v-if="!photoTaken"
            @click="capturePhoto"
            class="w-full md:w-auto px-4 py-2 bg-teal-600 text-white rounded-lg hover:bg-teal-700 transition font-medium text-sm md:text-base"
          >
            Ambil Foto
          </button>

          <div v-else class="flex gap-2 w-full md:w-auto">
            <button
              @click="retakePhoto"
              class="flex-1 md:flex-none px-4 py-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50 transition font-medium text-sm md:text-base"
            >
              Ulang
            </button>
            <button
              @click="acceptPhoto"
              class="flex-1 md:flex-none px-4 py-2 bg-teal-600 text-white rounded-lg hover:bg-teal-700 transition font-medium text-sm md:text-base"
            >
              Gunakan Foto
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, nextTick, onBeforeUnmount } from 'vue'
import apiClient from '../services/api'

export default {
  name: 'ImageUpload',
  props: {
    initialImage: { type: String, default: null }
  },
  emits: ['update:image'],
  setup(props, { emit }) {
    const previewImage = ref(props.initialImage)
    const fileName = ref('')
    const isDragging = ref(false)
    const uploadProgress = ref(0)
    const fileInput = ref(null)

    // Camera related
    const showCameraModal = ref(false)
    const videoElement = ref(null)
    const canvasElement = ref(null)
    const offscreenCanvas = ref(null) // Offscreen canvas untuk capture
    const stream = ref(null)
    const photoTaken = ref(false)

    // Handle file selection
    const handleFileSelect = async (event) => {
      const file = event.target.files?.[0]
      if (file) {
        await handleFile(file)
      }
      // Reset input so same file can be selected again
      event.target.value = ''
    }

    // Handle drop
    const handleDrop = async (event) => {
      isDragging.value = false
      const files = event.dataTransfer.files
      if (files.length > 0) {
        await handleFile(files[0])
      }
    }

    // Main file handler
    const handleFile = async (file) => {
      // Validate file type
      if (!file.type.startsWith('image/')) {
        alert('Mohon pilih file gambar')
        return
      }

      // Validate file size (max 5MB)
      const maxSize = 5 * 1024 * 1024
      if (file.size > maxSize) {
        alert('Ukuran file maksimal 5MB')
        return
      }

      // Show preview
      const reader = new FileReader()
      reader.onload = (e) => {
        previewImage.value = e.target.result
        fileName.value = file.name
      }
      reader.readAsDataURL(file)

      // Upload file
      await uploadFile(file)
    }

    // Upload file to server
    const uploadFile = async (file) => {
      try {
        uploadProgress.value = 0
        const formData = new FormData()
        formData.append('image', file)

        const response = await apiClient.post('/upload/image', formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          },
          onUploadProgress: (progressEvent) => {
            uploadProgress.value = Math.round(
              (progressEvent.loaded * 100) / progressEvent.total
            )
          }
        })

        // Emit the uploaded file path
        if (response.data.data) {
          emit('update:image', response.data.data.url)
        }

        // Reset progress after upload
        setTimeout(() => {
          uploadProgress.value = 0
        }, 1000)
      } catch (error) {
        console.error('Upload error:', error)
        alert('Gagal mengupload gambar: ' + (error.response?.data?.message || error.message))
        previewImage.value = null
        fileName.value = ''
      }
    }

    // Camera functions
    const openCamera = async () => {
      // Check browser support
      if (!navigator.mediaDevices || !navigator.mediaDevices.getUserMedia) {
        alert('Browser Anda tidak mendukung akses kamera')
        return
      }

      showCameraModal.value = true
      // Tunggu DOM ter-render
      await nextTick()
      await startCamera()
    }

    const closeCamera = () => {
      stopCamera()
      offscreenCanvas.value = null
      videoElement.value = null
      canvasElement.value = null
      showCameraModal.value = false
      photoTaken.value = false
      console.log('Camera modal closed')
    }

    const startCamera = async () => {
      try {
        // Stop previous stream completely
        if (stream.value) {
          console.log('Stopping previous stream...')
          stream.value.getTracks().forEach(track => {
            track.stop()
          })
          stream.value = null
        }

        console.log('Requesting camera access...')
        stream.value = await navigator.mediaDevices.getUserMedia({
          video: { 
            facingMode: 'environment',
            width: { ideal: 1280 },
            height: { ideal: 720 }
          },
          audio: false
        })

        console.log('Stream obtained, waiting for video element...')
        
        // Tunggu video element ready
        let retries = 0
        while (!videoElement.value && retries < 10) {
          await new Promise(resolve => setTimeout(resolve, 50))
          retries++
        }
        
        if (!videoElement.value) {
          console.error('Video element tidak ditemukan')
          throw new Error('Video element tidak ditemukan')
        }

        console.log('Video element ready, setting stream...')
        videoElement.value.srcObject = stream.value
        
        // Tunggu metadata
        await new Promise((resolve, reject) => {
          const timeout = setTimeout(() => {
            reject(new Error('Timeout waiting for video metadata'))
          }, 3000)
          
          const handleMetadata = () => {
            clearTimeout(timeout)
            videoElement.value.removeEventListener('loadedmetadata', handleMetadata)
            resolve()
          }
          
          videoElement.value.addEventListener('loadedmetadata', handleMetadata, { once: true })
        })

        console.log('Video metadata loaded, playing...')
        const playPromise = videoElement.value.play()
        if (playPromise !== undefined) {
          await playPromise
        }
        
        console.log('Camera started successfully')
      } catch (error) {
        console.error('Camera error:', error)
        alert('Tidak dapat mengakses kamera: ' + error.message)
        closeCamera()
      }
    }

    const stopCamera = () => {
      console.log('Stopping camera...')
      if (stream.value) {
        stream.value.getTracks().forEach(track => {
          console.log('Stopping track:', track.kind)
          track.stop()
        })
        stream.value = null
      }
      
      if (videoElement.value) {
        videoElement.value.srcObject = null
      }
      
      console.log('Camera stopped')
    }

    const capturePhoto = async () => {
      // Ensure video element is ready
      await nextTick()
      
      console.log('capturePhoto called')
      console.log('videoElement:', videoElement.value)
      
      if (!videoElement.value) {
        console.error('Video element tidak ditemukan')
        alert('Video element tidak ditemukan')
        return
      }

      // Pastikan video sudah siap
      if (videoElement.value.videoWidth === 0 || videoElement.value.videoHeight === 0) {
        console.warn('Video belum siap, dimensions:', videoElement.value.videoWidth, 'x', videoElement.value.videoHeight)
        alert('Video belum siap, coba lagi dalam beberapa saat')
        return
      }

      try {
        // Gunakan offscreen canvas (tidak perlu menunggu DOM render)
        const canvas = document.createElement('canvas')
        canvas.width = videoElement.value.videoWidth
        canvas.height = videoElement.value.videoHeight
        
        const context = canvas.getContext('2d')
        context.drawImage(
          videoElement.value,
          0,
          0,
          videoElement.value.videoWidth,
          videoElement.value.videoHeight
        )
        
        // Simpan reference ke offscreen canvas
        offscreenCanvas.value = canvas
        
        console.log('Photo captured successfully, canvas:', offscreenCanvas.value)
        photoTaken.value = true
        
        // Update template canvas setelah render
        await nextTick()
        if (canvasElement.value && offscreenCanvas.value) {
          const templateContext = canvasElement.value.getContext('2d')
          canvasElement.value.width = offscreenCanvas.value.width
          canvasElement.value.height = offscreenCanvas.value.height
          templateContext.drawImage(offscreenCanvas.value, 0, 0)
        }
      } catch (error) {
        console.error('Capture photo error:', error)
        alert('Gagal mengambil foto: ' + error.message)
      }
    }

    const retakePhoto = async () => {
      try {
        console.log('retakePhoto called')
        console.log('Current stream state:', stream.value ? 'active' : 'inactive')
        
        // Clear offscreen canvas
        offscreenCanvas.value = null
        
        // Step 1: Reset state untuk re-render DOM
        photoTaken.value = false
        console.log('photoTaken set to false')
        
        // Step 2: Tunggu DOM render kembali ke video element
        await nextTick()
        console.log('nextTick done, video element should be in DOM now')
        
        // Step 3: Tunggu video element benar-benar ada di DOM
        let retries = 0
        while (!videoElement.value && retries < 10) {
          await new Promise(resolve => setTimeout(resolve, 50))
          retries++
        }
        
        if (!videoElement.value) {
          console.error('Video element tidak ditemukan setelah retries')
          throw new Error('Video element tidak ditemukan')
        }
        
        console.log('Video element found in DOM')
        
        // Step 4: Reconnect stream ke video element yang baru
        if (stream.value) {
          console.log('Connecting stream to video element...')
          videoElement.value.srcObject = stream.value
          
          // Step 5: Tunggu metadata
          let metadataLoaded = false
          const metadataPromise = new Promise((resolve, reject) => {
            const timeout = setTimeout(() => {
              reject(new Error('Timeout waiting for video metadata'))
            }, 3000)
            
            const handleMetadata = () => {
              clearTimeout(timeout)
              videoElement.value.removeEventListener('loadedmetadata', handleMetadata)
              metadataLoaded = true
              resolve()
            }
            
            videoElement.value.addEventListener('loadedmetadata', handleMetadata, { once: true })
          })
          
          await metadataPromise
          console.log('Video metadata loaded')
          
          // Step 6: Play video
          const playPromise = videoElement.value.play()
          if (playPromise !== undefined) {
            await playPromise
          }
          
          console.log('Video playing successfully')
        } else {
          // Stream hilang, restart camera
          console.log('Stream tidak aktif, restarting camera...')
          await startCamera()
        }
      } catch (error) {
        console.error('Retake photo error:', error)
        console.error('Error message:', error.message)
        
        // Restart camera jika ada error
        try {
          console.log('Attempting to restart camera...')
          await startCamera()
        } catch (restartError) {
          console.error('Failed to restart camera:', restartError)
          alert('Gagal mengulang foto: ' + restartError.message)
          closeCamera()
        }
      }
    }

    const acceptPhoto = async () => {
      // Gunakan offscreen canvas atau template canvas
      const sourceCanvas = offscreenCanvas.value || canvasElement.value
      
      if (!sourceCanvas) {
        alert('Canvas tidak ditemukan')
        return
      }

      try {
        const blob = await new Promise((resolve, reject) => {
          sourceCanvas.toBlob(
            (blob) => {
              if (blob) {
                resolve(blob)
              } else {
                reject(new Error('Gagal mengonversi canvas ke blob'))
              }
            },
            'image/jpeg',
            0.95
          )
        })

        // Generate filename
        const filename = `camera_${Date.now()}.jpg`
        const file = new File([blob], filename, { type: 'image/jpeg' })

        // Create preview
        const reader = new FileReader()
        reader.onload = (e) => {
          previewImage.value = e.target.result
          fileName.value = filename
        }
        reader.readAsDataURL(file)

        // Upload file
        await uploadFile(file)

        // Clear offscreen canvas
        offscreenCanvas.value = null

        // Close camera
        closeCamera()
      } catch (error) {
        console.error('Accept photo error:', error)
        alert('Gagal memproses foto: ' + error.message)
      }
    }

    // Remove image
    const removeImage = () => {
      previewImage.value = null
      fileName.value = ''
      emit('update:image', null)
    }

    // Trigger file input safely
    const triggerFileInput = () => {
      if (fileInput.value) {
        fileInput.value.click()
      } else {
        console.warn('File input ref tidak ditemukan')
      }
    }

    // Cleanup when component unmounts
    onBeforeUnmount(() => {
      console.log('Component unmounting, cleaning up...')
      stopCamera()
      offscreenCanvas.value = null
      videoElement.value = null
      canvasElement.value = null
    })

    return {
      previewImage,
      fileName,
      isDragging,
      uploadProgress,
      fileInput,
      showCameraModal,
      videoElement,
      canvasElement,
      photoTaken,
      handleFileSelect,
      handleDrop,
      openCamera,
      closeCamera,
      capturePhoto,
      retakePhoto,
      acceptPhoto,
      removeImage,
      triggerFileInput
    }
  }
}
</script>
