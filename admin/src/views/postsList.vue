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
                    <span style="color: #01AAED;cursor:pointer;" @click="postsDelete(row)">删除</span>
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
            label-width="100px"
            class="demo-ruleForm"
        >
            <el-form-item label="标题:" prop="title">
                <el-input v-model="centerAddForm.title" placeholder="请输入标题"/>
            </el-form-item>
            <el-form-item label="描述:" prop="desc">
                <el-input :rows="5" type="textarea" v-model="centerAddForm.desc" placeholder="请输入描述"/>
            </el-form-item>
            <el-form-item label="文章内容:" prop="content">
                <el-input :rows="5" type="textarea" v-model="centerAddForm.content" placeholder="请输入文章内容"/>
            </el-form-item>
            <el-form-item label="作者:" prop="author">
                <el-input v-model="centerAddForm.author" placeholder="请输入作者"/>
            </el-form-item>
            <el-form-item label="标签:" prop="tags">
                <el-select v-model="centerAddForm.tags" class="m-2" placeholder="请选择标签" size="large" clearable>
                    <el-option
                        v-for="(item,index) in tagsList"
                        :key="index"
                        :label="item.Name"
                        :value="item.ID"
                    />
                </el-select>
            </el-form-item>
            <el-form-item label="分类:" prop="category">
                <el-select v-model="centerAddForm.category" class="m-2" placeholder="请选择分类" size="large" clearable>
                    <el-option
                        v-for="(item,index) in categoryList"
                        :key="index"
                        :label="item.Name"
                        :value="item.ID"
                    />
                </el-select>
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
import { getPostsList,getCateList,getTagsList, createPost, updatePostById, delPostById} from '@/api/api'
    const loading = ref(false)

    const height = ref(window.innerHeight - 170)//表格的高度
    const tableData = ref([])
    const optiobsC = ref([
        {label: '标题', prop: 'Title'},
        {label: '作者', prop: 'Author'},
        {label: '描述', prop: 'Desc'},
        {label: '标签', prop: 'Tags'},
        {label: '分类', prop: 'Category'},
        {label: '点赞人数', prop: 'Love'},
        {label: '查看人数', prop: 'Watch'},
        {label: '操作', prop: 'action'}
    ])
    const quotientList = async () => {
        loading.value = false
        let res = await getPostsList()
        if(res){
            tableData.value = res.data
            loading.value = false
        }
    }
    quotientList()

    // 获取分类和标签
    const tagsList = ref([])
    const categoryList = ref([])
    const getCateLists = async () => {
        loading.value = false
        let res = await getCateList()
        if(res){
            tagsList.value = res.data
            loading.value = false
        }
    }
    const getTagsLists = async () => {
        loading.value = false
        let res = await getTagsList()
        if(res){
            categoryList.value = res.data
            loading.value = false
        }
    }
    getCateLists()
    getTagsLists()

    //新增
    const centerDialogVisible = ref(false)
    const centerAddForm = reactive({
        id:'',
        title: '',
        desc: '',
        content: '',
        author: '',
        tags: '',
        category: '',
    })
    const centerRules = reactive({
        title: { required: true, message: '请输入', trigger: 'blur' },
        desc: { required: true, message: '请输入', trigger: 'blur' },
        content: { required: true, message: '请输入', trigger: 'blur' },
        author: { required: true, message: '请输入', trigger: 'blur' },
        tags: { required: true, message: '请输入', trigger: 'blur' },
        category: { required: true, message: '请输入', trigger: 'blur' }
    })
    const centerVisible = (row,type) => {
        centerDialogVisible.value = true
        if(type == 1){
            centerAddForm.id = ""
            centerAddForm.title = ""
            centerAddForm.desc = ""
            centerAddForm.content = ""
            centerAddForm.author = ""
            centerAddForm.tags = ""
            centerAddForm.category = ""
        }else{
            centerAddForm.id = row.ID
            centerAddForm.title = row.Title
            centerAddForm.desc = row.Desc
            centerAddForm.content = row.Content
            centerAddForm.author = row.Author
            centerAddForm.tags = Number(row.Tags)
            centerAddForm.category = row.Category
        }
    }
    const centerAddRef = ref(null)
    const onAdd = () => {
        centerAddRef.value.validate(async valid => {
            if(valid){
                if(centerAddForm.id != ""){
                    const res = await updatePostById(centerAddForm)
                    if(res){
                        ElMessage.success("修改成功")
                        centerDialogVisible.value = false
                        quotientList()
                    }
                }else{
                    const res = await createPost(centerAddForm)
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
    const postsDelete = (row) => {
        ElMessageBox.confirm(
            '确实删除文章：'+ row.Title +'?',
            '删除',
            {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning',
            }
        ).then(async () => {
            const res = await delPostById({id: row.ID})
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
