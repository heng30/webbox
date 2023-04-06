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
          <el-input v-model="search" placeholder="搜索" clearable />
        </template>
        <template #default="scope">
          <FileDialog
            @accept="handleUpload"
            v-if="scope.row.type === '目录'"
            buttonLabel="上传"
            :attach="scope.row"
          >
          </FileDialog>
          <el-button type="success" v-if="scope.row.type === '文件'">
            <el-link
              :underline="false"
              :href="downloadUrl(scope.row)"
              style="color: white"
              target="_blank"
              >下载</el-link
            >
          </el-button>

          <div style="margin: 0 10px; display: inline-block">
            <el-button
              v-if="canDelete"
              type="danger"
              @click="handleDelete(scope.$index, scope.row)"
              >删除</el-button
            >
          </div>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script lang="ts" setup>
import config from '../ts/config';
import util from '../ts/util';
import store from '../ts/store';
import FileDialog from './cbase/FileDialog.vue';
import { onMounted, computed, ref, watch } from 'vue';
import { v4 as uuidv4 } from 'uuid';
import { ElMessage } from 'element-plus';

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
const canDelete = ref(false);
const authtoken = store.AuthToken;

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

watch(store.AuthToken, async (newp) => {
  if (newp.length > 0) {
    await updateConfig();
    await updateCatalog('/');
  }
});

onMounted(async () => {
  /*
  await updateConfig();
  await updateCatalog('/');
  */
});

const handleUpload = async (file: any, row: Catalog) => {
  const fileSize = file.size;
  const chunkSize = 1024 * 1024;
  let startByte: number = 0;
  let chunkNumber: number = 0;

  const uuid = uuidv4();
  const uploadChunk = async (
    endByte: number,
    chunkNumber: number,
    isEndChunk: boolean
  ) => {
    const totalPart = isEndChunk ? chunkNumber : -1;
    const url = `${config.svraddr}/mupload?authtoken=${authtoken.value}&&uuid=${uuid}&&partindex=${chunkNumber}&&totalpart=${totalPart}&&filename=${row.path}/${file.name}`;

    const chunk = file.slice(startByte, endByte);
    const formData = new FormData();
    formData.append('chunkNumber', String(chunkNumber));
    formData.append('file', chunk);
    try {
      const resp = await fetch(url, { method: 'POST', body: formData });
      const jdata = await resp.json();
      if (jdata.code !== 0) {
        throw new Error(jdata.data);
      }

      store.UploadProgress.uploadedBytes += endByte - startByte;
      startByte = endByte;
      return true;
    } catch (err) {
      ElMessage.error(`${err}`);
      return false;
    }
  };

  store.UploadProgress.totalBytes += fileSize;

  let ret = true;
  while (startByte < fileSize) {
    const endByte = Math.min(startByte + chunkSize, fileSize);
    ret = await uploadChunk(endByte, ++chunkNumber, endByte == fileSize);
    if (!ret) {
      store.UploadProgress.uploadedBytes -= startByte;
      store.UploadProgress.totalBytes -= fileSize;
      break;
    }
  }

  if (ret) {
    ElMessage({
      message: `上传${file.name}成功!`,
      type: 'success',
    });
  }

  if (store.UploadProgress.uploadedBytes === store.UploadProgress.totalBytes) {
    store.UploadProgress.uploadedBytes = 0;
    store.UploadProgress.totalBytes = 0;
  }

  updateCatalog(row.path, false);
};

const handleDelete = async (_index: number, row: Catalog) => {
  let url: string = '';
  if (row.type == '目录') {
    url = `${config.svraddr}/deldir?authtoken=${authtoken.value}&&dirname=${row.path}`;
  } else {
    url = `${config.svraddr}/delfile?authtoken=${authtoken.value}&&filename=${row.path}`;
  }

  try {
    const res = await fetch(url);
    let jdata = await res.json();
    if (jdata.code !== 0) {
      throw new Error(jdata.data);
    }

    const path = util.GetParentPath(row.path);
    updateCatalog(path);

    ElMessage({
      message: `删除${row.name}成功!`,
      type: 'success',
    });
  } catch (err) {
    ElMessage.error(`${err}`);
  }
};

const updateConfig = async () => {
  const url = `${config.svraddr}/configure?authtoken=${authtoken.value}`;
  try {
    const res = await fetch(url);
    let jdata = await res.json();
    if (jdata.code !== 0) {
      throw new Error(jdata.data);
    }

    canDelete.value = jdata.data.canDelete;
  } catch (err) {
    ElMessage.error(`${err}`);
  }
};

const updateCatalog = async (path: string, isEnterPath: boolean = true) => {
  const url = `${config.svraddr}/readdir?authtoken=${authtoken.value}&&dirname=${path}`;
  try {
    const res = await fetch(url);
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

    if (isEnterPath) {
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
    } else {
      let td = [];
      data.forEach((item) => {
        td.push({
          type: item.type === 0 ? '文件' : '目录',
          name: item.path,
          path: path === '/' ? `/${item.path}` : `${path}/${item.path}`,
          size: item.type === 0 ? util.PrettyFileSize(item.size) : '--',
          sizeBytes: item.size,
        });
      });
      catalogCache[path] = td;
    }
  } catch (err) {
    ElMessage.error(`${err}`);
  }
};

const handleClick = async (row: Catalog, column: any) => {
  if (row.type === '文件' || column.no === 3) return;
  search.value = '';
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

const downloadUrl = (row: Catalog) => {
  return `${config.svraddr}/download?authtoken=${authtoken.value}&&filename=${row.path}`;
};

const Mkdir = async (pdir: string, dirname: string) => {
  const url = `${config.svraddr}/mkdir?authtoken=${authtoken.value}&&dirname=${dirname}`;

  try {
    const resp = await fetch(url);
    const jdata = await resp.json();
    if (jdata.code !== 0) {
      throw new Error(jdata.data);
    }

    ElMessage({
      message: `创建${dirname}成功!`,
      type: 'success',
    });

    await updateCatalog(pdir);
  } catch (err) {
    ElMessage.error(`${err}`);
  }
};
</script>
