<template>
  <div style="display: inline-block">
    <el-button
      type="primary"
      :size="size"
      :disabled="disabled"
      @click="openFileDialog"
    >
      {{ buttonLabel }}
    </el-button>
    <input id="file_dialog" type="file" style="display: none" ref="inputFile" />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
const props = defineProps(['buttonLabel', 'size', 'disabled', 'attach']);
const emit = defineEmits(['accept']); // 返回文件名和 ArrayBuffer 数据

const inputFile = ref(null);

function openFileDialog() {
  inputFile.value.click();
}

onMounted(() => {
  inputFile.value.addEventListener('change', () => {
    const file = inputFile.value.files[0];
    emit('accept', file, props.attach);
    inputFile.value.value = '';
  });
});
</script>
