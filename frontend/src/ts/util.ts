
const PrettyFileSize = (bytes: Number) => {
    const K = 1024;
    const M = 1024 * K;
    const G = 1024 * M;

    if (bytes < K) {
        return `${bytes}B`
    } else if (bytes < M) {
        return `${bytes}/${K}}K`
    } else if (bytes < G) {
        return `${bytes}/${M}M`
    } else {
        return `${bytes}/${G}G`
    }
}

export default { PrettyFileSize }
