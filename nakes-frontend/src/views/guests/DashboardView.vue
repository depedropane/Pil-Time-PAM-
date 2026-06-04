<script setup>
import { useRouter } from "vue-router";
import NavbarView from "./NavbarView.vue";
import FooterView from "./FooterView.vue";

const router = useRouter();

const goToInstall = () => {
  router.push("/install");
};

const goToPanduan = () => {
  router.push("/panduan");
};
</script>

<template>

  <!-- REUSABLE NAVBAR -->
  <NavbarView />

  <!-- HERO -->
  <section class="landing">

    <!-- DECORATIVE ORBS -->
    <div class="orb orb-1"></div>
    <div class="orb orb-2"></div>
    <div class="orb orb-3"></div>

    <!-- LEFT -->
    <div class="hero-left">

      <div class="badge">
        Aplikasi Pengingat Obat Modern
      </div>

      <h1>
        Jangan Pernah
        <br />
        Lupa Minum
        <br />
        Obat Bersama
        <span class="green">PilTime</span>
      </h1>

      <p>
        Pengingat minum obat yang <strong>hangat</strong>,
        sederhana, dan mudah digunakan
        untuk membantu lansia menjaga
        kesehatan setiap hari.
      </p>

      <div class="hero-buttons">

        <button
          class="primary-btn"
          @click="goToInstall"
          id="btn-download-hero"
        >
          Download Gratis
        </button>

        <button
          class="secondary-btn"
          @click="goToPanduan"
          id="btn-demo-hero"
        >
          Lihat Demo
        </button>

      </div>

      <!-- TRUST BADGES -->
      <div class="trust-row">
        <div class="trust-item">
          <span class="trust-num">5.000+</span>
          <span class="trust-label">Pengguna Aktif</span>
        </div>
        <div class="trust-divider"></div>
        <div class="trust-item">
          <span class="trust-num">98%</span>
          <span class="trust-label">Kepatuhan Meningkat</span>
        </div>
        <div class="trust-divider"></div>
        <div class="trust-item">
          <span class="trust-num">100%</span>
          <span class="trust-label">Gratis & Aman</span>
        </div>
      </div>

    </div>

    <!-- RIGHT -->
    <div class="hero-right">

      <div class="phone-scene">

        <div class="circle-bg"></div>

        <!-- FLOATING BADGES -->
        <div class="float-badge fb-top">
          Obat diminum tepat waktu
        </div>
        <div class="float-badge fb-bottom">
          Reminder aktif 3 obat
        </div>

        <div class="phone-wrapper">

          <div class="phone-shadow"></div>

          <div class="phone">

            <div class="notch"></div>

            <div class="phone-screen">

              <!-- TOP -->
              <div class="phone-top">

                <div>
                  <div class="hello">
                    Selamat Siang
                  </div>

                  <div class="username">
                    Hikzha Cristian
                  </div>
                </div>

                <div class="notif">
                  <svg viewBox="0 0 20 20" fill="none" style="width:18px;height:18px" xmlns="http://www.w3.org/2000/svg"><path d="M10 2a6 6 0 00-6 6v1.17l-.7 1.42A1 1 0 004.2 12h11.6a1 1 0 00.9-1.41L16 9.17V8a6 6 0 00-6-6zM8 14a2 2 0 004 0H8z" fill="#1D9E75"/></svg>
                </div>

              </div>

              <!-- CARD -->
              <div class="hero-card">

                <div class="hero-label">
                  KEPATUHAN HARI INI
                </div>

                <div class="hero-value">
                  1 dari 2 Selesai
                </div>

                <div class="hero-progress-bar">
                  <div class="hero-progress-fill"></div>
                </div>

                <div class="hero-desc">
                  Ayo selesaikan sesi malam ya!
                </div>

              </div>

              <!-- MENU -->
              <div class="menu-title">
                Menu Utama
              </div>

              <div class="menu-grid">

                <div class="menu-item">
                  <div class="menu-icon">
                    <svg viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M10 2a6 6 0 00-6 6v1.17l-.7 1.42A1 1 0 004.2 12h11.6a1 1 0 00.9-1.41L16 9.17V8a6 6 0 00-6-6zM8 14a2 2 0 004 0H8z" fill="#1D9E75"/></svg>
                  </div>
                  <div class="menu-text">Reminder</div>
                </div>

                <div class="menu-item">
                  <div class="menu-icon">
                    <svg viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg"><rect x="4" y="3" width="12" height="15" rx="2" stroke="#1D9E75" stroke-width="1.5"/><path d="M7 8h6M7 11h4" stroke="#1D9E75" stroke-width="1.5" stroke-linecap="round"/></svg>
                  </div>
                  <div class="menu-text">Info Obat</div>
                </div>

                <div class="menu-item">
                  <div class="menu-icon">
                    <svg viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg"><polyline points="3,15 7,9 11,12 15,5" stroke="#1D9E75" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg>
                  </div>
                  <div class="menu-text">Riwayat</div>
                </div>

                <div class="menu-item">
                  <div class="menu-icon">
                    <svg viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M10 17s-7-4.35-7-9a5 5 0 0110 0" stroke="#1D9E75" stroke-width="1.5" stroke-linecap="round"/><path d="M13 8a3 3 0 010 6" stroke="#1D9E75" stroke-width="1.5" stroke-linecap="round"/></svg>
                  </div>
                  <div class="menu-text">Rutinitas</div>
                </div>

              </div>

              <div class="bottom-nav">

                <div class="nav-item active">
                  <svg viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg" style="width:16px;height:16px"><path d="M3 9.5L10 3l7 6.5V17a1 1 0 01-1 1H4a1 1 0 01-1-1V9.5z" stroke="#1D9E75" stroke-width="1.5" stroke-linejoin="round"/></svg>
                </div>

                <div class="add-btn">
                  +
                </div>

                <div class="nav-item">
                  <svg viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg" style="width:16px;height:16px"><circle cx="10" cy="7" r="3" stroke="#64748b" stroke-width="1.5"/><path d="M4 17c0-3.314 2.686-6 6-6s6 2.686 6 6" stroke="#64748b" stroke-width="1.5" stroke-linecap="round"/></svg>
                </div>

              </div>

            </div>
          </div>
        </div>

      </div>

    </div>

  </section>

  <!-- FEATURES — REDESIGNED GRID -->
  <section class="feature-section">

    <div class="feature-header">
      <span class="section-tag">FITUR UNGGULAN</span>
      <h2 class="section-title">Semua Yang Anda Butuhkan<br/>Ada di PilTime</h2>
      <p class="section-subtitle">Dirancang khusus agar mudah digunakan oleh semua kalangan, termasuk lansia</p>
    </div>

    <div class="feature-grid">

      <div class="feature-card fc-green">
        <div class="fc-icon-wrap">
          <svg class="fc-svg" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M12 3a7 7 0 00-7 7v1.5l-.9 1.8A1 1 0 005 15h14a1 1 0 00.9-1.7L19 11.5V10a7 7 0 00-7-7zM9 17a3 3 0 006 0H9z" fill="#1D9E75"/></svg>
        </div>
        <div class="fc-number">01</div>
        <h3 class="fc-title">Reminder & Alarm</h3>
        <p class="fc-desc">
          Alarm pintar otomatis setiap jadwal minum obat — pagi, siang, dan malam. Tidak akan terlewat satu dosis pun.
        </p>
        <div class="fc-arrow">→</div>
      </div>

      <div class="feature-card fc-teal">
        <div class="fc-icon-wrap">
          <svg class="fc-svg" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg"><rect x="5" y="3" width="14" height="18" rx="2" stroke="#0F6E56" stroke-width="1.5"/><path d="M8 9h8M8 13h5" stroke="#0F6E56" stroke-width="1.5" stroke-linecap="round"/></svg>
        </div>
        <div class="fc-number">02</div>
        <h3 class="fc-title">Info Obat Lengkap</h3>
        <p class="fc-desc">
          Informasi obat yang jelas, dosis yang tepat, dan panduan mudah dipahami untuk semua usia.
        </p>
        <div class="fc-arrow">→</div>
      </div>

      <div class="feature-card fc-blue">
        <div class="fc-icon-wrap">
          <svg class="fc-svg" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg"><polyline points="3,18 8,11 13,14 18,6" stroke="#3874CB" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/><circle cx="18" cy="6" r="2" fill="#3874CB"/></svg>
        </div>
        <div class="fc-number">03</div>
        <h3 class="fc-title">Riwayat Kepatuhan</h3>
        <p class="fc-desc">
          Pantau progres kepatuhan minum obat secara realtime. Laporan harian yang mudah dibaca.
        </p>
        <div class="fc-arrow">→</div>
      </div>

      <div class="feature-card fc-warm">
        <div class="fc-icon-wrap">
          <svg class="fc-svg" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M12 21C12 21 4 13.5 4 8.5a5 5 0 0110 0 5 5 0 0110 0C24 13.5 12 21 12 21z" fill="none" stroke="#ef4444" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg>
        </div>
        <div class="fc-number">04</div>
        <h3 class="fc-title">Rutinitas Sehat</h3>
        <p class="fc-desc">
          Bangun kebiasaan hidup sehat yang stabil setiap hari dengan panduan rutinitas yang terstruktur.
        </p>
        <div class="fc-arrow">→</div>
      </div>

    </div>

  </section>


  <FooterView />

