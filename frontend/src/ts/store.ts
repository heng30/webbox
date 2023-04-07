import {ref, reactive} from 'vue';

const CurrentPath = ref([])
const AuthToken = ref('')
const IsShowMkdirDialog = ref(false);
const IsShowRenameDialog = ref(false);
const UploadProgress = reactive({
    uploadedBytes: 0,
    totalBytes: 0,
})

let GlobalFunc = { }

export default {
    CurrentPath,
    AuthToken,
    UploadProgress,
    IsShowMkdirDialog,
    IsShowRenameDialog,

    GlobalFunc,
}
