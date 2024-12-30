<template>
  <div class="card card-bordered shadow-xl">
    <figure>
      <img :src="recipe.imageUrl" />
    </figure>
    <div class="card-body">
      <h2 class="card-title">{{ recipe.title }}</h2>
      <div>tag1</div>
      <div class="card-actions justify-end">
        <ThreeDotMenu :id="recipe.id" @delete="deleteRecipe" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Recipe } from '@/types/recipe'
import ThreeDotMenu from '@/components/ThreeDotMenu.vue'

const props = defineProps<{
  recipe: Recipe
}>()

const emit = defineEmits<{
  (event: 'delete', id: string): void
}>()

const deleteRecipe = () => {
  emit('delete', props.recipe.id)
}
</script>
