<template>

  <footer class="guest-footer">
    <div class="footer-inner">

      <!-- BRAND -->
      <div class="footer-brand">
        <div class="footer-logo">
          <!-- Logo SVG: Pill icon -->
          <svg class="footer-logo-svg" viewBox="0 0 40 40" fill="none" xmlns="http://www.w3.org/2000/svg">
            <rect width="40" height="40" rx="12" fill="#1D9E75"/>
            <path d="M12 20C12 15.582 15.582 12 20 12s8 3.582 8 8-3.582 8-8 8-8-3.582-8-8z" stroke="white" stroke-width="2"/>
            <path d="M20 12v16M12 20h16" stroke="white" stroke-width="2" stroke-linecap="round"/>
          </svg>
          <h2>Pil Time</h2>
        </div>

        <p>
          Platform pengingat minum obat yang dirancang untuk membantu pasien,
          keluarga, dan tenaga kesehatan menjaga kepatuhan pengobatan setiap hari.
        </p>

        <div class="footer-badges">
          <div class="fbadge">
            <svg viewBox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M8 1.5a6.5 6.5 0 100 13 6.5 6.5 0 000-13zM0 8a8 8 0 1116 0A8 8 0 010 8z" fill="#4ade80"/>
              <path d="M8 4a.75.75 0 01.75.75v3.5h2a.75.75 0 010 1.5h-2.75A.75.75 0 017.25 9V4.75A.75.75 0 018 4z" fill="#4ade80"/>
            </svg>
            Medis Terverifikasi
          </div>
          <div class="fbadge">
            <svg viewBox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path fill-rule="evenodd" d="M8 1a7 7 0 100 14A7 7 0 008 1zM5.5 8a.5.5 0 000 1h5a.5.5 0 000-1h-5z" fill="#4ade80"/>
              <path d="M8 5.5a.5.5 0 01.5.5v4a.5.5 0 01-1 0V6a.5.5 0 01.5-.5z" fill="#4ade80"/>
            </svg>
            Privasi Terlindungi
          </div>
          <div class="fbadge">
            <svg viewBox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M8 .5a7.5 7.5 0 100 15A7.5 7.5 0 008 .5zm3.53 5.03l-4 4a.75.75 0 01-1.06 0l-2-2a.75.75 0 011.06-1.06l1.47 1.47 3.47-3.47a.75.75 0 011.06 1.06z" fill="#4ade80"/>
            </svg>
            Gratis Selamanya
          </div>
        </div>
      </div>

      <!-- NAVIGASI -->
      <div class="footer-links-group">
        <h4>Navigasi</h4>
        <nav class="footer-links">
          <router-link to="/">
            <svg viewBox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M6.5 14.5v-3.505c0-.245.25-.495.5-.495h2c.25 0 .5.25.5.5v3.5a.5.5 0 00.5.5h3.5a.5.5 0 00.5-.5V7.636a.5.5 0 00-.153-.357l-6-5.744a.5.5 0 00-.694 0l-6 5.744A.5.5 0 001 7.636V14.5a.5.5 0 00.5.5H5a.5.5 0 00.5-.5V11h-1v3.5h-3V7.915l6-5.741 6 5.741V14.5h-3V11h-1v3.5h-1.5a.5.5 0 01-.5-.5z" fill="currentColor"/>
            </svg>
            Beranda
          </router-link>
          <router-link to="/tentang">
            <svg viewBox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M8 15A7 7 0 118 1a7 7 0 010 14zm0 1A8 8 0 108 0a8 8 0 000 16z" fill="currentColor"/>
              <path d="M8.93 6.588l-2.29.287-.082.38.45.083c.294.07.352.176.288.469l-.738 3.468c-.194.897.105 1.319.808 1.319.545 0 1.178-.252 1.465-.598l.088-.416c-.2.176-.492.246-.686.246-.275 0-.375-.193-.304-.533L8.93 6.588zm.252-2.096a1 1 0 11-2 0 1 1 0 012 0z" fill="currentColor"/>
            </svg>
            Tentang
          </router-link>
          <router-link to="/install">
            <svg viewBox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M.5 9.9a.5.5 0 01.5.5v2.5a1 1 0 001 1h12a1 1 0 001-1v-2.5a.5.5 0 011 0v2.5a2 2 0 01-2 2H2a2 2 0 01-2-2v-2.5a.5.5 0 01.5-.5z" fill="currentColor"/>
              <path d="M7.646 11.854a.5.5 0 00.708 0l3-3a.5.5 0 00-.708-.708L8.5 10.293V1.5a.5.5 0 00-1 0v8.793L5.354 8.146a.5.5 0 10-.708.708l3 3z" fill="currentColor"/>
            </svg>
            Unduh Aplikasi
          </router-link>
          <router-link to="/panduan">
            <svg viewBox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M1 2.828c.885-.37 2.154-.769 3.388-.893 1.33-.134 2.458.063 3.112.752v9.746c-.935-.53-2.12-.603-3.213-.493-1.18.12-2.37.461-3.287.811V2.828zm7.5-.141c.654-.689 1.782-.886 3.112-.752 1.234.124 2.503.523 3.388.893v9.923c-.918-.35-2.107-.692-3.287-.81-1.094-.111-2.278-.039-3.213.492V2.687zM8 1.783C7.015.936 5.587.81 4.287.94c-1.514.153-3.042.672-3.994 1.105A.5.5 0 000 2.5v11a.5.5 0 00.707.455c.882-.4 2.303-.881 3.68-1.02 1.409-.142 2.59.087 3.223.877a.5.5 0 00.78 0c.633-.79 1.814-1.019 3.222-.877 1.378.139 2.8.62 3.681 1.02A.5.5 0 0016 13.5v-11a.5.5 0 00-.293-.455c-.952-.433-2.48-.952-3.994-1.105C10.413.809 8.985.936 8 1.783z" fill="currentColor"/>
            </svg>
            Panduan
          </router-link>
        </nav>
      </div>

      <!-- KONTAK -->
      <div class="footer-links-group">
        <h4>Kontak & Info</h4>
        <div class="footer-contact">
          <div class="fcontact-item">
            <svg viewBox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M0 4a2 2 0 012-2h12a2 2 0 012 2v8a2 2 0 01-2 2H2a2 2 0 01-2-2V4zm2-1a1 1 0 00-1 1v.217l7 4.2 7-4.2V4a1 1 0 00-1-1H2zm13 2.383l-4.758 2.855L15 11.114v-5.73zm-.034 6.878L9.271 8.82 8 9.583 6.728 8.82l-5.694 3.44A1 1 0 002 13h12a1 1 0 00.966-.739zM1 11.114l4.758-2.876L1 5.383v5.73z" fill="currentColor"/>
            </svg>
            <span>piltime@gmail.com</span>
          </div>
          <div class="fcontact-item">
            <svg viewBox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M2 2a2 2 0 012-2h8a2 2 0 012 2v13.5a.5.5 0 01-.777.416L8 13.101l-5.223 2.815A.5.5 0 012 15.5V2zm2-1a1 1 0 00-1 1v12.566l4.723-2.482a.5.5 0 01.554 0L13 14.566V2a1 1 0 00-1-1H4z" fill="currentColor"/>
            </svg>
            <span>Institut Teknologi Del</span>
          </div>
          <div class="fcontact-item">
            <svg viewBox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M8 16s6-5.686 6-10A6 6 0 002 6c0 4.314 6 10 6 10zm0-7a3 3 0 110-6 3 3 0 010 6z" fill="currentColor"/>
            </svg>
            <span>Desa Nauli, Sumatera Utara</span>
          </div>
        </div>
      </div>

    </div>

    <!-- BOTTOM BAR -->
    <div class="footer-bottom">
      <p>© 2026 Pil Time. Seluruh hak cipta dilindungi undang-undang.</p>
      <p>Dikembangkan oleh Tim Pil Time &mdash; Institut Teknologi Del</p>
    </div>
  </footer>
