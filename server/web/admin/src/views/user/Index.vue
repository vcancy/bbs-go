
<template>
    <section class="page-container">
        
        <el-col :span="24" class="toolbar">
            <el-form :inline="true" :model="filters">
                <el-form-item>
                    <el-input v-model="filters.name" placeholder="名称"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" v-on:click="list">查询</el-button>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="handleAdd">新增</el-button>
                </el-form-item>
            </el-form>
        </el-col>

        
        <el-table :data="results" highlight-current-row border v-loading="listLoading"
                  style="width: 100%;" @selection-change="handleSelectionChange">
            <el-table-column type="selection" width="55"></el-table-column>
            <el-table-column prop="id" label="编号"></el-table-column>
            
			<el-table-column prop="username" label="username"></el-table-column>
            
			<el-table-column prop="email" label="email"></el-table-column>
            
			<el-table-column prop="nickname" label="nickname"></el-table-column>
            
			<el-table-column prop="avatar" label="avatar"></el-table-column>
            
			<el-table-column prop="password" label="password"></el-table-column>
            
			<el-table-column prop="status" label="status"></el-table-column>
            
			<el-table-column prop="roles" label="roles"></el-table-column>
            
			<el-table-column prop="type" label="type"></el-table-column>
            
			<el-table-column prop="description" label="description"></el-table-column>
            
			<el-table-column prop="createTime" label="createTime"></el-table-column>
            
			<el-table-column prop="updateTime" label="updateTime"></el-table-column>
            
            <el-table-column label="操作" width="150">
                <template slot-scope="scope">
                    <el-button size="small" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
                </template>
            </el-table-column>
        </el-table>

        
        <el-col :span="24" class="toolbar">
            <el-pagination layout="total, sizes, prev, pager, next, jumper" :page-sizes="[20, 50, 100, 300]"
                           @current-change="handlePageChange"
                           @size-change="handleLimitChange"
                           :current-page="page.page"
                           :page-size="page.limit"
                           :total="page.total"
                           style="float:right;">
            </el-pagination>
        </el-col>


        
        <el-dialog title="新增" :visible.sync="addFormVisible" :close-on-click-modal="false">
            <el-form :model="addForm" label-width="80px" ref="addForm">
                
				<el-form-item label="username">
					<el-input v-model="addForm.username"></el-input>
				</el-form-item>
                
				<el-form-item label="email">
					<el-input v-model="addForm.email"></el-input>
				</el-form-item>
                
				<el-form-item label="nickname">
					<el-input v-model="addForm.nickname"></el-input>
				</el-form-item>
                
				<el-form-item label="avatar">
					<el-input v-model="addForm.avatar"></el-input>
				</el-form-item>
                
				<el-form-item label="password">
					<el-input v-model="addForm.password"></el-input>
				</el-form-item>
                
				<el-form-item label="status">
					<el-input v-model="addForm.status"></el-input>
				</el-form-item>
                
				<el-form-item label="roles">
					<el-input v-model="addForm.roles"></el-input>
				</el-form-item>
                
				<el-form-item label="type">
					<el-input v-model="addForm.type"></el-input>
				</el-form-item>
                
				<el-form-item label="description">
					<el-input v-model="addForm.description"></el-input>
				</el-form-item>
                
				<el-form-item label="createTime">
					<el-input v-model="addForm.createTime"></el-input>
				</el-form-item>
                
				<el-form-item label="updateTime">
					<el-input v-model="addForm.updateTime"></el-input>
				</el-form-item>
                
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click.native="addFormVisible = false">取消</el-button>
                <el-button type="primary" @click.native="addSubmit" :loading="addLoading">提交</el-button>
            </div>
        </el-dialog>

        
        <el-dialog title="编辑" :visible.sync="editFormVisible" :close-on-click-modal="false">
            <el-form :model="editForm" label-width="80px" ref="editForm">
                <el-input v-model="editForm.id" type="hidden"></el-input>
                
				<el-form-item label="username">
					<el-input v-model="editForm.username"></el-input>
				</el-form-item>
                
				<el-form-item label="email">
					<el-input v-model="editForm.email"></el-input>
				</el-form-item>
                
				<el-form-item label="nickname">
					<el-input v-model="editForm.nickname"></el-input>
				</el-form-item>
                
				<el-form-item label="avatar">
					<el-input v-model="editForm.avatar"></el-input>
				</el-form-item>
                
				<el-form-item label="password">
					<el-input v-model="editForm.password"></el-input>
				</el-form-item>
                
				<el-form-item label="status">
					<el-input v-model="editForm.status"></el-input>
				</el-form-item>
                
				<el-form-item label="roles">
					<el-input v-model="editForm.roles"></el-input>
				</el-form-item>
                
				<el-form-item label="type">
					<el-input v-model="editForm.type"></el-input>
				</el-form-item>
                
				<el-form-item label="description">
					<el-input v-model="editForm.description"></el-input>
				</el-form-item>
                
				<el-form-item label="createTime">
					<el-input v-model="editForm.createTime"></el-input>
				</el-form-item>
                
				<el-form-item label="updateTime">
					<el-input v-model="editForm.updateTime"></el-input>
				</el-form-item>
                
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click.native="editFormVisible = false">取消</el-button>
                <el-button type="primary" @click.native="editSubmit" :loading="editLoading">提交</el-button>
            </div>
        </el-dialog>
    </section>
