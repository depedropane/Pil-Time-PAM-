<template>
  <div class="panduan-page">

    <!-- NAVBAR -->
    <NavbarView />

    <!-- HERO -->
    <section class="hero">

      <div class="hero-orb hero-orb-1"></div>
      <div class="hero-orb hero-orb-2"></div>

      <div class="hero-content">

        <span class="badge">
          Panduan Penggunaan
        </span>

        <h1>
          Pelajari Cara Menggunakan
          <span>Pil Time</span>
        </h1>

        <p>
          Ikuti langkah-langkah mudah dan dokumentasi penggunaan
          untuk memahami seluruh fitur Pil Time — cocok untuk pasien
          maupun tenaga kesehatan.
        </p>

      </div>

    </section>


    <!-- PDF SECTION -->
    <section class="pdf-section">

      <div class="section-header">
        <span class="section-tag">DOKUMENTASI PDF</span>
        <h2 class="section-title">Panduan Lengkap Penggunaan Aplikasi</h2>
        <p class="section-subtitle">Unduh atau baca langsung panduan penggunaan sesuai peran Anda</p>
      </div>

      <!-- PANDUAN NAKES -->
      <div class="pdf-box">

        <div class="pdf-header">
          <div class="pdf-header-icon">
            <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" style="width:28px;height:28px"><path d="M17 21v-2a4 4 0 00-4-4H5a4 4 0 00-4 4v2" stroke="#1D9E75" stroke-width="1.5" stroke-linecap="round"/><circle cx="9" cy="7" r="4" stroke="#1D9E75" stroke-width="1.5"/><path d="M23 21v-2a4 4 0 00-3-3.87M16 3.13a4 4 0 010 7.75" stroke="#1D9E75" stroke-width="1.5" stroke-linecap="round"/></svg>
          </div>
          <div>
            <h3 class="pdf-title">Panduan Tenaga Kesehatan</h3>
            <p class="pdf-subtitle">Untuk dokter, perawat, dan tenaga medis yang mengelola pasien melalui dashboard nakes</p>
          </div>
        </div>

        <div class="pdf-container">
          <iframe
            :src="panduanNakes"
            width="100%"
            height="800px"
            class="pdf-viewer"
            title="Panduan Tenaga Kesehatan Pil Time"
          >
          </iframe>
        </div>

        <div class="download-box">
          <a
            :href="panduanNakes"
            download
            class="download-btn"
            id="btn-download-nakes"
          >
            Download Panduan Tenaga Kesehatan
          </a>
        </div>

      </div>

      <!-- PANDUAN PASIEN -->
      <div class="pdf-box">

        <div class="pdf-header">
          <div class="pdf-header-icon">
            <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" style="width:28px;height:28px"><path d="M17 21v-2a4 4 0 00-4-4H5a4 4 0 00-4 4v2" stroke="#1D9E75" stroke-width="1.5" stroke-linecap="round"/><circle cx="9" cy="7" r="4" stroke="#1D9E75" stroke-width="1.5"/></svg>
          </div>
          <div>
            <h3 class="pdf-title">Panduan Pasien</h3>
            <p class="pdf-subtitle">Untuk pasien dan keluarga yang menggunakan aplikasi Pil Time di smartphone Android</p>
          </div>
        </div>

        <div class="pdf-container">
          <iframe
            :src="panduanPasien"
            width="100%"
            height="800px"
            class="pdf-viewer"
            title="Panduan Pasien Pil Time"
          >
          </iframe>
        </div>

        <div class="download-box">
          <a
            :href="panduanPasien"
            download
            class="download-btn"
            id="btn-download-pasien"
          >
            Download Panduan Pasien
          </a>
        </div>

      </div>

    </section>

    <!-- CTA BOTTOM -->
    <section class="cta-section">
      <div class="cta-inner">
        <h2>Siap Mulai Menggunakan Pil Time?</h2>
        <p>Unduh aplikasinya sekarang dan mulai jaga kepatuhan minum obat Anda bersama Pil Time.</p>
        <a href="/install" class="cta-btn" id="btn-install-from-panduan">
          Download Pil Time Gratis
        </a>
      </div>
    </section>

    <!-- FAQ SECTION -->
    <section class="faq-section">

      <div class="section-header">
        <span class="section-tag">FAQ</span>
        <h2 class="section-title">Pertanyaan yang Sering Ditanyakan</h2>
      </div>

      <div class="faq-list">

        <div
          class="faq-item"
          v-for="(item, index) in faqs"
          :key="index"
          :class="{ open: openFaq === index }"
          @click="toggleFaq(index)"
        >
          <div class="faq-question">
            <span>{{ item.q }}</span>
            <div class="faq-icon">{{ openFaq === index ? '−' : '+' }}</div>
          </div>
          <div class="faq-answer" v-show="openFaq === index">
            <p>{{ item.a }}</p>
          </div>
        </div>

      </div>

    </section>

    <FooterView />

  </div>
