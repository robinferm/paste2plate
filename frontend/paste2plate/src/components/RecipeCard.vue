<template>
  <div class="card lg:card-side bg-base-100 shadow-xl">
    <figure>
      <img :src="recipe.imageurl" />
    </figure>
    <div class="card-body">
      <h2 class="card-title">{{ recipe.title }}</h2>
      <div>tag1</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Recipe } from '@/types/recipe'

defineProps<{
  recipe: Recipe
}>()
</script>
