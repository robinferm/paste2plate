<template>
  <div class="card lg:card-side bg-base-100 shadow-xl">
    <figure>
      <img
        src="https://img.daisyui.com/images/stock/photo-1494232410401-ad00d5433cfa.webp"
        alt="Album"
      />
    </figure>
    <div class="card-body">
      <h2 class="card-title">{{recipe.title}}</h2>
      <p>wapdjawdoandoandalwdnaklwdn</p>
      <div>tag1</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Recipe } from '@/types/recipe';
import { defineProps } from 'vue';

defineProps<{
  recipe: Recipe;
}>();

</script>
