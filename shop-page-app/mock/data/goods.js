const Mock = require('mockjs');
const utils = require('./utils');
const host = utils.getLocalHost();

let res_data = {
    code: 200,
    msg: 'success',
    data: []
}

exports.getGoodsDetail = (req, res) => {
    let tmp = Object.assign({}, res_data)
    tmp.data = Mock.mock({
        id: 1,
        gallery: [
            'http://' + host + ':8080/image/goods/1?w=80&h=80',
            'http://' + host + ':8080/image/goods/2?w=80&h=80',
            'http://' + host + ':8080/image/goods/3?w=80&h=80',
        ],
        price: 98.0,
        marked_price: 198.0,
        name: '竹筒酒竹筒酒',
        brief: '入口绵甜温和，竹味浓香，自然清纯，舌尖留味无穷，尽情回味',
        detail: '<p>哈哈</p><img width="100%" src=\"http://' + host + ':8080/image/goods/1?w=80&h=80\" _src=\"http://yanxuan.nosdn.127.net/2597f9e2e41093f50761837eb4c2e6be.jpg\" style=\"\"/>',
        // 选择n条最先评论默认显示
        comment: [{
            user_id: 1,
            avatar: 'http://' + host + ':8080/image/goods/1?w=80&h=80',
            nickname: '张三',
            star: 4,
            content: '挺好挺好挺好挺好挺好挺好挺好挺好挺好，挺好挺好挺好挺好挺好挺好挺好挺好挺好，挺好挺好挺好挺好挺好挺好挺好挺好挺好',
            pic_urls: [
                'http://' + host + ':8080/image/goods/1?w=80&h=80',
                'http://' + host + ':8080/image/goods/2?w=80&h=80',
                'http://' + host + ':8080/image/goods/3?w=80&h=80',
                'http://' + host + ':8080/image/goods/1?w=80&h=80',
                'http://' + host + ':8080/image/goods/2?w=80&h=80',
                'http://' + host + ':8080/image/goods/3?w=80&h=80',
            ]
        }]
    })
    res.status(200).json(tmp)
}
