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
      <!--机器生成，请勿修改-->
      <el-table-column prop="id" label="主键" min-width="60" :show-overflow-tooltip=" true" align="center"/>
      <el-table-column
          column-key="create_user"
          :filters="nameList"
          :filter-method="filterHandler"
          prop="create_user" label="所属用户" min-width="80" align="center" />
      <el-table-column prop="application_id" label="应用id" min-width="80" align="center" />
      <el-table-column prop="version" label="版本号" min-width="100" align="center" :show-overflow-tooltip=" true"/>
      <el-table-column prop="config_ids" label="配置ids" min-width="100" align="center" :show-overflow-tooltip=" true"/>
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
      <!--机器生成，请勿修改-->
      <el-table-column label="操作" min-width="100" fixed="right" align="center">
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
      <template #header="{ close, titleId, titleClass }">
        <div style="display: flex;flex-direction: row;justify-content: space-between;align-items: baseline;">
          <h4 :id="titleId" :class="titleClass">{{ dialogTitle }}</h4>
          <div>
            <el-button type="primary" v-show="dialogType === 'add'" @click="beforeReferenc" >参考之前</el-button>
            <el-button :icon="CloseBold" @click="closeDialog" />
          </div>
        </div>
      </template>
      <el-form :model="formData" :rules="rules" label-width="80px">
          <!--机器生成，请勿修改-->
                  <el-form-item label="所属用户" prop="create_user" min-width="100">
          <el-input-number v-model="formData.create_user" autocomplete="off" />
        </el-form-item>
        <el-form-item label="应用id" prop="application_id" min-width="100">
          <el-input-number v-model="formData.application_id" autocomplete="off" />
        </el-form-item>
        <el-form-item label="版本号" prop="version" :show-overflow-tooltip="true">
          <el-input v-model="formData.version" autocomplete="off" />
        </el-form-item>
        <el-form-item label="配置ids" prop="config_ids" :show-overflow-tooltip="true">
          <el-input v-model="formData.config_ids" autocomplete="off" />
        </el-form-item>
        <el-form-item label="状态" prop="state">
          <el-switch
              v-model="formData.state"
              inline-prompt
              :active-value="1"
              :inactive-value="-1"
          />
        </el-form-item>
        <el-form-item label="创建时间" prop="create_time" v-if="dialogType === 'edit'" >
          <el-input v-model="formData.create_time" autocomplete="off"/>
        </el-form-item>
        <el-form-item label="更新时间" prop="update_time" v-if="dialogType === 'edit'" >
          <el-input v-model="formData.update_time" autocomplete="off"/>
        </el-form-item>
          <!--机器生成，请勿修改-->
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
  name: "ApplicationInfoList"
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
const service = axios.create({
  baseURL: baseUrl,
  timeout: 99999
})
// http response 拦截器
service.interceptors.response.use(
    response => {
      if (response.data.code === 200 || response.headers.success === 'true') {
        if (response.headers.msg) {
          response.data.msg = decodeURI(response.headers.msg)
        }
        return response.data
      } else {
        ElMessage({
          showClose: true,
          message: response.data.msg || decodeURI(response.headers.msg),
          type: 'error'
        })
        return response.data.msg ? response.data : response
      }
    },
    error => {
      if (!error.response) {
        ElMessageBox.confirm(`
        <p>检测到请求错误</p>
        <p>${error}</p>
        `, '请求报错', {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: '稍后重试',
          cancelButtonText: '取消'
        })
        return
      }
      switch (error.response.status) {
        case 500:
          ElMessageBox.confirm(`
        <p>检测到接口错误${error}</p>
        <p>错误码<span style="color:red"> 500 </span>：此类错误内容常见于后台panic，请先查看后台日志，如果影响您正常使用可强制登出清理缓存</p>
        `, '接口报错', {
            dangerouslyUseHTMLString: true,
            distinguishCancelAndClose: true,
            confirmButtonText: '清理缓存',
            cancelButtonText: '取消'
          })
              .then(() => {
                const userStore = useUserStore()
                userStore.token = ''
                localStorage.clear()
                router.push({ name: 'Login', replace: true })
              })
          break
        case 404:
          ElMessageBox.confirm(`
          <p>检测到接口错误${error}</p>
          <p>错误码<span style="color:red"> 404 </span>：此类错误多为接口未注册（或未重启）或者请求路径（方法）与api路径（方法）不符--如果为自动化代码请检查是否存在空格</p>
          `, '接口报错', {
            dangerouslyUseHTMLString: true,
            distinguishCancelAndClose: true,
            confirmButtonText: '我知道了',
            cancelButtonText: '取消'
          })
          break
      }
      return error
    }
)

