module.exports = (app) => {
    const home = require('./data/home');
    const image = require('./data/image');
    const category = require('./data/category');
    const cart = require('./data/cart');
    const goods = require('./data/goods');
    const common = require('./data/common');

    // 首页
    app.get('/home/banner', home.getBannerData);
    app.get('/home/category', home.getCategoryData);
    app.get('/home/coupon', home.getCouponData);
    app.get('/home/goods', home.getGoodsData);
    // 类目列表
    app.get('/category', category.getCategoryData);
    app.get('/category/:id', category.getGoodsData);

    // 图片资源
    app.get('/image/banner/:id', image.banner);
    app.get('/image/category/:id', image.category);
    app.get('/image/goods/:id', image.goods);

    // 购物车
    app.get('/cart', cart.getGoods);
    app.post('/cart', cart.addGoods);
    app.put('/cart/:id/stock', cart.updateGoodsStock);
    app.put('/cart/:id/num', cart.updateGoodsNum);
    app.delete('/cart', cart.delCartGoods);
    app.put('/cart/:id/switch-checked', cart.switchCartGoodsItem);
    app.put('/cart/switch-checked', cart.switchCartGoodsAll);
    app.post('/cart/checkout', cart.checkoutCartGoods);

    // 商品
    app.get('/goods/:id', goods.getGoodsDetail);

    // 公共
    app.get('/common/question', common.getCommonQuestion)
};
