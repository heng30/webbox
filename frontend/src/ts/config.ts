const hostname = window.location.hostname;
const apiport = 8002;
const prefixpath = ''; // if use proxy_pass in nginx, set this variable.
const svraddr = `http://${hostname}:${apiport}${prefixpath}`

export default { hostname, apiport, svraddr};
