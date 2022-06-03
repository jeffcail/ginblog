<template>
   <div class="quotient-order-info">
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
            </el-table-column>
        </el-table>
    </div>
</template>
<script setup>
import { ref } from 'vue'
import { getCommentList } from '@/api/api'
    const loading = ref(false)

    const height = ref(window.innerHeight - 110)//表格的高度
    const tableData = ref([])
    const optiobsC = ref([
        {label: '编号', prop: 'ID'},
        {label: '文章id', prop: 'PostID'},
        {label: '评论内容', prop: 'CommentContent'},
        {label: '评论人邮箱', prop: 'Email'}
    ])
    const quotientList = async () => {
        loading.value = false
        let res = await getCommentList()
        if(res){
            tableData.value = res.data
            loading.value = false
        }
    }
    quotientList()


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
