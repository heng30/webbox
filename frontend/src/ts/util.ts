
const PrettyFileSize = (bytes: number) => {
    const K = 1024;
    const M = 1024 * K;
    const G = 1024 * M;

    if (bytes < K) {
        return `${bytes}B`
    } else if (bytes < M) {
        return (bytes/K).toFixed(0) + 'K';
    } else if (bytes < G) {
        return (bytes/M).toFixed(0) + 'M';
    } else {
        return (bytes/G).toFixed(2) + 'G';
    }
}

const  GetParentPath = (path: string) => {
  let index = path.lastIndexOf('/')
  if ( index < 0) {
    return path;
  }
  if (index === 0) return '/'
  return path.substring(0, index)
}

export default { PrettyFileSize, GetParentPath }