</template>

<script>
import NavbarView from './NavbarView.vue'
import FooterView from './FooterView.vue'

// IMPORT PDF
import panduanNakes from './../../assets/pdf/panduannakes.pdf'
import panduanPasien from './../../assets/pdf/panduanpasien.pdf'

export default {
  name: "PanduanView",

  components: {
    NavbarView,
    FooterView
  },

  data() {
    return {
      panduanNakes,
      panduanPasien,
      openFaq: null,
      faqs: [
        {
          q: "Apakah Pil Time tersedia untuk iOS?",
          a: "Saat ini Pil Time tersedia sebagai aplikasi Android (APK). Versi iOS sedang dalam tahap pengembangan dan akan segera tersedia."
        },
        {
          q: "Apakah aplikasi Pil Time gratis digunakan?",
          a: "Ya, Pil Time sepenuhnya gratis untuk diunduh dan digunakan oleh pasien. Tidak ada biaya langganan atau pembelian dalam aplikasi."
        },
        {
          q: "Bagaimana cara tenaga kesehatan mendaftar dan mengelola pasien?",
          a: "Tenaga kesehatan dapat login melalui web dashboard di halaman Login. Setelah masuk, Anda dapat menambahkan pasien, mengatur jadwal obat, dan memantau riwayat kepatuhan mereka secara real-time."
        },
        {
          q: "Apakah alarm akan berbunyi jika ponsel dalam mode senyap?",
          a: "Alarm Pil Time dirancang untuk menembus mode senyap pada kondisi tertentu. Namun, kami menyarankan untuk mengizinkan notifikasi prioritas tinggi dari aplikasi agar alarm selalu berbunyi tepat waktu."
        },
        {
          q: "Bagaimana jika saya lupa mengkonfirmasi bahwa saya sudah minum obat?",
          a: "Aplikasi akan mencatat dosis sebagai 'Terlewat' jika tidak dikonfirmasi dalam jangka waktu tertentu setelah alarm. Riwayat ini dapat dilihat oleh tenaga kesehatan untuk evaluasi pengobatan."
        },
        {
          q: "Bisakah lebih dari satu obat diatur jadwalnya sekaligus?",
          a: "Ya, Pil Time mendukung pengaturan jadwal untuk beberapa obat sekaligus dengan waktu dan dosis yang berbeda-beda, sesuai resep dari tenaga kesehatan."
        }
      ]
    }
  },

  methods: {
    toggleFaq(index) {
      this.openFaq = this.openFaq === index ? null : index
    }
  }
}
</script>

<style scoped>

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.panduan-page {
  font-family: "Plus Jakarta Sans", sans-serif;
  background: #ffffff;
  color: #1e293b;
}

/* =========================
   HERO
========================= */

