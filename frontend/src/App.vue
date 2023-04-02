<script setup>
import TopBar from './components/TopBar.vue';
import Login from './components/Login.vue';
import Catalog from './components/Catalog.vue';
</script>

<template>
  <div :style="{ margin: 0, width: '100%', display: 'flex', flexDirection: 'column' }">
    <TopBar msg="You did it!" />
    <Catalog />
  </div>
</template>
