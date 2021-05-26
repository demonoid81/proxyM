<template>
    <div style="display: flex; flex-direction: column">
        <Button type="primary" style="width: 150px" @click="modal = true">Add Proxy</Button>
        <Modal
                v-model="modal"
                title="Add proxy"
                ok-text="OK"
                cancel-text="Cancel"
                @on-ok="ok"
                @on-cancel="cancel">
            <Form ref="formCustom" :model="formItem" :label-width="100">
                <FormItem label="name">
                    <Input v-model="formItem.name" placeholder="Enter something..."></Input>
                </FormItem>
                <FormItem label="region">
                    <Select v-model="formItem.region">
                        <Option value="ch">ch</Option>
                        <Option value="ru">ru</Option>
                        <Option value="en">en</Option>
                        <Option value="all">all</Option>
                    </Select>
                </FormItem>
            </Form>
        </Modal>
        <div style="height: 10px"></div>
        <Table border :columns="columns" :data="data">
            <template slot-scope="{ row }" slot="name">
                <strong>{{ row.name }}</strong>
            </template>
            <template slot-scope="{ row }" slot="stat">
                <div v-if="row.status === 'online'" >
                    <Tag color="success">OnLine</Tag>
                </div>
                <div v-else>
                    <Tag color="error">Error</Tag>
                </div>
            </template>
            <template slot-scope="{ row, index }" slot="action">
                <Button type="error" size="small" @click="remove(index)">Delete</Button>
            </template>
        </Table>
    </div>
</template>
<script>
  export default {
    data () {
      return {
        modal: false,
        formItem: {
          name: '',
          region: ''
        },
        columns: [
          {
            title: 'Name',
            slot: 'name'
          },
          {
            title: 'Region',
            key: 'region'
          },
          {
            title: 'Status',
            slot: 'stat',
            width: 150,
            align: 'center'
          },
          {
            title: 'Action',
            slot: 'action',
            width: 150,
            align: 'center'
          }
        ],
        data: [
          {
            name: 'ch proxy',
            region: 'ch',
            status: 'online'
          },
          {
            name: 'ch proxy 2',
            region: 'ch',
            status: 'online'
          },
          {
            name: 'defailt',
            region: 'all',
            status: 'error'
          },
          {
            name: 'reserved',
            region: 'zh',
            status: 'online'
          }
        ]
      }
    },
    methods: {
      ok () {
        this.data.push({
          name: this.formItem.name,
          region:  this.formItem.region
        })
        this.$refs["formCustom"].resetFields()
        this.modal = false
      },
      cancel () {
        this.modal = false
      },
      remove (index) {
        this.data.splice(index, 1);
      }
    }
  }
</script>