</template>

<script>
export default {
  name: "FooterView",
}
</script>

<style scoped>

/* FOOTER */
.guest-footer {
  padding: 64px 8% 0;
  background: #0f1f17;
  color: #f0fdf8;
  font-family: 'Plus Jakarta Sans', sans-serif;
}

.footer-inner {
  display: grid;
  grid-template-columns: 1.8fr 1fr 1.2fr;
  gap: 64px;
  padding-bottom: 52px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.07);
  align-items: start;
}

/* ── BRAND ────────────────── */
.footer-brand {
  max-width: 360px;
}

.footer-logo {
  display: flex;
  align-items: center;
  gap: 14px;
  margin-bottom: 18px;
}

.footer-logo-svg {
  width: 40px;
  height: 40px;
  flex-shrink: 0;
  border-radius: 10px;
}

.footer-brand h2 {
  font-size: 24px;
  font-weight: 900;
  color: #f0fdf8;
  letter-spacing: -0.3px;
}

.footer-brand > p {
  color: #7aaa8e;
  line-height: 1.9;
  font-size: 14px;
  margin-bottom: 24px;
}

/* BADGES */
.footer-badges {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.fbadge {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  font-weight: 600;
  color: #94b8a4;
}

.fbadge svg {
  width: 14px;
  height: 14px;
  flex-shrink: 0;
}

/* ── LINKS GROUP ─────────── */
.footer-links-group h4 {
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 2px;
  color: #4ade80;
  text-transform: uppercase;
  margin-bottom: 22px;
}

.footer-links {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.footer-links a {
  display: flex;
  align-items: center;
  gap: 10px;
  color: #7aaa8e;
  text-decoration: none;
  font-size: 14px;
  font-weight: 500;
  transition: color 0.2s ease;
}

.footer-links a svg {
  width: 14px;
  height: 14px;
  flex-shrink: 0;
  opacity: 0.6;
  transition: opacity 0.2s ease;
}

.footer-links a:hover {
  color: #f0fdf8;
}

.footer-links a:hover svg {
  opacity: 1;
}

/* ── CONTACT ─────────────── */
.footer-contact {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.fcontact-item {
  display: flex;
  align-items: flex-start;
  gap: 10px;
  color: #7aaa8e;
  font-size: 14px;
  line-height: 1.6;
}

.fcontact-item svg {
  width: 14px;
  height: 14px;
  flex-shrink: 0;
  margin-top: 2px;
  opacity: 0.6;
}

/* ── BOTTOM BAR ──────────── */
.footer-bottom {
  padding: 22px 0 26px;
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  justify-content: space-between;
  align-items: center;
  color: #4d7a5f;
  font-size: 13px;
}

/* ── RESPONSIVE ──────────── */
@media (max-width: 1024px) {
  .footer-inner {
    grid-template-columns: 1fr 1fr;
    gap: 48px;
  }

  .footer-brand {
    grid-column: 1 / -1;
    max-width: 100%;
  }
}

@media (max-width: 640px) {
  .footer-inner {
    grid-template-columns: 1fr;
    gap: 36px;
  }

  .footer-bottom {
    flex-direction: column;
    align-items: flex-start;
    gap: 6px;
  }
}

</style>
