import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import './style.css'
// Import Excalidraw styles
import './components/document/excalidraw/excalidraw-vendor.css'
import './components/document/excalidraw/excalidraw.css'

const app = createApp(App)
const pinia = createPinia()
app.use(pinia)
app.use(router)

// Initialiser l'authentification après que Pinia soit configuré
import { useAuthStore } from './stores/auth'
const authStore = useAuthStore()
authStore.initializeAuth()

app.mount('#app')
