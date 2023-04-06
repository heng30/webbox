<template>
  <el-dialog
    v-model="popupVisible"
    width="450"
    :show-close="false"
    :close-on-click-modal="false"
    :close-on-press-escape="false"
    title="欢迎使用WebBox"
  >
    <el-form :model="form" :rules="rules" label-width="100px">
      <el-form-item label="用户名" prop="username">
        <el-input v-model="form.username" />
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input v-model="form.password" type="password" />
      </el-form-item>
    </el-form>
    <div
      slot="footer"
      style="display: flex; align-items: center; justify-content: center"
    >
      <el-button @click="handleCancel" style="margin: 0 20px"> 取消 </el-button>
      <el-button type="primary" @click="handleLogin"> 登录 </el-button>
    </div>
  </el-dialog>
</template>

<script lang="ts" setup>
import { ElMessage } from 'element-plus';
import { reactive, ref, watch } from 'vue';
import md5 from 'js-md5';
import store from '../ts/store';
import config from '../ts/config';

const popupVisible = ref(true);

const form = reactive({
  username: '',
  password: '',
});

const rules = {
  username: [{ required: true, message: ' 请输入用户名 ', trigger: 'blur' }],
  password: [{ required: true, message: ' 请输入密码 ', trigger: 'blur' }],
};

watch(store.AuthToken, (newp) => {
  popupVisible.value = !(newp.length > 0);
});

async function handleLogin() {
  const password = md5(form.password);
  const url = `${config.svraddr}/login?username=${form.username}&&password=${password}`;

  try {
    const res = await fetch(url);
    let jdata = await res.json();
    if (jdata.code !== 0) {
      throw new Error(jdata.data);
    }
    store.AuthToken.value = jdata.data;

    ElMessage({
      message: '登陆成功!',
      type: 'success',
    });
  } catch (err) {
    ElMessage.error(`${err}`);
  }
}

function handleCancel() {
  ElMessage.error('请登陆！否则无法使用WebBox!');
}
</script>
