const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');

const app = express();
app.use(bodyParser.urlencoded({
    extended: true
}));
app.use(bodyParser.json());
app.use(cors());

require('./router')(app)

//  指定端口
const PORT = 8080;

const utils = require('./data/utils')

app.listen(PORT, () => {
    const host = utils.getLocalHost()
    console.log(`服务器启动，运行为http://${host}:${PORT}`);
});
