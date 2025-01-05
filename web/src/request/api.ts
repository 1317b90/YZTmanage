import service from "@/request/index";


// --------- 任务 --------- 任务 --------- 任务 --------- 任务 --------- 任务 --------- 任务 --------- 任务


// 从数据库获取任务
export async function getTaskDB(state:string) {
    if (state=="all"){
        return service({
            url: "task/db",
            method: "GET",
        })
    }else{
        return service({
        url: "task/db?state="+state,
        method: "GET",
    })}

}

// 创建任务
export async function createTask(Type: string, Input: object) {
    return service({
        url: "task",
        method: "POST",
        data: {
            "Type": Type,
            "Input": Input
        }
    })
}

// 统计任务数据
export async function countTask() {
    return service({
        url: "task/count",
        method: "GET",
    })
}

// --------- redis --------- redis --------- redis --------- redis --------- redis --------- redis --------- redis --------- redis --------- redis

// 获取所有redis数据
export async function getRedis() {
    return service({
        url: "redis",
        method: "GET",
    })
}


// 删除redis数据
export async function delRedis(key: string) {
    return service({
        url: "redis?key=" + key,
        method: "DELETE",
    })
}


// --------- 用户 --------- 用户 --------- 用户 --------- 用户 --------- 用户 --------- 用户 --------- 用户 --------- 用户
// 动态查询用户
export async function getUser(params: { Phone?: string; Enable?: boolean; Uscid?: string; DsjUsername?: string; DsjPassword?: string; CompanyName?: string; BankName?: string; BankID?: string }) {
    const queryString = Object.entries(params)
        .filter(([_, value]) => value !== undefined)
        .map(([key, value]) => `${key}=${value}`)
        .join('&');
    return service({
        url: `user?${queryString}`,
        method: "GET",
    });
}

// 根据手机号查询用户
export async function getUserByPhone(phone: string) {
    return service({
        url: "user/" + phone,
        method: "GET",
    })
}

// 更新用户
export async function putUser(data: object) {
    return service({
        url: "user",
        method: "PUT",
        data: data
    })
}


// 删除用户
export async function delUser(phone: string) {
    return service({
        url: "user?phone=" + phone,
        method: "DELETE",
    })
}

// --------- 文件 --------- 文件 --------- 文件 --------- 文件 --------- 文件 --------- 文件 --------- 文件

// 亿企代账上传文件
export async function upFileBatch(formData:FormData){
    return service({
        url: "upfile/batch",
        method: "POST",
        data: formData,
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });
}

// 下载模板文件
export async function downFileTemplate(head:object){
    return service({
        url: "downfile/template",
        method: "POST",
        data: head
    })
}
// 上传通用文件
export async function upFileCommon(formData:FormData){
    return service({
        url: "upfile/common",
        method: "POST",
        data: formData,
    })
}

// --------- 亿企代账 --------- 亿企代账 --------- 亿企代账 --------- 亿企代账 --------- 亿企代账 --------- 亿企代账 --------- 亿企代账

// 获取亿企代账账号列表
export async function getYqdzAccount() {
    return service({
        url: "yqdz/account",
        method: "GET",
    })
}

// 选择亿企代账账号
export async function useYqdzAccount(id: number) {
    return service({
        url: "yqdz/useAccount?id=" + String(id),
        method: "GET",
    })
}


// --------- 微搭 --------- 微搭 --------- 微搭 --------- 微搭 --------- 微搭 --------- 微搭 --------- 微搭


// 获取微搭数据
export async function getWeda(filter:string){
    return service({
        url: "weda?filter="+filter,
        method: "GET",
    })
}

// --------- 定时任务 --------- 定时任务 --------- 定时任务 --------- 定时任务 --------- 定时任务 --------- 定时任务 --------- 定时任务

// 获取所有定时任务
export async function getCron(){
    return service({
        url: "cron",
        method: "GET",
    })
}

// 创建定时任务
export async function addCron(data: object) {
    return service({
        url: "cron",
        method: "POST",
        data: data
    })
}

// 删除定时任务
export async function delCron(id: string) {
    return service({
        url: "cron?id=" + id,
        method: "DELETE",
    })
}

// --------- 日志 --------- 日志 --------- 日志 --------- 日志 --------- 日志 --------- 日志 --------- 日志

// 获取日志
export async function getLog(params:object){
    return service({
        url: "log",
        method: "GET",
        params: params
    })
}


// --------- 企业微信 --------- 企业微信 --------- 企业微信 --------- 企业微信 --------- 企业微信 --------- 企业微信 --------- 企业微信

// 发送消息
export async function sendWecomMessage(data: object) {
    return service({
        url: "wecom",
        method: "POST",
        data: data
    })
}


// 获取消息队列
export async function getWecomMessageList() {
    return service({
        url: "wecom/list",
        method: "GET",
    })
}

// 删除消息
export async function delWecomMessage(id: string) {
    return service({
        url: "wecom/" + id,
        method: "DELETE",
    })
}

