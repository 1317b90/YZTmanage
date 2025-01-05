export const menuDict = [
    {
        type: "make_invoice",
        name: "开票",
        input: [
            { "var": "uscid", "name": "税号", required: true, default: "" },
            { "var": "dsj_username", "name": "用户名", required: true, default: "" },
            { "var": "dsj_password", "name": "密码", required: true, default: "" },
            { "var": "buy_name", "name": "购买方名称", required: true, default: "" },
            { "var": "invoice_name", "name": "项目名称", required: true, default: "" },
            { "var": "invoice_amount", "name": "金额", required: true, default: "" },
            { "var": "invoice_type", "name": "发票类型", required: true, default: "普通发票" },
            { "var": "buy_id", "name": "购买方ID", required: false, default: "" },
            { "var": "buy_address", "name": "购买方地址", required: false, default: "" },
            { "var": "buy_phone", "name": "购买方电话", required: false, default: "" },
            { "var": "buy_bank_name", "name": "购买方银行名称", required: false, default: "" },
            { "var": "buy_bank_id", "name": "购买方银行账号", required: false, default: "" },
            { "var": "sell_bank_name", "name": "销售方银行名称", required: false, default: "" },
            { "var": "sell_bank_id", "name": "销售方银行账号", required: false, default: "" },
            { "var": "invoice_model", "name": "规格型号", required: false, default: "" },
            { "var": "invoice_unit", "name": "单位", required: false, default: "" },
            { "var": "invoice_num", "name": "数量", required: false, default: "" },
            { "var": "invoice_price", "name": "单价", required: false, default: "" },
            { "var": "invoice_code", "name": "项目编码", required: false, default: "" },
            { "var": "is_preview", "name": "是否预览发票", required: true, default: "true" },

            { "var": "wecome_id", "name": "企业微信ID", required: false, default: "" },
            { "var": "task_id", "name": "任务ID", required: false, default: "" },
            { "var": "buy_email", "name": "购买方邮箱", required: false, default: "" },
            { "var": "company_name", "name": "销售方公司名称", required: false, default: "" },
        ]
        }
]