// 获取列表信息
const getList = async (val) => {
  console.log("getList",val)

  let res = await service({
    url: '/admin/ApplicationInfo/v1/gets',
    method: 'post',
    data:val
  })
  if (res.code === 200) {
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
      item.update_time = formatDate(item.update_time)
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
  }
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
  return service({
    url: '/admin/ApplicationInfo/v1/',
    method: 'post',
    data:val
  })
}
//del
const del = (val) =>{
  return service({
    url: '/admin/ApplicationInfo/v1/',
    method: 'delete',
    data:val
  })
}
//update
const update = (val) =>{
  if (val.hasOwnProperty("create_user")){
    val.create_user = parseInt(val.create_user)
  }
  return service({
    url: '/admin/ApplicationInfo/v1/',
    method: 'put',
    data:val
  })
}

//状态更改
const handleStateChange = (row) => {
  console.log("handleStateChange", row)
  update({
    "id": row.id,
    "state": row.state,
  })
  ElMessage({
    showClose: true,
    message: row.id+"操作成功",
    type: 'success'
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
      ElMessage({
        showClose: true,
        message: item.id+"操作成功",
        type: 'success'
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
  console.log("handleDelete", index, row)
  let res = await del({
        "id": row.id
  })
  console.log("handleDelete-res", res)
  if (res.code === 200) {
    ElMessage({
      showClose: true,
      message: row.id+"操作成功",
      type: 'success'
    })
    getList({
      "current": 1,
      "orderBy": "create_time",
      "pageSize": pageSize.value,
      "query": [],
      "sort": "desc"
    })
  }
  row.visible = false
}
//批量删除
const handleBatchDelete = async() => {
  console.log("handleBatchDelete", multipleSelection)
  for (let item of multipleSelection.value) {
    let res = await del({
      "id": item.id
    })
    if (res.code === 200) {
      ElMessage({
        showClose: true,
        message: item.id+"操作成功",
        type: 'success'
      })
      console.log("handleBatchDelete-res", res)
    }
  }

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
  let query = []
  if(searchKW.value != ""){
    query = [
      {
        "handle": "like",
        "key": "zh_name",
        "nextHandle": "and",
        "val": searchKW.value
      }
    ]
  }
  getList({
    "current": currentPage.value,
    "orderBy": "create_time",
    "pageSize": pageSize.value,
    "query": query,
    "sort": "desc"
  })
}
const handleCurrentChange = (val) => {
  currentPage.value = val
  let query = []
  if(searchKW.value != ""){
    query = [
      {
        "handle": "like",
        "key": "zh_name",
        "nextHandle": "and",
        "val": searchKW.value
      }
    ]
  }
  getList({
    "current": currentPage.value,
    "orderBy": "create_time",
    "pageSize": pageSize.value,
    "query": query,
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
  create_user: [
    { required: true, message: '请输入创建人id', trigger: 'blur' },
    { validator: mustUint, trigger: 'blur', message: '必须为正整数' }
  ]
})
const enterDialog = async() => {
  console.log("enterDialog", formData.value)
  if (dialogType.value == "add"){
    let time_ = new Date().Format("yyyy-MM-dd hh:mm:ss")
    let res = await add(formData.value)
    console.log("enterDialog-add-res", res)
    if (res.code === 200) {
      ElMessage({
        showClose: true,
        message: "新增操作成功",
        type: 'success'
      })
      formData.value["id"] = res.data["id"]
      formData.value["create_time"] = time_
      tableData.value.push(JSON.parse(JSON.stringify(formData.value)) )
    }
  } else if (dialogType.value == "edit"){
    let data = {"id":formData.value["id"] }
    for (let key in formData.value ){
      if ( formData.value[key] != rowHistory[key] ){
        data[key] = formData.value[key]
      }
    }
    console.log("data", data)
    let time_ = new Date().Format("yyyy-MM-dd hh:mm:ss")
    let res = await update(data)
    console.log("enterDialog-update-res", res)
    if (res.code === 200) {
      ElMessage({
        showClose: true,
        message: "编辑操作成功",
        type: 'success'
      })
      formData.value["id"] = res.data["id"]
      formData.value["create_time"] = time_
      tableData.value[editIndex] = JSON.parse(JSON.stringify(formData.value))
    }
  }
  // 历史记录
  rowHistory = JSON.parse(JSON.stringify(formData.value))
  dialogFormVisible.value = false
}
const closeDialog = () => {
  dialogFormVisible.value = false
}
const beforeReferenc = () => {
  formData.value = rowHistory
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