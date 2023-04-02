const hostname = window.location.hostname;
const apiport = 8002;
const svraddr = `http://${hostname}:${apiport}`;

export default { hostname, apiport, svraddr };
