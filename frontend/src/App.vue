<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'
import { RouterView } from 'vue-router'
import router from '@/router'
const authStore = useAuthStore()

router.beforeEach(async (to) => {
  if (authStore.token) {
    await useAuthStore().authRefresh()
  }

  if (authStore.isLoggedIn() && to.name === 'Home') {
    return { name: 'RecipeList' }
  }

  if (!authStore.isLoggedIn() && to.name !== 'Home') {
    return { name: 'Home' }
  }
})
</script>

<template>
  <header class="relative flex items-center justify-center">
    <button
      class="btn btn-ghost absolute left-0"
      v-if="
        router.currentRoute.value.name !== 'Home' && router.currentRoute.value.name !== 'RecipeList'
      "
      @click="router.back"
    >
      <i class="fa-solid fa-arrow-left"></i>
    </button>
    <div class="m-2">Paste2Plate</div>
  </header>
  <main class="m-2">
    <RouterView />
  </main>
</template>
