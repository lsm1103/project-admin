<template>
  <div class="gva-table-box">
    <div class="gva-btn-list" style="text-align: justify;">
      <div style="width: 80%; height: 100%">
        <el-button size="small" type="primary" icon="plus" @click="handleAdd()">新增</el-button>
        <el-button size="small" type="danger" icon="delete" @click="handleDelete()">批量删除</el-button>
      </div>
      <div style="width: 20%; height: 100%">
        <el-input clearable :prefix-icon="Search" size="large" v-model="input" min-width="140" placeholder="模糊搜索"/>
      </div>
    </div>
    <el-table
        row-key="remark"
        :data="tableData"
        :default-sort="{ prop: 'date', order: 'descending' }"
        @selection-change="handleSelectionChange"
        border style="width: 100%" height="600">
      <el-table-column type="selection" width="40" />
      <el-table-column fixed sortable prop="create_time" label="日期" width="180"/>
      <el-table-column fixed prop="id" label="id" :show-overflow-tooltip=" true" min-width="50" />
      <el-table-column
          column-key="create_user"
          :filters="nameList"
          :filter-method="filterHandler"
          prop="create_user" label="创建人" min-width="100" />
      <el-table-column prop="en_name" label="en应用名" min-width="80" />
      <el-table-column prop="zh_name" label="zh应用名" min-width="80" />
      <el-table-column prop="info" label="介绍" :show-overflow-tooltip=" true" min-width="120" />
      <el-table-column prop="project_id" label="项目id" min-width="100" />
      <el-table-column label="状态" min-width="40">
        <template #default="scope">
          <el-switch
              style="--el-switch-on-color: #13ce66; --el-switch-off-color: #797777"
              v-model="scope.row.state"
              inline-prompt
              :active-value="1"
              :inactive-value="-1"
              @change="()=>{handleStateChange(scope.row, v)}"
          />
        </template>
      </el-table-column>
      <el-table-column label="操作" min-width="140" fixed="right">
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
    <!-- 新增弹窗 -->
    <el-dialog v-model="dialogFormVisible" :title="dialogTitle">
      <el-form :model="form" :rules="rules" label-width="80px">
        <el-form-item label="中文名称" prop="zh_name">
          <el-input v-model="form.zh_name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="英文名称" prop="en_name">
          <el-input v-model="form.en_name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="图标" prop="ico" min-width="120">
          <el-input v-model="form.ico" autocomplete="off" />
        </el-form-item>
        <el-form-item label="简介" prop="info" :show-overflow-tooltip="true" min-width="50">
          <el-input v-model="form.info" autocomplete="off" />
        </el-form-item>
        <el-form-item label="创建者id" prop="create_user" min-width="100">
          <el-input v-model="form.create_user" autocomplete="off" />
        </el-form-item>
        <el-form-item label="需求组ids" prop="demand_ids" min-width="80">
          <el-input v-model="form.demand_ids" autocomplete="off" />
        </el-form-item>
        <el-form-item label="文档组ids" prop="doc_ids" min-width="80">
          <el-input v-model="form.doc_ids" autocomplete="off" />
        </el-form-item>
        <el-form-item label="参与者ids" prop="join_users" :show-overflow-tooltip="true" min-width="120">
          <el-input v-model="form.join_users" autocomplete="off" />
        </el-form-item>
        <el-form-item label="参与组ids" prop="join_groups" min-width="100">
          <el-input v-model="form.join_groups" autocomplete="off" />
        </el-form-item>
        <el-form-item label="所属项目id" prop="project_id" min-width="100">
          <el-input v-model="form.project_id" autocomplete="off" />
        </el-form-item>
        <el-form-item label="备注" prop="remark" min-width="100">
          <el-input v-model="form.remark" autocomplete="off" />
        </el-form-item>
        <el-form-item label="排序" prop="rank" min-width="100">
          <el-input v-model="form.rank" autocomplete="off" />
        </el-form-item>
        <el-form-item label="状态" prop="state" min-width="40">
          <el-input v-model="form.state" autocomplete="off" />
        </el-form-item>
        <el-form-item label="创建时间" prop="create_time" min-width="40">
          <el-input v-model="form.create_time" autocomplete="off"/>
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
// https://element-plus.gitee.io/zh-CN/component/table.html#%E8%87%AA%E5%AE%9A%E4%B9%89%E8%A1%A8%E5%A4%B4
import { ref } from 'vue'
import axios from "axios";
import {ElMessage} from "element-plus";
import { formatDate } from '@/utils/format'
import { Search } from '@element-plus/icons-vue'

const input = ref([])
const tableData = ref([])
const search = ref('')
const selectd = ref([])
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(0)
const nameList = ref([])

const baseUrl = "http://127.0.0.1:810"
// 获取列表信息
const getList = (val) => {
  console.log("getList",val)
  axios.post(baseUrl+"/admin/Application/v1/gets", {
    "current": currentPage.value,
    "orderBy": "id",
    "pageSize": pageSize.value,
    "query": [],
    "sort": "desc"
  }, {
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
getList("1")

//列表多选事件触发处理
const multipleSelection = ref([])
const handleSelectionChange = (val) => {
  multipleSelection.value = val
  console.log("multipleSelection",multipleSelection.value, val)
}
//人名筛选功能
const filterHandler = (value, row, column) => {
  const property = column['property']
  return row[property] === value
}

//状态更改
const handleStateChange = (row) => {
  console.log("row", row)
}
//新增
const handleAdd = () => {
  dialogType.value = "add"
  dialogFormVisible.value = true
}
//修改
const handleEdit = (index, row) => {
  console.log(index, row, row.create_time, formatDate(row.create_time))
  form.value = row
  // form.value.create_time = formatDate(row.create_time)
  dialogType.value = "edit"
  dialogFormVisible.value = true
}
//删除
const handleDelete = (index, row) => {
  console.log(index, row)
  console.log("selectd", selectd)
}

//分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getList("handleSizeChange")
}
const handleCurrentChange = (val) => {
  currentPage.value = val
  getList("handleCurrentChange")
}

// 弹窗
const dialogType = ref('add')
const dialogTitle = ref('新增角色')
const dialogFormVisible = ref(false)
const form = ref({
  create_time: '',
  create_user: 0,
  demand_ids: '',
  doc_ids: '',
  en_name: '',
  ico: '',
  id: '',
  info: '',
  join_groups: '',
  join_users: '',
  project_id: '',
  rank: '',
  remark: '',
  state: '',
  update_time: '',
  zn_name: '',
})
const mustUint = (rule, value, callback) => {
  if (!/^[0-9]*[1-9][0-9]*$/.test(value)) {
    return callback(new Error('请输入正整数'))
  }
  return callback()
}
const rules = ref({
  authorityId: [
    { required: true, message: '请输入角色ID', trigger: 'blur' },
    { validator: mustUint, trigger: 'blur', message: '必须为正整数' }
  ],
  authorityName: [
    { required: true, message: '请输入角色名', trigger: 'blur' }
  ],
  parentId: [
    { required: true, message: '请选择父角色', trigger: 'blur' },
  ]
})
const enterDialog = () => {
  console.log("enterDialog", form)
  dialogFormVisible.value = false
}
const closeDialog = () => {
  dialogFormVisible.value = false
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