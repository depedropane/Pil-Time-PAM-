<template>
  <nav class="navbar">

    <!-- BACKGROUND BLUR -->
    <div class="navbar-glow"></div>

    <!-- LOGO -->
    <div class="logo-section">

      <div class="logo-wrapper">

        <div class="logo-ring"></div>

        <div class="logo-icon">
          <img
            src="./../../assets/images/logo.jpeg"
            alt="Logo"
          />
        </div>

      </div>

      <div class="logo-content">

        <h1 class="logo-text">
          Pil Time
        </h1>

        <span class="logo-subtitle">
          Smart Medicine Reminder
        </span>

      </div>

    </div>

    <!-- MENU -->
    <ul class="nav-links">

      <li>
        <a href="/" class="active-link">
          Dashboard
        </a>
      </li>

      <li>
        <a href="/tentang">
          Tentang
        </a>
      </li>

      <li>
        <a href="/install">
          Install
        </a>
      </li>

      <li>
        <a href="/panduan">
          Panduan
        </a>
      </li>

      <li>
        <a href="/login" class="login-btn">
          Login
        </a>
      </li>

    </ul>

    <!-- MOBILE TOGGLE -->
    <div class="menu-toggle" @click="toggleMenu">

      <span></span>
      <span></span>
      <span></span>

    </div>

    <!-- MOBILE MENU -->
    <div class="mobile-menu" :class="{ active: isOpen }">

      <a href="/">Dashboard</a>
      <a href="/tentang">Tentang</a>
      <a href="/install">Install</a>
      <a href="/panduan">Panduan</a>

      <a href="/login" class="mobile-login">
        Login
      </a>

    </div>

  </nav>
</template>

<script>
export default {
  name: "NavbarView",

  data() {
    return {
      isOpen: false,
    };
  },

  methods: {
    toggleMenu() {
      this.isOpen = !this.isOpen;
    },
  },
};
</script>

<style scoped>

/* =========================
   RESET
========================= */

*{
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

/* =========================
   NAVBAR
========================= */

.navbar{
  width: 100%;
  height: 88px;

  position: fixed;
  top: 0;
  left: 0;
  z-index: 999;

  padding: 0 8%;

  display: flex;
  align-items: center;
  justify-content: space-between;

  background: rgba(255,255,255,0.92);

  backdrop-filter: blur(20px);

  border-bottom: 1px solid rgba(29,158,117,0.08);

  box-shadow:
    0 4px 24px rgba(15,23,42,0.07);

  font-family: 'Plus Jakarta Sans', sans-serif;
}

/* =========================
   GLOW EFFECT
========================= */

.navbar-glow{
  position: absolute;

  width: 300px;
  height: 300px;

  background:
    radial-gradient(
      circle,
      rgba(29,158,117,0.15),
      transparent 70%
    );

  top: -120px;
  left: -100px;

  z-index: -1;
}

/* =========================
   LOGO
========================= */

.logo-section{
  display: flex;
  align-items: center;
  gap: 16px;
}

.logo-wrapper{
  position: relative;
}

.logo-ring{
  position: absolute;

  inset: -5px;

  border-radius: 50%;

  background:
    linear-gradient(
      135deg,
      #1D9E75,
      #34d399,
      #0F766E
    );

  animation: rotateRing 6s linear infinite;
}

.logo-icon{
  width: 62px;
  height: 62px;

  border-radius: 50%;

  overflow: hidden;

  position: relative;
  z-index: 2;

  border: 3px solid white;

  background: white;

  box-shadow:
    0 10px 25px rgba(29,158,117,0.25);
}

.logo-icon img{
  width: 100%;
  height: 100%;

  object-fit: cover;
}

.logo-content{
  display: flex;
  flex-direction: column;
}

.logo-text{
  font-size: 28px;
  font-weight: 900;

  color: #0f172a;

  line-height: 1;
}

.logo-subtitle{
  font-size: 12px;
  font-weight: 600;

  color: #64748b;

  margin-top: 4px;
}

/* =========================
   NAVIGATION
========================= */

.nav-links{
  display: flex;
  align-items: center;
  gap: 38px;

  list-style: none;
}

.nav-links a{
  text-decoration: none;

  color: #374151;

  font-size: 16px;
  font-weight: 700;

  position: relative;

  transition: 0.3s ease;
}

.nav-links a:hover{
  color: #1D9E75;
}

/* UNDERLINE */

.nav-links a::after{
  content: "";

  position: absolute;

  left: 0;
  bottom: -8px;

  width: 0%;
  height: 3px;

  border-radius: 999px;

  background:
    linear-gradient(
      90deg,
      #1D9E75,
      #34d399
    );

  transition: 0.35s ease;
}

.nav-links a:hover::after,
.active-link::after{
  width: 100%;
}

/* =========================
   LOGIN BUTTON
========================= */

.login-btn{
  padding: 15px 32px;

  border-radius: 14px;

  background:
    linear-gradient(
      135deg,
      #1D9E75,
      #0F6E56
    );

  color: white !important;

  font-size: 16px;

  box-shadow:
    0 10px 28px rgba(29,158,117,0.3);

  transition: 0.35s ease;
}

.login-btn:hover{
  transform:
    translateY(-4px)
    scale(1.03);

  box-shadow:
    0 18px 35px rgba(29,158,117,0.35);
}

.login-btn::after{
  display: none;
}

/* =========================
   MENU TOGGLE
========================= */

.menu-toggle{
  display: none;

  flex-direction: column;
  gap: 5px;

  cursor: pointer;
}

.menu-toggle span{
  width: 28px;
  height: 3px;

  border-radius: 999px;

  background: #0f172a;

  transition: 0.3s;
}

/* =========================
   MOBILE MENU
========================= */

.mobile-menu{
  position: absolute;

  top: 88px;
  left: 0;

  width: 100%;

  background: rgba(255,255,255,0.96);

  backdrop-filter: blur(18px);

  display: flex;
  flex-direction: column;

  gap: 20px;

  padding: 28px;

  transform: translateY(-150%);
  opacity: 0;

  transition: 0.4s ease;

  box-shadow:
    0 20px 40px rgba(0,0,0,0.08);
}

.mobile-menu.active{
  transform: translateY(0%);
  opacity: 1;
}

.mobile-menu a{
  text-decoration: none;

  color: #374151;

  font-size: 17px;
  font-weight: 700;

  transition: 0.3s;
}

.mobile-menu a:hover{
  color: #1D9E75;
}

.mobile-login{
  background:
    linear-gradient(
      135deg,
      #1D9E75,
      #0F766E
    );

  color: white !important;

  text-align: center;

  padding: 15px;

  border-radius: 18px;

  margin-top: 10px;
}

/* =========================
   ANIMATION
========================= */

@keyframes rotateRing{

  0%{
    transform: rotate(0deg);
  }

  100%{
    transform: rotate(360deg);
  }

}

/* =========================
   RESPONSIVE
========================= */

@media(max-width: 950px){

  .nav-links{
    display: none;
  }

  .menu-toggle{
    display: flex;
  }

}

@media(max-width: 768px){

  .navbar{
    padding: 0 5%;
  }

  .logo-text{
    font-size: 22px;
  }

  .logo-subtitle{
    font-size: 10px;
  }

  .logo-icon{
    width: 55px;
    height: 55px;
  }

}

</style>