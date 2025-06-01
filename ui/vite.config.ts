import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  build: {
    outDir: resolve(__dirname, '../pkg/app/static/files'),
    emptyOutDir: true
  },
  assetsInclude: ['**/*.excalidrawlib'], // Traiter les fichiers .excalidrawlib comme des assets
  server: {
    fs: {
      // Permettre de servir les fichiers en dehors du répertoire racine
      allow: ['..']
    }
  }
})
