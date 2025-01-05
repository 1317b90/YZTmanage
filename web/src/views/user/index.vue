<template>

  <a-table class="varTable" :columns="columns"  :rowKey="(record: userI) => record.Phone" :data-source="data" style="height: 70vh;" :pagination="{ pageSize: 10 }">
    <template #headerCell="{ column }">
      <template v-if="column.dataIndex === 'operation'">
        <a-button class="editable-add-btn" @click="onAddVar" type="link">添加</a-button>
      </template>
    </template>

    <template #bodyCell="{ column, record}">
      <template v-if="column.dataIndex === 'Enable'">
        <a-switch v-model:checked="record.Enable"  size="small" disabled/>
      </template>
 
      <template v-if="column.dataIndex === 'operation'">
        <a-button type="primary" size="small" :icon="h(FormOutlined)" @click="onEdit(record)" />
        <a-popconfirm title="确定删除?" placement="leftTop" ok-text="确定" cancel-text="取消" @confirm="onDel(record)">

          <a-button size="small" :icon="h(DeleteOutlined)" class="delButton" />
        </a-popconfirm>
      </template>
    </template>
  </a-table>

  <!-- 添加v-if 让每一次打开都重新初始化 -->

  <a-modal v-if="isModal" v-model:open="isModal" title="Title">
    <VarForm :data="modalData"  :RPANameDict="RPANameDict" :RPAGroupDict="RPAGroupDict" :getData="getData" />
    <!-- 去除默认的地步按钮 -->
    <template #footer>
    </template>
  </a-modal>

</template>
<script lang="ts" setup>
import { message, type TableColumnType } from 'ant-design-vue';
import { onMounted, ref, type Ref } from 'vue';
import { h } from 'vue';
import { DeleteOutlined, FormOutlined } from '@ant-design/icons-vue';
import { getUser,getUserByPhone,putUser,delUser } from "@/request/api";
import VarForm from './form.vue';
import type { filterI, userI } from '@/interface';

// 状态和引用
const data: Ref<userI[]> = ref([]);
const RPANameDict: Ref<filterI[]> = ref([]);
const RPAGroupDict: Ref<filterI[]> = ref([]);
const isModal = ref(false);
const modalData = ref<userI>();

// 表格列定义
const columns: TableColumnType<userI>[] = [
  { title: '手机号', dataIndex: 'Phone', width: 50, align: "center" },
  { title: '公司名称', dataIndex: 'CompanyName', align: "center" },
  { title: '社会信用代码', dataIndex: 'Uscid', align: "center" },
  { title: '电税局账号', dataIndex: 'DsjUsername', align: "center" },
  { title: '电税局密码', dataIndex: 'DsjPassword', align: "center" },
  { title: '开户银行', dataIndex: 'BankName', align: "center" },
  { title: '银行卡号', dataIndex: 'BankID', align: "center" },
  { title: '是否启用', dataIndex: 'Enable', align: "center" },
  { title: '操作', dataIndex: 'operation', fixed: 'right', width: 100, align: "center" },
];

// 获取所有的变量
const getData = () => {
  getUser({}).then(res => {
    data.value = res.data.data;
    isModal.value = false;
  });
};

// 点击新增变量
const onAddVar = () => {
  isModal.value = true;
  modalData.value = {
    Phone: "",
    Enable: true,
    Uscid: "",
    DsjUsername: "",
    DsjPassword: "",
    CompanyName: "",
    BankName: "",
    BankID: ""
  }
};

// 点击编辑变量
const onEdit = (rowData: userI) => {
  modalData.value = rowData;
  isModal.value = true;
};

// 点击删除变量
const onDel = (rowData: userI) => {
  delUser(rowData.Phone || "").then(() => {
    getData();
    message.success("删除成功！");
  }).catch(err => {
    message.error(err);
  });
};




// 初始化
onMounted(() => {
  getData()
})


</script>

<style scoped>
.delButton {
  margin-left: 10px;
}
</style>