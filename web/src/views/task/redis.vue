<template>
  <div class="cards">
    <div v-if="data.length==0">
      <img src="@/assets/nil.jpg" alt="" style="height:180px">
    </div>
    
    <!-- <div v-else v-for="item in data" class="card">
      <a-tag :bordered="false" color="success" class="card-tag"> {{ item.ID }}</a-tag>
      <a-tag :bordered="false" color="processing" class="card-tag"> {{ item.Key }}</a-tag>

      <p v-if="getJSON(item.Input)==''">
          {{ item.Input }}
        </p>
      <json-viewer v-else :value="getJSON(item.Input)"></json-viewer>
      
      <a-button shape="circle" type="dashed" class="cancelButton" @click="onCancel(item.ID.toString())">
          <template #icon>
            <CloseOutlined />
          </template>
      </a-button>
    </div> -->


    <a-button type="primary" shape="circle" :icon="h(ReloadOutlined)" @click="getData" id="reloadButton" />
  </div>


</template>
<script lang="ts" setup>
import { getRedis,delRedis } from "@/request/api"
import { ref, type Ref } from 'vue';

import { h } from 'vue';
import { ReloadOutlined,CloseOutlined } from '@ant-design/icons-vue';
import JsonViewer from "vue-json-viewer";

import "vue-json-viewer/style.css";
import { message } from "ant-design-vue";
import {formatTime} from "@/func"

// 定义数据类型接口
interface RedisData {
  ID: string;
  Key: string;
  Input: string;
}

// 修改 data 的类型声明
const data: Ref<RedisData[]> = ref([]);

function getData(){
  getRedis().then(res => {
  data.value = []
  res.data.data.forEach((redata: RedisData) => {
    data.value.push(redata)
  });
}).catch(res=>{
  message.error("数据更新失败,"+res)
})
}
getData()

// 格式化json
function getJSON(jsonValue: any) {
  if (typeof jsonValue == "string") {
    try {
      var obj = JSON.parse(jsonValue);
      if (typeof obj == "object" && obj) {
        return obj;
      }
    } catch (e) {
      console.log("error：" + jsonValue + "!!!" + e);
      return "";
    }
  } else {
    return jsonValue;
  }
}


// 删除内存任务
function onCancel(id:string){
  delRedis(id).then(res=>{
    getData()
    message.success("删除成功！")
  }).catch(err=>{
    message.success("删除失败，"+err)
  })
}

</script>

<style scoped>
.cards{
  height: 80vh;
  overflow-y: auto;
}
.card{
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  padding: 10px;
  margin-bottom: 10px;
}
.card-tag{
  font-size: 16px;
}
#reloadButton{
  position: fixed;
  bottom: 30px;
  right: 30px;
  z-index: 1000;
}
.cancelButton{
  margin-top: -40px;
  margin-right: 10px;
  float: right;
}
</style>