</template>

<style scoped>

/* =========================
   FONT & GLOBAL
========================= */

* {
  font-family: 'Plus Jakarta Sans', sans-serif;
  box-sizing: border-box;
}

.green {
  color: #1D9E75;
}

/* =========================
   HERO
========================= */

.landing {
  min-height: 100vh;

  display: flex;
  align-items: center;
  gap: 60px;

  padding: 120px 8% 80px;

  position: relative;
  overflow: hidden;

  background:
    radial-gradient(ellipse at 10% 50%, rgba(29,158,117,0.06) 0%, transparent 60%),
    radial-gradient(ellipse at 90% 20%, rgba(29,158,117,0.05) 0%, transparent 50%),
    linear-gradient(180deg, #f8fffe 0%, #f0fdf8 60%, #eafaf3 100%);
}

/* DECORATIVE ORBS */
.orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(60px);
  pointer-events: none;
  z-index: 0;
}

.orb-1 {
  width: 400px;
  height: 400px;
  background: rgba(29, 158, 117, 0.08);
  top: -100px;
  left: -100px;
  animation: floatOrb 12s ease-in-out infinite;
}

.orb-2 {
  width: 300px;
  height: 300px;
  background: rgba(52, 211, 153, 0.07);
  bottom: 0px;
  right: 40%;
  animation: floatOrb 9s ease-in-out infinite reverse;
}

