<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'
import { RouterView } from 'vue-router'
import router from '@/router'
const authStore = useAuthStore()

router.beforeEach(async (to) => {
  if (authStore.token) {
    await useAuthStore().authRefresh()
  }
  if (!authStore.isLoggedIn() && to.name !== 'Home') {
    return { name: 'Home' }
  }
})
</script>

<template>
  <main class="m-2">
    <RouterView />
  </main>
</template>
