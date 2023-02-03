<template>
  <div class="gva-table-box">
    <div class="gva-btn-list" style="text-align: justify;">
      <div style="width: 80%; height: 100%">
        <el-button size="small" type="primary" icon="plus" @click="handleAdd">新增</el-button>
        <el-button size="small" type="primary" icon="turnOff" @click="handleBatchStateChange()">批量禁用</el-button>
        <el-button size="small" type="danger" icon="delete" @click="handleBatchDelete()">批量删除</el-button>
      </div>
      <div style="width: 20%; height: 100%">
        <el-input
            clearable
            :prefix-icon="Search"
            size="small"
            v-model="searchKW"
            show-word-limit
            min-width="160"
            minlength="1"
            maxlength="100"
            placeholder="zh_name模糊搜索"
            @keyup.enter.native="toSearch"
        />
      </div>
    </div>
    <el-table
        row-key="id"
        :data="tableData"
        :default-sort="{ prop: 'date', order: 'descending' }"
        @selection-change="handleSelectionChange"
        border style="width: 100%" height="600" >
      <el-table-column type="selection" width="40" />
      <el-table-column fixed sortable prop="create_time" label="创建时间" min-width="80" :show-overflow-tooltip=" true" align="center" />
      <el-table-column prop="id" label="主键" min-width="60" :show-overflow-tooltip=" true" align="center" />
      <el-table-column
          column-key="create_user"
          :filters="nameList"
          :filter-method="filterHandler"
          prop="create_user" label="创建者id" min-width="50" align="center" />
      <el-table-column prop="en_name" label="英文名称" min-width="80" align="center" />
      <el-table-column prop="zh_name" label="中文名称" min-width="80" align="center" />
      <el-table-column prop="info" label="简介" :show-overflow-tooltip=" true" min-width="100" align="center" />
      <el-table-column prop="project_id" label="所属项目id" min-width="80" align="center" />
      <el-table-column label="状态" min-width="40" align="center">
        <template #default="scope">
          <el-switch
              v-model="scope.row.state"
              inline-prompt
              :active-value="1"
              :inactive-value="-1"
              @change="handleStateChange(scope.row, v)"
          />
        </template>
      </el-table-column>
      <el-table-column label="操作" min-width="100" fixed="right" align="center">
<!--        <template #header>-->
<!--          <el-input v-model="search" size="small" placeholder="模糊搜索" />-->
<!--        </template>-->
        <template #default="scope">
          <el-popover :visible="scope.row.visible" placement="top" width="160">
            <p>确定要删除此用户吗</p>
            <div style="text-align: right; margin-top: 8px;">
              <el-button size="small" type="primary" link @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="small"  @click="handleDelete(scope.$index, scope.row)">确定</el-button>
            </div>
            <template #reference>
              <el-button type="danger" link icon="delete" size="small" @click="scope.row.visible = true">删除</el-button>
            </template>
          </el-popover>
          <el-button type="primary" link icon="edit" size="small" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!--分页-->
    <div class="pagination">
      <el-pagination background
          layout="total, sizes, prev, pager, next"
          :page-sizes="[10, 20, 30, 50]"
          :total="total"
          :current-page="currentPage"
          :page-size="pageSize"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
      />
    </div>
    <!-- 弹窗 -->
    <el-dialog v-model="dialogFormVisible" :show-close="false" width="40%" >
