<template>
   <div class="quotient-order-info">
       <div style="background-color: #FFFFFF;">
           <el-form :inline="true" class="demo-form-inline">
                <el-form-item>
                    <el-button type="primary" @click="centerVisible">新增</el-button>
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
                <template v-slot ="{ row }" v-if="item.prop === 'State'">
                    <el-switch v-model="row.State" width="60" inline-prompt active-text="开启" inactive-text="关闭" @change="changePower(row)"/>
                </template>
                <template v-slot ="{ row }" v-else-if="item.prop === 'action'">
                    <span style="color: #01AAED;cursor:pointer;" @click="accountView(row)">查看&nbsp;&nbsp;|&nbsp;&nbsp;</span>
                    <span style="color: #01AAED;cursor:pointer;" @click="accountUpdate(row)">修改&nbsp;&nbsp;|&nbsp;&nbsp;</span>
                    <span style="color: #01AAED;cursor:pointer;" @click="accountDelete(row)">删除</span>
                </template>
            </el-table-column>
        </el-table>
   </div>
    <!-- 新增 -->
    <el-dialog v-model="centerDialogVisible" title="新增" width="30%" center>
        <el-form
            ref="centerAddRef"
            :model="centerAddForm"
            status-icon
            :rules="centerRules"
            label-width="120px"
            class="demo-ruleForm"
        >
            <el-form-item label="商户名称:" prop="username">
                <el-input v-model="centerAddForm.username" placeholder="请输入商户名称"/>
            </el-form-item>
            <el-form-item label="密码:" prop="password">
                <el-input v-model="centerAddForm.password" placeholder="请输入密码"/>
            </el-form-item>
            <el-form-item label="手机号码:" prop="phone">
                <el-input v-model="centerAddForm.phone" placeholder="请输入手机号码"/>
            </el-form-item>
            <el-form-item label="邮箱:" prop="email">
                <el-input v-model="centerAddForm.email" placeholder="请输入邮箱"/>
            </el-form-item>
            <el-form-item label="状态:" prop="state">
                <el-select v-model="centerAddForm.state" class="m-2" placeholder="请选择状态" size="large" clearable>
                    <el-option label="开启" value="1" />
                    <el-option label="关闭" value="2" />
                </el-select>
            </el-form-item>
        </el-form>
        <template #footer>
        <span class="dialog-footer">
            <el-button type="primary" @click="onAdd">确定</el-button>
        </span>
        </template>
    </el-dialog>
    <!-- 查看 -->
    <el-dialog v-model="viewDialogVisible" title="查看" width="30%" center>
        <el-form
            ref="viewRef"
            :model="viewForm"
            status-icon
            label-width="120px"
            class="demo-ruleForm"
        >
            <el-form-item label="商户名称:">
                <el-input v-model="viewForm.Username" disabled/>
            </el-form-item>
            <el-form-item label="密码:">
                <el-input v-model="viewForm.Password" disabled/>
            </el-form-item>
            <el-form-item label="手机号码:">
                <el-input v-model="viewForm.Phone" disabled/>
            </el-form-item>
            <el-form-item label="邮箱:">
                <el-input v-model="viewForm.Email" disabled/>
            </el-form-item>
            <el-form-item label="状态:">
                <el-select v-model="viewForm.State" class="m-2" disabled size="large" clearable>
                    <el-option label="开启" value="1" />
                    <el-option label="关闭" value="2" />
                </el-select>
            </el-form-item>
        </el-form>
        <template #footer>
        <span class="dialog-footer">
            
        </span>
        </template>
    </el-dialog>
    <!-- 编辑 -->
    <el-dialog v-model="updateDialogVisible" title="修改" width="30%" center>
        <el-form
            ref="updateRef"
            :model="updateForm"
            status-icon
            :rules="updateRules"
            label-width="120px"
            class="demo-ruleForm"
        >
            <el-form-item label="商户名称:" prop="username">
                <el-input v-model="updateForm.username" placeholder="请输入商户名称"/>
            </el-form-item>
            <el-form-item label="密码:">
                <el-input v-model="updateForm.password" placeholder="请输入密码"/>
            </el-form-item>
            <el-form-item label="手机号码:" prop="phone">
                <el-input v-model="updateForm.phone" placeholder="请输入手机号码"/>
            </el-form-item>
            <el-form-item label="邮箱:" prop="email">
                <el-input v-model="updateForm.email" placeholder="请输入邮箱"/>
            </el-form-item>
            <el-form-item label="状态:" prop="state">
                <el-select v-model="updateForm.state" class="m-2" placeholder="请选择状态" size="large" clearable>
                    <el-option label="开启" value="1" />
                    <el-option label="关闭" value="2" />
                </el-select>
            </el-form-item>
        </el-form>
        <template #footer>
        <span class="dialog-footer">
            <el-button type="primary" @click="onUpdate">确定</el-button>
        </span>
        </template>
    </el-dialog>
