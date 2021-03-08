const Mock = require('mockjs');
const utils = require('./utils');
const host = utils.getLocalHost();

const res_data = {
    code: 200,
    msg: 'success',
    data: []
}

let goods_list = [{
        id: 1,
        goods_id: 1,
        goods_stock_id: 1,
        is_checked: true,
        pic_url: 'http://' + host + ':8080/image/goods/1?w=80&h=80',
        goods_name: '123123123123123123',
        goods_stock: '已选择的规格描述',
        number: 2,
        price: 99.9
    },
    {
        id: 2,
        goods_id: 1,
        goods_stock_id: 2,
        is_checked: true,
        pic_url: 'http://' + host + ':8080/image/goods/1?w=80&h=80',
        goods_name: '456',
        goods_stock: '已选择的规格描述',
        number: 1,
        price: 99.9
    }
]

// 获取购物车的商品数据
exports.getGoods = (req, res) => {
    let tmp = Object.assign({}, res_data)
    let tmp_fee = 0.0
    goods_list.forEach(function (item) {
        console.log(item)
        if (item.is_checked) {
            tmp_fee += item.number * item.price
        }
    })
    tmp.data = Mock.mock({
        goods_list: goods_list,
        total_fee: parseFloat(tmp_fee).toFixed(2)
    })
    res.status(200).json(tmp)
};

// 添加指定商品到购物车
exports.addGoods = (req, res) => {
    console.log(req)
    const tmp = Object.assign({}, res_data)
    res.status(200).json(tmp)
}

// 更新指定id的库存
exports.updateGoodsStock = (req, res) => {
    goods_list = goods_list.map(item => {
        if (item.id == req.body.id) {
            item.goods_stock_id = req.body.goods_stock_id
        }
        return item
    })
    const tmp = Object.assign({}, res_data)
    res.status(200).json(tmp)
}

// 更新指定id数量
exports.updateGoodsNum = (req, res) => {
    goods_list = goods_list.map(item => {
        if (item.id == req.body.id) {
            item.number = req.body.num
        }
        return item
    })
    const tmp = Object.assign({}, res_data)
    res.status(200).json(tmp)
}

// 删除指定商品
exports.delCartGoods = (req, res) => {
    let arr = goods_list.filter(item => {
        return item.id != req.body.id
    })
    goods_list = arr
    const tmp = Object.assign({}, res_data)
    res.status(200).json(tmp)
}

// 选中或取消指定id
exports.switchCartGoodsItem = (req, res) => {
    goods_list = goods_list.map(item => {
        if (item.id == req.body.id) {
            item.is_checked = req.body.is_checked
        }
        return item
    })
    const tmp = Object.assign({}, res_data)
    res.status(200).json(tmp)
}

// 全选或取消所有商品
exports.switchCartGoodsAll = (req, res) => {
    goods_list = goods_list.map(item => {
        item.is_checked = req.body.is_checked
        return item
    })
    const tmp = Object.assign({}, res_data)
    res.status(200).json(tmp)
}

// 根据选中商品建立订单
exports.checkoutCartGoods = (req, res) => {
    goods_list = goods_list.map(item => {
        item.is_checked = req.body.is_checked
        return item
    })
    const tmp = Object.assign({}, res_data)
    res.status(200).json(tmp)
}
