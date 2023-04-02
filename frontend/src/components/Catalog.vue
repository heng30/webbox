<template>
  <el-table :data="filterTableData" style="width: 100%">
    <el-table-column label="序号" prop="index" width="80" />
    <el-table-column label="名称" prop="name" />
    <el-table-column label="类型" prop="type" width="80" />
    <el-table-column label="大小" prop="size" width="100" />
    <el-table-column align="right" width="300">
      <template #header>
        <el-input v-model="search" size="small" placeholder="Type to search" />
      </template>
      <template #default="scope">
        <el-button size="small" @click="handleUpload(scope.$index, scope.row)"
          >上传</el-button
        >
        <el-button size="small" @click="handleDownload(scope.$index, scope.row)"
          >下载</el-button
        >
        <el-button
          size="small"
          type="danger"
          @click="handleDelete(scope.$index, scope.row)"
          >删除</el-button
        >
      </template>
    </el-table-column>
  </el-table>
</template>

<script lang="ts" setup>
import config from '../js/config.js';
import { onMounted, computed, ref } from 'vue';

interface Catalog {
  index: number;
  type: string;
  name: string;
  size: string;
}

const tableData = ref([]);
const search = ref('');

const filterTableData = computed(() =>
  tableData.value.filter(
    (data) =>
      !search.value ||
      data.name.toLowerCase().includes(search.value.toLowerCase())
  )
);

onMounted(async () => {
  await fetchCatalog('/');
});

const handleDownload = (index: number, row: Catalog) => {
  console.log(index, row);
};

const handleUpload = (index: number, row: Catalog) => {
  console.log(index, row);
};

const handleDelete = (index: number, row: Catalog) => {
  console.log(index, row);
};

const fetchCatalog = async (path: string) => {
  const url = `${config.svraddr}/readdir?path=${path}`;
  console.log(url);
  try {
    const res = await fetch(url, {
      headers: {
        Origin: '*',
      },
    });
    if (!res.ok) {
      throw new Error('Network response was not ok');
    }

    let jdata = await res.json();
    if (jdata.code !== 0) {
      throw new Error(jdata.data);
    }

    const data: any[] = JSON.parse(jdata.data);
    tableData.value.splice(0, tableData.value.length);
    data.forEach((item, index) => {
      tableData.value.push({
        index: index + 1,
        type: item.type === 0 ? '文件' : '目录',
        name: path === '/' ? item.path : path + item.path,
        size: String(item.size),
      });
    });
  } catch (err) {
    console.log('error', err);
  }
};

const handleClick = async (data: Catalog) => {
  console.log(data.name);
  await fetchCatalog(data.name);
  //  let d: Tree[] = await fetchCatalog(data.name);
  // data.children.push(d[0]);
};
</script>
