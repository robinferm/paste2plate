<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'
import { RouterView } from 'vue-router'
import { onMounted } from 'vue'
const authStore = useAuthStore()

const loginWithGoogle = async () => {
  try {
    await authStore.loginWithGoogle()
    console.log('Login successful!')
  } catch (err) {
    console.error('Login failed:', err)
  }
}

onMounted(async () => {
  await useAuthStore().authRefresh()
})
</script>

<template>
  <header>
    <button class="btn" v-if="!authStore.user" @click="loginWithGoogle">Login with Google</button>
    <button class="btn" v-else @click="authStore.logout">Logout</button>
  </header>

  <main>
    <RouterView v-if="authStore.user" />
  </main>
</template>
