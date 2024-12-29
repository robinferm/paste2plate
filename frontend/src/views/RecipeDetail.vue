<template>
  <button class="btn" @click="goBack">back</button>
  <div v-if="recipe">
    <h2 class="text-2xl my-2">{{ recipe.title }}</h2>
    <figure>
      <img :src="recipe.imageurl" />
    </figure>
    <ul>
      <li class="columns-2" v-for="ingredient in recipe.ingredients" :key="ingredient.name">
        {{ ingredient.name }}
        <div class="divider-horizontal"></div>
        {{ ingredient.amount }}
      </li>
    </ul>
    <div class="divider"></div>
    <ol>
      <li v-for="instruction in recipe.steps" :key="instruction">
        {{ instruction }}
        <div class="divider"></div>
      </li>
    </ol>
  </div>
</template>

<script setup lang="ts">
import { getRecipe } from '@/api/api'
import type { Recipe } from '@/types/recipe'
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

const props = defineProps<{ id: string }>()
const recipe = ref<Recipe | null>(null)

const fetchRecipe = async (id: string) => {
  recipe.value = await getRecipe(id)
}

const router = useRouter()
const goBack = () => {
  router.back()
}

onMounted(() => {
  fetchRecipe(props.id)
})
</script>