</template>
<script setup>
import { ref,reactive } from 'vue'
import { getUsersList, createUsers, getUserById, disableUserById, enableUserById, updateUserById, delUserById } from '@/api/api'
import { ElMessage, ElMessageBox } from 'element-plus'
    const loading = ref(false)

    const height = ref(window.innerHeight - 160)//表格的高度
    const tableData = ref([])
    const optiobsC = ref([
        {label: '编号', prop: 'ID'},
        {label: '用户名', prop: 'Username'},
        {label: '手机号码', prop: 'Phone'},
        {label: '邮箱', prop: 'Email'},
        {label: '创建时间', prop: 'CreatedAt'},
        {label: '上一次修改时间', prop: 'UpdatedAt'},
        {label: '状态', prop: 'State'},
        {label: '操作', prop: 'action'}
    ])
    const quotientList = async () => {
        loading.value = false
        let res = await getUsersList()
        if(res){
            tableData.value = res.data
            for(let i=0; i<tableData.value.length ;i++){
                tableData.value[i].State = tableData.value[i].State == 1 ? true : false
            }
            loading.value = false
        }
    }
    quotientList()

    //新增
    const centerDialogVisible = ref(false)
    const centerAddForm = reactive({
        id: 0,
        username: '',
        password: '',
        phone: "",
        email: "",
        state: "",
    })
    const centerRules = reactive({
        username: { required: true, message: '请输入', trigger: 'blur' },
        password: { required: true, message: '请输入', trigger: 'blur' },
        phone: { required: true, message: '请输入', trigger: 'blur' },
        email: { required: true, message: '请输入', trigger: 'blur' },
        state: { required: true, message: '请选择', trigger: 'blur' }
    })
    const centerVisible = () => {
        centerDialogVisible.value = true
        centerAddForm.id = 0
        centerAddForm.username = ""
        centerAddForm.password = ""
        centerAddForm.phone = ""
        centerAddForm.email = ""
        centerAddForm.state = ""
    }
    const centerAddRef = ref(null)
    const onAdd = () => {
        centerAddRef.value.validate(async valid => {
            if(valid){
                const res = await createUsers(centerAddForm)
                if(res){
                    ElMessage.success("新增成功")
                    centerDialogVisible.value = false
                    quotientList()
                }
            }else{
                ElMessage.error("请填写完整")
            }
        })
    }

    // 查看
    const viewDialogVisible = ref(false)
    const viewForm = ref([])
    const accountView = async (row) => {
        viewDialogVisible.value = true
        const res = await getUserById({id: row.ID})
        if(res){
            viewForm.value = res.data[0]
            viewForm.value.State = viewForm.value.State+""
        }
    }

    // 点击开启关闭
    const changePower = async (row) => {
        if(!row.State){
            const res = await disableUserById({id: row.ID})
            if(res){
                ElMessage.success("修改成功")
                quotientList()
            }
        }else{
            const res = await enableUserById({id: row.ID})
            if(res){
                ElMessage.success("修改成功")
                quotientList()
            } 
        }
    }

    // 修改
    const updateDialogVisible = ref(false)
    const updateForm = reactive({
        id: 0,
        username: '',
        password: '',
        phone: "",
        email: "",
        state: "",
    })
    const updateRules = reactive({
        username: { required: true, message: '请输入', trigger: 'blur' },
        phone: { required: true, message: '请输入', trigger: 'blur' },
        email: { required: true, message: '请输入', trigger: 'blur' },
        state: { required: true, message: '请选择', trigger: 'blur' }
    })
    const accountUpdate = (row) => {
        updateDialogVisible.value = true
        updateForm.id = row.ID
        updateForm.username = row.Username
        updateForm.password = ""
        updateForm.phone = row.Phone
        updateForm.email = row.Email
        updateForm.state = row.State ? "1" : "2"
    }
    const updateRef = ref(null)
    const onUpdate = () => {
        updateRef.value.validate(async valid => {
            if(valid){
                const res = await updateUserById(updateForm)
                if(res){
                    ElMessage.success("新增成功")
                    updateDialogVisible.value = false
                    quotientList()
                }
            }else{
                ElMessage.error("请填写完整")
            }
        })
    }
    
    // 删除
    const accountDelete = (row) => {
        ElMessageBox.confirm(
            '确实删除用户：'+ row.Username +'?',
            '删除',
            {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning',
            }
        ).then(async () => {
            const res = await delUserById({id: row.ID})
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
