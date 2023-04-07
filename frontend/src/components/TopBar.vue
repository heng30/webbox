<template>
  <div
    style="
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin: 0;
      width: 100%;
      height: 40px;
      border-bottom: 1px solid steelblue;
    "
  >
   <MkdirDialog :current-dir="currentDir"/>
    <el-page-header title="返回" @back="goBack">
      <template #content>
        <el-breadcrumb :separator-icon="ArrowRight">
          <el-breadcrumb-item v-for="(item, index) in bitems"
            ><el-button
              plain
              text
              style="color: black; padding: 4px"
              @click="handleClick(index)"
              >{{ item }}</el-button
            ></el-breadcrumb-item
          >
        </el-breadcrumb>
      </template>
    </el-page-header>

    <div style="display: flex; margin-right: 10px; align-items: center">
      <div style="display: flex; align-items: center" v-if="isUploading">
        <el-text type="primary" class="mx-1" style="margin-right: 10px"
          >上传进度</el-text
        >
        <el-progress
          style="width: 200px"
          :text-inside="true"
          :stroke-width="20"
          :percentage="uploadRate"
        />
      </div>
      <div style="display: flex; margin-left: 20px">
        <el-tooltip content="添加目录" placement="bottom">
          <el-button
            type="primary"
            size="small"
            :icon="Plus"
            circle
            @click="handleMkdir"
          />
        </el-tooltip>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ArrowRight, Plus } from '@element-plus/icons-vue';
import { watch, ref } from 'vue';
import store from '../ts/store';
import MkdirDialog from './MkdirDialog.vue'

const isShowMkdirDialog = store.IsShowMkdirDialog;
const bitems = store.CurrentPath;
const isUploading = ref(false);
const uploadRate = ref(0);
const currentDir = ref('')

watch(store.UploadProgress, () => {
  if (store.UploadProgress.totalBytes === 0) {
    isUploading.value = false;
  } else {
    isUploading.value = true;
    uploadRate.value = Number(
      (
        (store.UploadProgress.uploadedBytes * 100) /
        store.UploadProgress.totalBytes
      ).toFixed(0)
    );
  }
});

const handleClick = (index: number) => {
  if (index >= bitems.value.length - 1) return;
  bitems.value = bitems.value.slice(0, index + 1);
};

const goBack = () => {
  if (bitems.value.length <= 1) return;
  bitems.value = bitems.value.slice(0, bitems.value.length - 1);
};

const handleMkdir = () => {
  if (bitems.value.length === 1) {
    currentDir.value = '/';
  } else {
    currentDir.value = bitems.value.join('/').substring(1);
  }
  isShowMkdirDialog.value = true;
};
</script>
