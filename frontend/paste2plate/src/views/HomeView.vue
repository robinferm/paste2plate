<template>
  <main>
    <div class="grid lg:grid-cols-3 gap-4">
      <RecipeCard></RecipeCard>
      <RecipeCard></RecipeCard>
      <RecipeCard></RecipeCard>
      <RecipeCard></RecipeCard>
    </div>
  </main>
</template>

<script setup lang="ts">
import RecipeCard from '@/components/RecipeCard.vue'
</script>
