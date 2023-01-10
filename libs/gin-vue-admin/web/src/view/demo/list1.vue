<template>
  <div class="gva-table-box">
    <el-table
        :data="tableData" border
        :default-sort="{ prop: 'date', order: 'descending' }"
        style="width: 100%" height="600">
      <el-table-column fixed prop="date" label="Date" sortable min-width="150" />
      <el-table-column prop="name" label="姓名" min-width="120" />
      <el-table-column prop="state" label="状态" min-width="120" />
      <el-table-column prop="city" label="城市" min-width="120" />
      <el-table-column prop="address" label="地址" min-width="200" />
      <el-table-column prop="zip" label="压缩" min-width="120" />
      <el-table-column fixed="right" min-width="300"> <!--label="操作"-->
        <template #header>
          <el-input v-model="search" size="small" placeholder="模糊搜索" />
        </template>
        <template #default="scope">
          <el-button link type="primary" size="small" @click="handleClick">详情</el-button>
          <el-button link type="primary" size="small">编辑</el-button>
          <el-button link type="danger" size="small" @click="handleDelete(scope.$index, scope.row)" >删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: "List1"
}
</script>

<script setup>
// https://element-plus.gitee.io/zh-CN/component/table.html#%E8%87%AA%E5%AE%9A%E4%B9%89%E8%A1%A8%E5%A4%B4
const handleEdit = (index, row) => {
  console.log(index, row)
}
const handleDelete = (index, row) => {
  console.log(index, row)
}

const tableData = [
  {
    date: '2016-05-03',
    name: 'Tom',
    state: 'California',
    city: 'Los Angeles',
    address: 'No. 189, Grove St, Los Angeles',
    zip: 'CA 90036',
    tag: 'Home',
  },
  {
    date: '2016-05-02',
    name: 'Tom',
    state: 'California',
    city: 'Los Angeles',
    address: 'No. 189, Grove St, Los Angeles',
    zip: 'CA 90036',
    tag: 'Office',
  },
  {
    date: '2016-05-04',
    name: 'Tom',
    state: 'California',
    city: 'Los Angeles',
    address: 'No. 189, Grove St, Los Angeles',
    zip: 'CA 90036',
    tag: 'Home',
  },
  {
    date: '2016-05-01',
    name: 'Tom',
    state: 'California',
    city: 'Los Angeles',
    address: 'No. 189, Grove St, Los Angeles',
    zip: 'CA 90036',
    tag: 'Office',
  },
]

for (var i=0;i<10;i++){
  tableData.push(
      {
        date: '2016-05-03',
        name: 'Tom',
        state: 'California',
        city: 'Los Angeles',
        address: 'No. 189, Grove St, Los Angeles',
        zip: 'CA 90036',
        tag: 'Home',
      }
  )
}
</script>