.orb-3 {
  width: 200px;
  height: 200px;
  background: rgba(29, 158, 117, 0.06);
  top: 30%;
  right: 0;
  animation: floatOrb 14s ease-in-out infinite 2s;
}

@keyframes floatOrb {
  0%, 100% { transform: translateY(0) scale(1); }
  50% { transform: translateY(-30px) scale(1.05); }
}

.hero-left {
  flex: 1;
  position: relative;
  z-index: 1;
  animation: fadeInUp 0.7s ease both;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(28px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.badge {
  display: inline-flex;
  align-items: center;
  gap: 8px;

  padding: 10px 20px;

  border-radius: 999px;

  background: #e9faf3;
  border: 1.5px solid rgba(29,158,117,0.2);

  color: #1D9E75;

  font-size: 14px;
  font-weight: 700;

  margin-bottom: 28px;

  animation: fadeInUp 0.7s ease 0.1s both;
}

.hero-left h1 {
  font-size: 68px;
  line-height: 1.08;
  font-weight: 900;
  color: #0f172a;
  margin-bottom: 22px;
  animation: fadeInUp 0.7s ease 0.2s both;
}

.hero-left p {
  max-width: 500px;
  font-size: 18px;
  line-height: 1.85;
  color: #374151;
  margin-bottom: 36px;
  animation: fadeInUp 0.7s ease 0.3s both;
}

.hero-left p strong {
  color: #1D9E75;
  font-weight: 700;
}

/* BUTTONS */
.hero-buttons {
  display: flex;
  gap: 16px;
  flex-wrap: wrap;
  animation: fadeInUp 0.7s ease 0.4s both;
}

.btn-icon {
  font-size: 16px;
}

.primary-btn {
  border: none;
  padding: 16px 32px;
  border-radius: 14px;

  background: linear-gradient(135deg, #1D9E75, #0F6E56);

  color: white;
  font-size: 16px;
  font-weight: 700;
  font-family: 'Plus Jakarta Sans', sans-serif;

  cursor: pointer;
  transition: all 0.3s ease;

  display: flex;
  align-items: center;
  gap: 10px;

  box-shadow:
    0 8px 24px rgba(29,158,117,0.32),
    0 2px 6px rgba(29,158,117,0.2);
}

.primary-btn:hover {
  transform: translateY(-4px);
  box-shadow:
    0 16px 36px rgba(29,158,117,0.38),
    0 4px 10px rgba(29,158,117,0.22);
}

.primary-btn:active {
  transform: translateY(-1px);
}

.secondary-btn {
  border: 2px solid #1D9E75;
  padding: 16px 32px;
  border-radius: 14px;
  background: transparent;
  color: #1D9E75;
  font-size: 16px;
  font-weight: 700;
  font-family: 'Plus Jakarta Sans', sans-serif;

  cursor: pointer;
  transition: all 0.3s ease;

  display: flex;
  align-items: center;
  gap: 10px;
}

.secondary-btn:hover {
  background: #1D9E75;
  color: white;
  transform: translateY(-4px);
  box-shadow: 0 12px 28px rgba(29,158,117,0.25);
}

/* TRUST ROW */
.trust-row {
  display: flex;
  align-items: center;
  gap: 20px;
  margin-top: 40px;
  animation: fadeInUp 0.7s ease 0.5s both;
}

.trust-item {
  display: flex;
  flex-direction: column;
  gap: 3px;
}

.trust-num {
  font-size: 20px;
  font-weight: 900;
  color: #1D9E75;
  line-height: 1;
}

.trust-label {
  font-size: 12px;
  color: #6b7280;
  font-weight: 500;
}

.trust-divider {
  width: 1px;
  height: 40px;
  background: rgba(29,158,117,0.2);
}

/* =========================
   PHONE
========================= */

.hero-right {
  flex: 1;
  display: flex;
  justify-content: center;
  position: relative;
  z-index: 1;
  animation: fadeInUp 0.7s ease 0.2s both;
}

.phone-scene {
  position: relative;
  width: 340px;
  height: 600px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.circle-bg {
  position: absolute;
  width: 380px;
  height: 380px;
  border-radius: 50%;
  background: linear-gradient(135deg, #d4f5e8, #b8edda);
  box-shadow: inset 0 0 60px rgba(29,158,117,0.1);
}

/* FLOATING BADGES */
.float-badge {
  position: absolute;
  background: white;
  border-radius: 40px;
  padding: 10px 18px;
  font-size: 13px;
  font-weight: 700;
  color: #0f172a;
  box-shadow: 0 8px 24px rgba(0,0,0,0.1);
  white-space: nowrap;
  z-index: 10;
}

.fb-top {
  top: 30px;
  left: -10px;
  animation: floatBadge 5s ease-in-out infinite;
}

.fb-bottom {
  bottom: 80px;
  right: -10px;
  animation: floatBadge 5s ease-in-out infinite 1.5s;
}

@keyframes floatBadge {
  0%, 100% { transform: translateY(0); }
  50% { transform: translateY(-8px); }
}

.phone-wrapper {
  position: relative;
  z-index: 2;
  transform: rotate(4deg);
  transition: transform 0.4s ease;
}

.phone-wrapper:hover {
  transform: rotate(0deg) translateY(-6px);
}

.phone-shadow {
  position: absolute;
  inset: 0;
  border-radius: 50px;
  background: rgba(0,0,0,0.18);
  filter: blur(48px);
  transform: translateY(50px);
}

.phone {
  width: 265px;
  height: 545px;
  padding: 10px;
  border-radius: 44px;
  background: #111827;
  box-shadow: 0 0 0 1.5px rgba(255,255,255,0.08) inset;
}

.notch {
  position: absolute;
  top: 14px;
  left: 50%;
  transform: translateX(-50%);
  width: 90px;
  height: 22px;
  border-radius: 999px;
  background: #020617;
}

.phone-screen {
  width: 100%;
  height: 100%;
  padding: 36px 16px 16px;
  border-radius: 36px;
  background: white;
  overflow: hidden;
  position: relative;
}

.phone-top {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.hello {
  font-size: 12px;
  color: #6b7280;
  font-weight: 500;
}

.username {
  font-size: 17px;
  font-weight: 800;
  color: #0f172a;
}

.notif {
  width: 38px;
  height: 38px;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #f0fdf8;
  border: 1.5px solid rgba(29,158,117,0.15);
}

.hero-card {
  padding: 18px;
  border-radius: 20px;
  color: white;
  background: linear-gradient(135deg, #1D9E75, #0F6E56);
  box-shadow: 0 8px 20px rgba(29,158,117,0.3);
}

.hero-label {
  font-size: 9px;
  font-weight: 800;
  letter-spacing: 1px;
  opacity: 0.85;
}

.hero-value {
  margin-top: 8px;
  font-size: 24px;
  font-weight: 900;
}

.hero-progress-bar {
  width: 100%;
  height: 6px;
  background: rgba(255,255,255,0.25);
  border-radius: 999px;
  margin: 12px 0 10px;
  overflow: hidden;
}

.hero-progress-fill {
  width: 50%;
  height: 100%;
  background: rgba(255,255,255,0.8);
  border-radius: 999px;
}

.hero-desc {
  font-size: 12px;
  opacity: 0.9;
}

.menu-title {
  margin-top: 16px;
  margin-bottom: 10px;
  font-size: 14px;
  font-weight: 800;
  color: #0f172a;
}

.menu-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 9px;
}

.menu-item {
  padding: 12px;
  border-radius: 16px;
  background: #f8fffe;
  border: 1.5px solid rgba(29,158,117,0.08);
  box-shadow: 0 4px 12px rgba(0,0,0,0.04);
  transition: 0.3s;
  cursor: pointer;
}

.menu-item:hover {
  transform: translateY(-3px);
  box-shadow: 0 8px 20px rgba(29,158,117,0.12);
}

.menu-icon {
  font-size: 20px;
  margin-bottom: 6px;
}

.menu-text {
  font-size: 11px;
  font-weight: 700;
  color: #374151;
}

.bottom-nav {
  position: absolute;
  left: 14px;
  right: 14px;
  bottom: 12px;
  padding: 8px 16px;
  border-radius: 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: rgba(255,255,255,0.95);
  backdrop-filter: blur(10px);
  box-shadow: 0 4px 20px rgba(0,0,0,0.08);
}

.add-btn {
  width: 48px;
  height: 48px;
  margin-top: -28px;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  background: linear-gradient(135deg, #1D9E75, #0F6E56);
  color: white;
  font-size: 26px;
  font-weight: 800;
  box-shadow: 0 6px 16px rgba(29,158,117,0.4);
}

/* =========================
   FEATURES — GRID CARD
========================= */

.feature-section {
  padding: 100px 8%;
  background: #f8fffe;
}

.feature-header {
  text-align: center;
  margin-bottom: 64px;
}

.section-tag {
  display: inline-block;
  color: #1D9E75;
  font-size: 12px;
  font-weight: 800;
  letter-spacing: 2.5px;
  margin-bottom: 14px;
  padding: 6px 16px;
  border-radius: 999px;
  background: #e9faf3;
}

.section-title {
  font-size: 42px;
  font-weight: 900;
  color: #0f172a;
  line-height: 1.2;
  margin-bottom: 16px;
}

.section-subtitle {
  font-size: 17px;
  color: #6b7280;
  max-width: 520px;
  margin: 0 auto;
  line-height: 1.7;
}

.feature-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 24px;
  max-width: 1200px;
  margin: 0 auto;
}

.feature-card {
  position: relative;
  padding: 36px 28px;
  border-radius: 28px;
  background: white;
  overflow: hidden;
  cursor: pointer;
  transition: all 0.4s ease;
  border: 1.5px solid transparent;
}

.feature-card::before {
  content: '';
  position: absolute;
  width: 200px;
  height: 200px;
  border-radius: 50%;
  top: -80px;
  right: -60px;
  opacity: 0.08;
  transition: all 0.4s ease;
}

.feature-card:hover {
  transform: translateY(-10px);
  border-color: transparent;
}

.feature-card:hover::before {
  transform: scale(1.4);
  opacity: 0.14;
}

/* GREEN CARD */
.fc-green {
  box-shadow: 0 10px 30px rgba(29,158,117,0.08);
}
.fc-green::before { background: #1D9E75; }
.fc-green:hover {
  box-shadow: 0 24px 56px rgba(29,158,117,0.2);
  border-color: rgba(29,158,117,0.15);
}

/* TEAL CARD */
.fc-teal {
  box-shadow: 0 10px 30px rgba(15,110,86,0.08);
}
.fc-teal::before { background: #0F6E56; }
.fc-teal:hover {
  box-shadow: 0 24px 56px rgba(15,110,86,0.18);
  border-color: rgba(15,110,86,0.15);
}

/* BLUE CARD */
.fc-blue {
  box-shadow: 0 10px 30px rgba(56,116,203,0.08);
}
.fc-blue::before { background: #3874CB; }
.fc-blue:hover {
  box-shadow: 0 24px 56px rgba(56,116,203,0.18);
  border-color: rgba(56,116,203,0.15);
}

/* WARM CARD */
.fc-warm {
  box-shadow: 0 10px 30px rgba(239,68,68,0.08);
}
.fc-warm::before { background: #ef4444; }
.fc-warm:hover {
  box-shadow: 0 24px 56px rgba(239,68,68,0.15);
  border-color: rgba(239,68,68,0.12);
}

.fc-icon-wrap {
  width: 64px;
  height: 64px;
  border-radius: 20px;
  background: #f0fdf8;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 20px;
  transition: transform 0.4s ease;
}

.fc-green .fc-icon-wrap { background: #e9faf3; }
.fc-teal .fc-icon-wrap { background: #e6f7f2; }
.fc-blue .fc-icon-wrap { background: #eff6ff; }
.fc-warm .fc-icon-wrap { background: #fef2f2; }

.feature-card:hover .fc-icon-wrap {
  transform: scale(1.1) rotate(-6deg);
}

.fc-svg {
  width: 32px;
  height: 32px;
}

.fc-number {
  font-size: 11px;
  font-weight: 800;
  letter-spacing: 1.5px;
  margin-bottom: 10px;
  color: #9ca3af;
}

.fc-green .fc-number { color: #1D9E75; }
.fc-teal .fc-number { color: #0F6E56; }
.fc-blue .fc-number { color: #3874CB; }
.fc-warm .fc-number { color: #ef4444; }

.fc-title {
  font-size: 20px;
  font-weight: 800;
  color: #0f172a;
  margin-bottom: 12px;
  line-height: 1.3;
}

.fc-desc {
  font-size: 14px;
  line-height: 1.75;
  color: #6b7280;
  margin-bottom: 20px;
}

.fc-arrow {
  font-size: 18px;
  font-weight: 700;
  color: #d1d5db;
  transition: all 0.3s ease;
}

.feature-card:hover .fc-arrow {
  transform: translateX(6px);
}

.fc-green:hover .fc-arrow { color: #1D9E75; }
.fc-teal:hover .fc-arrow { color: #0F6E56; }
.fc-blue:hover .fc-arrow { color: #3874CB; }
.fc-warm:hover .fc-arrow { color: #ef4444; }

/* =========================
   CTA SECTION
========================= */

.cta-section {
  position: relative;
  padding: 100px 8%;
  overflow: hidden;
  text-align: center;

  background: linear-gradient(135deg, #1D9E75 0%, #0F6E56 60%, #0a4d3c 100%);
}

.cta-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(70px);
  pointer-events: none;
}

.cta-orb-1 {
  width: 400px;
  height: 400px;
  background: rgba(255,255,255,0.06);
  top: -100px;
  left: -100px;
}

.cta-orb-2 {
  width: 300px;
  height: 300px;
  background: rgba(255,255,255,0.04);
  bottom: -80px;
  right: -60px;
}

.cta-inner {
  position: relative;
  z-index: 1;
  max-width: 700px;
  margin: 0 auto;
}

.cta-icon {
  font-size: 52px;
  margin-bottom: 22px;
  animation: floatBadge 4s ease-in-out infinite;
}

.cta-title {
  font-size: 42px;
  font-weight: 900;
  color: white;
  line-height: 1.2;
  margin-bottom: 18px;
}

.cta-desc {
  font-size: 17px;
  color: rgba(255,255,255,0.85);
  line-height: 1.8;
  margin-bottom: 36px;
  max-width: 560px;
  margin-left: auto;
  margin-right: auto;
}

.cta-buttons {
  display: flex;
  justify-content: center;
  gap: 16px;
  flex-wrap: wrap;
}

.cta-btn-primary {
  border: none;
  padding: 18px 36px;
  border-radius: 14px;
  background: white;
  color: #1D9E75;
  font-size: 16px;
  font-weight: 800;
  font-family: 'Plus Jakarta Sans', sans-serif;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 8px 24px rgba(0,0,0,0.15);
}

.cta-btn-primary:hover {
  transform: translateY(-4px);
  box-shadow: 0 16px 36px rgba(0,0,0,0.2);
  background: #f0fdf8;
}

.cta-btn-secondary {
  border: 2px solid rgba(255,255,255,0.5);
  padding: 18px 36px;
  border-radius: 14px;
  background: transparent;
  color: white;
  font-size: 16px;
  font-weight: 700;
  font-family: 'Plus Jakarta Sans', sans-serif;
  cursor: pointer;
  transition: all 0.3s ease;
}

.cta-btn-secondary:hover {
  background: rgba(255,255,255,0.12);
  border-color: white;
  transform: translateY(-4px);
}

/* =========================
   RESPONSIVE
========================= */

@media (max-width: 1200px) {
  .feature-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 1024px) {

  .landing {
    flex-direction: column;
    text-align: center;
    padding: 120px 5% 80px;
  }

  .hero-buttons {
    justify-content: center;
  }

  .trust-row {
    justify-content: center;
  }

  .hero-left h1 {
    font-size: 52px;
  }

  .hero-left p {
    margin: 0 auto 36px;
  }

  .cta-title {
    font-size: 34px;
  }

}

@media (max-width: 768px) {

  .landing {
    padding: 100px 5% 60px;
    gap: 40px;
  }

  .feature-grid {
    grid-template-columns: 1fr;
    gap: 18px;
  }

  .hero-left h1 {
    font-size: 40px;
  }

  .section-title {
    font-size: 30px;
  }

  .cta-title {
    font-size: 28px;
  }

  .trust-row {
    flex-wrap: wrap;
    gap: 16px;
  }

  .phone-scene {
    transform: scale(0.8);
  }

}

@media (max-width: 480px) {
  .hero-left h1 {
    font-size: 32px;
  }
  
  .phone-scene {
    transform: scale(0.65);
  }

  .section-title {
    font-size: 26px;
  }
}

</style>
