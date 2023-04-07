<template>
  <el-dialog v-model="isShow" width="450" title="文件重命名">
    <el-form :model="form" :rules="rules" label-width="100px">
      <el-form-item label="新文件名" prop="newname">
        <el-input v-model="form.newname" clearable />
      </el-form-item>
    </el-form>
    <div
      slot="footer"
      style="display: flex; align-items: center; justify-content: center"
    >
      <el-button @click="handleCancel" style="margin: 0 20px">
        取消
      </el-button>
      <el-button type="primary" @click="handleOk"> 确定 </el-button>
    </div>
  </el-dialog>
</template>

<script lang="ts" setup>
import { reactive, defineProps } from 'vue';
import store from '../ts/store';

const props = defineProps({
  oldname: {
    type: String,
    required: true,
    default: '',
  },
});

const isShow = store.IsShowRenameDialog;
const form = reactive({ newname: '' });

const rules = {
  newname: [{ required: true, message: '请输入文件名', trigger: 'blur' }],
};

function handleCancel() {
  form.newname = '';
  isShow.value = false;
}

function handleOk() {
  store.GlobalFunc['rename'](props.oldname, form.newname);
  form.newname = '';
  isShow.value = false;
}
</script>
