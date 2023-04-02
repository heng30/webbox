<template>
  <div
    :style="{
      margin: 0,
      width: '100%',
      backgroundColor: 'steelblue',
    }"
  >
    <h1 :style="{ textAlign: 'center' }">{{ msg }}</h1>
  </div>
</template>

<script lang="ts" setup>
defineProps({
  msg: {
    type: String,
    required: true,
  },
});
</script>
