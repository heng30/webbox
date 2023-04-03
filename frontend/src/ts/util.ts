
const PrettyFileSize = (bytes: number) => {
    const K = 1024;
    const M = 1024 * K;
    const G = 1024 * M;

    if (bytes < K) {
        return `${bytes}B`
    } else if (bytes < M) {
        return (bytes/K).toFixed(2) + 'K';
    } else if (bytes < G) {
        return (bytes/M).toFixed(2) + 'M';
    } else {
        return (bytes/G).toFixed(2) + 'G';
    }
}

export default { PrettyFileSize }
