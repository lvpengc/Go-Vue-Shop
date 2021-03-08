// 获取本机的局域网IP
var os = require('os'),
    iptable = {},
    ifaces = os.networkInterfaces();
for (var dev in ifaces) {
    ifaces[dev].forEach(function (details, alias) {
        if (details.family == 'IPv4') {
            iptable[dev + (alias ? ':' + alias : '')] = details.address;
        }
    });
}

exports.getLocalHost = function () {
    for (const key in iptable) {
        if (key.startsWith('WLAN')) {
            return iptable[key]
        }
    }
    return ""
}
