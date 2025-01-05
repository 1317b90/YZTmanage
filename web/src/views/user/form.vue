<template>
    <a-form :model="data" name="basic" :label-col="{ span: 6 }" @finish="onFinish">
        <a-form-item label="手机号" name="Phone" :rules="[{ required: true, message: '请输入对应手机号' }]">
            <a-input v-model:value="data.Phone"  />
        </a-form-item>

        <a-form-item label="公司名称" name="CompanyName" :rules="[{ required: true, message: '请输入对应公司名称' }]">
            <a-input v-model:value="data.CompanyName"  />
        </a-form-item>

        <a-form-item label="社会信用代码" name="Uscid" :rules="[{ required: true, message: '请输入对应社会信用代码' }]">
            <a-input v-model:value="data.Uscid"  />
        </a-form-item>

        <a-form-item label="电税局账号" name="DsjUsername" :rules="[{ required: true, message: '请输入对应电税局账号' }]">
            <a-input v-model:value="data.DsjUsername"  />
        </a-form-item>

        <a-form-item label="电税局密码" name="DsjPassword" :rules="[{ required: true, message: '请输入对应电税局密码' }]">
            <a-input v-model:value="data.DsjPassword"  />
        </a-form-item>

        <a-form-item label="开户银行" name="BankName" :rules="[{ required: true, message: '请输入对应开户银行' }]">
            <a-input v-model:value="data.BankName" />
        </a-form-item>

        <a-form-item label="银行账号" name="BankID" :rules="[{ required: true, message: '请输入对应银行账号' }]">
            <a-input v-model:value="data.BankID" />
        </a-form-item>

        <a-form-item label="是否启用" name="Enable">
            <a-switch v-model:checked="data.Enable" />
        </a-form-item>

        <a-form-item style="text-align: right;">
            <a-button type="primary" html-type="submit">提交</a-button>
        </a-form-item>
    </a-form>
</template>
<script lang="ts" setup>
import { putUser } from "@/request/api"
import { message } from 'ant-design-vue';

const prop = defineProps(['data', 'RPANameDict', 'RPAGroupDict', 'getData'])


const onFinish = (values: any) => {
    console.log(values)
    putUser(values).then(res => {
        prop.getData()
        message.success('提交成功！');
    }).catch(err => {
        message.error(err)
    })

};

</script>
