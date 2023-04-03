<template>
  <div
    style="
      margin: 0;
      width: 100%;
      height: 40px;
      borderBottom: 2px solid steelblue;
    "
  >
    <el-page-header>
      <template #content>
        <div style="display: flex">
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
        </div>
      </template>
    </el-page-header>
  </div>
</template>

<script lang="ts" setup>
import { ArrowRight } from '@element-plus/icons-vue';
import store from '../ts/store';
const bitems = store.CurrentPath;

const handleClick = (index: number) => {
  if (index >= bitems.value.length - 1) return;
  bitems.value = bitems.value.slice(0, index + 1);
};
</script>

<!-- defineProps({ -->
<!--   msg: { -->
<!--     type: String, -->
<!--     required: true, -->
<!--   }, -->
<!-- }); -->
