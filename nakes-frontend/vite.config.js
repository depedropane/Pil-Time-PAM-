import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { VitePWA } from 'vite-plugin-pwa'

export default defineConfig({
  plugins: [
    vue(),

    VitePWA({

      registerType: 'autoUpdate',

      manifest: {
        name: 'Admin Panel - Pil Time',
        short_name: 'Pil Time Admin',
        description: 'Admin Panel untuk Tracking Jadwal & Obat',

        theme_color: '#1e40af',
        background_color: '#ffffff',

        start_url: '/',
        scope: '/',

        display: 'standalone',
        orientation: 'portrait-primary',

        icons: [
          {
            src: '/img/icon-192x192.svg',
            sizes: '192x192',
            type: 'image/svg+xml',
            purpose: 'any'
          },
          {
            src: '/img/icon-512x512.svg',
            sizes: '512x512',
            type: 'image/svg+xml',
            purpose: 'any'
          },
          {
            src: '/img/icon-192x192.svg',
            sizes: '192x192',
            type: 'image/svg+xml',
            purpose: 'maskable'
          },
          {
            src: '/img/icon-512x512.svg',
            sizes: '512x512',
            type: 'image/svg+xml',
            purpose: 'maskable'
          }
        ]
      },

      /* =========================
         WORKBOX FIX
      ========================= */

      workbox: {

        navigateFallback: '/index.html',

        navigateFallbackAllowlist: [
          /^\/.*/
        ]

      },

      devOptions: {
        enabled: true
      }

    })
  ],

  server: {
    port: 5173,

    proxy: {
      '/api': {
        target: 'https://pil-time-pam-production.up.railway.app',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '/api/v1')
      }
    }
  }
})