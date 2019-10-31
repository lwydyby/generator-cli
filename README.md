# LCli
[![Go Report Card](https://goreportcard.com/badge/github.com/lwydyby/generator-cli)](https://goreportcard.com/report/github.com/lwydyby/generator-cli)

spring boot + vue项目代码生成命令行工具

## 简介


``` bash
$ genratorv1 --help
A code generator for spring boot program.


For example:
$ genratorv1 -f ./data.json
          -t ./templates \
          -o ./output


Usage:
  genratorv1 [flags]

Flags:
  -f, --file string       setting文件及data文件所在位置
   
  -o, --output string     输出到的文件夹.
  -t, --template string   模板文件所在文件夹
  -v, --version           版本信息
```

## 安装及使用

项目根目录存在mac上编译好的二进制文件（genratorv1） 其他系统请下载源码自行安装

### 源码安装步骤

``` bash
$ git clone git@github.com:lwydyby/generator-cli.git
$ glide install
$ make install
```


### 说明

1. setting.yaml 

配置了模板名称和相应的输出文件名

```yaml
template:
  format: Go
template_files:
  web:
    file_path: web.tmpl
    output_file_naming: index.vue
  dao:
    file_path: dao.tmpl
    output_file_naming: userRepository.java
  api:
    file_path: api.tmpl
    output_file_naming: user.js
  service:
    file_path: service.tmpl
    output_file_naming: UserService.java
  controller:
    file_path: controller.tmpl
    output_file_naming: UserController.java
  entity:
    file_path: entity.tmpl
    output_file_naming: User.java
  search:
    file_path: search.tmpl
    output_file_naming: UserSearch.java


```

2. data.yaml

模板中所依赖的数据

```yaml
data:
  format: mongo
  # 提供搜索功能的字段
  search_column:
    - prop: "name"
      label: "名称"
      type: "input"
    - prop: "type"
      label: "类型"
      type: "select"
  #创建时所需的字段
  create_column:
    - prop:  "name"
      type: "String"
      label: "名称"
  #列表展示的字段
  table_column:
    - prop: "name"
      label: "名称"
      type: "String"
      #数据库中的名称
      column_name: "name"
  #除列表展示外数据库要保存的字段
  extra_column:
    - prop: "createDate"
      type: "Date"
      label: "创建日期"
      column_name: "create_date"
  #按钮及表单名
  button_name: "用户"
  #生成的英文名
  name: "User"
  package: "com.example"

```

3. *.tmpl

下面展示了生成vue的模板，其余模板请参照templates/example

```javascript
{{$Data := .Data}}
{{$pre:= "{{"}}
{{$end:= "}}"}}
<template>
    <d2-container>
        <template slot="header">
            <el-form :inline="true" :model="listQuery" ref="searchForm" size="mini" style="margin-bottom: -18px;">
            {{ range $Data.SearchColumn }}
                <el-form-item label="{{.Label}}" prop="{{.Prop}}">
                    {{if eq .Type "input"}}
                    <el-input v-model="listQuery.{{.Prop}}" placeholder="{{.Label}}" style="width: 100px;"/>
                    {{else if eq .Type "select"}}
                    <el-select v-model="listQuery.{{.Prop}}" filterable placeholder="请选择{{.Label}}">
                            <el-option label="0" value="0"></el-option>
                     </el-select>
                    {{end}}
                </el-form-item>
             {{- end }}
                <el-form-item>
                    <el-button type="primary" @click="handleFilter">
                        <d2-icon name="search" /> 查询{{$Data.ButtonName}}
                    </el-button>
                </el-form-item>
                <el-form-item>
                    <el-button @click="create{{$Data.Name}}">
                        <d2-icon name="create" /> 创建{{$Data.ButtonName}}
                    </el-button>
                </el-form-item>
            </el-form>
        </template>

        <el-table
                v-loading="listLoading"
                :data="list"
                border
                fit
                highlight-current-row
                style="width: 100%;margin-top: 20px"
        >
            <el-table-column type="index" align="center" width="50">
            </el-table-column>
            {{ range $Data.TableColumn }}
                 <el-table-column label="{{.Label}}" prop="{{.Prop}}" align="center" width="200">
                            <template slot-scope="scope">
                                <span>{{ $pre }}scope.row.{{.Prop}}{{$end}}</span>
                            </template>
                  </el-table-column>
            {{- end }}

            <el-table-column label="操作" align="center" fixed="right" width="400" class-name="small-padding fixed-width">
                <template slot-scope="scope">
                    <el-button type="primary" size="mini" @click="edit{{$Data.Name}}(scope.row)">编辑{{$Data.ButtonName}}</el-button>
                    <el-button type="danger" size="mini" @click="delete{{$Data.Name}}(scope.row)">删除{{$Data.ButtonName}}</el-button>

                </template>
            </el-table-column>
        </el-table>
        <template slot="footer">
            <el-pagination
                    :current-page="page.current"
                    :page-size="page.size"
                    :total="page.total"
                    :page-sizes="[1,100, 200, 300, 400]"
                    layout="total, sizes, prev, pager, next, jumper"
                    style="margin: -10px;"
                    @size-change="handleSizeChange"
                    @current-change="handleCurrentChange"
            >
            </el-pagination>
        </template>

        <el-dialog title="{{$Data.ButtonName}}管理" :visible.sync="dialogFormVisible" :close-on-click-modal=false>
            <el-form :model="temp" status-icon :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
                 {{ range $Data.CreateColumn }}
                    <el-form-item label="{{.Label}}" prop="{{.Prop}}">
                         {{if eq .Type "input"}}
                         <el-input v-model="temp.{{.Prop}}" placeholder="{{.Label}}" style="width: 100px;"/>
                         {{else if eq .Type "select"}}
                         <el-select v-model="temp.{{.Prop}}" filterable placeholder="请选择{{.Label}}">
                                 <el-option label="0" value="0"></el-option>
                          </el-select>
                         {{end}}
                   </el-form-item>                 
                   {{- end }}             
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button type="primary" @click="save{{$Data.Name}}('ruleForm')">提交</el-button>
                <el-button @click="reset{{$Data.Name}}Form('ruleForm')">取消</el-button>
            </div>
        </el-dialog>

    </d2-container>
</template>

<script>
   
    import {findPage, create{{$Data.Name}}, modify{{$Data.Name}}, delete{{$Data.Name}}} from '@/api/{{$Data.Name}}'
    export default {
        name: '{{$Data.Name}}',      
        data() {
            return {
                page: {
                    current: 1,
                    size: 100,
                    total: 0
                },
                temp: {
                    id: null,
                     {{ range $Data.CreateColumn }}
                     {{.Prop}}:null,
                     {{- end}}

                },
                list: null,
                total: 0,
                listLoading: true,
                listQuery: {
                    page: 1,
                    limit: 5,
                    {{ range $Data.SearchColumn }}
                     {{.Prop}}:null,
                    {{- end}}
                },
                dialogFormVisible: false,
                editFlag: 0,
                rules: {
                {{ range $Data.CreateColumn }}
                {{.Prop}}: [
                {required: true, message: '请输入{{.Label}}', trigger: 'blur'}
                                    ],
                 {{- end}}

                }
            }
        },
        created() {
            this.getList()
        },
        methods: {
            handleSizeChange(val) {
                this.page.size = val;
                this.getList();
            },
            handleCurrentChange(val) {
                this.page.current = val;
                this.getList();
            },
            handleFilter() {
                this.listQuery.page = 1
                this.getList()
            },
            save{{$Data.Name}}(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        if (this.editFlag === 0) {
                            create{{$Data.Name}}(
                                this.temp
                            ).then(response => {
                                this.$message({
                                    message: '保存成功',
                                    type: 'success'
                                });
                                this.getList();
                                this.reset{{$Data.Name}}Form('ruleForm')
                            })
                        } else {
                            modify{{$Data.Name}}(
                                this.temp
                            ).then(response => {
                                this.$message({
                                    message: '保存成功',
                                    type: 'success'
                                });
                                this.getList();
                                this.reset{{$Data.Name}}Form('ruleForm')
                            })

                        }


                    } else {
                        this.$message.error('保存失败,请重试');
                        return false;
                    }
                });
            },
            reset{{$Data.Name}}Form(formName) {
                this.$refs[formName].resetFields();
                this.dialogFormVisible = false;
            },
            getList() {
                this.listLoading = true
                this.listQuery.page=this.page.current
                this.listQuery.limit=this.page.size
                findPage(this.listQuery).then(response => {
                    this.list = response.content;
                    this.page.total = response.totalElements;
                    this.listLoading = false
                })
            },
            create{{$Data.Name}}() {
                this.resetTemp()
                this.dialogFormVisible = true;
                this.editFlag = 0;
            },
            edit{{$Data.Name}}(item) {
                this.resetTemp()
                this.temp.id = item.id;
                this.temp$Data.Name = item$Data.Name;
                this.temp.rpId=item.rpId
                this.dialogFormVisible = true;
                this.editFlag = 1;
            },

            resetTemp() {
                this.temp = {
                    id: null,
                    name: null,
                    rpId:null
                }
            },
            delete{{$Data.Name}}(row) {
                this.$confirm('此操作将永久删除该{{$Data.ButtonName}}, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    delete{{$Data.Name}}(row.id).then(response => {
                        this.getList();
                        this.$message({
                            type: 'success',
                            message: '删除成功!'
                        });
                    })

                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消删除'
                    });
                });
            }


        }

    }
</script>

<style scoped>

</style>
```

4. 使用实例

```bash
./generatorv1 -f ./templates/example/ -o ./templates/example/test/ -t ./templates/example/

```