</template>

<script>
  import HttpClient from '../../apis/HttpClient'

  export default {
    name: "List",
    data() {
      return {
        results: [],
        listLoading: false,
        page: {},
        filters: {},
        selectedRows: [],

        addForm: {
          
          'username': '',
          
          'email': '',
          
          'nickname': '',
          
          'avatar': '',
          
          'password': '',
          
          'status': '',
          
          'roles': '',
          
          'type': '',
          
          'description': '',
          
          'createTime': '',
          
          'updateTime': '',
          
        },
        addFormVisible: false,
        addLoading: false,

        editForm: {
          'id': '',
          
          'username': '',
          
          'email': '',
          
          'nickname': '',
          
          'avatar': '',
          
          'password': '',
          
          'status': '',
          
          'roles': '',
          
          'type': '',
          
          'description': '',
          
          'createTime': '',
          
          'updateTime': '',
          
        },
        editFormVisible: false,
        editLoading: false,
      }
    },
    mounted() {
      this.list();
    },
    methods: {
      list() {
        let me = this
        me.listLoading = true
		let params = Object.assign(me.filters, {
          page: me.page.page,
          limit: me.page.limit
        })
        HttpClient.post('/api/admin/user/list', params)
          .then(data => {
            me.results = data.results
            me.page = data.page
          })
          .finally(() => {
            me.listLoading = false
          })
      },
      handlePageChange (val) {
        this.page.page = val
        this.list()
      },
      handleLimitChange (val) {
        this.page.limit = val
        this.list()
      },
      handleAdd() {
        this.addForm = {
          name: '',
          description: '',
        }
        this.addFormVisible = true
      },
      addSubmit() {
        let me = this
        HttpClient.post('/api/admin/user/create', this.addForm)
          .then(data => {
            me.$message({message: '提交成功', type: 'success'});
            me.addFormVisible = false
            me.list()
          })
          .catch(rsp => {
            me.$notify.error({title: '错误', message: rsp.message})
          })
      },
      handleEdit(index, row) {
        let me = this
        HttpClient.get('/api/admin/user/' + row.id)
          .then(data => {
            me.editForm = Object.assign({}, data);
            me.editFormVisible = true
          })
          .catch(rsp => {
            me.$notify.error({title: '错误', message: rsp.message})
          })
      },
      editSubmit() {
        let me = this
        HttpClient.post('/api/admin/user/update', me.editForm)
          .then(data => {
            me.list()
            me.editFormVisible = false
          })
          .catch(rsp => {
            me.$notify.error({title: '错误', message: rsp.message})
          })
      },

      handleSelectionChange(val) {
        this.selectedRows = val
      },
    }
  }
</script>

<style lang="scss" scoped>

</style>

