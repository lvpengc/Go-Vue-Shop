const Mock = require('mockjs');
const utils = require('./utils');
const host = utils.getLocalHost();

let res_data = {
    code: 200,
    msg: 'success',
    data: []
}

exports.getCommonQuestion = (req, res) => {
    let tmp = Object.assign({}, res_data)
    tmp.data = Mock.mock([{
            subject: '运费怎么算？',
            answer: '满99包邮，不满则邮费8元。'
        },
        {
            subject: '快递是哪家？',
            answer: '默认使用顺丰发货，特殊要求请订单备注。'
        },
        {
            subject: '退货和退款？？',
            answer: '随时支持退货，款项将原路返回。'
        },
    ])
    res.status(200).json(tmp)
}
