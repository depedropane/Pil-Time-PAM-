import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/auth'

/* =========================
   GUEST PAGES
========================= */
import LandingView from '../views/guests/DashboardView.vue'
import TentangView from '../views/guests/TentangView.vue'
import InstallView from '../views/guests/InstallView.vue'
import PanduanView from '../views/guests/PanduanView.vue'

/* =========================
   AUTH
========================= */
import LoginView from '../views/LoginView.vue'

/* =========================
   ADMIN PAGES
========================= */
import DashboardView from '../views/DashboardView.vue'
import PasienView from '../views/PasienView.vue'
import JadwalView from '../views/jadwal/JadwalView.vue'
import JadwalDetail from '../views/jadwal/JadwalDetail.vue'
import JadwalEdit from '../views/jadwal/JadwalEdit.vue'
import RiwayatView from '../views/RiwayatView.vue'
import ObatView from '../views/ObatView.vue'

const routes = [

  /* =========================
     GUEST ROUTES
  ========================= */

  {
    path: '/',
    name: 'landing',
    component: LandingView,
    meta: {
      requiresAuth: false,
      guest: true
    }
  },

  {
    path: '/tentang',
    name: 'tentang',
    component: TentangView,
    meta: {
      requiresAuth: false,
      guest: true
    }
  },

  {
    path: '/install',
    name: 'install',
    component: InstallView,
    meta: {
      requiresAuth: false,
      guest: true
    }
  },

  {
    path: '/panduan',
    name: 'panduan',
    component: PanduanView,
    meta: {
      requiresAuth: false,
      guest: true
    }
  },

  /* =========================
     LOGIN
  ========================= */

  {
    path: '/login',
    name: 'login',
    component: LoginView,
    meta: {
      requiresAuth: false
    }
  },

  /* =========================
     ADMIN ROUTES
  ========================= */

  {
    path: '/dashboard',
    name: 'dashboard',
    component: DashboardView,
    meta: {
      requiresAuth: true,
      role: 'admin'
    }
  },

  {
    path: '/pasien',
    name: 'pasien',
    component: PasienView,
    meta: {
      requiresAuth: true,
      role: 'admin'
    }
  },

  {
    path: '/jadwal',
    name: 'jadwal',
    component: JadwalView,
    meta: {
      requiresAuth: true,
      role: 'admin'
    }
  },

  {
    path: '/jadwal/:id',
    name: 'jadwal-detail',
    component: JadwalDetail,
    meta: {
      requiresAuth: true,
      role: 'admin'
    }
  },

  {
    path: '/jadwal/:id/edit',
    name: 'jadwal-edit',
    component: JadwalEdit,
    meta: {
      requiresAuth: true,
      role: 'admin'
    }
  },

  {
    path: '/riwayat',
    name: 'riwayat',
    component: RiwayatView,
    meta: {
      requiresAuth: true,
      role: 'admin'
    }
  },

  {
    path: '/obat',
    name: 'obat',
    component: ObatView,
    meta: {
      requiresAuth: true,
      role: 'admin'
    }
  },

  /* =========================
     404 REDIRECT
  ========================= */

  {
    path: '/:pathMatch(.*)*',
    redirect: '/'
  }

]

const router = createRouter({
  history: createWebHistory(),
  routes
})

/* =========================
   NAVIGATION GUARD
========================= */

router.beforeEach((to, from, next) => {

  const authStore = useAuthStore()

  const requiresAuth = to.meta.requiresAuth
  const isLoggedIn = !!authStore.token

  // Jika halaman butuh login
  if (requiresAuth && !isLoggedIn) {
    next('/login')
  }

  // Jika sudah login dan buka login lagi
  else if (to.path === '/login' && isLoggedIn) {
    next('/dashboard')
  }

  // Lanjut
  else {
    next()
  }

})

export default router