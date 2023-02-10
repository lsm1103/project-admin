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
        row-key="{{ .PrimaryKey }}"
        :data="tableData"
        :default-sort="{ prop: 'date', order: 'descending' }"
        @selection-change="handleSelectionChange"
        border style="width: 100%" height="600" >
      <el-table-column type="selection" width="40" />
      <el-table-column fixed sortable prop="create_time" label="创建时间" min-width="80" :show-overflow-tooltip=" true" align="center" />
      <el-table-column prop="id" label="主键" min-width="60" :show-overflow-tooltip=" true" align="center" />

      <el-table-column prop="en_name" label="英文名称" min-width="80" align="center" />

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
        <el-form-item label="中文名称" prop="zh_name">
          <el-input v-model="formData.zh_name" autocomplete="off" />
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