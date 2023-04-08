<template>
  <el-dialog v-model="isShow" width="450" title="新建文件夹">
    <el-form :model="form" :rules="rules" label-width="100px" @submit.prevent>
      <el-form-item label="文件名" prop="filename">
        <el-input v-model="form.filename" clearable />
      </el-form-item>
    </el-form>
    <div
      slot="footer"
      style="display: flex; align-items: center; justify-content: center"
    >
      <el-button @click="handleCancel" style="margin: 0 20px"> 取消 </el-button>
      <el-button type="primary" @click="handleOk"> 创建 </el-button>
    </div>
  </el-dialog>
</template>

<script lang="ts" setup>
import { ElInput, ElForm, ElFormItem, ElButton, ElDialog } from 'element-plus';
import { reactive } from 'vue';
import store from '../ts/store';

const props = defineProps({
  currentDir: {
    type: String,
    required: true,
    default: '',
  },
});

const isShow = store.IsShowMkdirDialog;
const form = reactive({ filename: '' });

const rules = {
  filename: [{ required: true, message: '请输入文件名', trigger: 'blur' }],
};

function handleCancel() {
  form.filename = '';
  isShow.value = false;
}

function handleOk() {
  store.GlobalFunc['mkdir'](props.currentDir, form.filename);
  form.filename = '';
  isShow.value = false;
}
</script>
