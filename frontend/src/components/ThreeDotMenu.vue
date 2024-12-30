<template>
  <details class="dropdown dropdown-top">
    <summary class="btn m-1" @click.stop>
      <i class="fa-solid fa-ellipsis-vertical"></i>
    </summary>
    <ul class="menu dropdown-content bg-base-100 rounded-box z-[1] w-52 p-2 shadow">
      <li @click.stop><a>Edit</a></li>
      <li @click.stop="handleDeleteRecipe"><a>Delete</a></li>
    </ul>
  </details>
</template>

<script setup lang="ts">
const props = defineProps<{
  id: string
}>()

const emit = defineEmits<{
  (event: 'delete', id: string): void
}>()

const handleDeleteRecipe = async () => {
  emit('delete', props.id)
}
</script>