.hero {
  position: relative;
  min-height: 70vh;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  padding: 140px 20px 100px;
  overflow: hidden;

  background:
    radial-gradient(circle at top left, #d4f5e8 0%, transparent 35%),
    radial-gradient(circle at bottom right, #b8edda 0%, transparent 40%),
    linear-gradient(180deg, #f8fffe 0%, #f0fdf8 100%);
}

.hero-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(70px);
  pointer-events: none;
}

.hero-orb-1 {
  width: 350px;
  height: 350px;
  background: rgba(29, 158, 117, 0.09);
  top: -80px;
  left: -80px;
}

.hero-orb-2 {
  width: 250px;
  height: 250px;
  background: rgba(29, 158, 117, 0.07);
  bottom: -60px;
  right: -40px;
}

.hero-content {
  position: relative;
  z-index: 1;
  max-width: 850px;
  animation: fadeInUp 0.7s ease both;
}

@keyframes fadeInUp {
  from { opacity: 0; transform: translateY(24px); }
  to   { opacity: 1; transform: translateY(0); }
}

.badge {
  display: inline-block;
  background: #e9faf3;
  border: 1.5px solid rgba(29,158,117,0.2);
  color: #1D9E75;
  padding: 10px 20px;
  border-radius: 999px;
  font-size: 14px;
  font-weight: 700;
  margin-bottom: 24px;
}

.hero-content h1 {
  font-size: 58px;
  line-height: 1.2;
  margin-bottom: 22px;
  font-weight: 900;
  color: #0f172a;
}

.hero-content h1 span {
  color: #1D9E75;
}

.hero-content p {
  font-size: 18px;
  line-height: 1.9;
  color: #374151;
  max-width: 640px;
  margin: 0 auto;
}

/* =========================
   SECTION HEADER (shared)
========================= */

.section-header {
  text-align: center;
  margin-bottom: 60px;
}

.section-tag {
  display: inline-block;
  color: #1D9E75;
  font-size: 12px;
  font-weight: 800;
  letter-spacing: 2.5px;
  padding: 6px 16px;
  border-radius: 999px;
  background: #e9faf3;
  margin-bottom: 14px;
}

.section-title {
  font-size: 40px;
  font-weight: 900;
  color: #0f172a;
  margin-bottom: 14px;
  line-height: 1.25;
}

.section-subtitle {
  font-size: 17px;
  color: #6b7280;
  max-width: 560px;
  margin: 0 auto;
  line-height: 1.75;
}

/* =========================
   TUTORIAL SECTION
========================= */

.tutorial-section {
  padding: 100px 8%;
  background: #ffffff;
}

/* VIDEO */
.video-container {
  overflow: hidden;
  border-radius: 28px;
  box-shadow: 0 20px 50px rgba(0,0,0,0.08);
  margin-bottom: 60px;
  border: 1.5px solid rgba(29,158,117,0.1);
}

.video-placeholder {
  width: 100%;
  height: 420px;
  background: linear-gradient(135deg, #f0fdf8, #e9faf3);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 20px;
  text-align: center;
  padding: 40px;
}

.vp-icon {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  background: linear-gradient(135deg, #1D9E75, #0F6E56);
  color: white;
  font-size: 28px;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 12px 32px rgba(29,158,117,0.3);
}

.vp-text strong {
  display: block;
  font-size: 20px;
  font-weight: 800;
  color: #0f172a;
  margin-bottom: 8px;
}

.vp-text span {
  font-size: 15px;
  color: #6b7280;
  line-height: 1.7;
  max-width: 480px;
  display: block;
}

/* STEP CARDS */
.tutorial-description {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 24px;
}

.tutorial-card {
  background: #f8fffe;
  border: 1.5px solid rgba(29,158,117,0.09);
  padding: 28px;
  border-radius: 24px;
  display: flex;
  gap: 20px;
  align-items: flex-start;
  transition: all 0.35s ease;
  box-shadow: 0 8px 22px rgba(0,0,0,0.04);
}

.tutorial-card:hover {
  transform: translateY(-6px);
  box-shadow: 0 20px 44px rgba(29,158,117,0.12);
  border-color: rgba(29,158,117,0.2);
}

.step-num {
  min-width: 52px;
  height: 52px;
  border-radius: 16px;
  background: linear-gradient(135deg, #1D9E75, #0F6E56);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 800;
  font-size: 20px;
  flex-shrink: 0;
  box-shadow: 0 6px 16px rgba(29,158,117,0.3);
}

.step-content h3 {
  font-size: 17px;
  font-weight: 800;
  color: #0f172a;
  margin-bottom: 10px;
}

.step-content p {
  color: #6b7280;
  line-height: 1.8;
  font-size: 14.5px;
}

.step-content a {
  color: #1D9E75;
  font-weight: 700;
  text-decoration: none;
}

.step-content a:hover {
  text-decoration: underline;
}

/* =========================
   FAQ SECTION
========================= */

.faq-section {
  padding: 100px 8%;
  background: #f8fffe;
}

.faq-list {
  max-width: 820px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.faq-item {
  background: white;
  border-radius: 18px;
  border: 1.5px solid rgba(29,158,117,0.1);
  overflow: hidden;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 16px rgba(0,0,0,0.04);
}

.faq-item:hover {
  border-color: rgba(29,158,117,0.25);
  box-shadow: 0 8px 24px rgba(29,158,117,0.08);
}

.faq-item.open {
  border-color: rgba(29,158,117,0.3);
  box-shadow: 0 10px 28px rgba(29,158,117,0.1);
}

.faq-question {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  padding: 22px 26px;
}

.faq-question span {
  font-size: 16px;
  font-weight: 700;
  color: #0f172a;
  line-height: 1.5;
}

.faq-icon {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background: #e9faf3;
  color: #1D9E75;
  font-size: 22px;
  font-weight: 700;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  transition: all 0.3s ease;
}

.faq-item.open .faq-icon {
  background: #1D9E75;
  color: white;
}

.faq-answer {
  padding: 0 26px 22px;
}

.faq-answer p {
  color: #6b7280;
  line-height: 1.85;
  font-size: 15px;
  padding-top: 4px;
  border-top: 1px solid rgba(29,158,117,0.1);
  padding-top: 16px;
}

/* =========================
   PDF SECTION
========================= */

.pdf-section {
  padding: 100px 8%;
  background: #ffffff;
}

.pdf-box {
  margin-bottom: 80px;
  background: #f8fffe;
  border-radius: 30px;
  padding: 36px;
  border: 1.5px solid rgba(29,158,117,0.1);
}

.pdf-header {
  display: flex;
  align-items: center;
  gap: 20px;
  margin-bottom: 28px;
}

.pdf-header-icon {
  font-size: 42px;
  width: 72px;
  height: 72px;
  border-radius: 22px;
  background: #e9faf3;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.pdf-title {
  font-size: 24px;
  font-weight: 800;
  color: #0f172a;
  margin-bottom: 6px;
}

.pdf-subtitle {
  font-size: 14.5px;
  color: #6b7280;
  line-height: 1.7;
}

.pdf-container {
  overflow: hidden;
  border-radius: 20px;
  background: white;
  padding: 12px;
  box-shadow: 0 12px 36px rgba(0,0,0,0.06);
}

.pdf-viewer {
  border: none;
  border-radius: 14px;
  display: block;
}

.download-box {
  margin-top: 32px;
  text-align: center;
}

.download-btn {
  display: inline-flex;
  align-items: center;
  gap: 10px;

  background: linear-gradient(135deg, #1D9E75, #0F6E56);
  color: white;
  text-decoration: none;

  padding: 16px 36px;
  border-radius: 14px;

  font-weight: 700;
  font-size: 15px;

  transition: all 0.3s ease;
  box-shadow: 0 8px 24px rgba(29,158,117,0.28);
}

.download-btn:hover {
  transform: translateY(-4px);
  box-shadow: 0 16px 36px rgba(29,158,117,0.35);
}

/* =========================
   CTA BOTTOM
========================= */

.cta-section {
  padding: 90px 8%;
  background: linear-gradient(135deg, #1D9E75 0%, #0F6E56 60%, #0a4d3c 100%);
  text-align: center;
}

.cta-inner {
  max-width: 620px;
  margin: 0 auto;
}

.cta-emoji {
  font-size: 48px;
  margin-bottom: 20px;
  display: block;
}

.cta-section h2 {
  font-size: 36px;
  font-weight: 900;
  color: white;
  margin-bottom: 16px;
  line-height: 1.25;
}

.cta-section p {
  font-size: 16px;
  color: rgba(255,255,255,0.85);
  line-height: 1.8;
  margin-bottom: 32px;
}

.cta-btn {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  background: white;
  color: #1D9E75;
  text-decoration: none;
  padding: 17px 36px;
  border-radius: 14px;
  font-size: 16px;
  font-weight: 800;
  transition: all 0.3s ease;
  box-shadow: 0 8px 24px rgba(0,0,0,0.15);
}

.cta-btn:hover {
  transform: translateY(-4px);
  box-shadow: 0 16px 36px rgba(0,0,0,0.2);
  background: #f0fdf8;
}

/* =========================
   RESPONSIVE
========================= */

@media (max-width: 768px) {

  .hero-content h1 {
    font-size: 38px;
  }

  .section-title {
    font-size: 30px;
  }

  .tutorial-card {
    flex-direction: column;
    gap: 14px;
  }

  .pdf-box {
    padding: 22px;
  }

  .pdf-header {
    flex-direction: column;
    align-items: flex-start;
  }

  iframe {
    height: 350px;
  }

  .cta-section h2 {
    font-size: 28px;
  }

}

@media (max-width: 480px) {
  .hero-content h1 {
    font-size: 32px;
  }

  .section-title {
    font-size: 26px;
  }

  .vp-text strong {
    font-size: 18px;
  }

  iframe {
    height: 300px;
  }

  .cta-section h2 {
    font-size: 24px;
  }
}

</style>