<!--    <el-dialog v-model="dialogFormVisible" :title="dialogTitle">-->
      <template #header="{ close, titleId, titleClass }">
        <div style="display: flex;flex-direction: row;justify-content: space-between;align-items: baseline;">
          <h4 :id="titleId" :class="titleClass">{{ dialogTitle }}</h4>
          <div>
            <el-button type="primary" v-show="dialogType === 'add'" @click="beforeReferenc" >参考之前</el-button>
            <el-button :icon="CloseBold" @click="closeDialog" />
          </div>
        </div>
      </template>
      <el-form :model="formData" :rules="rules" label-width="80px" @change="formDataChange">
        <el-form-item label="中文名称" prop="zh_name">
          <el-input v-model="formData.zh_name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="英文名称" prop="en_name">
          <el-input v-model="formData.en_name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="图标" prop="ico" min-width="120">
          <el-input v-model="formData.ico" autocomplete="off" />
        </el-form-item>
        <el-form-item label="简介" prop="info" :show-overflow-tooltip="true" min-width="50">
          <el-input v-model="formData.info" autocomplete="off" />
        </el-form-item>
        <el-form-item label="创建者id" prop="create_user" min-width="100">
          <el-input v-model="formData.create_user" autocomplete="off" maxlength="36" />
        </el-form-item>
        <el-form-item label="需求组ids" prop="demand_ids" min-width="80">
          <el-input v-model="formData.demand_ids" autocomplete="off" />
        </el-form-item>
        <el-form-item label="文档组ids" prop="doc_ids" min-width="80">
          <el-input v-model="formData.doc_ids" autocomplete="off" />
        </el-form-item>
        <el-form-item label="参与者ids" prop="join_users" :show-overflow-tooltip="true" min-width="120">
          <el-input v-model="formData.join_users" autocomplete="off" />
        </el-form-item>
        <el-form-item label="参与组ids" prop="join_groups" min-width="100">
          <el-input v-model="formData.join_groups" autocomplete="off" />
        </el-form-item>
        <el-form-item label="所属项目id" prop="project_id" min-width="100">
          <el-input v-model="formData.project_id" autocomplete="off" />
        </el-form-item>
        <el-form-item label="备注" prop="remark" min-width="100">
          <el-input v-model="formData.remark" autocomplete="off" />
        </el-form-item>
        <el-form-item label="排序" prop="rank" min-width="100">
          <el-input-number v-model="formData.rank" autocomplete="off" />
        </el-form-item>
        <el-form-item label="状态" prop="state" min-width="40">
          <el-switch
              v-model="formData.state"
              inline-prompt
              :active-value="1"
              :inactive-value="-1"
          />
        </el-form-item>
        <el-form-item label="创建时间" prop="create_time" min-width="40" v-if="dialogType === 'edit'" >
          <el-input v-model="formData.create_time" autocomplete="off"/>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "List1"
}
</script>

<script setup>
import { ref } from 'vue'
import axios from "axios";
import {ElMessage} from "element-plus";
import { formatDate } from '@/utils/format'
import { Search,CloseBold } from '@element-plus/icons-vue'

const tableData = ref([])
const searchKW = ref('')
const selectd = ref([])
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(0)
const nameList = ref([])
let searchHistoryKW = ""
let rowHistory = {}

const baseUrl = "http://127.0.0.1:810"
// 获取列表信息
const getList = (val) => {
  console.log("getList",val)
  axios.post(baseUrl+"/admin/Application/v1/gets", val, {
    timeout: 99999,
    headers: {
      'Content-Type': 'application/json; charset=UTF-8',
      'token':'000'
    }
  }).catch(function (error) {
    ElMessage({
      showClose: true,
      message: error,
      type: 'error'
    })
  }).then(function (response) {
    if (!response){
      ElMessage({
        showClose: true,
        message: "数据获取失败",
        type: 'error'
      })
      return
    }
    if (response.status >= 300){
      ElMessage({
        showClose: true,
        message: "数据获取失败",
        type: 'error'
      })
      return
    }
    var res = response.data;
    console.log("res", res)
    if (res.code != 200){
      ElMessage({
        showClose: true,
        message: "数据获取出错,"+res.msg,
        type: 'warning'
      })
      return
    }
    if (!res.data.list){
      tableData.value = []
      ElMessage({
        showClose: true,
        message: "找不到数据",
        type: 'warning'
      })
      return
    }

    const tmp = []
    const tmp_ = []
    res.data.list.forEach((item) => {
      item.create_time = formatDate(item.create_time)
      //人名筛选功能
      if ( !tmp_.includes(item.create_user )){
        tmp.push({ text: item.create_user, value: item.create_user })
        tmp_.push(item.create_user)
      }
    });
    nameList.value = tmp

    tableData.value = res.data.list
    currentPage.value = res.data.current
    const total_ = res.data.current===1? res.data.list.length : (res.data.pageSize*(res.data.current-1))+res.data.list.length
    total.value = res.data.isNext? total_+1 : total_
    pageSize.value = res.data.pageSize
    console.log("total", total_, total.value, currentPage.value, pageSize.value)
  })
}
getList({
  "current": currentPage.value,
  "orderBy": "create_time",
  "pageSize": pageSize.value,
  "query": [],
  "sort": "desc"
})

