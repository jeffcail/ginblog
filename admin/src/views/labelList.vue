<template>
   <div class="quotient-order-info">
       <div style="background-color: #FFFFFF;">
           <el-form :inline="true" class="demo-form-inline">
                <el-form-item>
                    <el-button type="primary" @click="centerVisible('',1)">新增</el-button>
                </el-form-item>
            </el-form>
       </div>
       <el-table 
            ref="tableC"
            :data="tableData" 
            border 
            style="width: 100%"
            :height="height"
            resizable = true
            :row-class-name="tableRowClassName"
            v-loading="loading"
            element-loading-text="Loading..."
        >
            <el-table-column type="index" show-overflow-tooltip></el-table-column>
            <el-table-column 
                :prop="item.prop" 
                :label="item.label" 
                :width="item.width"
                :fixed="item.fixed"
                v-for="(item,index) in optiobsC"
                :key="index"
                show-overflow-tooltip
            >
                <template v-slot ="{ row }" v-if="item.prop === 'action'">
                    <span style="color: #01AAED;cursor:pointer;" @click="centerVisible(row,2)">修改&nbsp;&nbsp;|&nbsp;&nbsp;</span>
                    <span style="color: #01AAED;cursor:pointer;" @click="labelDelete(row)">删除</span>
                </template>
            </el-table-column>
        </el-table>
    </div>
    <!-- 新增 -->
    <el-dialog v-model="centerDialogVisible" title="标签" width="30%" center>
        <el-form
            ref="centerAddRef"
            :model="centerAddForm"
            status-icon
            :rules="centerRules"
            label-width="120px"
            class="demo-ruleForm"
        >
            <el-form-item label="标签名称:" prop="name">
                <el-input v-model="centerAddForm.name" placeholder="请输入标签名称"/>
            </el-form-item>
        </el-form>
        <template #footer>
        <span class="dialog-footer">
            <el-button type="primary" @click="onAdd">确定</el-button>
        </span>
        </template>
    </el-dialog>
</template>
<script setup>
import { ref,reactive } from 'vue'
import { getTagsList, createTags, updateTagsById, delTagsById} from '@/api/api'
    const loading = ref(false)

    const height = ref(window.innerHeight - 170)//表格的高度
    const tableData = ref([])
    const optiobsC = ref([
        {label: '标签名称', prop: 'Name'},
        {label: '创建时间', prop: 'CreatedAt'},
        {label: '上一次修改时间', prop: 'UpdatedAt'},
        {label: '操作', prop: 'action'}
    ])
    const quotientList = async () => {
        loading.value = false
        let res = await getTagsList()
        if(res){
            tableData.value = res.data
            loading.value = false
        }
    }
    quotientList()

    //新增
    const centerDialogVisible = ref(false)
    const centerAddForm = reactive({
        id:'',
        name: '',
    })
    const centerRules = reactive({
        name: { required: true, message: '请输入', trigger: 'blur' },
    })
    const centerVisible = (row,type) => {
        centerDialogVisible.value = true
        if(type == 1){
            centerAddForm.id = ""
            centerAddForm.name = ""
        }else{
            centerAddForm.id = row.ID
            centerAddForm.name = row.Name
        }
    }
    const centerAddRef = ref(null)
    const onAdd = () => {
        centerAddRef.value.validate(async valid => {
            if(valid){
                if(centerAddForm.id != ""){
                    const res = await updateTagsById(centerAddForm)
                    if(res){
                        ElMessage.success("修改成功")
                        centerDialogVisible.value = false
                        quotientList()
                    }
                }else{
                    const res = await createTags(centerAddForm)
                    if(res){
                        ElMessage.success("新增成功")
                        centerDialogVisible.value = false
                        quotientList()
                    }
                }
            }else{
                ElMessage.error("请填写完整")
            }
        })
    }

    // 删除
    const labelDelete = (row) => {
        ElMessageBox.confirm(
            '确实删除标签：'+ row.Name +'?',
            '删除',
            {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning',
            }
        ).then(async () => {
            const res = await delTagsById({id: row.ID})
            if(res){
                ElMessage.success("删除成功")
                quotientList()
            }
        }).catch(() => {
            
        })
    }

    //换行变颜色
    const tableRowClassName = (row,rowIndex) => {
        if (row.rowIndex%2>0) {
            return 'warning-row'
        } else{
            return ''
        }
    }

</script>
<style lang="scss" scoped>
.quotient-order-info{
    width: 100%;
    border-radius: 2px;
    background-color: #FFFFFF;
    padding: 10px 15px;
}
::v-deep .el-table_1_column_1{
    text-align: center;
}
::v-deep .el-table__expand-icon .el-icon svg{
    display: none;
}
::v-deep .el-select {
    width: 100%;
}

</style>
