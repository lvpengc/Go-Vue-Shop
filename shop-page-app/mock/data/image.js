exports.banner = (req, res) => {
    res.status(200).sendFile(__dirname + '/banner/' + req.params.id + '.png')
}
exports.category = (req, res) => {
    res.status(200).sendFile(__dirname + '/icon/category-' + req.params.id + '.png')
}
exports.goods = (req, res) => {
    res.status(200).sendFile(__dirname + '/goods/' + req.params.id + '.png')
}
