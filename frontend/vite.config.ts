import { fileURLToPath, URL } from 'node:url'

import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'
import { VitePWA } from 'vite-plugin-pwa'

// https://vite.dev/config/
export default defineConfig(({ command, mode }) => {
  const env = loadEnv(mode, process.cwd(), '')
  return {
    plugins: [
      vue(),
      vueDevTools(),
      VitePWA({
        registerType: 'autoUpdate',
        includeAssets: ['favicon.ico', 'robots.txt'],
        manifest: {
          name: 'Paste2Plate PWA',
          short_name: 'Paste2Plate',
          description: 'Paste2Plate',
          theme_color: '#ffffff',
          icons: [
            {
              src: '/pwa/pwa-192x192.png',
              sizes: '192x192',
              type: 'image/png',
            },
            {
              src: '/pwa/pwa-512x512.png',
              sizes: '512x512',
              type: 'image/png',
            },
          ],
          screenshots: [
            {
              src: '/pwa/vue-wide.png',
              sizes: '2048x1766',
              type: 'image/png',
              form_factor: 'wide',
            },
            {
              src: '/pwa/pwa-192x192.png',
              sizes: '192x192',
              type: 'image/png',
            },
          ],
        },
      }),
    ],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url)),
      },
    },
    server: {
      proxy: {
        '/api': env.VITE_BACKEND_URL,
      },
    },
  }
})
