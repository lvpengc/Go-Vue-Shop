// Package casbin model.go 定义casbin的Models
package casbin

/** 声明基于角色的访问控制
 * matchers：
 * g(r.sub, p.sub) == true 判断当前请求用户所归属的角色是否和请求资源需要的角色一致
 * keyMatch2(r.obj, p.obj) 根据 restful 规范判断 url
 * regexMatch(r.act, p.act) 正则匹配提交的方式是否在范围内，如 GET|PUT|POST|DELETE
 */
const casbinModel = `[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) == true \
        && keyMatch2(r.obj, p.obj) == true \
        && regexMatch(r.act, p.act) == true`
