<template>
  <el-table
    :data="filterTableData"
    @row-click="handleClick"
    style="width: 100%"
  >
    <el-table-column label="名称" prop="name" />
    <el-table-column label="类型" prop="type" width="80" />
    <el-table-column label="大小" prop="size" width="100" />
    <el-table-column align="right" width="300">
      <template #header>
        <el-input v-model="search" size="small" placeholder="Type to search" />
      </template>
      <template #default="scope">
        <el-button
          size="small"
          type="primary"
          @click="handleUpload(scope.$index, scope.row)"
          >上传</el-button
        >
        <el-button
          size="small"
          type="primary"
          @click="handleDownload(scope.$index, scope.row)"
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
import config from '../ts/config';
import util from '../ts/util';
import { onMounted, computed, ref } from 'vue';

interface Catalog {
  type: string;
  name: string;
  size: string;
  path: string;
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
  await updateCatalog('/');
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

const updateCatalog = async (path: string) => {
  const url = `${config.svraddr}/readdir?dirname=${path}`;
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

    const data: any[] = jdata.data;
    tableData.value = [];
    data.forEach((item) => {
      tableData.value.push({
        type: item.type === 0 ? '文件' : '目录',
        name: item.path,
        path: path === '/' ? `/${item.path}` : `${path}/${item.path}`,
        size: item.type === 0 ? util.PrettyFileSize(item.size) : '--',
      });
    });
  } catch (err) {
    console.log('Error:', err);
  }
};

const handleClick = async (row: Catalog) => {
  if (row.type === '文件') return;
  await updateCatalog(row.path);
};
</script>
