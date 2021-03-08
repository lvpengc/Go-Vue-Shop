// 去除字符串两端特殊字符
const trim = (str, char) => {
    str = str ? str : ""
    char = char ? char : " "
    if (str == "") {
        return ""
    }
    const lchar = char.length
    const lstr = str.length
    let s = false
    let e = false
    if (str.indexOf(char) === 0) {
        s = true
    }
    if (str.lastIndexOf(char) === lstr - lchar) {
        e = true
    }
    var r = str
    if (e) {
        if (r.length > 0) {
            r = r.substring(0, lstr - lchar)
            return trim(r, char)
        } else {
            return ""
        }
    }
    if (s) {
        if (r.length > 0) {
            r = r.substring(lchar, r.length)
            return trim(r, char)
        } else {
            return ""
        }
    }
    return r
}

export {
    trim
}