//搜索
const toSearch = () => {
  console.log("toSearch",searchKW.value, tableData.value.length == 0, tableData.value)
  if(searchKW.value === "") {
    if (tableData.value.length == 0 || searchHistoryKW != ""){
      getList({
        "current": 1,
        "orderBy": "create_time",
        "pageSize": pageSize.value,
        "query": [],
        "sort": "desc"
      })
    }
    searchHistoryKW = searchKW.value
    return;
  } else if (searchHistoryKW == searchKW.value){
    return;
  }else {
    // 请求查询接口，将列表展现出来
    getList({
      "current": 1,
      "orderBy": "create_time",
      "pageSize": pageSize.value,
      "query": [
        {
          "handle": "like",
          "key": "zh_name",
          "nextHandle": "and",
          "val": searchKW.value
        }
      ],
      "sort": "desc"
    })
    searchHistoryKW = searchKW.value
  }
}

//列表多选事件触发处理
const multipleSelection = ref([])
const handleSelectionChange = (val) => {
  console.log("multipleSelection",multipleSelection.value, val)
  multipleSelection.value = val
}
//人名筛选功能
const filterHandler = (value, row, column) => {
  const property = column['property']
  return row[property] === value
}

//add
const add = (val) => {
  if (val.hasOwnProperty("create_user")){
    val.create_user = parseInt(val.create_user)
  }
  axios.post(baseUrl+"/admin/Application/v1/", val, {
    timeout: 99999,
    headers: {
      'Content-Type': 'application/json; charset=UTF-8',
      'token':'000'
    }
  }).catch(function (error) {
    ElMessage({
      showClose: true,
      message: error,
      type: 'error'
    })
  }).then(function (response) {
    if (!response) {
      ElMessage({
        showClose: true,
        message: "请求接口失败",
        type: 'error'
      })
      return
    }
    if (response.status >= 300) {
      ElMessage({
        showClose: true,
        message: "操作失败",
        type: 'error'
      })
      return
    }
    var res = response.data;
    console.log("res", res)
    if (res.code != 200) {
      ElMessage({
        showClose: true,
        message: "操作出错," + res.msg,
        type: 'warning'
      })
      return
    }
    ElMessage({
      showClose: true,
      message: "操作成功",
      type: 'success'
    })
    getList({
      "current": 1,
      "orderBy": "create_time",
      "pageSize": pageSize.value,
      "query": [],
      "sort": "desc"
    })
  })
}
//del
const del = async(val) =>{
  axios.delete(baseUrl+"/admin/Application/v1/", { data:val }, {
    timeout: 99999,
    headers: {
      'Content-Type': 'application/json; charset=UTF-8',
      'token':'000'
    }
  }).catch(function (error) {
    ElMessage({
      showClose: true,
      message: error,
      type: 'error'
    })
  }).then(function (response) {
    if (!response) {
      ElMessage({
        showClose: true,
        message: "请求接口失败",
        type: 'error'
      })
      return
    }
    if (response.status >= 300) {
      ElMessage({
        showClose: true,
        message: "操作失败",
        type: 'error'
      })
      return
    }
    var res = response.data;
    console.log("res", res)
    if (res.code != 200) {
      ElMessage({
        showClose: true,
        message: "操作出错," + res.msg,
        type: 'warning'
      })
      return
    }
    ElMessage({
      showClose: true,
      message: "操作成功",
      type: 'success'
    })
  })
}
//update
const update = (val) =>{
  if (val.hasOwnProperty("create_user")){
    val.create_user = parseInt(val.create_user)
  }
  axios.put(baseUrl+"/admin/Application/v1/", val, {
    timeout: 99999,
    headers: {
      'Content-Type': 'application/json; charset=UTF-8',
      'token':'000'
    }
  }).catch(function (error) {
    ElMessage({
      showClose: true,
      message: error,
      type: 'error'
    })
  }).then(function (response) {
    if (!response) {
      ElMessage({
        showClose: true,
        message: "请求接口失败",
        type: 'error'
      })
      return
    }
    if (response.status >= 300) {
      ElMessage({
        showClose: true,
        message: "操作失败",
        type: 'error'
      })
      return
    }
    var res = response.data;
    console.log("res", res)
    if (res.code != 200) {
      ElMessage({
        showClose: true,
        message: "操作出错," + res.msg,
        type: 'warning'
      })
      return
    }
    ElMessage({
      showClose: true,
      message: "操作成功",
      type: 'success'
    })
  })
}

