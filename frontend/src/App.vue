<script setup>
import TopBar from './components/TopBar.vue';
import Login from './components/Login.vue';
import Catalog from './components/Catalog.vue';
</script>

<template>
  <div :style="{ margin: 0, width: '100%', height: '100%', display: 'flex', flexDirection: 'column' }">
    <TopBar />
    <Catalog />
  </div>
</template>
