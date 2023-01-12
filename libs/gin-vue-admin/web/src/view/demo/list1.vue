<template>
  <div class="gva-table-box">
    <div class="gva-btn-list">
      <el-button size="small" type="primary" icon="plus" @click="handleAdd()">新增</el-button>
      <el-button size="small" type="danger" icon="delete" @click="handleDelete()">批量删除</el-button>
    </div>
    <el-table
        row-key="data"
        :data="tableData"
        :default-sort="{ prop: 'date', order: 'descending' }"
        @selection-change="handleSelectionChange"
        border style="width: 100%" height="600">
      <el-table-column type="selection" width="40" />
      <el-table-column fixed prop="date" label="Date" sortable min-width="120" />
      <el-table-column
          column-key="name"
          :filters="nameList"
          :filter-method="filterHandler"
          prop="name" label="姓名" min-width="100" />
      <el-table-column prop="state" label="状态" min-width="100" />
      <el-table-column prop="city" label="城市" :show-overflow-tooltip="true" min-width="100" />
      <el-table-column prop="address" label="地址" :show-overflow-tooltip="true" min-width="120" />
      <el-table-column prop="zip" label="压缩" min-width="100" />
      <el-table-column label="操作" min-width="140" fixed="right">
        <template #header>
          <el-input v-model="search" size="small" placeholder="模糊搜索" />
        </template>
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
          :page-sizes="pageSizes"
          :total="total"
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
      />
    </div>
    <!-- 新增弹窗 -->
    <el-dialog v-model="dialogFormVisible" :title="dialogTitle">
      <el-form ref="authorityForm" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="父级角色" prop="parentId">
          <el-cascader
              v-model="form.parentId"
              style="width:100%"
              :disabled="dialogType==='add'"
              :options="AuthorityOption"
              :props="{ checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
              :show-all-levels="false"
              filterable
          />
        </el-form-item>
        <el-form-item label="角色ID" prop="authorityId">
          <el-input v-model="form.authorityId" :disabled="dialogType=='edit'" autocomplete="off" />
        </el-form-item>
        <el-form-item label="角色姓名" prop="authorityName">
          <el-input v-model="form.authorityName" autocomplete="off" />
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

const tableData = ref([])
const search = ref('')
const selectd = ref([])
const pageSizes = ref([10, 20, 30, 50])
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(25)

const nameList = [
  { text: '小明', value: '小明' },
  { text: '小红', value: '小红' },
  { text: '小白', value: '小白' },
  { text: '小蓝', value: '小蓝' },
]
for (var i=0;i<total.value;i++){
  tableData.value.push(
      {
        date: '2016-05-0'+i,
        name: nameList[i % 4]["text"],
        state: 'California',
        city: 'Los Angeles',
        address: 'No. 189, Grove St, Los Angeles',
        zip: 'CA 90036',
        tag: 'Home',
      }
  )
}

// 弹窗
const mustUint = (rule, value, callback) => {
  if (!/^[0-9]*[1-9][0-9]*$/.test(value)) {
    return callback(new Error('请输入正整数'))
  }
  return callback()
}
const dialogType = ref('add')
const dialogTitle = ref('新增角色')
const dialogFormVisible = ref(false)
const form = ref({
  authorityId: 0,
  authorityName: '',
  parentId: 0
})
const AuthorityOption = ref([
  {
    authorityId: 0,
    authorityName: '根角色'
  }
])
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

const multipleSelection = ref([])
const handleSelectionChange = (val) => {
  multipleSelection.value = val
  console.log("multipleSelection",multipleSelection.value, val)
}
const handleSizeChange = () => {

}
const handleCurrentChange = () => {

}

const filterHandler = (value, row, column) => {
  const property = column['property']
  return row[property] === value
}
const handleAdd = () => {
  dialogType.value = "add"
  dialogFormVisible.value = true
}
const enterDialog = () => {
  
}
const closeDialog = () => {
  
}
const handleEdit = (index, row) => {
  console.log(index, row)
  form.value = {
    authorityId: 12,
    authorityName: '晓红',
    parentId: 0
  }
  dialogType.value = "edit"
  dialogFormVisible.value = true
}
const handleDelete = (index, row) => {
  console.log(index, row)
  console.log("selectd", selectd)
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