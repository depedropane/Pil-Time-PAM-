### Fitur PWA yang Sudah Diimplementasikan:
1. **Service Worker Registration** ✅
   - Auto-register service worker dengan `vite-plugin-pwa`
   - Auto-update ketika ada perubahan kode
   - Offline support dengan asset precaching

2. **Web App Manifest** ✅
   - `manifest.webmanifest` auto-generated
   - App metadata (name, short name, description, theme color)
   - Icons untuk berbagai ukuran (192x192, 512x512 SVG)

3. **Installable App** ✅
   - Mobile Web App Capable (iOS & Android)
   - Standalone mode (fullscreen tanpa browser UI)
   - Custom status bar & app title

4. **Responsive Icons** ✅
   - SVG icons di `/public/img/`
   - Maskable icons untuk berbagai bentuk device
   - Dual purpose: `any` dan `maskable`

### Cara Install PWA di Devices:

**Android Chrome:**
1. Buka admin panel di Chrome
2. Tap menu (3 dots) → "Install app"
3. Confirm installation
4. App akan muncul di home screen

**iOS Safari:**
1. Buka admin panel di Safari
2. Tap Share → "Add to Home Screen"
3. Confirm dengan nama
4. App akan siap digunakan offline