//状态更改
const handleStateChange = (row) => {
  console.log("handleStateChange", row)
  update({
    "id": row.id,
    "state": row.state,
  })
}
//状态更改
const handleBatchStateChange = () => {
  console.log("handleBatchStateChange",multipleSelection.value)
  multipleSelection.value.forEach((item) => {
    if (item.state != -1){
      update({
        "id": item.id,
        "state": -1,
      })
      item.state = -1
    }
  })
}
//新增
const handleAdd = () => {
  dialogType.value = "add"
  formData.value = {}
  dialogFormVisible.value = true
}
//修改
let editIndex = 0
const handleEdit = (index, row) => {
  console.log("handleEdit", index, row)
  editIndex = index
  dialogType.value = "edit"
  formData.value = JSON.parse(JSON.stringify(row))
  // 历史记录
  rowHistory = JSON.parse(JSON.stringify(row))
  dialogFormVisible.value = true
}
//删除
const handleDelete = async(index, row) => {
  console.log(index, row)
  await del({
        "id": row.id
  })
  getList({
    "current": 1,
    "orderBy": "create_time",
    "pageSize": pageSize.value,
    "query": [],
    "sort": "desc"
  })
}
//批量删除
const handleBatchDelete = async() => {
  console.log("handleBatchDelete", multipleSelection)
  for (var item in multipleSelection.value) {
    await del({
      "id": item.id
    })
  }
  // multipleSelection.value.forEach((item) => {
  //   await del({
  //     "id": item.id
  //   })
  // })
  getList({
    "current": 1,
    "orderBy": "create_time",
    "pageSize": pageSize.value,
    "query": [],
    "sort": "desc"
  })
}

//分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getList({
    "current": currentPage.value,
    "orderBy": "create_time",
    "pageSize": pageSize.value,
    "query": [],
    "sort": "desc"
  })
}
const handleCurrentChange = (val) => {
  currentPage.value = val
  getList({
    "current": currentPage.value,
    "orderBy": "create_time",
    "pageSize": pageSize.value,
    "query": [],
    "sort": "desc"
  })
}

// 弹窗
const dialogType = ref('add')
const dialogTitle = ref('新增角色')
const dialogFormVisible = ref(false)
const formData = ref({})
const mustUint = (rule, value, callback) => {
  if (!/^[0-9]*[1-9][0-9]*$/.test(value)) {
    return callback(new Error('请输入正整数'))
  }
  return callback()
}
const rules = ref({
  // authorityId: [
  //   { required: true, message: '请输入角色ID', trigger: 'blur' },
  //   { validator: mustUint, trigger: 'blur', message: '必须为正整数' }
  // ],
  // authorityName: [
  //   { required: true, message: '请输入角色名', trigger: 'blur' }
  // ],
  // parentId: [
  //   { required: true, message: '请选择父角色', trigger: 'blur' },
  // ],
  create_user: [
    { required: true, message: '请输入创建人id', trigger: 'blur' },
    { validator: mustUint, trigger: 'blur', message: '必须为正整数' }
  ]
})
const enterDialog = () => {
  console.log("enterDialog", formData.value)
  if (dialogType.value == "add"){
    add(formData.value)
  } else if (dialogType.value == "edit"){
    update(formData.value)
  }
  // 历史记录
  rowHistory = JSON.parse(JSON.stringify(formData.value))
  tableData.value[editIndex] = JSON.parse(JSON.stringify(formData.value))
  dialogFormVisible.value = false
}
const closeDialog = () => {
  dialogFormVisible.value = false
}
const beforeReferenc = () => {
  formData.value = rowHistory
}
//xx
const formChangeData = {}
const formDataChange = (env) => {
  console.log("formDataChange",env, env.target.value, env.target.key)
  // formChangeData[]
  // 因为vue生成的原生html代码里面，没有key，所以在事件里面找不到key；要么通过中文名称进行转一下取值，做成代码生成的话，也要解决从key到中文的问题

  var fieldDict = {
    "id":"主键",
    "zh_name":"中文名称",
    "en_name":"英文名称",
    "ico":"图标",
    "info":"简介",
    "create_user":"创建者id",
    "demand_ids":"需求组ids",
    "doc_ids":"文档组ids",
    "join_users":"参与者ids",
    "join_groups":"参与组ids",
    "project_id":"所属项目id",
    "remark":"备注",
    "rank":"排序",
    "state":"状态",
    "create_time":"创建时间",
    "update_time":"更新时间"
  }

  fieldDict


}

</script>

<style lang="scss">
.pagination {
  display: flex;
  justify-content: flex-end;
  .el-pagination {
    padding: 10px 0 0 0 !important;
  }
}
</style>