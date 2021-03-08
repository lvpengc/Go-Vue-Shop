const Mock = require('mockjs');

const utils = require('./utils');
const host = utils.getLocalHost();

let res_data = {
    code: 200,
    msg: 'success',
    data: []
}

// 获取 banner 数据
exports.getBannerData = (req, res) => {
    let tmp = Object.assign({}, res_data)
    tmp.data = Mock.mock([{
            id: 1,
            img_url: 'http://' + host + ':8080/image/banner/1?w=450&h=300',
        },
        {
            id: 2,
            img_url: 'http://' + host + ':8080/image/banner/2?w=450&h=300',
        },
        {
            id: 3,
            img_url: 'http://' + host + ':8080/image/banner/3?w=450&h=300',
        },
    ])
    res.status(200).json(tmp)
};

// 获取 类目 数据
exports.getCategoryData = (req, res) => {
    let tmp = Object.assign({}, res_data)
    tmp.data = Mock.mock([{
            id: 1,
            icon_url: 'http://' + host + ':8080/image/category/1?w=30&h=30',
            name: '分类一'
        },
        {
            id: 2,
            icon_url: 'http://' + host + ':8080/image/category/2?w=30&h=30',
            name: '分类二'
        },
        {
            id: 3,
            icon_url: 'http://' + host + ':8080/image/category/3?w=30&h=30',
            name: '分类三'
        },
        {
            id: 4,
            icon_url: 'http://' + host + ':8080/image/category/4?w=30&h=30',
            name: '分类四'
        },
    ])
    res.status(200).json(tmp)
};

exports.getCouponData = (req, res) => {
    let tmp = Object.assign({}, res_data)
    tmp.data = Mock.mock([{
        id: 1,
        amount: 5.00,
        name: '限时满减券',
        consum_min: '99',
        limit_msg: '仅限品类一',
        days: '5',
    }, {
        id: 2,
        amount: 10.00,
        name: '注册满减券',
        consum_min: '11',
        limit_msg: '全场通用',
        days: '10',
    }])
    res.status(200).json(tmp)
}

exports.getGoodsData = (req, res) => {
    let tmp = Object.assign({}, res_data)
    tmp.data = Mock.mock([{
        id: 1,
        pic_url: 'http://' + host + ':8080/image/goods/1?w=80&h=80',
        price: 98.0,
        name: '竹筒酒',
        brief: '入口绵甜温和，竹味浓香，自然清纯，舌尖留味无穷，尽情回味入口绵甜温和，竹味浓香，自然清纯，舌尖留味无穷，尽情回味'
    }, {
        id: 2,
        pic_url: 'http://' + host + ':8080/image/goods/2?w=80&h=80',
        price: 58.0,
        name: '腊豆腐',
        brief: '土家腊豆腐，黄豆磨制，柴火烟熏特色'
    }, {
        id: 3,
        pic_url: 'http://' + host + ':8080/image/goods/3?w=80&h=80',
        price: 58.0,
        name: '湘西腊肠',
        brief: '风味鲜美，醇厚浓郁，回味绵长，越嚼越香'
    }])
    res.status(200).json(tmp)

}
