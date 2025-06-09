<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import { 
  EnvelopeIcon, 
  EyeIcon,
  EyeSlashIcon,
} from '@heroicons/vue/24/outline'

const email = ref('')
const password = ref('')
const error = ref('')
const showPassword = ref(false)
const authStore = useAuthStore()
const router = useRouter()
const route = useRoute()
const showSuccessMessage = ref(false)

onMounted(() => {
  showSuccessMessage.value = route.query.registered === 'true'
})

const handleSubmit = async () => {
  try {
    await authStore.login(email.value, password.value)
    router.push('/')
  } catch (err: any) {
    error.value = err.message
  }
}

const togglePassword = () => {
  showPassword.value = !showPassword.value
}
</script>

<template>
  <div class="font-[sans-serif]">
    <div class="min-h-screen flex fle-col items-center justify-center py-6 px-4 bg-white">
      <div class="grid md:grid-cols-2 items-center gap-6 max-w-6xl w-full">
        <div class="border border-gray-300 rounded-lg p-6 max-w-md shadow-[0_2px_22px_-4px_rgba(93,96,127,0.2)] max-md:mx-auto">
          <!-- Afficher le message de succès -->
          <div v-if="showSuccessMessage" class="bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded mb-4">
            Registration successful! Please login with your new account.
          </div>
          
          <form @submit.prevent="handleSubmit" class="space-y-4">
            <div class="mb-8">
              <h3 class="text-gray-800 text-3xl font-bold">Login</h3>
              <p class="text-gray-500 text-sm mt-4">
                Don't have an account? 
                <router-link to="/auth/register" class="text-blue-600 hover:underline">Register</router-link>
              </p>
            </div>

            <div>
              <label class="text-gray-800 text-sm mb-2 block">Email</label>
              <div class="relative flex items-center">
                <input 
                  v-model="email"
                  type="email" 
                  required 
                  class="w-full text-sm text-gray-800 border border-gray-300 pl-4 pr-10 py-3 rounded-lg outline-blue-600" 
                  placeholder="Enter email" 
                />
                <EnvelopeIcon class="w-[18px] h-[18px] absolute right-4 text-gray-400" />
              </div>
            </div>

            <div>
              <label class="text-gray-800 text-sm mb-2 block">Password</label>
              <div class="relative flex items-center">
                <input 
                  v-model="password"
                  :type="showPassword ? 'text' : 'password'"
                  required 
                  class="w-full text-sm text-gray-800 border border-gray-300 pl-4 pr-10 py-3 rounded-lg outline-blue-600" 
                  placeholder="Enter password" 
                />
                <button 
                  type="button"
                  @click="togglePassword"
                  class="absolute right-4"
                >
                  <EyeIcon v-if="!showPassword" class="w-[18px] h-[18px] text-gray-400" />
                  <EyeSlashIcon v-else class="w-[18px] h-[18px] text-gray-400" />
                </button>
              </div>
            </div>

            <div v-if="error" class="text-red-500 text-sm">
              {{ error }}
            </div>

            <div class="!mt-8">
              <button 
                type="submit"
                class="w-full shadow-xl py-2.5 px-4 text-sm font-medium tracking-wide rounded-lg text-white bg-blue-600 hover:bg-blue-700 focus:outline-none"
              >
                Sign in
              </button>
            </div>
          </form>
        </div>

        <div class="max-md:mt-8">
          <img src="https://readymadeui.com/login-image.webp" class="w-full aspect-[71/50] max-md:w-4/5 mx-auto block object-cover" alt="Dining Experience" />
        </div>
      </div>
    </div>
  </div>
</template>