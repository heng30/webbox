import {ref, reactive} from 'vue';

const CurrentPath = ref([])
const AuthToken = ref('')
const UploadProgress = reactive({
    uploadedBytes: 0,
    totalBytes: 0,
})



export default {
    CurrentPath,
    AuthToken,
    UploadProgress
}
