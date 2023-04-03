<template>
  <div style="overflow-y: scroll; height: calc(100vh - 50px)">
    <el-table
      :data="filterTableData"
      @row-click="handleClick"
      style="width: 100%"
    >
      <el-table-column
        label="名称"
        prop="name"
        sortable
        :sort-method="sortByName"
      />
      <el-table-column
        label="类型"
        prop="type"
        width="80"
        sortable
        :sort-method="sortByType"
      />
      <el-table-column
        label="大小"
        prop="size"
        width="100"
        sortable
        :sort-method="sortBySize"
      />
      <el-table-column align="right" width="300">
        <template #header>
          <el-input
            v-model="search"
            size="small"
            placeholder="Type to search"
          />
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
  </div>
</template>

<script lang="ts" setup>
import config from '../ts/config';
import util from '../ts/util';
import store from '../ts/store';
import { onMounted, computed, ref, watch } from 'vue';

interface Catalog {
  type: string;
  name: string;
  size: string;
  path: string;
  sizeBytes: number;
}

type Dictionary = {
  [key: string]: Catalog[];
};

const tableData = ref([]);
const search = ref('');
const catalogCache: Dictionary = {};

const filterTableData = computed(() =>
  tableData.value.filter(
    (data) =>
      !search.value ||
      data.name.toLowerCase().includes(search.value.toLowerCase())
  )
);

watch(store.CurrentPath, async (newp) => {
  let path = '';
  if (newp.length === 1) {
    path = '/';
  } else {
    path = String(newp.join('/')).substring(1);
  }
  tableData.value = catalogCache[path];
});

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
    data.sort((a, b) => {
      if (a.type === b.type) {
        return a.path.toLowerCase().localeCompare(b.path.toLowerCase());
      } else {
        return b.type - a.type;
      }
    });

    tableData.value = [];
    data.forEach((item) => {
      tableData.value.push({
        type: item.type === 0 ? '文件' : '目录',
        name: item.path,
        path: path === '/' ? `/${item.path}` : `${path}/${item.path}`,
        size: item.type === 0 ? util.PrettyFileSize(item.size) : '--',
        sizeBytes: item.size,
      });
    });

    if (path !== '/') {
      store.CurrentPath.value = path.split('/');
    }
    store.CurrentPath.value[0] = '/';
    catalogCache[path] = tableData.value;
  } catch (err) {
    console.log('Error:', err);
  }
};

const handleClick = async (row: Catalog) => {
  if (row.type === '文件') return;
  if (!!catalogCache[row.path]) {
    store.CurrentPath.value = row.path.split('/');
    store.CurrentPath.value[0] = '/';
    return;
  }

  await updateCatalog(row.path);
};

const sortByName = (a: Catalog, b: Catalog) => {
  return a.name.toLowerCase().localeCompare(b.name.toLowerCase());
};

const sortByType = (a: Catalog, b: Catalog) => {
  if (a.type === b.type) {
    return sortByName(a, b);
  } else {
    return b.type.localeCompare(a.type);
  }
};

const sortBySize = (a: Catalog, b: Catalog) => {
  if (a.type === b.type) {
    return a.sizeBytes - b.sizeBytes;
  } else {
    return sortByType(a, b);
  }
};
</script>
