<template>
  <div class="form-control gap-4 w-fit mx-auto">
    <form @submit.prevent="loginWithPassword" class="form-control gap-4">
      <label class="input input-bordered flex items-center gap-2 w-72">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 16 16"
          fill="currentColor"
          class="h-4 w-4 opacity-70"
        >
          <path
            d="M8 8a3 3 0 1 0 0-6 3 3 0 0 0 0 6ZM12.735 14c.618 0 1.093-.561.872-1.139a6.002 6.002 0 0 0-11.215 0c-.22.578.254 1.139.872 1.139h9.47Z"
          />
        </svg>
        <input type="text" class="grow" placeholder="email" v-model="email" required />
      </label>
      <label class="input input-bordered flex items-center gap-2 w-72">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 16 16"
          fill="currentColor"
          class="h-4 w-4 opacity-70"
        >
          <path
            fill-rule="evenodd"
            d="M14 6a4 4 0 0 1-4.899 3.899l-1.955 1.955a.5.5 0 0 1-.353.146H5v1.5a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1-.5-.5v-2.293a.5.5 0 0 1 .146-.353l3.955-3.955A4 4 0 1 1 14 6Zm-4-2a.75.75 0 0 0 0 1.5.5.5 0 0 1 .5.5.75.75 0 0 0 1.5 0 2 2 0 0 0-2-2Z"
            clip-rule="evenodd"
          />
        </svg>
        <input type="password" class="grow" placeholder="password" v-model="password" required />
      </label>
      <button class="btn w-72" type="submit">Login</button>
    </form>
    <button class="btn w-72" @click="loginWithGoogle">Login with Google</button>
  </div>
</template>

<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'
import { ref } from 'vue'

const authStore = useAuthStore()

const loginWithGoogle = async () => {
  try {
    await authStore.loginWithGoogle()
  } catch (err) {
    console.error('Login failed:', err)
  }
}

const email = ref('')
const password = ref('')

const loginWithPassword = async () => {
  try {
    await authStore.loginWithPassword(email.value, password.value)
  } catch (err) {
    console.error('Login failed:', err)
  }
}
</script>
