
<template>
    <section>

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


        <el-table :data="results" highlight-current-row stripe v-loading="listLoading"
                  style="width: 100%;" @selection-change="handleSelectionChange">
            <el-table-column type="selection" width="55"></el-table-column>
            <el-table-column prop="id" label="编号"></el-table-column>

                <el-table-column prop="topicId" label="topicId"></el-table-column>

                <el-table-column prop="tagId" label="tagId"></el-table-column>

                <el-table-column prop="createTime" label="createTime"></el-table-column>

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
            <el-form :model="addForm" label-width="80px" :rules="addFormRules" ref="addForm">

                    <el-form-item label="topicId" prop="rule">
                        <el-input v-model="addForm.topicId"></el-input>
                    </el-form-item>

                    <el-form-item label="tagId" prop="rule">
                        <el-input v-model="addForm.tagId"></el-input>
                    </el-form-item>

                    <el-form-item label="createTime" prop="rule">
                        <el-input v-model="addForm.createTime"></el-input>
                    </el-form-item>

            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click.native="addFormVisible = false">取消</el-button>
                <el-button type="primary" @click.native="addSubmit" :loading="addLoading">提交</el-button>
            </div>
        </el-dialog>


        <el-dialog title="编辑" :visible.sync="editFormVisible" :close-on-click-modal="false">
            <el-form :model="editForm" label-width="80px" :rules="editFormRules" ref="editForm">
                <el-input v-model="editForm.id" type="hidden"></el-input>

                    <el-form-item label="topicId" prop="rule">
                        <el-input v-model="editForm.topicId"></el-input>
                    </el-form-item>

                    <el-form-item label="tagId" prop="rule">
                        <el-input v-model="editForm.tagId"></el-input>
                    </el-form-item>

                    <el-form-item label="createTime" prop="rule">
                        <el-input v-model="editForm.createTime"></el-input>
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
import HttpClient from '@/apis/HttpClient';

export default {
  name: 'List',
  data() {
    return {
      results: [],
      listLoading: false,
      page: {},
      filters: {},
      selectedRows: [],

      addForm: {

        topicId: '',

        tagId: '',

        createTime: '',

      },
      addFormVisible: false,
      addFormRules: {},
      addLoading: false,

      editForm: {
        id: '',

        topicId: '',

        tagId: '',

        createTime: '',

      },
      editFormVisible: false,
      editFormRules: {},
      editLoading: false,
    };
  },
  mounted() {
    this.list();
  },
  methods: {
    list() {
      const me = this;
      me.listLoading = true;
      const params = Object.assign(me.filters, {
        page: me.page.page,
        limit: me.page.limit,
      });
      HttpClient.post('/api/admin/topic-tag/list', params)
        .then((data) => {
          me.results = data.results;
          me.page = data.page;
        })
        .finally(() => {
          me.listLoading = false;
        });
    },
    handlePageChange(val) {
      this.page.page = val;
      this.list();
    },
    handleLimitChange(val) {
      this.page.limit = val;
      this.list();
    },
    handleAdd() {
      this.addForm = {
        name: '',
        description: '',
      };
      this.addFormVisible = true;
    },
    addSubmit() {
      const me = this;
      HttpClient.post('/api/admin/topic-tag/create', this.addForm)
        .then((data) => {
          me.$message({ message: '提交成功', type: 'success' });
          me.addFormVisible = false;
          me.list();
        })
        .catch((rsp) => {
          me.$notify.error({ title: '错误', message: rsp.message });
        });
    },
    handleEdit(index, row) {
      const me = this;
      HttpClient.get(`/api/admin/topic-tag/${row.id}`)
        .then((data) => {
          me.editForm = Object.assign({}, data);
          me.editFormVisible = true;
        })
        .catch((rsp) => {
          me.$notify.error({ title: '错误', message: rsp.message });
        });
    },
    editSubmit() {
      const me = this;
      HttpClient.post('/api/admin/topic-tag/update', me.editForm)
        .then((data) => {
          me.list();
          me.editFormVisible = false;
        })
        .catch((rsp) => {
          me.$notify.error({ title: '错误', message: rsp.message });
        });
    },

    handleSelectionChange(val) {
      this.selectedRows = val;
    },
  },
};
</script>

<style scoped>

</style>
