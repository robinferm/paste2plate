<template>
  <button class="btn" @click="loginWithGoogle">Login with Google</button>
</template>

<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()

const loginWithGoogle = async () => {
  try {
    await authStore.loginWithGoogle()
    console.log('Login successful!')
  } catch (err) {
    console.error('Login failed:', err)
  }
}
</script>
