const hostname = window.location.hostname;
const apiport = 8002;
const prefixpath = '';
//const prefixpath = '/webbox/api'; // if use proxy_pass for apiserver in nginx, set this variable.

const svraddr = 'http://' + hostname + (prefixpath === '' ? (":" + apiport) : prefixpath);

export default { hostname, apiport, svraddr};
